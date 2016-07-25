package goat_test

import (
	"bytes"
	"fmt"
	"github.com/cycloidio/goat/auth0"
	"net/http"
	"net/http/httptest"
	"testing"
)

const (
	BODY_SAMPLE = "{'return': 'ok'}"
)

var JSON_OK = []byte(BODY_SAMPLE)
var ts *httptest.Server
var a *goat.Auth0

func init() {
	ts = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, string(BODY_SAMPLE))
	}))
	a = goat.Auth0New(ts.URL, "", "")
}

type MethodCheck struct {
	test_name     string
	user_id       string
	body          []byte
	params        map[string]string // Unused currently
	json_expected []byte
	err_expected  bool
}

var CheckUsersMethods = map[string][]MethodCheck{
	"CreateUser": {
		{"Body should not be nil", "", nil, nil, nil, true},
		{"Body should not be empty", "", []byte{}, nil, nil, true},
		{"Body is valid", "", []byte{1, 7, 42}, nil, JSON_OK, false},
	},
	"GetUser": {
		{"User id cannot be empty", "", nil, nil, nil, true},
		{"User is valid", "user-id", nil, nil, JSON_OK, false},
	},
	"GetUsers": {
		{"With parameters", "", nil, map[string]string{"page": "1", "per_page": "42"}, JSON_OK, false},
		{"Without parameters", "", nil, nil, JSON_OK, false},
	},
	"UpdateUser": {
		{"User id and body cannot be empty", "", []byte{}, nil, nil, true},
		{"User id empty and body cannot be nil", "", nil, nil, nil, true},
		{"User id cannot be empty", "", []byte{1, 7, 42}, nil, nil, true},
		{"Body cannot be nil", "user-id", nil, nil, nil, true},
		{"Body cannot be empty", "user-id", []byte{}, nil, nil, true},
		{"Body and user id are valid", "user-id", []byte{1, 7, 42}, nil, JSON_OK, false},
	},
	"DeleteUser": {
		{"User id cannot be empty", "", nil, nil, nil, true},
		{"User id is valid", "user-id", nil, nil, JSON_OK, false},
	},
}

func TestMethodCreateUser(t *testing.T) {
	for _, call := range CheckUsersMethods["CreateUser"] {
		fmt.Println("\t" + call.test_name)
		json, err := a.CreateUser(call.body)
		if bytes.Compare(json, call.json_expected) != 0 {
			t.Errorf("Failed test %#v:\nexpected json: %#v\ngot json: %#v", call.test_name, string(call.json_expected), string(json))
		}
		if (err != nil) != call.err_expected {
			t.Errorf("Failed test %#v:\nexpected err: %#v\ngot err: %#v", call.test_name, call.err_expected, err)
		}
	}
}

func TestMethodGetUser(t *testing.T) {
	for _, call := range CheckUsersMethods["GetUser"] {
		fmt.Println("\t" + call.test_name)
		json, err := a.GetUser(call.user_id)
		if bytes.Compare(json, call.json_expected) != 0 {
			t.Errorf("Failed test %#v:\nexpected json: %#v\ngot json: %#v", call.test_name, string(call.json_expected), string(json))
		}
		if (err != nil) != call.err_expected {
			t.Errorf("Failed test %#v:\nexpected err: %#v\ngot err: %#v", call.test_name, call.err_expected, err)
		}
	}
}

func TestMethodGetUsers(t *testing.T) {
	for _, call := range CheckUsersMethods["GetUsers"] {
		fmt.Println("\t" + call.test_name)
		json, err := a.GetUsers(call.params)
		if bytes.Compare(json, call.json_expected) != 0 {
			t.Errorf("Failed test %#v:\nexpected json: %#v\ngot json: %#v", call.test_name, string(call.json_expected), string(json))
		}
		if (err != nil) != call.err_expected {
			t.Errorf("Failed test %#v:\nexpected err: %#v\ngot err: %#v", call.test_name, call.err_expected, err)
		}
	}
}

func TestMethodUpdateUser(t *testing.T) {
	for _, call := range CheckUsersMethods["UpdateUser"] {
		fmt.Println("\t" + call.test_name)
		json, err := a.UpdateUser(call.user_id, call.body)
		if bytes.Compare(json, call.json_expected) != 0 {
			t.Errorf("Failed test %#v:\nexpected json: %#v\ngot json: %#v", call.test_name, string(call.json_expected), string(json))
		}
		if (err != nil) != call.err_expected {
			t.Errorf("Failed test %#v:\nexpected err: %#v\ngot err: %#v", call.test_name, call.err_expected, err)
		}
	}
}

func TestMethodDeleteUser(t *testing.T) {
	for _, call := range CheckUsersMethods["DeleteUser"] {
		fmt.Println("\t" + call.test_name)
		json, err := a.DeleteUser(call.user_id)
		if bytes.Compare(json, call.json_expected) != 0 {
			t.Errorf("Failed test %#v:\nexpected json: %#v\ngot json: %#v", call.test_name, string(call.json_expected), string(json))
		}
		if (err != nil) != call.err_expected {
			t.Errorf("Failed test %#v:\nexpected err: %#v\ngot err: %#v", call.test_name, call.err_expected, err)
		}
	}
}
