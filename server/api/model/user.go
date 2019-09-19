package model

import (
	"github.com/jinzhu/gorm"
)

// User UserModel
type User struct {
	gorm.Model
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Email    string  `json:"email"`
	Password string  `json:"password"`
	Rooms    []Room  `gorm:"many2many:room_users;"`
	Skills   []Skill `gorm:"many2many:user_skills;"`
}

// Users Slice of User
var Users []User
