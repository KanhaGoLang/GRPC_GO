package common

import (
	"database/sql"
)

func NewDatabaseConnection() (*sql.DB, error) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "password"
	dbName := "sanjeev"
	dbAddr := "tcp(localhost:3306)"

	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@"+dbAddr+"/"+dbName)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
