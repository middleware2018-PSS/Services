package controller

// @Summary Delete Account
// @Param id path int true "Account ID"
// @Tags Accounts
// @Router /accounts/{username} [delete]
// @Security ApiKeyAuth
func (r Controller) DeleteAccount(username string, whoKind string) (interface{}, error) {
	if whoKind == AdminUser {
		r.repo.DeleteAccount(username)
	} else {
		return nil, ErrorNotAuthorized
	}
}

// @Summary Delete Parent
// @Param id path int true "Parent ID"
// @Tags Parents
// @Router /parents/{id} [delete]
// @Security ApiKeyAuth
func (r Controller) DeleteParent(id int, who int, whoKind string) (interface{}, error) {
	if whoKind == AdminUser {
		r.repo.DeleteParent(id)
	} else {
		return nil, ErrorNotAuthorized
	}
}

// @Summary Delete Teacher
// @Param id path int true "Teacher ID"
// @Tags Teachers
// @Router /teachers/{id} [delete]
// @Security ApiKeyAuth
func (r Controller) DeleteTeacher(id int, who int, whoKind string) (interface{}, error) {
	if whoKind == AdminUser {
		r.repo.DeleteTeacher(id)
	} else {
		return nil, ErrorNotAuthorized
	}
}

// @Summary Delete Appointment
// @Param id path int true "Appointment ID"
// @Tags Appointments
// @Router /appointments/{id} [delete]
// @Security ApiKeyAuth
func (r Controller) DeleteAppointment(id int, who int, whoKind string) (interface{}, error) {
	if whoKind == AdminUser {
		r.repo.DeleteAppointment(id)
	} else {
		return nil, ErrorNotAuthorized
	}
}

// @Summary Delete Student
// @Param id path int true "Student ID"
// @Tags Students
// @Router /students/{id} [delete]
// @Security ApiKeyAuth
func (r Controller) DeleteStudent(id int, who int, whoKind string) (interface{}, error) {
	if whoKind == AdminUser {
		r.repo.DeleteStudent(id)
	} else {
		return nil, ErrorNotAuthorized
	}
}

// @Summary Delete Notification
// @Param id path int true "Notification ID"
// @Tags Notifications
// @Router /notifications/{id} [delete]
// @Security ApiKeyAuth
func (r Controller) DeleteNotification(id int, who int, whoKind string) (interface{}, error) {
	if whoKind == AdminUser {
		r.repo.DeleteNotification(id)
	} else {
		return nil, ErrorNotAuthorized
	}
}

// @Summary Delete Payment
// @Param id path int true "Payment ID"
// @Tags Payments
// @Router /payments/{id} [delete]
// @Security ApiKeyAuth
func (r Controller) DeletePayment(id int, who int, whoKind string) (interface{}, error) {
	if whoKind == AdminUser {
		r.repo.DeletePayment(id)
	} else {
		return nil, ErrorNotAuthorized
	}
}

// @Summary Delete Class
// @Param id path int true "Class ID"
// @Tags Classes
// @Router /classes/{id} [delete]
// @Security ApiKeyAuth
func (r Controller) DeleteClass(id int, who int, whoKind string) (interface{}, error) {
	if whoKind == AdminUser {
		r.repo.DeleteClass(id)
	} else {
		return nil, ErrorNotAuthorized
	}
}

// @Summary Delete Grade
// @Param id path int true "Grade ID"
// @Tags Grades
// @Router /grades/{id} [delete]
// @Security ApiKeyAuth
func (r Controller) DeleteGrade(id int, who int, whoKind string) (interface{}, error) {
	if whoKind == AdminUser {
		r.repo.DeleteGrade(id)
	} else {
		return nil, ErrorNotAuthorized
	}
}

// @Summary Delete Lecture
// @Param id path int true "Lecture ID"
// @Tags Lectures
// @Router /lectures/{id} [delete]
// @Security ApiKeyAuth
func (r Controller) DeleteLecture(id int, who int, whoKind string) (interface{}, error) {
	if whoKind == AdminUser {
		r.repo.DeleteLecture(id)
	} else {
		return nil, ErrorNotAuthorized
	}
}
