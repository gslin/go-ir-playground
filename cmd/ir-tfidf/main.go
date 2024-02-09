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

		for _, w := range bag {
			// Handle TF:
			if _, ok := tf[w]; !ok {
				tf[w] = make(map[string]int)
			}
			tf[w][article.Id] += strings.Count(str, w)

			// Handle DF:
			df[w] += 1
		}
	}

	fmt.Println("TF & DF Built")
}
