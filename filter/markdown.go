package filter

import (
	blackfriday "gopkg.in/russross/blackfriday.v2"

	"github.com/jianyuan/go-htmlpipeline"
)

var _ (htmlpipeline.Filter) = (*MarkdownFilter)(nil)

type MarkdownFilter struct {
	options []blackfriday.Option
}

func NewMarkdownFilter(options ...blackfriday.Option) htmlpipeline.Filter {
	return &MarkdownFilter{
		options: append([]blackfriday.Option{}, options...),
	}
}

func (md *MarkdownFilter) Render(ctx *htmlpipeline.Context) error {
	html, err := ctx.HTML()
	if err != nil {
		return err
	}

	input := []byte(html)
	output := blackfriday.Run(input, md.options...)
	return ctx.WriteHTML(string(output))
}
