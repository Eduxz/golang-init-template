package controllers

import (
	"net/http"
	"time"

	"github.com/firebase-golang/app/models"
	"github.com/firebase-golang/database/connections"

	"github.com/dgrijalva/jwt-go"
	"github.com/firebase-golang/app/utils"
	"github.com/labstack/echo"
)

// Claim to auth
type Claim struct {
	Username string `json:"username"`
	Usertype string `json:"role"`
	jwt.StandardClaims
}

type payloadLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Login User
func Login(c echo.Context) error {
	u := new(payloadLogin)
	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusBadRequest, utils.ResponseError(err, "Failed"))
	}

	user := models.User{Email: u.Username, Password: u.Password}
	user.CheckLogin()
	if !user.IsValid() {
		return echo.ErrUnauthorized
	}
	t, err := generateToken(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ResponseError(err, "Failed"))
	}

	return c.JSON(http.StatusOK, utils.ResponseSuccess(echo.Map{
		"token": t,
		"user":  user,
	}, "Success"))
}

// Register new users
func Register(c echo.Context) error {
	u := new(models.User)
	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusBadRequest, utils.ResponseError(err, "Failed"))
	}
	u.Role = "low"
	if s := connections.BooleanQueryPG(u.Create); !s {
		return c.JSON(http.StatusBadRequest, utils.ResponseError(echo.Map{"message": "Error Creating"}, "Failed"))
	}
	t, _ := generateToken(*u)

	return c.JSON(http.StatusCreated, utils.ResponseSuccess(
		echo.Map{"token": t, "user": *u}, "Success"))
}

func generateToken(u models.User) (string, error) {
	claims := &Claim{
		u.Email,
		u.Role,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}

	return t, nil
}

func accessible(c echo.Context) error {
	return c.String(http.StatusOK, "Accessible")
}

func restricted(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*Claim)
	name := claims.Username
	return c.String(http.StatusOK, "Welcome "+name+"!")
}
