package factories

import (
	"unifriends-db-script/models"

	"github.com/go-faker/faker/v4"
)

func CreateOptions(quantity int, question_id int64) []models.OptionTable {
	var options []models.OptionTable

	for i := 0; i < quantity; i++ {
		newOption := models.OptionTable{
			Text:       faker.Sentence(),
			QuestionId: question_id,
		}

		newOption.InsertOptions()

		options = append(options, newOption)
	}

	return options
}
