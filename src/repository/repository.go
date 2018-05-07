package repository

import "github.com/middleware2018-PSS/Services/src/models"

type Repository interface {

	ClassesByID(id int64) (class models.Class)

	NotificationByID(id int64) (notification models.Notification)

	PaymentByID(id int64) (payment models.Payment)


	// Parents
	// see/modify their personal data
	ParentById(id int64) (parent *models.Parent)
	UpdateParent(id int64)

	// see/modify the personal data of their registered children
	ChildrenByParent(id int64, offset int, limit int) (children []models.Student)
	StudentById(id int) (student *models.Student)
	UpdateStudent(id int64)

	// see the grades obtained by their children
	GradesByStudent(id int64, offset int, limit int) (grades []models.Grade)

	// see the monthly payments that have been made to the school in the past
	PaymentsByParent(id int64, offset int, limit int) (payments []models.Payment)

	// see general/personal notifications coming from the school
	NotificationsByParent(id int64, offset int, limit int) (list []models.Notification)

	// see/modify appointments that they have with their children's teachers
	// (calendar-like support for requesting appointments)
	AppointmentsByParent(id int64, offset int, limit int) (appointments []models.Appointment)
	UpdateAppointments(id int64)
	AppointmentById(id int64) (appointment models.Appointment)

	// see/modify their personal data
	TeacherByID(id int64) (teacher *models.Teacher)
	UpdateTeacher(id int64)

	// see the classrooms in which they teach, with information regarding the argument that they teach
	// in that class, the students that make up the class, and the complete lesson timetable for that
	// class
	ClassesByTeacher(id int64) (classes map[models.Subject][]models.Class)
	StudentByClass(id int64, offset int, limit int) (students []models.Student)
	LectureByClass(id int64, offset int, limit int) (lectures []models.TimeTable)

	// LectureByClass(id int64, offset int, limit int) (students []models.TimeTable)
	AppointmentsByTeacher(id int64, offset int, limit int) (appointments []models.Appointment)
	// UpdateAppointments(id int64)
	NotificationsByTeacher(id int64, offset int, limit int) (notifications []models.Notification)
	LectureByTeacher(id int64, offset int, limit int) (lectures []models.TimeTable)


	// LectureByClass(id int64, offset int, limit int) (students []models.TimeTable)
	GradeStudent(grade models.Grade)

	// TODO
	// parents:
	// see/pay (fake payment) upcoming scheduled payments (monthly, material, trips)
	// admins:
	// everything

}
