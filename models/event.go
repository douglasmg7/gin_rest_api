package models

import (
	"time"

	"github.com/douglasmg7/gin_rest_api.git/db"
)

type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int
}

var events = []Event{}

func (e Event) Save() error {
	query := ` INSERT INTO events(name, description, location, dateTime, user_id)
				VALUES(?, ?, ?, ?, ?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		return err
	}

	_, err = result.LastInsertId()
	// id, err := result.LastInsertId()
	// e.ID = id

	return err
}

func GetAllEvents() []Event {
	// fmt.Println(events)
	return events
}
