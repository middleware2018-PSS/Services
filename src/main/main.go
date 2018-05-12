package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/middleware2018-PSS/Services/src/controller"
	"github.com/middleware2018-PSS/Services/src/repository"
	"log"
	"net/http"
	"strconv"
	"github.com/pkg/errors"
)

var	(
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
	/*	l, _ := r.LectureByClass(1,0,100)
		l2, _:= r.LectureByClass(1,0,100)
		j1, _ := json.Marshal(l)
		j2, _ := json.Marshal(l2)
		fmt.Printf("%s ==  %s", j1, j2)*/

	api := gin.Default()
	// TODO implement AUTH with data from db
	parents := api.Group("/parents/:id", gin.BasicAuth(gin.Accounts{"3":"prova"}), Access())

	// TODO add hypermedia
	parents.GET("", byID("id", con.GetParentByID))
	parents.GET("/students", byIDWithOffsetAndLimit(con.StudentsByParent))
	// TODO redirect or deeper links? e.g. /parents/id/students/student/grades...
	parents.GET("/students/:student", byID("student",con.GetStudentByID))
	parents.GET("/appointments", byIDWithOffsetAndLimit(con.AppointmentsByParent))
	parents.GET("/payments", byIDWithOffsetAndLimit(con.PaymentsByParent))
	parents.GET("/notifications", byIDWithOffsetAndLimit(con.NotificationsByParent))

	// TODO /admin/...?
	api.GET("/students/:id", byID("id", con.GetStudentByID))
	api.GET("/students/:id/grades", byIDWithOffsetAndLimit(con.GradesByStudent))
	api.GET("/notifications/:id", byID("id", con.GetNotificationByID))
	api.GET("/payments/:id", byID("id", con.GetNotificationByID))


	teachers := api.Group("/teachers/:id", gin.BasicAuth(gin.Accounts{"3":"prova"}), Access())
	teachers.GET("", byID("id", con.GetTeacherByID))
	teachers.GET("/lectures", byIDWithOffsetAndLimit(con.LecturesByTeacher))
	teachers.GET("/appointments", byIDWithOffsetAndLimit(con.AppointmentsByTeacher))
	teachers.GET("/notifications", byIDWithOffsetAndLimit(con.NotificationsByTeacher))
	teachers.GET("/subjects", byIDWithOffsetAndLimit(con.SubjectByTeacher))
	teachers.GET("/subjects/:subject", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		subj := c.Param("subject")
		offset, err := strconv.Atoi(c.DefaultQuery("offset", "0"))
		limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
		res, err := con.ClassesBySubjectAndTeacher(int64(id), subj, limit, offset)
		handleErr(err, res, c)
	})

	api.GET("/parents", getOffsetLimit(con.Parents))
	api.GET("/students", getOffsetLimit(con.Students))
	api.GET("/notifications", getOffsetLimit(con.Notifications))
	api.GET("/payments", getOffsetLimit(con.Payments))
	api.GET("/teachers", getOffsetLimit(con.Teachers))
	api.GET("/classes", getOffsetLimit(con.Classes))

	api.Run(":5000")
}

func byID(key string, f func(int64) (interface{}, error)) func(c *gin.Context) {
	return func(c *gin.Context) {
		// TODO: check err
		id, err := strconv.Atoi(c.Param(key))
		res, err := f(int64(id))
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
		if limit > 0{
			res, err := f(limit, offset)
			handleErr(err, res, c)
		} else {
			handleErr( LimitError, nil, c)
		}

	}
}

func byIDWithOffsetAndLimit(f func(int64, int, int) ([]interface{}, error)) func(c *gin.Context) {
	return func(c *gin.Context) {
		//TODO Check id and errors (4 real)
		id, err := strconv.Atoi(c.Param("id"))
		offset, limit := offsetLimit(c)
		res, err := f(int64(id), limit, offset)
		handleErr(err, res, c)
	}
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

func handleErr(err error, res interface{}, c *gin.Context) {
	switch err {
	case nil:
		c.JSON(http.StatusOK, res)
	case repository.ErrNoResult:
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}
