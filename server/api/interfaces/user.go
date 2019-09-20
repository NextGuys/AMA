package interfaces

import (
	"net/http"

	repo "github.com/NextGuys/AMA/server/api/repository"

	"github.com/NextGuys/AMA/server/api/model"
	"github.com/NextGuys/AMA/server/api/utils"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/rs/xid"
)

// NewUserRepo Initialize user repository
func NewUserRepo(conn *gorm.DB) repo.UserRepo {
	return &UserRepo{
		Conn: conn,
	}
}

// UserRepo Handler with DB
type UserRepo struct {
	Conn *gorm.DB
}

// SignUp SignUp
func (h *UserRepo) SignUp(u *model.User) error {

	// Validate
	if u.Email == "" || u.Password == "" {
		return echo.NewHTTPError(http.StatusUnauthorized, "Please provide valid cred")
	}

	pwd := []byte(u.Password)
	hash := utils.HashAndSalt(pwd)
	uid := xid.New()

	u.UID = uid
	u.Password = hash
	if !h.Conn.NewRecord(&u) {
		panic("could not create new record")
	}
	if err := h.Conn.Create(&u).Error; err != nil {
		panic(err.Error())
	}

	return nil
}

// SignIn SignIn
func (h *UserRepo) SignIn(u *model.User) error {

	// Validate
	if u.Email == "" || u.Password == "" {
		return echo.NewHTTPError(http.StatusUnauthorized, "Please provide valid cred")
	}

	pwd := []byte(u.Password)
	hash := utils.HashAndSalt(pwd)
	uid := xid.New()

	u.UID = uid
	u.Password = hash

	h.Conn.Find(&u, model.User{Name: u.Name})

	return nil
}

// Create Create
func (h *UserRepo) Create(s *[]model.Skill, params *model.UserParams) error {
	var user model.User
	var skill model.Skill

	h.Conn.Find(&user, model.User{UID: params.UID})
	h.Conn.Model(&user).Related(&s, "Skills")

	skill.Name = params.Skill

	h.Conn.Model(&user).Association("Skills").Append(&skill)
	h.Conn.Model(&user).Association("Skills").Find(&s)

	return nil
}

// Read Read
func (h *UserRepo) Read(users *[]model.User) error {
	//var skill model.Skill

	h.Conn.Find(&users)

	// h.Conn.Find(&user, model.User{UID: params.UID})
	// h.Conn.Model(&user).Related(&s, "Skills")

	// skill.Name = params.Skill

	// h.Conn.Model(&user).Association("Skills").Append(&skill)
	// h.Conn.Model(&user).Association("Skills").Find(&s)

	return nil
}

func (h *UserRepo) GetUser(u *model.User, uid xid.ID) error {
	//var skill model.Skill

	h.Conn.First(&u, model.User{UID: uid})

	// h.Conn.Find(&user, model.User{UID: params.UID})
	// h.Conn.Model(&user).Related(&s, "Skills")

	// skill.Name = params.Skill

	// h.Conn.Model(&user).Association("Skills").Append(&skill)
	// h.Conn.Model(&user).Association("Skills").Find(&s)

	return nil
}
