package services

import "database/sql"

type ServicesDB struct {
	DB *sql.DB
}

func CreateServicesDB(db *sql.DB) *ServicesDB {
	return &ServicesDB{DB: db}
}
