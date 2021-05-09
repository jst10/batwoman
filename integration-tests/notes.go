package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"made.by.jst10/celtra/batwoman/cmd/database"
	"net/http"
)

func createNote(client *http.Client, note *database.Note, expectedStatusCode int) *database.Note {
	var jsonStr, _ = json.Marshal(note)
	url := fmt.Sprintf("%s/notes", BASE_URL)
	req := createARequestObject("POST", url, bytes.NewBuffer(jsonStr))
	resp := makeARequest(client, req, expectedStatusCode)
	if expectedStatusCode == http.StatusCreated {
		var dst = database.Note{}
		extractBody(resp.Body, &dst)
		return &dst
	}
	return nil
}

func getNotes(client *http.Client, expectedNumberOfNotes int) *database.PageOfNotes {
	url := fmt.Sprintf("%s/notes", BASE_URL)
	req := createARequestObject("GET", url, nil)
	resp := makeARequest(client, req, http.StatusOK)
	var dst = database.PageOfNotes{}
	extractBody(resp.Body, &dst)
	if len(dst.Items) != expectedNumberOfNotes {
		log.Fatal("Invalid number of received notes", len(dst.Items), expectedNumberOfNotes)
	}
	if dst.Count != expectedNumberOfNotes {
		log.Fatal("Invalid number of received notes", len(dst.Items), expectedNumberOfNotes)
	}
	return &dst
}

func updateNote(client *http.Client, noteId uint, noteData *database.Note, expectedStatusCode int) *database.Note {
	var jsonStr, _ = json.Marshal(noteData)
	url := fmt.Sprintf("%s/notes/%d", BASE_URL, noteId)
	req := createARequestObject("PUT", url, bytes.NewBuffer(jsonStr))
	resp := makeARequest(client, req, expectedStatusCode)
	if expectedStatusCode == http.StatusOK {
		var dst = database.Note{}
		extractBody(resp.Body, &dst)
		return &dst
	}
	return nil
}

func deleteNote(client *http.Client, noteId uint, expectedStatusCode int) {
	url := fmt.Sprintf("%s/notes/%d", BASE_URL, noteId)
	req := createARequestObject("DELETE", url, nil)
	makeARequest(client, req, expectedStatusCode)
}
