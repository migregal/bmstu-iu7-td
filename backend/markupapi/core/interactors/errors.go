package interactors

import "fmt"

var (
	ErrNotFound = fmt.Errorf("not found")
	ErrExists   = fmt.Errorf("exists")
)
