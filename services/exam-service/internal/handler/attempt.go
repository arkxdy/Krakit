package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/krakit/exam-service/internal/dto"
	"github.com/krakit/exam-service/internal/service"
)

// =========================
// ATTEMPT HANDLER
// =========================

type AttemptHandler struct {
	attemptSvc service.AttemptService
}

func NewAttemptHandler(attemptSvc service.AttemptService) *AttemptHandler {
	return &AttemptHandler{attemptSvc: attemptSvc}
}

func (h *AttemptHandler) StartAttempt(c *gin.Context) {
	var req dto.StartAttemptReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// userID comes from JWT middleware — same pattern as auth-service
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	attempt, err := h.attemptSvc.StartAttempt(c.Request.Context(), userID.(string), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, attempt)
}

func (h *AttemptHandler) GetAttempt(c *gin.Context) {
	attemptID := c.Param("id")

	attempt, err := h.attemptSvc.GetAttempt(c.Request.Context(), attemptID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, attempt)
}

func (h *AttemptHandler) GetAttemptQuestions(c *gin.Context) {
	attemptID := c.Param("id")

	questions, err := h.attemptSvc.GetAttemptQuestions(c.Request.Context(), attemptID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, questions)
}

func (h *AttemptHandler) GetAttemptQuestionsBySection(c *gin.Context) {
	attemptID := c.Param("id")
	sectionID := c.Param("section_id")

	questions, err := h.attemptSvc.GetAttemptQuestionsBySection(c.Request.Context(), attemptID, sectionID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, questions)
}

func (h *AttemptHandler) SaveAnswer(c *gin.Context) {
	var ans dto.Answer
	if err := c.ShouldBindJSON(&ans); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ans.AttemptID = c.Param("id")

	if err := h.attemptSvc.SaveAnswer(c.Request.Context(), ans); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "answer saved"})
}

func (h *AttemptHandler) GetAnswers(c *gin.Context) {
	attemptID := c.Param("id")

	answers, err := h.attemptSvc.GetAnswers(c.Request.Context(), attemptID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, answers)
}

func (h *AttemptHandler) SubmitAttempt(c *gin.Context) {
	attemptID := c.Param("id")

	result, err := h.attemptSvc.SubmitAttempt(c.Request.Context(), attemptID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

func (h *AttemptHandler) GetResult(c *gin.Context) {
	attemptID := c.Param("id")

	result, err := h.attemptSvc.GetResult(c.Request.Context(), attemptID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

func (h *AttemptHandler) GetUserAttempts(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	attempts, err := h.attemptSvc.GetUserAttempts(c.Request.Context(), userID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, attempts)
}
