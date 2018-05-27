package htmlpipeline

type Pipeline struct {
	filters []Filter
}

func New(filters ...Filter) *Pipeline {
	return &Pipeline{
		filters: append(([]Filter)(nil), filters...),
	}
}

func (p *Pipeline) Render(input string) (string, error) {
	ctx := NewContext(input)
	for _, filter := range p.filters {
		if err := filter.Render(ctx); err != nil {
			return "", err
		}
	}
	return ctx.HTML()
}
