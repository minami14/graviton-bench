package main

import (
	"testing"

	"github.com/google/uuid"
)

func BenchmarkUUIDDisableRandPool(b *testing.B) {
	uuid.DisableRandPool()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		uuid.New()
	}
}

func BenchmarkUUIDEnableRandPool(b *testing.B) {
	uuid.EnableRandPool()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		uuid.New()
	}
}
