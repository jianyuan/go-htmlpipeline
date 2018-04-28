package main

import (
	"fmt"

	"github.com/jianyuan/htmlpipeline"
	"github.com/jianyuan/htmlpipeline/filter"
)

func main() {
	pipeline := htmlpipeline.New(
		filter.NewMarkdownFilter(),
	)
	input := `
	# H1

	` + "```" + `python
	def main():
		print('Hello world!')

	` + "```" + `
`
	output := string(pipeline.Render([]byte(input)))
	fmt.Println(output)
}
