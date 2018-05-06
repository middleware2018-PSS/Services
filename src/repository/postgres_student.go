package repository

import (
	"github.com/middleware2018-PSS/Services/src/models"
	"log"
)

func (r *postgresRepository) StudentById(id int) (student *models.Student) {

	// get basic data
	student = &models.Student{}
	err := r.QueryRow(`SELECT id,	name, surname, mail 
								FROM back2school.students WHERE id = $1`, id).Scan(&student.ID,
		&student.Name, &student.Surname, &student.Mail)
	if err != nil {
		log.Print(err)
	}

	// Payments

	student.Payments = r.PaymentByStudent(student.ID)

	// Grades
	student.Grades = r.GradesByStudent(student.ID)

	// Notifications

	student.Notifications = r.NotificationsByStudent(student.ID)

	// Classes
	student.Classes = r.ClassesByStudent(student.ID)

	// Appointments
	student.Appointments = r.AppointmentsByStudent(student.ID)

	return student
}

func (r *postgresRepository) PaymentByStudent(id int64) (payments []models.Payment) {
	rows, err := r.Query(`SELECT id
								FROM back2school.payments WHERE student = $1 `, id)
	defer rows.Close()
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		payment := models.Payment{}
		rows.Scan(&payment.ID)
		payments = append(payments, payment)
	}
	return payments
}

func (r *postgresRepository) GradesByStudent(id int64) (grades []models.Grade) {
	rows, err := r.Query(`SELECT student, subject, date
									FROM back2school.grades WHERE student = $1`, id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		grade := models.Grade{}
		rows.Scan(&grade.Student.ID, &grade.Subject, &grade.Date)
		grades = append(grades, grade)
	}
	return grades
}

func (r *postgresRepository) NotificationsByStudent(id int64) (notifications []models.Notification) {
	rows, err := r.Query(`SELECT id FROM back2school.Notification 
								WHERE (receiver = $1 and receiver_kind = 'student') or receiver_kind = 'general'`, id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		notif := models.Notification{}
		rows.Scan(&notif.ID)
		notifications = append(notifications, notif)
	}
	return notifications
}

func (r *postgresRepository) ClassesByStudent(id int64) (classes []models.Class) {
	rows, err := r.Query(`SELECT class FROM back2school.enrolled
			WHERE student = $1`, id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		class := models.Class{}
		rows.Scan(&class.ID)
		classes = append(classes, class)
	}
	return classes
}

func (r *postgresRepository) AppointmentsByStudent(id int64) (appointments []models.Appointment) {

	rows, err := r.Query(`SELECT id FROM back2school.appointments 
								WHERE student = $1`, id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		app := models.Appointment{}
		rows.Scan(&app.ID)
		appointments = append(appointments, app)
	}
	return appointments
}
