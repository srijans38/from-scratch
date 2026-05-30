package parsers

import (
	"errors"
	"log"
	"slices"
	"strings"
)

var ErrInvalidRequestMethod = errors.New("invalid request method")
var ErrInvalidRequestPath = errors.New("invalid request path")

type Header struct {
	Key   string
	Value string
}

type Request struct {
	Method  Method
	Path    string
	Headers []Header
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

	headers := getRequestHeaders(request)

	return Request{
		Method:  method,
		Path:    path,
		Headers: headers,
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

func getRequestHeaders(request string) []Header {
	lines := strings.Split(request, "\r\n\r\n")

	requestLineWithHeaders := lines[0]

	_, headersPart, _ := strings.Cut(requestLineWithHeaders, "\r\n")

	headerLines := strings.Split(headersPart, "\r\n")

	var headers []Header

	for _, headerLine := range headerLines {
		if headerLine == "" {
			continue
		}

		key, value, found := strings.Cut(headerLine, ": ")

		if !found {
			log.Printf("Invalid header line: %s\n", headerLine)
			continue
		}

		headers = append(headers, Header{
			Key:   key,
			Value: value,
		})
	}

	log.Printf("Parsed headers: %v\n", headers)

	return headers
}
