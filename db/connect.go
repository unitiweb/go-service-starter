package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/unitiweb/go-service-starter/database/models"
	"time"
)

var Conn *gorm.DB

func Connect() {

	//db, err := gorm.Open("mysql", "root:password1@(mysql)/starter?parseTime=true")

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

	//if err != nil {
	//	panic("faield to connect to database")
	//}
}

func Migrate() {
	Conn.AutoMigrate(&models.Player{})
}

func Close() {
	Conn.Close()
}
