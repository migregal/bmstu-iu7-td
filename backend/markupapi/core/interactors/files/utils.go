package files

import (
	"io"
	"strconv"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/ast"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"

	stripmd "github.com/writeas/go-strip-markdown"
)

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

func mdToHTML(md []byte) []byte {
	htmlHeader := `
<html lang="en">
<head>
   <meta charset="UTF-8">
   <meta http-equiv="X-UA-Compatible" content="IE=edge">
   <meta name="viewport" content="width=device-width, initial-scale=1.0">
   <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bulma@0.9.3/css/bulma.min.css">
   <title>Blog Post Example</title>
</head>
<body>
<br>
<div class="container is-max-desktop">
<div class="content">`

	htmlFooter := `
</div>
</div>
<br>
</body>
</html>`

	// create markdown parser with extensions
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock | parser.SuperSubscript
	p := parser.NewWithExtensions(extensions)
	doc := p.Parse(md)

	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{
		Flags: htmlFlags,
		// CSS:   "https://unpkg.com/@picocss/pico@1.*/css/pico.min.css",
		RenderNodeHook: renderHook,
	}
	renderer := html.NewRenderer(opts)

	content := markdown.Render(doc, renderer)
	return append([]byte(htmlHeader), append(content, []byte(htmlFooter)...)...)
}

func mdToPlain(data []byte) []byte {
	return []byte(stripmd.Strip(string(data)))
}
