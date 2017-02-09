package hello

import (
	"github.com/darenegade/SimpleGoKitService/util"
)



type HelloWorldInterface interface {
	helloService(HelloWorld) (string, error)
}

type HelloWorldService struct{}

type HelloWorld struct {
	Name string `json:"name"`
}

func (HelloWorldService) HelloService(name HelloWorld) (string, error) {
	if name.Name == "" {
		return "", util.ErrEmpty
	}
	return "Hello " + name.Name, nil
}


