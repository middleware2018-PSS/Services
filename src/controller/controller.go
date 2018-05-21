package controller

import (
	"github.com/middleware2018-PSS/Services/src/models"
	"github.com/middleware2018-PSS/Services/src/repository"
)

type controller struct {
	r repository.Repository
}

func NewController(r repository.Repository) *controller {
	return &controller{r}
}

// @Summary Get a class by id
// @Param id path int true "Class ID"
// @Tags Classes
// @Router /classes/{id} [get]
func (c controller) ClassByID(id int64) (interface{}, error) {
	return c.r.ClassByID(id)
}

// @Summary Get a notification by id
// @Param id path int true "Notification ID"
// @Tags Notifications
// @Router /notifications/{id} [get]
func (c *controller) NotificationByID(id int64) (interface{}, error) {
	return c.r.NotificationByID(id)
}
// Get payment by id
// @Summary Get a payment by id
// @Param id path int true "Payment ID"
// @Tags Payments
// @Router /payments/{id} [get]
func (c *controller) PaymentByID(id int64) (interface{}, error) {
	return c.r.PaymentByID(id)
}

// Parents
// see/modify their personal data
// @Summary Get a parent by id
// @Param id path int true "Account ID"
// @Tags Parents
// @Router /parents/{id} [get]
func (c *controller) ParentByID(id int64) (interface{}, error) {
	return c.r.ParentByID(id)
}
// @Summary Update parents's data
// @Param id path int true "Parent ID"
// @Param parent body models.Parent true "data"
// @Tags Parents
// @Router /parents/{id} [put]
func (c *controller) UpdateParent(p models.Parent) error {
	return c.r.UpdateParent(p)
}

// @Summary Update student's data
// @Param id path int true "Student ID"
// @Param student body models.Student true "data"
// @Tags Students
// @Router /students/{id} [put]
func (c *controller) UpdateStudent(student models.Student) error {
	return c.r.UpdateStudent(student)
}

// see/modify the personal data of their registered children
// @Summary Get children of the parent
// @Param id path int true "Parent ID"
// @Param limit query int false "number of elements to return"
// @Param offset query int false "offset in the list of elements to return"
// @Tags Parents
// @Router /parents/{id}/students [get]
func (c *controller) ChildrenByParent(id int64, limit int, offset int) ([]interface{}, error) {
	return c.r.ChildrenByParent(id, limit, offset)
}
// Get student by id
// @Summary Get a student by id
// @Param id path int true "Student ID"
// @Tags Students
// @Router /students/{id} [get]
func (c *controller) StudentByID(id int64) (interface{}, error) {
	return c.r.StudentByID(id)
}

// see the grades obtained by their children
// @Summary Get grades of the student
// @Param id path int true "Student ID"
// @Param limit query int false "number of elements to return"
// @Param offset query int false "offset in the list of elements to return"
// @Tags Students
// @Router /students/{id}/grades [get]
func (c *controller) GradesByStudent(id int64, limit int, offset int) ([]interface{}, error) {
	return c.r.GradesByStudent(id, limit, offset)
}

// see the monthly payments that have been made to the school in the past
// @Summary Get payments of the parent
// @Param id path int true "Parent ID"
// @Param limit query int false "number of elements to return"
// @Param offset query int false "offset in the list of elements to return"
// @Tags Parents
// @Router /parents/{id}/payments [get]
func (c *controller) PaymentsByParent(id int64, limit int, offset int) ([]interface{}, error) {
	return c.r.PaymentsByParent(id, limit, offset)
}

// see general/personal notifications coming from the school
// @Summary Get notifications of the parent
// @Param id path int true "Parent ID"
// @Param limit query int false "number of elements to return"
// @Param offset query int false "offset in the list of elements to return"
// @Tags Parents
// @Router /parents/{id}/notifications [get]
func (c *controller) NotificationsByParent(id int64, limit int, offset int) ([]interface{}, error) {
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
// @Router /parents/{id}/appointments [get]
func (c *controller) AppointmentsByParent(id int64, limit int, offset int) ([]interface{}, error) {
	return c.r.AppointmentsByParent(id, limit, offset)
}

// @Summary Update appointment's data
// @Param id path int true "Appointment ID"
// @Param appointment body models.Appointment true "data"
// @Tags Appointments
// @Router /appointments/{id} [put]
func (c *controller) UpdateAppointment(appointment models.Appointment) error {
	return c.r.UpdateAppointment(appointment)
}
// @Summary Get a appointment by id
// @Param id path int true "Appointment ID"
// @Tags Classes
// @Router /classes/{id} [get]
func (c *controller) AppointmentByID(id int64) (interface{}, error) {
	return c.r.AppointmentByID(id)
}

// see/modify their personal data
// Get teacher by id
// @Summary Get a teacher by id
// @Param id path int true "Teacher ID"
// @Tags Teachers
// @Router /teachers/{id} [get]
func (c *controller) TeacherByID(id int64) (teacher interface{}, err error) {
	return c.r.TeacherByID(id)
}

// @Summary Update teacher's data
// @Param id path int true "Teacher ID"
// @Param teacher body models.Teacher true "data"
// @Tags Teachers
// @Router /teachers/{id} [put]
func (c *controller) UpdateTeacher(teacher models.Teacher) (err error) {
	return c.r.UpdateTeacher(teacher)
}

// see the classrooms in which they teach, with information regarding the argument that they teach
// in that class, the students that make up the class, and the complete lesson timetable for that
// class
// @Summary Get subject taught by the teacher
// @Param id path int true "Teacher ID"
// @Param limit query int false "number of elements to return"
// @Param offset query int false "offset in the list of elements to return"
// @Tags Teachers
// @Router /teachers/{id}/subjects [get]
func (c *controller) SubjectsByTeacher(id int64, limit int, offset int) ([]interface{}, error) {
	return c.r.SubjectsByTeacher(id, limit, offset)
}

// @Summary Get classes in which the subject is taught by the teacher
// @Param id path int true "Teacher ID"
// @Param subject path int true "Subject ID"
// @Param limit query int false "number of elements to return"
// @Param offset query int false "offset in the list of elements to return"
// @Tags Teachers
// @Router /teachers/{id}/subjects/{subject} [get]
func (c *controller) ClassesBySubjectAndTeacher(teacher int64, subject string, limit int, offset int) ([]interface{}, error) {
	return c.r.ClassesBySubjectAndTeacher(teacher, subject, limit, offset)
}

// @Summary Get a student by class
// @Param id path int true "Class ID"
// @Param limit query int false "number of elements to return"
// @Param offset query int false "offset in the list of elements to return"
// @Tags Classes
// @Router /classes/{id}/students [get]
func (c *controller) StudentsByClass(id int64, limit int, offset int) ([]interface{}, error) {
	return c.r.StudentsByClass(id, limit, offset)
}
// @Summary Get a lecture by class
// @Param limit query int false "number of elements to return"
// @Param offset query int false "offset in the list of elements to return"
// @Param id path int true "Class ID"
// @Tags Students
// @Router /students/{id} [get]
func (c *controller) LectureByClass(id int64, limit int, offset int) ([]interface{}, error) {
	return c.r.LectureByClass(id, limit, offset)
}

// LectureByClass(id int64, limit int, offset int) (students []interface{}, err error)

// @Summary Get appointments of the teacher
// @Param id path int true "Teacher ID"
// @Param limit query int false "number of elements to return"
// @Param offset query int false "offset in the list of elements to return"
// @Tags Teachers
// @Router /teachers/{id}/appointments [get]
func (c *controller) AppointmentsByTeacher(id int64, limit int, offset int) ([]interface{}, error) {
	return c.r.AppointmentsByTeacher(id, limit, offset)
}

// @Summary Get notifications of the teacher
// @Param id path int true "Teacher ID"
// @Param limit query int false "number of elements to return"
// @Param offset query int false "offset in the list of elements to return"
// @Tags Teachers
// @Router /teachers/{id}/notifications [get]
func (c *controller) NotificationsByTeacher(id int64, limit int, offset int) ([]interface{}, error) {
	return c.r.NotificationsByTeacher(id, limit, offset)
}

// @Summary Get lectures taught by the teacher
// @Param id path int true "Teacher ID"
// @Param limit query int false "number of elements to return"
// @Param offset query int false "offset in the list of elements to return"
// @Tags Teachers
// @Router /teachers/{id}/lectures [get]
func (c *controller) LecturesByTeacher(id int64, limit int, offset int) ([]interface{}, error) {
	return c.r.LecturesByTeacher(id, limit, offset)
}

// @Summary Get classes in which the teacher teaches
// @Param id path int true "Teacher ID"
// @Param limit query int false "number of elements to return"
// @Param offset query int false "offset in the list of elements to return"
// @Tags Teachers
// @Router /teachers/{id}/classes [get]
func (c *controller) ClassesByTeacher(id int64, limit int, offset int) ([]interface{}, error) {
	return c.r.ClassesByTeacher(id, limit, offset)
}

// LectureByClass(id int64, limit int, offset int) (students []interface{}, err error)
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
// @Router /students [get]
func (c *controller) Students(limit int, offset int) ([]interface{}, error) {
	return c.r.Students(limit, offset)
}

// List all teachers
// @Summary Get all teachers
// @Param limit query int false "number of elements to return"
// @Param offset query int false "offset in the list of elements to return"
// @Tags Teachers
// @Router /teachers [get]
func (c *controller) Teachers(limit int, offset int) ([]interface{}, error) {
	return c.r.Teachers(limit, offset)
}
// @Summary Get all parents
// @Param limit query int false "number of elements to return"
// @Param offset query int false "offset in the list of elements to return"
// @Tags Parents
// @Router /parents [get]
func (c *controller) Parents(limit int, offset int) ([]interface{}, error) {
	return c.r.Parents(limit, offset)
}
// @Summary Get all payments
// @Param limit query int false "number of elements to return"
// @Param offset query int false "offset in the list of elements to return"
// @Tags Payments
// @Router /payments [get]
func (c *controller) Payments(limit int, offset int) ([]interface{}, error) {
	return c.r.Payments(limit, offset)
}
// List all notifications
// @Summary Get all notifications
// @Param limit query int false "number of elements to return"
// @Param offset query int false "offset in the list of elements to return"
// @Tags Notifications
// @Router /notifications [get]
func (c *controller) Notifications(limit int, offset int) ([]interface{}, error) {
	return c.r.Notifications(limit, offset)
}
// @Summary Get all classes
// @Param limit query int false "number of elements to return"
// @Param offset query int false "offset in the list of elements to return"
// @Tags Classes
// @Router /classes [get]
func (c *controller) Classes(limit int, offset int) ([]interface{}, error) {
	return c.r.Classes(limit, offset)
}
// @Summary Get a login token
// @Param account body models.Account true "Add account"
// @Tags Auth
// @Router /login [post]
func (c *controller) CheckUser(id string, pass string) (string, bool) {
	return c.r.CheckUser(id, pass)
}
func (c *controller) UserKind(userID string) map[string]interface{} {
	return c.r.UserKind(userID)
}
// @Summary Create parent
// @Tags Parents
// @Param parent body models.Parent true "data"
// @Tags Parents
// @Router /parents [post]
func (c *controller) CreateParent(parent models.Parent) (int64, error) {
	return c.r.CreateParent(parent)
}
// @Summary Create appointment
// @Param id path int true "Appointment ID"
// @Param appointment body models.Appointment true "data"
// @Tags Appointments
// @Router /appointments [post]

func (c *controller) CreateAppointment(appointment models.Appointment) (int64, error) {
	return c.r.CreateAppointment(appointment)
}

// @Summary Create teacher
// @Param teacher body models.Teacher true "data"
// @Tags Teachers
// @Router /teachers [post]
func (c *controller) CreateTeacher(teacher models.Teacher) (int64, error) {
	return c.r.CreateTeacher(teacher)
}

// @Summary Get all appointments
// @Param limit query int false "number of elements to return"
// @Param offset query int false "offset in the list of elements to return"
// @Tags Appointments
// @Router /appointments [get]
func (c *controller) Appointments(limit int, offset int) ([]interface{}, error) {
	return c.r.Appointments(limit, offset)
}
// @Summary Get all grades
// @Param limit query int false "number of elements to return"
// @Param offset query int false "offset in the list of elements to return"
// @Tags Grades
// @Router /grades [get]
func (c *controller) Grades(limit int, offset int) ([]interface{}, error) {
	return c.r.Grades(limit, offset)
}
