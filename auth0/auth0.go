package goat // import "github.com/cycloidio/goat"

import (
	"bytes"
	"io/ioutil"
	"net"
	"net/http"
	"runtime"
	"time"
)

type Auth0 struct {
	domain      string
	client      *http.Client
	apiBasePath string
	respStatus  string
	respHeaders map[string][]string
	respBody    []byte
	token       string
}

func NewAuth0(auth0Domain, apiBasePath, token string) *Auth0 {
	var netTransport = &http.Transport{
		Dial: (&net.Dialer{
			Timeout: 5 * time.Second,
		}).Dial,
		TLSHandshakeTimeout: 5 * time.Second,
	}

	return &Auth0{
		client: &http.Client{
			Timeout:   time.Second * 10,
			Transport: netTransport,
		},
		domain:      auth0Domain,
		apiBasePath: apiBasePath,
		token:       token,
	}
}

func (a *Auth0) Call(apiEndPoint string, method string, body []byte) ([]byte, error) {
	var uri = a.domain + a.apiBasePath + apiEndPoint

	_, err := a.client.Get(uri)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(method, uri, bytes.NewBuffer(body))
	req.Header.Add("Authorization", "Bearer "+a.token)
	req.Header.Add("Content-Type", "application/json")
	resp, err := a.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	a.respStatus = resp.Status
	a.respHeaders = resp.Header
	res_body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	a.respBody = res_body
	return res_body, nil
}

// getFuncName is a helper function used to compose the error messages
func getFuncName() string {
	pc, _, _, _ := runtime.Caller(1)
	return runtime.FuncForPC(pc).Name()
}
