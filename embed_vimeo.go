package embed

import (
	"github.com/yuin/goldmark/ast"
)

// Vimeo struct represents a Vimeo Video embed of the markdown text.
type Vimeo struct {
	ast.Image
	Video string
}

// KindVimeo is a NodeKind of the Vimeo node.
var KindVimeo = ast.NewNodeKind("Vimeo")

// implements Node.Kind.
func (node *Vimeo) Kind() ast.NodeKind {
	return KindVimeo
}

// New Vimeo returns a new Vimeo node.
func NewVimeo(img *ast.Image, video string) *Vimeo {
	vimeo := &Vimeo{
		Image: *img,
		Video: video,
	}
	vimeo.Destination = img.Destination
	vimeo.Title = img.Title

	return vimeo
}
