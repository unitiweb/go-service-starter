package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

var Conn *gorm.DB

func Connect() {

	connected := false
	tryCount := 0
	for connected == false {
		db, err := gorm.Open("mysql", "root:password1@(mysql)/starter?parseTime=true")
		if err != nil {
			tryCount++
			fmt.Println("Trying to connect to database: try #", tryCount)
			time.Sleep(2000 * time.Millisecond)
		} else {
			Conn = db
			connected = true
		}
	}
}

func Migrate() {
	Conn.AutoMigrate(
		&Player{},
		&Course{},
	)
}

func Close() {
	Conn.Close()
}
