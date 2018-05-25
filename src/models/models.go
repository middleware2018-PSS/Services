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

	// remove connections
	Parent struct {
		ID      int     `json:"-,omitempty" xml:"id" example:"1"`
		Name    *string `json:"name,omitempty" xml:"name"`
		Surname *string `json:"surname,omitempty" xml:"surname"`
		Mail    *string `json:"mail,omitempty" xml:"mail"`
		Info    *string `json:"info,omitempty" xml:"info"`
	}

	Teacher struct {
		ID      int     `json:"-,omitempty" xml:"id" example:"1"`
		Name    *string `json:"name,omitempty" xml:"name"`
		Surname *string `json:"surname,omitempty" xml:"surname"`
		Mail    *string `json:"mail,omitempty" xml:"mail"`
		Info    *string `json:"info,omitempty" xml:"info"`
	}

	Student struct {
		ID      int     `json:"-,omitempty" xml:"id" example:"1"`
		Name    *string `json:"name,omitempty" xml:"name"`
		Surname *string `json:"surname,omitempty" xml:"surname"`
		Mail    *string `json:"mail,omitempty" xml:"mail"`
		Info    *string `json:"info,omitempty" xml:"info"`
	}

	Class struct {
		ID      int     `json:"-,omitempty" xml:"id" example:"1"`
		Year    *int    `json:"year,omitempty" xml:"year"`
		Section *string `json:"section,omitempty" xml:"section"` // as "A" in 5'A
		Grade   *int    `json:"grade,omitempty" xml:"grade"`     // as "5" in 5'A
		Info    *string `json:"info,omitempty" xml:"info"`
	}

	Notification struct {
		ID           int        `json:"-,omitempty" xml:"id" example:"1"`
		Receiver     *int       `json:"receiver,omitempty" xml:"receiver"`
		Time         *time.Time `json:"time,omitempty" xml:"time"`
		Message      *string    `json:"message,omitempty" xml:"message"`
		ReceiverKind *string    `json:"receiver_kind,omitempty" xml:"receiver_kind"`
	}

	Appointment struct {
		ID       int        `json:"-,omitempty" xml:"id" example:"1"`
		Time     *time.Time `json:"time,omitempty" xml:"time"`
		Location *string    `json:"location,omitempty" xml:"location" example:"Aula Magna"`
		Student  Student    `json:"student,omitempty" xml:"student"`
		Teacher  Teacher    `json:"student,omitempty" xml:"teacher"`
	}

	Payment struct {
		ID      int        `json:"-,omitempty" xml:"id" example:"1"`
		Amount  *int       `json:"amount,omitempty" xml:"amount"`
		Payed   *bool      `json:"payed,omitempty" xml:"payed"`
		Emitted *time.Time `json:"emitted,omitempty" xml:"emitted"`
		Reason  *string    `json:"reason,omitempty" xml:"reason"`
		Student Student    `json:"student,omitempty" xml:"student"`
	}

	Grade struct {
		ID      int        `json:"-,omitempty" xml:"id" example:"1"`
		Student Student    `json:"student,omitempty" xml:"student"`
		Subject *string    `json:"subject,omitempty" xml:"subject" example:"science"`
		Date    *time.Time `json:"date,omitempty" xml:"date"`
		Grade   *int       `json:"grade,omitempty" xml:"grade"`
		Teacher Teacher    `json:"teacher,omitempty" xml:"teacher"`
	}

	TimeTable struct {
		ID       int        `json:"-,omitempty" xml:"id" example:"1"`
		Class    Class      `json:"class,omitempty" xml:"class"`
		Location *string    `json:"location,omitempty" xml:"location" example:"Aula Magna"`
		Subject  *string    `json:"subject,omitempty" xml:"subject" example:"science"`
		Start    *time.Time `json:"start,omitempty" xml:"start"`
		End      *time.Time `json:"end,omitempty" xml:"end"`
		Info     *string    `json:"info,omitempty" xml:"info"`
	}

	Account struct {
		Username string `form:"username" json:"username,omitempty" binding:"required" example:"John"`
		Password string `form:"password" json:"password,omitempty" binding:"required" example:"Password"`
	}
)

func (r TimeTable) GetMap() hal.Entry {
	return hal.Entry{
		"location": r.Location,
		"subject":  r.Subject,
		"start":    r.Start,
		"end":      r.End,
		"info":     r.Info,
	}
}

func (r Appointment) GetMap() hal.Entry {
	return hal.Entry{
		"time":     r.Time,
		"location": r.Location,
	}
}

func (r Grade) GetMap() hal.Entry {
	return hal.Entry{
		"grade":   r.Grade,
		"subject": r.Subject,
		"date":    r.Date,
	}
}

func (r List) GetMap() hal.Entry {
	hen := hal.Entry{}
	if r.Previous != "" {
		hen["prev"] = r.Previous
	}
	if r.Next != "" {
		hen["next"] = r.Next
	}
	return hen
}

type (
	StudentRepr struct {
		SelfAndData
		Grades string `json:"grades",xml:"grades"`
	}

	ParentRepr struct {
		SelfAndData
		Children      string `json:"children",xml:"children"`
		Appointments  string `json:"appointments",xml:"appointments"`
		Payments      string `json:"payments",xml:"payments"`
		Notifications string `json:"notifications",xml:"notifications"`
	}

	TeacherRepr struct {
		SelfAndData
		Lectures      string `json:"lectures",xml:"lectures"`
		Appointments  string `json:"appointments",xml:"appointments"`
		Notifications string `json:"notifications",xml:"notifications"`
		Subjects      string `json:"subjects",xml:"subjects"`
		Classes       string `json:"classes",xml:"classes"`
	}

	SelfAndData struct {
		Self string      `json:"self",xml:"self"`
		Data interface{} `json:"data",xml:"data"`
	}

	ClassRepr struct {
		SelfAndData
		Students string `json:"students",xml:"students"`
	}

	List struct {
		Self     string        `json:"self",xml:"self"`
		Data     []interface{} `json:"data",xml:"data"`
		Next     string        `json:"next,omitempty",xml:"next"`
		Previous string        `json:"previous,omitempty",xml:"previous"`
	}

	Repr interface {
		GetRepresentation(bool) (interface{}, error)
	}
)

func (r Parent) GetRepresentation(halF bool) (interface{}, error) {
	self := "/parents/" + fmt.Sprintf("%d", r.ID)
	if !halF {
		return &ParentRepr{
			SelfAndData{self,
				r},
			self + "/students",
			self + "/appointments",
			self + "/payments",
			self + "/notifications",
		}, nil
	} else {
		h := hal.NewResource(r, self)
		h.AddNewLink("children", self+"/students")
		h.AddNewLink("appointments", self+"/appointments")
		h.AddNewLink("payments", self+"/payments")
		h.AddNewLink("notifications", self+"/notifications")
		return h, nil
	}
}

func (r Teacher) GetRepresentation(halF bool) (interface{}, error) {
	self := "/teachers/" + fmt.Sprintf("%d", r.ID)
	if !halF {
		return &TeacherRepr{
			SelfAndData{self,
				r},
			self + "/lectures",
			self + "/appointments",
			self + "/notifications",
			self + "/subjects",
			self + "/classes",
		}, nil
	} else {
		h := hal.NewResource(r, self)
		h.AddNewLink("lectures", self+"/lectures")
		h.AddNewLink("appointments", self+"/appointments")
		h.AddNewLink("subjects", self+"/subjects")
		h.AddNewLink("notifications", self+"/notifications")
		h.AddNewLink("classes", self+"/classes")
		return h, nil
	}
}

func (r Student) GetRepresentation(halF bool) (interface{}, error) {
	self := "/students/" + fmt.Sprintf("%d", r.ID)
	if !halF {
		return &StudentRepr{
			SelfAndData{self,
				r},
			self + "/grades",
		}, nil
	} else {
		h := hal.NewResource(r, self)
		h.AddNewLink("grades", self+"/grades")
		return h, nil
	}
}
func (r Class) GetRepresentation(halF bool) (interface{}, error) {
	self := "/classes/" + fmt.Sprintf("%d", r.ID)
	if !halF {
		return &ClassRepr{
			SelfAndData{self,
				r},
			self + "/students",
		}, nil
	} else {
		h := hal.NewResource(r, self)
		h.AddNewLink("students", self+"/students")
		return h, nil
	}
}

func (r Notification) GetRepresentation(halF bool) (interface{}, error) {
	self := "/notifications/" + fmt.Sprintf("%d", r.ID)
	if !halF {
		return &SelfAndData{
			self,
			r,
		}, nil
	} else {
		h := hal.NewResource(r, self)
		return h, nil
	}
}
func (r Appointment) GetRepresentation(halF bool) (interface{}, error) {
	self := "/appointments/" + fmt.Sprintf("%d", r.ID)
	s, _ := r.Student.GetRepresentation(false)
	t, _ := r.Teacher.GetRepresentation(false)
	if !halF {
		return &SelfAndData{
			self,
			struct {
				Appointment
				Student *StudentRepr `json:"student",xml:"student"`
				Teacher *TeacherRepr `json:"teacher",xml:"teacher"`
			}{r, s.(*StudentRepr), t.(*TeacherRepr)}}, nil
	} else {
		h := hal.NewResource(r, self)
		h.AddNewLink("student", s.(*StudentRepr).Self)
		h.AddNewLink("teacher", t.(*TeacherRepr).Self)
		return h, nil
	}
}
func (r Payment) GetRepresentation(halF bool) (interface{}, error) {
	self := "/payments/" + fmt.Sprintf("%d", r.ID)
	s, _ := r.Student.GetRepresentation(false)
	if !halF {
		return &SelfAndData{
			self,
			struct {
				Payment
				*StudentRepr
			}{r, s.(*StudentRepr)}}, nil
	} else {
		h := hal.NewResource(r, self)
		h.AddNewLink("student", s.(*StudentRepr).Self)
		return h, nil
	}
}
func (r Grade) GetRepresentation(halF bool) (interface{}, error) {
	self := "/grades/" + fmt.Sprintf("%d", r.ID)
	s, _ := r.Student.GetRepresentation(false)
	t, _ := r.Teacher.GetRepresentation(false)
	if !halF {
		return &SelfAndData{
			self,
			struct {
				Grade
				*StudentRepr `json:"student"`
				*TeacherRepr `json:"teacher"`
			}{r, s.(*StudentRepr), t.(*TeacherRepr)},
		}, nil
	} else {
		h := hal.NewResource(r, self)
		h.AddNewLink("student", s.(*StudentRepr).Self)
		h.AddNewLink("teacher", t.(*TeacherRepr).Self)
		return h, nil
	}
}
func (r TimeTable) GetRepresentation(halF bool) (interface{}, error) {
	self := "/lectures/" + fmt.Sprintf("%d", r.ID)
	class, _ := r.Class.GetRepresentation(false)
	if !halF {
		return &SelfAndData{
			self,
			struct {
				TimeTable
				*ClassRepr `json:"class"`
			}{r, class.(*ClassRepr)},
		}, nil
	} else {
		h := hal.NewResource(r, self)
		h.AddNewLink("class", class.(*ClassRepr).Self)
		return h, nil
	}
}
