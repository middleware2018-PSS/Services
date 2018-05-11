package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
)

type postgresRepository struct {
	*sql.DB
}

var (
	ErrNoResult      = errors.New("No result found.")
	ErrorNotBlocking = errors.New("Something went wrong but no worriez.")
)

func NewPostgresRepository(DB *sql.DB) *postgresRepository {
	return &postgresRepository{DB}
}

func switchError(err error) error {
	switch err {
	case sql.ErrNoRows:
		err = ErrNoResult
	default:
		if fmt.Sprintf("%v", err)[:len("sql: Scan error")] == "sql: Scan error" {
			err = ErrorNotBlocking
		}
	}
	return err
}

func (r *postgresRepository) listByParams(query string, f func(*sql.Rows) (interface{}, error), limit int, offset int, params ...interface{}) (list []interface{}, err error) {
	query = query + fmt.Sprintf(" limit $%d offset $%d", len(params)+1, len(params)+2)
	params = append(params, limit, offset)
	rows, err := r.Query(query, params...)
	defer rows.Close()
	if err != nil {
		log.Print(err)
	}
	for rows.Next() {
		el, err := f(rows)
		if err != nil {
			//TODO check error
		}
		list = append(list, el)
	}
	return list, switchError(err)
}
