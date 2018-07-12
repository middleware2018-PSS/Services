package repository

import "github.com/middleware2018-PSS/Services/models"

func (r *postgresRepository) AppointmentForParent(id int, who int) (interface{}, error) {
	appointment := models.Appointment{}

	query := "SELECT id, student, teacher, time, location " +
		"FROM back2school.appointments natural join back2school.isParent WHERE id = $1 and parent = $2"

	err := r.QueryRow(query, id, who).Scan(
		&appointment.ID, &appointment.Student, &appointment.Teacher, &appointment.Time, &appointment.Location)
	return switchResult(appointment, err)
}
func (r *postgresRepository) AppointmentForTeacher(id int, who int) (interface{}, error) {
	appointment := models.Appointment{}

	query := "SELECT id, student, teacher, time, location " +
		"FROM back2school.appointments WHERE id = $1 and teacher = $2 "

	err := r.QueryRow(query, id, who).Scan(
		&appointment.ID, &appointment.Student, &appointment.Teacher, &appointment.Time, &appointment.Location)
	return switchResult(appointment, err)
}
func (r *postgresRepository) AppointmentForAdmin(id int) (interface{}, error) {
	appointment := models.Appointment{}

	query := "SELECT id, student, teacher, time, location " +
		"FROM back2school.appointments WHERE id = $1 "

	err := r.QueryRow(query, id).Scan(
		&appointment.ID, &appointment.Student, &appointment.Teacher, &appointment.Time, &appointment.Location)
	return switchResult(appointment, err)
}

func (r *postgresRepository) GradeByIDForParent(id int, who int) (interface{}, error) {
	grade := models.Grade{}

	query := "SELECT id, student, teacher, subject, date, grade " +
		"FROM back2school.grades natural join back2school.isParent WHERE id = $1 and parent = $2 "
	err := r.QueryRow(query, id, who).Scan(
		&grade.ID, &grade.Student, &grade.Teacher, &grade.Subject, &grade.Date, &grade.Grade)
	return switchResult(grade, err)
}
func (r *postgresRepository) GradeByIDForTeacher(id int, who int) (interface{}, error) {
	grade := models.Grade{}

	query := "SELECT id, student, teacher, subject, date, grade " +
		"FROM back2school.grades WHERE id = $1 and teacher = $2 "
	err := r.QueryRow(query, id, who).Scan(
		&grade.ID, &grade.Student, &grade.Teacher, &grade.Subject, &grade.Date, &grade.Grade)
	return switchResult(grade, err)
}
func (r *postgresRepository) GradeByIDForAdmin(id int) (interface{}, error) {
	grade := models.Grade{}

	query := "SELECT id, student, teacher, subject, date, grade " +
		" FROM back2school.grades WHERE id = $1 "
	err := r.QueryRow(query, id).Scan(
		&grade.ID, &grade.Student, &grade.Teacher, &grade.Subject, &grade.Date, &grade.Grade)
	return switchResult(grade, err)
}

func (r *postgresRepository) ClassByIDForTeacher(id int, who int) (interface{}, error) {
	class := models.Class{}

	query := "SELECT id, year, section, info, grade FROM back2school.classes join back2school.teaches on class = id" +
		"WHERE id = $1 and teacher = $2 "

	err := r.QueryRow(query, id, who).Scan(&class.ID, &class.Year, &class.Section, &class.Info, &class.Grade)
	return switchResult(class, err)
}
func (r *postgresRepository) ClassByIDForAdmin(id int) (interface{}, error) {
	class := models.Class{}

	query := "SELECT id, year, section, info, grade FROM back2school.classes " +
		"WHERE id = $1 "

	err := r.QueryRow(query, id).Scan(&class.ID, &class.Year, &class.Section, &class.Info, &class.Grade)
	return switchResult(class, err)
}

func (r *postgresRepository) NotificationByIDForTeacher(id int, who int, whoKind string) (interface{}, error) {
	n := models.Notification{}

	query := "SELECT id, receiver, message, time, receiver_kind " +
		"FROM back2school.notification " +
		"WHERE id = $1 and receiver = $2 and receiver_kind = $3 "

	err := r.QueryRow(query, id, who, whoKind).Scan(&n.ID,
		&n.Receiver, &n.Message, &n.Time, &n.ReceiverKind)
	return switchResult(n, err)
}
func (r *postgresRepository) NotificationByIDForParent(id int, who int, whoKind string) (interface{}, error) {
	n := models.Notification{}
	query := "SELECT id, receiver, message, time, receiver_kind " +
		"FROM back2school.notification WHERE id = $1 and receiver = $2 and receiver_kind = $3 "

	err := r.QueryRow(query, id, who, whoKind).Scan(&n.ID,
		&n.Receiver, &n.Message, &n.Time, &n.ReceiverKind)
	return switchResult(n, err)
}
func (r *postgresRepository) NotificationByIDForAdmin(id int, who int, whoKind string) (interface{}, error) {
	n := models.Notification{}
	query := "SELECT id, receiver, message, time, receiver_kind " +
		"FROM back2school.notification WHERE id = $1 "

	err := r.QueryRow(query, id).Scan(&n.ID,
		&n.Receiver, &n.Message, &n.Time, &n.ReceiverKind)
	return switchResult(n, err)
}

func (r *postgresRepository) ParentByID(id int) (interface{}, error) {
	p := models.Parent{}

	query := "SELECT id,	name, surname, mail, info " +
		"FROM back2school.parents WHERE id = $1 "

	err := r.QueryRow(query, id).Scan(&p.ID, &p.Name, &p.Surname, &p.Mail, &p.Info)
	return switchResult(p, err)
}

func (r *postgresRepository) StudentByIDForParent(id int, who int) (student interface{}, err error) {
	s := models.Student{}

	query := "SELECT id,	name, surname, mail, info  " +
		"FROM back2school.students join back2school.isParent on student = id " +
		"WHERE id = $1 and parent = $2 "

	err = r.QueryRow(query, id, who).Scan(&s.ID,
		&s.Name, &s.Surname, &s.Mail, &s.Info)
	return switchResult(s, err)
}
func (r *postgresRepository) StudentByIDForAdmin(id int) (student interface{}, err error) {
	s := models.Student{}

	query := "SELECT id,	name, surname, mail, info  " +
		"FROM back2school.students WHERE id = $1 "

	err = r.QueryRow(query, id).Scan(&s.ID,
		&s.Name, &s.Surname, &s.Mail, &s.Info)
	return switchResult(s, err)
}

func (r *postgresRepository) LectureByIDForParent(id int, who int) (interface{}, error) {
	lecture := models.TimeTable{}

	query := "select id, class, subject, \"start\", \"end\", location, info " +
		"from back2school.timetable natural join back2school.enrolled natural join back2school.isParent " +
		"where id = $1 and parent = $2 " +
		"order by \"start\" desc "

	err := r.QueryRow(query, id, who).Scan(&lecture.ID, lecture.Class, &lecture.Subject, &lecture.Location, &lecture.Start, &lecture.End, &lecture.Info)
	return switchResult(lecture, err)
}

func (r *postgresRepository) LectureByIDForTeacher(id int, who int) (interface{}, error) {
	lecture := models.TimeTable{}

	query := "select id, class, subject, \"start\", \"end\", location, info " +
		"from back2school.timetable natural join back2school.teaches " +
		"where id = $1 and teacher = $2 " +
		"order by \"start\" desc "

	err := r.QueryRow(query, id, who).Scan(&lecture.ID, lecture.Class, &lecture.Subject, &lecture.Location, &lecture.Start, &lecture.End, &lecture.Info)
	return switchResult(lecture, err)
}

func (r *postgresRepository) LectureByIDForAdmin(id int) (interface{}, error) {
	lecture := models.TimeTable{}

	query := "select id, class, subject, \"start\", \"end\", location, info " +
		"from back2school.timetable " +
		"where id = $1 " +
		"order by \"start\" desc "

	err := r.QueryRow(query, id).Scan(&lecture.ID, lecture.Class, &lecture.Subject, &lecture.Location, &lecture.Start, &lecture.End, &lecture.Info)
	return switchResult(lecture, err)
}
