package files

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
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

var renderer render.Renderer

var renderers map[string]func([]byte, any) []byte

type Opts struct {
	Format string
	Style  string
}

func (i *Interactor) Get(ctx context.Context, id string, opts Opts) ([]byte, error) {
	if _, found := formats[opts.Format]; !found {
		return nil, interactors.ErrNotFound
	}

	reader, title, err := i.repo.Get(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get file from db: %w", err)
	}

	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, fmt.Errorf("failed to read file contents: %w", err)
	}

	var renderOpts any
	switch opts.Format {
	case "html":
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

func (i *Interactor) Add(ctx context.Context, title string, in io.Reader) (string, error) {
	return i.repo.Add(ctx, title, in)
}
