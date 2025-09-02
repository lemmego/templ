package templ

import (
	"github.com/a-h/templ"
	"github.com/lemmego/api/app"
	"io"
	"net/http"
)

type TemplResponse struct {
	component templ.Component
	ctx       app.Context
}

func New(c app.Context, component templ.Component) *TemplResponse {
	return &TemplResponse{component: component, ctx: c}
}

func Respond(c app.Context, component templ.Component) error {
	tr := New(c, component)
	return tr.ctx.Render(tr)
}

func (t *TemplResponse) Render(w io.Writer) error {
	t.ctx.SetHeader("content-type", "text/html")
	if t.ctx.Status() == 0 {
		t.ctx.SetStatus(http.StatusOK)
	}
	t.ctx.ResponseWriter().WriteHeader(t.ctx.Status())
	return t.component.Render(t.ctx.RequestContext(), t.ctx.ResponseWriter())
}
