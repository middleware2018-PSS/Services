package repository

import (
	"github.com/middleware2018-PSS/Services/src/models"
	"log"
)

func (r *postgresRepository) NotificationByID(id int64) (notifications []models.Notification) {
	rows, err := r.Query(`SELECT id, receiver, message, time, receiver_kind
								FROM back2school.notification WHERE id = $1 `, id)
	defer rows.Close()
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		notification := &models.Notification{}
		rows.Scan(notification.ID, notification.Receiver, notification.Message, notification.Time, notification.ReceiverKind)
		notifications = append(notifications, *notification)
	}
	return notifications
}
