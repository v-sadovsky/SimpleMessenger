package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"v-sadovsky/gateway/internal"
	"v-sadovsky/gateway/server/models"
)

const host = "http://localhost:8080"

type GatewayClient struct {
	CreateUserRequest *models.CreateUserRequest
	UpdateUserRequest *models.UpdateUserRequest
}

func NewGatewayClient() *GatewayClient {
	return &GatewayClient{}
}

func (gc *GatewayClient) createUser() (*models.CreateUserResponse, error) {
	const uri = "/messenger/v1/profiles"

	endpoint := fmt.Sprintf("%s%s", host, uri)
	status, respBody, err := internal.Do(http.MethodPost, endpoint, nil, gc.CreateUserRequest)
	if err != nil {
		return nil, err
	}

	if status == http.StatusBadRequest {
		var errorMsg models.ErrorMessage
		if err := json.NewDecoder(bytes.NewReader(respBody)).Decode(&errorMsg); err != nil {
			return nil, fmt.Errorf("user not created: status code %d", status)
		}

		return nil, fmt.Errorf("user not created: status code %d, error: %s", status, errorMsg.Message)
	}

	if status == http.StatusConflict {
		var errorMsg models.ErrorMessage
		if err := json.NewDecoder(bytes.NewReader(respBody)).Decode(&errorMsg); err != nil {
			return nil, fmt.Errorf("user not created: status code %d", status)
		}

		return nil, fmt.Errorf("user not created: status code %d, error: %s", status, errorMsg.Message)
	}

	if status == http.StatusInternalServerError {
		var errorMsg models.ErrorMessage
		if err := json.NewDecoder(bytes.NewReader(respBody)).Decode(&errorMsg); err != nil {
			return nil, fmt.Errorf("user not created: status code %d", status)
		}

		return nil, fmt.Errorf("user not created: status code %d, error: %s", status, errorMsg.Message)
	}

	if status == http.StatusCreated {
		response := new(models.CreateUserResponse)
		if err := json.NewDecoder(bytes.NewReader(respBody)).Decode(response); err != nil {
			return nil, fmt.Errorf("can't unmarshal reponse: %s", err.Error())
		}
		return response, nil
	}

	return nil, fmt.Errorf("unknown status code: %d", status)
}

func (gc *GatewayClient) getUser(id int) (*models.GetUserResponse, error) {
	uri := fmt.Sprintf("/messenger/v1/profiles/%d", id)

	endpoint := fmt.Sprintf("%s%s", host, uri)
	status, respBody, err := internal.Do(http.MethodGet, endpoint, nil, nil)
	if err != nil {
		return nil, err
	}

	if status == http.StatusNoContent {
		return &models.GetUserResponse{}, nil
	}

	if status == http.StatusInternalServerError {
		var errorMsg models.ErrorMessage
		if err := json.NewDecoder(bytes.NewReader(respBody)).Decode(&errorMsg); err != nil {
			return nil, fmt.Errorf("user not retrieved: status code %d", status)
		}

		return nil, fmt.Errorf("user not retrieved: status code %d, error: %s", status, errorMsg.Message)
	}

	if status == http.StatusOK {
		response := new(models.GetUserResponse)
		if err := json.NewDecoder(bytes.NewReader(respBody)).Decode(response); err != nil {
			return nil, fmt.Errorf("can't unmarshal reponse: %s", err.Error())
		}
		return response, nil
	}

	return nil, fmt.Errorf("unknown status code: %d", status)
}

func (gc *GatewayClient) updateUser(id int) (*models.UpdateUserResponse, error) {
	uri := fmt.Sprintf("/messenger/v1/profiles/%d", id)

	endpoint := fmt.Sprintf("%s%s", host, uri)
	status, respBody, err := internal.Do(http.MethodPut, endpoint, nil, gc.UpdateUserRequest)
	if err != nil {
		return nil, err
	}

	if status == http.StatusNoContent {
		return &models.UpdateUserResponse{}, nil
	}

	if status == http.StatusForbidden {
		var errorMsg models.ErrorMessage
		if err := json.NewDecoder(bytes.NewReader(respBody)).Decode(&errorMsg); err != nil {
			return nil, fmt.Errorf("user not updated: status code %d", status)
		}

		return nil, fmt.Errorf("user not updated: status code %d, error: %s", status, errorMsg.Message)
	}

	if status == http.StatusInternalServerError {
		var errorMsg models.ErrorMessage
		if err := json.NewDecoder(bytes.NewReader(respBody)).Decode(&errorMsg); err != nil {
			return nil, fmt.Errorf("user not updated: status code %d", status)
		}

		return nil, fmt.Errorf("user not updated: status code %d, error: %s", status, errorMsg.Message)
	}

	if status == http.StatusOK {
		response := new(models.UpdateUserResponse)
		if err := json.NewDecoder(bytes.NewReader(respBody)).Decode(response); err != nil {
			return nil, fmt.Errorf("can't unmarshal reponse: %s", err.Error())
		}
		return response, nil
	}

	return nil, fmt.Errorf("unknown status code: %d", status)
}

func (gc *GatewayClient) deleteUser(id int) (*models.GetUserResponse, error) {
	uri := fmt.Sprintf("/messenger/v1/profiles/%d", id)

	endpoint := fmt.Sprintf("%s%s", host, uri)
	status, respBody, err := internal.Do(http.MethodDelete, endpoint, nil, nil)
	if err != nil {
		return nil, err
	}

	if status == http.StatusNoContent {
		return &models.GetUserResponse{}, nil
	}

	if status == http.StatusForbidden {
		var errorMsg models.ErrorMessage
		if err := json.NewDecoder(bytes.NewReader(respBody)).Decode(&errorMsg); err != nil {
			return nil, fmt.Errorf("user not deleted: status code %d", status)
		}

		return nil, fmt.Errorf("user not deleted: status code %d, error: %s", status, errorMsg.Message)
	}

	if status == http.StatusInternalServerError {
		var errorMsg models.ErrorMessage
		if err := json.NewDecoder(bytes.NewReader(respBody)).Decode(&errorMsg); err != nil {
			return nil, fmt.Errorf("user not deleted: status code %d", status)
		}

		return nil, fmt.Errorf("user not deleted: status code %d, error: %s", status, errorMsg.Message)
	}

	if status == http.StatusOK {
		response := new(models.GetUserResponse)
		if err := json.NewDecoder(bytes.NewReader(respBody)).Decode(response); err != nil {
			return nil, fmt.Errorf("can't unmarshal reponse: %s", err.Error())
		}
		return response, nil
	}

	return nil, fmt.Errorf("unknown status code: %d", status)
}

func newString(s string) *string {
	return &s
}

func main() {

	client := NewGatewayClient()

	// Create user
	//client.CreateUserRequest = &models.CreateUserRequest{
	//	Email:    newString("tet@mail.ru"),
	//	Name:     newString("test_user"),
	//	Password: newString("1234"),
	//}
	//
	//resp, err := client.createUser()
	//if err != nil {
	//	fmt.Printf("createUser error: %+v\n", err)
	//	return
	//}

	// Get user
	//resp, err := client.getUser(1)
	//if err != nil {
	//	fmt.Printf("getUser error: %+v\n", err)
	//	return
	//}

	// Update user
	//client.UpdateUserRequest = &models.UpdateUserRequest{
	//	Email:    "tet@mail.ru",
	//	Name:     "test_user",
	//	Password: "1234",
	//}
	//
	//resp, err := client.updateUser(3)
	//if err != nil {
	//	fmt.Printf("updateUser error: %+v\n", err)
	//	return
	//}

	// Delete user
	resp, err := client.deleteUser(1)
	if err != nil {
		fmt.Printf("deleteUser error: %+v\n", err)
		return
	}

	fmt.Printf("success: %+v\n", resp)
}
