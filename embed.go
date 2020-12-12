package embed

import (
	"fmt"
	"net/url"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/text"
	"github.com/yuin/goldmark/util"
)

type astTransformer struct {
}

var defaultASTTransformer = &astTransformer{}

func (a *astTransformer) Transform(node *ast.Document, reader text.Reader, pc parser.Context) {
	replaceImages := func(n ast.Node, entering bool) (ast.WalkStatus, error) {
		if !entering {
			return ast.WalkContinue, nil
		}
		if n.Kind() != ast.KindImage {
			return ast.WalkContinue, nil
		}

		img := n.(*ast.Image)
		u, err := url.Parse(string(img.Destination))
		if err != nil {
			msg := ast.NewString([]byte(fmt.Sprintf("<!-- %s -->", err)))
			msg.SetCode(true)
			n.Parent().InsertAfter(n.Parent(), n, msg)
			return ast.WalkContinue, nil
		}

		if u.Host != "www.youtube.com" || u.Path != "/watch" {
			return ast.WalkContinue, nil
		}
		v := u.Query().Get("v")
		if v == "" {
			return ast.WalkContinue, nil
		}
		yt := NewYouTube(img, v)
		n.Parent().ReplaceChild(n.Parent(), n, yt)

		return ast.WalkContinue, nil
	}

	ast.Walk(node, replaceImages)

}

// Option is a functional option type for this extension.
type Option func(*embed)

type embed struct {
}

// Embed is a extension for goldmark.
var Embed = &embed{}

// New returns a new Embed extension.
func New(opts ...Option) goldmark.Extender {
	e := &embed{}
	for _, opt := range opts {
		opt(e)
	}
	return e
}

func (e *embed) Extend(m goldmark.Markdown) {
	m.Parser().AddOptions(
		parser.WithASTTransformers(
			util.Prioritized(defaultASTTransformer, 500),
		),
	)
	m.Renderer().AddOptions(
		renderer.WithNodeRenderers(
			util.Prioritized(NewHTMLRenderer(), 500),
		),
	)
}

// YouTube struct represents a YouTube Video embed of the Markdown text.
type YouTube struct {
	ast.Image
	Video string
}

// KindYouTube is a NodeKind of the YouTube node.
var KindYouTube = ast.NewNodeKind("YouTube")

// Kind implements Node.Kind.
func (n *YouTube) Kind() ast.NodeKind {
	return KindYouTube
}

// NewYouTube returns a new YouTube node.
func NewYouTube(img *ast.Image, v string) *YouTube {
	c := &YouTube{
		Image: *img,
		Video: v,
	}
	c.Destination = img.Destination
	c.Title = img.Title

	return c
}

// HTMLRenderer struct is a renderer.NodeRenderer implementation for the extension.
type HTMLRenderer struct{}

// NewHTMLRenderer builds a new HTMLRenderer with given options and returns it.
func NewHTMLRenderer() renderer.NodeRenderer {
	r := &HTMLRenderer{}
	return r
}

// RegisterFuncs implements NodeRenderer.RegisterFuncs.
func (r *HTMLRenderer) RegisterFuncs(reg renderer.NodeRendererFuncRegisterer) {
	reg.Register(KindYouTube, r.renderYouTubeVideo)
}

func (r *HTMLRenderer) renderYouTubeVideo(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	if entering {
		return ast.WalkContinue, nil
	}

	yt := node.(*YouTube)

	w.Write([]byte(`<iframe width="560" height="315" src="https://www.youtube.com/embed/` + yt.Video + `" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture" allowfullscreen></iframe>`))
	return ast.WalkContinue, nil
}
