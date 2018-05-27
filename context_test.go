package htmlpipeline

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmptyContext(t *testing.T) {
	ctx := NewContext("")

	html, err := ctx.HTML()
	assert.Empty(t, html)
	assert.NoError(t, err)

	doc, err := ctx.Document()
	assert.NotNil(t, doc)
	assert.NoError(t, err)
}
