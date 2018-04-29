package filter

import (
	"github.com/jianyuan/htmlpipeline"
	blackfriday "gopkg.in/russross/blackfriday.v2"
)

type MarkdownFilter struct {
	options []blackfriday.Option
}

func NewMarkdownFilter(options ...blackfriday.Option) htmlpipeline.Filter {
	return &MarkdownFilter{
		options: append([]blackfriday.Option{}, options...),
	}
}

func (md *MarkdownFilter) Render(ctx *htmlpipeline.Context) {
	input := []byte(ctx.HTML())
	output := blackfriday.Run(input, md.options...)
	ctx.WriteHTML(string(output))
}
