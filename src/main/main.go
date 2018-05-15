package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/middleware2018-PSS/Services/src/controller"
	"github.com/middleware2018-PSS/Services/src/models"
	"github.com/middleware2018-PSS/Services/src/repository"
	"github.com/middleware2018-PSS/Services/src/representations"
	"github.com/pkg/errors"
	"log"
	"net/http"
	"strconv"
	"strings"
)

var (
	LimitError = errors.New("Limit Must Be Greater Than Zero.")
)

func main() {
	db, err := sql.Open("postgres", "user=postgres dbname=postgres sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	r := repository.NewPostgresRepository(db)
	con := controller.NewController(r)

	api := gin.Default()
	// TODO implement AUTH with data from db
	// TODO implement AUTH instead of relying on url
	parent := api.Group("/parents/:id") //, gin.BasicAuth(gin.Accounts{"3": "prova"}), Access())
	{
		parent.GET("", byID("id", con.ParentByID))
		parent.PUT("", func(c *gin.Context) {
			// not possible to refactor (at the best of my knowledge)
			var p models.Parent
			if err := c.ShouldBind(&p); err == nil {
				id, _ := strconv.Atoi(c.Param("id"))
				p.ID = int64(id)
				if err := con.UpdateParent(p); err == nil {
					c.JSON(http.StatusCreated, p)
				}
			}

		})
		parent.GET("/students", byIDWithOffsetAndLimit("id", con.ChildrenByParent))
		// TODO redirect or deeper links? e.g. /parents/id/students/student/grades...
		parent.GET("/students/:student", byID("student", con.StudentByID))
		parent.GET("/students/:student/grades", byIDWithOffsetAndLimit("student", con.GradesByStudent))
		parent.GET("/appointments", byIDWithOffsetAndLimit("id", con.AppointmentsByParent))
		parent.GET("/payments", byIDWithOffsetAndLimit("id", con.PaymentsByParent))
		parent.GET("/notifications", byIDWithOffsetAndLimit("id", con.NotificationsByParent))
	}

	// TODO add hypermedia
	teachers := api.Group("/teachers/:id") //, gin.BasicAuth(gin.Accounts{"1": "prova"}), Access())
	{
		teachers.GET("", byID("id", con.TeacherByID))
		teachers.PUT("", func(c *gin.Context) {
			var t models.Teacher
			if err := c.ShouldBind(&t); err == nil {
				id, _ := strconv.Atoi(c.Param("id"))
				t.ID = int64(id)
				if err := con.UpdateTeacher(t); err != nil {
					c.JSON(http.StatusNoContent, nil)
				}
			}

		})
		teachers.GET("/lectures", byIDWithOffsetAndLimit("id", con.LecturesByTeacher))
		teachers.GET("/appointments", byIDWithOffsetAndLimit("id", con.AppointmentsByTeacher))
		teachers.GET("/notifications", byIDWithOffsetAndLimit("id", con.NotificationsByTeacher))
		teachers.GET("/subjects", byIDWithOffsetAndLimit("id", con.SubjectsByTeacher))
		teachers.GET("/subjects/:subject", func(c *gin.Context) {
			id, err := strconv.Atoi(c.Param("id"))
			subj := c.Param("subject")
			offset, err := strconv.Atoi(c.DefaultQuery("offset", "0"))
			limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
			res, err := con.ClassesBySubjectAndTeacher(int64(id), subj, limit, offset)
			handleErr(err, res, c)
		})
		teachers.GET("/classes", byIDWithOffsetAndLimit("id", con.ClassesByTeacher))
	}

	// TODO remove admin from path and use token
	admin := api.Group("/admins/:id", gin.BasicAuth(gin.Accounts{"1": "prova"}), Access())

	admin.GET("/parents", getOffsetLimit(con.Parents))
	admin.GET("/parents/:parent", byID("parent", con.ParentByID))

	admin.GET("/students", getOffsetLimit(con.Students))
	api.GET("/students/:student/grades", byIDWithOffsetAndLimit("student", con.GradesByStudent))

	admin.GET("/notifications", getOffsetLimit(con.Notifications))
	admin.GET("/notifications/:notification", byID("notification", con.NotificationByID))

	admin.GET("/payments", getOffsetLimit(con.Payments))
	admin.GET("/payments/:payment", byID("payment", con.PaymentByID))

	admin.GET("/teachers", getOffsetLimit(con.Teachers))

	admin.GET("/classes", getOffsetLimit(con.Classes))
	admin.GET("/classes/:class", byID("class", con.ClassByID))
	admin.GET("/classes/:class/students", byIDWithOffsetAndLimit("class", con.StudentsByClass))

	api.Run(":5000")
}

func byID(key string, f func(int64) (interface{}, error)) func(c *gin.Context) {
	return func(c *gin.Context) {
		// TODO: check err
		id, err := strconv.Atoi(c.Param(key))
		res, err := f(int64(id))
		res, _ = representations.ToRepresentation(res, c)
		handleErr(err, res, c)
	}

}

func offsetLimit(c *gin.Context) (int, int) {
	// TODO check err
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	return offset, limit
}

func getOffsetLimit(f func(int, int) ([]interface{}, error)) func(c *gin.Context) {
	return func(c *gin.Context) {
		//TODO Check id and errors
		offset, limit := offsetLimit(c)
		if limit > 0 {
			res, err := f(limit, offset)
			handleErr(err, res, c)
		} else {
			handleErr(LimitError, nil, c)
		}

	}
}

func byIDWithOffsetAndLimit(id string, f func(int64, int, int) ([]interface{}, error)) func(c *gin.Context) {
	return func(c *gin.Context) {
		//TODO Check id and errors (4 real)
		id, err := strconv.Atoi(c.Param(id))
		offset, limit := offsetLimit(c)
		res, err := f(int64(id), limit, offset)
		for i, el := range res {
			res[i], _ = representations.ToRepresentation(el, c)
		}
		result := representations.List{res, next(c.Request.RequestURI, offset, limit, res)}
		handleErr(err, result, c)
	}
}

func next(uri string, offset int, limit int, input []interface{}) (res string) {
	if n := strings.Index(uri, "?"); n >= 0 {
		res = uri[:n]
	} else {
		res = uri
	}
	if l := len(input); l < limit {
		return ""
	}
	return strings.Join([]string{res, fmt.Sprintf("?offset=%d&limit=%d", offset+limit, limit)}, "")
}

func Access() gin.HandlerFunc {
	return func(c *gin.Context) {
		if id := c.Param("id"); id == c.GetString("user") {
			c.Next()
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Not Authorized User."})
		}

	}
}

func negotiate(c *gin.Context, data interface{}) gin.Negotiate {
	return gin.Negotiate{
		Offered: []string{gin.MIMEJSON, gin.MIMEXML},
		Data:    data,
	}
}

func handleErr(err error, res interface{}, c *gin.Context) {
	if res != nil {
		switch err {
		case nil:
			c.Negotiate(http.StatusOK, negotiate(c, res))
		case repository.ErrNoResult:
			c.Negotiate(http.StatusNotFound, negotiate(c, gin.H{"error": err.Error()}))
		default:
			c.Negotiate(http.StatusBadRequest, negotiate(c, gin.H{"error": err.Error()}))
		}
	} else {
		c.JSON(http.StatusNoContent, nil)
	}
}
