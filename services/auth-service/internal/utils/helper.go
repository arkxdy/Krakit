package utils

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"os"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

type Config struct {
	GoogleClientID string
	TokenSecret    string
	TokenDuration  int
}

func LoadConfig() *Config {
	tokenDuration, err := strconv.Atoi(os.Getenv("TOKEN_DURATION"))
	if err != nil {
		tokenDuration = 60 * 60
	}
	return &Config{
		GoogleClientID: os.Getenv("GOOGLE_CLIENT_ID"),
		TokenSecret:    os.Getenv("TOKEN_SECRET"),
		TokenDuration:  tokenDuration,
	}
}

func GenerateRefreshToken() (string, error) {
	b := make([]byte, 32) // 256-bit

	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	return base64.RawURLEncoding.EncodeToString(b), nil
}

func HashToken(token string, secret string) string {
	h := sha256.New()
	h.Write([]byte(token))
	h.Write([]byte(secret)) // pepper
	return hex.EncodeToString(h.Sum(nil))
}

func HashPassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost, // usually 10–12
	)
	if err != nil {
		return "", err
	}

	return string(hashed), nil
}

func CheckPassword(password string, hashed string) error {
	return bcrypt.CompareHashAndPassword(
		[]byte(hashed),
		[]byte(password),
	)
}
