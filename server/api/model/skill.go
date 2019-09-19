package model

import (
	"github.com/jinzhu/gorm"
	"github.com/rs/xid"
)

// Skill UserModel
type Skill struct {
	gorm.Model
	UID      xid.ID `json:"uid"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Users    []User `gorm:"many2many:user_skills;"`
}

// Skills Slice of User
var Skills []Skill
