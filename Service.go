package main

import "errors"

type HelloWorldService interface {
	helloService(string) (string, error)
}

type helloWorldService struct{}

func (helloWorldService) helloService(name string) (string, error) {
	if name == "" {
		return "", ErrEmpty
	}
	return "Hello " + name, nil
}

var ErrEmpty = errors.New("Empty string")
