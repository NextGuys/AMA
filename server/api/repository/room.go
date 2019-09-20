package repo

import "github.com/NextGuys/AMA/server/api/model"

// UserRepo UserRepo
type RoomRepo interface {
	Read(rooms *[]model.Room) error
	Create(params *model.RoomParams) error
}
