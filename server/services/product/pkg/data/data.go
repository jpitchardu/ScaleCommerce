package data

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "127.0.0.1"
	port     = "5432"
	user     = "scale_commerce"
	password = "S3cre7P4ssw0rd"
	dbname   = "product"
)

func ConnectToDB() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Println("Successfully connected to database!")

	return db, err
}

func CloseDB(db *sql.DB) {
	db.Close()
}
