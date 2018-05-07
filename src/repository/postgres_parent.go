package repository

import (
	"github.com/middleware2018-PSS/Services/src/models"
	"log"
)

func (r *postgresRepository) UpdateParent(id int64) (err error) {
	// TODO
	return
}

func (r *postgresRepository) ParentById(id int64) (parent models.Parent, err error) {
	// get basic data}

	err = r.QueryRow(`SELECT id,	name, surname, mail, info
								FROM back2school.parents WHERE id = $1`,
		id).Scan(&parent.ID, &parent.Name, &parent.Surname, &parent.Mail, &parent.Info)
	if err != nil {
		log.Print(err)
	}
	return parent, err
}

func (r *postgresRepository) ChildrenByParent(id int64, offset int, limit int) (children []models.Student, err error) {
	rows, err := r.Query(`SELECT s.id, s.name, s.surname, s.mail, s.info
								FROM back2school.isparent join back2school.students as s on student = s.id 
								WHERE parent = $1
								order by s.name desc
								limit $2 offset $3`, id, limit, offset)
	if err != nil {
		log.Print(err)
	}
	defer rows.Close()
	for rows.Next() {
		student := models.Student{}
		rows.Scan(&student.ID, &student.Name, &student.Surname, &student.Mail, &student.Info)
		children = append(children, student)
	}
	return children, err
}

func (r *postgresRepository) PaymentsByParent(id int64, offset int, limit int) (payments []models.Payment, err error) {
	rows, err := r.Query(`select p.id, p.amount, p.student, p.payed, p.reason, p.emitted
		from back2school.payments as p natural join back2school.isparent
		where parent = $1
		order by p.emitted desc
		limit $2 offset $3`, id, limit, offset)
	defer rows.Close()
	if err != nil {
		log.Print(err)
	}
	for rows.Next() {
		payment := models.Payment{}
		rows.Scan(&payment.ID, &payment.Amount, &payment.Student.ID, &payment.Payed, &payment.Reason, &payment.Emitted)
		payments = append(payments, payment)
	}
	return payments, err
}

func (r *postgresRepository) NotificationsByParent(id int64, offset int, limit int) (list []models.Notification, err error) {
	query := `select * from (
				select n.id, n.receiver, n.message, n.receiver_kind, n.time
				from back2school.notification as n join back2school.isparent on n.receiver = student
				where parent = $1 and receiver_kind = 'student'
				union all 
				select n.id, n.receiver, n.message, n.receiver_kind, n.time
				from back2school.notification as n
				where receiver_kind = 'general'
				) as a order by time desc
				limit $2 offset $3`
	rows, err := r.Query(query, id, limit, offset)
	defer rows.Close()
	if err != nil {
		log.Print(err)
	}
	for rows.Next() {
		notification := models.Notification{}
		rows.Scan(&notification.ID, &notification.Receiver, &notification.Message,
			&notification.ReceiverKind, &notification.Time)
		list = append(list, notification)

	}
	return list, err
}


func (r *postgresRepository) AppointmentsByParent(id int64, offset int, limit int) (appointments []models.Appointment, err error) {
	query := `select a.id, a.student, a.teacher, a.location, a.time
				from back2school.appointments as a natural join back2school.isparent 
				where parent = $1
				order by a.time desc
				limit $2 offset $3`
	rows, err := r.Query(query, id, limit, offset)
	defer rows.Close()
	if err != nil {
		log.Print(err)
	}
	for rows.Next() {
		appointment := models.Appointment{}
		rows.Scan(&appointment.ID, &appointment.Student, &appointment.Teacher, &appointment.Location, &appointment.Time)
		appointments = append(appointments, appointment)

	}
	return appointments, err
}
