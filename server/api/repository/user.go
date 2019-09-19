package repo

import "github.com/NextGuys/AMA/server/api/model"

type UserRepo interface {
	SignUp(u *model.User) error
	Login(u *model.User) error
}
