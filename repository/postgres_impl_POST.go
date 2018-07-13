package repository

import (
	"log"

	"github.com/middleware2018-PSS/Services/models"
)

func (r *Repository) CreateAccount(username string, password []byte, id int, kind string, cost int) error {

	query := `INSERT INTO back2school.accounts (username, "password", id, kind) VALUES ($1, $2, $3, $4)`

	_, err := r.Exec(query, username, password, id, kind)
	if err != nil {
		log.Printf("%v", err.Error())
	}
	return err

}

func (r *Repository) CreateAppointmentForParent(appointment models.Appointment, who int) (id int, err error) {

	query := "if $1 in (select parent from back2school.isParent where student = $2) INSERT INTO back2school.appointments " +
		" (student, teacher, location, time) VALUES ($3, $4, $5, $6) "

	return r.exec(query, who, appointment.Student, appointment.Teacher, appointment.Location, appointment.Time)
}

func (r *Repository) CreateAppointmentForTeacher(appointment models.Appointment, who int) (id int, err error) {

	query := "INSERT INTO back2school.appointments " +
		" (student, teacher, location, time) VALUES ($1, $2, $3, $4) "
	return r.exec(query, appointment.Student, who, appointment.Location, appointment.Time)
}

func (r *Repository) CreateAppointmentForAdmin(appointment models.Appointment) (id int, err error) {

	query := "INSERT INTO back2school.appointments " +
		" (student, teacher, location, time) VALUES ($1, $2, $3, $4) "
	return r.exec(query, appointment.Student, appointment.Teacher, appointment.Location, appointment.Time)
}

func (r *Repository) CreateParentForAdmin(parent models.Parent) (int, error) {
	query := "INSERT INTO back2school.parents " +
		"(name, surname, mail, info) VALUES ($1, $2, $3, $4) "
	return r.exec(query, parent.Name, parent.Surname, parent.Mail, parent.Info)
}

func (r *Repository) CreateTeacherForAdmin(teacher models.Teacher) (int, error) {
	query := "INSERT INTO back2school.teachers " +
		" (name, surname, mail, info) " +
		" VALUES ($1, $2, $3, $4) "
	return r.exec(query, teacher.Name, teacher.Surname, teacher.Mail, teacher.Info)

}

func (r *Repository) CreateStudentForAdmin(student models.Student) (int, error) {
	query := "INSERT INTO back2school.students " +
		" (name, surname, mail, info) " +
		" VALUES ($1, $2, $3, $4) "
	return r.exec(query, student.Name, student.Surname, student.Mail, student.Info)

}

func (r *Repository) CreateClassForAdmin(class models.Class) (int, error) {
	query := "INSERT INTO back2school.classes " +
		" (year, section, info, grade) " +
		" VALUES ($1, $2, $3, $4) "
	return r.exec(query, class.Year, class.Section, class.Info, class.Grade)
}

func (r *Repository) CreateNotificationForAdmin(notification models.Notification) (int, error) {
	query := "insert into back2school.classes " +
		" (receiver, message, time, receiver_kind) " +
		" VALUES ($1, $2, $3, $4) "
	return r.exec(query, notification.Receiver, notification.Message, notification.Time, notification.ReceiverKind)

}

func (r *Repository) CreateGradeForTeacher(grade models.Grade) (int, error) {

	query := "if ($5, $1, $3) in " +
		"(select t.teacher, t.subject, e.student from back2school.teaches as t natural join back2school.enrolled as e ) " +
		" insert into back2school.grades " +
		" (student, grade, subject, date, teacher) " +
		" VALUES ($1, $2, $3, $4, $5) "

	return r.exec(query, grade.Student, grade.Grade, grade.Subject, grade.Date, grade.Teacher)
}

func (r *Repository) CreateGradeForAdmin(grade models.Grade) (int, error) {
	query := "insert into back2school.grades " +
		" (student, grade, subject, date, teacher) " +
		" VALUES ($1, $2, $3, $4, $5) "

	return r.exec(query, grade.Student, grade.Grade, grade.Subject, grade.Date, grade.Teacher)
}

func (r *Repository) CreatePaymentForAdmin(payment models.Payment) (int, error) {
	query := "insert into back2school.payments " +
		" (amount, student, paid, reason, emitted) " +
		" VALUES ($1, $2, $3, $4, $5) "
	return r.exec(query, payment.Amount, payment.Student, payment.Paid, payment.Reason, payment.Emitted)

}

func (r *Repository) CreateLectureForAdmin(lecture models.TimeTable) (int, error) {
	query := "insert into back2school.timetable " +
		" (class, subject, location, start, end, info) " +
		" VALUES ($1, $2, $3, $4, $5, $6) "
	return r.exec(query, lecture.Class, lecture.Subject, lecture.Location, lecture.Start, lecture.End, lecture.Info)
}
