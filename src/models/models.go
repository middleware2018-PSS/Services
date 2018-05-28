package models

import (
	"fmt"
	"github.com/phisco/hal"
	"time"
)

type (
	ID = int

	Subject = string

	Location = string

	Token struct {
		Code   int       `json:"code"`
		Token  string    `json:"token"`
		Expire time.Time `json:"expire"`
	}

	Login struct {
		Username string `form:"username" json:"username" binding:"required"`
		Password string `form:"password" json:"password" binding:"required"`
	}

	// remove connections
	Parent struct {
		ID      int     `json:"-" xml:"id" example:"1"`
		Name    *string `json:"name,omitempty" xml:"name"`
		Surname *string `json:"surname,omitempty" xml:"surname"`
		Mail    *string `json:"mail,omitempty" xml:"mail"`
		Info    *string `json:"info,omitempty" xml:"info"`
	}

	Teacher struct {
		ID      int     `json:"-" xml:"id" example:"1"`
		Name    *string `json:"name,omitempty" xml:"name"`
		Surname *string `json:"surname,omitempty" xml:"surname"`
		Mail    *string `json:"mail,omitempty" xml:"mail"`
		Info    *string `json:"info,omitempty" xml:"info"`
	}

	Student struct {
		ID      int     `json:"-" xml:"id" example:"1"`
		Name    *string `json:"name,omitempty" xml:"name"`
		Surname *string `json:"surname,omitempty" xml:"surname"`
		Mail    *string `json:"mail,omitempty" xml:"mail"`
		Info    *string `json:"info,omitempty" xml:"info"`
	}

	Class struct {
		ID      int     `json:"-" xml:"id" example:"1"`
		Year    *int    `json:"year,omitempty" xml:"year"`
		Section *string `json:"section,omitempty" xml:"section"` // as "A" in 5'A
		Grade   *int    `json:"grade,omitempty" xml:"grade"`     // as "5" in 5'A
		Info    *string `json:"info,omitempty" xml:"info"`
	}

	Notification struct {
		ID           int        `json:"-" xml:"id" example:"1"`
		Receiver     *int       `json:"receiver,omitempty" xml:"receiver"`
		Time         *time.Time `json:"time,omitempty" xml:"time"`
		Message      *string    `json:"message,omitempty" xml:"message"`
		ReceiverKind *string    `json:"receiver_kind,omitempty" xml:"receiver_kind"`
	}

	Appointment struct {
		ID       int        `json:"-" xml:"id" example:"1"`
		Time     *time.Time `json:"time,omitempty" xml:"time"`
		Location *string    `json:"location,omitempty" xml:"location" example:"Aula Magna"`
		Student  *int       `json:"studentID,omitempty" xml:"student"`
		Teacher  *int       `json:"teacherID,omitempty" xml:"teacher"`
	}

	Payment struct {
		ID      int        `json:"-" xml:"id" example:"1"`
		Amount  *int       `json:"amount,omitempty" xml:"amount"`
		Payed   *bool      `json:"payed,omitempty" xml:"payed"`
		Emitted *time.Time `json:"emitted,omitempty" xml:"emitted"`
		Reason  *string    `json:"reason,omitempty" xml:"reason"`
		Student *int       `json:"studentID,omitempty" xml:"student"`
	}

	Grade struct {
		ID      int        `json:"-" xml:"id" example:"1"`
		Student *int       `json:"studentID,omitempty" xml:"student"`
		Subject *string    `json:"subject,omitempty" xml:"subject" example:"science"`
		Date    *time.Time `json:"date,omitempty" xml:"date"`
		Grade   *int       `json:"grade,omitempty" xml:"grade"`
		Teacher *int       `json:"teacherID,omitempty" xml:"teacher"`
	}

	TimeTable struct {
		ID       int        `json:"-" xml:"id" example:"1"`
		Class    *int       `json:"classID,omitempty" xml:"class"`
		Location *string    `json:"location,omitempty" xml:"location" example:"Aula Magna"`
		Subject  *string    `json:"subject,omitempty" xml:"subject" example:"science"`
		Start    *time.Time `json:"start,omitempty" xml:"start"`
		End      *time.Time `json:"end,omitempty" xml:"end"`
		Info     *string    `json:"info,omitempty" xml:"info"`
	}

	Account struct {
		Username string `form:"username" json:"username,omitempty" binding:"required" example:"John"`
		Password string `form:"password" json:"password,omitempty" binding:"required" example:"Password"`
		Kind     string `json:"kind" binding:"required" example:"Parent"`
		ID       int    `json:"id" example:"1"`
	}

	User struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Role     string `json:"role"`
		ID       int    `json:"id"`
	}

	List struct {
		Self     string        `json:"self",xml:"self"`
		Data     []interface{} `json:"data,omitempty",xml:"data"`
		Next     string        `json:"next,omitempty",xml:"next"`
		Previous string        `json:"previous,omitempty",xml:"previous"`
	}

	Repr interface {
		GetRepresentation(bool) (interface{}, error)
	}

	Selfable interface {
		GetSelfLink() string
	}
)

func (r Parent) GetSelfLink() string {
	return "/parents/" + fmt.Sprintf("%d", r.ID)
}

func (r Parent) GetRepresentation(halF bool) (interface{}, error) {
	if !halF {
		return r, nil
	} else {
		self := r.GetSelfLink()
		h := hal.NewResource(r, self)
		h.AddNewLink("children", self+"/students")
		h.AddNewLink("appointments", self+"/appointments")
		h.AddNewLink("payments", self+"/payments")
		h.AddNewLink("notifications", self+"/notifications")
		return h, nil
	}
}

func (r Teacher) GetSelfLink() string {
	return "/teachers/" + fmt.Sprintf("%d", r.ID)
}

func (r Teacher) GetRepresentation(halF bool) (interface{}, error) {
	if !halF {
		return r, nil
	} else {
		self := r.GetSelfLink()
		h := hal.NewResource(r, self)
		h.AddNewLink("lectures", self+"/lectures")
		h.AddNewLink("appointments", self+"/appointments")
		h.AddNewLink("subjects", self+"/subjects")
		h.AddNewLink("notifications", self+"/notifications")
		h.AddNewLink("classes", self+"/classes")
		return h, nil
	}
}
func (r Student) GetSelfLink() string {
	return "/students/" + fmt.Sprintf("%d", r.ID)
}
func (r Student) GetRepresentation(halF bool) (interface{}, error) {
	if !halF {
		return r, nil
	} else {
		self := r.GetSelfLink()
		h := hal.NewResource(r, self)
		h.AddNewLink("grades", self+"/grades")
		return h, nil
	}
}

func (r Class) GetSelfLink() string {
	return "/classes/" + fmt.Sprintf("%d", r.ID)
}
func (r Class) GetRepresentation(halF bool) (interface{}, error) {
	if !halF {
		return r, nil
	} else {
		self := r.GetSelfLink()
		h := hal.NewResource(r, self)
		h.AddNewLink("students", self+"/students")
		return h, nil
	}
}

func (r Notification) GetSelfLink() string {
	return "/notifications/" + fmt.Sprintf("%d", r.ID)
}

func (r Notification) GetRepresentation(halF bool) (interface{}, error) {
	if !halF {
		return r, nil
	} else {
		self := r.GetSelfLink()
		h := hal.NewResource(r, self)
		return h, nil
	}
}

func (r Appointment) GetSelfLink() string {
	return "/appointments/" + fmt.Sprintf("%d", r.ID)
}

func (r Appointment) GetRepresentation(halF bool) (interface{}, error) {
	if !halF {
		return r, nil
	} else {
		self := r.GetSelfLink()
		h := hal.NewResource(r, self)
		h.AddNewLink("student", Student{ID: *r.Student}.GetSelfLink())
		h.AddNewLink("teacher", Teacher{ID: *r.Teacher}.GetSelfLink())
		return h, nil
	}
}

func (r Payment) GetSelfLink() string {
	return "/payments/" + fmt.Sprintf("%d", r.ID)
}

func (r Payment) GetRepresentation(halF bool) (interface{}, error) {
	if !halF {
		return r, nil
	} else {
		self := r.GetSelfLink()
		h := hal.NewResource(r, self)
		h.AddNewLink("student", Student{ID: *r.Student}.GetSelfLink())
		return h, nil
	}
}

func (r Grade) GetSelfLink() string {
	return "/grades/" + fmt.Sprintf("%d", r.ID)
}

func (r Grade) GetRepresentation(halF bool) (interface{}, error) {
	if !halF {
		return r, nil
	} else {
		self := r.GetSelfLink()
		h := hal.NewResource(r, self)
		h.AddNewLink("student", Student{ID: *r.Student}.GetSelfLink())
		h.AddNewLink("teacher", Teacher{ID: *r.Teacher}.GetSelfLink())
		return h, nil
	}
}

func (r TimeTable) GetSelfLink() string {
	return "/lectures/" + fmt.Sprintf("%d", r.ID)
}

func (r TimeTable) GetRepresentation(halF bool) (interface{}, error) {
	self := r.GetSelfLink()
	if !halF {
		return r, nil
	} else {
		h := hal.NewResource(r, self)
		h.AddNewLink("class", Class{ID: *r.Class}.GetSelfLink())
		return h, nil
	}
}
