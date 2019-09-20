package repo

import "github.com/NextGuys/AMA/server/api/model"

// UserRepo UserRepo
type RoomRepo interface {
	Create(params *model.RoomParams) error
}
