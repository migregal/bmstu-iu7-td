package files

import "go.mongodb.org/mongo-driver/bson"

type gridfsFile struct {
	ID     string `bson:"_id"`
	Meta   bson.D `bson:"metadata"`
	Length int64  `bson:"length"`
}
