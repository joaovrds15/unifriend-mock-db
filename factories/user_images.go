package factories

import (
	"unifriends-db-script/models"

	"github.com/go-faker/faker/v4"
)

func CreateUserImages(quantity int, user models.User) []models.UsersImages {
	var finalUserList []models.UsersImages
	for i := 0; i < quantity; i++ {
		userImage := models.UsersImages{
			ImageUrl: faker.URL(),
			UserID:   user.Id,
		}
		userImage.InsertUserImage()

		finalUserList = append(finalUserList, userImage)
	}

	return finalUserList
}
