package cryptoutils

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

const HashCost = 14

var jwtSecret = []byte("I am secret")

func NewJwt(Claims jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, Claims)
	// https://github.com/golang-jwt/jwt/blob/main/hmac.go
	// here SigningMethodHS512 = &SigningMethodHMAC{"HS512", crypto.SHA512}
	// so HS512 is SHA512 i guess
	str, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}
	return str, nil
}

func ParseJwt(tokenString string) (jwt.Claims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("there's an error with the signing method")
		}
		return jwtSecret, nil

	})
	if err != nil {
		return nil, err
	}
	return token.Claims, nil
}

func GenerateSecureToken() string {
	str := rand.Text() //sine it's in base32 so no errors
	return str
}
func HashToken(Token string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(Token), HashCost)
	return string(bytes), err
}

func CheckTokenHash(token, hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(token))
	return err
}

func EncodeToBase64(Token string) string {
	Encoded := base64.RawStdEncoding.EncodeToString([]byte(Token))
	return Encoded
}

func DecodeFromBase64(Encoded string) (string, error) {
	Decoded, err := base64.RawStdEncoding.DecodeString(Encoded)
	if err != nil {
		return "", err
	}
	return string(Decoded), nil
}
