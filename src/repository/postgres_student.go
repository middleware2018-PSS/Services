package repository

import (
	"database/sql"
	"github.com/middleware2018-PSS/Services/src/models"
)

func (r *postgresRepository) Students(limit int, offset int) (student []interface{}, err error) {
	return r.listByParams(`select id, name, surname, mail, info 
						from back2school.students
						order by name desc, surname desc`, func(rows *sql.Rows) (interface{}, error) {
		student := models.Student{}
		err := rows.Scan(&student.ID, &student.Name, &student.Surname, &student.Mail, &student.Info)
		return student, err
	}, limit, offset)
}

func (r *postgresRepository) StudentById(id int64) (student interface{}, err error) {
	s := models.Student{}
	err = r.QueryRow(`SELECT id,	name, surname, mail, info 
								FROM back2school.students WHERE id = $1`, id).Scan(&s.ID,
		&s.Name, &s.Surname, &s.Mail, &s.Info)
	return s, switchError(err)
}

func (r *postgresRepository) GradesByStudent(id int64, limit int, offset int) ([]interface{}, error) {
	return r.listByParams(`SELECT student, subject, date, grade, teacher
									FROM back2school.grades
									WHERE student = $1
									order by date desc`,
		func(rows *sql.Rows) (interface{}, error) {
			grade := models.Grade{}
			err := rows.Scan(&grade.Student.ID, &grade.Subject, &grade.Date, &grade.Grade, &grade.Teacher.ID)
			return grade, err
		}, limit, offset, id)
}

func (r *postgresRepository) GradeStudent(grade models.Grade) (err error) {
	// TODO
	return
}

func (r *postgresRepository) UpdateStudent(id int64) (err error) {
	// TODO
	return
}
