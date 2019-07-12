package main

import(
	"fmt"
	"net/http"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"strconv"
	"strings"
)

type (

	user struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Password string `json:"password"`
		PhoneNumber string `json:"phone_number"`
		EmailAddress string `json:"email_address"`
		IsAdministrator bool `json:"is_administrator"`
	}
)

var (
	users = map[int]*user{}
	seq = 1
)

//--------------
// Handlers
//--------------

func createUser(c echo.Context) error {
	u := &user {
		ID: seq,
	}

	if err := c.Bind(u); err != nil {
		return err
	}

	users[u.ID] = u
	seq++
	return c.JSON(http.StatusCreated, u)
}

func getUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	return c.JSON(http.StatusOK, users[id])
}

func updateUser(c echo.Context) error {
	u := new(user)
	if err := c.Bind(u); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Param("id"))
	users[id].Name = u.Name
	return c.JSON(http.StatusOK, users[id])
}

func deleterUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	delete(users, id)
	return c.NoContent(http.StatusNoContent)
}

func main() {
	e := echo.New();

	e.Use(middleware.Logger());
	e.Use(middleware.Recover())

	// Route => handler
	e.GET("/", func(c echo.Context) error {
		return
	})
}
