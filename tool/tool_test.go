package tool

import "testing"

func Benchmark_GetRandomString(b *testing.B) {
	tool := NewTool()
	for i := 0; i < b.N; i++ {
		tool.GetRandomString(64)
	}
}