package models

import (
	"fmt"
	"log"
)

type Major struct {
	Id   int64
	Name string
}

func (major *Major) InsertMajor() (*Major, error) {
	stmt, err := Db.Prepare("INSERT INTO majors (name) VALUES (?)")

	if err != nil {
		log.Println(err.Error())
		return nil, fmt.Errorf("FindUser: failed executing db query: %w", err)
	}

	result, errInsert := stmt.Exec(major.Name)

	if errInsert != nil {
		log.Println(errInsert.Error())
		return nil, fmt.Errorf("FindUser: failed executing db query: %w", err)
	}

	major.Id, _ = result.LastInsertId()

	return major, nil
}
