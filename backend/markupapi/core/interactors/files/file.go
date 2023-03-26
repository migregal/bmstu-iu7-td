package files

import (
	"fmt"
	"markup2/markupapi/core/interactors"
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
	cfg Config
}

func New(cfg Config) (Interactor, error) {
	renderer, err := render.New(render.Config{HTML: render.HTMLOpts(cfg)})
	if err != nil {
		return Interactor{}, fmt.Errorf("failed to init renderer: %w", err)
	}

	renderers = map[string]func([]byte, string) []byte{
		"md":    func(d []byte, _ string) []byte { return d },
		"html":  renderer.MDToHTML,
		"plain": renderer.MDToPlain,
	}

	return Interactor{cfg: cfg}, nil
}

var formats = map[string]struct{}{
	"md":    {},
	"html":  {},
	"plain": {},
}

var renderer render.Renderer

var renderers map[string]func([]byte, string) []byte

type Opts struct {
	Format string
	Style  string
}

func (i *Interactor) Get(opts Opts) ([]byte, error) {
	if _, found := formats[opts.Format]; !found {
		return nil, interactors.ErrNotFound
	}

	data := []byte(`
# Hello, world

![image](https://thumb.tildacdn.com/tild6465-6132-4937-b964-336163313261/-/format/webp/mem-2-1024x683.jpg)
`)

	return renderers[opts.Format](data, opts.Style), nil
}
