package quizzee_db

import (
	"testing"

	"github.com/WindomZ/testify/assert"
)

func TestNewQuiz(t *testing.T) {
	q := NewQuiz(" ")
	assert.True(t, q.Valid())
	q.trim()
	assert.False(t, q.Valid())
}
