package model

import (
	"github.com/jinzhu/gorm"
	"github.com/rs/xid"
)

// Room UserModel
type Room struct {
	gorm.Model
	Title    string    `json:"title"`
	Users    []User    `gorm:"many2many:room_users;"`
	Messages []Message `gorm:"many2many:room_messages;"`
}

// Rooms List user
var Rooms []Room

type RoomParams struct {
	UID   xid.ID `json:"uid"`
	Title string `json:"title"`
}
