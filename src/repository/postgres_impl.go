package repository

import (
	"database/sql"
	"github.com/middleware2018-PSS/Services/src/models"
	"log"
)

func (r *postgresRepository) UpdateAppointments(id int64) (err error) {
	//TODO
	return
}

func (r *postgresRepository) AppointmentById(id int64) (interface{}, error) {
	appointment := models.Appointment{}
	err := r.QueryRow(`SELECT id, student, teacher, time, location 
								FROM back2school.appointments WHERE id = $1 `, id).Scan(
		&appointment.ID, &appointment.Student.ID, &appointment.Teacher.ID, &appointment.Time, &appointment.Location)
	return appointment, switchError(err)
}

func (r *postgresRepository) ClassByID(id int64) (interface{}, error) {
	class := models.Class{}
	err := r.QueryRow(`SELECT id, year, section, info, grade FROM back2school.classes 
								WHERE id = $1`, id).Scan(&class.ID, &class.Year, &class.Section, &class.Info, &class.Grade)
	return class, switchError(err)
}

func (r *postgresRepository) Classes(limit int, offset int) ([]interface{}, error) {
	return r.listByParams(`select id, year, section, info, grade
						from back2school.classes
						order by year desc, grade asc, section asc`, func(rows *sql.Rows) (interface{}, error) {
		class := models.Class{}
		err := rows.Scan(&class.ID, &class.Year, &class.Section, &class.Info, &class.Grade)
		return class, err
	}, limit, offset)
}

func (r *postgresRepository) StudentByClass(id int64, limit int, offset int) (students []interface{}, err error) {
	return r.listByParams(`select id, name, surname, mail, info 
						from back2school.students join back2school.enrolled on student = id 
						where class = $1
						order by name desc, surname desc`,
		func(rows *sql.Rows) (interface{}, error) {
			student := models.Student{}
			err := rows.Scan(&student.ID, &student.Name, &student.Surname, &student.Mail, &student.Info)
			return student, err
		}, limit, offset, id)
}

func (r *postgresRepository) LectureByClass(id int64, limit int, offset int) ([]interface{}, error) {
	return r.listByParams(
		`select id, class, subject, "start", "end", location, info
				from back2school.timetable natural join back2school.teaches
				where teacher = $1
				order by "start" desc`,
		func(rows *sql.Rows) (interface{}, error) {
			lecture := models.TimeTable{}
			err := rows.Scan(&lecture.ID, &lecture.Class, &lecture.Subject, &lecture.Start, &lecture.End, &lecture.Location, &lecture.Info)
			return lecture, err
		}, limit, offset, id)
}

func (r *postgresRepository) NotificationByID(id int64) (interface{}, error) {
	n := models.Notification{}
	err := r.QueryRow(`SELECT id, receiver, message, time, receiver_kind
								FROM back2school.notification WHERE id = $1 `, id).Scan(&n.ID,
		&n.Receiver, &n.Message, &n.Time, &n.ReceiverKind)
	return n, switchError(err)
}

func (r *postgresRepository) Notifications(limit int, offset int) ([]interface{}, error) {
	return r.listByParams(`select id, receiver, message, time, receiver_kind 
						from back2school.notification
						order by time desc, receiver_kind desc`,
		func(rows *sql.Rows) (interface{}, error) {
			notification := models.Notification{}
			err := rows.Scan(&notification.ID, &notification.Receiver, &notification.Message, &notification.Time, &notification.ReceiverKind)
			return notification, err
		}, limit, offset)
}

func (r *postgresRepository) UpdateParent(id int64) (err error) {
	// TODO
	return
}

func (r *postgresRepository) ParentById(id int64) (interface{}, error) {
	p := models.Parent{}
	err := r.QueryRow(`SELECT id,	name, surname, mail, info
								FROM back2school.parents WHERE id = $1`,
		id).Scan(&p.ID, &p.Name, &p.Surname, &p.Mail, &p.Info)
	return p, switchError(err)
}

func (r *postgresRepository) Parents(limit int, offset int) ([]interface{}, error) {
	return r.listByParams(`select id, name, surname, mail, info 
						from back2school.parents
						order by name desc, surname desc`,
		func(rows *sql.Rows) (interface{}, error) {
			parent := models.Parent{}
			err := rows.Scan(&parent.ID, &parent.Name, &parent.Surname, &parent.Mail, &parent.Info)
			return parent, err
		}, limit, offset)
}

func (r *postgresRepository) ChildrenByParent(id int64, limit int, offset int) (children []interface{}, err error) {
	return r.listByParams(`SELECT s.id, s.name, s.surname, s.mail, s.info
								FROM back2school.isparent join back2school.students as s on student = s.id 
								WHERE parent = $1
								order by s.name desc`,
		func(rows *sql.Rows) (interface{}, error) {
			student := models.Student{}
			err := rows.Scan(&student.ID, &student.Name, &student.Surname, &student.Mail, &student.Info)
			return student, err
		}, limit, offset, id)
}

func (r *postgresRepository) PaymentsByParent(id int64, limit int, offset int) (payments []interface{}, err error) {
	return r.listByParams(`select p.id, p.amount, p.student, p.payed, p.reason, p.emitted
		from back2school.payments as p natural join back2school.isparent
		where parent = $1
		order by p.emitted desc`, func(rows *sql.Rows) (interface{}, error) {
		payment := models.Payment{}
		err := rows.Scan(&payment.ID, &payment.Amount, &payment.Student.ID, &payment.Payed, &payment.Reason, &payment.Emitted)
		return payment, err
	}, limit, offset, id)
}

func (r *postgresRepository) NotificationsByParent(id int64, limit int, offset int) (list []interface{}, err error) {
	return r.listByParams(`select * from (
				select n.id, n.receiver, n.message, n.receiver_kind, n.time
				from back2school.notification as n join back2school.isparent on n.receiver = student
				where parent = $1 and receiver_kind = 'student'
				union all 
				select n.id, n.receiver, n.message, n.receiver_kind, n.time
				from back2school.notification as n
				where receiver_kind = 'general'
				) as a order by time desc`,
		func(rows *sql.Rows) (interface{}, error) {
			notification := models.Notification{}
			err := rows.Scan(&notification.ID, &notification.Receiver, &notification.Message,
				&notification.ReceiverKind, &notification.Time)
			return notification, err
		}, limit, offset, id)
}

func (r *postgresRepository) AppointmentsByParent(id int64, limit int, offset int) (appointments []interface{}, err error) {
	return r.listByParams(`select a.id, a.student, a.teacher, a.location, a.time
				from back2school.appointments as a natural join back2school.isparent 
				where parent = $1
				order by a.time desc`,
		func(rows *sql.Rows) (interface{}, error) {
		appointment := models.Appointment{}
		err := rows.Scan(&appointment.ID, &appointment.Student, &appointment.Teacher, &appointment.Location, &appointment.Time)
		return appointment, err
	}, limit, offset, id)
}

func (r *postgresRepository) PaymentByID(id int64) (interface{}, error) {
	payment := &models.Payment{}
	err := r.QueryRow(`SELECT id, amount, payed, emitted, reason
								FROM back2school.payments WHERE id = $1 `, id).Scan(payment.ID, payment.Amount, payment.Payed, payment.Emitted, payment.Reason)
	return payment, switchError(err)
}

func (r *postgresRepository) Payments(limit int, offset int) ([]interface{}, error) {
	return r.listByParams(`select id, amount, student, payed, reason, emitted
						from back2school.payments
						order by payed asc, emitted asc`,
		func(rows *sql.Rows) (interface{}, error) {
			payment := models.Payment{}
			err := rows.Scan(&payment.ID, &payment.Amount, &payment.Student.ID, &payment.Payed, &payment.Reason, &payment.Emitted)
			return payment, err
		}, limit, offset)
}

func (r *postgresRepository) Students(limit int, offset int) (student []interface{}, err error) {
	return r.listByParams(`select id, name, surname, mail, info 
						from back2school.students
						order by name desc, surname desc`, func(rows *sql.Rows) (interface{}, error) {
		student := models.Student{}
		err := rows.Scan(&student.ID, &student.Name, &student.Surname, &student.Mail, &student.Info)
		return student, err
	}, limit, offset)
}

func (r *postgresRepository) StudentById(id int64) (student interface{}, err error) {
	s := models.Student{}
	err = r.QueryRow(`SELECT id,	name, surname, mail, info 
								FROM back2school.students WHERE id = $1`, id).Scan(&s.ID,
		&s.Name, &s.Surname, &s.Mail, &s.Info)
	return s, switchError(err)
}

func (r *postgresRepository) GradesByStudent(id int64, limit int, offset int) ([]interface{}, error) {
	return r.listByParams(`SELECT student, subject, date, grade, teacher
									FROM back2school.grades
									WHERE student = $1
									order by date desc`,
		func(rows *sql.Rows) (interface{}, error) {
			grade := models.Grade{}
			err := rows.Scan(&grade.Student.ID, &grade.Subject, &grade.Date, &grade.Grade, &grade.Teacher.ID)
			return grade, err
		}, limit, offset, id)
}

func (r *postgresRepository) GradeStudent(grade models.Grade) (err error) {
	// TODO
	return
}

func (r *postgresRepository) UpdateStudent(id int64) (err error) {
	// TODO
	return
}

func (r *postgresRepository) TeacherByID(id int64) (interface{}, error) {
	teacher := models.Teacher{}
	err := r.QueryRow(`SELECT id, name, surname, mail 
							FROM back2school.teachers
							WHERE id = $1`,
		id).Scan(&teacher.ID, &teacher.Name, &teacher.Surname, &teacher.Mail)
	return teacher, switchError(err)

}

func (r *postgresRepository) Teachers(limit int, offset int) ([]interface{}, error) {
	return r.listByParams(`select id, name, surname, mail, info 
						from back2school.teachers
						order by name desc, surname desc`, func(rows *sql.Rows) (interface{}, error) {
		teacher := models.Teacher{}
		err := rows.Scan(&teacher.ID, &teacher.Name, &teacher.Surname, &teacher.Mail, &teacher.Info)
		return teacher, err
	}, limit, offset)
}

func (r *postgresRepository) ClassesPerSubjectByTeacher(id int64) (classes map[models.Subject][]models.Class, err error) {
	// TODO check errors
	classes = make(map[models.Subject][]models.Class)
	rows, err := r.Query(`SELECT subject, id, year, section, info, grade
								FROM back2school.teaches join back2school.classes on id = class 
								WHERE teacher = $1
								order by subject asc, year desc, grade asc, section desc `, id)
	if err != nil {
		log.Print(err)
	}
	defer rows.Close()
	for rows.Next() {
		class := models.Class{}
		var subj models.Subject
		rows.Scan(&subj, &class.ID, &class.Year, &class.Section, &class.Info, &class.Grade)
		classes[subj] = append(classes[subj], class)
	}
	return classes, switchError(err)
}

func (r *postgresRepository) AppointmentsByTeacher(id int64, limit int, offset int) (appointments []interface{}, err error) {
	return r.listByParams(`SELECT id, student, teacher, location, time
										FROM back2school.appointments 
										WHERE teacher = $1
										order by time desc`,
		func(rows *sql.Rows) (interface{}, error) {
			app := models.Appointment{}
			err := rows.Scan(&app.ID, &app.Student.ID, &app.Teacher.ID, &app.Location, &app.Time)
			return app, err
		}, limit, offset, id)
}

func (r *postgresRepository) NotificationsByTeacher(id int64, limit int, offset int) (notifications []interface{}, err error) {
	return r.listByParams(`SELECT id, receiver, message, receiver_kind, time 
								FROM back2school.notification 
								WHERE (receiver = $1 and receiver_kind = 'teacher') or receiver_kind = 'general'
								order by time desc`,
		func(rows *sql.Rows) (interface{}, error) {
			notif := models.Notification{}
			err := rows.Scan(&notif.ID, &notif.Receiver, &notif.Message, &notif.ReceiverKind, &notif.Time)
			return notif, err
		}, limit, offset, id)
}

func (r *postgresRepository) LectureByTeacher(id int64, limit int, offset int) (lectures []interface{}, err error) {
	return r.listByParams(`SELECT id, class, subject, location, start, "end", info	
								from back2school.timetable natural join back2school.teaches as t
								where t.teacher = $1
								order by start desc`,
		func(rows *sql.Rows) (interface{}, error) {
			lecture := models.TimeTable{}
			err := rows.Scan(&lecture.ID, &lecture.Class.ID, &lecture.Subject, &lecture.Location, &lecture.Start, &lecture.End, &lecture.Info)
			return lecture, err

		}, limit, offset, id)
}

func (r *postgresRepository) UpdateTeacher(id int64) (err error) {
	//TODO
	return
}