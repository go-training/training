package main

import "C"
import "fmt"

func SayHello(name string) {
	fmt.Printf("Golang says: Hello, %s!\n", name)
}

func SayBye() {
	fmt.Println("Golang says: Bye!")
}
