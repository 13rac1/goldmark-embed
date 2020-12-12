package embed_test

import (
	"bytes"
	"testing"

	embed "github.com/13rac1/goldmark-embed"
	"github.com/yuin/goldmark"
)

func TestMeta(t *testing.T) {
	markdown := goldmark.New(
		goldmark.WithExtensions(
			embed.Embed,
		),
	)
	source := `# Hello goldmark-embed

![](https://www.youtube.com/watch?v=dQw4w9WgXcQ)
`
	var buf bytes.Buffer
	if err := markdown.Convert([]byte(source), &buf); err != nil {
		panic(err)
	}
	if buf.String() != `<h1>Hello goldmark-embed</h1>
<p><iframe width="560" height="315" src="https://www.youtube.com/embed/dQw4w9WgXcQ" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture" allowfullscreen></iframe></p>
` {
		t.Error("Invalid HTML output")
		t.Log(buf.String())
	}
}
