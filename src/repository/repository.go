package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/middleware2018-PSS/Services/src/models"
	"log"
)

var (
	ErrNoResult      = errors.New("No result found.")
	ErrorNotBlocking = errors.New("Something went wrong but no worriez.")
)

type Repository interface {
	ClassesByID(id int64) (class models.Class, err error)

	NotificationByID(id int64) (notification interface{}, err error)

	PaymentByID(id int64) (payment interface{}, err error)

	// Parents
	// see/modify their personal data
	ParentById(id int64) (parent interface{}, err error)
	UpdateParent(id int64) (err error)

	// see/modify the personal data of their registered children
	ChildrenByParent(id int64, offset int, limit int) ([]interface{}, error)
	StudentById(id int64) (student interface{}, err error)
	UpdateStudent(id int64) (err error)

	// see the grades obtained by their children
	GradesByStudent(id int64, offset int, limit int) (grades []interface{}, err error)

	// see the monthly payments that have been made to the school in the past
	PaymentsByParent(id int64, offset int, limit int) (payments []interface{}, err error)

	// see general/personal notifications coming from the school
	NotificationsByParent(id int64, offset int, limit int) (list []interface{}, err error)

	// see/modify appointments that they have with their children's teachers
	// (calendar-like support for requesting appointments, err error)
	AppointmentsByParent(id int64, offset int, limit int) (appointments []interface{}, err error)
	UpdateAppointments(id int64) (err error)
	AppointmentById(id int64) (appointment models.Appointment, err error)

	// see/modify their personal data
	TeacherByID(id int64) (teacher interface{}, err error)
	UpdateTeacher(id int64) (err error)

	// see the classrooms in which they teach, with information regarding the argument that they teach
	// in that class, the students that make up the class, and the complete lesson timetable for that
	// class
	ClassesPerSubjectByTeacher(id int64) (classes map[models.Subject][]models.Class, err error)
	StudentByClass(id int64, offset int, limit int) ([]interface{}, error)
	LectureByClass(id int64, offset int, limit int) (lectures []interface{}, err error)

	// LectureByClass(id int64, offset int, limit int) (students []interface{}, err error)
	AppointmentsByTeacher(id int64, offset int, limit int) (appointments []interface{}, err error)
	// UpdateAppointments(id int64, err error)
	NotificationsByTeacher(id int64, offset int, limit int) (notifications []interface{}, err error)
	LectureByTeacher(id int64, offset int, limit int) (lectures []interface{}, err error)

	// LectureByClass(id int64, offset int, limit int) (students []interface{}, err error)
	GradeStudent(grade models.Grade) (err error)

	// TODO
	// parents:
	// see/pay (fake payment) upcoming scheduled payments (monthly, material, trips, err error)
	// admins:
	// everything

}

func switchError(err error) error {
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			err = ErrNoResult
		default:
			if fmt.Sprintf("%v", err)[:len("sql: Scan error")] == "sql: Scan error" {
				err = ErrorNotBlocking
			}
		}
	}
	return err
}

func (r *postgresRepository) listByID(id int64, offset int, limit int, query string, f func(*sql.Rows) (interface{}, error)) (list []interface{}, err error) {
	rows, err := r.Query(query, id, limit, offset)
	defer rows.Close()
	if err != nil {
		log.Print(err)
	}
	for rows.Next() {
		el, err := f(rows)
		if err != nil {
			//TODO check error
		}
		list = append(list, el)
	}
	return list, err

}
