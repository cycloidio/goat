package goat_test

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	goat "github.com/cycloidio/goat/auth0"
)

const (
	bodySample = "{'return': 'ok'}"
)

var jsonOK = []byte(bodySample)
var ts *httptest.Server
var a *goat.Auth0

func init() {
	ts = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, string(bodySample))
	}))
	a = goat.NewAuth0(ts.URL, "", "")
}

type MethodCheck struct {
	testName     string
	userID       string
	body         []byte
	params       url.Values
	jsonExpected []byte
	errExpected  bool
}

var CheckUsersMethods = map[string][]MethodCheck{
	"CreateUser": {
		{"Body should not be nil", "", nil, nil, nil, true},
		{"Body should not be empty", "", []byte{}, nil, nil, true},
		{"Body is valid", "", []byte{1, 7, 42}, nil, jsonOK, false},
	},
	"GetUser": {
		{"User id cannot be empty", "", nil, nil, nil, true},
		{"User is valid", "user-id", nil, nil, jsonOK, false},
	},
	"GetUsers": {
		{"With parameters", "", nil, url.Values{"page": []string{"1"}, "per_page": []string{"42"}}, jsonOK, false},
		{"Without parameters", "", nil, nil, jsonOK, false},
	},
	"UpdateUser": {
		{"User id and body cannot be empty", "", []byte{}, nil, nil, true},
		{"User id empty and body cannot be nil", "", nil, nil, nil, true},
		{"User id cannot be empty", "", []byte{1, 7, 42}, nil, nil, true},
		{"Body cannot be nil", "user-id", nil, nil, nil, true},
		{"Body cannot be empty", "user-id", []byte{}, nil, nil, true},
		{"Body and user id are valid", "user-id", []byte{1, 7, 42}, nil, jsonOK, false},
	},
	"DeleteUser": {
		{"User id cannot be empty", "", nil, nil, nil, true},
		{"User id is valid", "user-id", nil, nil, jsonOK, false},
	},
}

func TestMethodCreateUser(t *testing.T) {
	for _, call := range CheckUsersMethods["CreateUser"] {
		fmt.Println("\t" + call.testName)
		json, err := a.CreateUser(call.body)
		if bytes.Compare(json, call.jsonExpected) != 0 {
			t.Errorf("Failed test %#v:\nexpected json: %#v\ngot json: %#v", call.testName, string(call.jsonExpected), string(json))
		}
		if (err != nil) != call.errExpected {
			t.Errorf("Failed test %#v:\nexpected err: %#v\ngot err: %#v", call.testName, call.errExpected, err)
		}
	}
}

func TestMethodGetUser(t *testing.T) {
	for _, call := range CheckUsersMethods["GetUser"] {
		fmt.Println("\t" + call.testName)
		json, err := a.GetUser(call.userID)
		if bytes.Compare(json, call.jsonExpected) != 0 {
			t.Errorf("Failed test %#v:\nexpected json: %#v\ngot json: %#v", call.testName, string(call.jsonExpected), string(json))
		}
		if (err != nil) != call.errExpected {
			t.Errorf("Failed test %#v:\nexpected err: %#v\ngot err: %#v", call.testName, call.errExpected, err)
		}
	}
}

func TestMethodGetUsers(t *testing.T) {
	for _, call := range CheckUsersMethods["GetUsers"] {
		fmt.Println("\t" + call.testName)
		json, err := a.GetUsers(call.params)
		if bytes.Compare(json, call.jsonExpected) != 0 {
			t.Errorf("Failed test %#v:\nexpected json: %#v\ngot json: %#v", call.testName, string(call.jsonExpected), string(json))
		}
		if (err != nil) != call.errExpected {
			t.Errorf("Failed test %#v:\nexpected err: %#v\ngot err: %#v", call.testName, call.errExpected, err)
		}
	}
}

func TestMethodUpdateUser(t *testing.T) {
	for _, call := range CheckUsersMethods["UpdateUser"] {
		fmt.Println("\t" + call.testName)
		json, err := a.UpdateUser(call.userID, call.body)
		if bytes.Compare(json, call.jsonExpected) != 0 {
			t.Errorf("Failed test %#v:\nexpected json: %#v\ngot json: %#v", call.testName, string(call.jsonExpected), string(json))
		}
		if (err != nil) != call.errExpected {
			t.Errorf("Failed test %#v:\nexpected err: %#v\ngot err: %#v", call.testName, call.errExpected, err)
		}
	}
}

func TestMethodDeleteUser(t *testing.T) {
	for _, call := range CheckUsersMethods["DeleteUser"] {
		fmt.Println("\t" + call.testName)
		json, err := a.DeleteUser(call.userID)
		if bytes.Compare(json, call.jsonExpected) != 0 {
			t.Errorf("Failed test %#v:\nexpected json: %#v\ngot json: %#v", call.testName, string(call.jsonExpected), string(json))
		}
		if (err != nil) != call.errExpected {
			t.Errorf("Failed test %#v:\nexpected err: %#v\ngot err: %#v", call.testName, call.errExpected, err)
		}
	}
}
