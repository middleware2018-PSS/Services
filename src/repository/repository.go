package repository

import (
	"github.com/middleware2018-PSS/Services/src/models"
)

type Repository interface {
	ClassByID(id int64) (interface{}, error)

	NotificationByID(id int64) (interface{}, error)

	PaymentByID(id int64) (interface{}, error)

	// Parents
	// see/modify their personal data
	ParentById(id int64) (interface{}, error)
	UpdateParent(id int64) error

	// see/modify the personal data of their registered children
	ChildrenByParent(id int64, limit int, offset int) ([]interface{}, error)
	StudentById(id int64) (interface{}, error)
	UpdateStudent(id int64) error

	// see the grades obtained by their children
	GradesByStudent(id int64, limit int, offset int) ([]interface{}, error)

	// see the monthly payments that have been made to the school in the past
	PaymentsByParent(id int64, limit int, offset int) ([]interface{}, error)

	// see general/personal notifications coming from the school
	NotificationsByParent(id int64, limit int, offset int) ([]interface{}, error)

	// see/modify appointments that they have with their children's teachers
	// (calendar-like support for requesting appointments, err error)
	AppointmentsByParent(id int64, limit int, offset int) ([]interface{}, error)
	UpdateAppointments(id int64) error
	AppointmentById(id int64) (interface{}, error)

	// see/modify their personal data
	TeacherByID(id int64) (teacher interface{}, err error)
	UpdateTeacher(id int64) (err error)

	// see the classrooms in which they teach, with information regarding the argument that they teach
	// in that class, the students that make up the class, and the complete lesson timetable for that
	// class
	ClassesPerSubjectByTeacher(id int64) (classes map[models.Subject][]models.Class, err error)
	StudentByClass(id int64, limit int, offset int) ([]interface{}, error)
	LectureByClass(id int64, limit int, offset int) ([]interface{}, error)

	// LectureByClass(id int64, limit int, offset int) (students []interface{}, err error)
	AppointmentsByTeacher(id int64, limit int, offset int) ([]interface{}, error)
	// UpdateAppointments(id int64, err error)
	NotificationsByTeacher(id int64, limit int, offset int) ([]interface{}, error)
	LectureByTeacher(id int64, limit int, offset int) ([]interface{}, error)

	// LectureByClass(id int64, limit int, offset int) (students []interface{}, err error)
	GradeStudent(grade models.Grade) error
	// TODO
	// parents:
	// see/pay (fake payment) upcoming scheduled payments (monthly, material, trips, err error)
	// admins:
	// everything

	Students(limit int, offset int) ([]interface{}, error)
	Teachers(limit int, offset int) ([]interface{}, error)
	Parents(limit int, offset int) ([]interface{}, error)
	Payments(limit int, offset int) ([]interface{}, error)
	Notifications(limit int, offset int) ([]interface{}, error)
	Classes(limit int, offset int) ([]interface{}, error)
}
