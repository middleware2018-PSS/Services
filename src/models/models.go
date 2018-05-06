package models

import (
	"time"
)

type Subject = string

type Location = string

type Student struct {
	ID            int64          `json:"id"`
	Name          string         `json:"name"`
	Surname       string         `json:"surname"`
	Mail          string         `json:"mail"`
	Payments      []Payment      `json:"payments"`
	Grades        []Grade        `json:"grades"`
	Notifications []Notification `json:"notifications"`
	Classes       []Class        `json:"classes"` // enrolled -> classes
	Appointments  []Appointment  `json:"appointments"`
}

type Grade struct {
	Subject Subject   `json:"subject"`
	Date    time.Time `json:"date"`
	Grade   int       `json:"grade"`
}

type Appointment struct {
	Time     time.Time `json:"time"`
	Location Location  `json:"location"`
	Student  Student   `json:"student,omitempty"`
	Teacher  Teacher   `json:"student,omitempty"`
}

type Notification struct {
	ID       int64  `json:"id"`
	Receiver int64  `json:"receiver"`
	Message  string `json:"message"`
}

type Parents struct {
	ID       int64     `json:"id"`
	Name     string    `json:"name"`
	Surname  string    `json:"surname"`
	Mail     string    `json:"mail"`
	ParentOf []Student `json:"ParentOf"`
}

type Teacher struct {
	ID            int64               `json:"id"`
	Name          string              `json:"name"`
	Surname       string              `json:"surname"`
	Mail          string              `json:"mail"`
	Classes       map[Subject][]Class `json:"classes"`
	Appointments  []Appointment       `json:"appointments"`
	Lectures      []TimeTable         `json:"lectures"`
	Notifications []Notification      `json:"notifications"`
}

type TimeTable struct {
	Class    Class     `json:"class"`
	Location Location  `json:"location"`
	Date     time.Time `json:"date"`
	Start    time.Time `json:"start"`
	End      time.Time `json:"end"`
}

type Payment struct {
	ID      int64     `json:"id"`
	Amount  int64     `json:"amount"`
	Payed   bool      `json:"payed"`
	Emitted time.Time `json:"emitted"`
	Reason  string    `json:"reason"`
}

type Class struct {
	ID       int64     `json:"id"`
	Year     int       `json:"year"`
	Section  string    `json:"section"` // as "A" in 5'A
	Grade    int       `json:"grade"`   // as "5" in 5'A
	Info     string    `json:"info"`
	Students []Student `json:"students,omitempty"`
}
