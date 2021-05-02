package api

import (
	"made.by.jst10/celtra/batwoman/cmd/auth"
	"made.by.jst10/celtra/batwoman/cmd/structs"
	"net/http"
)

func authenticate(w http.ResponseWriter, r *http.Request) {
	var authData structs.AuthData
	err := decodeJSONBody(w, r, &authData)
	if err != nil {
		respondError(w, http.StatusBadRequest, err)
		return
	}
	tokenWrapper, refreshTokenWrapper, err := auth.AuthenticateUser(&authData)
	if err != nil {
		respondError(w, http.StatusBadRequest, err)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:    tokenCookieName,
		Value:   tokenWrapper.Token,
		Expires: *tokenWrapper.Expiration,
	})

	http.SetCookie(w, &http.Cookie{
		Name:    refreshCookieName,
		Value:   refreshTokenWrapper.Token,
		Expires: *refreshTokenWrapper.Expiration,
	})
	respondJSON(w, http.StatusOK, structs.TokensResponse{Token: tokenWrapper.Token, RefreshToken: refreshTokenWrapper.Token})
}
func reAuthenticate(w http.ResponseWriter, r *http.Request) {
	cookie, err := getCookieFromRequest(r, refreshCookieName)
	if err != nil {
		respondError(w, http.StatusUnauthorized, err)
		return
	}
	refreshToken := cookie.Value
	tokenData, err := auth.VerifyJWRT(refreshToken)
	if err != nil {
		respondError(w, http.StatusUnauthorized, err)
		return
	}
	tokenWrapper, err := auth.ReAuthenticateUser(tokenData.SessionId)
	if err != nil {
		respondError(w, http.StatusBadRequest, err)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:    tokenCookieName,
		Value:   tokenWrapper.Token,
		Expires: *tokenWrapper.Expiration,
	})
	respondJSON(w, http.StatusOK, structs.TokensResponse{Token: tokenWrapper.Token})
}

func deAuthenticate(w http.ResponseWriter, r *http.Request) {
	tokenData, err := authenticateRequest(r)
	if err != nil {
		respondError(w, http.StatusUnauthorized, err)
		return
	}
	err = auth.DeAuthenticateUser(tokenData.UserId)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:   tokenCookieName,
		Value:  "",
		MaxAge: -1,
	})

	http.SetCookie(w, &http.Cookie{
		Name:   refreshCookieName,
		Value:  "",
		MaxAge: -1,
	})
	respondEmpty(w)
}
