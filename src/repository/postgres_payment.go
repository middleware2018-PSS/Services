package repository

import (
	"github.com/middleware2018-PSS/Services/src/models"
	"log"
)

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
