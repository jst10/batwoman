package api

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"made.by.jst10/celtra/batwoman/cmd/database"
	"net/http"
	"strconv"
)

func createUser(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respondError(w, http.StatusUnprocessableEntity, err)
	}
	user := database.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		respondError(w, http.StatusUnprocessableEntity, err)
		return
	}

	userCreated, err := user.SaveUser()

	if err != nil {
		respondError(w, http.StatusInternalServerError, err)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, userCreated.ID))
	respondJSON(w, http.StatusCreated, userCreated)
}
func getUsers(w http.ResponseWriter, r *http.Request) {

	user := database.User{}

	users, err := user.FindAllUsers()
	if err != nil {
		respondError(w, http.StatusInternalServerError, err)
		return
	}
	respondJSON(w, http.StatusOK, users)
}
func GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		respondError(w, http.StatusBadRequest, err)
		return
	}
	user := database.User{}
	userGotten, err := user.FindUserByID(uint32(uid))
	if err != nil {
		respondError(w, http.StatusBadRequest, err)
		return
	}
	respondJSON(w, http.StatusOK, userGotten)
}
func updateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		respondError(w, http.StatusBadRequest, err)
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respondError(w, http.StatusUnprocessableEntity, err)
		return
	}
	user := database.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		respondError(w, http.StatusUnprocessableEntity, err)
		return
	}
	if err != nil {
		respondError(w, http.StatusUnprocessableEntity, err)
		return
	}
	updatedUser, err := user.UpdateAUser(uint32(uid))
	if err != nil {
		respondError(w, http.StatusInternalServerError, err)
		return
	}
	respondJSON(w, http.StatusOK, updatedUser)
}
func deleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user := database.User{}

	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		respondError(w, http.StatusBadRequest, err)
		return
	}
	_, err = user.DeleteAUser(uint32(uid))
	if err != nil {
		respondError(w, http.StatusInternalServerError, err)
		return
	}
	w.Header().Set("Entity", fmt.Sprintf("%d", uid))
	respondJSON(w, http.StatusNoContent, "")
}
