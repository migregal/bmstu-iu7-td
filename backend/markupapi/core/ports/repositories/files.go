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

type FilesRepo interface {
	Get(ctx context.Context, id string) (io.Reader, string, error)
	Add(ctx context.Context, title string, content io.Reader) (string, error)
}
