package main

import (
	"time"

	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
	"github.com/tylerb/graceful"
	"github.com/mc2soft/dating/echo/controllers"
)

func main() {
	e := echo.New()

	e.Use(mw.Logger())
	e.Use(mw.Recover())

	e.Post("/auth/register", controllers.AuthRegister)
	graceful.ListenAndServe(e.Server(":9000"), 10*time.Second)
}
