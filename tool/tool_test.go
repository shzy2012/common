package tool

import (
	"fmt"
	"testing"
	"time"
)

func Benchmark_GetRandomString(b *testing.B) {
	tool := NewTool()
	for i := 0; i < b.N; i++ {
		tool.GetRandomString(64)
	}
}

func Test_AutoIncrementing(t *testing.T) {
	tool := NewTool()
	//tool.SetCounter(999999999)
	for n := 0; n <= 10; n++ {
		go func(n int) {
			for i := 0; i < 1000000000; i++ {
				fmt.Println(n, "-", tool.AutoIncrementing())
			}
		}(n)
	}

	time.Sleep(time.Second * 10000)
}

func Benchmark_AutoIncrementing(b *testing.B) {
	tool := NewTool()
	for i := 0; i < b.N; i++ {
		fmt.Println(tool.AutoIncrementing())
	}
}
