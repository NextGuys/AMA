package interfaces

import (
	repo "github.com/NextGuys/AMA/server/api/repository"

	"github.com/NextGuys/AMA/server/api/model"
	"github.com/jinzhu/gorm"
)

// NewRoomRepo Initialize user repository
func NewRoomRepo(conn *gorm.DB) repo.RoomRepo {
	return &RoomRepo{
		Conn: conn,
	}
}

// RoomRepo Handler with DB
type RoomRepo struct {
	Conn *gorm.DB
}

// Create Create
func (h *RoomRepo) Create(params *model.RoomParams) error {
	//var users []model.User
	var room model.Room
	var user model.User

	room.Title = params.Title

	if !h.Conn.NewRecord(&room) {
		panic("could not create new record")
	}
	if err := h.Conn.Create(&room).Error; err != nil {
		panic(err.Error())
	}

	//h.Conn.Find(&user, model.User{UID: params.UID})
	//h.Conn.Model(&room).Related(&users, "Users")

	user.UID = params.UID

	h.Conn.Model(&room).Association("Users").Append(&user)

	return nil
}

// Read Read
func (h *RoomRepo) Read(rooms *[]model.Room) error {
	//var skill model.Skill

	h.Conn.Find(&rooms)

	// h.Conn.Find(&user, model.User{UID: params.UID})
	// h.Conn.Model(&user).Related(&s, "Skills")

	// skill.Name = params.Skill

	// h.Conn.Model(&user).Association("Skills").Append(&skill)
	// h.Conn.Model(&user).Association("Skills").Find(&s)

	return nil
}
