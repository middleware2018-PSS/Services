package repository

import "github.com/middleware2018-PSS/Services/src/models"

// @Summary Get a appointment by id
// @Param id path int true "Appointment ID"
// @Tags Appointments
// @Success 200 {object} models.Appointment
// @Router /appointments/{id} [get]
func (r *postgresRepository) AppointmentByID(id int, who int, whoKind string) (interface{}, error) {
	appointment := models.Appointment{}
	var args []interface{}
	var query string
	switch whoKind {
	case ParentUser:
		query = "SELECT id, student, teacher, time, location " +
			"FROM back2school.appointments natural join back2school.isParent WHERE id = $1 and parent = $2"
		args = append(args, id, who)
	case TeacherUser:
		query = "SELECT id, student, teacher, time, location " +
			"FROM back2school.appointments WHERE id = $1 and teacher = $2 "
		args = append(args, id, who)
	case AdminUser:
		query = "SELECT id, student, teacher, time, location " +
			"FROM back2school.appointments WHERE id = $1 "
		args = append(args, id)
	default:
		return nil, ErrorNotAuthorized
	}
	err := r.QueryRow(query, args...).Scan(
		&appointment.ID, &appointment.Student.ID, &appointment.Teacher.ID, &appointment.Time, &appointment.Location)
	return switchResult(appointment, err)
}


// @Summary Get a grade by id
// @Param id path int true "Grade ID"
// @Tags Grades
// @Success 200 {object} models.Grade
// @Router /grades/{id} [get]
func (r *postgresRepository) GradeByID(id int, who int, whoKind string) (interface{}, error) {
	grade := models.Grade{}
	var query string
	var args []interface{}
	switch whoKind {
	case ParentUser:
		query = "SELECT id, student, teacher, subject, date, grade " +
			"FROM back2school.grades natural join back2school.isParent WHERE id = $1 and parent = $2 "
		args = append(args, id, who)
	case TeacherUser:
		query = "SELECT id, student, teacher, subject, date, grade " +
			"FROM back2school.grades WHERE id = $1 and teacher = $2 "
		args = append(args, id, who)
	case AdminUser:
		query = "SELECT id, student, teacher, subject, date, grade " +
			" FROM back2school.grades WHERE id = $1 "
		args = append(args, id)
	default:
		return nil, ErrorNotAuthorized
	}
	err := r.QueryRow(query, id, who).Scan(
		&grade.ID, &grade.Student.ID, &grade.Teacher.ID, &grade.Subject, &grade.Date, &grade.Grade)
	return switchResult(grade, err)
}

// @Summary Get a class by id
// @Param id path int true "Class ID"
// @Tags Classes
// @Success 200 {object} models.Class
// @Router /classes/{id} [get]
func (r *postgresRepository) ClassByID(id int, who int, whoKind string) (interface{}, error) {
	class := models.Class{}
	var query string
	var args []interface{}
	switch whoKind {
	case TeacherUser:
		query = "SELECT id, year, section, info, grade FROM back2school.classes join back2school.teaches on class = id" +
			"WHERE id = $1 and teacher = $2 "
		args = append(args, id, who)
	case AdminUser:
		query = "SELECT id, year, section, info, grade FROM back2school.classes " +
			"WHERE id = $1 "
		args = append(args, id)
	default:
		return nil, ErrorNotAuthorized
	}
	err := r.QueryRow(query, args...).Scan(&class.ID, &class.Year, &class.Section, &class.Info, &class.Grade)
	return switchResult(class, err)
}


// @Summary Get a notification by id
// @Param id path int true "Notification ID"
// @Tags Notifications
// @Success 200 {object} models.Notification
// @Router /notifications/{id} [get]
func (r *postgresRepository) NotificationByID(id int, who int, whoKind string) (interface{}, error) {
	n := models.Notification{}
	var query string
	var args []interface{}
	switch whoKind {
	case TeacherUser:
		query = "SELECT id, receiver, message, time, receiver_kind " +
			"FROM back2school.notification " +
			"WHERE id = $1 and receiver = $2 and receiver_kind = $3 "
		args = append(args, id, who, whoKind)
	case ParentUser:
		query = "SELECT id, receiver, message, time, receiver_kind " +
			"FROM back2school.notification WHERE id = $1 and receiver = $2 and receiver_kind = $3 "
		args = append(args, id, who, whoKind)
	case AdminUser:
		query = "SELECT id, receiver, message, time, receiver_kind " +
			"FROM back2school.notification WHERE id = $1 "
		args = append(args, id)

	default:
		return nil, ErrorNotAuthorized
	}
	err := r.QueryRow(query, args...).Scan(&n.ID,
		&n.Receiver, &n.Message, &n.Time, &n.ReceiverKind)
	return switchResult(n, err)
}


// Parents
// see/modify their personal data
// @Summary Get a parent by id
// @Param id path int true "Account ID"
// @Tags Parents
// @Success 200 {object} models.Parent
// @Router /parents/{id} [get]
func (r *postgresRepository) ParentByID(id int, who int, whoKind string) (interface{}, error) {
	p := models.Parent{}
	var query string
	var args []interface{}
	switch whoKind {
	case ParentUser:
		if id == who {
			query = "SELECT id,	name, surname, mail, info " +
				"FROM back2school.parents WHERE id = $1 "
			args = append(args, who)
		} else {
			return nil, ErrorNotAuthorized
		}
	case AdminUser:
		query = "SELECT id,	name, surname, mail, info " +
			"FROM back2school.parents WHERE id = $1 "
		args = append(args, id)
	default:
		return nil, ErrorNotAuthorized
	}
	err := r.QueryRow(query,
		args...).Scan(&p.ID, &p.Name, &p.Surname, &p.Mail, &p.Info)
	return switchResult(p, err)
}


// Get student by id
// @Summary Get a student by id
// @Param id path int true "Student ID"
// @Tags Students
// @Success 200 {object} models.Student
// @Router /students/{id} [get]
func (r *postgresRepository) StudentByID(id int, who int, whoKind string) (student interface{}, err error) {
	s := models.Student{}
	var query string
	var args []interface{}
	switch whoKind {
	case ParentUser:
		query = "SELECT id,	name, surname, mail, info  " +
			"FROM back2school.students join back2school.isParent on student = id " +
			"WHERE id = $1 and parent = $2 "
		args = append(args, id, who)
	case AdminUser:
		query = "SELECT id,	name, surname, mail, info  " +
			"FROM back2school.students WHERE id = $1 "
		args = append(args, id)

	default:
		return nil, ErrorNotAuthorized
	}
	err = r.QueryRow(query, args...).Scan(&s.ID,
		&s.Name, &s.Surname, &s.Mail, &s.Info)
	return switchResult(s, err)
}

// @Summary Get a lecture by id
// @Param id path int true "Lecture ID"
// @Tags Grades
// @Success 200 {object} models.TimeTable
// @Router /lectures/{id} [get]
func (r *postgresRepository) LectureByID(id int, who int, whoKind string) (interface{}, error) {
	grade := models.Grade{}
	var query string
	var args []interface{}
	switch whoKind {
	case ParentUser:
		query = "select id, class, subject, \"start\", \"end\", location, info " +
			"from back2school.timetable natural join back2school.enrolled natural join back2school.isParent " +
			"where id = $1 and parent = $2 " +
			"order by \"start\" desc "
		args = append(args, id, who)
	case TeacherUser:
		query = "select id, class, subject, \"start\", \"end\", location, info " +
			"from back2school.timetable natural join back2school.teaches " +
			"where id = $1 and teacher = $2 " +
			"order by \"start\" desc "
		args = append(args, id, who)
	case AdminUser:
		query = "select id, class, subject, \"start\", \"end\", location, info " +
			"from back2school.timetable " +
			"where id = $1 " +
			"order by \"start\" desc "
		args = append(args, id)

	default:
		return nil, ErrorNotAuthorized
	}
	err := r.QueryRow(query, id, who).Scan(
		&grade.ID, &grade.Student.ID, &grade.Teacher.ID, &grade.Subject, &grade.Date, &grade.Grade)
	return switchResult(grade, err)
}
