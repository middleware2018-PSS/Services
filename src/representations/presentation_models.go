package representations

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/middleware2018-PSS/Services/src/models"
	"log"
	"github.com/nvellon/hal"
)
const HAL = "application/hal+json"


type Student struct {
	SelfAndData
	Grades string `json:"grades",xml:"grades"`
}

type Parent struct {
	SelfAndData
	Children      string `json:"children",xml:"children"`
	Appointments  string `json:"appointments",xml:"appointments"`
	Payments      string `json:"payments",xml:"payments"`
	Notifications string `json:"notifications",xml:"notifications"`
}

type Teacher struct {
	SelfAndData
	Lectures      string `json:"lectures",xml:"lectures"`
	Appointments  string `json:"appointments",xml:"appointments"`
	Notifications string `json:"notifications",xml:"notifications"`
	Subjects      string `json:"subjects",xml:"subjects"`
	Classes       string `json:"classes",xml:"classes"`
}

type SelfAndData struct {
	Self string      `json:"self",xml:"self"`
	Data interface{} `json:"data",xml:"data"`
}

type Class struct {
	SelfAndData
	Students string `json:"students",xml:"students"`
}

type List struct {
	Self     string        `json:"self",xml:"self"`
	Data     []interface{} `json:"data",xml:"data"`
	Next     string        `json:"next,omitempty",xml:"next"`
	Previous string        `json:"previous,omitempty",xml:"previous"`
}




func ToRepresentation(res interface{}, c *gin.Context, halF bool) (interface{}, error) {
	switch r := res.(type) {
	case models.Parent:
		self := "/parents/" + fmt.Sprintf("%d", r.ID)
		if !halF{
			return &Parent{
			SelfAndData{self,
				r},
			self + "/students",
			self + "/appointments",
			self + "/payments",
			self + "/notifications",
			}, nil
		} else {
			h := hal.NewResource(r, self)
			h.AddNewLink("children", self + "/students")
			h.AddNewLink("appointments", self + "/appointments")
			h.AddNewLink("payments", self + "/payments")
			h.AddNewLink("notifications", self + "/notifications")
			return h, nil
		}
	case models.Teacher:
		self := "/teachers/" + fmt.Sprintf("%d", r.ID)
		if !halF {
			return &Teacher{
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
			h.AddNewLink("lectures", self + "/lectures")
			h.AddNewLink("appointments", self + "/appointments")
			h.AddNewLink("subjects", self + "/subjects")
			h.AddNewLink("notifications", self + "/notifications")
			h.AddNewLink("classes", self + "/classes")
			return h, nil
		}
	case models.Student:
		self := "/students/" + fmt.Sprintf("%d", r.ID)
		if !halF {
			return &Student{
				SelfAndData{self,
					r},
				self + "/grades",
			}, nil
		} else {
			h := hal.NewResource(r, self)
			h.AddNewLink("grades", self + "/grades")
			return h, nil
		}
	case models.Class:
		self := "/classes/" + fmt.Sprintf("%d", r.ID)
		if !halF {
			return &Class{
				SelfAndData{self,
					r},
				self + "/students",
			}, nil
		} else {
			h := hal.NewResource(r, self)
			h.AddNewLink("students", self + "/students")
			return h, nil
		}
	case models.Notification:
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
	case models.Appointment:
		self := "/appointments/" + fmt.Sprintf("%d", r.ID)
		s, _ := ToRepresentation(r.Student, c, false)
		t, _ := ToRepresentation(r.Teacher, c, false)
		if !halF {
			return &SelfAndData{
				self,
				struct {
					models.Appointment
					Student *Student `json:"student",xml:"student"`
					Teacher *Teacher `json:"teacher",xml:"teacher"`
				}{r, s.(*Student), t.(*Teacher)}}, nil
		} else {
			h := hal.NewResource(r, self)
			h.AddNewLink("student", s.(*Student).Self)
			h.AddNewLink("teacher", t.(*Teacher).Self)
			return h, nil
		}
	case models.Payment:
		self := "/payments/" + fmt.Sprintf("%d", r.ID)
		s, _ := ToRepresentation(r.Student, c, halF)
		if !halF {
			return &SelfAndData{
				self,
				struct {
					models.Payment
					*Student
				}{r, s.(*Student)}}, nil
		} else {
			h := hal.NewResource(r, self)
			h.AddNewLink("student", s.(*Student).Self)
			return h, nil
		}
	case models.Grade:
		self := "/grades/" + fmt.Sprintf("%d", r.ID)
		s, _ := ToRepresentation(r.Student, c, false)
		t, _ := ToRepresentation(r.Teacher, c, false)
		if !halF {
			return &SelfAndData{
				self,
				struct {
					models.Grade
					*Student `json:"student"`
					*Teacher `json:"teacher"`
				}{r, s.(*Student), t.(*Teacher)},
			}, nil
		} else {
			h := hal.NewResource(r, self)
			h.AddNewLink("student", s.(*Student).Self)
			h.AddNewLink("teacher", t.(*Teacher).Self)
			return h, nil
		}
	case models.TimeTable:
		self := "/lectures/" + fmt.Sprintf("%d", r.ID)
		class, _ := ToRepresentation(r.Class, c, false)
		if !halF {
			return &SelfAndData{
				self,
				struct {
					models.TimeTable
					*Class `json:"class"`
				}{r, class.(*Class)},
			}, nil
		} else {
			h := hal.NewResource(r,self)
			h.AddNewLink("class", class.(*Class).Self)
			return h, nil
		}
	// TODO timetable endpoint
	default:
		log.Fatal(fmt.Sprintf("implement representation for %T", res))
		return &struct {
			Self string      `json:"self",xml:"self"`
			Data interface{} `json:"data",xml:"data"`
		}{
			c.Request.RequestURI,
			res,
		}, nil
	}
}