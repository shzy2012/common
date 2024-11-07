package tools

import (
	"fmt"
	"testing"
	"time"
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

func Test_TimeAgo(t *testing.T) {
	t1 := time.Now().Add(-50000 * time.Hour)
	fmt.Println(TimeAgo(t1))
}

func Benchmark_GetRandomString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetRandomString(64)
	}
}
