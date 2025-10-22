package tools

import (
	"fmt"
	"testing"
)

func Test_AutoIncrementing(t *testing.T) {
	for n := 0; n < 1; n++ {
		go func(n int) {
			for i := 0; i < 1000000; i++ {
				fmt.Println(n, "-", AutoIncrementing(-1))
			}
		}(n)
	}

}

func Test_Sum(t *testing.T) {
	list := []int{1, 3, 5}
	fmt.Printf("sum:%v\n", Sum(list))
}

func Test_Max(t *testing.T) {
	list := []int{1, 3, 5}
	fmt.Printf("max:%v\n", Max(list))
}

func Test_Min(t *testing.T) {
	list := []int{1, 3, 5}
	fmt.Printf("min:%v\n", Min(list))
}

func Benchmark_AutoIncrementing(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Println(AutoIncrementing(-1))
	}
}
