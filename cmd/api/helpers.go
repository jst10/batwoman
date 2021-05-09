package api

import (
	"encoding/json"
	"fmt"
	"made.by.jst10/celtra/batwoman/cmd/auth"
	"made.by.jst10/celtra/batwoman/cmd/custom_errors"
	"made.by.jst10/celtra/batwoman/cmd/structs"
	"net/http"
	"net/url"
	"strconv"
	"strings"
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



func getPageFromQuery(query url.Values) int {
	page, _ := strconv.Atoi(query.Get("page"))
	if page == 0 {
		page = 1
	}
	return page
}
func getPageSizeFromQuery(query url.Values) int {
	pageSize, _ := strconv.Atoi(query.Get("page_size"))
	switch {
	case pageSize > 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}
	return pageSize
}
func stringInSlice(text string, list []string) bool {
	for _, item := range list {
		if item == text {
			return true
		}
	}
	return false
}

func getOrderFromQuery(query url.Values, validFields []string) (string, string) {
	orderBy := query.Get("order_by")
	if len(orderBy) > 0 {
		var orderDirection string
		if strings.HasPrefix(orderBy, "-") {
			orderDirection = "DESC"
		} else {
			orderDirection = "ASC"
		}
		if strings.HasPrefix(orderBy, "-") || strings.HasPrefix(orderBy, "+") {
			orderBy = orderBy[1:]
		}
		if stringInSlice(orderBy, validFields) {
			return orderBy, orderDirection
		}
	}
	return "", ""
}
