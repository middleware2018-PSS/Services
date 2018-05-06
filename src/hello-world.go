package main

import (
	_ "github.com/lib/pq"
	"database/sql"
	"log"
	"github.com/middleware2018-PSS/Services/src/repository"
)

func test_repository(){
	connStr := "user=postgres dbname=postgres sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	repo := repository.NewPostgresRepository(db)
	repo.StudentById(1)
	repo.ParentById(3)


}


func main() {
	test_repository()

}

