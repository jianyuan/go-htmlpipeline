package filter

import blackfriday "gopkg.in/russross/blackfriday.v2"

type MarkdownFilter struct {
	options []blackfriday.Option
}

var _ Filter = (*MarkdownFilter)(nil)

func NewMarkdownFilter(options ...blackfriday.Option) Filter {
	return &MarkdownFilter{
		options: append([]blackfriday.Option{}, options...),
	}
}

func (md *MarkdownFilter) Render(input []byte) []byte {
	return blackfriday.Run(input, md.options...)
}
