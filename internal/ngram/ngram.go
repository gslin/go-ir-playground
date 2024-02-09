package ngram

import (
	"regexp"
)

var re1 *regexp.Regexp

func init() {
	re1 = regexp.MustCompile("(\\w+|\\p{L})")
}

func Bigram(s string) []string {
	bag := split(s)

	r := make([]string, 0)
	for i := 0; i < len(bag) - 1; i++ {
		r = append(r, bag[i] + bag[i + 1])
	}

	return r
}

func Unigram(s string) []string {
	return split(s)
}

func split(s string) []string {
	bag := make([]string, 0)
	for _, w := range re1.FindAllStringSubmatch(s, -1) {
		bag = append(bag, w[0])
	}
	return bag
}
