package libs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToken(t *testing.T) {
	result := NewToken("user", "user")

	assert.IsType(t, &claims{}, result)
}

func TestCreateToken(t *testing.T) {
	result := NewToken("user", "user")

	token, err := result.Create()

	assert.Nil(t, err, "error is not nil")
	assert.IsType(t, "string", token)
}

func TestCheckToken(t *testing.T) {
	result := NewToken("user", "user")

	token, err := result.Create()

	assert.Nil(t, err, "error is not nil")
	assert.IsType(t, "string", token)

	resultToken, err := CheckToken(token)

	assert.Nil(t, err, "error is not nil")
	assert.NotNil(t, resultToken, "token nil")
}
