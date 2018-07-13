package repository

// @Summary Delete Account
// @Param id path int true "Account ID"
// @Tags Accounts
// @Router /accounts/{username} [delete]
// @Security ApiKeyAuth
func (r *Repository) DeleteAccount(username string) (interface{}, error) {
	query := "DELETE FROM back2school.accounts where username = $1"
	return r.exec(query, username)

}

// @Summary Delete Parent
// @Param id path int true "Parent ID"
// @Tags Parents
// @Router /parents/{id} [delete]
// @Security ApiKeyAuth
func (r *Repository) DeleteParent(id int) (interface{}, error) {
	query := "DELETE FROM back2school.parents where id = $1"
	return r.exec(query, id)
}

// @Summary Delete Teacher
// @Param id path int true "Teacher ID"
// @Tags Teachers
// @Router /teachers/{id} [delete]
// @Security ApiKeyAuth
func (r *Repository) DeleteTeacher(id int) (interface{}, error) {
	query := "DELETE FROM back2school.teachers where id = $1"
	return r.exec(query, id)
}

// @Summary Delete Appointment
// @Param id path int true "Appointment ID"
// @Tags Appointments
// @Router /appointments/{id} [delete]
// @Security ApiKeyAuth
func (r *Repository) DeleteAppointment(id int) (interface{}, error) {
	query := "DELETE FROM back2school.appointments where id = $1"
	return r.exec(query, id)
}

// @Summary Delete Student
// @Param id path int true "Student ID"
// @Tags Students
// @Router /students/{id} [delete]
// @Security ApiKeyAuth
func (r *Repository) DeleteStudent(id int) (interface{}, error) {
	query := "DELETE FROM back2school.students where id = $1"
	return r.exec(query, id)
}

// @Summary Delete Notification
// @Param id path int true "Notification ID"
// @Tags Notifications
// @Router /notifications/{id} [delete]
// @Security ApiKeyAuth
func (r *Repository) DeleteNotification(id int) (interface{}, error) {
	query := "DELETE FROM back2school.notification where id = $1"
	return r.exec(query, id)
}

// @Summary Delete Payment
// @Param id path int true "Payment ID"
// @Tags Payments
// @Router /payments/{id} [delete]
// @Security ApiKeyAuth
func (r *Repository) DeletePayment(id int) (interface{}, error) {
	query := "DELETE FROM back2school.payments where id = $1"
	return r.exec(query, id)
}

// @Summary Delete Class
// @Param id path int true "Class ID"
// @Tags Classes
// @Router /classes/{id} [delete]
// @Security ApiKeyAuth
func (r *Repository) DeleteClass(id int) (interface{}, error) {
	query := "DELETE FROM back2school.classes where id = $1"
	return r.exec(query, id)

}

// @Summary Delete Grade
// @Param id path int true "Grade ID"
// @Tags Grades
// @Router /grades/{id} [delete]
// @Security ApiKeyAuth
func (r *Repository) DeleteGrade(id int) (interface{}, error) {
	query := "DELETE FROM back2school.grades where id = $1"
	return r.exec(query, id)
}

// @Summary Delete Lecture
// @Param id path int true "Lecture ID"
// @Tags Lectures
// @Router /lectures/{id} [delete]
// @Security ApiKeyAuth
func (r *Repository) DeleteLecture(id int) (interface{}, error) {
	query := "DELETE FROM back2school.timetable where id = $1"
	return r.exec(query, id)
}
