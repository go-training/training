package main

// #include <stdio.h>
// #include <stdlib.h>
/*
void print(char *str) {
  printf("%s\n", str);
}
*/
import "C"
import "unsafe"

func main() {
	s := "Hello Golang"
	cs := C.CString(s)
	C.print(cs)
	C.free(unsafe.Pointer(cs))
}
