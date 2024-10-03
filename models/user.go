package models

import (
	"errors"

	"github.com/douglasmg7/gin_rest_api.git/db"
	"github.com/douglasmg7/gin_rest_api.git/utils"
)

type User struct {
	ID       int64
	Email    string `binding: "required"`
	Password string `binding: "required"`
}

func (u *User) Save() error {
	query := ` INSERT INTO users (email, password) VALUES(?, ?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	result, err := stmt.Exec(u.Email, hashedPassword)
	if err != nil {
		return err
	}

	// _, err = result.LastInsertId()
	id, err := result.LastInsertId()
	u.ID = id

	return err
}

func (u *User) ValidateCredentials() error {
	query := "SELECT id, password FROM users WHERE email = ?"
	row := db.DB.QueryRow(query, u.Email)

	var retrivedPassword string
	err := row.Scan(&u.ID, &retrivedPassword)
	if err != nil {
		return errors.New("Credentials invalid")
	}

	if utils.CheckPassword(u.Password, retrivedPassword) {
		return nil
	}
	return errors.New("Credentials invalid")
}
