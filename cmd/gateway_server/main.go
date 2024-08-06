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

//func main() {
//	http.HandleFunc("/messenger/v1/profiles", createUser)
//
//	log.Println("Starting Gateway service on port :8080")
//	log.Fatal(http.ListenAndServe(":8080", nil))
//}
//
//func createUser(w http.ResponseWriter, r *http.Request) {
//	data, err := io.ReadAll(r.Body)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	fmt.Println(string(data))
//
//	var request models.CreateUserRequest
//	if err := json.NewDecoder(bytes.NewReader(data)).Decode(&request); err != nil {
//		fmt.Println(err)
//		http.Error(w, err.Error(), http.StatusBadRequest)
//	}
//
//	if err := request.Validate(strfmt.Default); err != nil {
//		http.Error(w, err.Error(), http.StatusBadRequest)
//	}
//
//	// Business logic
//	var userID int64 = 3
//	//fmt.Printf("create user: %+v\n", request)
//
//	response := models.CreateUserResponse{ID: userID}
//
//	fmt.Println(response)
//}

func main() {
	e := echo.New()

	e.POST("/messenger/v1/profiles", createUser)
	e.GET("/messenger/v1/profiles/:id", getUser)
	e.PUT("/messenger/v1/profiles/:id", updateUser)
	e.DELETE("/messenger/v1/profiles/:id", deleteUser)

	e.Logger.Fatal(e.Start(":8080"))
}

//curl -X POST -H 'Content-Type: application/json' -d '{"name":"user_123","email":"user_123@mail.ru","password":"asTR3k!90d","photo":"https://image_path"}' http://localhost:8080/messenger/v1/profiles
