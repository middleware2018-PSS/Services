package repository

import (
	"database/sql"
	"github.com/middleware2018-PSS/Services/src/models"
)

func (r *postgresRepository) ClassByID(id int64) (interface{}, error) {
	class := models.Class{}
	err := r.QueryRow(`SELECT id, year, section, info, grade FROM back2school.classes 
								WHERE id = $1`, id).Scan(&class.ID, &class.Year, &class.Section, &class.Info, &class.Grade)
	return class, switchError(err)
}

func (r *postgresRepository) Classes(limit int, offset int) ([]interface{}, error) {
	return r.listByParams(`select id, year, section, info, grade
						from back2school.classes
						order by year desc, grade asc, section asc`, func(rows *sql.Rows) (interface{}, error) {
		class := models.Class{}
		err := rows.Scan(&class.ID, &class.Year, &class.Section, &class.Info, &class.Grade)
		return class, err
	}, limit, offset)
}

func (r *postgresRepository) StudentByClass(id int64, limit int, offset int) (students []interface{}, err error) {
	return r.listByParams(`select id, name, surname, mail, info 
						from back2school.students join back2school.enrolled on student = id 
						where class = $1
						order by name desc, surname desc`,
		func(rows *sql.Rows) (interface{}, error) {
			student := models.Student{}
			err := rows.Scan(&student.ID, &student.Name, &student.Surname, &student.Mail, &student.Info)
			return student, err
		}, limit, offset, id)
}

func (r *postgresRepository) LectureByClass(id int64, limit int, offset int) ([]interface{}, error) {
	return r.listByParams(
		`select id, class, subject, "start", "end", location, info
				from back2school.timetable natural join back2school.teaches
				where teacher = $1
				order by "start" desc`,
		func(rows *sql.Rows) (interface{}, error) {
			lecture := models.TimeTable{}
			err := rows.Scan(&lecture.ID, &lecture.Class, &lecture.Subject, &lecture.Start, &lecture.End, &lecture.Location, &lecture.Info)
			return lecture, err
		}, limit, offset, id)
}
