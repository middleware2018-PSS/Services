package controller

import "github.com/middleware2018-PSS/Services/repository"

// @Summary Get a appointment by id
// @Param id path int true "Appointment ID"
// @Tags Appointments
// @Success 200 {object} models.Appointment
// @Router /appointments/{id} [get]
// @Security ApiKeyAuth
func (c Controller) AppointmentByID(id int, who int, whoKind string) (interface{}, error) {
	switch whoKind {
	case repository.ParentUser:
		return c.repo.AppointmentForParent(id, who)
	case repository.TeacherUser:
		return c.repo.AppointmentForTeacher(id, who)
	case repository.AdminUser:
		return c.repo.AppointmentForAdmin(id)
	default:
		return nil, ErrorNotAuthorized
	}
}

// @Summary Get a grade by id
// @Param id path int true "Grade ID"
// @Tags Grades
// @Success 200 {object} models.Grade
// @Router /grades/{id} [get]
// @Security ApiKeyAuth
func (c Controller) GradeByID(id int, who int, whoKind string) (interface{}, error) {
	switch whoKind {
	case repository.ParentUser:
		return c.repo.GradeByIDForParent(id, who)
	case repository.TeacherUser:
		return c.repo.GradeByIDForTeacher(id, who)
	case repository.AdminUser:
		return c.repo.GradeByIDForAdmin(id)
	default:
		return nil, ErrorNotAuthorized
	}
}

// @Summary Get a class by id
// @Param id path int true "Class ID"
// @Tags Classes
// @Success 200 {object} models.Class
// @Router /classes/{id} [get]
// @Security ApiKeyAuth
func (c Controller) ClassByID(id int, who int, whoKind string) (interface{}, error) {
	switch whoKind {
	case repository.TeacherUser:
		return c.repo.ClassByIDForTeacher(id, who)
	case repository.AdminUser:
		return c.repo.ClassByIDForAdmin(id)
	default:
		return nil, ErrorNotAuthorized
	}
}

// @Summary Get a notification by id
// @Param id path int true "Notification ID"
// @Tags Notifications
// @Success 200 {object} models.Notification
// @Router /notifications/{id} [get]
// @Security ApiKeyAuth
func (c Controller) NotificationByID(id int, who int, whoKind string) (interface{}, error) {
	switch whoKind {
	case repository.TeacherUser:
		return c.repo.NotificationByIDForTeacher(id, who, whoKind)
	case repository.ParentUser:
		return c.repo.NotificationByIDForParent(id, who, whoKind)
	case repository.AdminUser:
		return c.repo.NotificationByIDForAdmin(id, who, whoKind)
	default:
		return nil, ErrorNotAuthorized
	}
}

// Parents
// see/modify their personal data
// @Summary Get a parent by id
// @Param id path int true "Account ID"
// @Tags Parents
// @Success 200 {object} models.Parent
// @Router /parents/{id} [get]
// @Security ApiKeyAuth
func (c Controller) ParentByID(id int, who int, whoKind string) (interface{}, error) {
	switch whoKind {
	case repository.ParentUser:
		if id == who {
			return c.repo.ParentByID(id)
		} else {
			return nil, ErrorNotAuthorized
		}
	case repository.AdminUser:
		return c.repo.ParentByID(id)
	default:
		return nil, ErrorNotAuthorized
	}
}

// Get student by id
// @Summary Get a student by id
// @Param id path int true "Student ID"
// @Tags Students
// @Success 200 {object} models.Student
// @Router /students/{id} [get]
// @Security ApiKeyAuth
func (c Controller) StudentByID(id int, who int, whoKind string) (student interface{}, err error) {
	switch whoKind {
	case repository.ParentUser:
		return c.repo.StudentByIDForParent(id, who)
	case repository.AdminUser:
		return c.repo.StudentByIDForAdmin(id)
	default:
		return nil, ErrorNotAuthorized
	}
}

// @Summary Get a lecture by id
// @Param id path int true "Lecture ID"
// @Tags Lectures
// @Success 200 {object} models.TimeTable
// @Router /lectures/{id} [get]
// @Security ApiKeyAuth
func (c Controller) LectureByID(id int, who int, whoKind string) (interface{}, error) {

	switch whoKind {
	case repository.ParentUser:
		return c.repo.LectureByIDForParent(id, who)
	case repository.TeacherUser:
		return c.repo.LectureByIDForTeacher(id, who)

	case repository.AdminUser:
		return c.repo.LectureByIDForAdmin(id)

	default:
		return nil, ErrorNotAuthorized
	}
}
