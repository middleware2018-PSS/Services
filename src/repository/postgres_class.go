package repository

import (
	"github.com/middleware2018-PSS/Services/src/models"
	"log"
)

func (r *postgresRepository) ClassesByID(id int64) (class models.Class) {
	err := r.QueryRow(`SELECT id, year, section, info, grade FROM back2school.classes 
								WHERE id = $1`, id).Scan(&class.ID, &class.Year, &class.Section, &class.Info, &class.Grade)
	if err != nil {
		log.Print(err)
	}
	return class
}
