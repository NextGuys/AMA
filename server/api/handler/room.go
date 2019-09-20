package handler

import (
	"net/http"

	"github.com/NextGuys/AMA/server/api/interfaces"
	"github.com/NextGuys/AMA/server/api/model"
	repo "github.com/NextGuys/AMA/server/api/repository"

	"github.com/jinzhu/gorm"

	"github.com/labstack/echo"
)

// NewUserHandler Initialize user repository
func NewRoomHandler(conn *gorm.DB) *RoomHandler {
	return &RoomHandler{
		repo: interfaces.NewRoomRepo(conn),
	}
}

// UserHandler Handler with DB
type RoomHandler struct {
	repo repo.RoomRepo
}

// Create Create
func (h *RoomHandler) Create(c echo.Context) (err error) {
	r := []model.Room{}
	params := &model.RoomParams{}

	if err = c.Bind(params); err != nil {
		return
	}

	h.repo.Create(params)

	if err != nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "invalid email or password"}
	}
	return c.JSON(http.StatusCreated, r)
}

// SignUp SignUp
func (h *RoomHandler) Read(c echo.Context) (err error) {
	var r model.Room
	rooms := &[]model.Room{}
	if err := c.Bind(r); err != nil {
		return err
	}

	if err := h.repo.Read(rooms); err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, rooms)
}

//SignIn log in

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
