package runtimeutils

import "errors"

var (
	ErrNoFramesWereFound    = errors.New("no caller frames were provided")
	ErrNoCallerInfoWasFound = errors.New("no caller information was found")
	ErrMaxStackIsLowerThan1 = errors.New("maxStack argument is lower than 1")
)
