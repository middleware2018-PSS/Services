package models

import (
	"time"
)

type ID = int64

type Subject = string

type Location = string

// remove connections

type Student struct {
	ID      int64      `json:"id" xml:"id" example:"1"`
	Name    *string `json:"name" xml:"name"`
	Surname *string `json:"surname" xml:"surname"`
	Mail    *string `json:"mail" xml:"mail"`
	Info    *string `json:"info" xml:"info"`
}

type Grade struct {
	ID      int64      `json:"id" xml:"id" example:"1"`
	Student Student    `json:"student" xml:"student"`
	Subject *string   `json:"subject" xml:"subject" example:"science"`
	Date    *time.Time `json:"date" xml:"date"`
	Grade   *int       `json:"grade" xml:"grade"`
	Teacher Teacher    `json:"teacher" xml:"teacher"`
}

type Appointment struct {
	ID      int64      `json:"id" xml:"id" example:"1"`
	Time     *time.Time `json:"time" xml:"time"`
	Location *string  `json:"location" xml:"location" example:"Aula Magna"`
	Student  Student    `json:"student" xml:"student"`
	Teacher  Teacher    `json:"student" xml:"teacher"`
}

type Notification struct {
	ID      int64      `json:"id" xml:"id" example:"1"`
	Receiver     *int64     `json:"receiver" xml:"receiver"`
	Time         *time.Time `json:"time" xml:"time"`
	Message      *string    `json:"message" xml:"message"`
	ReceiverKind *string    `json:"receiver_kind" xml:"receiver_kind"`
}

type Parent struct {
	ID      int64      `json:"id" xml:"id" example:"1"`
	Name    *string `json:"name" xml:"name"`
	Surname *string `json:"surname" xml:"surname"`
	Mail    *string `json:"mail" xml:"mail"`
	Info    *string `json:"info" xml:"info"`
}

type Teacher struct {
	ID      int64      `json:"id" xml:"id" example:"1"`
	Name    *string `json:"name" xml:"name"`
	Surname *string `json:"surname" xml:"surname"`
	Mail    *string `json:"mail" xml:"mail"`
	Info    *string `json:"info" xml:"info"`
}

type TimeTable struct {
	ID      int64      `json:"id" xml:"id" example:"1"`
	Class    Class      `json:"class" xml:"class"`
	Location *string  `json:"location" xml:"location" example:"Aula Magna"`
	Subject *string   `json:"subject" xml:"subject" example:"science"`
	Start    *time.Time `json:"start" xml:"start"`
	End      *time.Time `json:"end" xml:"end"`
	Info     *string    `json:"info" xml:"info"`
}

type Payment struct {
	ID      int64      `json:"id" xml:"id" example:"1"`
	Amount  *int       `json:"amount" xml:"amount"`
	Payed   *bool      `json:"payed" xml:"payed"`
	Emitted *time.Time `json:"emitted" xml:"emitted"`
	Reason  *string    `json:"reason" xml:"reason"`
	Student Student    `json:"student" xml:"student"`
}

type Class struct {
	ID      int64      `json:"id" xml:"id" example:"1"`
	Year    *int    `json:"year" xml:"year"`
	Section *string `json:"section" xml:"section"` // as "A" in 5'A
	Grade   *int    `json:"grade" xml:"grade"`     // as "5" in 5'A
	Info    *string `json:"info" xml:"info"`
}

type Account struct {
	Username string `form:"username" json:"username" binding:"required" example:"John"`
	Password string `form:"password" json:"password" binding:"required" example:"Password"`
}
