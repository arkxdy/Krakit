package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type JWTClaims struct {
	Sub         string   `json:"sub"`
	Email       string   `json:"email"`
	Role        string   `json:"role"`
	Plan        string   `json:"plan"`
	Permissions []string `json:"permissions"`
	JTI         string   `json:"jti"`
	jwt.RegisteredClaims
}

type JWTMaker struct {
	keyStore *KeyStore
	issuer   string
	duration time.Duration
}

func NewJWTMaker(ks *KeyStore, issuer string, duration time.Duration) *JWTMaker {
	return &JWTMaker{keyStore: ks, issuer: issuer, duration: duration}
}

func (j *JWTMaker) GenerateToken(
	userID, email, role, plan string,
	perms []string,
) (string, string, error) {

	key := j.keyStore.GetActiveKey()
	if key == nil {
		return "", "", errors.New("no active key")
	}

	jti := uuid.NewString()

	claims := JWTClaims{
		Sub:         userID,
		Email:       email,
		Role:        role,
		Plan:        plan,
		Permissions: perms,
		JTI:         jti,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    j.issuer,
			Subject:   userID,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.duration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	token.Header["kid"] = key.Kid

	signed, err := token.SignedString(key.PrivateKey)
	return signed, jti, err
}

func (j *JWTMaker) VerifyToken(tokenStr string) (*JWTClaims, error) {

	token, err := jwt.ParseWithClaims(tokenStr, &JWTClaims{}, func(t *jwt.Token) (interface{}, error) {

		kid, ok := t.Header["kid"].(string)
		if !ok {
			return nil, errors.New("missing kid")
		}

		key := j.keyStore.GetKey(kid)
		if key == nil {
			return nil, errors.New("unknown kid")
		}

		return key.PublicKey, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*JWTClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
