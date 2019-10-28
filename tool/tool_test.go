package tool

import (
	"fmt"
	"testing"
)

func Test_AutoIncrementing(t *testing.T) {
	tool := T()
	//tool.SetCounter(999999999)
	for n := 0; n <= 10; n++ {
		go func(n int) {
			for i := 0; i < 1000000000; i++ {
				fmt.Println(n, "-", tool.AutoIncrementing())
			}
		}(n)
	}
}

func Benchmark_AutoIncrementing(b *testing.B) {
	tool := T()
	for i := 0; i < b.N; i++ {
		fmt.Println(tool.AutoIncrementing())
	}
}
