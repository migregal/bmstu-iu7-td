package files

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"markup2/markupapi/core/ports/repositories"

	"github.com/labstack/gommon/log"
	"github.com/labstack/gommon/random"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository struct {
	cfg repositories.FilesConfig
}

func New(cfg repositories.FilesConfig) (*Repository, error) {
	return &Repository{cfg: cfg}, nil
}

func (r *Repository) newConn(ctx context.Context) (*mongo.Client, func(), error) {
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%d",
		r.cfg.User, r.cfg.Passsword, r.cfg.Host, r.cfg.Port)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, nil, fmt.Errorf("failed to connect to db: %w", err)
	}

	cancel := func() {
		if err = client.Disconnect(ctx); err != nil {
			log.Errorf("failed to disconnect from db: %v", err)
		}
	}

	return client, cancel, nil
}

func (r *Repository) Get(ctx context.Context, id string) (io.Reader, string, error) {
	client, cancel, err := r.newConn(ctx)
	if err != nil {
		return nil, "", fmt.Errorf("failed to connect to db: %w", err)
	}
	defer cancel()

	db := client.Database(r.cfg.Name)
	bucket, err := gridfs.NewBucket(db)
	if err != nil {
		return nil, "", fmt.Errorf("failed to connect to db: %w", err)
	}

	fileID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, "", fmt.Errorf("invalid file id: %w", err)
	}

	// extract exact file title from metadata of gridfs
	filter := bson.D{{Key: "_id", Value: bson.D{{Key: "$eq", Value: fileID}}}}
	cursor, err := bucket.Find(filter)
	if err != nil {
		return nil, "", fmt.Errorf("failed to find file metadata: %w", err)
	}
	type gridfsFile struct {
		ID   string `bson:"_id"`
		Meta bson.D `bson:"metadata"`
	}
	var foundFiles []gridfsFile
	if err = cursor.All(context.TODO(), &foundFiles); err != nil {
		return nil, "", fmt.Errorf("failed to retrieve file metadata: %w", err)
	}
	if len(foundFiles) == 0 {
		return nil, "", fmt.Errorf("failed to find file by id: %w", err)
	}

	fileBuffer := bytes.NewBuffer(nil)
	if _, err := bucket.DownloadToStream(fileID, fileBuffer); err != nil {
		return nil, "", fmt.Errorf("failed to get file from db: %w", err)
	}

	title := foundFiles[0].Meta.Map()["title"].(string)

	return fileBuffer, title, nil
}

func (r *Repository) Add(ctx context.Context, title string, content io.Reader) (string, error) {
	client, cancel, err := r.newConn(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to connect to db: %w", err)
	}
	defer cancel()

	db := client.Database(r.cfg.Name)
	bucket, err := gridfs.NewBucket(db)
	if err != nil {
		return "", fmt.Errorf("failed to connect to db: %w", err)
	}

	uploadOpts := options.GridFSUpload().SetMetadata(bson.D{{Key: "title", Value: title}})
	objectID, err := bucket.UploadFromStream(random.String(64), content, uploadOpts)

	return objectID.Hex(), nil
}
