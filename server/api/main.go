package main

import (
	"github.com/NextGuys/AMA/server/api/db"
	"github.com/NextGuys/AMA/server/api/handler"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func init() {
	db.New()
	db.Init()
}

func main() {
	db := db.GetDB()
	defer db.Close()

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.CORS())
	e.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(handler.Key),
		Skipper: func(c echo.Context) bool {
			// Skip authentication for and signup login requests
			if c.Path() == "/signin" || c.Path() == "/signup" {
				return true
			}
			return false
		},
	}))

	// Initialize handler
	handler.InitializeRouter(db, e)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
