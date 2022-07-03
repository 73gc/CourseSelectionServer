package sqlcontroller

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func Init() (err error) {
	Db, err = sql.Open("mysql", "gouser:acmicpc@(localhost)/courseselection")
	return
}
