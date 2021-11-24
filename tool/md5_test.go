package tool

import (
	"fmt"
	"testing"
)

func TestSha512(t *testing.T) {
	passwd := Sha256("123")
	fmt.Println(passwd)

	passwd = Sha256("123")
	fmt.Println(passwd)
}
