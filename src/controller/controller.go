package controller

import (
	"github.com/middleware2018-PSS/Services/src/models"
	"github.com/middleware2018-PSS/Services/src/repository"
)

type Controller struct {
	r repository.Repository
}
type Token struct {
	Token  string `json:"token"`
	Expire string `json:"expire"`
}


func NewController(r repository.Repository) *Controller {
	return &Controller{r}
}

// @Summary Get a class by id
// @Param id path int true "Class ID"
// @Tags Classes
// @Success 200 {object} models.Class
// @Router /classes/{id} [get]
func (c Controller) ClassByID(id int) (interface{}, error) {
	return c.r.ClassByID(id)
}

// @Summary Get a grade by id
// @Param id path int true "Grade ID"
// @Tags Grades
// @Success 200 {object} models.Grade
// @Router /grades/{id} [get]
func (c Controller) GradeByID(id int) (interface{}, error) {
	return c.r.GradeByID(id)
}

// @Summary Get a notification by id
// @Param id path int true "Notification ID"
// @Tags Notifications
// @Success 200 {object} models.Notification
// @Router /notifications/{id} [get]
func (c *Controller) NotificationByID(id int) (interface{}, error) {
	return c.r.NotificationByID(id)
}

// Get payment by id
// @Summary Get a payment by id
// @Param id path int true "Payment ID"
// @Tags Payments
// @Success 200 {object} models.Payment
// @Router /payments/{id} [get]
func (c *Controller) PaymentByID(id int) (interface{}, error) {
	return c.r.PaymentByID(id)
}

// Parents
// see/modify their personal data
// @Summary Get a parent by id
// @Param id path int true "Account ID"
// @Tags Parents
// @Success 200 {object} models.Parent
// @Router /parents/{id} [get]
func (c *Controller) ParentByID(id int) (interface{}, error) {
	return c.r.ParentByID(id)
}

// @Summary Update parents's data
// @Param id path int true "Parent ID"
// @Param parent body models.Parent true "data"
// @Tags Parents
// @Success 201 {object} models.Parent
// @Router /parents/{id} [put]
func (c *Controller) UpdateParent(p models.Parent) error {
	return c.r.UpdateParent(p)
}

// @Summary Update student's data
// @Param id path int true "Student ID"
// @Param student body models.Student true "data"
// @Tags Students
// @Success 201 {object} models.Student
// @Router /students/{id} [put]
func (c *Controller) UpdateStudent(student models.Student) error {
	return c.r.UpdateStudent(student)
}

// see/modify the personal data of their registered children
// @Summary Get children of the parent
// @Param id path int true "Parent ID"
// @Param limit query int false "number of elements to return"
// @Param offset query int false "offset in the list of elements to return"
// @Tags Parents
// @Success 200 {array} models.Student
// @Router /parents/{id}/students [get]
func (c *Controller) ChildrenByParent(id int, limit int, offset int) ([]interface{}, error) {
	return c.r.ChildrenByParent(id, limit, offset)
}

// Get student by id
// @Summary Get a student by id
// @Param id path int true "Student ID"
// @Tags Students
// @Success 200 {object} models.Student
// @Router /students/{id} [get]
func (c *Controller) StudentByID(id int) (interface{}, error) {
	return c.r.StudentByID(id)
}

// see the grades obtained by their children
// @Summary Get grades of the student
// @Param id path int true "Student ID"
// @Param limit query int false "number of elements to return"
// @Param offset query int false "offset in the list of elements to return"
// @Tags Students
// @Success 200 {array} models.Grade
// @Router /students/{id}/grades [get]
func (c *Controller) GradesByStudent(id int, limit int, offset int) ([]interface{}, error) {
	return c.r.GradesByStudent(id, limit, offset)
}

// see the monthly payments that have been made to the school in the past
// @Summary Get payments of the parent
// @Param id path int true "Parent ID"
// @Param limit query int false "number of elements to return"
// @Param offset query int false "offset in the list of elements to return"
// @Tags Parents
// @Success 200 {array} models.Payment
// @Router /parents/{id}/payments [get]
func (c *Controller) PaymentsByParent(id int, limit int, offset int) ([]interface{}, error) {
	return c.r.PaymentsByParent(id, limit, offset)
}

// see general/personal notifications coming from the school
// @Summary Get notifications of the parent
// @Param id path int true "Parent ID"
// @Param limit query int false "number of elements to return"
// @Param offset query int false "offset in the list of elements to return"
// @Tags Parents
// @Success 200 {array} models.Notification
// @Router /parents/{id}/notifications [get]
func (c *Controller) NotificationsByParent(id int, limit int, offset int) ([]interface{}, error) {
	return c.r.NotificationsByParent(id, limit, offset)
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
func (c *Controller) AppointmentsByParent(id int, limit int, offset int) ([]interface{}, error) {
	return c.r.AppointmentsByParent(id, limit, offset)
}

// @Summary Update appointment's data
// @Param id path int true "Appointment ID"
// @Param appointment body models.Appointment true "data"
// @Tags Appointments
// @Success 201 {object} models.Appointment
// @Router /appointments/{id} [put]
func (c *Controller) UpdateAppointment(appointment models.Appointment) error {
	return c.r.UpdateAppointment(appointment)
}

// @Summary Get a appointment by id
// @Param id path int true "Appointment ID"
// @Tags Appointments
// @Success 200 {object} models.Appointment
// @Router /appointments/{id} [get]
func (c *Controller) AppointmentByID(id int) (interface{}, error) {
	return c.r.AppointmentByID(id)
}

// see/modify their personal data
// Get teacher by id
// @Summary Get a teacher by id
// @Param id path int true "Teacher ID"
// @Tags Teachers
// @Success 200 {object} models.Teacher
// @Router /teachers/{id} [get]
func (c *Controller) TeacherByID(id int) (teacher interface{}, err error) {
	return c.r.TeacherByID(id)
}

// @Summary Update teacher's data
// @Param id path int true "Teacher ID"
// @Param teacher body models.Teacher true "data"
// @Tags Teachers
// @Success 204 {object} models.Teacher
// @Router /teachers/{id} [put]
func (c *Controller) UpdateTeacher(teacher models.Teacher) (err error) {
	return c.r.UpdateTeacher(teacher)
}

// see the classrooms in which they teach, with information regarding the argument that they teach
// in that class, the students that make up the class, and the complete lesson timetable for that
// class
type Subjects struct {
	Subjects []string `json:"subjects" example:"science"`
}

// @Summary Get subject taught by the teacher
// @Param id path int true "Teacher ID"
// @Param limit query int false "number of elements to return"
// @Param offset query int false "offset in the list of elements to return"
// @Success 200 {object} Controller.Subjects
// @Tags Teachers
// @Router /teachers/{id}/subjects [get]
func (c *Controller) SubjectsByTeacher(id int, limit int, offset int) ([]interface{}, error) {
	return c.r.SubjectsByTeacher(id, limit, offset)
}

// @Summary Get classes in which the subject is taught by the teacher
// @Param id path int true "Teacher ID"
// @Param subject path int true "Subject ID"
// @Param limit query int false "number of elements to return"
// @Param offset query int false "offset in the list of elements to return"
// @Success 200 {array} models.Class
// @Tags Teachers
// @Router /teachers/{id}/subjects/{subject} [get]
func (c *Controller) ClassesBySubjectAndTeacher(teacher int, subject string, limit int, offset int) ([]interface{}, error) {
	return c.r.ClassesBySubjectAndTeacher(teacher, subject, limit, offset)
}

// @Summary Get a student by class
// @Param id path int true "Class ID"
// @Param limit query int false "number of elements to return"
// @Param offset query int false "offset in the list of elements to return"
// @Tags Classes
// @Success 200 {array} models.Student
// @Router /classes/{id}/students [get]
func (c *Controller) StudentsByClass(id int, limit int, offset int) ([]interface{}, error) {
	return c.r.StudentsByClass(id, limit, offset)
}

// @Summary Get a lecture by class
// @Param limit query int false "number of elements to return"
// @Param offset query int false "offset in the list of elements to return"
// @Param id path int true "Class ID"
// @Tags Students
// @Success 200 {array} models.Appointment
// @Router /students/{id} [get]
func (c *Controller) LectureByClass(id int, limit int, offset int) ([]interface{}, error) {
	return c.r.LectureByClass(id, limit, offset)
}

// LectureByClass(id int, limit int, offset int) (students []interface{}, err error)

// @Summary Get appointments of the teacher
// @Param id path int true "Teacher ID"
// @Param limit query int false "number of elements to return"
// @Param offset query int false "offset in the list of elements to return"
// @Tags Teachers
// @Success 200 {array} models.Appointment
// @Router /teachers/{id}/appointments [get]
func (c *Controller) AppointmentsByTeacher(id int, limit int, offset int) ([]interface{}, error) {
	return c.r.AppointmentsByTeacher(id, limit, offset)
}

// @Summary Get notifications of the teacher
// @Param id path int true "Teacher ID"
// @Param limit query int false "number of elements to return"
// @Param offset query int false "offset in the list of elements to return"
// @Tags Teachers
// @Success 200 {array} models.TimeTable
// @Router /teachers/{id}/notifications [get]
func (c *Controller) NotificationsByTeacher(id int, limit int, offset int) ([]interface{}, error) {
	return c.r.NotificationsByTeacher(id, limit, offset)
}

// @Summary Get lectures taught by the teacher
// @Param id path int true "Teacher ID"
// @Param limit query int false "number of elements to return"
// @Param offset query int false "offset in the list of elements to return"
// @Tags Teachers
// @Success 200 {array} models.TimeTable
// @Router /teachers/{id}/lectures [get]
func (c *Controller) LecturesByTeacher(id int, limit int, offset int) ([]interface{}, error) {
	return c.r.LecturesByTeacher(id, limit, offset)
}

// @Summary Get classes in which the teacher teaches
// @Param id path int true "Teacher ID"
// @Param limit query int false "number of elements to return"
// @Param offset query int false "offset in the list of elements to return"
// @Tags Teachers
// @Success 200 {array} models.Class
// @Router /teachers/{id}/classes [get]
func (c *Controller) ClassesByTeacher(id int, limit int, offset int) ([]interface{}, error) {
	return c.r.ClassesByTeacher(id, limit, offset)
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
func (c *Controller) Students(limit int, offset int) ([]interface{}, error) {
	return c.r.Students(limit, offset)
}

// List all teachers
// @Summary Get all teachers
// @Param limit query int false "number of elements to return"
// @Param offset query int false "offset in the list of elements to return"
// @Tags Teachers
// @Success 200 {array} models.Teacher
// @Router /teachers [get]
func (c *Controller) Teachers(limit int, offset int) ([]interface{}, error) {
	return c.r.Teachers(limit, offset)
}

// @Summary Get all parents
// @Param limit query int false "number of elements to return"
// @Param offset query int false "offset in the list of elements to return"
// @Tags Parents
// @Success 200 {array} models.Parent
// @Router /parents [get]
func (c *Controller) Parents(limit int, offset int) ([]interface{}, error) {
	return c.r.Parents(limit, offset)
}

// @Summary Get all payments
// @Param limit query int false "number of elements to return"
// @Param offset query int false "offset in the list of elements to return"
// @Tags Payments
// @Success 200 {array} models.Payment
// @Router /payments [get]
func (c *Controller) Payments(limit int, offset int) ([]interface{}, error) {
	return c.r.Payments(limit, offset)
}

// List all notifications
// @Summary Get all notifications
// @Param limit query int false "number of elements to return"
// @Param offset query int false "offset in the list of elements to return"
// @Tags Notifications
// @Success 200 {array} models.Notification
// @Router /notifications [get]
func (c *Controller) Notifications(limit int, offset int) ([]interface{}, error) {
	return c.r.Notifications(limit, offset)
}

// @Summary Get all classes
// @Param limit query int false "number of elements to return"
// @Param offset query int false "offset in the list of elements to return"
// @Tags Classes
// @Success 200 {array} models.Class
// @Router /classes [get]
func (c *Controller) Classes(limit int, offset int) ([]interface{}, error) {
	return c.r.Classes(limit, offset)
}

// @Summary Get a login token
// @Param account body models.Account true "Add account"
// @Tags Auth
// @Success 200 {object} Controller.Token
// @Router /login [post]
func (c *Controller) CheckUser(id string, pass string) (int, string, bool) {
	// TODO save kind and id in context
	return c.r.CheckUser(id, pass)
}


// @Summary Create parent
// @Tags Parents
// @Param parent body models.Parent true "data"
// @Tags Parents
// @Success 201 {object} models.Parent
// @Router /parents [post]
func (c *Controller) CreateParent(parent models.Parent) (int, error) {
	return c.r.CreateParent(parent)
}

// @Summary Create appointment
// @Param id path int true "Appointment ID"
// @Param appointment body models.Appointment true "data"
// @Tags Appointments
// @Router /appointments [post]
// @Success 201 {object} models.Appointment
func (c *Controller) CreateAppointment(appointment models.Appointment) (int, error) {
	return c.r.CreateAppointment(appointment)
}

// @Summary Create teacher
// @Param teacher body models.Teacher true "data"
// @Tags Teachers
// @Router /teachers [post]
// @Success 201 {object} models.Teacher
func (c *Controller) CreateTeacher(teacher models.Teacher) (int, error) {
	return c.r.CreateTeacher(teacher)
}

// @Summary Get all appointments
// @Param limit query int false "number of elements to return"
// @Param offset query int false "offset in the list of elements to return"
// @Tags Appointments
// @Router /appointments [get]
// @Success 200 {object} representations.List
// @Security ApiKeyAuth
func (c *Controller) Appointments(limit int, offset int) ([]interface{}, error) {
	return c.r.Appointments(limit, offset)
}

// @Summary Get all grades
// @Param limit query int false "number of elements to return"
// @Param offset query int false "offset in the list of elements to return"
// @Tags Grades
// @Success 200 {array} models.Grade
// @Router /grades [get]
func (c *Controller) Grades(limit int, offset int) ([]interface{}, error) {
	return c.r.Grades(limit, offset)
}

// @Summary Create student
// @Param student body models.Student true "data"
// @Tags Students
// @Router /students [post]
// @Success 201 {object} models.Student
func (c *Controller) CreateStudent(student models.Student) (int, error) {
	return c.r.CreateStudent(student)
}

// @Summary Create class
// @Param class body models.Class true "data"
// @Tags Classes
// @Router /classes [post]
// @Success 201 {object} models.Class
func (c *Controller) CreateClass(class models.Class) (int, error) {
	return c.r.CreateClass(class)
}

// @Summary Update Class's data
// @Param id path int true "Class ID"
// @Param parent body models.Class true "data"
// @Tags Classes
// @Success 201 {object} models.Class
// @Router /classes/{id} [put]
func (c *Controller) UpdateClass(class models.Class) error {
	return c.r.UpdateClass(class)
}

// @Summary Create notification
// @Param class body models.Notification true "data"
// @Tags Notifications
// @Router /notifications [post]
// @Success 201 {object} models.Notification
func (c *Controller) CreateNotification(notification models.Notification) (int, error) {
	return c.r.CreateNotification(notification)
}

// @Summary Update notification
// @Param id path int true "Notification ID"
// @Param class body models.Notification true "data"
// @Tags Notifications
// @Router /notifications/{id} [put]
// @Success 201 {object} models.Notification
func (c *Controller) UpdateNotification(notification models.Notification) error {
	return c.r.UpdateNotification(notification)
}

// @Summary Update Grade
// @Param id path int true "Grade ID"
// @Param class body models.Grade true "data"
// @Tags Grades
// @Router /grades/{id} [put]
// @Success 201 {object} models.Grade
func (c *Controller) UpdateGrade(grade models.Grade) error {
	return c.r.UpdateGrade(grade)
}

// @Summary Create grade
// @Param class body models.Grade true "data"
// @Tags Grades
// @Router /grades [post]
// @Success 201 {object} models.Grade
func (c *Controller) CreateGrade(grade models.Grade) (int, error) {
	return c.r.CreateGrade(grade)
}

// @Summary Update payment
// @Param id path int true "Payment ID"
// @Param class body models.Payment true "data"
// @Tags Payments
// @Router /payments/{id} [put]
// @Success 201 {object} models.Payment
func (c *Controller) UpdatePayment(payment models.Payment) error {
	return c.r.UpdatePayment(payment)
}

// @Summary Create payment
// @Param class body models.Payment true "data"
// @Tags Payments
// @Router /payments [post]
// @Success 201 {object} models.Payment
func (c *Controller) CreatePayment(payment models.Payment) (int, error) {
	return c.r.CreatePayment(payment)
}
