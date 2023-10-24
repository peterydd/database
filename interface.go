package database

import "database/sql"

type Database interface {
	Open() (*sql.DB, error)
}
