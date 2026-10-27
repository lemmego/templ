# lemmego/templ

Renders [templ](https://templ.guide) components from a Lemmego handler. templ
compiles your templates to Go, so a missing field or a renamed struct is a
build error rather than a blank space on the page.

Two functions, because that is all this needs to be.

```go
import (
    lemmetempl "github.com/lemmego/templ"
    "myapp/resources/views"
)

func Show(c app.Context) error {
    return lemmetempl.Respond(c, views.PostShow(post))
}
```

## Install

```bash
go get github.com/lemmego/templ
```

You need the `templ` binary to compile your templates. A project from
`lemmego new --frontend templ` has it in `make deps`, and `lemmego dev`
watches and recompiles them.

```bash
go install github.com/a-h/templ/cmd/templ@latest
templ generate
```

## Responding

`Respond` writes the component to the response with `text/html`:

```go
return lemmetempl.Respond(c, views.Layout(views.PostShow(post)))
```

Composition is ordinary Go, so a layout is a component that takes a component.

## Status and headers

`Respond` sets `content-type: text/html` and, if nothing has set a status
yet, `200`. To send anything else, set it first:

```go
c.SetStatus(http.StatusNotFound)
return lemmetempl.Respond(c, views.NotFound())
```

`New` returns the response object if you want to hold it before sending, but
its `Render` always writes to the request's `ResponseWriter` — it does not
render into a buffer. To render a component somewhere else, such as an email
body, call templ directly:

```go
var buf bytes.Buffer
err := views.PostShow(post).Render(ctx, &buf)
```

## Choosing between this and the alternatives

| | Reach for it when |
|---|---|
| **templ** | you want server-rendered HTML with compile-time checking |
| Go templates | you want server-rendered HTML with no build step |
| [Inertia](https://github.com/lemmego/inertia) | you want React or Vue with routing and data loading still in Go |

templ and Inertia coexist in one application, and `lemmego new` offers the
combination. A page is one or the other.

## Testing

```bash
go test ./...
GOWORK=off go test ./...   # as a published module
```
