package repository

import (
	"database/sql"
	"encoding/json"
	"fmt"
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

	rows, err = r.Query(`SELECT id FROM back2school.payments WHERE student = $1`, id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		payment := models.Payment{}
		rows.Scan(&payment.ID)
		student.Payments = append(student.Payments, payment)
	}

	// Grades

	rows, err = r.Query(`SELECT student, subject, grade 
								FROM back2school.grades WHERE student = $1`, id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		grade := models.Grade{}
		rows.Scan(&grade.Student.ID, &grade.Subject, &grade.Grade)
		student.Grades = append(student.Grades, grade)
	}

	// Notifications

	rows, err = r.Query(`SELECT id FROM back2school.Notification 
								WHERE (receiver = $1 and receiver_kind = 'student') or receiver_kind = 'general'`, id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		notif := models.Notification{}
		rows.Scan(&notif.ID)
		student.Notifications = append(student.Notifications, notif)
	}

	// Classes

	rows, err = r.Query(`SELECT class FROM back2school.enrolled 
								WHERE student = $1`, id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		class := models.Class{}
		rows.Scan(&class.ID)
		student.Classes = append(student.Classes, class)
	}

	// Appointments

	rows, err = r.Query(`SELECT id FROM back2school.appointments 
								WHERE student = $1`, id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		app:= models.Appointment{}
		rows.Scan(&app.ID)
		student.Appointments = append(student.Appointments, app)
	}


	s, _ := json.MarshalIndent(student, " ", "  ")
	fmt.Printf("%s\n", s)



	return student
}


func (r *postgresRepository) ParentById(id int) (parent *models.Parent) {

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

	rows, err = r.Query(`SELECT student
								FROM back2school.isparent WHERE parent = $1`, id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		student := models.Student{}
		rows.Scan(&student.ID)
		parent.ParentOf = append(parent.ParentOf, student)
	}

	s, _ := json.MarshalIndent(parent, " ", "  ")
	fmt.Printf("%s\n", s)

	return parent
}


