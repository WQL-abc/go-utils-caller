// Package caller is return information about the caller.
package caller

import (
	"runtime"
	"strings"
)

const defaultSkip = 2

// GetCallFuncName - get the function name of the caller
func GetCallFuncName() string {
	sp := GetCallFuncRoute(defaultSkip)
	return sp[len(sp)-1]
}

// GetCallFuncRoute - get the route to the function name of the caller and return it as a slice
func GetCallFuncRoute(skip int) []string {
	pc, _, _, _ := runtime.Caller(skip)
	return strings.Split(runtime.FuncForPC(pc).Name(), ".")
}
