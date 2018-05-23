package repository

import (
	"errors"
	"github.com/middleware2018-PSS/Services/src/models"
)

const (
	USER = "User"
	KIND = "Kind"
	AdminUser = "Admin"
	ParentUser = "Parent"
	TeacherUser = "Teacher"
	GET = "GET"
	PUT = "PUT"
	POST = "POST"
	ALL = iota
	CLASS
	PARENT
	TEACHER
	ADM
	STUDENT
	APPOINTMENT

)

var (
	ErrNoResult      = errors.New("No results found.")
	ErrorNotBlocking = errors.New("Something went wrong but no worriez.")
	ErrorNotAuthorized = errors.New("No authorization for this resource")
)

type Repository interface {
	ClassByID(id int, who int, whoKind string) (interface{}, error)
	GradeByID(id int,who int, whoKind string) (interface{}, error)

	NotificationByID(id int,who int, whoKind string) (interface{}, error)

	PaymentByID(id int,who int, whoKind string) (interface{}, error)

	// Parents
	// see/modify their personal data
	ParentByID(id int, who int, whoKind string) (interface{}, error)
	UpdateParent(models.Parent, int, string) error
	UpdateStudent(student models.Student, who int, whoKind string) error

	// see/modify the personal data of their registered children
	ChildrenByParent(id int, limit int, offset int,who int, whoKind string) ([]interface{}, error)
	StudentByID(id int,who int, whoKind string) (interface{}, error)

	// see the grades obtained by their children
	GradesByStudent(id int, limit int, offset int,who int, whoKind string) ([]interface{}, error)

	// see the monthly payments that have been made to the school in the past
	PaymentsByParent(id int, limit int, offset int,who int, whoKind string) ([]interface{}, error)

	// see general/personal notifications coming from the school
	NotificationsByParent(id int, limit int, offset int,who int, whoKind string) ([]interface{}, error)

	// see/modify appointments that they have with their children's teachers
	// (calendar-like support for requesting appointments, err error)
	AppointmentsByParent(id int, limit int, offset int,who int, whoKind string) ([]interface{}, error)
	UpdateAppointment(appointment models.Appointment, who int, whoKind string) error
	AppointmentByID(id int,who int, whoKind string) (interface{}, error)

	// see/modify their personal data
	TeacherByID(id int,who int, whoKind string) (teacher interface{}, err error)
	UpdateTeacher(teacher models.Teacher, who int, whoKind string) (err error)

	// see the classrooms in which they teach, with information regarding the argument that they teach
	// in that class, the students that make up the class, and the complete lesson timetable for that
	// class
	SubjectsByTeacher(id int, limit int, offset int, who int, whoKind string) ([]interface{}, error)
	ClassesBySubjectAndTeacher(teacher int, subject string, limit int, offset int, who int, whoKind string) ([]interface{}, error)

	StudentsByClass(id int, limit int, offset int,who int, whoKind string) ([]interface{}, error)
	LectureByClass(id int, limit int, offset int,who int, whoKind string) ([]interface{}, error)

	// LectureByClass(id int, limit int, offset int,who int, whoKind string) (students []interface{}, err error)
	AppointmentsByTeacher(id int, limit int, offset int,who int, whoKind string) ([]interface{}, error)
	// UpdateAppointments(id int, err error,who int, whoKind string)
	NotificationsByTeacher(id int, limit int, offset int,who int, whoKind string) ([]interface{}, error)
	LecturesByTeacher(id int, limit int, offset int,who int, whoKind string) ([]interface{}, error)
	ClassesByTeacher(id int, limit int, offset int,who int, whoKind string) ([]interface{}, error)

	// LectureByClass(id int, limit int, offset int,who int, whoKind string) (students []interface{}, err error)
	// TODO GradeStudent(grade models.Grade) error
	// TODO
	// parents:
	// see/pay (fake payment) upcoming scheduled payments (monthly, material, trips, err error)
	// admins:
	// everything

	Students(limit int, offset int, who int, whoKind string) ([]interface{}, error)
	Teachers(limit int, offset int, who int, whoKind string) ([]interface{}, error)
	Parents(limit int, offset int, who int, whoKind string) ([]interface{}, error)
	Payments(limit int, offset int, who int, whoKind string) ([]interface{}, error)
	Notifications(limit int, offset int, who int, whoKind string) ([]interface{}, error)
	Classes(limit int, offset int, who int, whoKind string) ([]interface{}, error)
	CheckUser(s string, s2 string) (int, string, bool)
	CreateParent(parent models.Parent, who int, whoKind string) (int, error)
	CreateAppointment(appointment models.Appointment, who int, whoKind string) (int, error)
	CreateTeacher(teacher models.Teacher, who int, whoKind string) (int, error)
	Appointments(int, int, int, string) ([]interface{}, error)
	Grades(int, int, int, string) ([]interface{}, error)
	CreateStudent(models.Student, int, string) (int, error)
	CreateClass(models.Class, int, string) (int, error)
	UpdateClass(models.Class, int, string) error
	CreateNotification(models.Notification, int, string) (int, error)
	UpdateNotification(models.Notification, int, string) error
	CreateGrade(models.Grade, int, string) (int, error)
	UpdateGrade(models.Grade, int, string) error
	CreatePayment(models.Payment, int, string) (int, error)
	UpdatePayment(models.Payment, int, string) error
	IsParent(int, int) bool
	ParentHasAppointment(int, int) bool
}
