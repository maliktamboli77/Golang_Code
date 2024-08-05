package calc

import "testing"

func BenchmarkCalculateSum(b *testing.B) {

	for i := 0; i < b.N; i++ {
		CalculateSum(100)
	}
}
