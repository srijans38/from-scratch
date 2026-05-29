package parsers

import (
	"errors"
	"slices"
	"strings"
)

var ErrInvalidRequestMethod = errors.New("invalid request method")

func ParseRequest(data []byte) (string, error) {
	request := string(data)

	requestLine, _, _ := strings.Cut(request, "\r\n")

	method, err := getRequestMethod(requestLine)

	if err != nil {
		return "", err
	}

	return method, nil
}

func getRequestMethod(request string) (string, error) {
	allowedMethods := []string{"GET", "POST", "PUT", "DELETE", "PATCH", "HEAD", "OPTIONS"}

	method, _, _ := strings.Cut(request, " ")

	if !slices.Contains(allowedMethods, method) {
		return "", ErrInvalidRequestMethod
	}

	return method, nil
}
