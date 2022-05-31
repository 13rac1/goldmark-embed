package embed

import (
	"github.com/yuin/goldmark/ast"
)

// YouTube struct represents a YouTube Video embed of the Markdown text.
type YouTube struct {
	ast.Image
	Video string
}

// KindYouTube is a NodeKind of the YouTube node.
var KindYouTube = ast.NewNodeKind("YouTube")

// Kind implements Node.Kind.
func (node *YouTube) Kind() ast.NodeKind {
	return KindYouTube
}

// NewYouTube returns a new YouTube node.
func NewYouTube(img *ast.Image, video string) *YouTube {
	vimeo := &YouTube{
		Image: *img,
		Video: video,
	}
	vimeo.Destination = img.Destination
	vimeo.Title = img.Title

	return vimeo
}
