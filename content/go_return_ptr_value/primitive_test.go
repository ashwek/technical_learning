package main

import "testing"

func BenchmarkReturnNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		value := returnNumber()
		_ = value
	}
}

func BenchmarkReturnNumberPtr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		value := returnNumberPtr()
		_ = value
	}
}
