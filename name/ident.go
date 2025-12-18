package name

import "github.com/riscdanger/flect"

// Ident represents the string and it's parts
type Ident struct {
	flect.Ident
}

// New creates a new Ident from the string
func New(s string) Ident {
	return Ident{flect.New(s)}
}
