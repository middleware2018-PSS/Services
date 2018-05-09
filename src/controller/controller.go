package controller

// shall check accesses and orchestrate actions of the

import (
	"database/sql"
	"fmt"
	"github.com/middleware2018-PSS/Services/src/repository"
	"github.com/pkg/errors"
)

type Controller struct {
	r repository.Repository
}

var (
	NoResult = errors.New("No result found")
)

func NewController(repo repository.Repository) Controller {
	return Controller{repo}
}

func (con Controller) GetParentByID(id int64) (parent interface{}, err error) {
	return getByID(id, con.r.ParentById)
}

func (con Controller) GetStudentByID(id int64) (student interface{}, err error) {
	return getByID(id, con.r.StudentById)
}

func getByID(id int64, f func(int64) (interface{}, error)) (res interface{}, err error) {
	res, e := f(id)
	//TODO: check err
	switch e {
	case sql.ErrNoRows:
		return res, NoResult
	default:
		if fmt.Sprintf("%v", e)[:len("sql: Scan error")] == "sql: Scan error" {
			return res, nil
		} else {
			return res, err
		}
	}
}
