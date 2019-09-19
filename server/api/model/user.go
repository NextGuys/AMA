package model

import (
	"github.com/jinzhu/gorm"
	"github.com/rs/xid"
)

// User UserModel
type User struct {
	gorm.Model
	UID      xid.ID  `json:"uid"`
	Name     string  `json:"name"`
	Email    string  `json:"email"`
	Password string  `json:"password"`
	Rooms    []Room  `gorm:"many2many:room_users;"`
	Skills   []Skill `gorm:"many2many:user_skills;"`
}

// Users Slice of User
var Users []User

// UserParams UserParams
type UserParams struct {
	UID   xid.ID `json:"uid"`
	Skill string `json:"skill"`
}
