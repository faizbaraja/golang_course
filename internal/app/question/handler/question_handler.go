package handler

import (
	"exercise/internal/app/domainQuestion"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type QuestionHandler struct {
	db *gorm.DB
}

func NewQuestionHandler(db *gorm.DB) *QuestionHandler {
	return &QuestionHandler{db: db}
}

func (uh QuestionHandler) NewQuestion(c *gin.Context) {
	exerciseIdString := c.Param("exerciseId")
	userId, ok := c.Request.Context().Value("user_id").(int)
	if !ok {
		c.JSON(http.StatusBadRequest, map[string]string{
			"message": "invalid userId",
		})
		return
	}

	exerciseId, err := strconv.Atoi(exerciseIdString)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{
			"message": "invalid exerciseId",
		})
		return
	}

	var newQuestion domainQuestion.NewQuestion
	if err := c.ShouldBind(&newQuestion); err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{
			"message": "invalid body",
		})
	}

	// setelah line ini adalah usecase
	question, err := domainQuestion.CreateNewquestion(
		exerciseId,
		userId,
		newQuestion.Body,
		newQuestion.OptionA,
		newQuestion.OptionB,
		newQuestion.OptionC,
		newQuestion.OptionD,
		newQuestion.CorrectAnswer)

	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}
	if err := uh.db.Create(question).Error; err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}
}
