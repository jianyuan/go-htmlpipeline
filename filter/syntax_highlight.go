package filter

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/alecthomas/chroma"
	"github.com/alecthomas/chroma/formatters/html"
	"github.com/alecthomas/chroma/lexers"
	"github.com/alecthomas/chroma/styles"
)

type SyntaxHighlightFilter struct {
}

var _ Filter = (*SyntaxHighlightFilter)(nil)

func NewSyntaxHighlightFilter() Filter {
	return &SyntaxHighlightFilter{}
}

func (sh *SyntaxHighlightFilter) Render(input []byte) []byte {
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(input))
	if err != nil {
		return input
	}

	doc.Find("pre").Each(func(i int, s *goquery.Selection) {
		// TODO: make this customizable

		codeSelection := s.Find("code")

		var lang string
		if raw, ok := codeSelection.Attr("class"); ok {
			if strings.HasPrefix(raw, "language-") {
				lang = raw[len("language-"):]
			}
		}

		var lexer chroma.Lexer
		if lang == "" {
			lexer = lexers.Analyse(s.Text())
		} else {
			lexer = lexers.Get(lang)
		}

		if lexer == nil {
			lexer = lexers.Fallback
		}

		lexer = chroma.Coalesce(lexer)

		style := styles.Get("swapoff")
		if style == nil {
			style = styles.Fallback
		}

		formatter := html.New(html.Standalone())

		var err error

		iterator, err := lexer.Tokenise(nil, codeSelection.Text())
		if err != nil {
			fmt.Println(err)
			return
		}

		var w bytes.Buffer
		err = formatter.Format(&w, style, iterator)
		if err != nil {
			fmt.Println(err)
			return
		}

		s.ReplaceWithHtml(w.String())
	})

	output, err := doc.Find("body").Html()
	if err != nil {
		return input
	}

	return []byte(output)
}