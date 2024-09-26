package factories

import (
	"unifriends-db-script/models"

	"github.com/go-faker/faker/v4"
)

func CreateQuestions(quantity int, quiz_id int64) []models.QuestionTable {
	var questions []models.QuestionTable

	for i := 0; i < quantity; i++ {
		newQuestion := models.QuestionTable{
			Text:    faker.Sentence(),
			Quiz_id: quiz_id,
		}

		newQuestion.InsertQuestion()

		questions = append(questions, newQuestion)

	}

	return questions
}
