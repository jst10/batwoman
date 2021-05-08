package api

import (
	"github.com/gorilla/mux"
	"made.by.jst10/celtra/batwoman/cmd/auth"
	"made.by.jst10/celtra/batwoman/cmd/custom_errors"
	"made.by.jst10/celtra/batwoman/cmd/database"
	"made.by.jst10/celtra/batwoman/cmd/structs"
	"net/http"
	"strconv"
)

func createUser(w http.ResponseWriter, r *http.Request) {
	userController := database.User{}
	var authDataFromRequestBody structs.AuthData
	err := decodeJSONBody(w, r, &authDataFromRequestBody)
	if err != nil {
		respondError(w, http.StatusUnprocessableEntity, err)
		return
	}
	_, err = userController.GetByUsername(authDataFromRequestBody.Username)
	if err == nil {
		respondError(w, http.StatusBadRequest, custom_errors.GetNotValidDataError("user by that username already exists"))
		return
	}
	user := &database.User{}
	user.Username = authDataFromRequestBody.Username
	user.Salt, err = auth.CreateSalt()
	user.Password = auth.CreateHash(authDataFromRequestBody.Password, user.Salt)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err)
		return
	}
	createdUser, err := user.Create()
	if err != nil {
		respondError(w, http.StatusInternalServerError, err)
		return
	}
	respondJSON(w, http.StatusCreated, createdUser)
}

func getMe(w http.ResponseWriter, r *http.Request) {
	tokenData := r.Context().Value("token").(*structs.TokenData)
	userController := database.User{}
	user, err := userController.GetByID(tokenData.UserId)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err)
		return
	}
	respondJSON(w, http.StatusOK, user)
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

func getUser(w http.ResponseWriter, r *http.Request) {
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
	tokenData := r.Context().Value("token").(*structs.TokenData)
	if tokenData.UserId != userFromRequestParam.ID {
		respondError(w, http.StatusForbidden, custom_errors.GetForbiddenError("update user"))
		return
	}
	var authDataFromRequestBody structs.AuthData
	err = decodeJSONBody(w, r, &authDataFromRequestBody)
	if err != nil {
		respondError(w, http.StatusUnprocessableEntity, err)
		return
	}
	userFromRequestParam.Username = authDataFromRequestBody.Username
	userFromRequestParam.Password = auth.CreateHash(authDataFromRequestBody.Password, userFromRequestParam.Salt)
	updatedUser, err := userFromRequestParam.Update(userFromRequestParam.ID)
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
	tokenData := r.Context().Value("token").(*structs.TokenData)
	if tokenData.UserId != oldUser.ID {
		respondError(w, http.StatusForbidden, custom_errors.GetForbiddenError("update user"))
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
