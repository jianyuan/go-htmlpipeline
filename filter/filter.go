package filter

type Filter interface {
	Render([]byte) []byte
}
