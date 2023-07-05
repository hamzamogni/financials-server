package postgres

import "database/sql"

type DB struct {
	db *sql.DB
}
