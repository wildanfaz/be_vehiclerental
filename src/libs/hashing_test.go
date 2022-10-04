package libs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashingPassword(t *testing.T) {
	myPassword := "hellotest123"
	result, err := HashingPassword(myPassword)

	//**check
	// err = errors.New("error")
	if err != nil {
		t.Fatal("msg :", err)
	}
	//**assert.Nil(t, err, err)

	//**check
	// result = myPassword
	if result == myPassword {
		t.Fatal("password has not been hashed")
	}

	assert.IsType(t, "string", result, "result is not string")
}

func TestCheckPassword(t *testing.T) {
	myPassword := "hellotest123"
	result, _ := HashingPassword(myPassword)
	err := CheckPassword(result, myPassword)

	assert.Nil(t, err, "incorrect password")
}
