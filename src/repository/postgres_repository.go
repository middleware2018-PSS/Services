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

func (r *postgresRepository) AppointmentsByStudent(id int64) (appointments []models.Appointment){

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

func (r *postgresRepository) ParentById(id int64) (parent *models.Parent) {

	// get basic data
	parent = &models.Parent{}

	err := r.QueryRow(`SELECT id,	name, surname, mail 
								FROM back2school.parents WHERE id = $1`, id).Scan(&parent.ID, &parent.Name, &parent.Surname, &parent.Mail)
	if err != nil {
		log.Print(err)
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

func (r *postgresRepository) TeacherByID(id int64) (teacher *models.Teacher, err error){

	teacher = &models.Teacher{}

	// general info

	err = r.QueryRow(`SELECT id, name, surname, mail FROM back2school.teachers WHERE id = $1`,
						id).Scan(&teacher.ID, &teacher.Name, &teacher.Surname, &teacher.Mail)
	if err != nil {
		log.Print(err)
	}
	if err == sql.ErrNoRows{
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

func (r *postgresRepository) LectureByTeacher(id int64) (lectures []models.TimeTable){
	rows, err := r.Query(`SELECT class, subject, date from back2school.timetable natural join back2school.teaches as t
								where t.teacher = $1`, id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		lecture := models.TimeTable{}
		rows.Scan(&lecture.Class, &lecture.Subject, &lecture.Date)
		lectures = append(lectures, lecture)
	}
	return lectures
}


func (r *postgresRepository) ClassesByTeacher(id int64) (classes map[models.Subject][]models.Class){
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
		rows.Scan(&subj,&class.ID)
		classes[subj] = append(classes[subj],class)
	}
	return classes
	}




func (r *postgresRepository) AppointmentsByTeacher(id int64) (appointments []models.Appointment){

	rows, err := r.Query(`SELECT id FROM back2school.appointments 
								WHERE teacher = $1`, id)
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


func (r *postgresRepository) NotificationsByTeacher(id int64) (notifications []models.Notification) {
	rows, err := r.Query(`SELECT id FROM back2school.Notification 
								WHERE (receiver = $1 and receiver_kind = 'teacher') or receiver_kind = 'general'`, id)
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


