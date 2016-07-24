package goat

import (
	"errors"
)

func (a *Auth0) UserCall(method string, user_id string, params map[string]string, body []byte) ([]byte, error) {
	var api_action = "/users"
	if (len(user_id) > 0) {
		api_action += "/" + user_id
	}
	api_action += BuildParamsURL(params)
	return a.Call(api_action, method, body)
}

func (a *Auth0) CreateUser(body []byte) ([]byte, error) {
	if (body == nil || len(body) == 0) {
		return nil, errors.New(GetFuncName() + ": 'body' cannot be nil nor empty.")
	}
	return a.UserCall(POST, "", nil, body)
}

func (a *Auth0) GetUser(user_id string) ([]byte, error) {
	if (len(user_id) == 0) {
		return nil, errors.New(GetFuncName() + ": 'user_id' cannot be empty.")
	}
	return a.UserCall(GET, user_id, nil, nil)
}

func (a *Auth0) GetUsers(params map[string]string) ([]byte, error) {
	return a.UserCall(GET, "", params, nil)
}

func (a *Auth0) UpdateUser(user_id string, body []byte) ([]byte, error) {
	if (len(user_id) == 0 || body == nil || len(body) == 0) {
		return nil, errors.New(GetFuncName() + ": 'user_id' or 'body' cannot be nil nor empty.")
	}
	return a.UserCall(PATCH, user_id, nil, body)
}

func (a *Auth0) DeleteUser(user_id string) ([]byte, error) {
	if (len(user_id) == 0) {
		return nil, errors.New(GetFuncName() + ": 'user_id' cannot be empty.")
	}
	return a.UserCall(DELETE, user_id, nil, nil)
}