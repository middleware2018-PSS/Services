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

	con := controller.NewController(repository.NewPostgresRepository(db))
	api := gin.Default()
	// TODO: add authentication
	//auth := api.Group("/parent", gin.BasicAuth(gin.Accounts{"3":"bar"}))
	//auth.Use(Auth())
	// auth.GET("/:id",getParent )
	api.GET("/parent/:id", func(c *gin.Context) {
		// TODO: check err
		id, _ := strconv.Atoi(c.Param("id"))
		p, _ := con.GetParentByID(int64(id))
		c.JSON(http.StatusOK, p)
	})
	api.GET("/student/:id", func(c *gin.Context) {
		// TODO: check err
		id, err := strconv.Atoi(c.Param("id"))
		p, err := con.GetStudentByID(int64(id))
		switch err {
		case nil:
			c.JSON(http.StatusOK, p)
		default:
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		}
	})
	api.Run(":5000")
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
