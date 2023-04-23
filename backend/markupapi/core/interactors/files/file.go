package files

import (
	"context"
	"errors"
	"fmt"
	"io"

	"markup2/markupapi/core/interactors"
	"markup2/markupapi/core/ports/repositories"
	"markup2/pkg/render"
)

type Config struct {
	Styles   map[string]string
	Wrappers map[string]struct {
		Begin string
		End   string
	}
}
type Interactor struct {
	cfg  Config
	repo repositories.FilesRepo
}

func New(cfg Config, repo repositories.FilesRepo) (Interactor, error) {
	renderer, err := render.New(render.Config{HTML: render.HTMLOpts(cfg)})
	if err != nil {
		return Interactor{}, fmt.Errorf("failed to init renderer: %w", err)
	}

	renderers = map[string]func([]byte, any) []byte{
		"md": func(d []byte, _ any) []byte { return d },
		"html": func(d []byte, opts any) []byte {
			o := opts.(render.MDToHTMLOpts)
			return renderer.MDToHTML(d, o)
		},
		"plain": func(d []byte, _ any) []byte {
			return renderer.MDToPlain(d, render.MDToPlainOpts{})
		},
	}

	return Interactor{cfg: cfg, repo: repo}, nil
}

var formats = map[string]struct{}{
	"md":    {},
	"html":  {},
	"plain": {},
}

var renderers map[string]func([]byte, any) []byte

type GetOpts struct {
	Format string
	Style  string
}

func (i *Interactor) Get(ctx context.Context, id string, opts GetOpts) ([]byte, error) {
	if _, found := formats[opts.Format]; !found {
		return nil, interactors.ErrNotFound
	}

	reader, title, err := i.repo.Get(ctx, id)
	if err != nil {
		if errors.Is(err, repositories.ErrNotFound) {
			err = interactors.ErrNotFound
		}

		return nil, fmt.Errorf("failed to get file from db: %w", err)
	}

	data, err := io.ReadAll(reader)
	if err != nil {
		return nil, fmt.Errorf("failed to read file contents: %w", err)
	}

	var renderOpts any
	switch opts.Format {
	case "html":
		if opts.Style == "" {
			for k := range i.cfg.Styles {
				opts.Style = k
				break
			}
		}
		renderOpts = render.MDToHTMLOpts{
			Style: opts.Style,
			Title: title,
		}
	case "plain":
		renderOpts = render.MDToPlainOpts{}
	default:
		renderOpts = nil
	}

	return renderers[opts.Format](data, renderOpts), nil
}

func (i *Interactor) Find(ctx context.Context, ownerID uint64) ([]File, error) {
	repoFiles, err := i.repo.Find(ctx, ownerID)
	if err != nil {
		return nil, fmt.Errorf("failed to get files info from db: %w", err)
	}

	files := make([]File, 0, len(repoFiles))
	for _, file := range repoFiles {
		files = append(files, File(file))
	}

	return files, nil
}

func (i *Interactor) Add(ctx context.Context, owner uint64, title string, in io.Reader) (string, error) {
	return i.repo.Add(ctx, owner, title, in)
}

func (i *Interactor) Update(ctx context.Context, owner uint64, id string, title string, in io.Reader) (string, error) {
	return i.repo.Update(ctx, owner, title, id, in)
}

func (i *Interactor) Delete(ctx context.Context, ownerID uint64, id string) error {
	err := i.repo.Delete(ctx, ownerID, id)
	if err != nil {
		if errors.Is(err, repositories.ErrNotFound) {
			return interactors.ErrNotFound
		}

		return fmt.Errorf("failed to delete file from repo: %w", err)
	}

	return nil
}
