package render

import (
	stripmd "github.com/writeas/go-strip-markdown"
)

type MDToPlainOpts struct {

}

func (r Renderer) MDToPlain(data []byte, opts MDToPlainOpts) []byte {
	return []byte(stripmd.Strip(string(data)))
}
