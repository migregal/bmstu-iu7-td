package render

import (
	stripmd "github.com/writeas/go-strip-markdown"
)

func MDToPlain(data []byte, style string) []byte {
	return []byte(stripmd.Strip(string(data)))
}
