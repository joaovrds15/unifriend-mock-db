package models

import (
	"log"
)

type UserResponse struct {
	Id         int64
	QuestionID int64
	OptionID   int64
	UserID     int64
}

func (userResponse *UserResponse) InsertUserResponse() (*UserResponse, error) {
	stmt, err := Db.Prepare("INSERT INTO user_responses (question_id, option_id, user_id) VALUES (?, ?, ?)")

	if err != nil {
		log.Fatal(err.Error())
	}

	result, errInsert := stmt.Exec(userResponse.QuestionID, userResponse.OptionID, userResponse.UserID)

	if errInsert != nil {
		log.Fatal(errInsert.Error())
	}

	userResponse.Id, _ = result.LastInsertId()

	return userResponse, errInsert
}
