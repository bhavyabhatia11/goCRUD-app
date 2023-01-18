package api

import (
	"net/http"
	"strconv"

	"github.com/bhavyaunacademy/goCRUD/handlers/utils"
	"github.com/labstack/echo/v4"
)

func GetAllUsers(c echo.Context) error {
	users := utils.ReadAllUsers()
	return c.JSON(http.StatusCreated, users)
}

func GetUserByUID(c echo.Context) error {
	uid, err := strconv.Atoi(c.Param("uid"))
	if err != nil {
		return c.JSON(403, "Invalid uid")
	}
	users := utils.ReadAllUsers()
	if uid > len(users.Users) || uid < 1 {
		return c.JSON(404, "User Not Found")
	}
	return c.JSON(http.StatusCreated, users.Users[uid-1])

}
