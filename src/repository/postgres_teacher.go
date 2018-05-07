package repository

import (
	"database/sql"
	"github.com/middleware2018-PSS/Services/src/models"
	"log"
)

func (r *postgresRepository) TeacherByID(id int64) (teacher *models.Teacher, err error) {

	teacher = &models.Teacher{}

	// general info

	err = r.QueryRow(`SELECT id, name, surname, mail FROM back2school.teachers WHERE id = $1`,
		id).Scan(&teacher.ID, &teacher.Name, &teacher.Surname, &teacher.Mail)
	if err != nil {
		log.Print(err)
	}
	if err == sql.ErrNoRows {
		return nil, err
	}

	// Classes

	teacher.Classes = r.ClassesByTeacher(teacher.ID)

	// Appointments
	teacher.Appointments = r.AppointmentsByTeacher(teacher.ID)

	// Lectures
	teacher.Lectures = r.LectureByTeacher(teacher.ID)

	// Notifications
	teacher.Notifications = r.NotificationsByTeacher(teacher.ID)

	return teacher, nil

}

func (r *postgresRepository) ClassesByTeacher(id int64) (classes map[models.Subject][]models.Class) {
	classes = make(map[models.Subject][]models.Class)
	rows, err := r.Query(`SELECT subject, class FROM back2school.teaches 
								WHERE teacher = $1`, id)
	if err != nil {
		log.Print(err)
	}
	defer rows.Close()
	for rows.Next() {
		class := models.Class{}
		var subj models.Subject
		rows.Scan(&subj, &class.ID)
		classes[subj] = append(classes[subj], class)
	}
	return classes
}

func (r *postgresRepository) AppointmentsByTeacher(id int64) (appointments []models.Appointment) {

	rows, err := r.Query(`SELECT id FROM back2school.appointments 
								WHERE teacher = $1`, id)
	if err != nil {
		log.Print(err)
	}
	defer rows.Close()
	for rows.Next() {
		app := models.Appointment{}
		rows.Scan(&app.ID)
		appointments = append(appointments, app)
	}
	return appointments
}

func (r *postgresRepository) NotificationsByTeacher(id int64) (notifications []models.Notification) {
	rows, err := r.Query(`SELECT id FROM back2school.Notification 
								WHERE (receiver = $1 and receiver_kind = 'teacher') or receiver_kind = 'general'`, id)
	if err != nil {
		log.Print(err)
	}
	defer rows.Close()
	for rows.Next() {
		notif := models.Notification{}
		rows.Scan(&notif.ID)
		notifications = append(notifications, notif)
	}
	return notifications
}

func (r *postgresRepository) LectureByTeacher(id int64) (lectures []models.TimeTable) {
	rows, err := r.Query(`SELECT class, subject, date from back2school.timetable natural join back2school.teaches as t
								where t.teacher = $1`, id)
	if err != nil {
		log.Print(err)
	}
	defer rows.Close()
	for rows.Next() {
		lecture := models.TimeTable{}
		rows.Scan(&lecture.Class, &lecture.Subject, &lecture.Date)
		lectures = append(lectures, lecture)
	}
	return lectures
}
