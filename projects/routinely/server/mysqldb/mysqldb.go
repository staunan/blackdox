package mysqldb

import (
	"database/sql"

	"github.com/go-sql-driver/mysql"
)

const (
	MySQLUserName     = "root"
	MySQLPassword     = "0000"
	MySQLHostName     = "127.0.0.1:3306"
	MySQLDatabaseName = "routinely"
)

func ConnectMySQL() (*sql.DB, error) {
	// Create a DB Connection --
	cfg := mysql.Config{
		User:   MySQLUserName,
		Passwd: MySQLPassword,
		Net:    "tcp",
		Addr:   MySQLHostName,
		DBName: MySQLDatabaseName,
	}
	// Get a database handle.
	return sql.Open("mysql", cfg.FormatDSN())
}
