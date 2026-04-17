package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/krakit/auth-service/internal/dto"
	"github.com/krakit/auth-service/internal/middleware"
	"github.com/krakit/auth-service/internal/service"
)

type SessionHandler struct {
	service service.SessionService
}

func (h *SessionHandler) ListSessions(c *gin.Context) {
	claims, valid := middleware.GetUserFromContext(c)
	if !valid {
		c.JSON(http.StatusUnauthorized, dto.Error(dto.ErrUnauthorized, ""))
		return
	}

	res, err := h.service.ListSessions(c, claims.Sub, claims.JTI)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Error(dto.ErrInternal, err.Error()))
		return
	}

	c.JSON(http.StatusOK, dto.Success(res, ""))
}

func (h *SessionHandler) RevokeSession(c *gin.Context) {
	claims, valid := middleware.GetUserFromContext(c)
	if !valid {
		c.JSON(http.StatusUnauthorized, dto.Error(dto.ErrUnauthorized, ""))
		return
	}

	sessionID := c.Param("id")
	if sessionID == "" {
		c.JSON(http.StatusBadRequest, dto.Error(dto.ErrBadRequest, "missing session id"))
		return
	}

	isAdmin := claims.Role == "admin" || claims.Role == "super_admin"
	if err := h.service.RevokeSession(c, claims.Sub, sessionID, isAdmin); err != nil {
		// Service decides which errors to expose; for now return a safe error.
		c.JSON(http.StatusBadRequest, dto.Error(dto.ErrBadRequest, err.Error()))
		return
	}

	c.JSON(http.StatusOK, dto.Success(nil, "Session revoked"))
}

func NewSessionHandler(s service.SessionService) *SessionHandler {
	return &SessionHandler{service: s}
}
