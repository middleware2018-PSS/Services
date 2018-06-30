package repository

// @Summary Delete Account
// @Param id path int true "Account ID"
// @Tags Accounts
// @Router /accounts/{username} [delete]
// @Security ApiKeyAuth
func (r *postgresRepository) DeleteAccount(username string, whoKind string) (interface{}, error) {
	if whoKind == AdminUser {
		query := "DELETE FROM back2school.accounts where username = $1"
		return r.exec(query, username)
	} else {
		return nil, ErrorNotAuthorized
	}
}

// @Summary Delete Parent
// @Param id path int true "Parent ID"
// @Tags Parents
// @Router /parents/{id} [delete]
// @Security ApiKeyAuth
func (r *postgresRepository) DeleteParent(id int, who int, whoKind string) (interface{}, error) {
	if whoKind == AdminUser {
		query := "DELETE FROM back2school.parents where id = $1"
		return r.exec(query, id)
	} else {
		return nil, ErrorNotAuthorized
	}
}

// @Summary Delete Teacher
// @Param id path int true "Teacher ID"
// @Tags Teachers
// @Router /teachers/{id} [delete]
// @Security ApiKeyAuth
func (r *postgresRepository) DeleteTeacher(id int, who int, whoKind string) (interface{}, error) {
	if whoKind == AdminUser {
		query := "DELETE FROM back2school.teachers where id = $1"
		return r.exec(query, id)
	} else {
		return nil, ErrorNotAuthorized
	}
}

// @Summary Delete Appointment
// @Param id path int true "Appointment ID"
// @Tags Appointments
// @Router /appointments/{id} [delete]
// @Security ApiKeyAuth
func (r *postgresRepository) DeleteAppointment(id int, who int, whoKind string) (interface{}, error) {
	if whoKind == AdminUser {
		query := "DELETE FROM back2school.appointments where id = $1"
		return r.exec(query, id)
	} else {
		return nil, ErrorNotAuthorized
	}
}

// @Summary Delete Student
// @Param id path int true "Student ID"
// @Tags Students
// @Router /students/{id} [delete]
// @Security ApiKeyAuth
func (r *postgresRepository) DeleteStudent(id int, who int, whoKind string) (interface{}, error) {
	if whoKind == AdminUser {
		query := "DELETE FROM back2school.students where id = $1"
		return r.exec(query, id)
	} else {
		return nil, ErrorNotAuthorized
	}
}

// @Summary Delete Notification
// @Param id path int true "Notification ID"
// @Tags Notifications
// @Router /notifications/{id} [delete]
// @Security ApiKeyAuth
func (r *postgresRepository) DeleteNotification(id int, who int, whoKind string) (interface{}, error) {
	if whoKind == AdminUser {
		query := "DELETE FROM back2school.notification where id = $1"
		return r.exec(query, id)
	} else {
		return nil, ErrorNotAuthorized
	}
}

// @Summary Delete Payment
// @Param id path int true "Payment ID"
// @Tags Payments
// @Router /payments/{id} [delete]
// @Security ApiKeyAuth
func (r *postgresRepository) DeletePayment(id int, who int, whoKind string) (interface{}, error) {
	if whoKind == AdminUser {
		query := "DELETE FROM back2school.payments where id = $1"
		return r.exec(query, id)
	} else {
		return nil, ErrorNotAuthorized
	}
}

// @Summary Delete Class
// @Param id path int true "Class ID"
// @Tags Classes
// @Router /classes/{id} [delete]
// @Security ApiKeyAuth
func (r *postgresRepository) DeleteClass(id int, who int, whoKind string) (interface{}, error) {
	if whoKind == AdminUser {
		query := "DELETE FROM back2school.classes where id = $1"
		return r.exec(query, id)
	} else {
		return nil, ErrorNotAuthorized
	}
}

// @Summary Delete Grade
// @Param id path int true "Grade ID"
// @Tags Grades
// @Router /grades/{id} [delete]
// @Security ApiKeyAuth
func (r *postgresRepository) DeleteGrade(id int, who int, whoKind string) (interface{}, error) {
	if whoKind == AdminUser {
		query := "DELETE FROM back2school.grades where id = $1"
		return r.exec(query, id)
	} else {
		return nil, ErrorNotAuthorized
	}
}

// @Summary Delete Lecture
// @Param id path int true "Lecture ID"
// @Tags Lectures
// @Router /lectures/{id} [delete]
// @Security ApiKeyAuth
func (r *postgresRepository) DeleteLecture(id int, who int, whoKind string) (interface{}, error) {
	if whoKind == AdminUser {
		query := "DELETE FROM back2school.timetable where id = $1"
		return r.exec(query, id)
	} else {
		return nil, ErrorNotAuthorized
	}
}
