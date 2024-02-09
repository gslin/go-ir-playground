package tokenize

import (
	"slices"

	"github.com/gslin/go-ir-playground/internal/ngram"
)

func Tokenize(s string) []string {
	bag := append(ngram.Unigram(s), ngram.Bigram(s)...)
	slices.Sort(bag)
	slices.Compact[[]string, string](bag)
	return bag
}
