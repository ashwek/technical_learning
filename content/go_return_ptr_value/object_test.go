package main

import "testing"

func BenchmarkReturnData(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a := returnData()
		_ = a
	}
}

func BenchmarkReturnDataPtr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a := returnDataPtr()
		_ = a
	}
}
