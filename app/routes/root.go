package routes

import (
	"github.com/firebase-golang/app/controllers"
	"github.com/firebase-golang/app/utils"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func getSomething(c echo.Context) error {
	test, _ := c.FormParams()
	algo := utils.ResponseSuccess(test, "Success")
	return c.JSON(200, algo)
}

// SetRoutes is a configuration root of routes in the app
func SetRoutes() *echo.Echo {
	e := echo.New()
	e.Use(middleware.Logger())
	//e.Use(middleware.Recover())

	publicRoutes(e)
	privateRoutes(e.Group("/private"))
	return e
}

func staticRoutes(r *echo.Echo) {
	r.Static("/", "public")
}

func publicRoutes(r *echo.Echo) {
	r.GET("/something", getSomething)
	PublicAuthRoutes(r)
}

func privateRoutes(r *echo.Group) {
	config := middleware.JWTConfig{
		Claims:     controllers.Claim{},
		SigningKey: []byte("secret"),
	}
	r.Use(middleware.JWTWithConfig(config))

}
