package factories

import (
	"unifriends-db-script/models"

	"github.com/go-faker/faker/v4"
)

func CreateUser(quantity int, major models.Major) []models.User {
	var finalUserList []models.User
	for i := 0; i < quantity; i++ {
		user := models.User{
			Email:             faker.Email(),
			Name:              faker.Name(),
			ProfilePictureURL: faker.URL(),
			IsAdmin:           false,
			MajorID:           major.Id,
		}

		user.InsertUser()
		finalUserList = append(finalUserList, user)
	}

	return finalUserList
}
