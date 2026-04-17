package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/krakit/auth-service/internal/dto"
	"github.com/krakit/auth-service/internal/middleware"
	"github.com/krakit/auth-service/internal/service"
)

type UserHandler struct {
	service service.UserService
}

func (h *UserHandler) GetCurrentUser(c *gin.Context) {
	claim, valid := middleware.GetUserFromContext(c)
	if !valid {
		c.JSON(http.StatusUnauthorized, dto.Error(dto.ErrValidationFailed, ""))
		return
	}
	user, err := h.service.GetCurrentUser(c, claim.Sub)
	if err != nil {
		c.JSON(http.StatusNotFound, dto.Error(dto.ErrUserNotFound, ""))
		return
	}

	res := &dto.UserResponse{
		ID:          user.ID,
		Email:       user.Email,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		FullName:    user.FullName,
		Role:        user.Role,
		Plan:        user.Plan,
		LastLoginAt: user.LastLoginAt,
	}
	c.JSON(http.StatusOK, dto.Success(res, ""))
}
func (h *UserHandler) UpdateProfile(c *gin.Context) {
	var req dto.UpdateProfileRequest
	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.Error(dto.ErrBadRequest, err.Error()))
		return
	}

	claim, valid := middleware.GetUserFromContext(c)
	if !valid {
		c.JSON(http.StatusUnauthorized, dto.Error(dto.ErrValidationFailed, ""))
		return
	}

	user, err := h.service.UpdateProfile(c, claim.Sub, req)
	if err != nil {
		c.JSON(http.StatusNotFound, dto.Error(dto.ErrUserNotFound, ""))
		return
	}

	res := &dto.UserResponse{
		ID:          user.ID,
		Email:       user.Email,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		FullName:    user.FullName,
		Role:        user.Role,
		Plan:        user.Plan,
		LastLoginAt: user.LastLoginAt,
	}

	c.JSON(http.StatusAccepted, dto.Success(res, ""))
}
func (h *UserHandler) ChangePassword(c *gin.Context) {
	var req dto.ChangePasswordRequest
	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.Error(dto.ErrBadRequest, err.Error()))
		return
	}

	claim, valid := middleware.GetUserFromContext(c)
	if !valid {
		c.JSON(http.StatusUnauthorized, dto.Error(dto.ErrValidationFailed, ""))
		return
	}

	err := h.service.ChangePassword(c, claim.Sub, req)
	if err != nil {
		c.JSON(http.StatusNotFound, dto.Error(dto.ErrValidationFailed, ""))
		return
	}

	c.JSON(http.StatusAccepted, dto.Success(nil, "Password updated successfully"))
}

func NewUserHandler(s service.UserService) *UserHandler {
	return &UserHandler{service: s}
}
