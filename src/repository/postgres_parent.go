package repository

import (
	"github.com/middleware2018-PSS/Services/src/models"
	"log"
)

func (r *postgresRepository) ParentById(id int64) (parent *models.Parent) {

	// get basic data
	parent = &models.Parent{}

	err := r.QueryRow(`SELECT id,	name, surname, mail 
								FROM back2school.parents WHERE id = $1`, id).Scan(&parent.ID, &parent.Name, &parent.Surname, &parent.Mail)
	if err != nil {
		log.Print(err)
	}

	// Children

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
		log.Print(err)
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
		log.Print(err)
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
		log.Print(err)
	}
	for rows.Next() {
		notification := models.Notification{}
		rows.Scan(&notification.ID)
		list = append(list, notification)

	}
	return list
}
