package main

import (
	answerHandler "exercise/internal/app/answer/handler"
	"exercise/internal/app/database"
	"exercise/internal/app/exercise/handler"
	questionHandler "exercise/internal/app/question/handler"
	userHandler "exercise/internal/app/user/handler"
	"exercise/internal/pkg/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]string{
			"message": "hello world",
		})
	})

	db := database.NewConnDatabase()
	exerciseHandler := handler.NewExerciseHandler(db)
	userHandler := userHandler.NewUserHandler(db)
	questionHandler := questionHandler.NewQuestionHandler(db)
	answerHandler := answerHandler.NewAnswerHandler(db)

	r.GET("/exercises/:id", middleware.WithAuh(), exerciseHandler.GetExerciseByID)
	r.GET("/exercises/:id/score", middleware.WithAuh(), exerciseHandler.GetScore)
	r.POST("/exercises", middleware.WithAuh(), exerciseHandler.NewExercise)

	r.POST("/exercises/:exerciseId/questions", middleware.WithAuh(), questionHandler.NewQuestion)

	r.POST("/exercises/:exerciseId/questions/:questionId/answer", middleware.WithAuh(), answerHandler.NewAnswer)

	r.POST("/register", userHandler.Register)
	r.POST("/login", userHandler.Login)
	r.Run(":1234")
}
