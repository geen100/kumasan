package mysql

import (
	"fmt"
	"os"

	"github.com/go-sql-driver/mysql"
)

func NewMySQLConfig() *mysql.Config {
	host := os.Getenv("MYSQL_HOST")
	port := os.Getenv("MYSQL_PORT")
	user := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASSWORD")
	dbName := os.Getenv("MYSQL_DATABASE")

	cfg := mysql.Config{
		User:                 user,
		Passwd:               password,
		Net:                  "tcp",
		Addr:                 fmt.Sprintf("%s:%s", host, port),
		DBName:               dbName,
		ParseTime:            true,
		AllowNativePasswords: true,
	}

	return &cfg
}
