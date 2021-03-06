package handler

import (
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

var router *mux.Router

type (
	// Handler Handler
	Handler struct {
		DB *gorm.DB
	}
)

const (
	// Key This should be imported from somewhere else
	Key = "secret"
)

// InitializeRouter Init Router
func InitializeRouter(db *gorm.DB, e *echo.Echo) *echo.Echo {
	// h := &Handler{DB: db}
	u := NewUserHandler(db)
	r := NewRoomHandler(db)
	// User
	e.GET("/users", u.Read)
	e.GET("/users/:id", u.GetUser)
	e.POST("/signup", u.SignUp)
	e.POST("/signin", u.SignIn)
	e.POST("/skills", u.Create)
	// Room
	e.GET("/rooms", r.Read)
	e.POST("/rooms", r.Create)
	// // User
	// e.GET("/users/:id", u.GetUserByID)
	// // Tag
	// e.POST("/tags", h.Create)
	//e.GET("/tags/:name/users", h.GetUserByTag)
	return e
}
