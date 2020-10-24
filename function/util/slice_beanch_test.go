package util

import (
	"fmt"
	"testing"
)

func generateTestData() []string {
	var data []string
	for i := 0; i < 2300; i++ {
		data = append(data, fmt.Sprintf("data%d", i))
	}
	return data
}

func Benchmark_subtractSliceWithContains(b *testing.B) {
	src := generateTestData()
	sub := generateTestData()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		subtractSliceWithContains(src, sub)
	}
}
