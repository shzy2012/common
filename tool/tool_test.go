package tool

import (
	"fmt"
	"testing"
)

func Benchmark_GetRandomString(b *testing.B) {

	for i := 0; i < b.N; i++ {
		GetRandomString(64)
	}
}

func Test_AutoIncrementing(t *testing.T) {

	//tool.SetCounter(999999999)
	for n := 0; n <= 10; n++ {
		go func(n int) {
			for i := 0; i < 1000000000; i++ {
				fmt.Println(n, "-", AutoIncrementing())
			}
		}(n)
	}
}

func Benchmark_AutoIncrementing(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Println(AutoIncrementing())
	}
}
