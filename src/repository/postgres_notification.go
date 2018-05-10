package repository

import (
	"database/sql"
	"github.com/middleware2018-PSS/Services/src/models"
)

func (r *postgresRepository) NotificationByID(id int64) (interface{}, error) {
	n := models.Notification{}
	err := r.QueryRow(`SELECT id, receiver, message, time, receiver_kind
								FROM back2school.notification WHERE id = $1 `, id).Scan(&n.ID,
		&n.Receiver, &n.Message, &n.Time, &n.ReceiverKind)
	return n, switchError(err)
}

func (r *postgresRepository) Notifications(limit int, offset int) ([]interface{}, error) {
	return r.listByParams(`select id, receiver, message, time, receiver_kind 
						from back2school.notification
						order by time desc, receiver_kind desc`,
		func(rows *sql.Rows) (interface{}, error) {
			notification := models.Notification{}
			err := rows.Scan(&notification.ID, &notification.Receiver, &notification.Message, &notification.Time, &notification.ReceiverKind)
			return notification, err
		}, limit, offset)
}
