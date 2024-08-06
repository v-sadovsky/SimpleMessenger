package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/go-openapi/strfmt"
	"github.com/labstack/echo/v4"
	"io"
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
	body, _ := io.ReadAll(c.Request().Body)
	fmt.Println(string(body))

	if err := json.NewDecoder(bytes.NewReader(body)).Decode(&request); err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, httpErrorMsg(err))
	}

	if err := request.Validate(strfmt.Default); err != nil {
		return c.JSON(http.StatusBadRequest, httpErrorMsg(err))
	}

	// Business logic
	var userID int64 = 3

	response := models.CreateUserResponse{ID: userID}
	return c.JSON(http.StatusCreated, response)
}

func getUser(c echo.Context) error {
	id := c.Param("id")
	userID, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, httpErrorMsg(err))
	}

	// Business logic

	response := models.GetUserResponse{
		ID:    int64(userID),
		Email: "test@mail.ru",
		Name:  "test",
	}
	return c.JSON(http.StatusOK, response)
}

func updateUser(c echo.Context) error {
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

	// Business logic

	response := models.UpdateUserResponse{
		ID:    int64(userID),
		Email: "test@mail.ru",
		Name:  "test",
	}
	return c.JSON(http.StatusOK, response)
}

func deleteUser(c echo.Context) error {
	id := c.Param("id")
	userID, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, httpErrorMsg(err))
	}

	_ = userID

	return c.JSON(http.StatusOK, empty)
}

func makeFriendship(c echo.Context) error {
	id := c.Param("id")
	userID, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, httpErrorMsg(err))
	}
	fmt.Println("UserId", userID)

	var request models.CreateUserFriendship
	body, _ := io.ReadAll(c.Request().Body)
	fmt.Println(string(body))

	if err := json.NewDecoder(bytes.NewReader(body)).Decode(&request); err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, httpErrorMsg(err))
	}

	if err := request.Validate(strfmt.Default); err != nil {
		return c.JSON(http.StatusBadRequest, httpErrorMsg(err))
	}

	// Business logic

	return c.JSON(http.StatusOK, empty)
}

func getUserFriends(c echo.Context) error {
	id := c.Param("id")
	userID, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, httpErrorMsg(err))
	}

	fmt.Println("UserId", userID)

	// Business logic

	response := []models.GetUserResponse{
		models.GetUserResponse{
			ID:    3,
			Email: "friend1@mail.ru",
			Name:  "friend1",
		},
		models.GetUserResponse{
			ID:    7,
			Email: "friend2@mail.ru",
			Name:  "friend2",
		},
	}
	return c.JSON(http.StatusOK, response)
}

func handleFriendshipRequest(c echo.Context) error {
	id := c.Param("id")
	userID, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, httpErrorMsg(err))
	}
	fmt.Println("UserId", userID)

	state := c.QueryParam("accepted")
	accepted, err := strconv.ParseBool(state)
	if err != nil {
		return c.JSON(http.StatusBadRequest, httpErrorMsg(err))
	}

	// Business logic

	response := models.AcceptResponse{
		Status: &accepted,
	}
	return c.JSON(http.StatusOK, response)
}

func deleteFriend(c echo.Context) error {
	id := c.Param("id")
	userID, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, httpErrorMsg(err))
	}

	friendName := c.QueryParam("name")

	// Business logic
	_ = userID
	_ = friendName

	return c.JSON(http.StatusOK, empty)
}

func sendMessage(c echo.Context) error {
	id := c.Param("id")
	userID, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, httpErrorMsg(err))
	}
	fmt.Println("UserId", userID)

	var request models.SendMessage
	body, _ := io.ReadAll(c.Request().Body)
	fmt.Println(string(body))

	if err := json.NewDecoder(bytes.NewReader(body)).Decode(&request); err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, httpErrorMsg(err))
	}

	if err := request.Validate(strfmt.Default); err != nil {
		return c.JSON(http.StatusBadRequest, httpErrorMsg(err))
	}

	// Business logic

	return c.JSON(http.StatusOK, empty)
}

func getMessages(c echo.Context) error {
	id := c.Param("id")
	userID, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, httpErrorMsg(err))
	}

	fmt.Println("UserId", userID)

	// Business logic

	response := []models.GetMessages{
		models.GetMessages{
			ID:       3,
			Message:  "hello",
			UserName: "friend1",
		},
		models.GetMessages{
			ID:       7,
			Message:  "hi",
			UserName: "friend2",
		},
	}
	return c.JSON(http.StatusOK, response)
}

func main() {
	e := echo.New()

	e.POST("/messenger/v1/profiles", createUser)
	e.GET("/messenger/v1/profiles/:id", getUser)
	e.PUT("/messenger/v1/profiles/:id", updateUser)
	e.DELETE("/messenger/v1/profiles/:id", deleteUser)

	e.POST("/messenger/v1/friends/:id", makeFriendship)
	e.GET("/messenger/v1/friends/:id", getUserFriends)
	e.PUT("/messenger/v1/friends/:id", handleFriendshipRequest)
	e.DELETE("/messenger/v1/friends/:id", deleteFriend)

	e.POST("/messenger/v1/chats/:id", sendMessage)
	e.GET("/messenger/v1/chats/:id", getMessages)

	e.Logger.Fatal(e.Start(":8080"))
}
