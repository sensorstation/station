package station

import "io"

// Move this in from stub
type Publisher struct {
	Path string
	io.Reader
}
