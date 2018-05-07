package repository

import (
	"github.com/middleware2018-PSS/Services/src/models"
	"log"
)

func (r *postgresRepository) NotificationByID(id int64) (notification models.Notification) {
	err := r.QueryRow(`SELECT id, receiver, message, time, receiver_kind
								FROM back2school.notification WHERE id = $1 `, id).Scan(notification.ID,
									notification.Receiver, notification.Message, notification.Time, notification.ReceiverKind)
	if err != nil {
		log.Print(err)
	}
	return notification
}
