package repository

import (
	"database/sql"

	_ "github.com/middleware2018-PSS/Services/docs"
	"github.com/middleware2018-PSS/Services/models"
)

type Subjects struct {
	Subjects []string `json:"subjects" example:"science"`
}

func (r *postgresRepository) ClassesForTeachers(limit int, offset int, who int) ([]interface{}, error) {
	query := "select id, year, section, info, grade " +
		" from back2school.classes join back2school.teaches on class = id " +
		" WHERE teacher = $1 " +
		" order by year desc, grade asc, section asc"
	return r.listByParams(query, func(rows *sql.Rows) (interface{}, error) {
		class := models.Class{}
		err := rows.Scan(&class.ID, &class.Year, &class.Section, &class.Info, &class.Grade)
		return class, err
	}, limit, offset, who)
}

func (r *postgresRepository) ClassesForAdmins(limit int, offset int) ([]interface{}, error) {
	query := "select id, year, section, info, grade " +
		"from back2school.classes " +
		"order by year desc, grade asc, section asc "
	return r.listByParams(query, func(rows *sql.Rows) (interface{}, error) {
		class := models.Class{}
		err := rows.Scan(&class.ID, &class.Year, &class.Section, &class.Info, &class.Grade)
		return class, err
	}, limit, offset)
}

func (r *postgresRepository) StudentsByClassForTeachers(id int, limit int, offset int, who int) (students []interface{}, err error) {
	query := "select distinct s.id, s.name, s.surname, s.mail, s.info " +
		"from back2school.students as s join back2school.enrolled as e " +
		" join back2school.teaches as t on s.id = e.student and t.class = e.class " +
		"where s.class = $1 and t.teacher = $2 " +
		"order by s.name desc, s.surname desc "
	return r.listByParams(query,
		func(rows *sql.Rows) (interface{}, error) {
			student := models.Student{}
			err := rows.Scan(&student.ID, &student.Name, &student.Surname, &student.Mail, &student.Info)
			return student, err
		}, limit, offset, id, who)
}
func (r *postgresRepository) StudentsByClassForAdmins(id int, limit int, offset int) (students []interface{}, err error) {
	query := "select distinct id, name, surname, mail, info " +
		"from back2school.students join back2school.enrolled on student = id " +
		"where class = $1 " +
		"order by name desc, surname desc "
	return r.listByParams(query,
		func(rows *sql.Rows) (interface{}, error) {
			student := models.Student{}
			err := rows.Scan(&student.ID, &student.Name, &student.Surname, &student.Mail, &student.Info)
			return student, err
		}, limit, offset, id)
}

func (r *postgresRepository) LectureByClassForTeacherOrParents(id int, limit int, offset int, who int) ([]interface{}, error) {
	query := "select id, class, subject, \"start\", \"end\", location, info " +
		"from back2school.timetable natural join back2school.teaches " +
		"where teacher = $1 and class = $2 " +
		"order by \"start\" desc"
	return r.listByParams(
		query,
		func(rows *sql.Rows) (interface{}, error) {
			lecture := models.TimeTable{}
			err := rows.Scan(&lecture.ID, lecture.Class, &lecture.Subject, &lecture.Start, &lecture.End, &lecture.Location, &lecture.Info)
			return lecture, err
		}, limit, offset, who, id)
}
func (r *postgresRepository) LectureByClassForAdmins(id int, limit int, offset int) ([]interface{}, error) {
	query := "select id, class, subject, \"start\", \"end\", location, info " +
		"from back2school.timetable " +
		"order by \"start\" desc "
	return r.listByParams(
		query,
		func(rows *sql.Rows) (interface{}, error) {
			lecture := models.TimeTable{}
			err := rows.Scan(&lecture.ID, lecture.Class, &lecture.Subject, &lecture.Start, &lecture.End, &lecture.Location, &lecture.Info)
			return lecture, err
		}, limit, offset)
}

func (r *postgresRepository) NotificationsForTeacherOrParents(limit int, offset int, who int, whoKind string) ([]interface{}, error) {
	query := "select id, receiver, message, time, receiver_kind " +
		"from back2school.notification " +
		" where receiver = $1 and receiver_kind = $2 " +
		"order by time desc, receiver_kind desc "
	return r.listByParams(query,
		func(rows *sql.Rows) (interface{}, error) {
			notification := models.Notification{}
			err := rows.Scan(&notification.ID, &notification.Receiver, &notification.Message, &notification.Time, &notification.ReceiverKind)
			return notification, err
		}, limit, offset, who, whoKind)
}

func (r *postgresRepository) NotificationsForAdmins(limit int, offset int) ([]interface{}, error) {
	query := "select id, receiver, message, time, receiver_kind " +
		"from back2school.notification " +
		"order by time desc, receiver_kind desc "
	return r.listByParams(query,
		func(rows *sql.Rows) (interface{}, error) {
			notification := models.Notification{}
			err := rows.Scan(&notification.ID, &notification.Receiver, &notification.Message, &notification.Time, &notification.ReceiverKind)
			return notification, err
		}, limit, offset)
}

func (r *postgresRepository) GradesForParent(limit int, offset int, who int) ([]interface{}, error) {
	query := "select id, student, grade, subject, date, teacher " +
		"from back2school.grades natural join back2school.isParent " +
		" where parent = $1" +
		"order by date desc, teacher asc"
	return r.listByParams(query,
		func(rows *sql.Rows) (interface{}, error) {
			g := models.Grade{}
			err := rows.Scan(&g.ID, g.Student, &g.Grade, &g.Subject, &g.Date, g.Teacher)
			return g, err
		}, limit, offset, who)
}
func (r *postgresRepository) GradesForAdmins(limit int, offset int) ([]interface{}, error) {
	query := "select id, student, grade, subject, date, teacher " +
		"from back2school.grades " +
		"order by date desc, teacher asc "
	return r.listByParams(query,
		func(rows *sql.Rows) (interface{}, error) {
			g := models.Grade{}
			err := rows.Scan(&g.ID, g.Student, &g.Grade, &g.Subject, &g.Date, g.Teacher)
			return g, err
		}, limit, offset)
}
func (r *postgresRepository) ParentsForParents(who int) ([]interface{}, error) {
	p, err := r.ParentByID(who)
	return []interface{}{p}, err
}
func (r *postgresRepository) ParentsForAdmins(limit int, offset int) ([]interface{}, error) {
	query := "select id, name, surname, mail, info " +
		"from back2school.parents " +
		"order by name desc, surname desc "
	return r.listByParams(query,
		func(rows *sql.Rows) (interface{}, error) {
			parent := models.Parent{}
			err := rows.Scan(&parent.ID, &parent.Name, &parent.Surname, &parent.Mail, &parent.Info)
			return parent, err
		}, limit, offset)
}
func (r *postgresRepository) ChildrenByParentForParent(id int, limit int, offset int) (children []interface{}, err error) {

	query := "SELECT distinct s.id, s.name, s.surname, s.mail, s.info " +
		"FROM back2school.isparent join back2school.students as s on student = s.id " +
		"WHERE parent = $1 " +
		"order by s.name desc "

	return r.listByParams(query,
		func(rows *sql.Rows) (interface{}, error) {
			student := models.Student{}
			err := rows.Scan(&student.ID, &student.Name, &student.Surname, &student.Mail, &student.Info)
			return student, err
		}, limit, offset, who)
}

func (r *postgresRepository) ChildrenByParentForAdmin(id int, limit int, offset int) (children []interface{}, err error) {
	query := "SELECT distinct s.id, s.name, s.surname, s.mail, s.info " +
		"FROM back2school.isparent join back2school.students as s on student = s.id " +
		"WHERE parent = $1 " +
		"order by s.name desc"

	return r.listByParams(query,
		func(rows *sql.Rows) (interface{}, error) {
			student := models.Student{}
			err := rows.Scan(&student.ID, &student.Name, &student.Surname, &student.Mail, &student.Info)
			return student, err
		}, limit, offset, id)
}

func (r *postgresRepository) PaymentsByParentForParent(id int, limit int, offset int) (payments []interface{}, err error) {
	query := "select p.id, p.amount, p.student, p.paid, p.reason, p.emitted " +
		"from back2school.payments as p natural join back2school.isParent " +
		"where parent = $1 " +
		"order by p.emitted desc"
	return r.listByParams(
		query,
		func(rows *sql.Rows) (interface{}, error) {
			lecture := models.TimeTable{}
			err := rows.Scan(&lecture.ID, lecture.Class, &lecture.Subject, &lecture.Start, &lecture.End, &lecture.Location, &lecture.Info)
			return lecture, err
		}, limit, offset, id)
}

func (r *postgresRepository) PaymentsByParentForAdmin(id int, limit int, offset int) (payments []interface{}, err error) {
	query := "select p.id, p.amount, p.student, p.paid, p.reason, p.emitted " +
		"from back2school.payments as p natural join back2school.isParent " +
		"where parent = $1 " +
		"order by p.emitted desc"

	return r.listByParams(query, func(rows *sql.Rows) (interface{}, error) {
		payment := models.Payment{}
		err := rows.Scan(&payment.ID, &payment.Amount, payment.Student, &payment.Paid, &payment.Reason, &payment.Emitted)
		return payment, err
	}, limit, offset, id)
}

func (r *postgresRepository) NotificationsByParentForParent(id int, limit int, offset int) (list []interface{}, err error) {

	query := "select * from ( " +
		"select n.id, n.receiver, n.message, n.receiver_kind, n.time " +
		"from back2school.notification as n join back2school.isparent on n.receiver = student " +
		"where parent = $1 and receiver_kind = 'student' " +
		"union all  " +
		"select n.id, n.receiver, n.message, n.receiver_kind, n.time " +
		"from back2school.notification as n " +
		"where receiver_kind = 'general' " +
		") as a order by time desc "

	return r.listByParams(query,
		func(rows *sql.Rows) (interface{}, error) {
			notification := models.Notification{}
			err := rows.Scan(&notification.ID, &notification.Receiver, &notification.Message,
				&notification.ReceiverKind, &notification.Time)
			return notification, err
		}, limit, offset, who)
}
func (r *postgresRepository) NotificationsByParentForAdmins(id int, limit int, offset int) (list []interface{}, err error) {

	query := "select * from ( " +
		"select n.id, n.receiver, n.message, n.receiver_kind, n.time " +
		"from back2school.notification as n join back2school.isparent on n.receiver = student " +
		"where parent = $1 and receiver_kind = 'student' " +
		"union all  " +
		"select n.id, n.receiver, n.message, n.receiver_kind, n.time " +
		"from back2school.notification as n " +
		"where receiver_kind = 'general' " +
		") as a order by time desc "

	return r.listByParams(query,
		func(rows *sql.Rows) (interface{}, error) {
			notification := models.Notification{}
			err := rows.Scan(&notification.ID, &notification.Receiver, &notification.Message,
				&notification.ReceiverKind, &notification.Time)
			return notification, err
		}, limit, offset, id)
}

func (r *postgresRepository) AppointmentsByParentForParent(id int, limit int, offset int) (appointments []interface{}, err error) {

	query := "select a.id, a.student, a.teacher, a.location, a.time " +
		"from back2school.appointments as a natural join back2school.isparent  " +
		"where parent = $1 " +
		"order by a.time desc"

	return r.listByParams(query,
		func(rows *sql.Rows) (interface{}, error) {
			appointment := models.Appointment{}
			err := rows.Scan(&appointment.ID, appointment.Student, appointment.Teacher, &appointment.Location, &appointment.Time)
			return appointment, err
		}, limit, offset, who)
}

func (r *postgresRepository) AppointmentsByParentForAdmin(id int, limit int, offset int) (appointments []interface{}, err error) {

	query := "select a.id, a.student, a.teacher, a.location, a.time " +
		"from back2school.appointments as a natural join back2school.isparent " +
		"where parent = $1 " +
		"order by a.time desc "
	return r.listByParams(query,
		func(rows *sql.Rows) (interface{}, error) {
			appointment := models.Appointment{}
			err := rows.Scan(&appointment.ID, appointment.Student, appointment.Teacher, &appointment.Location, &appointment.Time)
			return appointment, err
		}, limit, offset, id)
}

func (r *postgresRepository) PaymentByIDForParent(id int, who int) (interface{}, error) {
	payment := &models.Payment{}
	query := "SELECT id, amount, paid, emitted, reason " +
		"FROM back2school.payments natural join back2school.isParent" +
		" WHERE id = $1 and parent = $2 "

	err := r.QueryRow(query, id, who).Scan(payment.ID, payment.Amount, payment.Paid, payment.Emitted, payment.Reason)
	return switchResult(payment, err)
}

func (r *postgresRepository) PaymentByIDForAdmins(id int) (interface{}, error) {
	payment := &models.Payment{}
	query := "SELECT id, amount, paid, emitted, reason " +
		"FROM back2school.payments WHERE id = $1 "

	err := r.QueryRow(query, id).Scan(payment.ID, payment.Amount, payment.Paid, payment.Emitted, payment.Reason)
	return switchResult(payment, err)
}

func (r *postgresRepository) PaymentsForParent(limit int, offset int, who int) ([]interface{}, error) {

	query := "select id, amount, student, paid, reason, emitted " +
		"from back2school.payments natural join back2school.isParent " +
		" where parent = $1 " +
		"order by paid asc, emitted asc "
	return r.listByParams(query,
		func(rows *sql.Rows) (interface{}, error) {
			payment := models.Payment{}
			err := rows.Scan(&payment.ID, &payment.Amount, payment.Student, &payment.Paid, &payment.Reason, &payment.Emitted)
			return payment, err
		}, limit, offset, who)
}
func (r *postgresRepository) PaymentsForAdmin(limit int, offset int) ([]interface{}, error) {

	query := "select id, amount, student, paid, reason, emitted " +
		"from back2school.payments " +
		"order by paid asc, emitted asc "

	return r.listByParams(query,
		func(rows *sql.Rows) (interface{}, error) {
			payment := models.Payment{}
			err := rows.Scan(&payment.ID, &payment.Amount, payment.Student, &payment.Paid, &payment.Reason, &payment.Emitted)
			return payment, err
		}, limit, offset)
}

func (r *postgresRepository) AppointmentsForParents(limit int, offset int, who int) ([]interface{}, error) {

	query := "select id, student, teacher, location, time " +
		"from back2school.appointments natural join back2school.isParent " +
		" where parent = $1 " +
		"order by time desc "

	return r.listByParams(query,
		func(rows *sql.Rows) (interface{}, error) {
			appointment := models.Appointment{}
			err := rows.Scan(&appointment.ID, &appointment.Student, &appointment.Teacher, &appointment.Location, &appointment.Time)
			return appointment, err
		}, limit, offset, who)
}

func (r *postgresRepository) AppointmentsForAdmin(limit int, offset int) ([]interface{}, error) {

	query := "select id, student, teacher, location, time " +
		"from back2school.appointments " +
		"order by time desc, teacher asc "

	return r.listByParams(query,
		func(rows *sql.Rows) (interface{}, error) {
			appointment := models.Appointment{}
			err := rows.Scan(&appointment.ID, &appointment.Student, &appointment.Teacher, &appointment.Location, &appointment.Time)
			return appointment, err
		}, limit, offset)
}

func (r *postgresRepository) StudentsForParent(limit int, offset int, who int) (student []interface{}, err error) {

	query := "select id, name, surname, mail, info " +
		"from back2school.students join back2school.isParent on id=student " +
		"where parent = $1 " +
		"order by name desc, surname desc "

	return r.listByParams(query, func(rows *sql.Rows) (interface{}, error) {
		student := models.Student{}
		err := rows.Scan(&student.ID, &student.Name, &student.Surname, &student.Mail, &student.Info)
		return student, err
	}, limit, offset, who)
}
func (r *postgresRepository) StudentsForAdmins(limit int, offset int) (student []interface{}, err error) {

	query := "select id, name, surname, mail, info  " +
		"from back2school.students " +
		"order by name desc, surname desc "

	return r.listByParams(query, func(rows *sql.Rows) (interface{}, error) {
		student := models.Student{}
		err := rows.Scan(&student.ID, &student.Name, &student.Surname, &student.Mail, &student.Info)
		return student, err
	}, limit, offset)
}

func (r *postgresRepository) GradesByStudentForParent(id int, limit int, offset int, who int) ([]interface{}, error) {

	query := "SELECT id, student, subject, date, grade, teacher " +
		"FROM back2school.grades natural join back2school.isParent " +
		"WHERE student = $1 and parent = $2 " +
		"order by date desc "
	return r.listByParams(query,
		func(rows *sql.Rows) (interface{}, error) {
			grade := models.Grade{}
			err := rows.Scan(&grade.ID, grade.Student, &grade.Subject, &grade.Date, &grade.Grade, grade.Teacher)
			return grade, err
		}, limit, offset, id, who)
}
func (r *postgresRepository) GradesByStudentForAdmins(id int, limit int, offset int) ([]interface{}, error) {

	query := "SELECT id, student, subject, date, grade, teacher " +
		"FROM back2school.grades " +
		"WHERE student = $1 " +
		"order by date desc "

	return r.listByParams(query,
		func(rows *sql.Rows) (interface{}, error) {
			grade := models.Grade{}
			err := rows.Scan(&grade.ID, grade.Student, &grade.Subject, &grade.Date, &grade.Grade, grade.Teacher)
			return grade, err
		}, limit, offset, id)
}

func (r *postgresRepository) TeacherByIDForTeacher(id int) (interface{}, error) {
	teacher := models.Teacher{}
	query := "SELECT id, name, surname, mail " +
		"FROM back2school.teachers " +
		"WHERE id = $1 "

	err := r.QueryRow(query, id).Scan(&teacher.ID, &teacher.Name, &teacher.Surname, &teacher.Mail)
	return switchResult(teacher, err)

}

func (r *postgresRepository) TeacherByIDForAdmin(id int) (interface{}, error) {
	teacher := models.Teacher{}
	query := "SELECT id, name, surname, mail " +
		"FROM back2school.teachers " +
		"WHERE id = $1 "
	err := r.QueryRow(query, id).Scan(&teacher.ID, &teacher.Name, &teacher.Surname, &teacher.Mail)
	return switchResult(teacher, err)

}

func (r *postgresRepository) TeachersForTeacher(who int) ([]interface{}, error) {
	t, err := r.TeacherByIDForTeacher(who)
	return []interface{}{t}, err
}
func (r *postgresRepository) TeachersForAdmin(limit int, offset int) ([]interface{}, error) {

	query := "select id, name, surname, mail, info  " +
		"from back2school.teachers " +
		"order by name desc, surname desc"

	return r.listByParams(query, func(rows *sql.Rows) (interface{}, error) {
		teacher := models.Teacher{}
		err := rows.Scan(&teacher.ID, &teacher.Name, &teacher.Surname, &teacher.Mail, &teacher.Info)
		return teacher, err
	}, limit, offset)
}

func (r *postgresRepository) AppointmentsByTeacherForTeacher(id int, limit int, offset int, who int) (appointments []interface{}, err error) {
	query := "SELECT id, student, teacher, location, time " +
		"FROM back2school.appointments " +
		"WHERE teacher = $1 " +
		"order by time desc "

	return r.listByParams(query,
		func(rows *sql.Rows) (interface{}, error) {
			app := models.Appointment{}
			err := rows.Scan(&app.ID, app.Student, app.Teacher, &app.Location, &app.Time)
			return app, err
		}, limit, offset, who)
}

func (r *postgresRepository) AppointmentsByTeacherForAdmin(id int, limit int, offset int) (appointments []interface{}, err error) {

	query := "SELECT id, student, teacher, location, time " +
		"FROM back2school.appointments " +
		"WHERE teacher = $1 " +
		"order by time desc"

	return r.listByParams(query,
		func(rows *sql.Rows) (interface{}, error) {
			app := models.Appointment{}
			err := rows.Scan(&app.ID, app.Student, app.Teacher, &app.Location, &app.Time)
			return app, err
		}, limit, offset, id)
}

func (r *postgresRepository) NotificationsByTeacherForTeacher(id int, limit int, offset int, who int, whoKind string) (notifications []interface{}, err error) {

	query := "SELECT id, receiver, message, receiver_kind, time " +
		"FROM back2school.notification " +
		"WHERE (receiver = $1 and receiver_kind = $2) or receiver_kind = 'general' " +
		"order by time desc "

	return r.listByParams(query,
		func(rows *sql.Rows) (interface{}, error) {
			notif := models.Notification{}
			err := rows.Scan(&notif.ID, &notif.Receiver, &notif.Message, &notif.ReceiverKind, &notif.Time)
			return notif, err
		}, limit, offset, who, whoKind)
}

func (r *postgresRepository) NotificationsByTeacherForAdmin(id int, limit int, offset int) (notifications []interface{}, err error) {

	query := "SELECT id, receiver, message, receiver_kind, time " +
		"FROM back2school.notification " +
		"WHERE (receiver = $1 and receiver_kind = '" + TeacherUser + "') or receiver_kind = 'general' " +
		"order by time desc "

	return r.listByParams(query,
		func(rows *sql.Rows) (interface{}, error) {
			notif := models.Notification{}
			err := rows.Scan(&notif.ID, &notif.Receiver, &notif.Message, &notif.ReceiverKind, &notif.Time)
			return notif, err
		}, limit, offset, id)
}

func (r *postgresRepository) SubjectsByTeacher(id int, limit int, offset int) (notifications []interface{}, err error) {

	query := "SELECT DISTINCT subject FROM back2school.teaches where teacher = $1 order by subject "

	return r.listByParams(query,
		func(rows *sql.Rows) (interface{}, error) {
			subj := ""
			err := rows.Scan(&subj)
			return subj, err
		}, limit, offset, id)
}

func (r *postgresRepository) ClassesBySubjectAndTeacher(teacher int, subject string, limit int, offset int) ([]interface{}, error) {

	query := "SELECT id, year, section, info, grade " +
		"FROM back2school.teaches join back2school.classes on id = class " +
		"WHERE teacher = $1 and subject = $2 " +
		"order by year desc, grade asc, section desc "

	return r.listByParams(query,
		func(rows *sql.Rows) (interface{}, error) {
			class := models.Class{}
			err := rows.Scan(&class.ID, &class.Year, &class.Section, &class.Info, &class.Grade)
			return class, err
		}, limit, offset, teacher, subject)
}

func (r *postgresRepository) LecturesByTeacher(id int, limit int, offset int) (lectures []interface{}, err error) {

	query := "SELECT id, class, subject, location, start, \"end\", info " +
		"from back2school.timetable natural join back2school.teaches as t " +
		"where t.teacher = $1 " +
		"order by start desc"

	return r.listByParams(query,
		func(rows *sql.Rows) (interface{}, error) {
			lecture := models.TimeTable{}
			err := rows.Scan(&lecture.ID, lecture.Class, &lecture.Subject, &lecture.Location, &lecture.Start, &lecture.End, &lecture.Info)
			return lecture, err

		}, limit, offset, who)
}

func (r *postgresRepository) ClassesByTeacher(id int, limit int, offset int) ([]interface{}, error) {
	type Class struct {
		models.Class
		Subject models.Subject `json:"subject,omitempty"`
	}
	query := "SELECT id, year, section, info, grade, subject " +
		"FROM back2school.teaches join back2school.classes on id = class  " +
		"WHERE teacher = $1 " +
		"order by subject asc, year desc, grade asc, section desc "

	return r.listByParams(query,
		func(rows *sql.Rows) (interface{}, error) {
			class := Class{}
			err := rows.Scan(&class.ID, &class.Year, &class.Section, &class.Info, &class.Grade, &class.Subject)
			return class, err
		}, limit, offset, who)
}

func (r *postgresRepository) LecturesForParent(limit int, offset int, who int) ([]interface{}, error) {

	query := "select id, class, subject, \"start\", \"end\", location, info " +
		"from back2school.timetable natural join back2school.enrolled natural join back2school.isParent " +
		"where parent = $1 " +
		"order by \"start\" desc "

	return r.listByParams(query, func(rows *sql.Rows) (interface{}, error) {
		lecture := models.TimeTable{}
		err := rows.Scan(&lecture.ID, lecture.Class, &lecture.Subject, &lecture.Start, &lecture.End, &lecture.Location, &lecture.Info)
		return lecture, err
	}, limit, offset, who)
}

func (r *postgresRepository) LecturesForAdmin(limit int, offset int) ([]interface{}, error) {
	query := "select id, class, subject, \"start\", \"end\", location, info " +
		"from back2school.timetable " +
		"order by \"start\" desc "

	return r.listByParams(query, func(rows *sql.Rows) (interface{}, error) {
		lecture := models.TimeTable{}
		err := rows.Scan(&lecture.ID, lecture.Class, &lecture.Subject, &lecture.Start, &lecture.End, &lecture.Location, &lecture.Info)
		return lecture, err
	}, limit, offset)
}

func (r *postgresRepository) AccountsForAdmins(limit int, offset int) ([]interface{}, error) {
	query := `select username, kind, id from back2school.accounts`
	return r.listByParams(query, func(rows *sql.Rows) (interface{}, error) {
		account := models.Account{}
		err := rows.Scan(&account.Username, &account.Kind, &account.ID)
		return account, err
	}, limit, offset)

}
