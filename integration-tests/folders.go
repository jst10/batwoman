package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"made.by.jst10/celtra/batwoman/cmd/database"
	"net/http"
)

func createFolder(client *http.Client, folder *database.Folder, expectedStatusCode int) *database.Folder {
	var jsonStr, _ = json.Marshal(folder)
	url := fmt.Sprintf("%s/folders", BASE_URL)
	req := createARequestObject("POST", url, bytes.NewBuffer(jsonStr))
	resp := makeARequest(client, req, expectedStatusCode)
	if expectedStatusCode == http.StatusCreated {
		var dst = database.Folder{}
		extractBody(resp.Body, &dst)
		return &dst
	}
	return nil
}

func getFolders(client *http.Client, expectedNumberOfFolders int) *database.PageOfFolders {
	url := fmt.Sprintf("%s/folders", BASE_URL)
	req := createARequestObject("GET", url, nil)
	resp := makeARequest(client, req, http.StatusOK)
	var dst = database.PageOfFolders{}
	extractBody(resp.Body, &dst)
	if len(dst.Items) != expectedNumberOfFolders {
		log.Fatal("Invalid number of received folders", len(dst.Items), expectedNumberOfFolders)
	}
	if dst.Count != expectedNumberOfFolders {
		log.Fatal("Invalid number of received folders", len(dst.Items), expectedNumberOfFolders)
	}
	return &dst
}

func updateFolder(client *http.Client, folderId uint, folderData *database.Folder, expectedStatusCode int) *database.Folder {
	var jsonStr, _ = json.Marshal(folderData)
	url := fmt.Sprintf("%s/folders/%d", BASE_URL, folderId)
	req := createARequestObject("PUT", url, bytes.NewBuffer(jsonStr))
	resp := makeARequest(client, req, expectedStatusCode)
	if expectedStatusCode == http.StatusOK {
		var dst = database.Folder{}
		extractBody(resp.Body, &dst)
		return &dst
	}
	return nil
}

func deleteFolder(client *http.Client, folderId uint, expectedStatusCode int) {
	url := fmt.Sprintf("%s/folders/%d", BASE_URL, folderId)
	req := createARequestObject("DELETE", url, nil)
	makeARequest(client, req, expectedStatusCode)
}
