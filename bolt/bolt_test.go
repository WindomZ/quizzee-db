package bolt

import (
	"testing"

	"github.com/WindomZ/testify/assert"
)

var demo *DB
var (
	demoTable = []byte("question")
	demoKey   = []byte("key")
	demoValue = []byte("value")
)

func TestOpen(t *testing.T) {
	db := Open("../data/test-bolt.db")
	if v, ok := db.(*DB); ok {
		demo = v
	}
}

func TestRegister(t *testing.T) {
	assert.NoError(t, demo.Register(demoTable))
}

func TestPut(t *testing.T) {
	assert.NoError(t, demo.Put(demoTable, demoKey, demoValue))
}

func TestGet(t *testing.T) {
	assert.Equal(t, demoValue, demo.Get(demoTable, demoKey))
}

func TestCount(t *testing.T) {
	assert.Equal(t, 1, demo.Count(demoTable))
}

func TestClose(t *testing.T) {
	assert.NoError(t, demo.Close())
}
