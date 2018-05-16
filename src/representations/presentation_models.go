package representations

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/middleware2018-PSS/Services/src/models"
	"log"
)

type Student struct {
	SelfAndData
	Grades string         `json:"grades",xml:"grades"`
}

type Parent struct {
	SelfAndData
	Children      string        `json:"children",xml:"children"`
	Appointments  string        `json:"appointments",xml:"appointments"`
	Payments      string        `json:"payments",xml:"payments"`
	Notifications string        `json:"notifications",xml:"notifications"`
}

type Teacher struct {
	SelfAndData
	Lectures      string         `json:"lectures",xml:"lectures"`
	Appointments  string         `json:"appointments",xml:"appointments"`
	Notifications string         `json:"notifications",xml:"notifications"`
	Subjects      string         `json:"subjects",xml:"subjects"`
	Classes       string         `json:"classes",xml:"classes"`
}

type SelfAndData struct {
	Self          string         `json:"self",xml:"self"`
	Data          interface{} `json:"data",xml:"data"`
}

type Class struct {
	SelfAndData
	Students string       `json:"students",xml:"students"`
}

type List struct {
	Self     string        `json:"self",xml:"self"`
	Data  []interface{} `json:"data",xml:"data"`
	Next     string        `json:"next,omitempty",xml:"next"`
	Previous string        `json:"previous,omitempty",xml:"previous"`
}

func ToRepresentation(res interface{}, c *gin.Context) (interface{}, error) {
	switch r := res.(type) {
	case models.Parent:
		self := "/parents/" + fmt.Sprintf("%d", r.ID)
		return &Parent{
			SelfAndData{self,
			r},
			self + "/students",
			self + "/appointments",
			self + "/payments",
			self + "/notifications",
		}, nil
	case models.Teacher:
		self := "/teachers/" + fmt.Sprintf("%d", r.ID)
		return &Teacher{
			SelfAndData{self,
			r},
			self + "/lectures",
			self + "/appointments",
			self + "/notifications",
			self + "/subjects",
			self + "/classes",
		}, nil
	case models.Student:
		self := "/students/" + fmt.Sprintf("%d", r.ID)
		return &Student{
			SelfAndData{self,
			r},
			self + "/grades",
		}, nil
	case models.Class:
		self := "/classes/" + fmt.Sprintf("%d", r.ID)
		return &Class{
			SelfAndData{self,
			r},
			self + "/students",
		}, nil
	case models.Notification:
		self := "/notifications/" + fmt.Sprintf("%d", r.ID)
		return &SelfAndData{
			self,
			r,
		}, nil
	case models.Appointment:
		self := "/appointments/" + fmt.Sprintf("%d", r.ID)
		s, _ := ToRepresentation(r.Student, c)
		t, _ := ToRepresentation(r.Teacher, c)

		return &SelfAndData{
			self,
			struct {
				models.Appointment
				Student *Student `json:"student",xml:"student"`
				Teacher *Teacher `json:"teacher",xml:"teacher"`
			}{r, s.(*Student), t.(*Teacher)}}, nil
	case models.Payment:
		self := "/payments/" + fmt.Sprintf("%d", r.ID)
		s, _ := ToRepresentation(r.Student, c)
		return &SelfAndData{
			self,
			struct {
				models.Payment
				*Student
			}{r, s.(*Student)}}, nil
	case models.Grade:
		self := "/payments/" + fmt.Sprintf("%d", r.ID)
		s, _ := ToRepresentation(r.Student, c)
		t, _ := ToRepresentation(r.Teacher, c)
		return &SelfAndData{
			self,
			struct {
				models.Grade
				*Student `json:"student"`
				*Teacher `json:"teacher"`
			}{r, s.(*Student), t.(*Teacher)},
		}, nil

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
