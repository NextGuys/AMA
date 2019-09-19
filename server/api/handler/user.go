package handler

import (
	"net/http"
	"time"

	"github.com/NextGuys/AMA/server/api/interfaces"
	"github.com/NextGuys/AMA/server/api/model"
	repo "github.com/NextGuys/AMA/server/api/repository"
	"github.com/NextGuys/AMA/server/api/utils"
	"github.com/dgrijalva/jwt-go"

	"github.com/jinzhu/gorm"

	"github.com/labstack/echo"
)

// NewUserHandler Initialize user repository
func NewUserHandler(conn *gorm.DB) *UserHandler {
	return &UserHandler{
		repo: interfaces.NewUserRepo(conn),
	}
}

// UserHandler Handler with DB
type UserHandler struct {
	repo repo.UserRepo
}

// SignUp SignUp
func (h *UserHandler) SignUp(c echo.Context) (err error) {
	u := &model.User{}
	if err := c.Bind(u); err != nil {
		return err
	}

	if err := h.repo.SignUp(u); err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, u)
}

// Create Create
func (h *UserHandler) Create(c echo.Context) (err error) {
	s := []model.Skill{}
	params := &model.UserParams{}

	if err = c.Bind(params); err != nil {
		return
	}

	h.repo.Create(&s, params)

	if err != nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "invalid email or password"}
	}
	return c.JSON(http.StatusCreated, s)
}

//SignIn log in
func (h *UserHandler) SignIn(c echo.Context) (err error) {
	u := &model.User{}
	if err = c.Bind(u); err != nil {
		return
	}
	pwd := []byte(u.Password)

	h.repo.SignIn(u)
	if utils.ComparePasswords(u.Password, pwd) {
		token := jwt.New(jwt.SigningMethodHS256)
		claims := token.Claims.(jwt.MapClaims)
		claims["admin"] = true
		claims["sub"] = u.UID
		claims["name"] = u.Name
		claims["iat"] = time.Now()
		claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
		tokenString, _ := token.SignedString([]byte(Key))
		return c.JSON(http.StatusCreated, tokenString)
	}

	return &echo.HTTPError{Code: http.StatusBadRequest, Message: "invalid email or password"}
}

// // GetUserByID for getting user info by ID
// func (h *UserHandler) GetUserByID(c echo.Context) (err error) {
// 	var u model.User
// 	var tags []model.Tag
// 	var id string
// 	uid, _ := xid.FromString(id)

// 	h.Conn.First(&u, model.User{UID: uid})
// 	h.Conn.Model(&u).Association("Tags").Find(&tags)

// 	data := map[string]interface{}{"uid": u.UID, "name": u.Name, "tags": tags}

// 	if err != nil {
// 		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "invalid email or password"}
// 	}

// 	return c.JSON(http.StatusCreated, data)
// }

// // GetUserByTag UIDでユーザー情報を取得
// func (h *UserHandler) GetUserByTag(c echo.Context) (err error) {
// 	var users []model.User
// 	var uid []xid.ID
// 	t := &model.Tag{}
// 	if err = c.Bind(t); err != nil {
// 		return
// 	}

// 	h.Conn.Where("tags.name=?", t.Name).Select("DISTINCT(uid)").Joins("JOIN user_tags ON user_tags.user_id = users.id").
// 		Joins("JOIN tags ON user_tags.tag_id=tags.id").Find(&users)
// 	for _, v := range users {
// 		uid = append(uid, v.UID)
// 	}

// 	data := map[string]interface{}{"uid": uid}

// 	if err != nil {
// 		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "invalid email or password"}
// 	}

// 	return c.JSON(http.StatusCreated, data)
// }
