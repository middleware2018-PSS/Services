package repository

import "github.com/middleware2018-PSS/Services/src/models"

type Repository interface {
	ClassesByID(id int64) (class models.Class)

	NotificationByID(id int64) (notifications []models.Notification)

	PaymentByID(id int64) (payments []models.Payment)

	ParentById(id int64) (parent *models.Parent)
	ChildrenByParent(id int64) (child []models.Student)
	PaymentsByParent(id int64) (payments []models.Payment)
	NotificationsByParent(id int64) (list []models.Notification)

	StudentById(id int) (student *models.Student)
	PaymentByStudent(id int64) (payments []models.Payment)
	GradesByStudent(id int64) (grades []models.Grade)
	NotificationsByStudent(id int64) (notifications []models.Notification)
	ClassesByStudent(id int64) (classes []models.Class)
	AppointmentsByStudent(id int64) (appointments []models.Appointment)

	TeacherByID(id int64) (teacher *models.Teacher, err error)
	ClassesByTeacher(id int64) (classes map[models.Subject][]models.Class)
	AppointmentsByTeacher(id int64) (appointments []models.Appointment)
	NotificationsByTeacher(id int64) (notifications []models.Notification)
	LectureByTeacher(id int64) (lectures []models.TimeTable)
}
