package api

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

const tokenCookieName = "token"
const refreshCookieName = "refresh_token"

func StartApi() {
	myRouter := mux.NewRouter().StrictSlash(true)
	// Auth routes
	myRouter.HandleFunc("/api/auth", authenticate).Methods("POST")
	myRouter.HandleFunc("/api/auth", reAuthenticate).Methods("PUT")
	myRouter.HandleFunc("/api/auth", deAuthenticate).Methods("DELETE")

	//Users routes
	myRouter.HandleFunc("/api/users", createUser).Methods("POST")
	myRouter.HandleFunc("/api/users", setMiddlewareAuthentication(getUsers)).Methods("GET")
	myRouter.HandleFunc("/api/users/me", setMiddlewareAuthentication(getMe)).Methods("GET")
	myRouter.HandleFunc("/api/users/{id}", setMiddlewareAuthentication(getUser)).Methods("GET")
	myRouter.HandleFunc("/api/users/{id}", setMiddlewareAuthentication(updateUser)).Methods("PUT")
	myRouter.HandleFunc("/api/users/{id}", setMiddlewareAuthentication(deleteUser)).Methods("DELETE")

	//Folder routes
	myRouter.HandleFunc("/api/folders", setMiddlewareAuthentication(createFolder)).Methods("POST")
	myRouter.HandleFunc("/api/folders", setMiddlewareAuthentication(getFolders)).Methods("GET")
	myRouter.HandleFunc("/api/folders/{id}", setMiddlewareAuthentication(getFolder)).Methods("GET")
	myRouter.HandleFunc("/api/folders/{id}", setMiddlewareAuthentication(updateFolder)).Methods("PUT")
	myRouter.HandleFunc("/api/folders/{id}", setMiddlewareAuthentication(deleteFolder)).Methods("DELETE")

	//Note routes
	myRouter.HandleFunc("/api/notes", setMiddlewareAuthentication(createNote)).Methods("POST")
	myRouter.HandleFunc("/api/notes", setMiddlewareAuthentication(getNotes)).Methods("GET")
	myRouter.HandleFunc("/api/notes/{id}", setMiddlewareAuthentication(getNote)).Methods("GET")
	myRouter.HandleFunc("/api/notes/{id}", setMiddlewareAuthentication(updateNote)).Methods("PUT")
	myRouter.HandleFunc("/api/notes/{id}", setMiddlewareAuthentication(deleteNote)).Methods("DELETE")

	//Note routes
	myRouter.HandleFunc("/api/notes/{id}/collaboratives", setMiddlewareAuthentication(getNoteCollaboratives)).Methods("GET")
	myRouter.HandleFunc("/api/notes/{id}/collaboratives", setMiddlewareAuthentication(addNoteCollaborative)).Methods("POST")
	myRouter.HandleFunc("/api/notes/{id}/collaboratives", setMiddlewareAuthentication(removeNoteCollaborative)).Methods("DELETE")

	fmt.Println("api ok")
	http.ListenAndServe(":8888", myRouter)
}
