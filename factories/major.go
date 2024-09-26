package factories

import (
	"unifriends-db-script/models"

	"github.com/go-faker/faker/v4"
)

func CreateMajor() models.Major {
	major := models.Major{
		Name: faker.Word(),
	}

	major.InsertMajor()

	return major
}
