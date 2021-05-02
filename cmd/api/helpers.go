package api

import (
	"encoding/json"
	"fmt"
	"made.by.jst10/celtra/batwoman/cmd/auth"
	"made.by.jst10/celtra/batwoman/cmd/custom_errors"
	"made.by.jst10/celtra/batwoman/cmd/structs"
	"net/http"
)

func respondJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}
}
func respondEmpty(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNoContent)
}


func respondError(w http.ResponseWriter, statusCode int, err interface{}) { //*custom_errors.CustomError
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(err)
	fmt.Println("Error")
	fmt.Println(err)
	//fmt.Println(err.StackTrace)
	//fmt.Println(err.OriginalError)
}
func getCookieFromRequest(r *http.Request, cookieName string) (*http.Cookie, *custom_errors.CustomError) {
	cookie, err := r.Cookie(cookieName)
	if err != nil {
		return nil, custom_errors.GetErrorGettingCookie(err, cookieName)
	} else {
		return cookie, nil
	}
}

func authenticateRequest( r *http.Request) (*structs.TokenData, *custom_errors.CustomError) {
	cookie, err := getCookieFromRequest(r, tokenCookieName)
	if err != nil {
		return nil, custom_errors.GetCookieNotPresentError(tokenCookieName)
	}
	token := cookie.Value
	tokenData, err := auth.VerifyJWT(token)
	if err != nil {
		return nil, err
	}
	return tokenData, nil
}