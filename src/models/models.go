package models

import (
	"time"
)

type Subject = string

type Location = string

// remove connections

type Student struct {
	ID      int64  `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
	Surname string `json:"surname,omitempty"`
	Mail    string `json:"mail,omitempty"`
	Info    string `json:"info,omitempty"`
}

type Grade struct {
	Student Student    `json:"student,omitempty"`
	Subject Subject    `json:"subject,omitempty"`
	Date    *time.Time `json:"date,omitempty"`
	Grade   int        `json:"grade,omitempty"`
	Teacher Teacher    `json:"teacher,omitempty"`
}

type Appointment struct {
	ID       int        `json:"id,omitempty"`
	Time     *time.Time `json:"time,omitempty"`
	Location Location   `json:"location,omitempty"`
	Student  Student    `json:"student,omitempty"`
	Teacher  Teacher    `json:"student,omitempty"`
}

type Notification struct {
	ID           int64      `json:"id,omitempty"`
	Receiver     int64      `json:"receiver,omitempty"`
	Time         *time.Time `json:"time,omitempty"`
	Message      string     `json:"message,omitempty"`
	ReceiverKind string     `json:"receiver_kind,omitempty"`
}

type Parent struct {
	ID      int64  `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
	Surname string `json:"surname,omitempty"`
	Mail    string `json:"mail,omitempty"`
	Info    string `json:"info,omitempty"`
}

type Teacher struct {
	ID      int64  `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
	Surname string `json:"surname,omitempty"`
	Mail    string `json:"mail,omitempty"`
}

type TimeTable struct {
	ID int64 `json:"id,omitempty"`
	Class    Class      `json:"class,omitempty"`
	Location Location   `json:"location,omitempty"`
	Subject  Subject    `json:"subject,omitempty"`
	Start    *time.Time `json:"start,omitempty"`
	End      *time.Time `json:"end,omitempty"`
	Info     string     `json:"info,omitempty"`
}

type Payment struct {
	ID      int64      `json:"id,omitempty"`
	Amount  int64      `json:"amount,omitempty"`
	Payed   bool       `json:"payed,omitempty"`
	Emitted *time.Time `json:"emitted,omitempty"`
	Reason  string     `json:"reason,omitempty"`
	Student Student    `json:"student,omitempty"`
}

type Class struct {
	ID      int64  `json:"id,omitempty"`
	Year    int    `json:"year,omitempty"`
	Section string `json:"section,omitempty"` // as "A" in 5'A
	Grade   int    `json:"grade,omitempty"`   // as "5" in 5'A
	Info    string `json:"info,omitempty"`
}
