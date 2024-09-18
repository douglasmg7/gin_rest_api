package models

import "time"

type Event struct {
	ID          int
	Name        string
	Description string
	Place       string
	DateTime    time.Time
	UserID      int
}

var events = []*Event{}

func (e *Event) Save() {
	events = append(events, e)
}
