package goat

import (
	"net/url"
	"runtime"
)

func BuildParamsURL(params map[string]string) string {
	if params == nil || len(params) == 0 {
		return ""
	}
	var baseUrl = "?"
	params_url := url.Values{}

	for key, val := range params {
		params_url.Add(key, val)
	}
	return baseUrl + params_url.Encode()
}

func GetFuncName() string {
	pc, _, _, _ := runtime.Caller(1)
	return runtime.FuncForPC(pc).Name()
}
