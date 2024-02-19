package tokenizer_test

import (
	"testing"
	"github.com/stretchr/testify/assert"

	"github.com/gslin/go-ir-playground/internal/tokenizer"
)

func TestTokenize(t *testing.T) {
	a := tokenizer.Tokenize("test")
	assert.Equal(t, len(a), 1)
	assert.Equal(t, a[0], "test")
}
