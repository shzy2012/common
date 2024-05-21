package tools

import (
	"fmt"
	"net/url"
	"testing"
)

func Test_GetParam(t *testing.T) {

	url, _ := url.Parse("https://example.org/x?a=1&a=2&b=3&=3&&&&")
	fmt.Printf("[a]=%s\n", GetParams(url, "a"))
	fmt.Printf("[b]=%s\n", GetParam(url, "b"))
	fmt.Printf("[c]=%s\n", GetParam(url, "c"))
}
