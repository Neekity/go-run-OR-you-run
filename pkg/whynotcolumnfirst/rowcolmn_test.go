package whynotcolumnfirst

import "testing"

func BenchmarkRowFirst4096(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CopyRowFirst4096()
	}
}

func BenchmarkColumnFirst4096(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CopyColumnFirst4096()
	}
}

func BenchmarkRowFirst2048(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CopyRowFirst2048()
	}
}

func BenchmarkColumnFirst2048(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CopyColumnFirst2048()
	}
}
