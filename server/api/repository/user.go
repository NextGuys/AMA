package repo

import "github.com/NextGuys/AMA/server/api/model"

// UserRepo UserRepo
type UserRepo interface {
	SignUp(u *model.User) error
	SignIn(u *model.User) error
	Read(users *[]model.User) error
	Create(s *[]model.Skill, params *model.UserParams) error
}
