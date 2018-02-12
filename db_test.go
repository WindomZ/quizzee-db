package quizzee_db

import (
	"testing"

	"github.com/WindomZ/testify/assert"
)

var demoName = []byte("test_question")
var demo = &Quiz{
	Question: "十万个为什么？",
	Options: []string{
		"东1a", "南南2b", "西西西3c", "北北北北4d",
	},
	Answer: "中",
}

func init() {
	Open("./data/test.db")
	Register(demoName)
}

func TestOpen(t *testing.T) {
	assert.NoError(t, Close())
	assert.NoError(t, Close())
	assert.NoError(t, Open("./data/test.db"))
}

func TestRegister(t *testing.T) {
	assert.NoError(t, Register(demoName))
}

func TestPut(t *testing.T) {
	assert.True(t, demo.Completion())
	assert.NoError(t, demo.Store(demoName))
}

func TestGet(t *testing.T) {
	q := GetQuiz(demoName, "   十万个为什么？  ")
	assert.True(t, q.Completion())
	assert.Equal(t, demo.Question, q.Question)
	assert.Equal(t, demo.Options, q.Options)
	assert.Equal(t, demo.Answer, q.Answer)
	assert.NotEmpty(t, q.Update)
}

func TestCount(t *testing.T) {
	assert.Equal(t, 1, Count(demoName))
}

func TestClose(t *testing.T) {
	assert.NoError(t, Close())
}
