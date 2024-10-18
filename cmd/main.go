package main

import (
	"github.com/cslemes/hc-web/internal/handlers"

	"github.com/cslemes/hc-web/internal/utils"
	"github.com/cslemes/hc-web/internal/views"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	e := echo.New()

	e.Use(middleware.Logger())

	e.Renderer = views.NewTemplates()

	e.GET("/", func(c echo.Context) error {
		return c.Render(200, "layout", nil)
	})

	e.GET("/home", func(c echo.Context) error {
		return c.Render(200, "home", nil)
	})
	e.GET("/characters", func(c echo.Context) error {
		return c.Render(200, "characters", nil)
	})
	e.GET("/under", func(c echo.Context) error {
		return c.Render(200, "under", nil)
	})

	//	e.GET("/characters", handlers.Characters())

	e.GET("/cards", handlers.Cards())

	e.Static("/static/", "public")

	conf := utils.AppConfig()
	port := conf.Server.Port

	e.Logger.Printf("Server starting on http://localhost:%s", port)
	e.Logger.Fatal(e.Start(":" + port))
}
