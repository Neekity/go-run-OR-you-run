package damnifelse

import "testing"

func BenchmarkCountGtNumberElse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CountGtNumberElse()
	}
}

func BenchmarkCountGtNumberIf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CountGtNumberIf()
	}
}
