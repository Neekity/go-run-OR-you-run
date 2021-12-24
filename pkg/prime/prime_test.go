package prime

import "testing"

func BenchmarkPrimeA(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PrimeA(1000)
	}
}

func BenchmarkPrimeB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PrimeB(1000)
	}
}

func BenchmarkPrimeS(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PrimeS(1000)
	}
}

func BenchmarkPrimeSS(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PrimeSS(1000)
	}
}
