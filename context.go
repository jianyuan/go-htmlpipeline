package htmlpipeline

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

// Context carries the HTML and/or goquery.Document between pipelines.
type Context struct {
	html *string
	doc  *goquery.Document
}

func NewContext(html string) *Context {
	return &Context{
		html: &html,
	}
}

func (ctx *Context) HTML() (string, error) {
	if ctx.html != nil {
		return *ctx.html, nil
	}

	if ctx.doc != nil {
		html, err := ctx.doc.Html()
		if err != nil {
			return "", err
		}

		ctx.html = &html
		return *ctx.html, nil
	}

	return "", nil
}

func (ctx *Context) Document() (*goquery.Document, error) {
	if ctx.doc != nil {
		return ctx.doc, nil
	}

	if ctx.html != nil {
		bodyNode := &html.Node{
			Type:     html.ElementNode,
			Data:     "body",
			DataAtom: atom.Body,
		}

		nodes, err := html.ParseFragment(strings.NewReader(*ctx.html), bodyNode)
		if err != nil {
			return nil, nil
		}

		for _, node := range nodes {
			bodyNode.AppendChild(node)
		}

		doc := goquery.NewDocumentFromNode(bodyNode)
		ctx.doc = doc
		return ctx.doc, nil
	}

	// TODO: return empty document
	return nil, nil
}

func (ctx *Context) WriteHTML(input string) error {
	ctx.html = &input
	ctx.doc = nil
	return nil
}

func (ctx *Context) WriteDocument(doc *goquery.Document) error {
	ctx.doc = doc
	ctx.html = nil
	return nil
}
