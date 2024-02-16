package ngram_test

import (
	"testing"

	"github.com/gslin/go-ir-playground/internal/ngram"
)

func TestUnigram(t *testing.T) {
	a := ngram.Unigram("test")
	if len(a) != 1 || a[0] != "test" {
		t.Error()
	}

	a = ngram.Unigram("測試")
	if len(a) != 2 || a[0] != "測" || a[1] != "試" {
		t.Error()
	}
}
