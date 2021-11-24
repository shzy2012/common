package tool

import (
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
)

// EncodeMD5 md5 encryption
func EncodeMD5(value string) string {
	m := md5.New()
	m.Write([]byte(value))

	return hex.EncodeToString(m.Sum(nil))
}

//Sha512 encryption
//Package sha512 implements the SHA-384, SHA-512, SHA-512/224, and SHA-512/256 hash algorithms as defined in FIPS 180-4.
func Sha512(value string) string {
	m := sha512.New()
	m.Write([]byte(value))

	return hex.EncodeToString(m.Sum(nil))
}

//Sha256 encryption
func Sha256(value string) string {
	m := sha256.New()
	m.Write([]byte(value))

	return hex.EncodeToString(m.Sum(nil))
}
