package models

import (
	_ "github.com/doug-martin/goqu/v9/dialect/mysql"
	"github.com/jinzhu/gorm"
	"github.com/unitiweb/go-service-starter/database"
)

// The name of the players' database table
const PlayerTableName = "Players"

// The players data struct
type Player struct {
	//Id         *int       `json:"id"`
	PlayerName *string `json:"playerName"`
	gorm.Model
	//CreatedAt  *time.Time `json:"createdAt"`
	//UpdatedAt  *time.Time `json:"updatedAt"`
	//DeletedAt  *time.Time `json:"deletedAt"`
}

// Get and return the player's table name
func (p Player) TableName() string {
	return PlayerTableName
}

type playerStruct struct {
	id         int
	playerName string
	createdAt  string
	updatedAt  string
	deletedAt  string
}

// Find one player by id
func (p Player) Find(id int) {
	database.Conn.Where("id = ?", id).First(p)
}

// Persist the player to the database
func (p Player) Save() Player {
	return p
}

// Delete the player from the database
func (p Player) Delete(id int) Player {
	//p.UpdatedAt = Stamp()
	//p.DeletedAt = Stamp()
	return p
}
