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
	// TODO: add authentication
	//auth := api.Group("/parent", gin.BasicAuth(gin.Accounts{"3":"bar"}))
	//auth.Use(Auth())
	// auth.GET("/:id",getParent )
	// admin := api.Group("/admin")
	api.GET("/parents/:id", byID(con.GetParentByID))
	api.GET("/students/:id", byID(con.GetStudentByID))
	api.GET("/notifications/:id", byID(con.GetNotificationByID))
	api.GET("/payments/:id", byID(con.GetNotificationByID))
	api.GET("/teachers/:id", byID(con.GetTeacherByID))
	api.GET("/teachers/:id/lectures", byIDWithOffsetAndLimit(con.LecturesByTeacher))
	api.GET("/teachers/:id/appointments", byIDWithOffsetAndLimit(con.AppointmentsByTeacher))
	api.GET("/teachers/:id/notifications", byIDWithOffsetAndLimit(con.NotificationsByTeacher))
	api.GET("/teachers/:id/subjects", byIDWithOffsetAndLimit(con.SubjectByTeacher))
	api.GET("/teachers/:id/subjects/:subject", func(c *gin.Context) {
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

func byID(f func(int64) (interface{}, error)) func(c *gin.Context) {
	return func(c *gin.Context) {
		// TODO: check err
		id, err := strconv.Atoi(c.Param("id"))
		res, err := f(int64(id))
		handleErr(err, res, c)
	}

}

func getOffsetLimit(f func(int, int) ([]interface{}, error)) func(c *gin.Context) {
	return func(c *gin.Context) {
		//TODO Check id and errors
		offset, err := strconv.Atoi(c.DefaultQuery("offset", "0"))
		limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
		res, err := f(limit, offset)

		handleErr(err, res, c)
	}
}

func byIDWithOffsetAndLimit(f func(int64, int, int) ([]interface{}, error)) func(c *gin.Context) {
	return func(c *gin.Context) {
		//TODO Check id and errors (4 real)
		id, err := strconv.Atoi(c.Param("id"))
		offset, err := strconv.Atoi(c.DefaultQuery("offset", "0"))
		limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
		res, err := f(int64(id), limit, offset)
		handleErr(err, res, c)
	}
}

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		if id := c.Param("id"); id == c.GetString("user") {
			c.Next()
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "not allowed"})
			c.AbortWithStatus(401)
		}

	}
}

func handleErr(err error, res interface{}, c *gin.Context) {
	switch err {
	case nil:
		c.JSON(http.StatusOK, res)
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}
