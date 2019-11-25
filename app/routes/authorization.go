package routes

import (
	"github.com/firebase-golang/app/controllers"
	"github.com/labstack/echo"
)

// PublicAuthRoutes authentication routes
func PublicAuthRoutes(r *echo.Echo) {
	r.POST("/login", controllers.Login)
	r.POST("/register", controllers.Register)
}

// PrivateAuthRoutes authentication routes
func PrivateAuthRoutes(r *echo.Echo) {

}
