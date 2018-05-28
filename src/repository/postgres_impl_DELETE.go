package repository

func (r *postgresRepository) DeleteAccount(username string, whoKind string) (interface{}, error) {
	if whoKind == AdminUser {
		query := "DELETE FROM back2school.accounts where user = $1"
		return r.exec(query, username)
	} else {
		return nil, ErrorNotAuthorized
	}
}

func (r *postgresRepository) DeleteParent(id int, who int, whoKind string) (interface{}, error) {
	if whoKind == AdminUser {
		query := "DELETE FROM back2school.parents where id = $1"
		return r.exec(query, id)
	} else {
		return nil, ErrorNotAuthorized
	}
}
func (r *postgresRepository) DeleteTeacher(id int, who int, whoKind string) (interface{}, error) {
	if whoKind == AdminUser {
		query := "DELETE FROM back2school.teachers where id = $1"
		return r.exec(query, id)
	} else {
		return nil, ErrorNotAuthorized
	}
}
func (r *postgresRepository) DeleteAppointment(id int, who int, whoKind string) (interface{}, error) {
	if whoKind == AdminUser {
		query := "DELETE FROM back2school.appointments where id = $1"
		return r.exec(query, id)
	} else {
		return nil, ErrorNotAuthorized
	}
}
func (r *postgresRepository) DeleteStudent(id int, who int, whoKind string) (interface{}, error) {
	if whoKind == AdminUser {
		query := "DELETE FROM back2school.students where id = $1"
		return r.exec(query, id)
	} else {
		return nil, ErrorNotAuthorized
	}
}
func (r *postgresRepository) DeleteNotification(id int, who int, whoKind string) (interface{}, error) {
	if whoKind == AdminUser {
		query := "DELETE FROM back2school.notification where id = $1"
		return r.exec(query, id)
	} else {
		return nil, ErrorNotAuthorized
	}
}
func (r *postgresRepository) DeletePayment(id int, who int, whoKind string) (interface{}, error) {
	if whoKind == AdminUser {
		query := "DELETE FROM back2school.payments where id = $1"
		return r.exec(query, id)
	} else {
		return nil, ErrorNotAuthorized
	}
}
func (r *postgresRepository) DeleteClass(id int, who int, whoKind string) (interface{}, error) {
	if whoKind == AdminUser {
		query := "DELETE FROM back2school.classes where id = $1"
		return r.exec(query, id)
	} else {
		return nil, ErrorNotAuthorized
	}
}
