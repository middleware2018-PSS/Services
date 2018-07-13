package controller

import (
	"log"

	"github.com/middleware2018-PSS/Services/models"
	"github.com/middleware2018-PSS/Services/repository"
	"golang.org/x/crypto/bcrypt"
)

// @Summary Create an account
// @Param Account body models.User true "data"
// @Tags Accounts
// @Router /accounts [post]
// @Security ApiKeyAuth
func (c Controller) CreateAccount(username string, password string, id int, kind string, cost int, whoKind string) error {
	if whoKind == repository.AdminUser {
		if !allowedKind[kind] {
			return ErrorNoKindSpecified
		}
		cryptedPass, err := bcrypt.GenerateFromPassword([]byte(password), cost)
		if err != nil {
			return err
		}
		err = c.repo.CreateAccount(username, cryptedPass, id, kind, cost)
		if err != nil {
			log.Printf("%v", err.Error())
		}
		return err
	} else {
		return ErrorNotAuthorized
	}
}

// @Summary Create appointment
// @Param id path int true "Appointment ID"
// @Param appointment body models.Appointment true "data"
// @Tags Appointments
// @Router /appointments [post]
// @Success 201 {object} models.Appointment
// @Security ApiKeyAuth
func (c Controller) CreateAppointment(appointment models.Appointment, who int, whoKind string) (id int, err error) {
	switch whoKind {
	case repository.ParentUser:
		return c.repo.CreateAppointmentForParent(appointment, who)
	case repository.TeacherUser:
		if *appointment.Teacher == who {
			return c.repo.CreateAppointmentForTeacher(appointment, who)
		} else {
			return 0, ErrorNotAuthorized
		}
	case repository.AdminUser:
		return c.repo.CreateAppointmentForAdmin(appointment)
	default:
		return 0, ErrorNotAuthorized
	}
}

// @Summary Create parent
// @Tags Parents
// @Param parent body models.Parent true "data"
// @Tags Parents
// @Success 201 {object} models.Parent
// @Router /parents [post]
// @Security ApiKeyAuth
func (c Controller) CreateParent(parent models.Parent, who int, whoKind string) (int, error) {
	if whoKind == repository.AdminUser {
		return c.repo.CreateParentForAdmin(parent)
	} else {
		return 0, ErrorNotAuthorized
	}
}

// @Summary Create teacher
// @Param teacher body models.Teacher true "data"
// @Tags Teachers
// @Router /teachers [post]
// @Success 201 {object} models.Teacher
// @Security ApiKeyAuth
func (c Controller) CreateTeacher(teacher models.Teacher, who int, whoKind string) (int, error) {
	if whoKind == repository.AdminUser {
		return c.repo.CreateTeacherForAdmin(teacher)
	} else {
		return 0, ErrorNotAuthorized
	}
}

// @Summary Create student
// @Param student body models.Student true "data"
// @Tags Students
// @Router /students [post]
// @Success 201 {object} models.Student
// @Security ApiKeyAuth
func (c Controller) CreateStudent(student models.Student, who int, whoKind string) (int, error) {
	if whoKind == repository.AdminUser {
		return c.repo.CreateStudentForAdmin(student)
	} else {
		return 0, ErrorNotAuthorized
	}

}

// @Summary Create class
// @Param class body models.Class true "data"
// @Tags Classes
// @Router /classes [post]
// @Success 201 {object} models.Class
// @Security ApiKeyAuth
func (c Controller) CreateClass(class models.Class, who int, whoKind string) (int, error) {
	if whoKind == repository.AdminUser {
		return c.repo.CreateClassForAdmin(class)
	} else {
		return 0, ErrorNotAuthorized
	}

}

// @Summary Create notification
// @Param class body models.Notification true "data"
// @Tags Notifications
// @Router /notifications [post]
// @Success 201 {object} models.Notification
// @Security ApiKeyAuth
func (c Controller) CreateNotification(notification models.Notification, who int, whoKind string) (int, error) {
	if whoKind == repository.AdminUser {
		return c.repo.CreateNotificationForAdmin(notification)
	} else {
		return 0, ErrorNotAuthorized
	}

}

// @Summary Create grade
// @Param class body models.Grade true "data"
// @Tags Grades
// @Router /grades [post]
// @Success 201 {object} models.Grade
// @Security ApiKeyAuth
func (c Controller) CreateGrade(grade models.Grade, who int, whoKind string) (int, error) {
	switch whoKind {
	case repository.TeacherUser:
		if *grade.Teacher == who {
			return c.repo.CreateGradeForTeacher(grade)
		} else {
			return 0, ErrorNotAuthorized
		}
	case repository.AdminUser:
		return c.repo.CreateGradeForAdmin(grade)
	default:
		return 0, ErrorNotAuthorized
	}
}

// @Summary Create payment
// @Param class body models.Payment true "data"
// @Tags Payments
// @Router /payments [post]
// @Success 201 {object} models.Payment
// @Security ApiKeyAuth
func (c Controller) CreatePayment(payment models.Payment, who int, whoKind string) (int, error) {
	if whoKind == repository.AdminUser {
		return c.repo.CreatePaymentForAdmin(payment)
	} else {
		return 0, ErrorNotAuthorized
	}
}

// @Summary Create lecture
// @Param class body models.TimeTable true "data"
// @Tags Lectures
// @Router /lectures [post]
// @Success 201 {object} models.Payment
// @Security ApiKeyAuth
func (c Controller) CreateLecture(lecture models.TimeTable, who int, whoKind string) (int, error) {
	if whoKind == repository.AdminUser {
		return c.repo.CreateLectureForAdmin(lecture)
	} else {
		return 0, ErrorNotAuthorized
	}
}
