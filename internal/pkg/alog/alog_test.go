package alog

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewLogger(t *testing.T) {
	log := NewLogger()
	assert.IsType(t, &Alog{}, log)
}
