package controller

// shall check accesses and orchestrate actions of the

import (
	"github.com/middleware2018-PSS/Services/src/repository"
)

func getListByIDOffsetLimit(id int64, limit int, offset int, f func(int64, int, int) ([]interface{}, error)) ([]interface{}, error) {
	res, e := f(id, limit, offset)
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

type Controller struct {
	r repository.Repository
}

func NewController(repo repository.Repository) Controller {
	return Controller{repo}
}

func (con Controller) GetTeacherByID(id int64) (interface{}, error) {
	return getByID(id, con.r.TeacherByID)
}

func (con Controller) ClassByID(id int64) (interface{}, error) {
	return getByID(id, con.r.ClassByID)
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

func (con Controller) LecturesByTeacher(id int64, limit int, offset int) ([]interface{}, error) {
	return getListByIDOffsetLimit(id, limit, offset, con.r.LecturesByTeacher)
}

func (con Controller) AppointmentsByTeacher(id int64, limit int, offset int) ([]interface{}, error) {
	return getListByIDOffsetLimit(id, limit, offset, con.r.AppointmentsByTeacher)
}

func (con Controller) SubjectByTeacher(id int64, limit int, offset int) ([]interface{}, error) {
	return getListByIDOffsetLimit(id, limit, offset, con.r.SubjectsByTeacher)
}

func (con Controller) ClassesBySubjectAndTeacher(teacher int64, subject string, limit int, offset int) ([]interface{}, error) {
	return con.r.ClassesBySubjectAndTeacher(teacher, subject, limit, offset)
}

func (con Controller) NotificationsByTeacher(id int64, limit int, offset int) ([]interface{}, error) {
	return getListByIDOffsetLimit(id, limit, offset, con.r.NotificationsByTeacher)
}

func (con Controller) StudentsByParent(id int64, limit int, offset int) ([]interface{}, error) {
	return getListByIDOffsetLimit(id, limit, offset, con.r.ChildrenByParent)
}

func (con Controller) GradesByStudent(id int64, limit int, offset int) ([]interface{}, error) {
	return con.r.GradesByStudent(id, limit, offset)
}

func (con Controller) AppointmentsByParent(id int64, limit int, offset int) ([]interface{}, error) {
	return con.r.AppointmentsByParent(id, limit, offset)
}

func (con Controller) PaymentsByParent(id int64, limit int, offset int) ([]interface{}, error) {
	return con.r.PaymentsByParent(id, limit, offset)
}

func (con Controller) NotificationsByParent(id int64, limit int, offset int) ([]interface{}, error) {
	return con.r.NotificationsByParent(id, limit, offset)
}

func (con Controller) ClassesByTeacher(id int64, limit int, offset int) ([]interface{}, error) {
	return con.r.ClassesByTeacher(id, limit, offset)
}

func (con Controller) StudentsByClass(id int64, limit int, offset int) ([]interface{}, error) {
	return con.r.StudentByClass(id, limit, offset)
}

func (con Controller) PaymentByID(id int64) (interface{}, error) {
	return con.r.PaymentByID(id)
}
