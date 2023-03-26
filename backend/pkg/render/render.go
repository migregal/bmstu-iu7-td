package render

import (
	"fmt"
	"html/template"
)

type Config struct {
	HTML HTMLOpts
}

type Renderer struct {
	cfg Config
}

func New(cfg Config) (Renderer, error) {
	t, err := template.New("html-tmpl").Parse(htmlHeaderTemplate)
	if err != nil {
		return Renderer{}, fmt.Errorf("failed to parse templates: %w", err)
	}

	htmlTemplate = t

	return Renderer{cfg: cfg}, nil
}
