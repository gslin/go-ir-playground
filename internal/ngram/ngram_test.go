package ngram_test

import (
	"testing"
	"github.com/stretchr/testify/assert"

	"github.com/gslin/go-ir-playground/internal/ngram"
)

func testBigram(t *testing.T) {
	a := ngram.Bigram("test")
	assert.Equal(t, len(a), 1)
	assert.Equal(t, a[0], "test")

	a = ngram.Bigram("測試")
	assert.Equal(t, len(a), 1)
	assert.Equal(t, a[0], "測試")

	a = ngram.Bigram("中文測試")
	assert.Equal(t, len(a), 3)
	assert.Equal(t, a[0], "中文")
	assert.Equal(t, a[1], "文測")
	assert.Equal(t, a[2], "測試")
}

func TestUnigram(t *testing.T) {
	a := ngram.Unigram("test")
	assert.Equal(t, len(a), 1)
	assert.Equal(t, a[0], "test")

	a = ngram.Unigram("測試")
	assert.Equal(t, len(a), 2)
	assert.Equal(t, a[0], "測")
	assert.Equal(t, a[1], "試")
}
