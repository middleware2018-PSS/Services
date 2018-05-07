package repository

import (
	"github.com/middleware2018-PSS/Services/src/models"
	"log"
)

func (r *postgresRepository) PaymentByID(id int64) (payment models.Payment) {
	err := r.QueryRow(`SELECT id, amount, payed, emitted, reason
								FROM back2school.payments WHERE id = $1 `, id).Scan(payment.ID, payment.Amount, payment.Payed, payment.Emitted, payment.Reason)
	if err != nil {
		log.Print(err)
	}
	return payment
}
