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
	api.GET("/parent/:id", ById(con.GetParentByID))
	api.GET("/student/:id", ById(con.GetStudentByID))
	api.GET("/notification/:id", ById(con.GetNotificationByID))
	api.GET("/payment/:id", ById(con.GetNotificationByID))
	api.GET("/teacher/:id", ById(con.GetTeacherByID))

	api.Run(":5000")
}

func ById(f func(int64) (interface{}, error)) func(c *gin.Context) {
	return func(c *gin.Context) {
		// TODO: check err
		id, err := strconv.Atoi(c.Param("id"))
		p, err := f(int64(id))
		switch err {
		case nil:
			c.JSON(http.StatusOK, p)
		default:
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		}
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
