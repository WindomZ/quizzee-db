package buntdb

import (
	"testing"

	quiz "github.com/WindomZ/quizzee-db"
	"github.com/WindomZ/testify/assert"
)

var demoTable = []byte("test")
var demo = &quiz.Quiz{
	Question: "   何为四方？  ",
	Options: []string{
		"东1a", " 南南2b", " 西西西3c ", " 北北北北4d ",
	},
	Answer: " 北北北北4d",
}
var correct = &quiz.Quiz{
	Question: "何为四方？",
	Options: []string{
		"东1a", "南南2b", "西西西3c", "北北北北4d",
	},
	Answer: "北北北北4d",
}

func init() {
	quiz.Open(demoTable, "../data/buntdb.db")
}

func TestOpen(t *testing.T) {
	assert.NoError(t, quiz.Close())
	assert.NoError(t, quiz.Open(demoTable, "../data/buntdb.db"))
}

func TestPut(t *testing.T) {
	assert.True(t, demo.Completion())
	assert.NoError(t, demo.Store())
}

func TestGet(t *testing.T) {
	q := quiz.GetQuiz(correct.Question)
	assert.True(t, q.Completion())
	assert.Equal(t, correct.Question, q.Question)
	assert.Equal(t, correct.Options, q.Options)
	assert.Equal(t, correct.Answer, q.Answer)
	assert.NotEmpty(t, q.Update)
}

func TestParseQuiz(t *testing.T) {
	demo.Update = 0
	q := quiz.ParseQuiz(demo.Bytes())
	assert.True(t, q.Completion())
	assert.Equal(t, demo.Question, q.Question)
	assert.Equal(t, demo.Options, q.Options)
	assert.Equal(t, demo.Answer, q.Answer)
	assert.Empty(t, q.Update)
}

func TestCount(t *testing.T) {
	assert.Equal(t, 1, quiz.Count())
}

func TestIterator(t *testing.T) {
	cnt := 0
	assert.NoError(t, quiz.Iterator(func(q *quiz.Quiz) bool {
		cnt++
		return assert.Equal(t, correct.Question, q.Question) &&
			assert.Equal(t, correct.Options, q.Options) &&
			assert.Equal(t, correct.Answer, q.Answer)
	}))
	assert.Equal(t, 1, cnt)
}

func TestClose(t *testing.T) {
	assert.NoError(t, quiz.Close())
}
