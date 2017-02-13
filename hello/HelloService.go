package hello

import (
	"github.com/darenegade/SimpleGoKitService/util"
)

type HelloWorldRepository interface {
	helloService(HelloWorld) (string, error)
}

type HelloWorldService struct{}

type HelloWorld struct {
	Name string `json:"Name"`
}

func (HelloWorldService) helloService(name HelloWorld) (string, error) {
	if name.Name == "" {
		return "", util.ErrEmpty
	}
	return "Hello " + name.Name, nil
}
