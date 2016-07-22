package goat

import "runtime"

func BuildParamsURL(params map[string]string) string {
	if params == nil || len(params) == 0 {
		return ""
	}
	var params_url = ""
	for _, val := range params {
		params_url += "?" + val
	}
	return params_url
}

func GetFuncName() string {
	pc, _, _, _ := runtime.Caller(1)
	return runtime.FuncForPC(pc).Name()
}
