package models

import (
	"fmt"
	"log"
)

type QuizTable struct {
	Id          int64
	Title       string
	Description string
}

func (quiz *QuizTable) InsertQuiz() (*QuizTable, error) {
	stmt, err := Db.Prepare("INSERT INTO quiz_tables (title, description) VALUES (?, ?)")

	if err != nil {
		log.Fatal(err.Error())
		return nil, fmt.Errorf("InsertQuiz: failed executing db query: %w", err)
	}

	result, errInsert := stmt.Exec(quiz.Title, quiz.Description)

	if errInsert != nil {
		log.Println(errInsert.Error())
		return nil, fmt.Errorf("InsertQuiz: failed executing db query: %w", err)
	}

	quiz.Id, _ = result.LastInsertId()

	return quiz, errInsert
}
