package libs

import (
	"testing"
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
}
