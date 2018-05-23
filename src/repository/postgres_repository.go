package repository

import (
	"database/sql"
	"fmt"
	"log"
)

type postgresRepository struct {
	*sql.DB
}

func NewPostgresRepository(DB *sql.DB) *postgresRepository {
	//TODO prepare all statement at startup
	return &postgresRepository{DB}
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
		return switchResults(list, err)
	} else {
		return switchResults(list, sql.ErrNoRows)
	}
}

func (r *postgresRepository) exec(query string, params ...interface{}) (id int, err error) {
	res, err := r.DB.Exec(query, params...)
	if err != nil {
		log.Print(err.Error())
	}
	if id, e := res.LastInsertId(); e != nil {
		return int(id), switchErrors(err)
	} else {
		return 0, err
	}
}

func (r *postgresRepository) execUpdate(query string, params ...interface{}) (err error) {
	_, err = r.exec(query, params...)
	return err
}

func switchResult(res interface{}, e error) (interface{}, error) {
	//TODO: check err
	if e = switchErrors(e); e != nil {
		return nil, e
	} else {
		return res, nil
	}
}

func switchResults(res []interface{}, e error) ([]interface{}, error) {
	//TODO: check err
	if e = switchErrors(e); e != nil {
		return nil, e
	} else {
		return res, nil
	}
}

func switchErrors(e error) error {
	if e != nil {

		switch e {
		case sql.ErrNoRows:
			return ErrNoResult

		default:
			log.Printf("%v", e)
			return nil
		}
	} else {
		return nil
	}

}
