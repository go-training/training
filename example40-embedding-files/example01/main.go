package main

import "embed"

func main() {
	//go:embed hello.txt
	var s string
	print(s)

	//go:embed hello.txt
	var b []byte
	print(string(b))

	//go:embed hello.txt
	var f embed.FS
	data, _ := f.ReadFile("hello.txt")
	print(string(data))

}
