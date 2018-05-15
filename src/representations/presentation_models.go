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
	Results []interface{} `json:"results", xml:"results"`
	Next    string        `json:"next,omitempty",xml:"next"`
}

func ToRepresentation(res interface{}, c *gin.Context) (interface{}, error) {
	url := c.Request.RequestURI
	switch r := res.(type) {
	case models.Parent:

		return Parent{r,
			url + "/students",
			url,
			url + "/appointments",
			url + "/payments",
			url + "/notifications",
		}, nil
	case models.Teacher:
		return Teacher{r, url,
			url + "/lectures",
			url + "/appointments",
			url + "/notifications",
			url + "/subjects",
			url + "/classes"}, nil
	case models.Student:
		url = fixID(url, r.ID)
		return Student{
			r, url,
			url + "/grades",
		}, nil

	default:
		log.Print(fmt.Sprintf("implement representation for %T", res))
		return struct {
			Data interface{} `json:"data",xml:"data"`
			Self string      `json:"self",xml:"self"`
		}{res, url}, nil
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
