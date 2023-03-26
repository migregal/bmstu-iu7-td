package files

import (
	"markup2/markupapi/core/interactors"
	"markup2/pkg/render"
)

type Interactor struct {
}

func New() Interactor {
	return Interactor{}
}

var formats = map[string]struct{}{
	"md":    {},
	"html":  {},
	"plain": {},
}

var renderers = map[string]func([]byte, string) []byte{
	"md": func(d []byte, _ string) []byte { return d },
	"html": render.MDToHTML,
	"plain": render.MDToPlain,
}

type Opts struct {
	Format string
}

func (i *Interactor) Get(opts Opts) ([]byte, error) {
	if _, found := formats[opts.Format]; !found {
		return nil, interactors.ErrNotFound
	}

	data := []byte(`
# Hello, world

![image](https://thumb.tildacdn.com/tild6465-6132-4937-b964-336163313261/-/format/webp/mem-2-1024x683.jpg)
`)

	return renderers[opts.Format](data, "bulma"), nil
}
