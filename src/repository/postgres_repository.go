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
	ErrNoResult      = errors.New("No results found.")
	ErrorNotBlocking = errors.New("Something went wrong but no worriez.")
)

func NewPostgresRepository(DB *sql.DB) *postgresRepository {
	return &postgresRepository{DB}
}

func switchError(err error) error {
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			err = ErrNoResult
		default:
			if l := len("sql: Scan error"); len(err.Error()) >= l && err.Error()[:l] == "sql: Scan error" {
				err = ErrorNotBlocking
			}
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
	} else {
		for rows.Next() {
			el, err := f(rows)
			if err != nil {
				//TODO check error
			}
			list = append(list, el)
		}
	}
	if len(list) > 0 {
		return list, switchError(err)
	} else {
		return list, switchError(sql.ErrNoRows)
	}
}

func getListByIDOffsetLimit(id int64, limit int, offset int, f func(int64, int, int) ([]interface{}, error)) ([]interface{}, error) {
	res, e := f(id, limit, offset)
	//TODO: check err
	switch e {
	case ErrNoResult:
		return nil, e
	case ErrorNotBlocking:
		return res, nil
	default:
		return res, nil
	}
}

func getByID(id int64, f func(int64) (interface{}, error)) (interface{}, error) {
	res, e := f(id)
	//TODO: check err
	switch e {
	case ErrNoResult:
		return nil, e
	case ErrorNotBlocking:
		return res, nil
	default:
		return res, nil
	}
}
