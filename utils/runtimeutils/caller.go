/*
Package caller get information regarding caller stack.
*/
package runtimeutils

import (
	"runtime"
	"strings"
)

const (
	callerSkip int = 3
)

// GetCallerInfo returns a struct of CallerInfo regarding the caller frame
// requested by callerSkipFrames.
//
// With callerSkipFrames, when value is 0 - It will return runtime.Callers,
// When callerSkipFrames is 1, it will return GetCallerInfo.
// So it is recommended to set callerSkipFrames for 2 or beyond to get the
// current caller for information.
func GetCallerInfo(callerSkipFrames int) (CallerInfo, error) {
	fpcs := make([]uintptr, 1)

	// Skip levels to get the caller
	n := runtime.Callers(callerSkipFrames, fpcs)
	if n == 0 {
		return CallerInfo{}, ErrNoFramesWereFound
	}

	caller := runtime.FuncForPC(fpcs[0] - 1)
	if caller == nil {
		return CallerInfo{}, ErrNoCallerInfoWasFound
	}

	name := caller.Name()
	packageAndName := strings.Split(name, ".")
	length := len(packageAndName)
	packageName := ""
	funcName := name

	if length > 1 {
		funcName = packageAndName[length-1]
		packageName = strings.Join(packageAndName[0:length-1], ".")
	}

	fileName, line := caller.FileLine(fpcs[0] - 1)

	result := CallerInfo{
		Entry:   caller.Entry(),
		LineNo:  line,
		File:    fileName,
		Package: packageName,
		Func:    funcName,
	}

	return result, nil
}

// GetCallerStack provides a list of all stack level (from callerSkipFrames
// and up to most of maxStack or the latest element that was read).
//
// With callerSkipFrames, when value is 0 - It will return runtime.Callers,
// When callerSkipFrames is 1, it will return GetCallerInfo.
// So it is recommended to set callerSkipFrames for 2 or beyond to get the
// current caller for information.
//
// maxStack must be bigger than 0 or an error will be returned.
func GetCallerStack(callerSkipFrames, maxStack int) ([]CallerInfo, error) {
	if maxStack <= 0 {
		return nil, ErrMaxStackIsLowerThan1
	}

	fpcs := make([]uintptr, maxStack, maxStack+1)

	// Skip 2 levels to get the caller
	fpcsLength := runtime.Callers(callerSkipFrames, fpcs)
	if fpcsLength == 0 {
		return nil, ErrNoFramesWereFound
	}

	var result []CallerInfo

	for index := 0; index < fpcsLength; index++ {
		caller := runtime.FuncForPC(fpcs[index] - 1)
		if caller == nil {
			return nil, ErrNoCallerInfoWasFound
		}

		name := caller.Name()
		packageAndName := strings.Split(name, ".")
		length := len(packageAndName)
		packageName := ""
		funcName := name

		if length > 1 {
			funcName = packageAndName[length-1]
			packageName = strings.Join(packageAndName[0:length-1], ".")
		}

		fileName, line := caller.FileLine(fpcs[index] - 1)

		result = append(
			result, CallerInfo{
				Entry:   caller.Entry(),
				LineNo:  line,
				File:    fileName,
				Package: packageName,
				Func:    funcName,
			},
		)
	}

	return result, nil
}

// GetCallerFunctionName return the name of whom called this function
func GetCallerFunctionName() string {
	result, err := GetCallerInfo(callerSkip)
	if err != nil {
		return ""
	}

	return result.Func
}
