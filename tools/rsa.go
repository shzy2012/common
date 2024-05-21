package tools

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"os"
)

const (
	PKCS1 = "PKCS1"
	PKCS8 = "PKCS8"
)

// 生成密钥对(PKCS1格式)
func RSAGenKey(bits int) error {
	/*
		1.生成私钥
	*/
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return err
	}

	priByte := x509.MarshalPKCS1PrivateKey(privateKey)
	block1 := pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: priByte,
	}

	privatePem, err := os.Create("private_key.pem")
	if err != nil {
		return err
	}
	defer privatePem.Close()
	err = pem.Encode(privatePem, &block1)
	if err != nil {
		return err
	}

	/*
		2.生成公钥
	*/
	pubKey := privateKey.PublicKey
	publicBytes := x509.MarshalPKCS1PublicKey(&pubKey)
	block2 := pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: publicBytes,
	}
	pubPem, err := os.Create("public_key.pem")
	if err != nil {
		return err
	}
	defer pubPem.Close()
	pem.Encode(pubPem, &block2)
	return nil
}

// 加密(PKCS1格式)
/*
pubkey: 公钥,base64字符串,不带格式
plaintext: 明文

example:

	pubkey := `
MIGJAoGBAPKvW8kNX1UjX3l98LJeBGnxz7piPsVuHafcE5ygP+y+Q6AZD/bJHiMX
TP7fe7tC+AJMprH6Hjn596YaMT23Hn3xcAy/5PoJlsY0MRXjB7vnGWTSHvEe5ji/
o7KxRIFMF0yjmR+WwqTn6FuTXdpRemMybnRrYbSlMWf+yTvcy7xVAgMBAAE=
`
	plaintext := "abc"

*/
func RsaEncrypt(pubkey, plaintext string) (string, error) {
	pubkeyBytes, err := base64.StdEncoding.DecodeString(pubkey)
	if err != nil {
		return "", err
	}

	// go语言对格式有严格要求,使用pem构造避免格式问题
	pubPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: pubkeyBytes,
	})

	// 解析 PEM 结构
	block, _ := pem.Decode(pubPEM)
	if block == nil || block.Type != "RSA PUBLIC KEY" {
		return "", errors.New("failed to decode PEM block containing public key")
	}

	/* ParsePKIXPublicKey  与 ParsePKCS1PublicKey 区别
	ParsePKIXPublicKey: 函数用于解析PKIX格式的公钥。PKIX（Public Key Infrastructure using X.509）是一种公钥基础设施标准，通常使用X.509证书来描述公钥。
						这个函数可以解析DER编码或PEM编码的PKIX格式的公钥。输出的 PublicKey 类型是一个接口。

	ParsePKCS1PublicKey: 函数用于解析PKCS#1格式的RSA公钥。PKCS#1（Public Key Cryptography Standards #1）是一种专门为RSA密钥设计的格式，它定义了私钥和公钥的结构。
						 这个函数可以解析DER编码或PEM编码的PKCS#1格式的公钥。输出的是 *rsa.PublicKey 类型，表示RSA公钥。
	*/
	// 有 ParsePKIXPublicKey 和 ParsePKCS1PublicKey 两种格式
	pub, err := x509.ParsePKIXPublicKey(block.Bytes) //兼容更好
	// pub, err := x509.ParsePKCS1PublicKey(block.Bytes)
	if err != nil {
		return "", err
	}

	ciphertext, err := rsa.EncryptPKCS1v15(rand.Reader, pub.(*rsa.PublicKey), []byte(plaintext))
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

/*
	解密

prikey: 私钥,base64格式,不带格式
base64Ciphertext: base64格式的密文

example:

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
*/
func RsaDecode(prikey string, base64Ciphertext string) (string, error) {

	pubkeyBytes, err := base64.StdEncoding.DecodeString(prikey)
	if err != nil {
		return "", err
	}

	pubPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: pubkeyBytes,
	})

	block, _ := pem.Decode(pubPEM)
	if block == nil {
		return "", fmt.Errorf("failed to parse PEM block containing the private key")
	}

	// privateKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return "", err
	}

	// base64解密
	ciphertext, err := base64.StdEncoding.DecodeString(base64Ciphertext)
	if err != nil {
		return "", err
	}

	text, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, ciphertext)
	if err != nil {
		return "", err
	}

	return string(text), nil
}

/* PKIX格式 与 PKCS#1格式
PKIX（Public Key Infrastructure using X.509）格式和PKCS#1（Public Key Cryptography Standards #1）格式是两种不同的标准，用于表示公钥的结构。这两者之间的主要区别在于它们的设计目标和用途。

PKIX格式：
描述： PKIX是一个公共密钥基础设施（PKI）标准，它定义了使用X.509证书表示的公钥结构。
基础： PKIX基于X.509标准，该标准定义了用于在公共网络上进行通信的数字证书的格式。
用途： PKIX格式通常用于在互联网上建立信任链，支持SSL/TLS协议以及其他需要数字证书的应用程序。
证书： PKIX格式的公钥通常包含在X.509数字证书中，该证书除了包含公钥外，还包含有关所有者身份的信息和数字签名。
编码： PKIX格式的公钥通常采用DER编码或PEM编码。

PKCS#1格式：
描述： PKCS#1是一系列与公钥密码学有关的标准之一，其中包括了定义RSA公钥和私钥结构的规范。
基础： PKCS#1是由RSA Security公司提出的一系列标准，最初是为了支持RSA算法的使用而设计的。
用途： PKCS#1格式主要用于表示RSA密钥，包括RSA公钥和私钥。它通常用于数字签名、加密等与RSA密钥相关的操作。
结构： PKCS#1定义了RSA公钥和私钥的结构，包括了模数（Modulus）、指数（Exponent）等。
编码： PKCS#1格式的公钥通常采用DER编码，而私钥可能以DER编码或PEM编码。

PKIX是一个更广泛的标准，旨在建立可信任的PKI体系，它的应用不仅仅局限于RSA算法。
PKCS#1是专门用于表示RSA密钥的标准，适用于RSA算法的各种操作。
在实际使用中，PKIX格式的公钥通常包含在X.509数字证书中，而PKCS#1格式的公钥通常独立存在。


DER与PEM区别:
  DER编码是二进制格式
  PEM编码是Base64编码的文本格式
*/
