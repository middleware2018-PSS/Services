package repository

import (
	"database/sql"
	"github.com/middleware2018-PSS/Services/src/models"
)

func (r *postgresRepository) CheckUser(userID string, password string) (string, bool) {
	query := `select "user", password from back2school.accounts where "user" = $1 and password = $2`
	var id, pass string
	err := r.QueryRow(query, userID, password).Scan(&id, &pass)
	return userID, err == nil
}

func (r *postgresRepository) UserKind(userID string) map[string]interface{} {
	query := `select kind, id from back2school.accounts where "user" = $1`
	var kind string
	var id int
	r.QueryRow(query, userID).Scan(&kind, &id)
	return map[string]interface{}{"kind": kind, "dbID": id}
}

func (r *postgresRepository) AppointmentByID(id int64) (interface{}, error) {
	appointment := models.Appointment{}
	err := r.QueryRow("SELECT id, student, teacher, time, location "+
		"FROM back2school.appointments WHERE id = $1 ", id).Scan(
		&appointment.ID, &appointment.Student.ID, &appointment.Teacher.ID, &appointment.Time, &appointment.Location)
	return switchResult(appointment, err)
}

func (r *postgresRepository) GradeByID(id int64) (interface{}, error) {
	grade := models.Grade{}
	err := r.QueryRow("SELECT id, student, teacher, subject, date, grade "+
		"FROM back2school.grades WHERE id = $1 ", id).Scan(
			&grade.ID, &grade.Student.ID, &grade.Teacher.ID, &grade.Subject, &grade.Date, &grade.Grade)
	return switchResult(grade, err)
}

func (r *postgresRepository) ClassByID(id int64) (interface{}, error) {
	class := models.Class{}
	err := r.QueryRow("SELECT id, year, section, info, grade FROM back2school.classes "+
		"WHERE id = $1", id).Scan(&class.ID, &class.Year, &class.Section, &class.Info, &class.Grade)
	return switchResult(class, err)
}

func (r *postgresRepository) Classes(limit int, offset int) ([]interface{}, error) {
	return r.listByParams("select id, year, section, info, grade "+
		"from back2school.classes "+
		"order by year desc, grade asc, section asc", func(rows *sql.Rows) (interface{}, error) {
		class := models.Class{}
		err := rows.Scan(&class.ID, &class.Year, &class.Section, &class.Info, &class.Grade)
		return class, err
	}, limit, offset)
}

func (r *postgresRepository) StudentsByClass(id int64, limit int, offset int) (students []interface{}, err error) {
	return r.listByParams("select distinct id, name, surname, mail, info "+
		"from back2school.students join back2school.enrolled on student = id "+
		"where class = $1 "+
		"order by name desc, surname desc",
		func(rows *sql.Rows) (interface{}, error) {
			student := models.Student{}
			err := rows.Scan(&student.ID, &student.Name, &student.Surname, &student.Mail, &student.Info)
			return student, err
		}, limit, offset, id)
}

func (r *postgresRepository) LectureByClass(id int64, limit int, offset int) ([]interface{}, error) {
	return r.listByParams(
		"select id, class, subject, \"start\", \"end\", location, info "+
			"from back2school.timetable natural join back2school.teaches "+
			"where teacher = $1 "+
			"order by \"start\" desc",
		func(rows *sql.Rows) (interface{}, error) {
			lecture := models.TimeTable{}
			err := rows.Scan(&lecture.ID, &lecture.Class, &lecture.Subject, &lecture.Start, &lecture.End, &lecture.Location, &lecture.Info)
			return lecture, err
		}, limit, offset, id)
}

func (r *postgresRepository) NotificationByID(id int64) (interface{}, error) {
	n := models.Notification{}
	err := r.QueryRow("SELECT id, receiver, message, time, receiver_kind "+
		"FROM back2school.notification WHERE id = $1 ", id).Scan(&n.ID,
		&n.Receiver, &n.Message, &n.Time, &n.ReceiverKind)
	return switchResult(n, err)
}

func (r *postgresRepository) Notifications(limit int, offset int) ([]interface{}, error) {
	return r.listByParams("select id, receiver, message, time, receiver_kind "+
		"from back2school.notification "+
		"order by time desc, receiver_kind desc",
		func(rows *sql.Rows) (interface{}, error) {
			notification := models.Notification{}
			err := rows.Scan(&notification.ID, &notification.Receiver, &notification.Message, &notification.Time, &notification.ReceiverKind)
			return notification, err
		}, limit, offset)
}

func (r *postgresRepository) ParentByID(id int64) (interface{}, error) {
	p := models.Parent{}
	err := r.QueryRow("SELECT id,	name, surname, mail, info "+
		"FROM back2school.parents WHERE id = $1",
		id).Scan(&p.ID, &p.Name, &p.Surname, &p.Mail, &p.Info)
	return switchResult(p, err)
}

func (r *postgresRepository) Grades(limit int, offset int) ([]interface{}, error) {
	return r.listByParams("select id, student, grade, subject, date, teacher "+
		"from back2school.grades "+
		"order by date desc, teacher asc",
		func(rows *sql.Rows) (interface{}, error) {
			g := models.Grade{}
			err := rows.Scan(&g.ID, &g.Student.ID, &g.Grade, &g.Subject, &g.Date, &g.Teacher)
			return g, err
		}, limit, offset)
}

func (r *postgresRepository) Parents(limit int, offset int) ([]interface{}, error) {
	return r.listByParams("select id, name, surname, mail, info "+
		"from back2school.parents "+
		"order by name desc, surname desc",
		func(rows *sql.Rows) (interface{}, error) {
			parent := models.Parent{}
			err := rows.Scan(&parent.ID, &parent.Name, &parent.Surname, &parent.Mail, &parent.Info)
			return parent, err
		}, limit, offset)
}

func (r *postgresRepository) ChildrenByParent(id int64, limit int, offset int) (children []interface{}, err error) {
	return r.listByParams("SELECT distinct s.id, s.name, s.surname, s.mail, s.info "+
		"FROM back2school.isparent join back2school.students as s on student = s.id  "+
		"WHERE parent = $1 "+
		"order by s.name desc",
		func(rows *sql.Rows) (interface{}, error) {
			student := models.Student{}
			err := rows.Scan(&student.ID, &student.Name, &student.Surname, &student.Mail, &student.Info)
			return student, err
		}, limit, offset, id)
}

func (r *postgresRepository) PaymentsByParent(id int64, limit int, offset int) (payments []interface{}, err error) {
	return r.listByParams("select p.id, p.amount, p.student, p.payed, p.reason, p.emitted "+
		"from back2school.payments as p natural join back2school.isparent "+
		"where parent = $1 "+
		"order by p.emitted desc", func(rows *sql.Rows) (interface{}, error) {
		payment := models.Payment{}
		err := rows.Scan(&payment.ID, &payment.Amount, &payment.Student.ID, &payment.Payed, &payment.Reason, &payment.Emitted)
		return payment, err
	}, limit, offset, id)
}

func (r *postgresRepository) NotificationsByParent(id int64, limit int, offset int) (list []interface{}, err error) {
	return r.listByParams("select * from ( "+
		"select n.id, n.receiver, n.message, n.receiver_kind, n.time "+
		"from back2school.notification as n join back2school.isparent on n.receiver = student "+
		"where parent = $1 and receiver_kind = 'student' "+
		"union all  "+
		"select n.id, n.receiver, n.message, n.receiver_kind, n.time "+
		"from back2school.notification as n "+
		"where receiver_kind = 'general' "+
		") as a order by time desc",
		func(rows *sql.Rows) (interface{}, error) {
			notification := models.Notification{}
			err := rows.Scan(&notification.ID, &notification.Receiver, &notification.Message,
				&notification.ReceiverKind, &notification.Time)
			return notification, err
		}, limit, offset, id)
}

func (r *postgresRepository) AppointmentsByParent(id int64, limit int, offset int) (appointments []interface{}, err error) {
	return r.listByParams("select a.id, a.student, a.teacher, a.location, a.time "+
		"from back2school.appointments as a natural join back2school.isparent  "+
		"where parent = $1 "+
		"order by a.time desc",
		func(rows *sql.Rows) (interface{}, error) {
			appointment := models.Appointment{}
			err := rows.Scan(&appointment.ID, &appointment.Student, &appointment.Teacher, &appointment.Location, &appointment.Time)
			return appointment, err
		}, limit, offset, id)
}

func (r *postgresRepository) PaymentByID(id int64) (interface{}, error) {
	payment := &models.Payment{}
	err := r.QueryRow("SELECT id, amount, payed, emitted, reason "+
		"FROM back2school.payments WHERE id = $1 ", id).Scan(payment.ID, payment.Amount, payment.Payed, payment.Emitted, payment.Reason)
	return switchResult(payment, err)
}

func (r *postgresRepository) Payments(limit int, offset int) ([]interface{}, error) {
	return r.listByParams("select id, amount, student, payed, reason, emitted "+
		"from back2school.payments "+
		"order by payed asc, emitted asc",
		func(rows *sql.Rows) (interface{}, error) {
			payment := models.Payment{}
			err := rows.Scan(&payment.ID, &payment.Amount, &payment.Student.ID, &payment.Payed, &payment.Reason, &payment.Emitted)
			return payment, err
		}, limit, offset)
}

func (r *postgresRepository) Appointments(limit int, offset int) ([]interface{}, error) {
	return r.listByParams("select id, student, teacher, location, time "+
		"from back2school.appointments "+
		"order by time desc, teacher asc",
		func(rows *sql.Rows) (interface{}, error) {
			appointment := models.Appointment{}
			err := rows.Scan(&appointment.ID, &appointment.Student.ID, &appointment.Teacher.ID, &appointment.Location, &appointment.Time)
			return appointment, err
		}, limit, offset)
}

func (r *postgresRepository) Students(limit int, offset int) (student []interface{}, err error) {
	return r.listByParams("select id, name, surname, mail, info  "+
		"from back2school.students "+
		"order by name desc, surname desc", func(rows *sql.Rows) (interface{}, error) {
		student := models.Student{}
		err := rows.Scan(&student.ID, &student.Name, &student.Surname, &student.Mail, &student.Info)
		return student, err
	}, limit, offset)
}

func (r *postgresRepository) StudentByID(id int64) (student interface{}, err error) {
	s := models.Student{}
	err = r.QueryRow("SELECT id,	name, surname, mail, info  "+
		"FROM back2school.students WHERE id = $1", id).Scan(&s.ID,
		&s.Name, &s.Surname, &s.Mail, &s.Info)
	return switchResult(s, err)
}

func (r *postgresRepository) GradesByStudent(id int64, limit int, offset int) ([]interface{}, error) {
	return r.listByParams("SELECT id, student, subject, date, grade, teacher "+
		"FROM back2school.grades "+
		"WHERE student = $1 "+
		"order by date desc",
		func(rows *sql.Rows) (interface{}, error) {
			grade := models.Grade{}
			err := rows.Scan(&grade.ID, &grade.Student.ID, &grade.Subject, &grade.Date, &grade.Grade, &grade.Teacher.ID)
			return grade, err
		}, limit, offset, id)
}

func (r *postgresRepository) TeacherByID(id int64) (interface{}, error) {
	teacher := models.Teacher{}
	err := r.QueryRow("SELECT id, name, surname, mail  "+
		"FROM back2school.teachers "+
		"WHERE id = $1",
		id).Scan(&teacher.ID, &teacher.Name, &teacher.Surname, &teacher.Mail)
	return switchResult(teacher, err)

}

func (r *postgresRepository) Teachers(limit int, offset int) ([]interface{}, error) {
	return r.listByParams("select id, name, surname, mail, info  "+
		"from back2school.teachers "+
		"order by name desc, surname desc", func(rows *sql.Rows) (interface{}, error) {
		teacher := models.Teacher{}
		err := rows.Scan(&teacher.ID, &teacher.Name, &teacher.Surname, &teacher.Mail, &teacher.Info)
		return teacher, err
	}, limit, offset)
}

func (r *postgresRepository) AppointmentsByTeacher(id int64, limit int, offset int) (appointments []interface{}, err error) {
	return r.listByParams("SELECT id, student, teacher, location, time "+
		"FROM back2school.appointments  "+
		"WHERE teacher = $1 "+
		"order by time desc",
		func(rows *sql.Rows) (interface{}, error) {
			app := models.Appointment{}
			err := rows.Scan(&app.ID, &app.Student.ID, &app.Teacher.ID, &app.Location, &app.Time)
			return app, err
		}, limit, offset, id)
}

func (r *postgresRepository) NotificationsByTeacher(id int64, limit int, offset int) (notifications []interface{}, err error) {
	return r.listByParams("SELECT id, receiver, message, receiver_kind, time  "+
		"FROM back2school.notification  "+
		"WHERE (receiver = $1 and receiver_kind = 'teacher') or receiver_kind = 'general' "+
		"order by time desc",
		func(rows *sql.Rows) (interface{}, error) {
			notif := models.Notification{}
			err := rows.Scan(&notif.ID, &notif.Receiver, &notif.Message, &notif.ReceiverKind, &notif.Time)
			return notif, err
		}, limit, offset, id)
}

func (r *postgresRepository) SubjectsByTeacher(id int64, limit int, offset int) (notifications []interface{}, err error) {
	return r.listByParams("SELECT DISTINCT subject FROM back2school.teaches where teacher = $1 order by subject",
		func(rows *sql.Rows) (interface{}, error) {
			subj := ""
			err := rows.Scan(&subj)
			return subj, err
		}, limit, offset, id)
}

func (r *postgresRepository) ClassesBySubjectAndTeacher(teacher int64, subject string, limit int, offset int) ([]interface{}, error) {
	return r.listByParams("SELECT id, year, section, info, grade "+
		"FROM back2school.teaches join back2school.classes on id = class  "+
		"WHERE teacher = $1 and subject = $2 "+
		"order by year desc, grade asc, section desc ",
		func(rows *sql.Rows) (interface{}, error) {
			class := models.Class{}
			err := rows.Scan(&class.ID, &class.Year, &class.Section, &class.Info, &class.Grade)
			return class, err
		}, limit, offset, teacher, subject)
}

func (r *postgresRepository) LecturesByTeacher(id int64, limit int, offset int) (lectures []interface{}, err error) {
	return r.listByParams("SELECT id, class, subject, location, start, \"end\", info	 "+
		"from back2school.timetable natural join back2school.teaches as t "+
		"where t.teacher = $1 "+
		"order by start desc",
		func(rows *sql.Rows) (interface{}, error) {
			lecture := models.TimeTable{}
			err := rows.Scan(&lecture.ID, &lecture.Class.ID, &lecture.Subject, &lecture.Location, &lecture.Start, &lecture.End, &lecture.Info)
			return lecture, err

		}, limit, offset, id)
}

func (r *postgresRepository) ClassesByTeacher(id int64, limit int, offset int) ([]interface{}, error) {
	type Class struct {
		models.Class
		Subject models.Subject `json:"subject,omitempty"`
	}
	return r.listByParams("SELECT id, year, section, info, grade, subject "+
		"FROM back2school.teaches join back2school.classes on id = class  "+
		"WHERE teacher = $1 "+
		"order by subject asc, year desc, grade asc, section desc ",
		func(rows *sql.Rows) (interface{}, error) {
			class := Class{}
			err := rows.Scan(&class.ID, &class.Year, &class.Section, &class.Info, &class.Grade, &class.Subject)
			return class, err
		}, limit, offset, id)
}

func (r *postgresRepository) UpdateTeacher(teacher models.Teacher) (err error) {
	query := "UPDATE back2school.teachers" +
		" SET name = $1, surname = $2, mail = $3, info = $4 " +
		"where id = $5"
	return r.execUpdate(query, teacher.Name, teacher.Surname, teacher.Mail, teacher.Info, teacher.ID)
}

func (r *postgresRepository) UpdateParent(parent models.Parent) (err error) {
	query := "UPDATE back2school.parents" +
		" SET name = $1, surname = $2, mail = $3, info = $4 " +
		"where id = $5"
	return r.execUpdate(query, parent.Name, parent.Surname, parent.Mail, parent.Info, parent.ID)
}

func (r *postgresRepository) UpdateStudent(student models.Student) (err error) {
	query := "UPDATE back2school.student" +
		" SET name = $1, surname = $2, mail = $3, info = $4 " +
		"where id = $5"
	return r.execUpdate(query, student.Name, student.Surname, student.Mail, student.Info, student.ID)
}

func (r *postgresRepository) UpdateAppointment(appointment models.Appointment) (err error) {
	query := "UPDATE back2school.appointments " +
		"SET student = $1, teacher = $2, location = $3, time = $4 where id = $5"
	return r.execUpdate(query, appointment.Student, appointment.Teacher, appointment.Location, appointment.Time, appointment.ID)
}

func (r *postgresRepository) CreateAppointment(appointment models.Appointment) (id int64, err error) {
	query := "INSERT INTO back2school.appointments " +
		" (student, teacher, location, time) VALUES ($1, $2, $3, $4)"
	return r.exec(query, appointment.Student, appointment.Teacher, appointment.Location, appointment.Time)
}

func (r *postgresRepository) CreateParent(parent models.Parent) (int64, error) {
	query := "INSERT INTO back2school.parents " +
		"(name, surname, mail, info) VALUES ($1, $2, $3, $4)"
	return r.exec(query, parent.Name, parent.Surname, parent.Mail, parent.Info)

}

func (r *postgresRepository) CreateTeacher(teacher models.Teacher) (int64, error) {
	query := "INSERT INTO back2school.teachers" +
		" (name, surname, mail, info)" +
		" VALUES ($1, $2, $3, $4)"
	return r.exec(query, teacher.Name, teacher.Surname, teacher.Mail, teacher.Info)
}

func (r *postgresRepository) CreateStudent(student models.Student) (int64, error){
	query := "INSERT INTO back2school.students" +
		" (name, surname, mail, info) " +
		" VALUES ($1, $2, $3, $4)"
	return r.exec(query, student.Name, student.Surname, student.Mail, student.Info)
}

func (r *postgresRepository) CreateClass(class models.Class) (int64, error){
	query := "INSERT INTO back2school.classes" +
		" (year, section, info, grade) " +
		" VALUES ($1, $2, $3, $4)"
		return r.exec(query, class.Year, class.Section, class.Info, class.Grade)
}

func (r *postgresRepository) UpdateClass(class models.Class) (err error) {
	query := "UPDATE back2school.classes " +
		" SET year = $1, section = $2, info = $3, grade = $4" +
		" where id = $5"
	return r.execUpdate(query, class.Year, class.Section, class.Info, class.Grade, class.ID)
}

func (r *postgresRepository) CreateNotification(notification models.Notification) (int64, error)  {
	query := "insert into back2school.classes " +
		" (receiver, message, time, receiver_kind) " +
		" VALUES ($1, $2, $3, $4) "
	return r.exec(query, notification.Receiver, notification.Message, notification.Time, notification.ReceiverKind)
}
func (r *postgresRepository) UpdateNotification(notification models.Notification) error {
	query := "UPDATE back2school.notification " +
		"SET receiver = $1, message = $2, time = $3, receiver_kind = $4" +
		" where id = $5"
	return r.execUpdate(query, notification.Receiver, notification.Message, notification.Time, notification.ReceiverKind, notification.ID)
}


func (r *postgresRepository) CreateGrade(grade models.Grade) (int64, error)  {
	query := "insert into back2school.grades " +
		" (student, grade, subject, date, teacher) " +
		" VALUES ($1, $2, $3, $4, $5) "
		return r.exec(query, grade.Student.ID, grade.Grade, grade.Subject, grade.Date, grade.Teacher.ID)
}
func (r *postgresRepository) UpdateGrade(grade models.Grade) error {
	query := "UPDATE back2school.grades " +
		"SET student = $1, grade = $2, subject = $3, date = $4, teacher = $5) " +
		" where id = $6"
	return r.execUpdate(query, grade.Student.ID, grade.Grade, grade.Subject, grade.Date, grade.Teacher.ID, grade.ID)
}

func (r *postgresRepository) CreatePayment(payment models.Payment) (int64, error)  {
	query := "insert into back2school.payments " +
		" (amount, student, payed, reason, emitted) " +
		" VALUES ($1, $2, $3, $4, $5) "
	return r.exec(query, payment.Amount, payment.Student.ID, payment.Payed, payment.Reason, payment.Emitted)
}
func (r *postgresRepository) UpdatePayment(payment models.Payment) error {
	query := "UPDATE back2school.grades " +
		"SET amount = $1, student = $2, payed = $3, reason = $4, emitted = $5" +
		" where id = $6"
	return r.execUpdate(query, payment.Amount, payment.Student.ID, payment.Payed, payment.Reason, payment.Emitted, payment.ID)
}