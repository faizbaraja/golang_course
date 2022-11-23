package domain

import (
	"errors"
	"time"
)

type Exercise struct {
	ID          int        `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Questions   []Question `json:"questions"`
}

type NewExercise struct {
	Title       string
	Description string
}

type Question struct {
	ID            int    `json:"id"`
	ExerciseID    int    `json:"exerciseID"`
	Body          string `json:"body"`
	OptionA       string `json:"option_a"`
	OptionB       string `json:"option_b"`
	OptionC       string `json:"option_c"`
	OptionD       string `json:"option_d"`
	CorrectAnswer string `json:"correctAnswer"`
	Score         int    `json:"score"`
	// CreatorID     int       `json:"creatorID"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Answer struct {
	ID         int       `json:"id"`
	ExerciseID int       `json:"exerciseID"`
	QuestionID int       `json:"questionID"`
	UserID     int       `json:"userID"`
	Answer     string    `json:"answer"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}

func CreateNewExercise(title, description string) (*Exercise, error) {
	if title == "" {
		return nil, errors.New("title is required")
	}

	if description == "" {
		return nil, errors.New("description is required")
	}

	return &Exercise{
		Title:       title,
		Description: description,
	}, nil
}
