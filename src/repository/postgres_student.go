package repository

import (
	"github.com/middleware2018-PSS/Services/src/models"
	"log"
)

func (r *postgresRepository) StudentById(id int) (student *models.Student) {
	student = &models.Student{}
	err := r.QueryRow(`SELECT id,	name, surname, mail, info 
								FROM back2school.students WHERE id = $1`, id).Scan(&student.ID,
		&student.Name, &student.Surname, &student.Mail, &student.Info)
	if err != nil {
		log.Print(err)
	}
	return student
}

func (r *postgresRepository) GradesByStudent(id int64, offset int, limit int) (grades []models.Grade) {
	rows, err := r.Query(`SELECT student, subject, date, grade, teacher
									FROM back2school.grades
									WHERE student = $1
									order by date desc
									limit $2 offset $3
									`, id, limit, offset)
	if err != nil {
		log.Print(err)
	}
	defer rows.Close()
	for rows.Next() {
		grade := models.Grade{}
		rows.Scan(&grade.Student.ID, &grade.Subject, &grade.Date, &grade.Grade, &grade.Teacher.ID)
		grades = append(grades, grade)
	}
	return grades
}

func (r *postgresRepository) GradeStudent(grade models.Grade) {
	// TODO
}

func (r *postgresRepository) UpdateStudent(id int64) {
	// TODO
}

