package controller

// shall check accesses and orchestrate actions of the

import (
	"github.com/middleware2018-PSS/Services/src/repository"
)

type Controller struct {
	r repository.Repository
}

func NewController(repo repository.Repository) Controller {
	return Controller{repo}
}

func (con Controller) GetTeacherByID(id int64) (interface{}, error) {
	return getByID(id, con.r.TeacherByID)
}

func (con Controller) GetNotificationByID(id int64) (interface{}, error) {
	return getByID(id, con.r.NotificationByID)
}

func (con Controller) GetParentByID(id int64) (interface{}, error) {
	return getByID(id, con.r.ParentById)
}

func (con Controller) GetStudentByID(id int64) (interface{}, error) {
	return getByID(id, con.r.StudentById)
}

func (con Controller) Students(limit int, offset int) ([]interface{}, error) {
	return con.r.Students(limit, offset)
}

func (con Controller) Teachers(limit int, offset int) ([]interface{}, error) {
	return con.r.Teachers(limit, offset)
}

func (con Controller) Classes(limit int, offset int) ([]interface{}, error) {
	return con.r.Classes(limit, offset)
}

func (con Controller) Parents(limit int, offset int) ([]interface{}, error) {
	return con.r.Parents(limit, offset)
}

func (con Controller) Payments(limit int, offset int) ([]interface{}, error) {
	return con.r.Payments(limit, offset)
}

func (con Controller) Notifications(limit int, offset int) ([]interface{}, error) {
	return con.r.Notifications(limit, offset)
}

func getByID(id int64, f func(int64) (interface{}, error)) (interface{}, error) {
	res, e := f(id)
	//TODO: check err
	switch e {
	case repository.ErrNoResult:
		return nil, e
	case repository.ErrorNotBlocking:
		return res, nil
	default:
		return res, nil
	}
}
