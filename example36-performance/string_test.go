package main

import (
	"fmt"
	"strings"
	"testing"
	"unsafe"
)

func BenchmarkString01(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = string01(bar)
	}
}

func BenchmarkString02(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("%s %v", bar.ID, bar.Addr)
	}
}

func b2s(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func BenchmarkString03(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b := make([]byte, 0, 40)
		b = append(b, bar.ID...)
		b = append(b, ' ')
		b = append(b, bar.Addr...)
		_ = b2s(b)
	}
}

func BenchmarkString04(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var b strings.Builder
		b.WriteString(bar.ID)
		b.WriteString(" ")
		b.WriteString(bar.Addr)
		_ = b.String()
	}
}

func BenchmarkString05(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var b strings.Builder
		for i := 0; i < len(bar.ID); i++ {
			c := bar.ID[i]
			b.WriteByte(c)
		}
		b.WriteByte(' ')
		for i := 0; i < len(bar.Addr); i++ {
			c := bar.ID[i]
			b.WriteByte(c)
		}
		_ = b.String()
	}
}
