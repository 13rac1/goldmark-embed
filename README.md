# goldmark-embed

goldmark-embed is an extension for the [goldmark][goldmark] library that extends
the Markdown `![]()` image embed syntax to support additional media formats.

[goldmark]: http://github.com/yuin/goldmark

YouTube only at first.

## Demo

This markdown:

```md
# Hello goldmark-embed

![](https://www.youtube.com/watch?v=dQw4w9WgXcQ)
```

Becomes this HTML:

```html
<h1>Hello goldmark-embed</h1>
<p><iframe width="560" height="315" src="https://www.youtube.com/embed/dQw4w9WgXcQ" frameborder="0"
allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture"
allowfullscreen></iframe></p>
```

### Installation

```bash
go get github.com/13rac1/goldmark-embed
```

## Usage

```go
  markdown := goldmark.New(
    goldmark.WithExtensions(
      embed.Embed,
    ),
  )
  var buf bytes.Buffer
  if err := markdown.Convert([]byte(source), &buf); err != nil {
    panic(err)
  }
  fmt.Print(buf)
}
```

## TODO

* Embed Options
* Additional Data Sources

## License

MIT

## Author

Brad Erickson
