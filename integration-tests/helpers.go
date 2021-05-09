package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"made.by.jst10/celtra/batwoman/cmd/structs"
	"net/http"
)

func createARequestObject(method, url string, body io.Reader) *http.Request {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		log.Fatal("Error creating a request", err)
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	return req
}
func makeARequest(client *http.Client, req *http.Request, expectedStatusCode int) *http.Response {
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error at making request", err)
	}
	if resp.StatusCode != expectedStatusCode {
		log.Fatal("Invalid status ", resp.Status, expectedStatusCode)
	}
	return resp
}

func extractBody(body io.ReadCloser, dst interface{}) {
	dec := json.NewDecoder(body)
	dec.DisallowUnknownFields()
	err := dec.Decode(dst)
	if err != nil {
		log.Fatal("Error at decoding data", err)
	}
	body.Close()
}
func authenticate(client *http.Client, user *structs.AuthData) *structs.TokensResponse {
	var jsonStr, _ = json.Marshal(user)
	req, err := http.NewRequest("POST", "http://localhost:8888/api/auth", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		log.Fatal("Error authenticating user", err)
	}
	resp := makeARequest(client, req, http.StatusOK)
	dst := &structs.TokensResponse{}
	extractBody(resp.Body, dst)
	return dst
}

func assertEqualInt(fist int, second int) {
	if fist != second {
		log.Fatal("Not equal", fist, second)
	}
}

func assertEqualUint(fist uint, second uint) {
	if fist != second {
		log.Fatal("Not equal", fist, second)
	}
}
func assertEqualString(fist string, second string) {
	if fist != second {
		log.Fatal("Not equal", fist, second)
	}
}
