package main

import (
	"fmt"
	"strings"

	"github.com/gslin/go-ir-playground/internal/artifact"
	"github.com/gslin/go-ir-playground/internal/tokenizer"
)

func main() {
	articles := artifact.Read("data/articles.json")

	tokens := make(map[string][]string)
	tf := make(map[string]map[string]int)
	df := make(map[string]int)

	for _, article := range articles {
		str := article.Title + "\n" + article.Body

		bag := tokenizer.Tokenize(str)
		tokens[article.Id] = bag

		// Init TF:
		tf[article.Id] = make(map[string]int)

		for _, w := range bag {
			// Handle TF:
			tf[article.Id][w] = strings.Count(str, w)

			// Handle DF:
			df[w] += 1
		}
	}

	fmt.Println("TF & DF Built")
}
