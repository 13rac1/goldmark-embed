# goldmark-embed

goldmark-embed is based on 13rac1's extension for the [goldmark][goldmark] library that extends
the Markdown `![]()` image embed syntax to support additional media formats.

[goldmark]: http://github.com/yuin/goldmark

Supports YouTube and Vimeo links.

## Demo

This markdown:

```md
# Hello goldmark-embed for YouTube

![](https://www.youtube.com/watch?v=dQw4w9WgXcQ)

# Hello goldmark-embed for Vimeo

![](https://vimeo.com/148751763)
```

Becomes this HTML:

```html
<h1>Hello goldmark-embed for Youtube</h1>
<p><iframe width="560" height="315" src="https://www.youtube.com/embed/dQw4w9WgXcQ" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture" allowfullscreen></iframe></p>
<h1>Hello goldmark-embed for Vimeo</h1>
<p><iframe src="https://player.vimeo.com/video/148751763?&amp;badge=0&amp;autopause=0&amp;player_id=0&amp;app_id=58479" width="724" height="404" frameborder="0" allow="autoplay; fullscreen; picture-in-picture" allowfullscreen></iframe></p>
```

### Installation

```bash
go get github.com/PaperPrototype/goldmark-embed
```

## Usage

```go
  markdown := goldmark.New(
    goldmark.WithExtensions(
      embed.New(),
    ),
  )
  var buf bytes.Buffer
  if err := markdown.Convert([]byte(source), &buf); err != nil {
    panic(err)
  }
  
  // output html
  fmt.Print(buf)
}
```

## TODO

* Embed Options
* Additional Data Sources

## License

MIT

## Authors

![Brad Erickson](https://github.com/13rac1)
![Abdiel Lopez](https://github.com/PaperPrototype)
