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
	if len(os.Args) < 2 {
		fmt.Printf("You need to specify a keyword to search.\n")
		os.Exit(1)
	}

	articles := artifact.Read("data/articles.json")

	tokens := make(map[string][]string)
	token_num := 0
	tf := make(map[string]map[string]int)
	df := make(map[string]int)

	for _, article := range articles {
		id := article.Id
		str := strings.ToLower(article.Title + "\n" + article.Body)

		bag := tokenizer.Tokenize(str)
		tokens[id] = bag

		token_num += len(bag)

		for _, w := range bag {
			// Handle TF:
			if _, ok := tf[w]; !ok {
				tf[w] = make(map[string]int)
			}
			tf[w][id] += strings.Count(str, w)

			// Handle DF:
			df[w] += 1
		}
	}

	fmt.Println("TF & DF Built")

	n := len(articles)
	avgdl := float64(token_num / n)
	q := strings.ToLower(os.Args[1])
	q_tokens := tokenizer.Tokenize(q)

	for _, article := range articles {
		id := article.Id

		var score float64 = 0.0
		for _, w := range q_tokens {
			if tf[w] != nil {
				idf := math.Log2((float64(n - df[w]) + 0.5) / (float64(df[w]) + 0.5) + 1)
				score += idf * (float64(tf[w][id]) * 2.2) / (float64(tf[w][id]) + 1.2 * (0.25 + 0.75 * (float64(n) / avgdl)))
			}
		}

		if score > 0 {
			fmt.Printf("Article %v (https://blog.gslin.org/?p=%v): %v\n", id, id, score)
		}
	}
}
