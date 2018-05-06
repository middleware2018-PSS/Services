package repository

import (
	"database/sql"
	"github.com/middleware2018-PSS/Services/src/models"
	"log"
)

type postgresRepository struct {
	*sql.DB
}

func NewPostgresRepository(DB *sql.DB) *postgresRepository {
	return &postgresRepository{DB}
}

func (r *postgresRepository) StudentById(id int) (student *models.Student) {

	// get basic data

	rows, err := r.Query(`SELECT id,	name, surname, mail 
								FROM back2school.students WHERE id = $1`, id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	student = &models.Student{}
	for rows.Next() {
		rows.Scan(&student.ID, &student.Name, &student.Surname, &student.Mail)
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

	rows, err = r.Query(`SELECT id FROM back2school.appointments 
								WHERE student = $1`, id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		app := models.Appointment{}
		rows.Scan(&app.ID)
		student.Appointments = append(student.Appointments, app)
	}

	return student
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

func (r *postgresRepository) ParentById(id int64) (parent *models.Parent) {

	// get basic data

	rows, err := r.Query(`SELECT id,	name, surname, mail 
								FROM back2school.parents WHERE id = $1`, id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	parent = &models.Parent{}
	for rows.Next() {
		rows.Scan(&parent.ID, &parent.Name, &parent.Surname, &parent.Mail)
	}

	// Childrens
	parent.ParentOf = r.ChildrenByParent(parent.ID)

	// Payments

	parent.Payments = r.PaymentsByParent(parent.ID)

	// Notifications

	parent.Notifications = r.NotificationsByParent(parent.ID)

	return parent
}

func (r *postgresRepository) ChildrenByParent(id int64) (child []models.Student) {

	rows, err := r.Query(`SELECT student
								FROM back2school.isparent WHERE parent = $1`, id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		student := models.Student{}
		rows.Scan(&student.ID)
		child = append(child, student)
	}
	return child
}

func (r *postgresRepository) PaymentsByParent(id int64) (payments []models.Payment) {
	query := `select p.id
		from back2school.payments as p natural join back2school.isparent
		where parent = $1`
	rows, err := r.Query(query, id)
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

func (r *postgresRepository) NotificationsByParent(id int64) (list []models.Notification) {
	query := `select n.id
				from back2school.notification as n join back2school.isparent on n.receiver = student
				where parent = $1 and receiver_kind = 'student'
				union all select id from back2school.notification where receiver_kind = 'general'`
	rows, err := r.Query(query, id)
	defer rows.Close()
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		notification := models.Notification{}
		rows.Scan(&notification.ID)
		list = append(list, notification)

	}
	return list
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

func (r *postgresRepository) PaymentByID(id int64) (payments []models.Payment) {
	rows, err := r.Query(`SELECT id, amount, payed, emitted, reason
								FROM back2school.payments WHERE id = $1 `, id)
	defer rows.Close()
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		payment := &models.Payment{}
		rows.Scan(payment.ID, payment.Amount, payment.Payed, payment.Emitted, payment.Reason)
		payments = append(payments, *payment)
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
