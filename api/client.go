package api

import (
	"encoding/json"
	"net/http"

	"go-api-cli/models"
)

const baseURL = "https://jsonplaceholder.typicode.com"

func FetchUsers() []models.User {

	resp, err := http.Get(baseURL + "/users")
	if err != nil {
		return nil
	}
	defer resp.Body.Close()

	var users []models.User
	json.NewDecoder(resp.Body).Decode(&users)

	return users
}