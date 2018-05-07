package repository

import (
	"github.com/middleware2018-PSS/Services/src/models"
	"log"
)

func (r *postgresRepository) ClassesByID(id int64) (class models.Class) {
	err := r.QueryRow(`SELECT id, year, section, info, grade FROM back2school.classes 
								WHERE id = $1`, id).Scan(&class.ID, &class.Year, &class.Section, &class.Info, &class.Grade)
	if err != nil {
		log.Print(err)
	}
	return class
}

func (r *postgresRepository) StudentByClass(id int64, offset int, limit int) (students []models.Student){
	rows, err := r.Query(`select id, name, surname, mail, info 
						from back2school.student join back2school.enrolled on student = id 
						where class = $1
						order by name desc, surname desc
						limit $2 offset $3`, id, limit, offset)
	defer rows.Close()
	if err != nil {
		log.Print(err)
	}
	for rows.Next() {
		student := models.Student{}
		rows.Scan(&student.ID, &student.Name, &student.Surname, &student.Mail, &student.Info)
		students = append(students,student)

	}
	return students
}

func (r *postgresRepository) LectureByClass(id int64, offset int, limit int) (lectures []models.TimeTable){
	rows, err := r.Query(`select class, subject, start, end, location, info
						from back2school.timetable natural join teaches
						where teacher = $1
						order by start desc
						limit $2 offset $3`, id, limit, offset)
	defer rows.Close()
	if err != nil {
		log.Print(err)
	}
	for rows.Next() {
		lecture := models.TimeTable{}
		rows.Scan(&lecture.Class, &lecture.Subject, &lecture.Start, &lecture.End, &lecture.Location, &lecture.Info)
		lectures = append(lectures, lecture)
	}
	return lectures

}
