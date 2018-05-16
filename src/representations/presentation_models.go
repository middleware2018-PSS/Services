package representations

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/middleware2018-PSS/Services/src/models"
	"log"
)

type Student struct {
	Self   string         `json:"self",xml:"self"`
	Data   models.Student `json:"data",xml:"data"`
	Grades string         `json:"grades",xml:"grades"`
}

type Notification struct {
	Self string              `json:"self",xml:"self"`
	Data models.Notification `json:"data",xml:"data"`
}

type Parent struct {
	Self          string        `json:"self",xml:"self"`
	Data          models.Parent `json:"data",xml:"data"`
	Children      string        `json:"children",xml:"children"`
	Appointments  string        `json:"appointments",xml:"appointments"`
	Payments      string        `json:"payments",xml:"payments"`
	Notifications string        `json:"notifications",xml:"notifications"`
}

type Teacher struct {
	Self          string         `json:"self",xml:"self"`
	Data          models.Teacher `json:"data",xml:"data"`
	Lectures      string         `json:"lectures",xml:"lectures"`
	Appointments  string         `json:"appointments",xml:"appointments"`
	Notifications string         `json:"notifications",xml:"notifications"`
	Subjects      string         `json:"subjects",xml:"subjects"`
	Classes       string         `json:"classes",xml:"classes"`
}

type List struct {
	Self     string        `json:"self",xml:"self"`
	Results  []interface{} `json:"results",xml:"results"`
	Next     string        `json:"next,omitempty",xml:"next"`
	Previous string        `json:"previous,omitempty",xml:"previous"`
}

type Class struct {
	Self     string       `json:"self",xml:"self"`
	Data     models.Class `json:"data",xml:"data"`
	Students string       `json:"students",xml:"students"`
}

func ToRepresentation(res interface{}, c *gin.Context) (interface{}, error) {
	switch r := res.(type) {
	case models.Parent:
		self := "/parents/" + fmt.Sprintf("%d", r.ID)
		return Parent{
			self,
			r,
			self + "/students",
			self + "/appointments",
			self + "/payments",
			self + "/notifications",
		}, nil
	case models.Teacher:
		self := "/teachers/" + fmt.Sprintf("%d", r.ID)
		return Teacher{
			self,
			r,
			self + "/lectures",
			self + "/appointments",
			self + "/notifications",
			self + "/subjects",
			self + "/classes",
		}, nil
	case models.Student:
		self := "/students/" + fmt.Sprintf("%d", r.ID)
		return Student{
			self,
			r,
			self + "/grades",
		}, nil
	case models.Class:
		self := "/classes/" + fmt.Sprintf("%d", r.ID)
		return Class{
			self,
			r,
			self + "/students",
		}, nil
	case models.Notification:
		self := "/notifications/" + fmt.Sprintf("%d", r.ID)
		return Notification{
			self,
			r,
		}, nil

	default:
		log.Fatal(fmt.Sprintf("implement representation for %T", res))
		return struct {
			Self string      `json:"self",xml:"self"`
			Data interface{} `json:"data",xml:"data"`
		}{
			c.Request.RequestURI,
			res,
		}, nil
	}
}
