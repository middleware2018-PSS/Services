package repository

import (
	"errors"
	"github.com/middleware2018-PSS/Services/src/models"
)

var (
	ErrNoResult      = errors.New("No results found.")
	ErrorNotBlocking = errors.New("Something went wrong but no worriez.")
)

type Repository interface {
	ClassByID(id int) (interface{}, error)
	GradeByID(id int) (interface{}, error)

	NotificationByID(id int) (interface{}, error)

	PaymentByID(id int) (interface{}, error)

	// Parents
	// see/modify their personal data
	ParentByID(id int) (interface{}, error)
	UpdateParent(models.Parent) error
	UpdateStudent(student models.Student) error

	// see/modify the personal data of their registered children
	ChildrenByParent(id int, limit int, offset int) ([]interface{}, error)
	StudentByID(id int) (interface{}, error)

	// see the grades obtained by their children
	GradesByStudent(id int, limit int, offset int) ([]interface{}, error)

	// see the monthly payments that have been made to the school in the past
	PaymentsByParent(id int, limit int, offset int) ([]interface{}, error)

	// see general/personal notifications coming from the school
	NotificationsByParent(id int, limit int, offset int) ([]interface{}, error)

	// see/modify appointments that they have with their children's teachers
	// (calendar-like support for requesting appointments, err error)
	AppointmentsByParent(id int, limit int, offset int) ([]interface{}, error)
	UpdateAppointment(appointment models.Appointment) error
	AppointmentByID(id int) (interface{}, error)

	// see/modify their personal data
	TeacherByID(id int) (teacher interface{}, err error)
	UpdateTeacher(teacher models.Teacher) (err error)

	// see the classrooms in which they teach, with information regarding the argument that they teach
	// in that class, the students that make up the class, and the complete lesson timetable for that
	// class
	SubjectsByTeacher(id int, limit int, offset int) ([]interface{}, error)
	ClassesBySubjectAndTeacher(teacher int, subject string, limit int, offset int) ([]interface{}, error)

	StudentsByClass(id int, limit int, offset int) ([]interface{}, error)
	LectureByClass(id int, limit int, offset int) ([]interface{}, error)

	// LectureByClass(id int, limit int, offset int) (students []interface{}, err error)
	AppointmentsByTeacher(id int, limit int, offset int) ([]interface{}, error)
	// UpdateAppointments(id int, err error)
	NotificationsByTeacher(id int, limit int, offset int) ([]interface{}, error)
	LecturesByTeacher(id int, limit int, offset int) ([]interface{}, error)
	ClassesByTeacher(id int, limit int, offset int) ([]interface{}, error)

	// LectureByClass(id int, limit int, offset int) (students []interface{}, err error)
	// TODO GradeStudent(grade models.Grade) error
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
	CheckUser(s string, s2 string) (string, string, bool)
	UserKind(userID string) map[string]interface{}
	CreateParent(parent models.Parent) (int, error)
	CreateAppointment(appointment models.Appointment) (int, error)
	CreateTeacher(teacher models.Teacher) (int, error)
	Appointments(int, int) ([]interface{}, error)
	Grades(int, int) ([]interface{}, error)
	CreateStudent(models.Student) (int, error)
	CreateClass(models.Class) (int, error)
	UpdateClass(models.Class) error
	CreateNotification(models.Notification) (int, error)
	UpdateNotification(models.Notification) error
	CreateGrade(models.Grade) (int, error)
	UpdateGrade(models.Grade) error
	CreatePayment(models.Payment) (int, error)
	UpdatePayment(models.Payment) error
	IsParent(parent int, child int) bool
	ParentHasAppointment(parent int, appointment int) bool
}
