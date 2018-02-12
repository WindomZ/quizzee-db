package quizzee_db

import (
	"testing"

	"github.com/WindomZ/testify/assert"
)

func TestParseQuiz(t *testing.T) {
	demo.Update = 0
	q := ParseQuiz(demoName, demo.Bytes())
	assert.True(t, q.Completion())
	assert.Equal(t, demo.Question, q.Question)
	assert.Equal(t, demo.Options, q.Options)
	assert.Equal(t, demo.Answer, q.Answer)
	assert.Empty(t, q.Update)
}

func TestNewQuiz(t *testing.T) {
	q := NewQuiz(" ")
	assert.True(t, q.Valid())
	q.trim()
	assert.False(t, q.Valid())
}
