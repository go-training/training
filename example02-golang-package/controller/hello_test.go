package controller

import "testing"

func TestHelloWorld(t *testing.T) {
	hello := HelloWorld("appleboy")
	if hello != "Hi, appleboy" {
		t.Errorf("Testing fail")
	}

	hello = HelloWorld("appleboy ")
	if hello != "Hi, appleboy" {
		t.Errorf("Testing fail")
	}
}
