package models

import (
	"fmt"
	"log"
	"math/rand/v2"
)

type OptionTable struct {
	Id         int64
	Text       string
	QuestionId int64
}

func (option *OptionTable) InsertOptions() (*OptionTable, error) {
	stmt, err := Db.Prepare("INSERT INTO option_tables (text, question_id) VALUES (?, ?)")

	if err != nil {
		log.Fatal(err.Error())
		return nil, fmt.Errorf("InsertOptions: failed executing db query: %w", err)
	}

	result, errInsert := stmt.Exec(option.Text, option.QuestionId)

	if errInsert != nil {
		log.Fatal(errInsert.Error())
		return nil, fmt.Errorf("InsertOptions: failed executing db query: %w", err)
	}

	option.Id, _ = result.LastInsertId()

	return option, errInsert
}

func GenerateRandomOption(questionId int64) int64 {
	var questionCount int64
	errCount := Db.QueryRow("SELECT COUNT(id) FROM question_tables").Scan(&questionCount)

	if errCount != nil {
		log.Fatal(errCount.Error())
		return -1
	}

	optionsIds, err := Db.Query("SELECT id FROM option_tables WHERE question_id = ?", questionId)

	var optionSlice []int64
	if err != nil {
		log.Fatal(err.Error())
		return -1
	}

	for optionsIds.Next() {
		var value int64
		optionsIds.Scan(&value)
		optionSlice = append(optionSlice, value)
	}

	randomPosition := rand.IntN(int(questionCount))
	return optionSlice[randomPosition]
}
