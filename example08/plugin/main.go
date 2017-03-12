package main

import (
	"plugin"
)

func main() {
	p, err := plugin.Open("hello.so")
	if err != nil {
		panic(err)
	}

	sayHelloSymbol, err := p.Lookup("SayHello")
	if err != nil {
		panic(err)
	}

	// Note that because the compiler cannot know in advance
	// what data type or function signature you are trying
	// to use, type assertions must be used.
	sayHello := sayHelloSymbol.(func(string))
	sayHello("appleboy")
}
