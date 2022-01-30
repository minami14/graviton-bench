package main

import "testing"

func BenchmarkCopy64(b *testing.B) {
	src := make([]byte, 64)
	dst := make([]byte, len(src))
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		copy(dst, src)
	}
}

func BenchmarkCopy1024(b *testing.B) {
	src := make([]byte, 1024)
	dst := make([]byte, len(src))
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		copy(dst, src)
	}
}

func BenchmarkMake64(b *testing.B) {
	buf := make([]byte, 64)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf = make([]byte, len(buf))
	}
}

func BenchmarkMake1024(b *testing.B) {
	buf := make([]byte, 1024)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf = make([]byte, len(buf))
	}
}
