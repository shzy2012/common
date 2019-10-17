package tool

import (
	"testing"
)

func Test_StringBuilder(t *testing.T) {
	origin := "apple,iphone,apple"
	result := StringBuilder(origin).Replace("apple", "fruit").Replace("iphone", "phone").Build()
	if result != "fruit,phone,fruit" {
		t.Error()
	}
}
