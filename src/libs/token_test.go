package libs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToken(t *testing.T) {
	result := NewToken("user", "user")

	assert.IsType(t, &claims{}, result)
}
