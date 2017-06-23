package goat // import "github.com/cycloidio/goat"

import (
	"errors"
	"net/http"
	"net/url"
)

func (a *Auth0) UserCall(method string, userID string, params url.Values, body []byte) ([]byte, error) {
	var apiEndPoint = "/users"
	if len(userID) > 0 {
		apiEndPoint += "/" + userID
	}
	if params != nil {
		apiEndPoint += params.Encode()
	}

	return a.Call(apiEndPoint, method, body)
}

func (a *Auth0) CreateUser(body []byte) ([]byte, error) {
	if body == nil || len(body) == 0 {
		return nil, errors.New(getFuncName() + ": 'body' cannot be nil nor empty.")
	}
	return a.UserCall(http.MethodPost, "", nil, body)
}

func (a *Auth0) GetUser(userID string) ([]byte, error) {
	if len(userID) == 0 {
		return nil, errors.New(getFuncName() + ": 'userID' cannot be empty.")
	}
	return a.UserCall(http.MethodGet, userID, nil, nil)
}

func (a *Auth0) GetUsers(params url.Values) ([]byte, error) {
	return a.UserCall(http.MethodGet, "", params, nil)
}

func (a *Auth0) UpdateUser(userID string, body []byte) ([]byte, error) {
	if len(userID) == 0 || body == nil || len(body) == 0 {
		return nil, errors.New(getFuncName() + ": 'userID' or 'body' cannot be nil nor empty.")
	}
	return a.UserCall(http.MethodPatch, userID, nil, body)
}

func (a *Auth0) DeleteUser(userID string) ([]byte, error) {
	if len(userID) == 0 {
		return nil, errors.New(getFuncName() + ": 'userID' cannot be empty.")
	}
	return a.UserCall(http.MethodDelete, userID, nil, nil)
}
