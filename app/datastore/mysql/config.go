package mysql

import (
	"github.com/go-sql-driver/mysql"
)

// NewMySQLConfig は MySQL の設定を生成します。
func NewMySQLConfig() *mysql.Config {
	return &mysql.Config{
		User:   "your-username", // 必要に応じて変更
		Passwd: "your-password", // 必要に応じて変更
		Net:    "tcp",
		Addr:   "127.0.0.1:3306", // 必要に応じて変更
		DBName: "your-database",  // 必要に応じて変更
	}
}
