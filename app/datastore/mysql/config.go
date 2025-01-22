package mysql

import (
	"fmt"
	"os"

	"github.com/go-sql-driver/mysql"
)

// NewMySQLConfig は、環境変数から取得した情報をもとに
// MySQLドライバ（github.com/go-sql-driver/mysql）の設定を生成して返します。
func NewMySQLConfig() *mysql.Config {
	host := os.Getenv("MYSQL_HOST")
	port := os.Getenv("MYSQL_PORT")
	user := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASSWORD")
	dbName := os.Getenv("MYSQL_DATABASE")

	// github.com/go-sql-driver/mysql のConfig構造体を組み立てる
	cfg := mysql.Config{
		User:                 user,
		Passwd:               password,
		Net:                  "tcp",
		Addr:                 fmt.Sprintf("%s:%s", host, port),
		DBName:               dbName,
		ParseTime:            true, // 日付時刻型を time.Time で受け取れるように
		AllowNativePasswords: true, // MySQL 8.x系でパスワード認証の問題が出ないように
	}

	return &cfg
}
