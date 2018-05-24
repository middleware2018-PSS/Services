package main

import (
	"encoding/json"
	"fmt"

	"github.com/middleware2018-PSS/Services/src/models"
	"github.com/middleware2018-PSS/Services/src/representations"
	"github.com/gin-gonic/gin"
)

type Student struct {
	ID      int     `json:"id" xml:"id" example:"1"`
	Name    *string `json:"name" xml:"name"`
	Surname *string `json:"surname" xml:"surname"`
	Mail    *string `json:"mail" xml:"mail"`
	Info    *string `json:"info" xml:"info"`
}

func main() {
	r, _ := representations.ToHALRepresentation(models.Parent{ID:1}, &gin.Context{})
	s, _ := json.Marshal(r)
	fmt.Printf("%s", s)
}