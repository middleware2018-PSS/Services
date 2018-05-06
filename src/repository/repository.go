package repository

import "github.com/middleware2018-PSS/Services/src/models"

type StudentRepository interface {
	ById (id string) (models.Student)

}
