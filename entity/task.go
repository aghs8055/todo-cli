package entity

import (
	"time"
)

type Status int8

const (
	Waiting = Status(iota + 1)
	Doing
	Done
	Missed
)

type Task struct {
	ID          int64
	Title       string
	Description string
	Status      Status
	CreatedAt   time.Time
	Deadline    time.Time
	CategoryID  int64
}

func (s Status) GetTitle() string {
	switch s {
	case 1:
		return "Waiting"
	case 2:
		return "Doing"
	case 3:
		return "Done"
	case 4:
		return "Missed"
	default:
		return "Invalid status"
	}
}
