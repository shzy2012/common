package tool

import (
	"fmt"
	"math"
	"sync/atomic"
)

var counter int64
var iset bool

//AutoIncrementing 自动自增(base设置自增起始点)
func AutoIncrementing(base int64) int64 {
	if !iset {
		counter = base
		iset = true
	}
	return atomic.AddInt64(&counter, 1)
}

//Sum Get sum of numbers
func Sum(numbers []int) int {
	sum := 0
	for _, v := range numbers {
		sum = sum + v
	}

	return sum
}

//Max Get max value from numbers
func Max(numbers []int) int {
	max := math.MinInt64
	for _, n := range numbers {
		if n > max {
			max = n
		}
	}

	return max
}

//Min Get min value from numbers
func Min(numbers []int) int {
	min := math.MaxInt64
	for _, n := range numbers {
		if n < min {
			min = n
		}
	}

	return min
}

//F64SliceToStringSlice convert float64 slice to string slice
func F64SliceToStringSlice(list []float64) []string {
	tmp := make([]string, 0)
	for _, ele := range list {
		tmp = append(tmp, fmt.Sprintf("%v", ele))
	}
	return tmp
}
