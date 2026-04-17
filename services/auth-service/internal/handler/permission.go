package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/krakit/auth-service/internal/dto"
	"github.com/krakit/auth-service/internal/service"
)

type PermissionHandler struct {
	service service.PermissionService
}

func (h *PermissionHandler) ListPermissions(c *gin.Context) {
	perms, err := h.service.ListPermissions(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Error(dto.ErrInternal, err.Error()))
		return
	}

	c.JSON(http.StatusOK, dto.Success(perms, ""))
}

func (h *PermissionHandler) AssignPermissionToRole(c *gin.Context) {
	role := c.Param("role")
	if role == "" {
		c.JSON(http.StatusBadRequest, dto.Error(dto.ErrBadRequest, "missing role"))
		return
	}

	var req dto.AssignPermissionRequest
	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.Error(dto.ErrBadRequest, err.Error()))
		return
	}

	// Ensure it's a valid UUID even if binding tags are bypassed/changed.
	if _, err := uuid.Parse(req.PermissionID); err != nil {
		c.JSON(http.StatusBadRequest, dto.Error(dto.ErrValidationFailed, "invalid permission_id"))
		return
	}

	if err := h.service.AssignPermissionToRole(c, role, req.PermissionID); err != nil {
		c.JSON(http.StatusBadRequest, dto.Error(dto.ErrBadRequest, err.Error()))
		return
	}

	c.JSON(http.StatusOK, dto.Success(nil, "Permission assigned"))
}

func (h *PermissionHandler) RemovePermissionFromRole(c *gin.Context) {
	role := c.Param("role")
	permissionID := c.Param("permission_id")

	if role == "" || permissionID == "" {
		c.JSON(http.StatusBadRequest, dto.Error(dto.ErrBadRequest, "missing role or permission_id"))
		return
	}

	if _, err := uuid.Parse(permissionID); err != nil {
		c.JSON(http.StatusBadRequest, dto.Error(dto.ErrValidationFailed, "invalid permission_id"))
		return
	}

	if err := h.service.RemovePermissionFromRole(c, role, permissionID); err != nil {
		c.JSON(http.StatusBadRequest, dto.Error(dto.ErrBadRequest, err.Error()))
		return
	}

	c.JSON(http.StatusOK, dto.Success(nil, "Permission removed"))
}

func NewPermissionHandler(s service.PermissionService) *PermissionHandler {
	return &PermissionHandler{service: s}
}
