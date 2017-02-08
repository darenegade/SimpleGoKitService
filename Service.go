package main

import "errors"

var ErrEmpty = errors.New("Empty string")

type HelloWorldService interface {
	helloService(helloWorld) (string, error)
}

type helloWorldService struct{}

type helloWorld struct {
	Name string `json:"name"`
}

func (helloWorldService) helloService(name helloWorld) (string, error) {
	if name.Name == "" {
		return "", ErrEmpty
	}
	return "Hello " + name.Name, nil
}


