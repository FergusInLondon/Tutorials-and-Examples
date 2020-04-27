package main

import "errors"

//go:generate go run .generators/status_gen.go

// StatusCode is a simple alias for int
type StatusCode int

// StatusMessage is a simple alias for string.
type StatusMessage string

// Message retrieves the message for a given StatusCode
func (sc StatusCode) Message() (StatusMessage, error) {
	if val, ok := statusMap[sc]; ok {
		return val, nil
	}

	return "", errors.New("invalid status code")
}
