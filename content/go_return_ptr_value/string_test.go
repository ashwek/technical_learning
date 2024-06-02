package main

import "testing"

func BenchmarkReturnString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		value := returnString()
		_ = value
	}
}

func BenchmarkReturnStringPtr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		value := returnStringPtr()
		_ = value
	}
}
