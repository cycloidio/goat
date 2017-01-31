package goat

import (
	"errors"
)

func (a *Auth0) UserCall(method string, userID string, params map[string]string, body []byte) ([]byte, error) {
	var apiAction = "/users"
	if len(userID) > 0 {
		apiAction += "/" + userID
	}
	apiAction += BuildParamsURL(params)
	return a.Call(apiAction, method, body)
}

func (a *Auth0) CreateUser(body []byte) ([]byte, error) {
	if body == nil || len(body) == 0 {
		return nil, errors.New(GetFuncName() + ": 'body' cannot be nil nor empty.")
	}
	return a.UserCall(POST, "", nil, body)
}

func (a *Auth0) GetUser(userID string) ([]byte, error) {
	if len(userID) == 0 {
		return nil, errors.New(GetFuncName() + ": 'userID' cannot be empty.")
	}
	return a.UserCall(GET, userID, nil, nil)
}

func (a *Auth0) GetUsers(params map[string]string) ([]byte, error) {
	return a.UserCall(GET, "", params, nil)
}

func (a *Auth0) UpdateUser(userID string, body []byte) ([]byte, error) {
	if len(userID) == 0 || body == nil || len(body) == 0 {
		return nil, errors.New(GetFuncName() + ": 'userID' or 'body' cannot be nil nor empty.")
	}
	return a.UserCall(PATCH, userID, nil, body)
}

func (a *Auth0) DeleteUser(userID string) ([]byte, error) {
	if len(userID) == 0 {
		return nil, errors.New(GetFuncName() + ": 'userID' cannot be empty.")
	}
	return a.UserCall(DELETE, userID, nil, nil)
}
