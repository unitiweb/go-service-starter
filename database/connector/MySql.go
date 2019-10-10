package connector

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func MySql(u string, p string, h string, pt string, d string) *sql.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", u, p, h, pt, d)
	conn, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err.Error())
	}
	//defer conn.Close()
	return conn
}
