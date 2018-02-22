package db

import (
	"database/sql"
)

type rowScanner interface {
	Scan(dest ...interface{}) error
}

type dbActions interface {
	QueryRow(query string, args ...interface{}) *sql.Row
}
