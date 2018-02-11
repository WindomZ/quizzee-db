package quizzee_db

import (
	"testing"

	"github.com/WindomZ/testify/assert"
)

func TestConfigString(t *testing.T) {
	assert.Equal(t, "./data/data.db", ConfigString("db_path"))
}
