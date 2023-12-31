package config

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectToDatabase() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/read_file_csv?parseTime=true")
	if err != nil {
		return nil, err
	}

	return db, nil
}
