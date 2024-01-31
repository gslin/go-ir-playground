package main

import (
	"fmt"

	"github.com/gslin/go-ir-playground/internal/artifact"
	"github.com/gslin/go-ir-playground/internal/ngram"
)

func main() {
	articles := artifact.Read("data/articles.json")

	for _, article := range articles {
		title_bag := ngram.Bigram(article.Title)
		body_bag := ngram.Bigram(article.Body)

		fmt.Printf("title_bag = %v\n", title_bag)
		fmt.Printf("body_bag = %v\n", body_bag)
	}
}
