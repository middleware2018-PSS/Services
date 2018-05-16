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
	Name    *string `json:"name"`
	Surname *string `json:"surname"`
	Mail    *string `json:"mail"`
	Info    *string `json:"info"`
}

type Grade struct {
	Student Student    `json:"student"`
	Subject *Subject   `json:"subject"`
	Date    *time.Time `json:"date"`
	Grade   *int       `json:"grade"`
	Teacher *Teacher   `json:"teacher"`
}

type Appointment struct {
	ID       ID         `json:"id",xml:"id"`
	Time     *time.Time `json:"time"`
	Location *Location  `json:"location"`
	Student  *Student   `json:"student"`
	Teacher  *Teacher   `json:"student"`
}

type Notification struct {
	ID           ID         `json:"id",xml:"id"`
	Receiver     *int64     `json:"receiver"`
	Time         *time.Time `json:"time"`
	Message      *string    `json:"message"`
	ReceiverKind *string    `json:"receiver_kind"`
}

type Parent struct {
	ID      ID      `json:"id",xml:"id"`
	Name    *string `json:"name"`
	Surname *string `json:"surname"`
	Mail    *string `json:"mail"`
	Info    *string `json:"info"`
}

type Teacher struct {
	ID      ID      `json:"id",xml:"id"`
	Name    *string `json:"name"`
	Surname *string `json:"surname"`
	Mail    *string `json:"mail"`
	Info    *string `json:"info"`
}

type TimeTable struct {
	ID       ID         `json:"id",xml:"id"`
	Class    *Class     `json:"class"`
	Location *Location  `json:"location"`
	Subject  *Subject   `json:"subject"`
	Start    *time.Time `json:"start"`
	End      *time.Time `json:"end"`
	Info     *string    `json:"info"`
}

type Payment struct {
	ID      ID         `json:"id",xml:"id"`
	Amount  *int64     `json:"amount"`
	Payed   *bool      `json:"payed"`
	Emitted *time.Time `json:"emitted"`
	Reason  *string    `json:"reason"`
	Student *Student   `json:"student"`
}

type Class struct {
	ID      ID      `json:"id",xml:"id"`
	Year    *int    `json:"year"`
	Section *string `json:"section"` // as "A" in 5'A
	Grade   *int    `json:"grade"`   // as "5" in 5'A
	Info    *string `json:"info"`
}
