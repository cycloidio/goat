package goat

import "runtime"

func BuildParamsURL(params map[string]string) string {
	if params == nil || len(params) == 0 {
		return ""
	}
	var paramsURL = ""
	for _, val := range params {
		paramsURL += "?" + val
	}
	return paramsURL
}

func GetFuncName() string {
	pc, _, _, _ := runtime.Caller(1)
	return runtime.FuncForPC(pc).Name()
}
