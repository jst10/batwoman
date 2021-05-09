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

func createNote(w http.ResponseWriter, r *http.Request) {
	var noteData database.Note
	err := decodeJSONBody(w, r, &noteData)
	if err != nil {
		respondError(w, http.StatusUnprocessableEntity, err)
		return
	}
	tokenData := r.Context().Value("token").(*structs.TokenData)

	folderController := database.Folder{}
	folder, err := folderController.GetByID(noteData.FolderID)
	if err != nil {
		respondError(w, http.StatusForbidden, custom_errors.GetForbiddenError("get folder of not"))
	}

	if folder.OwnerID != tokenData.UserId {
		respondError(w, http.StatusForbidden, custom_errors.GetForbiddenError("create note in folder"))
		return
	}

	note := &database.Note{}
	note.Name = noteData.Name
	note.IsListNote = noteData.IsListNote
	note.IsPublic = noteData.IsPublic
	note.NoteBodies = noteData.NoteBodies
	note.FolderID = noteData.FolderID
	note.OwnerID = tokenData.UserId
	createdNote, err := note.Create()
	if err != nil {
		respondError(w, http.StatusInternalServerError, err)
		return
	}
	respondJSON(w, http.StatusCreated, createdNote)
}

func getNotes(w http.ResponseWriter, r *http.Request) {
	noteController := database.Note{}
	tokenData := r.Context().Value("token").(*structs.TokenData)
	queryOptions := &database.NoteQueryOptions{}
	query := r.URL.Query()
	queryOptions.OwnerId = fmt.Sprint(tokenData.UserId)
	queryOptions.Page = getPageFromQuery(query)
	queryOptions.PageSize = getPageSizeFromQuery(query)
	queryOptions.Name = query.Get("name")
	queryOptions.SharedOption = query.Get("is_public")
	queryOptions.FolderId = query.Get("folder_id")
	queryOptions.OrPublic = true
	queryOptions.OrderBy, queryOptions.OrderDirection = getOrderFromQuery(query, []string{"name", "id", "is_public"})
	pageOfNotes, err := noteController.List(queryOptions)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err)
		return
	}
	respondJSON(w, http.StatusOK, pageOfNotes)
}

func getNote(w http.ResponseWriter, r *http.Request) {
	note, err := getNoteFromRequest(r)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err)
		return
	}
	tokenData := r.Context().Value("token").(*structs.TokenData)
	if tokenData.UserId != note.OwnerID {
		respondError(w, http.StatusForbidden, custom_errors.GetForbiddenError("get note"))
		return
	}
	respondJSON(w, http.StatusOK, note)
}

func updateNote(w http.ResponseWriter, r *http.Request) {
	oldNote, err := getNoteFromRequest(r)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err)
		return
	}
	tokenData := r.Context().Value("token").(*structs.TokenData)
	if tokenData.UserId != oldNote.OwnerID {
		respondError(w, http.StatusForbidden, custom_errors.GetForbiddenError("update note"))
		return
	}
	var noteData database.Note
	err = decodeJSONBody(w, r, &noteData)
	if err != nil {
		respondError(w, http.StatusUnprocessableEntity, err)
		return
	}
	oldNote.Name = noteData.Name
	oldNote.IsPublic = noteData.IsPublic
	oldNote.NoteBodies = noteData.NoteBodies
	updatedNote, err := oldNote.Update(oldNote.ID)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err)
		return
	}
	respondJSON(w, http.StatusOK, updatedNote)
}

func deleteNote(w http.ResponseWriter, r *http.Request) {
	oldNote, err := getNoteFromRequest(r)
	if err != nil {
		respondError(w, http.StatusNotFound, err)
		return
	}
	tokenData := r.Context().Value("token").(*structs.TokenData)
	if tokenData.UserId != oldNote.OwnerID {
		respondError(w, http.StatusForbidden, custom_errors.GetForbiddenError("update note"))
		return
	}
	noteController := database.Note{}
	_, err = noteController.DeleteById(oldNote.ID)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err)
		return
	}
	respondJSON(w, http.StatusNoContent, "")
}
func getNoteFromRequest(r *http.Request) (*database.Note, *custom_errors.CustomError) {
	vars := mux.Vars(r)
	uid, parsingError := strconv.ParseUint(vars["id"], 10, 32)
	if parsingError != nil {
		return nil, custom_errors.GetParsingError(parsingError, "id")
	}
	noteController := database.Note{}
	note, err := noteController.GetByID(uint(uid))
	if err != nil {
		return nil, err
	}
	return note, nil
}
