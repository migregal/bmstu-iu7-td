package repositories

import (
	"context"
	"io"
)

type FilesConfig struct {
	Host      string
	Port      int
	User      string
	Passsword string
	Name      string
}

type File struct {
	ID     string
	Title  string
	Length int64
}

type FilesRepo interface {
	Get(ctx context.Context, id string) (io.Reader, string, error)
	Find(ctx context.Context, ownerID uint64) ([]File, error)
	Add(ctx context.Context, ownerID uint64, title string, content io.Reader) (string, error)
	Update(ctx context.Context, ownerID uint64, id string, title string, content io.Reader) (string, error)
	Delete(ctx context.Context, ownerID uint64, id string) error
}
