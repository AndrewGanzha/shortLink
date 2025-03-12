package jwt

import jwt5 "github.com/golang-jwt/jwt/v5"

type JWT struct {
	Secret string
}

func NewJWT(secret string) *JWT {
	return &JWT{
		Secret: secret,
	}
}
func (jwt *JWT) Create(email string) (string, error) {
	t := jwt5.NewWithClaims(jwt5.SigningMethodHS256, jwt5.MapClaims{
		"email": email,
	})
	s, err := t.SignedString([]byte(jwt.Secret))

	if err != nil {
		return "", err
	}

	return s, nil
}
