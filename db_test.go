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
	DBRegister(demoName)
}

func TestDBPut(t *testing.T) {
	assert.True(t, demo.Completion())
	assert.NoError(t, demo.Store(demoName))
}

func TestDBGet(t *testing.T) {
	q := GetQuiz(demoName, "   十万个为什么？  ")
	assert.True(t, q.Completion())
	assert.Equal(t, demo.Question, q.Question)
	assert.Equal(t, demo.Options, q.Options)
	assert.Equal(t, demo.Answer, q.Answer)
	assert.NotEmpty(t, q.Update)
}

func TestDBCount(t *testing.T) {
	assert.Equal(t, 1, DBCount(demoName))
}
