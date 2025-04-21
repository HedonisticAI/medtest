package cryptoutils

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"

	"golang.org/x/crypto/bcrypt"
)

const TokenSize = 16
const HashCost = 14

func GenerateSecureToken(length int) string {
	b := make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	return hex.EncodeToString(b)
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), HashCost)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func EncodeToBase64(Token string) string {
	Encoded := base64.StdEncoding.EncodeToString([]byte(Token))
	return Encoded
}

func DecodeFromBase64(Encoded string) (string, error) {
	Decoded, err := base64.StdEncoding.DecodeString(Encoded)
	if err != nil {
		return "", err
	}
	return string(Decoded), nil
}
