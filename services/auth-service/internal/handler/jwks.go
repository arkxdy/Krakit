package handler

import (
	"encoding/base64"
	"math/big"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/krakit/auth-service/internal/dto"
	"github.com/krakit/auth-service/internal/utils"
)

type JWKSHandler struct {
	ks *utils.KeyStore
}

func NewJWKSHandler(ks *utils.KeyStore) *JWKSHandler {
	return &JWKSHandler{ks: ks}
}

func (h *JWKSHandler) GetJWKS(c *gin.Context) {

	var keys []gin.H

	for _, k := range h.ks.GetAll() {
		keys = append(keys, gin.H{
			"kid": k.Kid,
			"kty": "RSA",
			"alg": "RS256",
			"use": "sig",
			"n":   base64.RawURLEncoding.EncodeToString(k.PublicKey.N.Bytes()),
			"e":   base64.RawURLEncoding.EncodeToString(big.NewInt(int64(k.PublicKey.E)).Bytes()),
		})
	}

	c.JSON(http.StatusOK, dto.Success(keys, ""))
}
