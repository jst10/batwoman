package main

import (
	"fmt"
	"log"
	"made.by.jst10/celtra/batwoman/cmd/structs"
	"net/http"
	"net/http/cookiejar"
)

var BASE_URL = "http://localhost:8888/api"
var user1Data = &structs.AuthData{Username: "user1", Password: "user1"}
var user2Data = &structs.AuthData{Username: "user2", Password: "user2"}
var user3Data = &structs.AuthData{Username: "user3", Password: "user3"}

func initClient() *http.Client {
	options := cookiejar.Options{}
	jar, err := cookiejar.New(&options)
	if err != nil {
		log.Fatal("error creating client", err)
	}
	return &http.Client{Jar: jar}
}

func main() {
	fmt.Println("Start tests")
	testUsers()
	fmt.Println("Finish all tests")
}

func testUsers() {
	client1 := initClient()
	client2 := initClient()
	fmt.Println("Start users tests")
	createUser(client1, user1Data, http.StatusCreated)
	createUser(client2, user2Data, http.StatusCreated)
	createUser(client1, user1Data, http.StatusBadRequest)
	authenticate(client1, user1Data)
	authenticate(client2, user2Data)
	getUsers(client1, 2)
	getUsers(client2, 2)
	user1 := getMe(client1)
	user2 := getMe(client2)
	updateUser(client2, user1.ID, user3Data, http.StatusForbidden)
	updateUser(client1, user2.ID, user3Data, http.StatusForbidden)
	updateMe(client1, user3Data)
	deleteUser(client2, user1.ID, http.StatusForbidden)
	deleteUser(client1, user2.ID, http.StatusForbidden)
	deleteMe(client1)
	deleteMe(client2)
	fmt.Println("Finish testing users")
}
