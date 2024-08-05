package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-openapi/strfmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"v-sadovsky/gateway/server/models"
)

var empty struct{}

func httpErrorMsg(err error) *models.ErrorMessage {
	if err == nil {
		return nil
	}
	return &models.ErrorMessage{
		Message: err.Error(),
	}
}

func createUser(c echo.Context) error {
	var request models.CreateUserRequest
	if err := json.NewDecoder(c.Request().Body).Decode(&request); err != nil {
		return c.JSON(http.StatusBadRequest, httpErrorMsg(err))
	}

	if err := request.Validate(strfmt.Default); err != nil {
		return c.JSON(http.StatusBadRequest, httpErrorMsg(err))
	}

	// Business logic
	var userID int64 = 3
	fmt.Printf("create user: %+v\n", request)

	response := models.CreateUserResponse{ID: userID}
	return c.JSON(http.StatusCreated, response)
}

func getUser(c echo.Context) error {
	// User ID from path `users/:id`
	id := c.Param("id")
	userID, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, httpErrorMsg(err))
	}

	// ...

	response := models.GetUserResponse{
		ID:    int64(userID),
		Email: "test@mail.ru",
		Name:  "test",
	}
	return c.JSON(http.StatusOK, response)
}

func updateUser(c echo.Context) error {
	// User ID from path `users/:id`
	id := c.Param("id")
	userID, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, httpErrorMsg(err))
	}

	var request models.UpdateUserRequest
	if err := json.NewDecoder(c.Request().Body).Decode(&request); err != nil {
		return c.JSON(http.StatusBadRequest, httpErrorMsg(err))
	}

	if err := request.Validate(strfmt.Default); err != nil {
		return c.JSON(http.StatusBadRequest, httpErrorMsg(err))
	}

	// ...

	response := models.UpdateUserResponse{
		ID:    int64(userID),
		Email: "test@mail.ru",
		Name:  "test",
	}
	return c.JSON(http.StatusOK, response)
}

func deleteUser(c echo.Context) error {
	// User ID from path `users/:id`
	id := c.Param("id")
	userID, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, httpErrorMsg(err))
	}

	_ = userID

	return c.JSON(http.StatusOK, empty)
}

func main() {
	e := echo.New()

	e.POST("/messenger/v1/profiles", createUser)
	e.GET("/messenger/v1/profiles/:id", getUser)
	e.PUT("/messenger/v1/profiles/:id", updateUser)
	e.DELETE("/messenger/v1/profiles/:id", deleteUser)

	e.Logger.Fatal(e.Start(":8080"))
}
