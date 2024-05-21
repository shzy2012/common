package tools

import (
	"fmt"
	"testing"
)

func TestRSAGenKey(t *testing.T) {

	// 生成密钥对
	// 密钥位数: 512\1024\2048\4096 bit等
	// 密钥格式: PKCS#1 PKCS#8
	err := RSAGenKey(1024)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func TestRsaEncrypt(t *testing.T) {

	pubkey := `
MFwwDQYJKoZIhvcNAQEBBQADSwAwSAJBALBGQ6yVSVjLNSompRSrkDReWYNdY4pP
	OoiabtCrDB+yKhMIIIHLJi5Z89cqvqjvkMvrXxXzl056IOQNM+Wdh60CAwEAAQ==
`
	plaintext := "abc"
	encryptText, err := RsaEncrypt(pubkey, plaintext)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("\n%s\n", encryptText)
}

func TestRsaDecode(t *testing.T) {

	prikey := `
MIICXQIBAAKBgQDyr1vJDV9VI195ffCyXgRp8c+6Yj7Fbh2n3BOcoD/svkOgGQ/2
yR4jF0z+33u7QvgCTKax+h45+femGjE9tx598XAMv+T6CZbGNDEV4we75xlk0h7x
HuY4v6OysUSBTBdMo5kflsKk5+hbk13aUXpjMm50a2G0pTFn/sk73Mu8VQIDAQAB
AoGALEUG1kclM8+vE+eAZ9k0rurYfOR9FODAciV4QmMNJi+TAHpx6g/H+pi+h+PW
m1NdEHZRFjhNGUBbB6bRgrOL0QywZkDi41ZUQV0plXxYx9YpWzncaFK9vxzzjKb3
GLSxWsikNl2yGecJQ+lk4TaXIAGG2sCwooi1XOOlPCtXvQECQQD7Dt4E1wKNcSpw
BPKaKxElO8XUXeWRmnzOpFT0MJxfKTtABydnVTBg1iaIURyYdGoHCDc553rAeYa4
L1H6+PSVAkEA93ZMMJHOQlFBiq2MRvNRl5dP0gONS0EzpmvKuRQogu5hAUKhvYoz
bV8E3qns838pZoaA3r+PBtUYx52fQ7/4wQJAXHW1PoMQ5ZZv0qF/11dVESlaSkPq
cB09Kb1LrELa1BETSRlZYaz2DDPSLRHyPhNhmQVlkWW2x3v6KYsD3jIhoQJBAO44
8auoEXmCI6hO3cXHovpd7bdtN+4EPKavCh8VqtIwjS3baTy/+DYHzPZVewgFmGNc
hF7q5dNb/VjdAl8ERYECQQDdYVV+YRUTIf6NOGFQJmKu+6DSLVbWSUwnvP70NbJw
XzpyxgRy9LOLORMuKY03QVhcPJqYBc8sf63kNOAneFRY
`
	encryptText := `PaIv3mpRbu2MnMZkmRtlmVctRthAd7fZprAANzls4HFHcuim9zMZqkO/ZypbfD0vJ+aHbrEZcaZAlPum9wfKhZgXZfs84jonG507B3kDKuLxsn21IkqKFCZ7XdT+cZ3jERdjzBIEC4wwgUT58ryuIlIiauCisa9ZU5kREpCzv2o=`
	text, err := RsaDecode(prikey, encryptText)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%s\n", text)
}
