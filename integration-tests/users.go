package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"made.by.jst10/celtra/batwoman/cmd/database"
	"made.by.jst10/celtra/batwoman/cmd/structs"
	"net/http"
)

func createUser(client *http.Client, authData *structs.AuthData, expectedStatusCode int) {
	var jsonStr, _ = json.Marshal(authData)
	url := fmt.Sprintf("%s/users", BASE_URL)
	req := createARequestObject("POST", url, bytes.NewBuffer(jsonStr))
	makeARequest(client, req, expectedStatusCode)
}

func getUsers(client *http.Client, expectedNumberOfUsers int) []database.User {
	url := fmt.Sprintf("%s/users", BASE_URL)
	req := createARequestObject("GET", url, nil)
	resp := makeARequest(client, req, http.StatusOK)
	var dst = make([]database.User, 0)
	extractBody(resp.Body, &dst)
	if len(dst) != expectedNumberOfUsers {
		log.Fatal("Invalid number of received users", len(dst), expectedNumberOfUsers)
	}
	return dst
}

func getMe(client *http.Client) *database.User {
	url := fmt.Sprintf("%s/users/me", BASE_URL)
	req := createARequestObject("GET", url, nil)
	resp := makeARequest(client, req, http.StatusOK)
	var dst = database.User{}
	extractBody(resp.Body, &dst)
	return &dst
}

func updateMe(client *http.Client, userData *structs.AuthData) {
	user := getMe(client)
	updateUser(client, user.ID, userData, http.StatusOK)
}

func updateUser(client *http.Client, userId uint, userData *structs.AuthData, expectedStatusCode int) {
	var jsonStr, _ = json.Marshal(userData)
	url := fmt.Sprintf("%s/users/%d", BASE_URL, userId)
	req := createARequestObject("PUT", url, bytes.NewBuffer(jsonStr))
	makeARequest(client, req, expectedStatusCode)
}

func deleteMe(client *http.Client) {
	user := getMe(client)
	deleteUser(client, user.ID, http.StatusNoContent)
}

func deleteUser(client *http.Client, userId uint, expectedStatusCode int) {
	url := fmt.Sprintf("%s/users/%d", BASE_URL, userId)
	req := createARequestObject("DELETE", url, nil)
	makeARequest(client, req, expectedStatusCode)
}
