package main

import (
	"fmt"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	project := os.Getenv("GOLANG_PROJECT")
	fmt.Println(project)
}
