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

var htmlTemplate *template.Template

type HTMLOpts struct {
	Styles   map[string]string
	Wrappers map[string]struct{ Begin, End string }
}

func renderHook(w io.Writer, node ast.Node, entering bool) (ast.WalkStatus, bool) {
	if _, ok := node.(*ast.Heading); ok {
		level := strconv.Itoa(node.(*ast.Heading).Level)

		var err error
		if entering && level == "1" {
			_, err = w.Write([]byte(`<h1 class="title is-1 has-text-centered">`))
		} else if entering {
			_, err = w.Write([]byte("<h" + level + ">"))
		} else {
			_, err = w.Write([]byte("</h" + level + ">"))
		}
		if err != nil {
			return ast.Terminate, true
		}

		return ast.GoToNext, true
	}

	if _, ok := node.(*ast.Image); ok {
		src := string(node.(*ast.Image).Destination)

		c := node.(*ast.Image).GetChildren()[0]
		alt := string(c.AsLeaf().Literal)

		var err error
		if entering && alt != "" {
			_, err = w.Write([]byte(`<figure class="image is-5by3"><img src="` + src + `" alt="` + alt + `">`))
		} else if entering {
			_, err = w.Write([]byte(`<figure class="image is-5by3"><img src="` + src + `">`))
		} else {
			_, err = w.Write([]byte(`</figure>`))
		}
		if err != nil {
			return ast.Terminate, true
		}

		return ast.SkipChildren, true
	}

	return ast.GoToNext, false
}

type MDToHTMLOpts struct {
	Style string
	Title string
}

func (r Renderer) MDToHTML(md []byte, opts MDToHTMLOpts) []byte {
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock | parser.SuperSubscript
	p := parser.NewWithExtensions(extensions)
	doc := p.Parse(md)

	renderer := html.NewRenderer(html.RendererOptions{
		Flags:          html.CommonFlags | html.HrefTargetBlank,
		RenderNodeHook: renderHook,
	})

	content := markdown.Render(doc, renderer)

	vars := struct {
		Style        string
		Title        string
		WrapperBegin template.HTML
		Content      template.HTML
		WrapperEnd   template.HTML
	}{
		Style:        r.cfg.HTML.Styles[opts.Style],
		Title:        opts.Title,
		Content:      template.HTML(content),
		WrapperBegin: template.HTML(r.cfg.HTML.Wrappers[opts.Style].Begin),
		WrapperEnd:   template.HTML(r.cfg.HTML.Wrappers[opts.Style].End),
	}

	var processed bytes.Buffer
	_ = htmlTemplate.Execute(&processed, vars)

	return processed.Bytes()
}
