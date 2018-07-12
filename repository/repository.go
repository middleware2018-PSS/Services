package repository

import (
	"errors"

	"github.com/middleware2018-PSS/Services/models"
)

const (
	USER        = "userID"
	KIND        = "kind"
	AdminUser   = "Admin"
	ParentUser  = "Parent"
	TeacherUser = "Teacher"
)

func init() {
	allowedKind = map[string]bool{AdminUser: true, ParentUser: true, TeacherUser: true}
}

var (
	allowedKind          map[string]bool
	ErrNoResult          = errors.New("No results found.")
	ErrorNotBlocking     = errors.New("Something went wrong but no worriez.")
	ErrorNotAuthorized   = errors.New("No authorization for this resource")
	ErrorNoKindSpecified = errors.New("No \"kind\" has been specified")
)

type Repository interface {
	ClassByID(id int) (interface{}, error)
	GradeByID(id int) (interface{}, error)

	NotificationByID(id int) (interface{}, error)

	PaymentByID(id int) (interface{}, error)

	// Parents
	// see/modify their personal data
	ParentByID(id int) (interface{}, error)
	UpdateParent(models.Parent, int, string) error
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

	// LectureByClass(id int, limit int, offset int,who int) (students []interface{}, err error)
	AppointmentsByTeacher(id int, limit int, offset int) ([]interface{}, error)
	// UpdateAppointments(id int, err error,who int)
	NotificationsByTeacher(id int, limit int, offset int) ([]interface{}, error)
	LecturesByTeacher(id int, limit int, offset int) ([]interface{}, error)
	ClassesByTeacher(id int, limit int, offset int) ([]interface{}, error)

	// LectureByClass(id int, limit int, offset int,who int) (students []interface{}, err error)
	// TODO GradeStudent(grade models.Grade) error
	// TODO
	// parents:
	// see/pay (fake payment) upcoming scheduled payments (monthly, material, trips, err error)
	// admins:
	// everything
	Lectures(limit int, offset int) ([]interface{}, error)
	Students(limit int, offset int) ([]interface{}, error)
	Teachers(limit int, offset int) ([]interface{}, error)
	Parents(limit int, offset int) ([]interface{}, error)
	Payments(limit int, offset int) ([]interface{}, error)
	Notifications(limit int, offset int) ([]interface{}, error)
	Classes(limit int, offset int) ([]interface{}, error)
	CheckUser(s string, s2 string) (int, string, bool)
	CreateParent(parent models.Parent) (int, error)
	CreateAppointment(appointment models.Appointment) (int, error)
	CreateTeacher(teacher models.Teacher) (int, error)
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

	CreateAccount(username string, password string, id int, kind string, cost int) error
	LectureByID(id int) (interface{}, error)
	DeleteParent(id int) (interface{}, error)
	DeleteTeacher(id int) (interface{}, error)
	DeleteAppointment(id int) (interface{}, error)
	DeleteStudent(id int) (interface{}, error)
	DeleteNotification(id int) (interface{}, error)
	DeletePayment(id int) (interface{}, error)
	DeleteClass(id int) (interface{}, error)
	DeleteAccount(username string) (interface{}, error)
	DeleteGrade(id int) (interface{}, error)
	DeleteLecture(id int) (interface{}, error)
	UpdateLecture(models.TimeTable, int, string) error
	CreateLecture(lecture models.TimeTable) (int, error)
	UpdateAccount(account models.Account, cost int) error
	Accounts(limit int, offset int) ([]interface{}, error)
}
