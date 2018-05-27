package filter

import (
	"fmt"
	"html/template"

	"github.com/jianyuan/go-htmlpipeline"
)

var _ htmlpipeline.Filter = (*PlainTextInputFilter)(nil)

type PlainTextInputFilter struct {
}

func NewPlainTextInputFilter() htmlpipeline.Filter {
	return &PlainTextInputFilter{}
}

func (*PlainTextInputFilter) Render(ctx *htmlpipeline.Context) error {
	html, err := ctx.HTML()
	if err != nil {
		return err
	}

	output := template.HTMLEscapeString(html)
	output = fmt.Sprintf("<div>%s</div>", output)
	return ctx.WriteHTML(output)
}
