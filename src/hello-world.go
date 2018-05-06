package main

import (
	"github.com/middleware2018-PSS/Services/src/models"
	"fmt"
	"encoding/json"
	"time"
	"github.com/gobuffalo/pop"
	"log"
)

func main() {
	student := models.Student{}
	student.Payments = []models.Payment{{0, 100, true, time.Now(), ""}}
	student.Appointments = []models.Appointment{
		{
			Time:time.Now(),
			Location:"",
			Teacher: models.Teacher{},
		},
	}
	str, _ := json.MarshalIndent(student," ", "  ")
	str2, _ := json.MarshalIndent(models.Teacher{}," ", "  ")
	fmt.Printf("%s\n",str)
	fmt.Printf("%s\n",str2)
	fmt.Println(time.Now().Date())
	db, err := pop.Connect("development")
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()
	fmt.Println(db.URL())

}

