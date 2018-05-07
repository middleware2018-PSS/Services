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
	s, _ := repo.StudentById(1)
	pprint(s)
	p, _ := repo.ParentById(4)
	pprint(p)
	t, _:= repo.TeacherByID(1)
	pprint(t)
	c, err:= repo.ClassesByID(2)
	pprint(c)
	fmt.Printf("%v",err)

}

func pprint(smt interface{}) {
	s, _ := json.MarshalIndent(smt, " ", "  ")
	fmt.Printf("%T : %s\n\n", smt, s)
}
func main() {
	test_repository()

}
