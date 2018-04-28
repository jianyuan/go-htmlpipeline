package htmlpipeline

import "github.com/jianyuan/htmlpipeline/filter"

type Pipeline struct {
	filters []filter.Filter
}

func New(filters ...filter.Filter) *Pipeline {
	return &Pipeline{
		filters: append(([]filter.Filter)(nil), filters...),
	}
}

func (p *Pipeline) Render(input []byte) []byte {
	output := input
	for _, filter := range p.filters {
		output = filter.Render(output)
	}
	return output
}
