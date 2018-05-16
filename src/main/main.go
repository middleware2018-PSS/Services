package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/middleware2018-PSS/Services/src/models"
	"github.com/middleware2018-PSS/Services/src/repository"
	"github.com/middleware2018-PSS/Services/src/representations"
	"github.com/pkg/errors"
	"log"
	"net/http"
	"strconv"
	"strings"
	"github.com/appleboy/gin-jwt"
	"time"
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

	con := repository.NewPostgresRepository(db)

	g := gin.Default()
	authMiddleware := jwt.GinJWTMiddleware{
		Realm: "test",
		Key: []byte("password"),
		Timeout: time.Hour,
		MaxRefresh:time.Hour,
		Authenticator: func(userID string, password string, c *gin.Context) (string, bool) {
			return con.CheckUser(userID, password)
		},
		PayloadFunc: con.UserKind,
		}
	g.POST("/login", authMiddleware.LoginHandler)

	api := g.Group("", authMiddleware.MiddlewareFunc())
	api.GET("/refresh_token", authMiddleware.RefreshHandler)

	parent := api.Group("/parents/:id", authAdminOrParent(authMiddleware.Realm) )
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
		parent.GET("/appointments", byIDWithOffsetAndLimit("id", con.AppointmentsByParent))
		parent.GET("/payments", byIDWithOffsetAndLimit("id", con.PaymentsByParent))
		parent.GET("/notifications", byIDWithOffsetAndLimit("id", con.NotificationsByParent))
	}

	// TODO add hypermedia

	teachers := api.Group("/teachers/:id",authAdminOrTeacher(authMiddleware.Realm))
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

	api.GET("/appointments", getOffsetLimit(con.Appointments))
	api.GET("/appointments/:appointment", byID("appointment", con.AppointmentByID))
	api.PUT("/appointments/:appointment", func(c *gin.Context) {
		var a models.Appointment
		if err := c.ShouldBind(&a); err == nil {
			id, _ := strconv.Atoi(c.Param("id"))
			a.ID = int64(id)
			if err := con.UpdateAppointments(a); err == nil {
				c.JSON(http.StatusCreated, a)
			}
		}
	})
	// TODO remove admin from path and use token

	api.GET("/parents", getOffsetLimit(con.Parents))
	api.GET("/grades", getOffsetLimit(con.Grades))

	api.GET("/students", getOffsetLimit(con.Students))
	api.GET("/students/:id", byID("id", con.StudentByID))
	api.GET("/students/:id/grades", byIDWithOffsetAndLimit("id", con.GradesByStudent))

	api.GET("/notifications", getOffsetLimit(con.Notifications))
	api.GET("/notifications/:id", byID("id", con.NotificationByID))

	api.GET("/payments", getOffsetLimit(con.Payments))
	api.GET("/payments/:id", byID("id", con.PaymentByID))

	api.GET("/teachers", getOffsetLimit(con.Teachers))

	api.GET("/classes", getOffsetLimit(con.Classes))
	api.GET("/classes/:id", byID("id", con.ClassByID))
	api.GET("/classes/:id/students", byIDWithOffsetAndLimit("id", con.StudentsByClass))

	g.Run(":5000")
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
			for i, el := range res {
				res[i], _ = representations.ToRepresentation(el, c)
			}
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
			//TODO handle err
			res[i], _ = representations.ToRepresentation(el, c)
		}
		result := representations.List{c.Request.RequestURI,
			res,
			next(c.Request.RequestURI, offset, limit, res),
			prev(c.Request.RequestURI, offset, limit),
		}
		handleErr(err, result, c)
	}
}

func prev(uri string, offset int, limit int) string {
	if offset == 0 {
		return ""
	} else if n := strings.Index(uri, "?"); n >= 0 {
		uri = uri[:n]
	}
	if prev := offset - limit; prev < 0 {
		offset = 0
	} else {
		offset = prev
	}
	return strings.Join([]string{uri, fmt.Sprintf("?offset=%d&limit=%d", offset, limit)}, "")
}

func next(uri string, offset int, limit int, input []interface{}) string {
	if l := len(input); l < limit {
		return ""
	}
	if n := strings.Index(uri, "?"); n >= 0 {
		uri = uri[:n]
	}
	return strings.Join([]string{uri, fmt.Sprintf("?offset=%d&limit=%d", offset+limit, limit)}, "")
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

func negotiate(data interface{}) gin.Negotiate {
	return gin.Negotiate{
		Offered: []string{gin.MIMEJSON, gin.MIMEXML},
		Data:    data,
	}
}

func handleErr(err error, res interface{}, c *gin.Context) {
	if res != nil {
		switch err {
		case nil:
			c.Negotiate(http.StatusOK, negotiate(res))
		case repository.ErrNoResult:
			c.Negotiate(http.StatusNotFound, negotiate(gin.H{"error": err.Error()}))
		default:
			c.Negotiate(http.StatusBadRequest, negotiate(gin.H{"error": err.Error()}))
		}
	} else {
		c.JSON(http.StatusNoContent, nil)
	}
}

func authAdminOrParent(realm string) func(c *gin.Context) {
	return func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		s, err := strconv.Atoi(c.Param("id"))
		if k := claims["kind"]; err != nil || k == "admin" || (k.(string) == "parent" && (claims["dbID"] == float64(s))) {
			c.Next()
		} else {
			c.Header("WWW-Authenticate", realm)
			c.AbortWithStatus(http.StatusForbidden)
			return
		}
	}
	}

func authAdminOrTeacher(realm string) func(c *gin.Context) {
	return func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		s, err := strconv.Atoi(c.Param("id"))
		if k := claims["kind"]; err != nil || k == "admin" || (k.(string) == "teacher" && (claims["dbID"] == float64(s))) {
			c.Next()
		} else {
			c.Header("WWW-Authenticate", realm)
			c.AbortWithStatus(http.StatusForbidden)
			return
		}
	}
	}

func authAdmin(realm string) func(c *gin.Context) {
	return func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		if k := claims["kind"]; k == "admin" {
			c.Next()
		} else {
			c.Header("WWW-Authenticate", realm)
			c.AbortWithStatus(http.StatusForbidden)
			return
		}
	}
	}