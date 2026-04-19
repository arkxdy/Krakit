package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/krakit/exam-service/internal/dto"
	"github.com/krakit/exam-service/internal/service"
)

// =========================
// QUESTION HANDLER
// =========================

type QuestionHandler struct {
	questionSvc service.QuestionService
}

func NewQuestionHandler(questionSvc service.QuestionService) *QuestionHandler {
	return &QuestionHandler{questionSvc: questionSvc}
}

func (h *QuestionHandler) AttachQuestionSet(c *gin.Context) {
	var req dto.AttachQuestionSetReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	req.ExamID = c.Param("id")

	if err := h.questionSvc.AttachQuestionSet(c.Request.Context(), req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "question set attached"})
}

func (h *QuestionHandler) GetQuestionSets(c *gin.Context) {
	examID := c.Param("id")

	sets, err := h.questionSvc.GetQuestionSets(c.Request.Context(), examID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, sets)
}

func (h *QuestionHandler) CreateQuestions(c *gin.Context) {
	var questions []dto.Question
	if err := c.ShouldBindJSON(&questions); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.questionSvc.CreateQuestions(c.Request.Context(), questions); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "questions created"})
}

func (h *QuestionHandler) BulkUpsertQuestions(c *gin.Context) {
	var questions []dto.Question
	if err := c.ShouldBindJSON(&questions); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.questionSvc.BulkUpsertQuestions(c.Request.Context(), questions); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "questions upserted"})
}

func (h *QuestionHandler) CreatePassages(c *gin.Context) {
	var passages []dto.Passage
	if err := c.ShouldBindJSON(&passages); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.questionSvc.CreatePassages(c.Request.Context(), passages); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "passages created"})
}

func (h *QuestionHandler) BulkUpsertPassages(c *gin.Context) {
	var passages []dto.Passage
	if err := c.ShouldBindJSON(&passages); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.questionSvc.BulkUpsertPassages(c.Request.Context(), passages); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "passages upserted"})
}

func (h *QuestionHandler) DeleteQuestionsBySet(c *gin.Context) {
	setID := c.Param("set_id")

	if err := h.questionSvc.DeleteQuestionsBySet(c.Request.Context(), setID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "questions deleted"})
}

func (h *QuestionHandler) DeletePassagesBySection(c *gin.Context) {
	sectionID := c.Param("section_id")

	if err := h.questionSvc.DeletePassagesBySection(c.Request.Context(), sectionID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "passages deleted"})
}
