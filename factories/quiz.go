package factories

import (
	"unifriends-db-script/models"

	"github.com/go-faker/faker/v4"
)

func CreateQuiz() models.QuizTable {
	quiz := models.QuizTable{
		Title:       faker.Word(),
		Description: faker.Sentence(),
	}

	quiz.InsertQuiz()

	return quiz
}
