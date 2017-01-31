package goat

import (
	"bytes"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

const (
	DELETE = "DELETE"
	GET    = "GET"
	PATCH  = "PATCH"
	POST   = "POST"
)

type Auth0 struct {
	domain      string
	client      *http.Client
	apiURL      string
	respStatus  string
	respHeaders map[string][]string
	respBody    []byte
	token       string
}

func Auth0New(auth0Domain, apiVersion, token string) *Auth0 {
	var netTransport = &http.Transport{
		Dial: (&net.Dialer{
			Timeout: 5 * time.Second,
		}).Dial,
		TLSHandshakeTimeout: 5 * time.Second,
	}

	a := Auth0{}
	a.client = &http.Client{
		Timeout:   time.Second * 10,
		Transport: netTransport,
	}
	a.domain = auth0Domain
	a.apiURL = apiVersion
	a.token = token
	return &a
}

func (a *Auth0) Call(apiAction string, method string, body []byte) ([]byte, error) {
	var uri = a.domain + a.apiURL + apiAction

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
