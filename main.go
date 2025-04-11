package main

//go:generate gobind -i build/libgobind-example.so

import (
	"github.com/soerlemans/gobind-example/gbmod/hello_world"
)

func main() {
	// Call the C++ code.
	hello_world.HelloWorld()
}
