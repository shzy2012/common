package tools

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
)

/*
AES（Advanced Encryption Standard，高级加密标准）是一种对称加密算法，用于保护电子数据。它由美国国家标准与技术研究院（NIST）于2001年发布，取代了之前的DES（Data Encryption Standard）。
	优点：算法公开、计算量小、加密速度快、加密效率高。
	缺点：发送方和接收方必须商定好密钥，然后使双方都能保存好密钥，密钥管理成为双方的负担。

Secret Key Length	AES Variant	Block Size
16 bytes (128 bits)	AES-128	16 bytes (128 bits)
24 bytes (192 bits)	AES-192	16 bytes (128 bits)
32 bytes (25 6bits)	AES-256	16 bytes (128 bits)

主要特点：
	对称加密：AES使用相同的密钥进行加密和解密，这意味着密钥必须在发送方和接收方之间安全共享。
	块加密：AES是一种分组密码算法，处理固定大小的数据块。标准块大小为128位。
	密钥长度：AES支持三种密钥长度：128位、192位和256位。密钥长度越长，安全性越高，但处理速度可能稍慢。
	安全性：AES被认为是非常安全的，目前没有已知的有效攻击方法能够在合理时间内破解AES。
	广泛应用：AES被广泛应用于各种安全协议和应用程序中，如TLS/SSL、VPN、文件加密等。

工作模式：
AES通常与不同的工作模式结合使用，以增强安全性和灵活性。常见的模式包括：
	ECB（电子密码本模式）：简单但不安全，因为相同的明文块会生成相同的密文块。
	CBC（密码分组链接模式）：每个明文块在加密前与前一个密文块进行异或操作，增加了安全性。
	CFB（密码反馈模式）和OFB（输出反馈模式）：将块密码转换为流密码。
	GCM（Galois/Counter Mode）：提供认证加密，确保数据的完整性和真实性。
	AES因其高效性和安全性，成为全球范围内的加密标准，被广泛应用于保护敏感数据
*/

// https://dev.to/breda/secret-key-encryption-with-go-using-aes-316d

// 加密
func AesEncrypt(secretKey []byte, plaintext string) string {
	aes, err := aes.NewCipher(secretKey)
	if err != nil {
		panic(err)
	}

	gcm, err := cipher.NewGCM(aes)
	if err != nil {
		panic(err)
	}

	// We need a 12-byte nonce for GCM (modifiable if you use cipher.NewGCMWithNonceSize())
	// A nonce should always be randomly generated for every encryption.
	nonce := make([]byte, gcm.NonceSize())
	_, err = rand.Read(nonce)
	if err != nil {
		panic(err)
	}

	// ciphertext here is actually nonce+ciphertext
	// So that when we decrypt, just knowing the nonce size
	// is enough to separate it from the ciphertext.
	ciphertext := gcm.Seal(nonce, nonce, []byte(plaintext), nil)

	return string(ciphertext)
}

// 解密
func AesDecrypt(secretKey []byte, ciphertext string) string {
	aes, err := aes.NewCipher(secretKey)
	if err != nil {
		panic(err)
	}

	gcm, err := cipher.NewGCM(aes)
	if err != nil {
		panic(err)
	}

	// Since we know the ciphertext is actually nonce+ciphertext
	// And len(nonce) == NonceSize(). We can separate the two.
	nonceSize := gcm.NonceSize()
	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]

	plaintext, err := gcm.Open(nil, []byte(nonce), []byte(ciphertext), nil)
	if err != nil {
		panic(err)
	}

	return string(plaintext)
}
