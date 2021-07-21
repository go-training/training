package main

import (
	"embed"
)

//go:embed hello.txt
var s string

//go:embed hello.txt
var b []byte

//go:embed hello.txt
var f embed.FS

func main() {
	print(s)
	print(string(b))
	data, _ := f.ReadFile("hello.txt")
	print(string(data))
}
