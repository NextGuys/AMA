package model

import (
	"github.com/jinzhu/gorm"
)

// Skill UserModel
type Skill struct {
	gorm.Model
	Name  string `json:"name"`
	Users []User `gorm:"many2many:user_skills;"`
}

// Skills Slice of User
var Skills []Skill
