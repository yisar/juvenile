package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var (
	dbConn *sql.DB
	err    error
)

func init() {
	str := "root:123456@tcp(localhost:3306)/acton?charset=utf8"
	dbConn, err = sql.Open("mysql", str)
	if err != nil {
		panic(err.Error())
	}
}