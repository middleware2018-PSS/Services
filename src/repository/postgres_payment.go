package repository

import (
	"database/sql"
	"github.com/middleware2018-PSS/Services/src/models"
)

func (r *postgresRepository) PaymentByID(id int64) (interface{}, error) {
	payment := &models.Payment{}
	err := r.QueryRow(`SELECT id, amount, payed, emitted, reason
								FROM back2school.payments WHERE id = $1 `, id).Scan(payment.ID, payment.Amount, payment.Payed, payment.Emitted, payment.Reason)
	return payment, switchError(err)
}

func (r *postgresRepository) Payments(limit int, offset int) ([]interface{}, error) {
	return r.listByParams(`select id, amount, student, payed, reason, emitted
						from back2school.payments
						order by payed asc, emitted asc`,
		func(rows *sql.Rows) (interface{}, error) {
			payment := models.Payment{}
			err := rows.Scan(&payment.ID, &payment.Amount, &payment.Student.ID, &payment.Payed, &payment.Reason, &payment.Emitted)
			return payment, err
		}, limit, offset)
}
