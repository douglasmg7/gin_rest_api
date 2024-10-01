package models

import (
	"github.com/douglasmg7/gin_rest_api.git/db"
)

type User struct {
	ID       int64
	Email    string `binding: "required"`
	Password string `binding: "required"`
}

func (u User) Save() error {
	query := ` INSERT INTO users (email, password) VALUES(?, ?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(u.Email, u.Password)
	if err != nil {
		return err
	}

	// _, err = result.LastInsertId()
	id, err := result.LastInsertId()
	u.ID = id

	return err
}
