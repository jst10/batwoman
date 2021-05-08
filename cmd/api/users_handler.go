package api

import (
	"github.com/gorilla/mux"
	"made.by.jst10/celtra/batwoman/cmd/custom_errors"
	"made.by.jst10/celtra/batwoman/cmd/database"
	"net/http"
	"strconv"
)

func createUser(w http.ResponseWriter, r *http.Request) {
	userFromRequestBody := database.User{}
	err := decodeJSONBody(w, r, userFromRequestBody)
	if err != nil {
		respondError(w, http.StatusUnprocessableEntity, err)
		return
	}
	createdUser, err := userFromRequestBody.Create()
	if err != nil {
		respondError(w, http.StatusInternalServerError, err)
		return
	}
	respondJSON(w, http.StatusCreated, createdUser)
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	userController := database.User{}
	users, err := userController.All()
	if err != nil {
		respondError(w, http.StatusInternalServerError, err)
		return
	}
	respondJSON(w, http.StatusOK, users)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	user, err := getUserFromRequest(r)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err)
		return
	}
	respondJSON(w, http.StatusOK, user)
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	userFromRequestParam, err := getUserFromRequest(r)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err)
		return
	}
	userFromRequestBody := database.User{}
	err = decodeJSONBody(w, r, userFromRequestBody)
	if err != nil {
		respondError(w, http.StatusUnprocessableEntity, err)
		return
	}
	updatedUser, err := userFromRequestBody.Update(userFromRequestParam.ID)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err)
		return
	}
	respondJSON(w, http.StatusOK, updatedUser)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	oldUser, err := getUserFromRequest(r)
	if err != nil {
		respondError(w, http.StatusNotFound, err)
		return
	}
	userController := database.User{}
	_, err = userController.DeleteById(oldUser.ID)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err)
		return
	}
	respondJSON(w, http.StatusNoContent, "")
}
func getUserFromRequest(r *http.Request) (*database.User, *custom_errors.CustomError) {
	vars := mux.Vars(r)
	uid, parsingError := strconv.ParseUint(vars["id"], 10, 32)
	if parsingError != nil {
		return nil, custom_errors.GetParsingError(parsingError, "id")
	}
	userController := database.User{}
	user, err := userController.GetByID(uint(uid))
	if err != nil {
		return nil, err
	}
	return user, nil
}
