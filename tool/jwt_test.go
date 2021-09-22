package tool

import (
	"fmt"
	"testing"
	"time"
)

func TestGenerateToken(t *testing.T) {
	username := "admin"
	password := "123456"
	token, _ := GenerateToken(username, password)
	fmt.Println(token)
}

func TestGenerateTokenForOutgoing(t *testing.T) {
	username := "admin"
	password := GetRandomString(25)
	expired := time.Now().Add(3 * time.Hour)
	token, _ := GenerateTokenForOutgoing(username, password, expired.Unix())
	fmt.Println(token)
}

func TestParseToken(t *testing.T) {
	token := `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6IjIxMjMyZjI5N2E1N2E1YTc0Mzg5NGEwZTRhODAxZmMzIiwicGFzc3dvcmQiOiIxODM2MmY5YWYxYTY2OTIwYzNkOTJmYmQ5N2M3NGZlMyIsImV4cCI6MTYxMTU3MDE4NywiaXNzIjoidGoiLCJzdWIiOiI0T0cifQ.8w0MxfkeRsT4WK8tQbDyFUI9DVOEn6bUHLJ4CUceIlg`
	claims, err := ParseToken(token)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
	}
	fmt.Printf("%+v\n", claims)
}
