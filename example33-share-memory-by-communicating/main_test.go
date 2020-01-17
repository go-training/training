package main

import "testing"

func BenchmarkAddByShareMemory(b *testing.B) {
	for i := 0; i < b.N; i++ {
		addByShareMemory(100)
	}
}

func BenchmarkAddByShareCommunicate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		addByShareCommunicate(100)
	}
}
