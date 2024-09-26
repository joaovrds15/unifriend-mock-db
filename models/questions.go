package models

import (
	"fmt"
	"log"
)

type QuestionTable struct {
	Id      int64
	Text    string
	Quiz_id int64
}

func (question *QuestionTable) InsertQuestion() (*QuestionTable, error) {
	stmt, err := Db.Prepare("INSERT INTO question_tables (text, quiz_id) VALUES (?, ?)")

	if err != nil {
		log.Println(err.Error())
		return nil, fmt.Errorf("InsertQuestion: failed executing db query: %w", err)
	}

	result, errInsert := stmt.Exec(question.Text, question.Quiz_id)

	if errInsert != nil {
		log.Println(errInsert.Error())
		return nil, fmt.Errorf("InsertQuestion: failed executing db query: %w", err)
	}

	question.Id, _ = result.LastInsertId()

	return question, nil
}
