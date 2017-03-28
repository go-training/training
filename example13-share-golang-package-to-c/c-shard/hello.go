package main

import "C"
import "fmt"

//export SayHello
func SayHello(name string) {
	fmt.Printf("Nautilus says: Hello, %s!\n", name)
}

//export SayBye
func SayBye() {
	fmt.Println("Nautilus says: Bye!")
}

func main() {
	// We need the main function to make possible
	// CGO compiler to compile the package as C shared library
}
