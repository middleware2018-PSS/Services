package controller

import (
	"github.com/middleware2018-PSS/Services/models"
	"github.com/middleware2018-PSS/Services/repository"
)

// @Summary Update teacher's data
// @Param id path int true "Teacher ID"
// @Param teacher body models.Teacher true "data"
// @Tags Teachers
// @Success 204 {object} models.Teacher
// @Router /teachers/{id} [put]
// @Security ApiKeyAuth
func (c Controller) UpdateTeacher(teacher models.Teacher, who int, whoKind string) (err error) {
	switch whoKind {
	case repository.TeacherUser:
		if teacher.ID == who {
			return c.repo.UpdateTeacherForTeacher(teacher, who)
		} else {
			return ErrorNotAuthorized
		}
	case repository.AdminUser:
		return c.repo.UpdateTeacherForAdmin(teacher)
	default:
		return ErrorNotAuthorized
	}
}

// @Summary Update parents's data
// @Param id path int true "Parent ID"
// @Param parent body models.Parent true "data"
// @Tags Parents
// @Success 201 {object} models.Parent
// @Router /parents/{id} [put]
// @Security ApiKeyAuth
func (c Controller) UpdateParent(parent models.Parent, who int, whoKind string) (err error) {

	switch whoKind {
	case repository.ParentUser:
		if parent.ID == who {
			return c.repo.UpdateParentForParent(parent, who)
		} else {
			return ErrorNotAuthorized
		}
	case repository.AdminUser:
		return c.repo.UpdateParentForAdmin(parent)

	default:
		return ErrorNotAuthorized
	}
}

// @Summary Update student's data
// @Param id path int true "Student ID"
// @Param student body models.Student true "data"
// @Tags Students
// @Success 201 {object} models.Student
// @Router /students/{id} [put]
// @Security ApiKeyAuth
func (c Controller) UpdateStudent(student models.Student, who int, whoKind string) (err error) {

	switch whoKind {
	case repository.ParentUser:
		return c.repo.UpdateStudentForParent(student, who)
	case repository.AdminUser:
		return c.repo.UpdateStudentForAdmin(student)
	default:
		return ErrorNotAuthorized
	}
}

// @Summary Update appointment's data
// @Param id path int true "Appointment ID"
// @Param appointment body models.Appointment true "data"
// @Tags Appointments
// @Success 201 {object} models.Appointment
// @Router /appointments/{id} [put]
// @Security ApiKeyAuth
func (c Controller) UpdateAppointment(appointment models.Appointment, who int, whoKind string) (err error) {

	switch whoKind {
	case repository.ParentUser:
		return c.repo.UpdateAppointmentForParent(appointment, who)
	case repository.TeacherUser:
		if *appointment.Teacher == who {
			return c.repo.UpdateAppointmentForTeacher(appointment, who)
		} else {
			return ErrorNotAuthorized
		}
	case repository.AdminUser:
		return c.repo.UpdateAppointmentForAdmin(appointment)
	default:
		return ErrorNotAuthorized
	}
}

// @Summary Update Class's data
// @Param id path int true "Class ID"
// @Param parent body models.Class true "data"
// @Tags Classes
// @Success 201 {object} models.Class
// @Router /classes/{id} [put]
// @Security ApiKeyAuth
func (c Controller) UpdateClass(class models.Class, who int, whoKind string) (err error) {
	if whoKind == repository.AdminUser {
		return c.repo.UpdateClassForAdmin(class)
	} else {
		return ErrorNotAuthorized
	}

}

// @Summary Update notification
// @Param id path int true "Notification ID"
// @Param class body models.Notification true "data"
// @Tags Notifications
// @Router /notifications/{id} [put]
// @Success 201 {object} models.Notification
// @Security ApiKeyAuth
func (c Controller) UpdateNotification(notification models.Notification, who int, whoKind string) error {
	if whoKind == repository.AdminUser {
		return c.repo.UpdateNotificationForAdmin(notification)
	} else {
		return ErrorNotAuthorized
	}

}

// @Summary Update Grade
// @Param id path int true "Grade ID"
// @Param class body models.Grade true "data"
// @Tags Grades
// @Router /grades/{id} [put]
// @Success 201 {object} models.Grade
// @Security ApiKeyAuth
func (c Controller) UpdateGrade(grade models.Grade, who int, whoKind string) error {

	switch whoKind {
	case repository.TeacherUser:
		if *grade.Teacher == who {
			return c.repo.UpdateGradeForTeacher(grade)
		} else {
			return ErrorNotAuthorized
		}
	case repository.AdminUser:
		return c.repo.UpdateGradeForAdmin(grade)
	default:
		return ErrorNotAuthorized
	}
}

// @Summary Update payment
// @Param id path int true "Payment ID"
// @Param class body models.Payment true "data"
// @Tags Payments
// @Router /payments/{id} [put]
// @Success 201 {object} models.Payment
// @Security ApiKeyAuth
func (c Controller) UpdatePayment(payment models.Payment, who int, whoKind string) error {

	switch whoKind {
	case repository.ParentUser:
		return c.repo.UpdatePaymentForParent(payment, who)

	case repository.AdminUser:
		return c.repo.UpdatePaymentForAdmin(payment)

	default:
		return ErrorNotAuthorized
	}
}

// @Summary Update lecture
// @Param id path int true "Lecture ID"
// @Param class body models.TimeTable true "data"
// @Tags Lectures
// @Router /lectures/{id} [put]
// @Success 201 {object} models.TimeTable
// @Security ApiKeyAuth
func (c Controller) UpdateLecture(lecture models.TimeTable, who int, whoKind string) error {

	switch whoKind {
	case repository.TeacherUser:
		return c.repo.UpdateLectureForTeacher(lecture, who)
	case repository.AdminUser:
		return c.repo.UpdateLectureForadmin(lecture)
	default:
		return ErrorNotAuthorized
	}
}

// @Summary Update Account
// @Param class body models.Account true "data"
// @Tags Accounts
// @Router /accounts [put]
// @Security ApiKeyAuth
func (c Controller) UpdateAccount(account models.Account, who int, whoKind string, cost int) error {

	switch whoKind {
	case repository.TeacherUser, repository.ParentUser:
		if account.Kind != whoKind || account.ID != who {
			return ErrorNotAuthorized
		}
		return c.repo.UpdateAccount(account, cost)
	case repository.AdminUser:
		return c.repo.UpdateAccount(account, cost)
	default:
		return ErrorNotAuthorized
	}
}
