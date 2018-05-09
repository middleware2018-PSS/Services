package controller

// shall check accesses and orchestrate actions of the

import (
	"github.com/middleware2018-PSS/Services/src/repository"
	"log"
)

type Controller struct {
	r repository.Repository
}

func NewController(repo repository.Repository) Controller {
	return Controller{repo}
}

func (con Controller) GetTeacherByID(id int64) (parent interface{}, err error) {
	return getByID(id, con.r.TeacherByID)
}

func (con Controller) GetNotificationByID(id int64) (parent interface{}, err error) {
	return getByID(id, con.r.NotificationByID)
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
	case repository.ErrNoResult:
		log.Print(e)
		return nil, e
	case repository.ErrorNotBlocking:
		log.Print(e)
		return res, nil
	default:
		return res, nil
	}
}


