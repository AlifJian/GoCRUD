package config

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func getDB() *sql.DB {
	dbName := "goCRUD"
	dbHost := "localhost"
	dbPort := "3306"

	db, err := sql.Open("mysql", "root:@tcp("+dbHost+":"+dbPort+")/"+dbName)
	if err != nil {
		panic(err)
	}
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}
