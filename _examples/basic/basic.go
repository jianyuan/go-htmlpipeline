package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jianyuan/go-htmlpipeline"
	"github.com/jianyuan/go-htmlpipeline/filter"
)

const TestInput = `
# H1

` + "```" + `python
def main():
	print('Hello world!')

` + "```" + `
`

func main() {
	pipeline := htmlpipeline.New(
		filter.NewMarkdownFilter(),
		filter.NewSyntaxHighlightFilter(),
	)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		output, err := pipeline.Render(TestInput)
		if err == nil {
			fmt.Fprint(w, output)
		} else {
			fmt.Fprintf(w, "Error: %v", err)
		}
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
