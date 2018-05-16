package models

import (
	"time"
)

type ID = int64

type Subject = string

type Location = string

// remove connections

type Student struct {
	ID      ID      `json:"id",xml:"id"`
	Name    *string `json:"name",xml:"name"`
	Surname *string `json:"surname",xml:"surname"`
	Mail    *string `json:"mail",xml:"mail"`
	Info    *string `json:"info",xml:"info"`
}

type Grade struct {
	ID ID `json:"id",xml:"id"`
	Student Student    `json:"student",xml:"student"`
	Subject *Subject   `json:"subject",xml:"subject"`
	Date    *time.Time `json:"date",xml:"date"`
	Grade   *int       `json:"grade",xml:"grade"`
	Teacher Teacher   `json:"teacher",xml:"teacher"`
}

type Appointment struct {
	ID       ID         `json:"id",xml:"id"`
	Time     *time.Time `json:"time",xml:"time"`
	Location *Location  `json:"location",xml:"location"`
	Student  Student   `json:"student",xml:"student"`
	Teacher  Teacher   `json:"student",xml:"teacher"`
}

type Notification struct {
	ID           ID         `json:"id",xml:"id"`
	Receiver     *int64     `json:"receiver",xml:"receiver"`
	Time         *time.Time `json:"time",xml:"time"`
	Message      *string    `json:"message",xml:"message"`
	ReceiverKind *string    `json:"receiver_kind",xml:"receiver_kind"`
}

type Parent struct {
	ID      ID      `json:"id",xml:"id"`
	Name    *string `json:"name",xml:"name"`
	Surname *string `json:"surname",xml:"surname"`
	Mail    *string `json:"mail",xml:"mail"`
	Info    *string `json:"info",xml:"info"`
}

type Teacher struct {
	ID      ID      `json:"id",xml:"id"`
	Name    *string `json:"name",xml:"name"`
	Surname *string `json:"surname",xml:"surname"`
	Mail    *string `json:"mail",xml:"mail"`
	Info    *string `json:"info",xml:"info"`
}

type TimeTable struct {
	ID       ID         `json:"id",xml:"id"`
	Class    Class     `json:"class",xml:"class"`
	Location *Location  `json:"location",xml:"location"`
	Subject  *Subject   `json:"subject",xml:"subject"`
	Start    *time.Time `json:"start",xml:"start"`
	End      *time.Time `json:"end",xml:"end"`
	Info     *string    `json:"info",xml:"info"`
}

type Payment struct {
	ID      ID         `json:"id",xml:"id"`
	Amount  *int    `json:"amount",xml:"amount"`
	Payed   *bool      `json:"payed",xml:"payed"`
	Emitted *time.Time `json:"emitted",xml:"emitted"`
	Reason  *string    `json:"reason",xml:"reason"`
	Student Student   `json:"student",xml:"student"`
}

type Class struct {
	ID      ID      `json:"id",xml:"id"`
	Year    *int    `json:"year",xml:"year"`
	Section *string `json:"section",xml:"section"` // as "A" in 5'A
	Grade   *int    `json:"grade",xml:"grade"`   // as "5" in 5'A
	Info    *string `json:"info",xml:"info"`
}
