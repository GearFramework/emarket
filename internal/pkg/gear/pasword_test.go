package gear

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPassword(t *testing.T) {
	p1 := GenerateRandomPassword(16)
	assert.NotEmpty(t, p1)
	assert.Equal(t, 16, len(p1))
	fmt.Printf("Random pass 1: %s\n", p1)
	p2 := GenerateRandomPassword(16)
	assert.NotEmpty(t, p1)
	assert.Equal(t, 16, len(p1))
	fmt.Printf("Random pass 2: %s\n", p2)
	fmt.Printf("%s <> %s\n", p1, p2)
	assert.NotEqual(t, p1, p2)
	hash, err := HashPassword(p1)
	assert.NoError(t, err)
	assert.NotEmpty(t, hash)
	fmt.Printf("Hash pass 1: %s\n", hash)
	assert.True(t, ValidatePassword(hash, p1))
	hash, err = HashPassword(p2)
	assert.NoError(t, err)
	assert.NotEmpty(t, hash)
	fmt.Printf("Hash pass 2: %s\n", hash)
	assert.True(t, ValidatePassword(hash, p2))
}
