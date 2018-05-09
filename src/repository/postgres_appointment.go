package repository

import (
	"github.com/middleware2018-PSS/Services/src/models"
	"log"
)

func (r *postgresRepository) UpdateAppointments(id int64) (err error) {
	//TODO
	return
}

func (r *postgresRepository) AppointmentById(id int64) (appointment models.Appointment, err error) {
	err = r.QueryRow(`SELECT id, student, teacher, time, location 
								FROM back2school.appointments WHERE id = $1 `, id).Scan(
		&appointment.ID, &appointment.Student.ID, &appointment.Teacher.ID, &appointment.Time, &appointment.Location)
	if err != nil {
		log.Print(err)
	}
	return appointment, err
}
