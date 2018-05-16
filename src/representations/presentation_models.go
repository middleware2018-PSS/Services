package representations

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/middleware2018-PSS/Services/src/models"
	"log"
	"strings"
)

type Student struct {
	Data   models.Student `json:"data",xml:"data"`
	Self   string         `json:"self",xml:"self"`
	Grades string         `json:"grades",xml:"grades"`
}

type Parent struct {
	Data          models.Parent `json:"data",xml:"data"`
	Childrens     string        `json:"childrens",xml:"childrens"`
	Self          string        `json:"self",xml:"self"`
	Appointments  string        `json:"appointments",xml:"appointments"`
	Payments      string        `json:"payments",xml:"payments"`
	Notifications string        `json:"notifications",xml:"notifications"`
}

type Teacher struct {
	Data          models.Teacher `json:"data", xml:"data"`
	Self          string         `json:"self",xml:"self"`
	Lectures      string         `json:"lectures",xml:"lectures"`
	Appointments  string         `json:"appointments",xml:"appointments"`
	Notifications string         `json:"notifications",xml:"notifications"`
	Subjects      string         `json:"subjects",xml:"subjects"`
	Classes       string         `json:"classes",xml:"classes"`
}

type List struct {
	Self    string        `json:"self",xml:"self"`
	Results []interface{} `json:"results", xml:"results"`
	Next    string        `json:"next,omitempty",xml:"next"`
}

type Class struct {
	Self     string       `json:"self",xml:"self"`
	Data     models.Class `json:"data", xml:"data"`
	Students string       `json:"students", xml:"students"`
}

func ToRepresentation(res interface{}, c *gin.Context) (interface{}, error) {
	switch r := res.(type) {
	case models.Parent:
		self := "/parents/" + fmt.Sprintf("%d", r.ID)
		return Parent{r,
			self + "/students",
			self,
			self + "/appointments",
			self + "/payments",
			self + "/notifications",
		}, nil
	case models.Teacher:
		self := "/teachers/" + fmt.Sprintf("%d", r.ID)
		return Teacher{r,
			self,
			self + "/lectures",
			self + "/appointments",
			self + "/notifications",
			self + "/subjects",
			self + "/classes"}, nil
	case models.Student:
		self := "/students/" + fmt.Sprintf("%d", r.ID)
		return Student{
			r,
			self,
			self + "/grades",
		}, nil
	case models.Class:
		self := "/classes/" + fmt.Sprintf("%d", r.ID)
		return Class{
			self,
			r,
			self + "/students",
		}, nil

	default:
		log.Print(fmt.Sprintf("implement representation for %T", res))
		return struct {
			Data interface{} `json:"data",xml:"data"`
			Self string      `json:"self",xml:"self"`
		}{res, c.Request.RequestURI}, nil
	}
}

func fixID(url string, ID models.ID) string {
	id := fmt.Sprintf("%d", ID)
	i := strings.LastIndex(url, "/")
	if url[i+1:] != id {
		url = url + "/" + id
	}
	return url
}
