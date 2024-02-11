package main

import (
	"fmt"
	"math"
	"os"
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
		str := strings.ToLower(article.Title + "\n" + article.Body)

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

	q := strings.ToLower(os.Args[1])
	q_tokens := tokenizer.Tokenize(q)

	for _, article := range articles {
		var score float64 = 0.0
		for _, w := range q_tokens {
			if tf[w] != nil {
				score += float64(tf[w][article.Id]) * math.Log2(float64(len(articles) / df[w]))
			}
		}
		fmt.Printf("Article %v: %v\n", article.Id, score)
	}
}
