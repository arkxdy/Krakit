package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/krakit/auth-service/internal/dto"
	"github.com/krakit/auth-service/internal/middleware"
	"github.com/krakit/auth-service/internal/service"
	"github.com/krakit/auth-service/internal/utils"
)

type AuthHandler struct {
	authservice service.AuthService
	config      utils.Config
}

func (h *AuthHandler) Signup(c *gin.Context) {
	var req dto.SignupRequest

	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.Error(dto.ErrBadRequest, err.Error()))
		return
	}

	res, rfshToken, err := h.authservice.Signup(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Error(dto.ErrInternal, err.Error()))
		return
	}
	c.SetCookieData(&http.Cookie{
		Name:     "refresh_token",
		Value:    rfshToken,
		Path:     "/",
		MaxAge:   int(h.config.TokenDuration),
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	})
	c.JSON(http.StatusOK, dto.Success(res, ""))
}
func (h *AuthHandler) Login(c *gin.Context) {
	var req dto.LoginRequest

	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.Error(dto.ErrBadRequest, err.Error()))
		return
	}

	res, rfshToken, err := h.authservice.Login(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Error(dto.ErrInternal, err.Error()))
		return
	}
	c.SetCookieData(&http.Cookie{
		Name:     "refresh_token",
		Value:    rfshToken,
		Path:     "/",
		MaxAge:   int(h.config.TokenDuration),
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	})
	c.JSON(http.StatusOK, dto.Success(res, ""))
}
func (h *AuthHandler) Refresh(c *gin.Context) {
	rfshToken, err := c.Cookie("refresh_token")
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Error(dto.ErrInternal, err.Error()))
		return
	}

	res, rfshToken, err := h.authservice.Refresh(c, rfshToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Error(dto.ErrInternal, err.Error()))
		return
	}
	c.SetCookieData(&http.Cookie{
		Name:     "refresh_token",
		Value:    rfshToken,
		Path:     "/",
		MaxAge:   int(h.config.TokenDuration),
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	})
	c.JSON(http.StatusOK, dto.Success(res, ""))
}
func (h *AuthHandler) Logout(c *gin.Context) {
	rfshToken, err := c.Cookie("refresh_token")
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Error(dto.ErrInternal, err.Error()))
		return
	}

	err = h.authservice.Logout(c, rfshToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Error(dto.ErrInternal, err.Error()))
		return
	}

	c.JSON(http.StatusOK, dto.Success(nil, "Logged out successfully"))
}
func (h *AuthHandler) LogoutAll(c *gin.Context) {
	user, valid := middleware.GetUserFromContext(c)
	if !valid {
		c.JSON(http.StatusUnauthorized, dto.Error(dto.ErrValidationFailed, ""))
		return
	}
	err := h.authservice.LogoutAll(c, user.Sub)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Error(dto.ErrInternal, err.Error()))
		return
	}

	c.JSON(http.StatusOK, dto.Success(nil, "Logged out from all devices"))
}
func (h *AuthHandler) GoogleLogin(c *gin.Context) {
	var req dto.GoogleAuthRequest

	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.Error(dto.ErrBadRequest, err.Error()))
		return
	}

	res, _, err := h.authservice.GoogleLogin(c, req.IDToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Error(dto.ErrInternal, err.Error()))
		return
	}

	c.JSON(http.StatusOK, dto.Success(res, ""))
}
func NewAuthHandler(a service.AuthService) *AuthHandler {
	return &AuthHandler{
		authservice: a,
	}
}
