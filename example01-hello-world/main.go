package main

import "fmt"

func main() {
	a := 1
	fmt.Println("===Fork Github form appleboy===")
	fmt.Printf(HelloWorld("pgluffy"))
	fmt.Println("一天就學會 Go 語言")

	if (a >= 1) { fmt.Println("a >= 1") }
}

func HelloWorld(user_name string) string {
	return fmt.Sprintf("Hi, %s ", user_name)
}
