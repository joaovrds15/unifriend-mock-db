package factories

import (
	"fmt"
	"unifriends-db-script/models"

	"github.com/go-faker/faker/v4"
	"golang.org/x/exp/rand"
)

func CreateUser(quantity int, major models.Major) []models.User {
	var finalUserList []models.User
	for i := 0; i < quantity; i++ {
		user := models.User{
			Email:             faker.Email(),
			Name:              faker.Name(),
			ProfilePictureURL: faker.URL(),
			PhoneNumber:       generateBrazilianPhoneNumber(),
			IsAdmin:           false,
			MajorID:           major.Id,
		}

		user.InsertUser()
		finalUserList = append(finalUserList, user)
	}

	return finalUserList
}

func generateBrazilianPhoneNumber() string {

	areaCode := rand.Intn(90) + 10

	firstDigit := 9

	remainingDigits := rand.Intn(90000000) + 10000000

	phoneNumber := fmt.Sprintf("%02d%d%d", areaCode, firstDigit, remainingDigits)
	return phoneNumber
}
