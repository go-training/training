package main

import "time"

type foo struct {
	ID        string
	Addr      string
	CreatedAt time.Time
}

var bar = foo{
	ID:   "1234567890",
	Addr: "1234567890",
}

func string01(bar foo) string {
	s := bar.ID
	s += " " + bar.Addr
	return s
}
