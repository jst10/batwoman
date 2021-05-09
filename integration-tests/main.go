package main

import (
	"fmt"
	"log"
	"made.by.jst10/celtra/batwoman/cmd/database"
	"made.by.jst10/celtra/batwoman/cmd/structs"
	"net/http"
	"net/http/cookiejar"
)

var BASE_URL = "http://localhost:8888/api"
var user1Data = &structs.AuthData{Username: "user1", Password: "user1"}
var user2Data = &structs.AuthData{Username: "user2", Password: "user2"}
var user3Data = &structs.AuthData{Username: "user3", Password: "user3"}

var folder1Data = &database.Folder{Name: "folder1"}
var folder2Data = &database.Folder{Name: "folder2"}
var folder3Data = &database.Folder{Name: "folder3"}

var note1Data = &database.Note{
	Name:       "note1",
	IsPublic:   false,
	IsListNote: true,
	NoteBodies: []database.NoteBody{database.NoteBody{Text: "sdsadsad"}}}

var note2Data = &database.Note{
	Name:       "note2",
	IsPublic:   false,
	IsListNote: true,
	NoteBodies: []database.NoteBody{
		database.NoteBody{Text: "note 21"},
		database.NoteBody{Text: "note 22"}}}

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
	testFolders()
	testNotes()
	fmt.Println("Finish all tests")
}

func testUsers() {
	fmt.Println("Start users tests")
	client1 := initClient()
	client2 := initClient()
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
func testFolders() {
	fmt.Println("Start folder tests")
	client1 := initClient()
	client2 := initClient()
	createUser(client1, user1Data, http.StatusCreated)
	createUser(client2, user2Data, http.StatusCreated)
	authenticate(client1, user1Data)
	authenticate(client2, user2Data)
	folder1 := createFolder(client1, folder1Data, http.StatusCreated)
	folder2 := createFolder(client2, folder2Data, http.StatusCreated)
	updateFolder(client2, folder1.ID, folder2Data, http.StatusForbidden)
	folder1 = updateFolder(client1, folder1.ID, folder2Data, http.StatusOK)
	getFolders(client1, 1)
	getFolders(client2, 1)
	deleteFolder(client2, folder1.ID, http.StatusForbidden)
	deleteFolder(client1, folder1.ID, http.StatusNoContent)
	deleteFolder(client2, folder2.ID, http.StatusNoContent)
	deleteMe(client1)
	deleteMe(client2)
	fmt.Println("Finish testing folders")
}

func testNotes() {
	fmt.Println("Start notes tests")
	client1 := initClient()
	createUser(client1, user1Data, http.StatusCreated)
	authenticate(client1, user1Data)
	folder1 := createFolder(client1, folder1Data, http.StatusCreated)
	getFolders(client1, 1)
	note1Data.FolderID = folder1.ID
	note1 := createNote(client1, note1Data, http.StatusCreated)
	getNotes(client1, 1)
	updateNote(client1, note1.ID, note2Data, http.StatusOK)
	getNotes(client1, 1)
	deleteMe(client1)
	fmt.Println("Finish testing notes")
}
