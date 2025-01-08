package mysql

import (
	"os"

	"github.com/go-sql-driver/mysql"
)

func NewMySQLConfig() *mysql.Config {
	return &mysql.Config{
		User:   os.Getenv("DB_USER"),
		Passwd: os.Getenv("DB_PASS"),
		Net:    "tcp",
		Addr:   os.Getenv("DB_HOST"),
		DBName: os.Getenv("DB_NAME"),
	}
}
