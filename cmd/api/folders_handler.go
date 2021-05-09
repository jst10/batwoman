package api

import (
	"fmt"
	"github.com/gorilla/mux"
	"made.by.jst10/celtra/batwoman/cmd/custom_errors"
	"made.by.jst10/celtra/batwoman/cmd/database"
	"made.by.jst10/celtra/batwoman/cmd/structs"
	"net/http"
	"strconv"
)

func createFolder(w http.ResponseWriter, r *http.Request) {
	var folderData database.Folder
	err := decodeJSONBody(w, r, &folderData)
	if err != nil {
		respondError(w, http.StatusUnprocessableEntity, err)
		return
	}
	tokenData := r.Context().Value("token").(*structs.TokenData)
	folder := &database.Folder{}
	folder.Name = folderData.Name
	folder.OwnerID = tokenData.UserId
	createdFolder, err := folder.Create()
	if err != nil {
		respondError(w, http.StatusInternalServerError, err)
		return
	}
	respondJSON(w, http.StatusCreated, createdFolder)
}

func getFolders(w http.ResponseWriter, r *http.Request) {
	folderController := database.Folder{}
	tokenData := r.Context().Value("token").(*structs.TokenData)
	queryOptions := &database.FolderQueryOptions{}
	query := r.URL.Query()
	queryOptions.OwnerId = fmt.Sprint(tokenData.UserId)
	queryOptions.Page = getPageFromQuery(query)
	queryOptions.PageSize = getPageSizeFromQuery(query)
	queryOptions.Name = query.Get("name")
	queryOptions.OrderBy, queryOptions.OrderDirection = getOrderFromQuery(query, []string{"name", "id"})
	pageOfFolders, err := folderController.List(queryOptions)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err)
		return
	}
	respondJSON(w, http.StatusOK, pageOfFolders)
}

func getFolder(w http.ResponseWriter, r *http.Request) {
	folder, err := getFolderFromRequest(r)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err)
		return
	}
	tokenData := r.Context().Value("token").(*structs.TokenData)
	if tokenData.UserId != folder.OwnerID {
		respondError(w, http.StatusForbidden, custom_errors.GetForbiddenError("get folder"))
		return
	}
	respondJSON(w, http.StatusOK, folder)
}

func updateFolder(w http.ResponseWriter, r *http.Request) {
	oldFolder, err := getFolderFromRequest(r)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err)
		return
	}
	tokenData := r.Context().Value("token").(*structs.TokenData)
	if tokenData.UserId != oldFolder.OwnerID {
		respondError(w, http.StatusForbidden, custom_errors.GetForbiddenError("update folder"))
		return
	}
	var folderData database.Folder
	err = decodeJSONBody(w, r, &folderData)
	if err != nil {
		respondError(w, http.StatusUnprocessableEntity, err)
		return
	}
	oldFolder.Name = folderData.Name
	updatedFolder, err := oldFolder.Update(oldFolder.ID)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err)
		return
	}
	respondJSON(w, http.StatusOK, updatedFolder)
}

func deleteFolder(w http.ResponseWriter, r *http.Request) {
	oldFolder, err := getFolderFromRequest(r)
	if err != nil {
		respondError(w, http.StatusNotFound, err)
		return
	}
	tokenData := r.Context().Value("token").(*structs.TokenData)
	if tokenData.UserId != oldFolder.OwnerID {
		respondError(w, http.StatusForbidden, custom_errors.GetForbiddenError("update folder"))
		return
	}
	folderController := database.Folder{}
	_, err = folderController.DeleteById(oldFolder.ID)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err)
		return
	}
	respondJSON(w, http.StatusNoContent, "")
}
func getFolderFromRequest(r *http.Request) (*database.Folder, *custom_errors.CustomError) {
	vars := mux.Vars(r)
	uid, parsingError := strconv.ParseUint(vars["id"], 10, 32)
	if parsingError != nil {
		return nil, custom_errors.GetParsingError(parsingError, "id")
	}
	folderController := database.Folder{}
	folder, err := folderController.GetByID(uint(uid))
	if err != nil {
		return nil, err
	}
	return folder, nil
}
