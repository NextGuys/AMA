package model

import (
	"github.com/jinzhu/gorm"
)

// Message UserModel
type Message struct {
	gorm.Model
	ID      int    `json:"id"`
	UserID  int    `json:"user_id"`
	RoomID  int    `json:"room_id"`
	Content string `json:"content"`
	Rooms   []Room `gorm:"many2many:room_users;"`
}

// Messages List users
var Messages []Message
