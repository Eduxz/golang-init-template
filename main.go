package main

import (
	"fmt"

	m "migration"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func getSomething(c echo.Context) error {
	test := c.QueryParams()
	fmt.Println(test)

	return c.JSON(400, test)
}

func main() {
	m.Migrate()
	e := echo.New()
	e.Use(middleware.Logger())
	//e.Use(middleware.Recover())

	e.Static("/", "public")
	e.GET("/something", getSomething)

	// Restricted group
	r := e.Group("/restricted")

	// Configure middleware with the custom claims type
	// config := middleware.JWTConfig{
	// 	Claims:     &jwtCustomClaims{},
	// 	SigningKey: []byte("secret"),
	// }
	//r.Use(middleware.JWTWithConfig(config))

	e.Logger.Fatal(e.Start(":1323"))
}
