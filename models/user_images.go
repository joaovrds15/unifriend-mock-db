package models

import "log"

type UsersImages struct {
	ID       int64
	ImageUrl string
	UserID   int64
}

func (userImage *UsersImages) InsertUserImage() (*UsersImages, error) {
	stmt, err := Db.Prepare("INSERT INTO users_images (image_url, user_id) VALUES (?, ?)")

	if err != nil {
		log.Fatal(err.Error())
	}

	result, errInsert := stmt.Exec(userImage.ImageUrl, userImage.UserID)

	if errInsert != nil {
		log.Fatal(errInsert.Error())
	}

	userImage.ID, _ = result.LastInsertId()

	return userImage, errInsert
}
