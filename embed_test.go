package embed_test

import (
	"bytes"
	"testing"

	embed "github.com/PaperPrototype/goldmark-embed"
	"github.com/yuin/goldmark"
)

func TestMeta(t *testing.T) {
	markdown := goldmark.New(
		goldmark.WithExtensions(
			embed.New(),
		),
	)
	source := `# Hello goldmark-embed for Youtube

![](https://www.youtube.com/watch?v=dQw4w9WgXcQ)
`
	var buf bytes.Buffer
	if err := markdown.Convert([]byte(source), &buf); err != nil {
		panic(err)
	}
	if buf.String() != `<h1>Hello goldmark-embed for Youtube</h1>
<p><iframe width="560" height="315" src="https://www.youtube.com/embed/dQw4w9WgXcQ" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture" allowfullscreen></iframe></p>
` {
		t.Error("Invalid HTML output for YouTube")
		t.Log(buf.String())
	}

	source2 := `# Hello goldmark-embed for Vimeo

![](https://vimeo.com/148751763)
`

	var buf2 bytes.Buffer
	if err2 := markdown.Convert([]byte(source2), &buf2); err2 != nil {
		panic(err2)
	}
	if buf2.String() != `<h1>Hello goldmark-embed for Vimeo</h1>
<p><iframe src="https://player.vimeo.com/video/148751763?&amp;badge=0&amp;autopause=0&amp;player_id=0&amp;app_id=58479" width="724" height="404" frameborder="0" allow="autoplay; fullscreen; picture-in-picture" allowfullscreen></iframe></p>
` {
		t.Error("Invalid HTML output for Vimeo")
		t.Log(buf2.String())
	}
}
