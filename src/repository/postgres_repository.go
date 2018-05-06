package repository

import (
	"database/sql"
)

type postgresRepository struct {
	*sql.DB
}

func NewPostgresRepository(DB *sql.DB) *postgresRepository {
	return &postgresRepository{DB}
}
