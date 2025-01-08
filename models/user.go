package models

import (
	"log"
)

type User struct {
	Id                int64
	Email             string
	Password          string
	Name              string
	PhoneNumber       string
	ProfilePictureURL string
	IsAdmin           bool
	MajorID           int64
}

func (user *User) InsertUser() (*User, error) {
	stmt, err := Db.Prepare("INSERT INTO users (email, password, name, profile_picture_url, phone_number, is_admin, major_id) VALUES (?, ?, ?, ?, ?, ?, ?)")

	if err != nil {
		log.Fatal(err.Error())
	}

	result, errInsert := stmt.Exec(user.Email, user.Password, user.Name, user.ProfilePictureURL, user.PhoneNumber, user.IsAdmin, user.MajorID)

	if errInsert != nil {
		log.Fatal(errInsert.Error())
	}

	user.Id, _ = result.LastInsertId()

	return user, errInsert
}
