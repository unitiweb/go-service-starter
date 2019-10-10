package migrations

import (
	"github.com/unitiweb/go-service-starter/database"
	"github.com/unitiweb/go-service-starter/database/models"
)

func Migrate001() {
	database.Db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&models.Player{})
}