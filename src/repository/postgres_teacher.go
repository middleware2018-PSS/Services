package repository

import (
	"database/sql"
	"github.com/middleware2018-PSS/Services/src/models"
	"log"
)

func (r *postgresRepository) TeacherByID(id int64) (interface{}, error) {
	teacher := models.Teacher{}
	err := r.QueryRow(`SELECT id, name, surname, mail 
							FROM back2school.teachers
							WHERE id = $1`,
		id).Scan(&teacher.ID, &teacher.Name, &teacher.Surname, &teacher.Mail)
	if err != nil {
		log.Print(err)
	}
	return teacher, switchError(err)

}

func (r *postgresRepository) ClassesPerSubjectByTeacher(id int64) (classes map[models.Subject][]models.Class, err error) {
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
	return classes, err
}

func (r *postgresRepository) AppointmentsByTeacher(id int64, offset int, limit int) (appointments []interface{}, err error) {
	return r.listByID(id, offset, limit, `SELECT id, student, teacher, location, time
										FROM back2school.appointments 
										WHERE teacher = $1
										order by time desc
										limit $2 offset $3`,
		func(rows *sql.Rows) (interface{}, error) {
			app := models.Appointment{}
			err := rows.Scan(&app.ID, &app.Student.ID, &app.Teacher.ID, &app.Location, &app.Time)
			return app, err
		})
}

func (r *postgresRepository) NotificationsByTeacher(id int64, offset int, limit int) (notifications []interface{}, err error) {
	return r.listByID(id, offset, limit, `SELECT id, receiver, message, receiver_kind, time 
								FROM back2school.notification 
								WHERE (receiver = $1 and receiver_kind = 'teacher') or receiver_kind = 'general'
								order by time desc
								limit $2 offset $3`,
		func(rows *sql.Rows) (interface{}, error) {
			notif := models.Notification{}
			err := rows.Scan(&notif.ID, &notif.Receiver, &notif.Message, &notif.ReceiverKind, &notif.Time)
			return notif, err
		})
}

func (r *postgresRepository) LectureByTeacher(id int64, offset int, limit int) (lectures []interface{}, err error) {
	return r.listByID(id, offset, limit, `SELECT id, class, subject, location, start, end, info	
								from back2school.timetable natural join back2school.teaches as t
								where t.teacher = $1
								order by start desc
								limit $2 offset $3`, func(rows *sql.Rows) (interface{}, error) {
		lecture := models.TimeTable{}
		err := rows.Scan(&lecture.ID, &lecture.Class.ID, &lecture.Subject, &lecture.Location, &lecture.Start, &lecture.End, &lecture.Info)
		return lecture, err

	})
}

func (r *postgresRepository) UpdateTeacher(id int64) (err error) {
	//TODO
	return
}
