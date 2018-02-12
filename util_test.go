package quizzee_db

import (
	"testing"

	"github.com/WindomZ/testify/assert"
)

func TestTrimQuestion(t *testing.T) {
	assert.Equal(t, "简化问题", TrimQuestion("1. 简化问题？"))
	assert.Equal(t, "简化问题", TrimQuestion("1. 简化问题?"))
	assert.Equal(t, "简化问题", TrimQuestion("1. 简化问题"))
	assert.Equal(t, "简化问题", TrimQuestion("1.简化问题"))
	assert.Equal(t, "简化问题", TrimQuestion("1.简化问题？"))
	assert.Equal(t, "简化问题", TrimQuestion("1.简化问题?"))
	assert.Equal(t, "简化问题", TrimQuestion("简化问题？"))
}
