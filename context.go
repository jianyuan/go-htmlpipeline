package htmlpipeline

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

type Context struct {
	html *string
	doc  *goquery.Document
}

func NewContext(html string) *Context {
	return &Context{
		html: &html,
	}
}

func (ctx *Context) HTML() string {
	if ctx.html != nil {
		return *ctx.html
	}

	if ctx.doc != nil {
		html, err := ctx.doc.Html()
		if err == nil {
			ctx.html = &html
			return html
		}
	}

	// TODO: panic?
	return ""
}

func (ctx *Context) Document() *goquery.Document {
	if ctx.doc != nil {
		return ctx.doc
	}

	if ctx.html != nil {
		bodyNode := &html.Node{
			Type:     html.ElementNode,
			Data:     "body",
			DataAtom: atom.Body,
		}

		nodes, err := html.ParseFragment(strings.NewReader(*ctx.html), bodyNode)
		if err != nil {
			return nil
		}

		for _, node := range nodes {
			bodyNode.AppendChild(node)
		}

		doc := goquery.NewDocumentFromNode(bodyNode)
		ctx.doc = doc
		return doc
	}

	// TODO: panic?
	return nil
}

func (ctx *Context) WriteHTML(input string) {
	ctx.html = &input
	ctx.doc = nil
}

func (ctx *Context) WriteDocument(doc *goquery.Document) {
	ctx.doc = doc
	ctx.html = nil
}
