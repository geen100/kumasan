package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

// NewDB はデータベース接続を初期化して返します
func NewDB() (*sql.DB, error) {
	dbConn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_PORT"),
		os.Getenv("MYSQL_DATABASE"),
	)

	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		log.Println("Failed to connect to database:", err)
		return nil, err
	}

	if err := db.Ping(); err != nil {
		log.Println("Failed to ping database:", err)
		return nil, err
	}

	return db, nil
}

// CloseDB はデータベース接続を閉じます
func CloseDB(db *sql.DB) {
	if err := db.Close(); err != nil {
		log.Println("Failed to close the database connection:", err)
	}
}
