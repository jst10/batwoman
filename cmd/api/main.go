package api

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func StartApi() {
	myRouter := mux.NewRouter().StrictSlash(true)
	// Auth routes
	myRouter.HandleFunc("/api/auth", authenticate).Methods("POST")
	myRouter.HandleFunc("/api/auth", reAuthenticate).Methods("PUT")
	myRouter.HandleFunc("/api/auth", deAuthenticate).Methods("DELETE")

	//Users routes
	myRouter.HandleFunc("/users", createUser).Methods("POST")
	myRouter.HandleFunc("/users", getUsers).Methods("GET")
	myRouter.HandleFunc("/users/{id}", GetUser).Methods("GET")
	myRouter.HandleFunc("/users/{id}", updateUser).Methods("PUT")
	myRouter.HandleFunc("/users/{id}", deleteUser).Methods("DELETE")

	//Folder routes
	myRouter.HandleFunc("/folders", createFolder).Methods("POST")
	myRouter.HandleFunc("/folders", getFolders).Methods("GET")
	myRouter.HandleFunc("/folders/{id}", GetFolder).Methods("GET")
	myRouter.HandleFunc("/folders/{id}", updateFolder).Methods("PUT")
	myRouter.HandleFunc("/folders/{id}", deleteFolder).Methods("DELETE")

	//Note routes
	myRouter.HandleFunc("/notes", createNote).Methods("POST")
	myRouter.HandleFunc("/notes", getNotes).Methods("GET")
	myRouter.HandleFunc("/notes/{id}", GetNote).Methods("GET")
	myRouter.HandleFunc("/notes/{id}", updateNote).Methods("PUT")
	myRouter.HandleFunc("/notes/{id}", deleteNote).Methods("DELETE")

	//Note routes
	myRouter.HandleFunc("/notes/{id}/collaboratives", getNoteCollaboratives).Methods("GET")
	myRouter.HandleFunc("/notes/{id}/collaboratives", addNoteCollaborative).Methods("POST")
	myRouter.HandleFunc("/notes/{id}/collaboratives", removeNoteCollaborative).Methods("DELETE")

	fmt.Println("api ok")
	http.ListenAndServe(":8888", myRouter)
}
