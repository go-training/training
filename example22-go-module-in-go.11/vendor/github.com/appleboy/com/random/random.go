package random

import (
	"math/rand"
	"time"
)

type (
	// Charset is string type
	Charset string
)

const (
	// Alphanumeric contain Alphabetic and Numeric
	Alphanumeric Charset = Alphabetic + Numeric
	// Alphabetic is \w+ \W
	Alphabetic Charset = "abcdefghijklmnopqrstuvwxyz" + "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	// Numeric is number list
	Numeric Charset = "0123456789"
	// Hex is Hexadecimal
	Hex Charset = Numeric + "abcdef"
)

var seededRand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

// StringWithCharset support rand string you defined
func StringWithCharset(length int, charset Charset) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

// String supply rand string
func String(length int) string {
	return StringWithCharset(length, Alphanumeric)
}
