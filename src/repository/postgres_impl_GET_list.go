package repository

import (
	"database/sql"
	_ "github.com/middleware2018-PSS/Services/src/docs"
	"github.com/middleware2018-PSS/Services/src/models"
)

type Subjects struct {
	Subjects []string `json:"subjects" example:"science"`
}

// @Summary Get all classes
// @Param limit query int false "number of elements to return"
// @Param offset query int false "offset in the list of elements to return"
// @Tags Classes
// @Success 200 {array} models.Class
// @Router /classes [get]
func (r *postgresRepository) Classes(limit int, offset int, who int, whoKind string) ([]interface{}, error) {
	var query string
	var args []interface{}
	switch whoKind {
	case TeacherUser:
		query = "select id, year, section, info, grade " +
			" from back2school.classes join back2school.teaches on class = id " +
			" WHERE teacher = $1 " +
			" order by year desc, grade asc, section asc"
		args = append(args, who)
	case AdminUser:
		query = "select id, year, section, info, grade " +
			"from back2school.classes " +
			"order by year desc, grade asc, section asc "
	default:
		return nil, ErrorNotAuthorized
	}
	return r.listByParams(query, func(rows *sql.Rows) (interface{}, error) {
		class := models.Class{}
		err := rows.Scan(&class.ID, &class.Year, &class.Section, &class.Info, &class.Grade)
		return class, err
	}, limit, offset, args...)
}

// @Summary Get a student by class
// @Param id path int true "Class ID"
// @Param limit query int false "number of elements to return"
// @Param offset query int false "offset in the list of elements to return"
// @Tags Classes
// @Success 200 {array} models.Student
// @Router /classes/{id}/students [get]
func (r *postgresRepository) StudentsByClass(id int, limit int, offset int, who int, whoKind string) (students []interface{}, err error) {
	var query string
	var args []interface{}
	switch whoKind {
	case TeacherUser:
		query = "select distinct s.id, s.name, s.surname, s.mail, s.info " +
			"from back2school.students as s join back2school.enrolled as e " +
			" join back2school.teaches as t on s.id = e.student and t.class = e.class " +
			"where s.class = $1 and t.teacher = $2 " +
			"order by s.name desc, s.surname desc "
		args = append(args, id, who)
	case AdminUser:
		query = "select distinct id, name, surname, mail, info " +
			"from back2school.students join back2school.enrolled on student = id " +
			"where class = $1 " +
			"order by name desc, surname desc "
		args = append(args, id)
	default:
		return nil, ErrorNotAuthorized
	}
	return r.listByParams(query,
		func(rows *sql.Rows) (interface{}, error) {
			student := models.Student{}
			err := rows.Scan(&student.ID, &student.Name, &student.Surname, &student.Mail, &student.Info)
			return student, err
		}, limit, offset, args...)
}

//TODO add entrypoint !!!
// @Summary Get lectures by class
// @Param limit query int false "number of elements to return"
// @Param offset query int false "offset in the list of elements to return"
// @Param id path int true "Class ID"
// @Tags Classes
// @Success 200 {array} models.TimeTable
// @Router /classes/{id}/lectures [get]
func (r *postgresRepository) LectureByClass(id int, limit int, offset int, who int, whoKind string) ([]interface{}, error) {
	var query string
	var args []interface{}
	switch whoKind {
	case TeacherUser:
		query = "select id, class, subject, \"start\", \"end\", location, info " +
			"from back2school.timetable natural join back2school.teaches " +
			"where teacher = $1 and class = $2 " +
			"order by \"start\" desc"
		args = append(args, who, id)
	case AdminUser:
		query = "select id, class, subject, \"start\", \"end\", location, info " +
			"from back2school.timetable " +
			"order by \"start\" desc "
	default:
		return nil, ErrorNotAuthorized
	}
	return r.listByParams(
		query,
		func(rows *sql.Rows) (interface{}, error) {
			lecture := models.TimeTable{}
			err := rows.Scan(&lecture.ID, &lecture.Class, &lecture.Subject, &lecture.Start, &lecture.End, &lecture.Location, &lecture.Info)
			return lecture, err
		}, limit, offset, args...)
}

// List all notifications
// @Summary Get all notifications
// @Param limit query int false "number of elements to return"
// @Param offset query int false "offset in the list of elements to return"
// @Tags Notifications
// @Success 200 {array} models.Notification
// @Router /notifications [get]
func (r *postgresRepository) Notifications(limit int, offset int, who int, whoKind string) ([]interface{}, error) {
	var query string
	var args []interface{}
	switch whoKind {
	case TeacherUser, ParentUser:
		query = "select id, receiver, message, time, receiver_kind " +
			"from back2school.notification " +
			" where receiver = $1 and receiver_kind = $2 " +
			"order by time desc, receiver_kind desc "
		args = append(args, who, whoKind)
	case AdminUser:
		query = "select id, receiver, message, time, receiver_kind " +
			"from back2school.notification " +
			"order by time desc, receiver_kind desc "
	default:
		return nil, ErrorNotAuthorized
	}
	return r.listByParams(query,
		func(rows *sql.Rows) (interface{}, error) {
			notification := models.Notification{}
			err := rows.Scan(&notification.ID, &notification.Receiver, &notification.Message, &notification.Time, &notification.ReceiverKind)
			return notification, err
		}, limit, offset, args...)
}

// @Summary Get all grades
// @Param limit query int false "number of elements to return"
// @Param offset query int false "offset in the list of elements to return"
// @Tags Grades
// @Success 200 {array} models.Grade
// @Router /grades [get]
func (r *postgresRepository) Grades(limit int, offset int, who int, whoKind string) ([]interface{}, error) {
	var query string
	var args []interface{}
	switch whoKind {
	case ParentUser:
		query = "select id, student, grade, subject, date, teacher " +
			"from back2school.grades natural join back2school.isParent " +
			" where parent = $1" +
			"order by date desc, teacher asc"
		args = append(args, who)
	case AdminUser:
		query = "select id, student, grade, subject, date, teacher " +
			"from back2school.grades " +
			"order by date desc, teacher asc "
	default:
		return nil, ErrorNotAuthorized
	}
	return r.listByParams(query,
		func(rows *sql.Rows) (interface{}, error) {
			g := models.Grade{}
			err := rows.Scan(&g.ID, &g.Student.ID, &g.Grade, &g.Subject, &g.Date, &g.Teacher.ID)
			return g, err
		}, limit, offset, args...)
}

// @Summary Get all parents
// @Param limit query int false "number of elements to return"
// @Param offset query int false "offset in the list of elements to return"
// @Tags Parents
// @Success 200 {array} models.Parent
// @Router /parents [get]
func (r *postgresRepository) Parents(limit int, offset int, who int, whoKind string) ([]interface{}, error) {
	var query string
	var args []interface{}
	switch whoKind {
	case ParentUser:
		p, err := r.ParentByID(who, who, whoKind)
		return []interface{}{p}, err
		args = append(args, who)
	case AdminUser:
		query = "select id, name, surname, mail, info " +
			"from back2school.parents " +
			"order by name desc, surname desc "
	default:
		return nil, ErrorNotAuthorized
	}
	return r.listByParams(query,
		func(rows *sql.Rows) (interface{}, error) {
			parent := models.Parent{}
			err := rows.Scan(&parent.ID, &parent.Name, &parent.Surname, &parent.Mail, &parent.Info)
			return parent, err
		}, limit, offset, args...)
}

// see/modify the personal data of their registered children
// @Summary Get children of the parent
// @Param id path int true "Parent ID"
// @Param limit query int false "number of elements to return"
// @Param offset query int false "offset in the list of elements to return"
// @Tags Parents
// @Success 200 {array} models.Student
// @Router /parents/{id}/students [get]
func (r *postgresRepository) ChildrenByParent(id int, limit int, offset int, who int, whoKind string) (children []interface{}, err error) {
	var query string
	var args []interface{}
	switch whoKind {
	case ParentUser:
		if id == who {
			query = "SELECT distinct s.id, s.name, s.surname, s.mail, s.info " +
				"FROM back2school.isparent join back2school.students as s on student = s.id " +
				"WHERE parent = $1 " +
				"order by s.name desc "
			args = append(args, who)
		} else {
			return nil, ErrorNotAuthorized
		}
	case AdminUser:
		query = "SELECT distinct s.id, s.name, s.surname, s.mail, s.info " +
			"FROM back2school.isparent join back2school.students as s on student = s.id " +
			"WHERE parent = $1 " +
			"order by s.name desc"
		args = append(args, id)

	default:
		return nil, ErrorNotAuthorized
	}
	return r.listByParams(query,
		func(rows *sql.Rows) (interface{}, error) {
			student := models.Student{}
			err := rows.Scan(&student.ID, &student.Name, &student.Surname, &student.Mail, &student.Info)
			return student, err
		}, limit, offset, args...)
}

// see the monthly payments that have been made to the school in the past
// @Summary Get payments of the parent
// @Param id path int true "Parent ID"
// @Param limit query int false "number of elements to return"
// @Param offset query int false "offset in the list of elements to return"
// @Tags Parents
// @Success 200 {array} models.Payment
// @Router /parents/{id}/payments [get]
func (r *postgresRepository) PaymentsByParent(id int, limit int, offset int, who int, whoKind string) (payments []interface{}, err error) {
	var query string
	var args []interface{}
	switch whoKind {
	case ParentUser:
		if id == who {
			query = "select p.id, p.amount, p.student, p.payed, p.reason, p.emitted " +
				"from back2school.payments as p natural join back2school.isParent " +
				"where parent = $1 " +
				"order by p.emitted desc"
			args = append(args, who)
		} else {
			return nil, ErrorNotAuthorized
		}
	case AdminUser:
		query = "select p.id, p.amount, p.student, p.payed, p.reason, p.emitted " +
			"from back2school.payments as p natural join back2school.isParent " +
			"where parent = $1 " +
			"order by p.emitted desc"
		args = append(args, id)

	default:
		return nil, ErrorNotAuthorized
	}
	return r.listByParams(query, func(rows *sql.Rows) (interface{}, error) {
		payment := models.Payment{}
		err := rows.Scan(&payment.ID, &payment.Amount, &payment.Student.ID, &payment.Payed, &payment.Reason, &payment.Emitted)
		return payment, err
	}, limit, offset, args...)
}

// see general/personal notifications coming from the school
// @Summary Get notifications of the parent
// @Param id path int true "Parent ID"
// @Param limit query int false "number of elements to return"
// @Param offset query int false "offset in the list of elements to return"
// @Tags Parents
// @Success 200 {array} models.Notification
// @Router /parents/{id}/notifications [get]
func (r *postgresRepository) NotificationsByParent(id int, limit int, offset int, who int, whoKind string) (list []interface{}, err error) {
	var query string
	var args []interface{}
	switch whoKind {
	case ParentUser:
		if id == who {
			query = "select * from ( " +
				"select n.id, n.receiver, n.message, n.receiver_kind, n.time " +
				"from back2school.notification as n join back2school.isparent on n.receiver = student " +
				"where parent = $1 and receiver_kind = 'student' " +
				"union all  " +
				"select n.id, n.receiver, n.message, n.receiver_kind, n.time " +
				"from back2school.notification as n " +
				"where receiver_kind = 'general' " +
				") as a order by time desc "
			args = append(args, who)
		} else {
			return nil, ErrorNotAuthorized
		}
	case AdminUser:
		query = "select * from ( " +
			"select n.id, n.receiver, n.message, n.receiver_kind, n.time " +
			"from back2school.notification as n join back2school.isparent on n.receiver = student " +
			"where parent = $1 and receiver_kind = 'student' " +
			"union all  " +
			"select n.id, n.receiver, n.message, n.receiver_kind, n.time " +
			"from back2school.notification as n " +
			"where receiver_kind = 'general' " +
			") as a order by time desc "
		args = append(args, id)

	default:
		return nil, ErrorNotAuthorized
	}
	return r.listByParams(query,
		func(rows *sql.Rows) (interface{}, error) {
			notification := models.Notification{}
			err := rows.Scan(&notification.ID, &notification.Receiver, &notification.Message,
				&notification.ReceiverKind, &notification.Time)
			return notification, err
		}, limit, offset, args...)
}

// see/modify appointments that they have with their children's teachers
// (calendar-like support for requesting appointments, err error)
// @Summary Get appointments of the parent
// @Param id path int true "Parent ID"
// @Tags Parents
// @Param limit query int false "number of elements to return"
// @Param offset query int false "offset in the list of elements to return"
// @Tags Parents
// @Success 200 {array} models.Appointment
// @Router /parents/{id}/appointments [get]
func (r *postgresRepository) AppointmentsByParent(id int, limit int, offset int, who int, whoKind string) (appointments []interface{}, err error) {
	var query string
	var args []interface{}
	switch whoKind {
	case ParentUser:
		if id == who {
			query = "select a.id, a.student, a.teacher, a.location, a.time " +
				"from back2school.appointments as a natural join back2school.isparent  " +
				"where parent = $1 " +
				"order by a.time desc"
			args = append(args, who)
		} else {
			return nil, ErrorNotAuthorized
		}
	case AdminUser:
		query = "select a.id, a.student, a.teacher, a.location, a.time " +
			"from back2school.appointments as a natural join back2school.isparent " +
			"where parent = $1 " +
			"order by a.time desc "
		args = append(args, id)

	default:
		return nil, ErrorNotAuthorized
	}
	return r.listByParams(query,
		func(rows *sql.Rows) (interface{}, error) {
			appointment := models.Appointment{}
			err := rows.Scan(&appointment.ID, &appointment.Student.ID, &appointment.Teacher.ID, &appointment.Location, &appointment.Time)
			return appointment, err
		}, limit, offset, args...)
}

// Get payment by id
// @Summary Get a payment by id
// @Param id path int true "Payment ID"
// @Tags Payments
// @Success 200 {object} models.Payment
// @Router /payments/{id} [get]
func (r *postgresRepository) PaymentByID(id int, who int, whoKind string) (interface{}, error) {
	payment := &models.Payment{}
	var query string
	var args []interface{}
	switch whoKind {
	case ParentUser:
		query = "SELECT id, amount, payed, emitted, reason " +
			"FROM back2school.payments natural join back2school.isParent" +
			" WHERE id = $1 and parent = $2 "
		args = append(args, id, who)
	case AdminUser:
		query = "SELECT id, amount, payed, emitted, reason " +
			"FROM back2school.payments WHERE id = $1 "
		args = append(args, id)

	default:
		return nil, ErrorNotAuthorized
	}
	err := r.QueryRow(query, args...).Scan(payment.ID, payment.Amount, payment.Payed, payment.Emitted, payment.Reason)
	return switchResult(payment, err)
}

// @Summary Get all payments
// @Param limit query int false "number of elements to return"
// @Param offset query int false "offset in the list of elements to return"
// @Tags Payments
// @Success 200 {array} models.Payment
// @Router /payments [get]
func (r *postgresRepository) Payments(limit int, offset int, who int, whoKind string) ([]interface{}, error) {
	var query string
	var args []interface{}
	switch whoKind {
	case ParentUser:
		query = "select id, amount, student, payed, reason, emitted " +
			"from back2school.payments natural join back2school.isParent " +
			" where parent = $1 " +
			"order by payed asc, emitted asc "
		args = append(args, who)
	case AdminUser:
		query = "select id, amount, student, payed, reason, emitted " +
			"from back2school.payments " +
			"order by payed asc, emitted asc "
	default:
		return nil, ErrorNotAuthorized
	}
	return r.listByParams(query,
		func(rows *sql.Rows) (interface{}, error) {
			payment := models.Payment{}
			err := rows.Scan(&payment.ID, &payment.Amount, &payment.Student.ID, &payment.Payed, &payment.Reason, &payment.Emitted)
			return payment, err
		}, limit, offset, args...)
}

// @Summary Get all appointments
// @Param limit query int false "number of elements to return"
// @Param offset query int false "offset in the list of elements to return"
// @Tags Appointments
// @Router /appointments [get]
// @Success 200 {object} models.List
// @Security ApiKeyAuth
func (r *postgresRepository) Appointments(limit int, offset int, who int, whoKind string) ([]interface{}, error) {
	var query string
	var args []interface{}
	switch whoKind {
	case ParentUser:
		query = "select id, student, teacher, location, time " +
			"from back2school.appointments natural join back2school.isParent " +
			" where parent = $1 " +
			"order by time desc "
		args = append(args, who)
	case AdminUser:
		query = "select id, student, teacher, location, time " +
			"from back2school.appointments " +
			"order by time desc, teacher asc "
	default:
		return nil, ErrorNotAuthorized
	}
	return r.listByParams(query,
		func(rows *sql.Rows) (interface{}, error) {
			appointment := models.Appointment{}
			err := rows.Scan(&appointment.ID, &appointment.Student.ID, &appointment.Teacher.ID, &appointment.Location, &appointment.Time)
			return appointment, err
		}, limit, offset, args...)
}

// LectureByClass(id int, limit int, offset int) (students []interface{}, err error)
// TODO GradeStudent(grade models.Grade) error
// TODO
// parents:
// see/pay (fake payment) upcoming scheduled payments (monthly, material, trips, err error)
// admins:
// everything
// @Summary Get all students
// @Param limit query int false "number of elements to return"
// @Param offset query int false "offset in the list of elements to return"
// @Tags Students
// @Success 200 {array} models.Student
// @Router /students [get]
func (r *postgresRepository) Students(limit int, offset int, who int, whoKind string) (student []interface{}, err error) {
	var query string
	var args []interface{}
	switch whoKind {
	case ParentUser:
		query = "select id, name, surname, mail, info " +
			"from back2school.students join back2school.isParent on id=student " +
			"where parent = $1 " +
			"order by name desc, surname desc "
		args = append(args, who)
	case AdminUser:
		query = "select id, name, surname, mail, info  " +
			"from back2school.students " +
			"order by name desc, surname desc "
	default:
		return nil, ErrorNotAuthorized
	}
	return r.listByParams(query, func(rows *sql.Rows) (interface{}, error) {
		student := models.Student{}
		err := rows.Scan(&student.ID, &student.Name, &student.Surname, &student.Mail, &student.Info)
		return student, err
	}, limit, offset, args...)
}

// see the grades obtained by their children
// @Summary Get grades of the student
// @Param id path int true "Student ID"
// @Param limit query int false "number of elements to return"
// @Param offset query int false "offset in the list of elements to return"
// @Tags Students
// @Success 200 {array} models.Grade
// @Router /students/{id}/grades [get]
func (r *postgresRepository) GradesByStudent(id int, limit int, offset int, who int, whoKind string) ([]interface{}, error) {
	var query string
	var args []interface{}
	switch whoKind {
	case ParentUser:
		query = "SELECT id, student, subject, date, grade, teacher " +
			"FROM back2school.grades natural join back2school.isParent " +
			"WHERE student = $1 and parent = $2 " +
			"order by date desc "
		args = append(args, id, who)
		//TODO case ParentUser: //GradesByStudent
	case AdminUser:
		query = "SELECT id, student, subject, date, grade, teacher " +
			"FROM back2school.grades " +
			"WHERE student = $1 " +
			"order by date desc "
		args = append(args, id)
	default:
		return nil, ErrorNotAuthorized
	}
	return r.listByParams(query,
		func(rows *sql.Rows) (interface{}, error) {
			grade := models.Grade{}
			err := rows.Scan(&grade.ID, &grade.Student.ID, &grade.Subject, &grade.Date, &grade.Grade, &grade.Teacher.ID)
			return grade, err
		}, limit, offset, args...)
}

// see/modify their personal data
// Get teacher by id
// @Summary Get a teacher by id
// @Param id path int true "Teacher ID"
// @Tags Teachers
// @Success 200 {object} models.Teacher
// @Router /teachers/{id} [get]
func (r *postgresRepository) TeacherByID(id int, who int, whoKind string) (interface{}, error) {
	teacher := models.Teacher{}
	var query string
	var args []interface{}
	switch whoKind {
	case TeacherUser:
		if id == who {
			query = "SELECT id, name, surname, mail " +
				"FROM back2school.teachers " +
				"WHERE id = $1 "
			args = append(args, who)
		} else {
			return nil, ErrorNotAuthorized
		}
	case AdminUser:
		query = "SELECT id, name, surname, mail " +
			"FROM back2school.teachers " +
			"WHERE id = $1 "
		args = append(args, id)
	default:
		return nil, ErrorNotAuthorized
	}
	err := r.QueryRow(query,
		id).Scan(&teacher.ID, &teacher.Name, &teacher.Surname, &teacher.Mail)
	return switchResult(teacher, err)

}

// List all teachers
// @Summary Get all teachers
// @Param limit query int false "number of elements to return"
// @Param offset query int false "offset in the list of elements to return"
// @Tags Teachers
// @Success 200 {array} models.Teacher
// @Router /teachers [get]
func (r *postgresRepository) Teachers(limit int, offset int, who int, whoKind string) ([]interface{}, error) {
	var query string
	var args []interface{}
	switch whoKind {
	case TeacherUser:
		t, err := r.TeacherByID(who, who, whoKind)
		return []interface{}{t}, err
	case AdminUser:
		query = "select id, name, surname, mail, info  " +
			"from back2school.teachers " +
			"order by name desc, surname desc"
	default:
		return nil, ErrorNotAuthorized
	}
	return r.listByParams(query, func(rows *sql.Rows) (interface{}, error) {
		teacher := models.Teacher{}
		err := rows.Scan(&teacher.ID, &teacher.Name, &teacher.Surname, &teacher.Mail, &teacher.Info)
		return teacher, err
	}, limit, offset, args...)
}

// @Summary Get appointments of the teacher
// @Param id path int true "Teacher ID"
// @Param limit query int false "number of elements to return"
// @Param offset query int false "offset in the list of elements to return"
// @Tags Teachers
// @Success 200 {array} models.Appointment
// @Router /teachers/{id}/appointments [get]
func (r *postgresRepository) AppointmentsByTeacher(id int, limit int, offset int, who int, whoKind string) (appointments []interface{}, err error) {
	var query string
	var args []interface{}
	switch whoKind {
	case TeacherUser:
		if id == who {
			query = "SELECT id, student, teacher, location, time " +
				"FROM back2school.appointments " +
				"WHERE teacher = $1 " +
				"order by time desc "
			args = append(args, who)
		} else {
			return nil, ErrorNotAuthorized
		}
	case AdminUser:
		query = "SELECT id, student, teacher, location, time " +
			"FROM back2school.appointments " +
			"WHERE teacher = $1 " +
			"order by time desc"
		args = append(args, id)

	default:
		return nil, ErrorNotAuthorized
	}
	return r.listByParams(query,
		func(rows *sql.Rows) (interface{}, error) {
			app := models.Appointment{}
			err := rows.Scan(&app.ID, &app.Student.ID, &app.Teacher.ID, &app.Location, &app.Time)
			return app, err
		}, limit, offset, args...)
}

// @Summary Get notifications of the teacher
// @Param id path int true "Teacher ID"
// @Param limit query int false "number of elements to return"
// @Param offset query int false "offset in the list of elements to return"
// @Tags Teachers
// @Success 200 {array} models.TimeTable
// @Router /teachers/{id}/notifications [get]
func (r *postgresRepository) NotificationsByTeacher(id int, limit int, offset int, who int, whoKind string) (notifications []interface{}, err error) {
	var query string
	var args []interface{}
	switch whoKind {
	case TeacherUser:
		if id == who {
			query = "SELECT id, receiver, message, receiver_kind, time " +
				"FROM back2school.notification " +
				"WHERE (receiver = $1 and receiver_kind = $2) or receiver_kind = 'general' " +
				"order by time desc "
			args = append(args, who, whoKind)
		} else {
			return nil, ErrorNotAuthorized
		}
	case AdminUser:
		query = "SELECT id, receiver, message, receiver_kind, time " +
			"FROM back2school.notification " +
			"WHERE (receiver = $1 and receiver_kind = '" + TeacherUser + "') or receiver_kind = 'general' " +
			"order by time desc "
		args = append(args, id)

	default:
		return nil, ErrorNotAuthorized
	}
	return r.listByParams(query,
		func(rows *sql.Rows) (interface{}, error) {
			notif := models.Notification{}
			err := rows.Scan(&notif.ID, &notif.Receiver, &notif.Message, &notif.ReceiverKind, &notif.Time)
			return notif, err
		}, limit, offset, args...)
}

// @Summary Get subject taught by the teacher
// @Param id path int true "Teacher ID"
// @Param limit query int false "number of elements to return"
// @Param offset query int false "offset in the list of elements to return"
// @Success 200 {object} repository.Subjects
// @Tags Teachers
// @Router /teachers/{id}/subjects [get]
func (r *postgresRepository) SubjectsByTeacher(id int, limit int, offset int, who int, whoKind string) (notifications []interface{}, err error) {
	var query string
	var args []interface{}
	switch whoKind {
	case TeacherUser:
		if id == who {
			query = "SELECT DISTINCT subject FROM back2school.teaches where teacher = $1 order by subject "
			args = append(args, who)
		} else {
			return nil, ErrorNotAuthorized
		}
	case AdminUser:
		query = "SELECT DISTINCT subject FROM back2school.teaches where teacher = $1 order by subject "
		args = append(args, id)
	default:
		return nil, ErrorNotAuthorized
	}
	return r.listByParams(query,
		func(rows *sql.Rows) (interface{}, error) {
			subj := ""
			err := rows.Scan(&subj)
			return subj, err
		}, limit, offset, args...)
}

// @Summary Get classes in which the subject is taught by the teacher
// @Param id path int true "Teacher ID"
// @Param subject path int true "Subject ID"
// @Param limit query int false "number of elements to return"
// @Param offset query int false "offset in the list of elements to return"
// @Success 200 {array} models.Class
// @Tags Teachers
// @Router /teachers/{id}/subjects/{subject} [get]
func (r *postgresRepository) ClassesBySubjectAndTeacher(teacher int, subject string, limit int, offset int, who int, whoKind string) ([]interface{}, error) {
	var query string
	var args []interface{}
	switch whoKind {
	case TeacherUser:
		if teacher == who {
			query = "SELECT id, year, section, info, grade " +
				"FROM back2school.teaches join back2school.classes on id = class " +
				"WHERE teacher = $1 and subject = $2 " +
				"order by year desc, grade asc, section desc "
			args = append(args, who, subject)
		} else {
			return nil, ErrorNotAuthorized
		}
	case AdminUser:
		query = "SELECT id, year, section, info, grade " +
			"FROM back2school.teaches join back2school.classes on id = class " +
			"WHERE teacher = $1 and subject = $2 " +
			"order by year desc, grade asc, section desc "
		args = append(args, teacher, subject)
	default:
		return nil, ErrorNotAuthorized
	}
	return r.listByParams(query,
		func(rows *sql.Rows) (interface{}, error) {
			class := models.Class{}
			err := rows.Scan(&class.ID, &class.Year, &class.Section, &class.Info, &class.Grade)
			return class, err
		}, limit, offset, args...)
}

// @Summary Get lectures taught by the teacher
// @Param id path int true "Teacher ID"
// @Param limit query int false "number of elements to return"
// @Param offset query int false "offset in the list of elements to return"
// @Tags Teachers
// @Success 200 {array} models.TimeTable
// @Router /teachers/{id}/lectures [get]
func (r *postgresRepository) LecturesByTeacher(id int, limit int, offset int, who int, whoKind string) (lectures []interface{}, err error) {
	var query string
	var args []interface{}
	switch whoKind {
	case TeacherUser:
		if id == who {
			query = "SELECT id, class, subject, location, start, \"end\", info " +
				"from back2school.timetable natural join back2school.teaches as t " +
				"where t.teacher = $1 " +
				"order by start desc"
			args = append(args, who)
		} else {
			return nil, ErrorNotAuthorized
		}
	case AdminUser:
		query = "SELECT id, class, subject, location, start, \"end\", info " +
			"from back2school.timetable natural join back2school.teaches as t " +
			"where t.teacher = $1 " +
			"order by start desc"
		args = append(args, id)
	default:
		return nil, ErrorNotAuthorized
	}
	return r.listByParams(query,
		func(rows *sql.Rows) (interface{}, error) {
			lecture := models.TimeTable{}
			err := rows.Scan(&lecture.ID, &lecture.Class.ID, &lecture.Subject, &lecture.Location, &lecture.Start, &lecture.End, &lecture.Info)
			return lecture, err

		}, limit, offset, args...)
}

// @Summary Get classes in which the teacher teaches
// @Param id path int true "Teacher ID"
// @Param limit query int false "number of elements to return"
// @Param offset query int false "offset in the list of elements to return"
// @Tags Teachers
// @Success 200 {array} models.Class
// @Router /teachers/{id}/classes [get]
func (r *postgresRepository) ClassesByTeacher(id int, limit int, offset int, who int, whoKind string) ([]interface{}, error) {
	type Class struct {
		models.Class
		Subject models.Subject `json:"subject,omitempty"`
	}
	var query string
	var args []interface{}
	switch whoKind {
	case TeacherUser:
		if id == who {
			query = "SELECT id, year, section, info, grade, subject " +
				"FROM back2school.teaches join back2school.classes on id = class  " +
				"WHERE teacher = $1 " +
				"order by subject asc, year desc, grade asc, section desc "
			args = append(args, who)
		} else {
			return nil, ErrorNotAuthorized
		}
	case AdminUser:
		query = "SELECT id, year, section, info, grade, subject " +
			"FROM back2school.teaches join back2school.classes on id = class  " +
			"WHERE teacher = $1 " +
			"order by subject asc, year desc, grade asc, section desc "
		args = append(args, id)
	default:
		return nil, ErrorNotAuthorized
	}
	return r.listByParams(query,
		func(rows *sql.Rows) (interface{}, error) {
			class := Class{}
			err := rows.Scan(&class.ID, &class.Year, &class.Section, &class.Info, &class.Grade, &class.Subject)
			return class, err
		}, limit, offset, args...)
}

// @Summary Get all lectures
// @Param limit query int false "number of elements to return"
// @Param offset query int false "offset in the list of elements to return"
// @Tags Lectures
// @Success 200 {array} models.TimeTable
// @Router /lectures [get]
func (r *postgresRepository) Lectures(limit int, offset int, who int, whoKind string) ([]interface{}, error) {
	var query string
	var args []interface{}
	switch whoKind {
	case ParentUser:
		query = "select id, class, subject, \"start\", \"end\", location, info " +
			"from back2school.timetable natural join back2school.enrolled natural join back2school.isParent " +
			"where parent = $1 " +
			"order by \"start\" desc "
		args = append(args, who)
	case TeacherUser:
		query = "select id, class, subject, \"start\", \"end\", location, info " +
			"from back2school.timetable natural join back2school.teaches " +
			"where teacher = $1 " +
			"order by \"start\" desc "
		args = append(args, who)
	case AdminUser:
		query = "select id, class, subject, \"start\", \"end\", location, info " +
			"from back2school.timetable " +
			"order by \"start\" desc "
	default:
		return nil, ErrorNotAuthorized
	}
	return r.listByParams(query, func(rows *sql.Rows) (interface{}, error) {
		lecture := models.TimeTable{}
		err := rows.Scan(&lecture.ID, &lecture.Class, &lecture.Subject, &lecture.Start, &lecture.End, &lecture.Location, &lecture.Info)
		return lecture, err
	}, limit, offset, args...)
}
