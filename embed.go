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

// Option is a functional option type for this extension.
type Option func(*embedExtension)

type embedExtension struct{}

// New returns a new Embed extension.
func New(opts ...Option) goldmark.Extender {
	e := &embedExtension{}
	for _, opt := range opts {
		opt(e)
	}
	return e
}

func (e *embedExtension) Extend(m goldmark.Markdown) {
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

type astTransformer struct{}

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
		// parse the url
		u, err := url.Parse(string(img.Destination))
		if err != nil {
			msg := ast.NewString([]byte(fmt.Sprintf("<!-- %s -->", err)))
			msg.SetCode(true)
			n.Parent().InsertAfter(n.Parent(), n, msg)
			return ast.WalkContinue, nil
		}

		if u.Host == "www.youtube.com" && u.Path == "/watch" {
			// if YouTube

			v := u.Query().Get("v")
			if v == "" {
				return ast.WalkContinue, nil
			}
			youtube := NewYouTube(img, v)
			n.Parent().ReplaceChild(n.Parent(), n, youtube)

		} else if u.Host == "vimeo.com" {
			// if Vimeo

			// remove the '/' from url
			path := u.Path[1:]
			vimeo := NewVimeo(img, path)
			n.Parent().ReplaceChild(n.Parent(), n, vimeo)
		}

		return ast.WalkContinue, nil
	}

	ast.Walk(node, replaceImages)
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
	reg.Register(KindVimeo, r.renderVimeoVideo)
}

func (r *HTMLRenderer) renderYouTubeVideo(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	if entering {
		return ast.WalkContinue, nil
	}

	youtube := node.(*YouTube)

	w.Write([]byte(`<iframe width="560" height="315" src="https://www.youtube.com/embed/` + youtube.Video + `" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture" allowfullscreen></iframe>`))
	return ast.WalkContinue, nil
}

func (r *HTMLRenderer) renderVimeoVideo(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	if entering {
		return ast.WalkContinue, nil
	}

	vimeo := node.(*Vimeo)

	w.Write([]byte(`<iframe src="https://player.vimeo.com/video/` + vimeo.Video + `?&amp;badge=0&amp;autopause=0&amp;player_id=0&amp;app_id=58479" width="724" height="404" frameborder="0" allow="autoplay; fullscreen; picture-in-picture" allowfullscreen></iframe>`))
	return ast.WalkContinue, nil
}
