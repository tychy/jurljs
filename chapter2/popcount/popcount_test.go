package popcount

import "testing"

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(11111110011)
	}
}

func BenchmarkPopCount16(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount16(11111110011)
	}
}
func BenchmarkPopCountLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountLoop(11111110011)
	}
}
