package user

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func InitUserRoutes(e *echo.Echo) {
	e.POST("/user/create", CreateNewUserHandler)
}

func CreateNewUserHandler(c echo.Context) error {
	type UserReq struct {
		Name string `json:"name"`
	}

	var userReq UserReq
	err := c.Bind(&userReq)
	if err != nil {
		return err
	}

	u, err := MakeNewUser(userReq.Name)
	if err != nil {
		return err
	}

	AddToUserList(&u)

	return c.JSON(http.StatusCreated, u)
}
