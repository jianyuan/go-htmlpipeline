package filter

import (
	"fmt"
	"html/template"

	"github.com/jianyuan/go-htmlpipeline"
)

type PlainTextInputFilter struct {
}

func NewPlainTextInputFilter() htmlpipeline.Filter {
	return &PlainTextInputFilter{}
}

func (*PlainTextInputFilter) Render(ctx *htmlpipeline.Context) {
	output := template.HTMLEscapeString(ctx.HTML())
	output = fmt.Sprintf("<div>%s</div>", output)
	ctx.WriteHTML(output)
}
