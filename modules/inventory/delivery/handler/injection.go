package handler

import "database/sql"

type DatabaseInjection struct {
	DB *sql.DB
}
