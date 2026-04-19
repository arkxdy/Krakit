package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/krakit/exam-service/internal/dto"
	"github.com/krakit/exam-service/internal/service"
)

// =========================
// EXAM HANDLER
// =========================

type ExamHandler struct {
	examSvc service.ExamService
}

func NewExamHandler(examSvc service.ExamService) *ExamHandler {
	return &ExamHandler{examSvc: examSvc}
}

func (h *ExamHandler) CreateExam(c *gin.Context) {
	var req dto.CreateExamReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	exam, err := h.examSvc.CreateExam(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, exam)
}

func (h *ExamHandler) UpdateExam(c *gin.Context) {
	var req dto.UpdateExamReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	req.ExamID = c.Param("id")

	if err := h.examSvc.UpdateExam(c.Request.Context(), req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "exam updated"})
}

func (h *ExamHandler) ListExams(c *gin.Context) {
	exams, err := h.examSvc.ListExams(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, exams)
}

func (h *ExamHandler) ListExamsPaginated(c *gin.Context) {
	var query struct {
		Limit  int32 `form:"limit"`
		Offset int32 `form:"offset"`
	}
	if err := c.ShouldBindQuery(&query); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if query.Limit == 0 {
		query.Limit = 20 // default page size
	}

	exams, err := h.examSvc.ListExamsPaginated(c.Request.Context(), query.Limit, query.Offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, exams)
}

func (h *ExamHandler) DisableExam(c *gin.Context) {
	examID := c.Param("id")

	if err := h.examSvc.DisableExam(c.Request.Context(), examID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "exam disabled"})
}

func (h *ExamHandler) PublishExam(c *gin.Context) {
	examID := c.Param("id")

	if err := h.examSvc.PublishExam(c.Request.Context(), examID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "exam published"})
}

func (h *ExamHandler) CreateExamSettings(c *gin.Context) {
	var req dto.ExamSettings
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	req.ExamID = c.Param("id")

	if err := h.examSvc.CreateExamSettings(c.Request.Context(), req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "settings saved"})
}

func (h *ExamHandler) CreateSection(c *gin.Context) {
	var req dto.CreateSectionReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	req.ExamID = c.Param("id")

	section, err := h.examSvc.CreateSection(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, section)
}

func (h *ExamHandler) UpdateSection(c *gin.Context) {
	var req dto.CreateSectionReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sectionID := c.Param("section_id")
	if err := h.examSvc.UpdateSection(c.Request.Context(), sectionID, req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "section updated"})
}

func (h *ExamHandler) GetSections(c *gin.Context) {
	examID := c.Param("id")

	sections, err := h.examSvc.GetSections(c.Request.Context(), examID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, sections)
}

func (h *ExamHandler) CreateSubject(c *gin.Context) {
	var req dto.CreateSubjectReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	subject, err := h.examSvc.CreateSubject(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, subject)
}

func (h *ExamHandler) GetSubjects(c *gin.Context) {
	subjects, err := h.examSvc.GetSubjects(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, subjects)
}
