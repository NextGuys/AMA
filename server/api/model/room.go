package model

import (
	"github.com/jinzhu/gorm"
)

// Room UserModel
type Room struct {
	gorm.Model
	ID       int       `json:"id"`
	Title    string    `json:"title"`
	Users    []User    `gorm:"many2many:room_users;"`
	Messages []Message `gorm:"many2many:room_messages;"`
}

// Rooms List user
var Rooms []Room
