package mysql

import "database/sql"

type dao struct {
	db *sql.DB
}

var Dao dao
