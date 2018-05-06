package models

import (
	"time"
)

type Subject = string

type Location = string

type Student struct {
	ID            int64          `json:"id,omitempty"`
	Name          string         `json:"name,omitempty"`
	Surname       string         `json:"surname,omitempty"`
	Mail          string         `json:"mail,omitempty"`
	Payments      []Payment      `json:"payments,omitempty"`
	Grades        []Grade        `json:"grades,omitempty"`
	Notifications []Notification `json:"notifications,omitempty"`
	Classes       []Class        `json:"classes,omitempty"` // enrolled -> classes
	Appointments  []Appointment  `json:"appointments,omitempty"`
}

type Grade struct {
	Student Student    `json:"student,omitempty"`
	Subject Subject    `json:"subject,omitempty"`
	Date    *time.Time `json:"date,omitempty"`
	Grade   int        `json:"grade,omitempty"`
}

type Appointment struct {
	ID       int        `json:"id,omitempty"`
	Time     *time.Time `json:"time,omitempty"`
	Location Location   `json:"location,omitempty"`
	Student  Student    `json:"student,omitempty"`
	Teacher  Teacher    `json:"student,omitempty"`
}

type Notification struct {
	ID       int64  `json:"id,omitempty"`
	Receiver int64  `json:"receiver,omitempty"`
	Message  string `json:"message,omitempty"`
}

type Parent struct {
	ID            int64          `json:"id,omitempty"`
	Name          string         `json:"name,omitempty"`
	Surname       string         `json:"surname,omitempty"`
	Mail          string         `json:"mail,omitempty"`
	ParentOf      []Student      `json:"ParentOf,omitempty"`
	Payments      []Payment      `json:"payments,omitempty"`
	Notifications []Notification `json:"notification,omitempty"`
}

type Teacher struct {
	ID            int64               `json:"id,omitempty"`
	Name          string              `json:"name,omitempty"`
	Surname       string              `json:"surname,omitempty"`
	Mail          string              `json:"mail,omitempty"`
	Classes       map[Subject][]Class `json:"classes,omitempty"`
	Appointments  []Appointment       `json:"appointments,omitempty"`
	Lectures      []TimeTable         `json:"lectures,omitempty"`
	Notifications []Notification      `json:"notifications,omitempty"`
}

type TimeTable struct {
	Class    Class      `json:"class,omitempty"`
	Location Location   `json:"location,omitempty"`
	Subject Subject `json:"subject,omitempty"`
	Date     *time.Time `json:"date,omitempty"`
	Start    *time.Time `json:"start,omitempty"`
	End      *time.Time `json:"end,omitempty"`
}

type Payment struct {
	ID      int64      `json:"id,omitempty"`
	Amount  int64      `json:"amount,omitempty"`
	Payed   bool       `json:"payed,omitempty"`
	Emitted *time.Time `json:"emitted,omitempty"`
	Reason  string     `json:"reason,omitempty"`
}

type Class struct {
	ID       int64     `json:"id,omitempty"`
	Year     int       `json:"year,omitempty"`
	Section  string    `json:"section,omitempty"` // as "A" in 5'A
	Grade    int       `json:"grade,omitempty"`   // as "5" in 5'A
	Info     string    `json:"info,omitempty"`
	Students []Student `json:"students,omitempty"`
}
