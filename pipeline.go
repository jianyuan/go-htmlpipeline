package htmlpipeline

type Pipeline struct {
	filters []Filter
}

func New(filters ...Filter) *Pipeline {
	return &Pipeline{
		filters: append(([]Filter)(nil), filters...),
	}
}

func (p *Pipeline) Render(input string) string {
	ctx := NewContext(input)
	for _, filter := range p.filters {
		filter.Render(ctx)
	}
	return ctx.HTML()
}
