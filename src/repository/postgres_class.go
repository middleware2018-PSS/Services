package repository

import (
	"database/sql"
	"github.com/middleware2018-PSS/Services/src/models"
)

func (r *postgresRepository) ClassesByID(id int64) (class models.Class, err error) {
	err = r.QueryRow(`SELECT id, year, section, info, grade FROM back2school.classes 
								WHERE id = $1`, id).Scan(&class.ID, &class.Year, &class.Section, &class.Info, &class.Grade)
	return class, err
}

func (r *postgresRepository) StudentByClass(id int64, offset int, limit int) (students interface{}, err error) {
	return r.listBySMTH(id, offset, limit,`select id, name, surname, mail, info 
						from back2school.student join back2school.enrolled on student = id 
						where class = $1
						order by name desc, surname desc
						limit $2 offset $3`,
		func(rows *sql.Rows) interface{} {
			student := models.Student{}
			rows.Scan(&student.ID, &student.Name, &student.Surname, &student.Mail, &student.Info)
			return student
		})
}

func (r *postgresRepository) LectureByClass(id int64, offset int, limit int) ([]interface{}, error) {
	return r.listBySMTH(id, offset, limit,
		`select id, class, subject, "start", "end", location, info
				from back2school.timetable natural join back2school.teaches
				where teacher = $1
				order by "start" desc
				limit $2 offset $3`,
		func(rows *sql.Rows) interface{} {
			lecture := models.TimeTable{}
			rows.Scan(&lecture.ID, &lecture.Class, &lecture.Subject, &lecture.Start, &lecture.End, &lecture.Location, &lecture.Info)
			return lecture
		})
}
