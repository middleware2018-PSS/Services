package repository

import (
	"github.com/middleware2018-PSS/Services/src/models"
	"golang.org/x/crypto/bcrypt"
	"log"
)

// @Summary Create an account
// @Param Account body models.User true "data"
// @Tags Accounts
// @Router /accounts [post]
// @Security ApiKeyAuth
func (r *postgresRepository) CreateAccount(username string, password string, id int, kind string, cost int, whoKind string) error {
	if whoKind == AdminUser {
		query := `INSERT INTO back2school.accounts ("user", "password", id, kind) VALUES ($1, $2, $3, $4)`
		cryptedPass, err := bcrypt.GenerateFromPassword([]byte(password), cost)
		if err != nil {
			return err
		}
		_, err = r.Exec(query, username, cryptedPass, id, kind)
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
func (r *postgresRepository) CreateAppointment(appointment models.Appointment, who int, whoKind string) (id int, err error) {
	var query string
	var args []interface{}
	switch whoKind {
	case ParentUser:
		query = "if $1 in (select parent from back2school.isParent where student = $2) INSERT INTO back2school.appointments " +
			" (student, teacher, location, time) VALUES ($3, $4, $5, $6) "
		args = append(args, who, appointment.Student, appointment.Teacher, appointment.Location, appointment.Time)
	case TeacherUser:
		if *appointment.Teacher == who {
			query = "INSERT INTO back2school.appointments " +
				" (student, teacher, location, time) VALUES ($1, $2, $3, $4) "
			args = append(args, appointment.Student, who, appointment.Location, appointment.Time)
		} else {
			return 0, ErrorNotAuthorized
		}
	case AdminUser:
		query = "INSERT INTO back2school.appointments " +
			" (student, teacher, location, time) VALUES ($1, $2, $3, $4) "
		args = append(args, appointment.Student, appointment.Teacher, appointment.Location, appointment.Time)
	default:
		return 0, ErrorNotAuthorized
	}
	return r.exec(query, args...)
}

// @Summary Create parent
// @Tags Parents
// @Param parent body models.Parent true "data"
// @Tags Parents
// @Success 201 {object} models.Parent
// @Router /parents [post]
// @Security ApiKeyAuth
func (r *postgresRepository) CreateParent(parent models.Parent, who int, whoKind string) (int, error) {
	if whoKind == AdminUser {
		query := "INSERT INTO back2school.parents " +
			"(name, surname, mail, info) VALUES ($1, $2, $3, $4) "
		return r.exec(query, parent.Name, parent.Surname, parent.Mail, parent.Info)
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
func (r *postgresRepository) CreateTeacher(teacher models.Teacher, who int, whoKind string) (int, error) {
	if whoKind == AdminUser {
		query := "INSERT INTO back2school.teachers " +
			" (name, surname, mail, info) " +
			" VALUES ($1, $2, $3, $4) "
		return r.exec(query, teacher.Name, teacher.Surname, teacher.Mail, teacher.Info)
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
func (r *postgresRepository) CreateStudent(student models.Student, who int, whoKind string) (int, error) {
	if whoKind == AdminUser {
		query := "INSERT INTO back2school.students " +
			" (name, surname, mail, info) " +
			" VALUES ($1, $2, $3, $4) "
		return r.exec(query, student.Name, student.Surname, student.Mail, student.Info)
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
func (r *postgresRepository) CreateClass(class models.Class, who int, whoKind string) (int, error) {
	if whoKind == AdminUser {
		query := "INSERT INTO back2school.classes " +
			" (year, section, info, grade) " +
			" VALUES ($1, $2, $3, $4) "
		return r.exec(query, class.Year, class.Section, class.Info, class.Grade)
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
func (r *postgresRepository) CreateNotification(notification models.Notification, who int, whoKind string) (int, error) {
	if whoKind == AdminUser {
		query := "insert into back2school.classes " +
			" (receiver, message, time, receiver_kind) " +
			" VALUES ($1, $2, $3, $4) "
		return r.exec(query, notification.Receiver, notification.Message, notification.Time, notification.ReceiverKind)
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
func (r *postgresRepository) CreateGrade(grade models.Grade, who int, whoKind string) (int, error) {
	var query string
	var args []interface{}
	switch whoKind {
	case TeacherUser:
		if *grade.Teacher == who {
			query = "if ($5, $1, $3) in " +
				"(select t.teacher, t.subject, e.student from back2school.teaches as t natural join back2school.enrolled as e ) " +
				" insert into back2school.grades " +
				" (student, grade, subject, date, teacher) " +
				" VALUES ($1, $2, $3, $4, $5) "
			args = append(args, grade.Student, grade.Grade, grade.Subject, grade.Date, grade.Teacher)
		} else {
			return 0, ErrorNotAuthorized
		}
	case AdminUser:
		query = "insert into back2school.grades " +
			" (student, grade, subject, date, teacher) " +
			" VALUES ($1, $2, $3, $4, $5) "
		args = append(args, grade.Student, grade.Grade, grade.Subject, grade.Date, grade.Teacher)
	default:
		return 0, ErrorNotAuthorized
	}

	return r.exec(query, args...)
}

// @Summary Create payment
// @Param class body models.Payment true "data"
// @Tags Payments
// @Router /payments [post]
// @Success 201 {object} models.Payment
// @Security ApiKeyAuth
func (r *postgresRepository) CreatePayment(payment models.Payment, who int, whoKind string) (int, error) {
	if whoKind == AdminUser {
		query := "insert into back2school.payments " +
			" (amount, student, payed, reason, emitted) " +
			" VALUES ($1, $2, $3, $4, $5) "
		return r.exec(query, payment.Amount, payment.Student, payment.Payed, payment.Reason, payment.Emitted)
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
func (r *postgresRepository) CreateLecture(lecture models.TimeTable, who int, whoKind string) (int, error) {
	if whoKind == AdminUser {
		query := "insert into back2school.timetable " +
			" (class, subject, location, start, end, info) " +
			" VALUES ($1, $2, $3, $4, $5, $6) "
		return r.exec(query, lecture.Class, lecture.Subject, lecture.Location, lecture.Start, lecture.End, lecture.Info)
	} else {
		return 0, ErrorNotAuthorized
	}
}
