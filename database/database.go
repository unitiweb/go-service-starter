package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//type Config struct {
//	Dialect  string
//	Host     string
//	Port     string
//	Database string
//	User     string
//	Password string
//}
//
//var Conn *sql.DB
//
//func Connect(c Config) {
//	switch c.Dialect {
//	case "mysql":
//		Conn = connector.MySql(c.User, c.Password, c.Host, c.Port, c.Database)
//	}
//	if Conn == nil {
//		panic("No valid database dialect was configured")
//	}
//}

var Conn *gorm.DB

func Connect() {
	db, err := gorm.Open("mysql", "root:password1@(mysql)/starter?parseTime=true")
	if err != nil {
		panic("faield to connect to database")
	}
	Conn = db
}

func Migrate() {
	//Conn.AutoMigrate(&models.Player{})
}

//func Db() *gorm.DB {
//	db, err := gorm.Open("mysql", "root:password1@(mysql)/starter?parseTime=true")
//	if err != nil {
//		panic("faield to connect to database")
//	}
//	//defer db.Close()
//	return db
//}
