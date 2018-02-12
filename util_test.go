package quizzee_db

import (
	"testing"

	"github.com/WindomZ/testify/assert"
)

func TestTrimQuestion(t *testing.T) {
	assert.Equal(t, "内容", TrimQuestion("1. 内容？"))
	assert.Equal(t, "内容", TrimQuestion("1. 内容?"))
	assert.Equal(t, "内容", TrimQuestion("1. 内容"))
	assert.Equal(t, "内容", TrimQuestion("1.内容"))
	assert.Equal(t, "内容", TrimQuestion("1.内容？"))
	assert.Equal(t, "内容", TrimQuestion("1.内容?"))
	assert.Equal(t, "内容", TrimQuestion("内容？"))
}
