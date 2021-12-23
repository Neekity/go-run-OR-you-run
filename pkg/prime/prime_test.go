package prime

import "testing"

func BenchmarkPrimeA(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PrimeA(10000)
	}
}

func BenchmarkPrimeB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PrimeB(10000)
	}
}

func BenchmarkPrimeS(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PrimeS(10000)
	}
}
