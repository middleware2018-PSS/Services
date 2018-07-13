package controller

import (
	_ "github.com/middleware2018-PSS/Services/docs"
	"github.com/middleware2018-PSS/Services/repository"
)

type Subjects struct {
	Subjects []string `json:"subjects" example:"science"`
}

// @Summary Get all classes
// @Param limit query int false "number of elements to return"
// @Param offset query int false "offset in the list of elements to return"
// @Tags Classes
// @Success 200 {array} models.Class
// @Router /classes [get]
// @Security ApiKeyAuth
func (c Controller) Classes(limit int, offset int, who int, whoKind string) ([]interface{}, error) {
	switch whoKind {
	case repository.TeacherUser:
		return c.repo.ClassesForTeachers(limit, offset, who)
	case repository.AdminUser:
		return c.repo.ClassesForAdmins(limit, offset)
	default:
		return nil, ErrorNotAuthorized
	}
}

// @Summary Get a student by class
// @Param id path int true "Class ID"
// @Param limit query int false "number of elements to return"
// @Param offset query int false "offset in the list of elements to return"
// @Tags Classes
// @Success 200 {array} models.Student
// @Router /classes/{id}/students [get]
// @Security ApiKeyAuth
func (c Controller) StudentsByClass(id int, limit int, offset int, who int, whoKind string) (students []interface{}, err error) {
	switch whoKind {
	case repository.TeacherUser:
		return c.repo.StudentsByClassForTeachers(id, limit, offset, who)
	case repository.AdminUser:
		return c.repo.StudentsByClassForAdmins(id, limit, offset)
	default:
		return nil, ErrorNotAuthorized
	}

}

// @Summary Get lectures by class
// @Param limit query int false "number of elements to return"
// @Param offset query int false "offset in the list of elements to return"
// @Param id path int true "Class ID"
// @Tags Classes
// @Success 200 {array} models.TimeTable
// @Router /classes/{id}/lectures [get]
// @Security ApiKeyAuth
func (c Controller) LectureByClass(id int, limit int, offset int, who int, whoKind string) ([]interface{}, error) {

	switch whoKind {
	case repository.TeacherUser:
		return c.repo.LectureByClassForTeacherOrParents(id, limit, offset, who)
	case repository.AdminUser:
		return c.repo.LectureByClassForAdmins(id, limit, offset)
	default:
		return nil, ErrorNotAuthorized
	}
}

// List all notifications
// @Summary Get all notifications
// @Param limit query int false "number of elements to return"
// @Param offset query int false "offset in the list of elements to return"
// @Tags Notifications
// @Success 200 {array} models.Notification
// @Router /notifications [get]
// @Security ApiKeyAuth
func (c Controller) Notifications(limit int, offset int, who int, whoKind string) ([]interface{}, error) {

	switch whoKind {
	case repository.TeacherUser, repository.ParentUser:
		return c.repo.NotificationsForTeacherOrParents(limit, offset, who, whoKind)
	case repository.AdminUser:
		return c.repo.NotificationsForAdmins(limit, offset)
	default:
		return nil, ErrorNotAuthorized
	}
}

// @Summary Get all grades
// @Param limit query int false "number of elements to return"
// @Param offset query int false "offset in the list of elements to return"
// @Tags Grades
// @Success 200 {array} models.Grade
// @Router /grades [get]
// @Security ApiKeyAuth
func (c Controller) Grades(limit int, offset int, who int, whoKind string) ([]interface{}, error) {

	switch whoKind {
	case repository.ParentUser:
		return c.repo.GradesForParent(limit, offset, who)
	case repository.AdminUser:
		return c.repo.GradesForAdmins(limit, offset)
	default:
		return nil, ErrorNotAuthorized
	}
}

// @Summary Get all parents
// @Param limit query int false "number of elements to return"
// @Param offset query int false "offset in the list of elements to return"
// @Tags Parents
// @Success 200 {array} models.Parent
// @Router /parents [get]
// @Security ApiKeyAuth
func (c Controller) Parents(limit int, offset int, who int, whoKind string) ([]interface{}, error) {

	switch whoKind {
	case repository.ParentUser:
		return c.repo.ParentsForParents(who)
	case repository.AdminUser:
		return c.repo.ParentsForAdmins(limit, offset)
	default:
		return nil, ErrorNotAuthorized
	}
}

// see/modify the personal data of their registered children
// @Summary Get children of the parent
// @Param id path int true "Parent ID"
// @Param limit query int false "number of elements to return"
// @Param offset query int false "offset in the list of elements to return"
// @Tags Parents
// @Success 200 {array} models.Student
// @Router /parents/{id}/students [get]
// @Security ApiKeyAuth
func (c Controller) ChildrenByParent(id int, limit int, offset int, who int, whoKind string) (children []interface{}, err error) {
	switch whoKind {
	case repository.ParentUser:
		if id == who {
			return c.repo.ChildrenByParentForParent(who, limit, offset)
		} else {
			return nil, ErrorNotAuthorized
		}
	case repository.AdminUser:
		return c.repo.ChildrenByParentForAdmin(id, limit, offset)
	default:
		return nil, ErrorNotAuthorized
	}
}

// see the monthly payments that have been made to the school in the past
// @Summary Get payments of the parent
// @Param id path int true "Parent ID"
// @Param limit query int false "number of elements to return"
// @Param offset query int false "offset in the list of elements to return"
// @Tags Parents
// @Success 200 {array} models.Payment
// @Router /parents/{id}/payments [get]
// @Security ApiKeyAuth
func (c Controller) PaymentsByParent(id int, limit int, offset int, who int, whoKind string) (payments []interface{}, err error) {

	switch whoKind {
	case repository.ParentUser:
		if id == who {
			return c.repo.PaymentsByParentForParent(id, limit, offset)
		} else {
			return nil, ErrorNotAuthorized
		}
	case repository.AdminUser:
		return c.repo.PaymentsByParentForAdmin(id, limit, offset)

	default:
		return nil, ErrorNotAuthorized
	}
}

// see general/personal notifications coming from the school
// @Summary Get notifications of the parent
// @Param id path int true "Parent ID"
// @Param limit query int false "number of elements to return"
// @Param offset query int false "offset in the list of elements to return"
// @Tags Parents
// @Success 200 {array} models.Notification
// @Router /parents/{id}/notifications [get]
// @Security ApiKeyAuth
func (c Controller) NotificationsByParent(id int, limit int, offset int, who int, whoKind string) (list []interface{}, err error) {

	switch whoKind {
	case repository.ParentUser:
		if id == who {
			return c.repo.NotificationsByParentForParent(id, limit, offset)
		} else {
			return nil, ErrorNotAuthorized
		}
	case repository.AdminUser:
		return c.repo.NotificationsByParentForAdmins(id, limit, offset)

	default:
		return nil, ErrorNotAuthorized
	}
}

// see/modify appointments that they have with their children's teachers
// (calendar-like support for requesting appointments, err error)
// @Summary Get appointments of the parent
// @Param id path int true "Parent ID"
// @Tags Parents
// @Param limit query int false "number of elements to return"
// @Param offset query int false "offset in the list of elements to return"
// @Tags Parents
// @Success 200 {array} models.Appointment
// @Router /parents/{id}/appointments [get]
// @Security ApiKeyAuth
func (c Controller) AppointmentsByParent(id int, limit int, offset int, who int, whoKind string) (appointments []interface{}, err error) {

	switch whoKind {
	case repository.ParentUser:
		if id == who {
			return c.repo.AppointmentsByParentForParent(id, limit, offset, who)
		} else {
			return nil, ErrorNotAuthorized
		}
	case repository.AdminUser:
		return c.repo.AppointmentsByParentForAdmin(id, limit, offset)

	default:
		return nil, ErrorNotAuthorized
	}
}

// Get payment by id
// @Summary Get a payment by id
// @Param id path int true "Payment ID"
// @Tags Payments
// @Success 200 {object} models.Payment
// @Router /payments/{id} [get]
// @Security ApiKeyAuth
func (c Controller) PaymentByID(id int, who int, whoKind string) (interface{}, error) {
	switch whoKind {
	case repository.ParentUser:
		return c.repo.PaymentByIDForParent(id, who)
	case repository.AdminUser:
		return c.repo.PaymentByIDForAdmins(id)
	default:
		return nil, ErrorNotAuthorized
	}
}

// @Summary Get all payments
// @Param limit query int false "number of elements to return"
// @Param offset query int false "offset in the list of elements to return"
// @Tags Payments
// @Success 200 {array} models.Payment
// @Router /payments [get]
// @Security ApiKeyAuth
func (c Controller) Payments(limit int, offset int, who int, whoKind string) ([]interface{}, error) {

	switch whoKind {
	case repository.ParentUser:
		return c.repo.PaymentsForParent(limit, offset, who)
	case repository.AdminUser:
		return c.repo.PaymentsForAdmin(limit, offset)

	default:
		return nil, ErrorNotAuthorized
	}
}

// @Summary Get all appointments
// @Param limit query int false "number of elements to return"
// @Param offset query int false "offset in the list of elements to return"
// @Tags Appointments
// @Router /appointments [get]
// @Success 200 {object} models.List
// @Security ApiKeyAuth
func (c Controller) Appointments(limit int, offset int, who int, whoKind string) ([]interface{}, error) {

	switch whoKind {
	case repository.ParentUser:
		return c.repo.AppointmentsForParents(limit, offset, who)

	case repository.AdminUser:
		return c.repo.AppointmentsForAdmin(limit, offset)

	default:
		return nil, ErrorNotAuthorized
	}
}

// LectureByClass(id int, limit int, offset int) (students []interface{}, err error)
// TODO GradeStudent(grade models.Grade) error
// TODO
// parents:
// see/pay (fake payment) upcoming scheduled payments (monthly, material, trips, err error)
// admins:
// everything
// @Summary Get all students
// @Param limit query int false "number of elements to return"
// @Param offset query int false "offset in the list of elements to return"
// @Tags Students
// @Success 200 {array} models.Student
// @Router /students [get]
// @Security ApiKeyAuth
func (c Controller) Students(limit int, offset int, who int, whoKind string) (student []interface{}, err error) {

	switch whoKind {
	case repository.ParentUser:
		return c.repo.StudentsForParent(limit, offset, who)
	case repository.AdminUser:
		return c.repo.StudentsForAdmins(limit, offset)
	default:
		return nil, ErrorNotAuthorized
	}
}

// see the grades obtained by their children
// @Summary Get grades of the student
// @Param id path int true "Student ID"
// @Param limit query int false "number of elements to return"
// @Param offset query int false "offset in the list of elements to return"
// @Tags Students
// @Success 200 {array} models.Grade
// @Router /students/{id}/grades [get]
// @Security ApiKeyAuth
func (c Controller) GradesByStudent(id int, limit int, offset int, who int, whoKind string) ([]interface{}, error) {

	switch whoKind {
	case repository.ParentUser:
		return c.repo.GradesByStudentForParent(id, limit, offset, who)
	case repository.AdminUser:
		return c.repo.GradesByStudentForAdmins(id, limit, offset)

	default:
		return nil, ErrorNotAuthorized
	}
}

// see/modify their personal data
// Get teacher by id
// @Summary Get a teacher by id
// @Param id path int true "Teacher ID"
// @Tags Teachers
// @Success 200 {object} models.Teacher
// @Router /teachers/{id} [get]
// @Security ApiKeyAuth
func (c Controller) TeacherByID(id int, who int, whoKind string) (interface{}, error) {
	switch whoKind {
	case repository.TeacherUser:
		if id == who {
			return c.repo.TeacherByIDForTeacher(id)
		} else {
			return nil, ErrorNotAuthorized
		}
	case repository.AdminUser:
		return c.repo.TeacherByIDForAdmin(id)
	default:
		return nil, ErrorNotAuthorized
	}
}

// List all teachers
// @Summary Get all teachers
// @Param limit query int false "number of elements to return"
// @Param offset query int false "offset in the list of elements to return"
// @Tags Teachers
// @Success 200 {array} models.Teacher
// @Router /teachers [get]
// @Security ApiKeyAuth
func (c Controller) Teachers(limit int, offset int, who int, whoKind string) ([]interface{}, error) {

	switch whoKind {
	case repository.TeacherUser:
		return c.repo.TeachersForTeacher(who)
	case repository.AdminUser:
		return c.repo.TeachersForAdmin(limit, offset)
	default:
		return nil, ErrorNotAuthorized
	}
}

// @Summary Get appointments of the teacher
// @Param id path int true "Teacher ID"
// @Param limit query int false "number of elements to return"
// @Param offset query int false "offset in the list of elements to return"
// @Tags Teachers
// @Success 200 {array} models.Appointment
// @Router /teachers/{id}/appointments [get]
// @Security ApiKeyAuth
func (c Controller) AppointmentsByTeacher(id int, limit int, offset int, who int, whoKind string) (appointments []interface{}, err error) {

	switch whoKind {
	case repository.TeacherUser:
		if id == who {
			return c.repo.AppointmentsByTeacherForTeacher(id, limit, offset, who)
		} else {
			return nil, ErrorNotAuthorized
		}
	case repository.AdminUser:
		return c.repo.AppointmentsByTeacherForAdmin(id, limit, offset)

	default:
		return nil, ErrorNotAuthorized
	}
}

// @Summary Get notifications of the teacher
// @Param id path int true "Teacher ID"
// @Param limit query int false "number of elements to return"
// @Param offset query int false "offset in the list of elements to return"
// @Tags Teachers
// @Success 200 {array} models.TimeTable
// @Router /teachers/{id}/notifications [get]
// @Security ApiKeyAuth
func (c Controller) NotificationsByTeacher(id int, limit int, offset int, who int, whoKind string) (notifications []interface{}, err error) {

	switch whoKind {
	case repository.TeacherUser:
		if id == who {
			return c.repo.NotificationsByTeacherForTeacher(id, limit, offset, who, whoKind)
		} else {
			return nil, ErrorNotAuthorized
		}
	case repository.AdminUser:
		return c.repo.NotificationsByTeacherForAdmin(id, limit, offset)

	default:
		return nil, ErrorNotAuthorized
	}
}

// @Summary Get subject taught by the teacher
// @Param id path int true "Teacher ID"
// @Param limit query int false "number of elements to return"
// @Param offset query int false "offset in the list of elements to return"
// @Success 200 {object} repository.Subjects
// @Tags Teachers
// @Router /teachers/{id}/subjects [get]
// @Security ApiKeyAuth
func (c Controller) SubjectsByTeacher(id int, limit int, offset int, who int, whoKind string) (notifications []interface{}, err error) {

	switch whoKind {
	case repository.TeacherUser:
		if id == who {
			return c.repo.SubjectsByTeacher(id, limit, offset)
		} else {
			return nil, ErrorNotAuthorized
		}
	case repository.AdminUser:
		return c.repo.SubjectsByTeacher(id, limit, offset)
	default:
		return nil, ErrorNotAuthorized
	}
}

// @Summary Get classes in which the subject is taught by the teacher
// @Param id path int true "Teacher ID"
// @Param subject path int true "Subject ID"
// @Param limit query int false "number of elements to return"
// @Param offset query int false "offset in the list of elements to return"
// @Success 200 {array} models.Class
// @Tags Teachers
// @Router /teachers/{id}/subjects/{subject} [get]
// @Security ApiKeyAuth
func (c Controller) ClassesBySubjectAndTeacher(teacher int, subject string, limit int, offset int, who int, whoKind string) ([]interface{}, error) {

	switch whoKind {
	case repository.TeacherUser:
		if teacher == who {
			return c.repo.ClassesBySubjectAndTeacher(teacher, subject, limit, offset)
		} else {
			return nil, ErrorNotAuthorized
		}
	case repository.AdminUser:
		return c.repo.ClassesBySubjectAndTeacher(teacher, subject, limit, offset)

	default:
		return nil, ErrorNotAuthorized
	}
}

// @Summary Get lectures taught by the teacher
// @Param id path int true "Teacher ID"
// @Param limit query int false "number of elements to return"
// @Param offset query int false "offset in the list of elements to return"
// @Tags Teachers
// @Success 200 {array} models.TimeTable
// @Router /teachers/{id}/lectures [get]
// @Security ApiKeyAuth
func (c Controller) LecturesByTeacher(id int, limit int, offset int, who int, whoKind string) (lectures []interface{}, err error) {

	switch whoKind {
	case repository.TeacherUser:
		if id == who {
			return c.repo.LecturesByTeacher(id, limit, offset)
		} else {
			return nil, ErrorNotAuthorized
		}
	case repository.AdminUser:
		return c.repo.LecturesByTeacher(id, limit, offset)
	default:
		return nil, ErrorNotAuthorized
	}

}

// @Summary Get classes in which the teacher teaches
// @Param id path int true "Teacher ID"
// @Param limit query int false "number of elements to return"
// @Param offset query int false "offset in the list of elements to return"
// @Tags Teachers
// @Success 200 {array} models.Class
// @Router /teachers/{id}/classes [get]
// @Security ApiKeyAuth
func (c Controller) ClassesByTeacher(id int, limit int, offset int, who int, whoKind string) ([]interface{}, error) {
	switch whoKind {
	case repository.TeacherUser:
		if id == who {
			return c.repo.ClassesByTeacher(id, limit, offset)
		} else {
			return nil, ErrorNotAuthorized
		}
	case repository.AdminUser:
		return c.repo.ClassesByTeacher(id, limit, offset)
	default:
		return nil, ErrorNotAuthorized
	}
}

// @Summary Get all lectures
// @Param limit query int false "number of elements to return"
// @Param offset query int false "offset in the list of elements to return"
// @Tags Lectures
// @Success 200 {array} models.TimeTable
// @Router /lectures [get]
// @Security ApiKeyAuth
func (c Controller) Lectures(limit int, offset int, who int, whoKind string) ([]interface{}, error) {

	switch whoKind {
	case repository.ParentUser:
		return c.repo.LecturesForParent(limit, offset, who)
	case repository.TeacherUser:
		return c.repo.LecturesByTeacher(who, limit, offset)

	case repository.AdminUser:
		return c.repo.LecturesForAdmin(limit, offset)

	default:
		return nil, ErrorNotAuthorized
	}
}

// @Summary Get all accounts
// @Param limit query int false "number of elements to return"
// @Param offset query int false "offset in the list of elements to return"
// @Tags Accounts
// @Success 200 {array} models.Login
// @Router /accounts [get]
// @Security ApiKeyAuth
func (c Controller) Accounts(limit int, offset int, who int, whoKind string) ([]interface{}, error) {
	if whoKind == repository.AdminUser {
		return c.repo.AccountsForAdmins(limit, offset)
	} else {
		return nil, ErrorNotAuthorized
	}
}
