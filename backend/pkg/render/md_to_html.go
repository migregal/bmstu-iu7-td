package render

import (
	"bytes"
	"html/template"
	"io"
	"strconv"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/ast"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

var htmlHeaderTemplate = `
<html lang="en">
<head>
	<meta charset="UTF-8">
	<meta http-equiv="X-UA-Compatible" content="IE=edge">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	{{if .Style}}<link rel="stylesheet" href="{{.Style}}">{{end}}
	{{if .Title}}<title>{{.Title}}</title>{{end}}
</head>
<body>
{{if .WrapperBegin}}{{.WrapperBegin}}{{end}}
{{if .Content}}{{.Content}}{{end}}
{{if .WrapperEnd}}{{.WrapperEnd}}{{end}}
</body>
</html>
`

type HTMLOpts struct {
	Styles   map[string]string
	Wrappers map[string]struct{ Begin, End string }
}

func renderHook(w io.Writer, node ast.Node, entering bool) (ast.WalkStatus, bool) {
	if _, ok := node.(*ast.Heading); ok {
		level := strconv.Itoa(node.(*ast.Heading).Level)

		if entering && level == "1" {
			w.Write([]byte(`<h1 class="title is-1 has-text-centered">`))
		} else if entering {
			w.Write([]byte("<h" + level + ">"))
		} else {
			w.Write([]byte("</h" + level + ">"))
		}

		return ast.GoToNext, true
	}

	if _, ok := node.(*ast.Image); ok {
		src := string(node.(*ast.Image).Destination)

		c := node.(*ast.Image).GetChildren()[0]
		alt := string(c.AsLeaf().Literal)

		if entering && alt != "" {
			w.Write([]byte(`<figure class="image is-5by3"><img src="` + src + `" alt="` + alt + `">`))
		} else if entering {
			w.Write([]byte(`<figure class="image is-5by3"><img src="` + src + `">`))
		} else {
			w.Write([]byte(`</figure>`))
		}

		return ast.SkipChildren, true
	}

	return ast.GoToNext, false
}

func (r Renderer) MDToHTML(md []byte, style string) []byte {
	t, err := template.New("html-tmpl").Parse(htmlHeaderTemplate)
	if err != nil {
		panic(err)
	}

	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock | parser.SuperSubscript
	p := parser.NewWithExtensions(extensions)
	doc := p.Parse(md)

	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{
		Flags:          htmlFlags,
		RenderNodeHook: renderHook,
	}
	renderer := html.NewRenderer(opts)

	content := markdown.Render(doc, renderer)

	vars := struct {
		Style        string
		Title        string
		WrapperBegin template.HTML
		Content      template.HTML
		WrapperEnd   template.HTML
	}{
		Style:        r.cfg.HTML.Styles[style],
		Title:        "TODO",
		Content:      template.HTML(content),
		WrapperBegin: template.HTML(r.cfg.HTML.Wrappers[style].Begin),
		WrapperEnd:   template.HTML(r.cfg.HTML.Wrappers[style].End),
	}

	var processed bytes.Buffer
	t.Execute(&processed, vars)

	return processed.Bytes()
}
