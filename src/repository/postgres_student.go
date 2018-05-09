package repository

import (
	"database/sql"
	"github.com/middleware2018-PSS/Services/src/models"
)

func (r *postgresRepository) StudentById(id int64) (student interface{}, err error) {
	s := models.Student{}
	err = r.QueryRow(`SELECT id,	name, surname, mail, info 
								FROM back2school.students WHERE id = $1`, id).Scan(&s.ID,
		&s.Name, &s.Surname, &s.Mail, &s.Info)
	return s, switchError(err)
}

func (r *postgresRepository) GradesByStudent(id int64, offset int, limit int) ([]interface{}, error) {
	return r.listByID(id, offset, limit, `SELECT student, subject, date, grade, teacher
									FROM back2school.grades
									WHERE student = $1
									order by date desc
									limit $2 offset $3
									`,
		func(rows *sql.Rows) (interface{}, error) {
			grade := models.Grade{}
			err := rows.Scan(&grade.Student.ID, &grade.Subject, &grade.Date, &grade.Grade, &grade.Teacher.ID)
			return grade, err
		})
}

func (r *postgresRepository) GradeStudent(grade models.Grade) (err error) {
	// TODO
	return
}

func (r *postgresRepository) UpdateStudent(id int64) (err error) {
	// TODO
	return
}
