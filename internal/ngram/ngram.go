package ngram

import (
	"regexp"
)

var re1, re2 *regexp.Regexp

func init() {
	re1 = regexp.MustCompile("\\PL+")
	re2 = regexp.MustCompile("")
}

func Bigram(s string) []string {
	bag := split(s)

	r := make([]string, 0)
	for i := 0; i < len(bag) - 1; i++ {
		r = append(r, bag[i] + bag[i + 1])
	}

	return r
}

func split(s string) []string {
	bag := make([]string, 0)
	for _, w := range re1.Split(s, -1) {
		bag = append(bag, re2.Split(w, -1)...)
	}
	return bag
}
