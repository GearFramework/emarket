package alog

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewLogger(t *testing.T) {
	log := NewLogger("Test")
	assert.IsType(t, &Alog{}, log)
}
