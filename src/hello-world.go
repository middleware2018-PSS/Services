package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/middleware2018-PSS/Services/src/repository"
	"log"
)

func test_repository() {
	connStr := "user=postgres dbname=postgres sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	repo := repository.NewPostgresRepository(db)
	student := repo.StudentById(1)
	s, _ := json.MarshalIndent(student, " ", "  ")
	fmt.Printf("student : %s\n\n", s)
	parent := repo.ParentById(3)
	s, _ = json.MarshalIndent(parent, " ", "  ")
	fmt.Printf("parent : %s\n\n", s)



}

func main() {
	test_repository()

}
