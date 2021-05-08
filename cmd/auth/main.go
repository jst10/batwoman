package auth

import (
	"made.by.jst10/celtra/batwoman/cmd/custom_errors"
	"made.by.jst10/celtra/batwoman/cmd/database"
	"made.by.jst10/celtra/batwoman/cmd/structs"
)

func AuthenticateUser(authData *structs.AuthData) (*structs.TokenWrapper, *structs.TokenWrapper, *custom_errors.CustomError) {
	userController := database.User{}
	user, err := userController.GetByUsername(authData.Username)
	if err != nil {
		return nil, nil, err.AST("authenticate user")
	}
	if !doPasswordsMatch(user, authData.Password) {
		return nil, nil, custom_errors.GetNotValidDataError("invalid auth data")
	}
	session := &database.Session{UserID: user.ID}
	session, err = session.Create()
	if err != nil {
		return nil, nil, err.AST("authenticate user")
	}
	tokenData := structs.TokenData{
		UserId:    user.ID,
		CreatedAt: user.CreatedAt,
		Username:  user.Username,
		SessionId: session.ID}

	token, tokenExpiration, err := createJWT(&tokenData)
	if err != nil {
		return nil, nil, err.AST("authenticate user")
	}

	refreshToken, refreshTokenExpiration, err := createJWRT(&tokenData)
	if err != nil {
		return nil, nil, err.AST("authenticate user")
	}

	tokenWrapper := structs.TokenWrapper{User: user, Token: token, Expiration: tokenExpiration}
	refreshTokenWrapper := structs.TokenWrapper{User: user, Token: refreshToken, Expiration: refreshTokenExpiration}
	return &tokenWrapper, &refreshTokenWrapper, nil
}
func ReAuthenticateUser(sessionId uint) (*structs.TokenWrapper, *custom_errors.CustomError) {
	userController := database.User{}
	sessionController := database.Session{}
	session, err := sessionController.GetByID(sessionId)
	if err != nil {
		return nil, err.AST("re authenticate user")
	}
	user, err := userController.GetByID(session.UserID)
	if err != nil {
		return nil, err.AST("re authenticate user")
	}
	tokenData := structs.TokenData{
		UserId:    user.ID,
		CreatedAt: user.CreatedAt,
		Username:  user.Username,
		SessionId: session.ID}

	token, tokenExpiration, err := createJWT(&tokenData)
	if err != nil {
		return nil, err.AST("re authenticate user")
	}
	return &structs.TokenWrapper{User: user, Token: token, Expiration: tokenExpiration}, nil
}
func DeAuthenticateUser(userId uint) (int64, *custom_errors.CustomError) {
	userController := database.User{}
	sessionController := database.Session{}
	user, err := userController.GetByID(userId)
	if err != nil {
		return 0, err.AST("de authenticate user")
	}
	return sessionController.DeleteByUserId(user.ID)
}

func VerifyJWT(token string) (*structs.TokenData, *custom_errors.CustomError) {
	return verifyJWT(token)
}
func VerifyJWRT(token string) (*structs.TokenData, *custom_errors.CustomError) {
	return verifyJWRT(token)
}
