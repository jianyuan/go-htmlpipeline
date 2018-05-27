package htmlpipeline

type Filter interface {
	Render(*Context) error
}
