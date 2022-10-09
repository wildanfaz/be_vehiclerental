package libs

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var mySecretKey = []byte(os.Getenv("JWT_KEYS"))

type claims struct {
	// Name string `json:"name"`
	Email string `json:"email"`
	Role string `json:"role"`
	jwt.StandardClaims
}

// **payload
func NewToken(email, role string) *claims {
	return &claims{
		// Name: name,
		Email: email,
		Role: role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour).Unix(),
		},
	}
}

func (c *claims) Create() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)

	return token.SignedString(mySecretKey)
}

func CheckToken(token string) (*claims, error) {
	tokens, err := jwt.ParseWithClaims(token, &claims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(mySecretKey), nil
	})

	if err != nil {
		return nil, err
	}

	claims := tokens.Claims.(*claims)
	return claims, nil
}
