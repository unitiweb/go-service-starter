package db

import (
	"time"
)

// The players data struct
type Player struct {
	Id         *int       `json:"id"`
	PlayerName *string    `json:"playerName" validate:"required"`
	CreatedAt  *time.Time `json:"createdAt"`
	UpdatedAt  *time.Time `json:"updatedAt"`
	DeletedAt  *time.Time `json:"deletedAt"`
}

// Get player by id
func (p Player) Find(id int) Player {
	Conn.First(&p, id)
	return p
}

func (p Player) Update() {
	Conn.Save(&p)
}

// Get player by player name
func (p Player) ByPlayerName(n string) Player {
	Conn.Where("playerName = ?", n).First(&p)
	return p
}
