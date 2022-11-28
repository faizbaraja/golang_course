package main

import (
	answerHandler "exercise/internal/app/answer/handler"
	"exercise/internal/app/database"
	"exercise/internal/app/exercise/handler"
	questionHandler "exercise/internal/app/question/handler"
	userHandler "exercise/internal/app/user/handler"
	"exercise/internal/pkg/middleware"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	r := gin.Default()
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]string{
			"message": "hello world",
		})
	})

	// gp := os.Getenv("GOPATH")
	// fmt.Println("path adalah " + gp)
	envErr := godotenv.Load()
	if envErr != nil {
		fmt.Printf("error load env file")
		os.Exit(1)
	}

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

	port := os.Getenv("PORT")
	// domain := os.Getenv("DOMAIN")
	appPort := "0.0.0.0:" + port
	// appPort := ":" + port
	fmt.Println("masuk " + appPort)
	r.Run(appPort)
}
