package tools

import (
	"fmt"
	"testing"
)

func Test_StringBuilder(t *testing.T) {
	origin := "apple,iphone,apple"
	result := StringBuilder(origin).Replace("apple", "fruit").Replace("iphone", "phone").Build()
	if result != "fruit,phone,fruit" {
		t.Error()
	}
}

func Test_Splits(t *testing.T) {
	// aba
	s := "a,b,c;d"
	result := Splits(s, ",;")
	fmt.Printf("%s\n", result)
}

func Benchmark_GetRandomString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetRandomString(64)
	}
}
