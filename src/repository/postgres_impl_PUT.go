package repository

import "github.com/middleware2018-PSS/Services/src/models"

// @Summary Update teacher's data
// @Param id path int true "Teacher ID"
// @Param teacher body models.Teacher true "data"
// @Tags Teachers
// @Success 204 {object} models.Teacher
// @Router /teachers/{id} [put]
func (r *postgresRepository) UpdateTeacher(teacher models.Teacher, who int, whoKind string) (err error) {
	var query string
	var args []interface{}
	switch whoKind {
	case TeacherUser:
		if teacher.ID == who {
			query = "UPDATE back2school.teachers " +
				"SET name = $1, surname = $2, mail = $3, info = $4 " +
				"where id = $5 "
			args = append(args, teacher.Name, teacher.Surname, teacher.Mail, teacher.Info, who)
		} else {
			return ErrorNotAuthorized
		}
	case AdminUser:
		query = "UPDATE back2school.teachers " +
			"SET name = $1, surname = $2, mail = $3, info = $4 " +
			"where id = $5 "
		args = append(args, teacher.Name, teacher.Surname, teacher.Mail, teacher.Info, teacher.ID)
	default:
		return ErrorNotAuthorized
	}
	return r.execUpdate(query, args...)
}

// @Summary Update parents's data
// @Param id path int true "Parent ID"
// @Param parent body models.Parent true "data"
// @Tags Parents
// @Success 201 {object} models.Parent
// @Router /parents/{id} [put]
func (r *postgresRepository) UpdateParent(parent models.Parent, who int, whoKind string) (err error) {
	var query string
	var args []interface{}
	switch whoKind {
	case ParentUser:
		if parent.ID == who {
			query = "UPDATE back2school.parents " +
				"SET name = $1, surname = $2, mail = $3, info = $4 " +
				"where id = $5 "
			args = append(args, parent.Name, parent.Surname, parent.Mail, parent.Info, who)
		} else {
			return ErrorNotAuthorized
		}
	case AdminUser:
		query = "UPDATE back2school.parents " +
			" SET name = $1, surname = $2, mail = $3, info = $4 " +
			"where id = $5 "
		args = append(args, parent.Name, parent.Surname, parent.Mail, parent.Info, parent.ID)
	default:
		return ErrorNotAuthorized
	}

	return r.execUpdate(query, args...)
}

// @Summary Update student's data
// @Param id path int true "Student ID"
// @Param student body models.Student true "data"
// @Tags Students
// @Success 201 {object} models.Student
// @Router /students/{id} [put]
func (r *postgresRepository) UpdateStudent(student models.Student, who int, whoKind string) (err error) {
	var query string
	var args []interface{}
	switch whoKind {
	case ParentUser:
		query = "IF $1 in (select parent from back2school.isParent where student = $2) UPDATE back2school.student " +
			" ET name = $3, surname = $4, mail = $5, info = $6 " +
			"where id = $7 "
		args = append(args, who, student.ID, student.Name, student.Surname, student.Mail, student.Info, student.ID)
	case AdminUser:
		query = "UPDATE back2school.student " +
			"SET name = $1, surname = $2, mail = $3, info = $4 " +
			"where id = $5 "
		args = append(args, student.Name, student.Surname, student.Mail, student.Info, student.ID)
	default:
		return ErrorNotAuthorized
	}
	return r.execUpdate(query, args...)
}

// @Summary Update appointment's data
// @Param id path int true "Appointment ID"
// @Param appointment body models.Appointment true "data"
// @Tags Appointments
// @Success 201 {object} models.Appointment
// @Router /appointments/{id} [put]
func (r *postgresRepository) UpdateAppointment(appointment models.Appointment, who int, whoKind string) (err error) {
	var query string
	var args []interface{}
	switch whoKind {
	case ParentUser:
		query = "if $1 in (select parent from back2school.isParent where student = $2) UPDATE back2school.appointments " +
			"SET student = $3, teacher = $4, location = $5, time = $6 where id = $7 "
		args = append(args, who, appointment.Student.ID, appointment.Student.ID, appointment.Teacher.ID, appointment.Location, appointment.Time, appointment.ID)
	case TeacherUser:
		if appointment.Teacher.ID == who {
			query = "UPDATE back2school.appointments " +
				"SET student = $1, teacher = $2, location = $3, time = $4 where id = $5 "
			args = append(args, appointment.Student.ID, who, appointment.Location, appointment.Time, appointment.ID)
		}
	case AdminUser:
		query = "UPDATE back2school.appointments " +
			"SET student = $1, teacher = $2, location = $3, time = $4 where id = $5 "
		args = append(args, appointment.Student.ID, appointment.Teacher.ID, appointment.Location, appointment.Time, appointment.ID)
	default:
		return ErrorNotAuthorized
	}
	return r.execUpdate(query, args...)
}

// @Summary Update Class's data
// @Param id path int true "Class ID"
// @Param parent body models.Class true "data"
// @Tags Classes
// @Success 201 {object} models.Class
// @Router /classes/{id} [put]
func (r *postgresRepository) UpdateClass(class models.Class, who int, whoKind string) (err error) {
	if whoKind == AdminUser {
		query := "UPDATE back2school.classes " +
			" SET year = $1, section = $2, info = $3, grade = $4 " +
			" where id = $5 "
		return r.execUpdate(query, class.Year, class.Section, class.Info, class.Grade, class.ID)
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
func (r *postgresRepository) UpdateNotification(notification models.Notification, who int, whoKind string) error {
	if whoKind == AdminUser {
		query := "UPDATE back2school.notification " +
			"SET receiver = $1, message = $2, time = $3, receiver_kind = $4 " +
			" where id = $5 "
		return r.execUpdate(query, notification.Receiver, notification.Message, notification.Time, notification.ReceiverKind, notification.ID)
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
func (r *postgresRepository) UpdateGrade(grade models.Grade, who int, whoKind string) error {
	var query string
	var args []interface{}
	switch whoKind {
	case TeacherUser:
		if grade.Teacher.ID == who {
			query = "UPDATE back2school.grades " +
				"SET student = $1, grade = $2, subject = $3, date = $4, teacher = $5) " +
				" where id = $6 "
			args = append(args, grade.Student.ID, grade.Grade, grade.Subject, grade.Date, grade.Teacher.ID, grade.ID)
		} else {
			return ErrorNotAuthorized
		}
	case AdminUser:
		query = "UPDATE back2school.grades " +
			"SET student = $1, grade = $2, subject = $3, date = $4, teacher = $5) " +
			" where id = $6 "
		args = append(args, grade.Student.ID, grade.Grade, grade.Subject, grade.Date, grade.Teacher.ID)
	default:
		return ErrorNotAuthorized
	}
	return r.execUpdate(query, args...)
}

// @Summary Update payment
// @Param id path int true "Payment ID"
// @Param class body models.Payment true "data"
// @Tags Payments
// @Router /payments/{id} [put]
// @Success 201 {object} models.Payment
func (r *postgresRepository) UpdatePayment(payment models.Payment, who int, whoKind string) error {
	var query string
	var args []interface{}
	switch whoKind {
	case TeacherUser:
		query = "if $7 in (select parent from back2school.isParent where student = $2) UPDATE back2school.grades " +
			"SET amount = $1, student = $2, payed = $3, reason = $4, emitted = $5 " +
			" where id = $6 "
		args = append(args, payment.Amount, payment.Student.ID, payment.Payed, payment.Reason, payment.Emitted, payment.ID, who)

	case AdminUser:
		query = "UPDATE back2school.grades " +
			"SET amount = $1, student = $2, payed = $3, reason = $4, emitted = $5 " +
			" where id = $6 "
		args = append(args, payment.Amount, payment.Student.ID, payment.Payed, payment.Reason, payment.Emitted, payment.ID)
	default:
		return ErrorNotAuthorized
	}
	return r.execUpdate(query, args...)
}
