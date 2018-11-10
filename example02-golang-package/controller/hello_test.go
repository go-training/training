package controller

import "testing"

func TestHelloWorld(t *testing.T) {
	hello := HelloWorld("pgluffy")
	if hello != "Hi, pgluffy" {
		t.Errorf("Testing fail")
	}

	hello = HelloWorld("pgluffy ")
	if hello != "Hi, pgluffy" {
		t.Errorf("Testing fail")
	}
}
