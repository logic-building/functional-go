package fp

import "testing"

func Fib(n int) int {
	if n < 2 {
		return n
	}
	return Fib(n-1) + Fib(n-2)
}

func BenchmarkPMapInt(b *testing.B) {
	for n := 0; n < b.N; n++ {
		PMapInt(Fib, RangeInt(2, 10))
	}
}

func BenchmarkPMapIntRandomOrder(b *testing.B) {
	for n := 0; n < b.N; n++ {
		PMapInt(Fib, RangeInt(2, 10), Optional{RandomOrder: true})
	}
}
