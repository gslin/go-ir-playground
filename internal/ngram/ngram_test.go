package ngram_test

import (
	"testing"
	"github.com/stretchr/testify/assert"

	"github.com/gslin/go-ir-playground/internal/ngram"
)

func TestUnigram(t *testing.T) {
	a := ngram.Unigram("test")
	assert.Equal(t, len(a), 1)
	assert.Equal(t, a[0], "test")

	a = ngram.Unigram("測試")
	assert.Equal(t, len(a), 2)
	assert.Equal(t, a[0], "測")
	assert.Equal(t, a[1], "試")
}
