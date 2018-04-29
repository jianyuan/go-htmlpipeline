package htmlpipeline

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
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
		body := ctx.doc.Find("body") // TODO: set fragment context
		html, err := body.Html()
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
		doc, err := goquery.NewDocumentFromReader(strings.NewReader(*ctx.html)) // TODO: set fragment context
		if err == nil {
			ctx.doc = doc
			return doc
		}
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
