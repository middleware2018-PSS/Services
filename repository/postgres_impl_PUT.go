package repository

import (
	"github.com/middleware2018-PSS/Services/models"
	"golang.org/x/crypto/bcrypt"
)

func (r *postgresRepository) UpdateTeacherForTeacher(teacher models.Teacher, who int) (err error) {

	query := "UPDATE back2school.teachers " +
		"SET name = $1, surname = $2, mail = $3, info = $4 " +
		"where id = $5 "
	return r.execUpdate(query, teacher.Name, teacher.Surname, teacher.Mail, teacher.Info, who)

}

func (r *postgresRepository) UpdateTeacherForAdmin(teacher models.Teacher) (err error) {
	query := "UPDATE back2school.teachers " +
		"SET name = $1, surname = $2, mail = $3, info = $4 " +
		"where id = $5 "

	return r.execUpdate(query, teacher.Name, teacher.Surname, teacher.Mail, teacher.Info, teacher.ID)
}

func (r *postgresRepository) UpdateParentForParent(parent models.Parent, who int) (err error) {

	query := "UPDATE back2school.parents " +
		"SET name = $1, surname = $2, mail = $3, info = $4 " +
		"where id = $5 "

	return r.execUpdate(query, parent.Name, parent.Surname, parent.Mail, parent.Info, who)
}

func (r *postgresRepository) UpdateParentForAdmin(parent models.Parent) (err error) {

	query := "UPDATE back2school.parents " +
		" SET name = $1, surname = $2, mail = $3, info = $4 " +
		"where id = $5 "

	return r.execUpdate(query, parent.Name, parent.Surname, parent.Mail, parent.Info, parent.ID)
}

func (r *postgresRepository) UpdateStudentForParent(student models.Student, who int) (err error) {

	query := "IF $1 in (select parent from back2school.isParent where student = $2) UPDATE back2school.student " +
		" ET name = $3, surname = $4, mail = $5, info = $6 " +
		"where id = $7 "

	return r.execUpdate(query, who, student.ID, student.Name, student.Surname, student.Mail, student.Info, student.ID)
}

func (r *postgresRepository) UpdateStudentForAdmin(student models.Student) (err error) {

	query := "UPDATE back2school.student " +
		"SET name = $1, surname = $2, mail = $3, info = $4 " +
		"where id = $5 "

	return r.execUpdate(query, student.Name, student.Surname, student.Mail, student.Info, student.ID)

}

func (r *postgresRepository) UpdateAppointmentForParent(appointment models.Appointment, who int) (err error) {

	query := "if $1 in (select parent from back2school.isParent where student = $2) UPDATE back2school.appointments " +
		"SET student = $3, teacher = $4, location = $5, time = $6 where id = $7 "

	return r.execUpdate(query, who, appointment.Student, appointment.Student, appointment.Teacher, appointment.Location, appointment.Time, appointment.ID)
}
func (r *postgresRepository) UpdateAppointmentForTeacher(appointment models.Appointment, who int) (err error) {
	query := "UPDATE back2school.appointments " +
		"SET student = $1, teacher = $2, location = $3, time = $4 where id = $5 "

	return r.execUpdate(query, appointment.Student, who, appointment.Location, appointment.Time, appointment.ID)

}

func (r *postgresRepository) UpdateAppointmentForAdmin(appointment models.Appointment) (err error) {

	query := "UPDATE back2school.appointments " +
		"SET student = $1, teacher = $2, location = $3, time = $4 where id = $5 "

	return r.execUpdate(query, appointment.Student, appointment.Teacher, appointment.Location, appointment.Time, appointment.ID)
}

func (r *postgresRepository) UpdateClassForAdmin(class models.Class) (err error) {
	query := "UPDATE back2school.classes " +
		" SET year = $1, section = $2, info = $3, grade = $4 " +
		" where id = $5 "
	return r.execUpdate(query, class.Year, class.Section, class.Info, class.Grade, class.ID)

}

func (r *postgresRepository) UpdateNotificationForAdmin(notification models.Notification) error {
	query := "UPDATE back2school.notification " +
		"SET receiver = $1, message = $2, time = $3, receiver_kind = $4 " +
		" where id = $5 "
	return r.execUpdate(query, notification.Receiver, notification.Message, notification.Time, notification.ReceiverKind, notification.ID)
}

func (r *postgresRepository) UpdateGradeForTeacher(grade models.Grade) error {

	query := "UPDATE back2school.grades " +
		"SET student = $1, grade = $2, subject = $3, date = $4, teacher = $5) " +
		" where id = $6 "

	return r.execUpdate(query, grade.Student, grade.Grade, grade.Subject, grade.Date, grade.Teacher, grade.ID)
}

func (r *postgresRepository) UpdateGradeForAdmin(grade models.Grade) error {

	query := "UPDATE back2school.grades " +
		"SET student = $1, grade = $2, subject = $3, date = $4, teacher = $5) " +
		" where id = $6 "

	return r.execUpdate(query, grade.Student, grade.Grade, grade.Subject, grade.Date, grade.Teacher)

}

func (r *postgresRepository) UpdatePaymentForParent(payment models.Payment, who int) error {

	query := "if $7 in (select parent from back2school.isParent where student = $2) UPDATE back2school.grades " +
		"SET amount = $1, student = $2, paid = $3, reason = $4, emitted = $5 " +
		" where id = $6 "

	return r.execUpdate(query, payment.Amount, payment.Student, payment.Paid, payment.Reason, payment.Emitted, payment.ID, who)

}

func (r *postgresRepository) UpdatePaymentForAdmin(payment models.Payment) error {

	query := "UPDATE back2school.grades " +
		"SET amount = $1, student = $2, paid = $3, reason = $4, emitted = $5 " +
		" where id = $6 "

	return r.execUpdate(query, payment.Amount, payment.Student, payment.Paid, payment.Reason, payment.Emitted, payment.ID)
}

func (r *postgresRepository) UpdateLectureForTeacher(lecture models.TimeTable, who int) error {

	query := "if $1 in (select teacher from back2school.teaches natural join timetable where id = $2) UPDATE back2school.timetable " +
		"SET location = $3, start = $4, end = $5, info = $6 " +
		" where id = $2"

	return r.execUpdate(query, who, lecture.ID, lecture.Location, lecture.Start, lecture.End, lecture.Info)
}

func (r *postgresRepository) UpdateLectureForadmin(lecture models.TimeTable) error {

	query := "UPDATE back2school.timetable " +
		"SET location = $2, start = $3, end = $4, info = $5 " +
		" where id = $1"

	return r.execUpdate(query, lecture.ID, lecture.Location, lecture.Start, lecture.End, lecture.Info)
}

func (r *postgresRepository) UpdateAccount(account models.Account, cost int) error {

	encryptedPassword, _ := bcrypt.GenerateFromPassword([]byte(account.Password), cost)
	query := "UPDATE back2school.accounts SET password = $1 " +
		"WHERE username = $2 AND id = $3 AND kind = $4"

	return r.execUpdate(query, string(encryptedPassword), account.Username, account.ID, account.Kind)
}
