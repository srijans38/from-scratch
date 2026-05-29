package parsers

import (
	"errors"
	"slices"
	"strings"
)

var ErrInvalidRequestMethod = errors.New("invalid request method")
var ErrInvalidRequestPath = errors.New("invalid request path")

type Request struct {
	Method Method
	Path   string
}

type Method string

const (
	GET     Method = "GET"
	POST    Method = "POST"
	PUT     Method = "PUT"
	DELETE  Method = "DELETE"
	PATCH   Method = "PATCH"
	HEAD    Method = "HEAD"
	OPTIONS Method = "OPTIONS"
)

func ParseRequest(data []byte) (Request, error) {
	request := string(data)

	requestLine, _, _ := strings.Cut(request, "\r\n")

	method, err := getRequestMethod(requestLine)

	if err != nil {
		return Request{}, err
	}

	path, err := getRequestPath(requestLine)

	if err != nil {
		return Request{}, err
	}

	return Request{
		Method: method,
		Path:   path,
	}, nil
}

func getRequestMethod(request string) (Method, error) {
	allowedMethods := []Method{GET, POST, PUT, DELETE, PATCH, HEAD, OPTIONS}

	method, _, _ := strings.Cut(request, " ")

	if !slices.Contains(allowedMethods, Method(method)) {
		return "", ErrInvalidRequestMethod
	}

	return Method(method), nil
}

func getRequestPath(request string) (string, error) {
	parts := strings.Split(request, " ")

	if len(parts) < 2 {
		return "", ErrInvalidRequestPath
	}

	return parts[1], nil
}
