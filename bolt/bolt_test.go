package bolt

import (
	"testing"

	quiz "github.com/WindomZ/quizzee-db"
	"github.com/WindomZ/testify/assert"
)

var demoTable = []byte("test_question")
var demo = &quiz.Quiz{
	Question: "十万个为什么？",
	Options: []string{
		"东1a", "南南2b", "西西西3c", "北北北北4d",
	},
	Answer: "中",
}

func init() {
	quiz.Open(demoTable, "../data/test.db")
}

func TestOpen(t *testing.T) {
	assert.NoError(t, quiz.Close())
	assert.NoError(t, quiz.Close())
	assert.NoError(t, quiz.Open(demoTable, "../data/test.db"))
}

func TestPut(t *testing.T) {
	assert.True(t, demo.Completion())
	assert.NoError(t, demo.Store(demoTable))
}

func TestGet(t *testing.T) {
	q := quiz.GetQuiz(demoTable, "   十万个为什么？  ")
	assert.True(t, q.Completion())
	assert.Equal(t, demo.Question, q.Question)
	assert.Equal(t, demo.Options, q.Options)
	assert.Equal(t, demo.Answer, q.Answer)
	assert.NotEmpty(t, q.Update)
}

func TestParseQuiz(t *testing.T) {
	demo.Update = 0
	q := quiz.ParseQuiz(demoTable, demo.Bytes())
	assert.True(t, q.Completion())
	assert.Equal(t, demo.Question, q.Question)
	assert.Equal(t, demo.Options, q.Options)
	assert.Equal(t, demo.Answer, q.Answer)
	assert.Empty(t, q.Update)
}

func TestCount(t *testing.T) {
	assert.Equal(t, 1, quiz.Count(demoTable))
}

func TestClose(t *testing.T) {
	assert.NoError(t, quiz.Close())
}
