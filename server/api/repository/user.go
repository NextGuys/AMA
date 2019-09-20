package repo

import (
	"github.com/NextGuys/AMA/server/api/model"
	"github.com/rs/xid"
)

// UserRepo UserRepo
type UserRepo interface {
	SignUp(u *model.User) error
	SignIn(u *model.User) error
	Read(users *[]model.User) error
	GetUser(u *model.User, uid xid.ID) error
	Create(s *[]model.Skill, params *model.UserParams) error
}
