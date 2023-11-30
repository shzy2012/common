package tool

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"math/rand"
	"regexp"
	"strings"
	"time"

	"github.com/shzy2012/common/log"
)

var r *rand.Rand

func init() {
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
}

// StringBuilderBody 对string进行链式处理
type StringBuilderBody struct {
	origin string
}

// StringBuilder new stringbuilder
func StringBuilder(origin string) *StringBuilderBody {
	return &StringBuilderBody{
		origin: origin,
	}
}

// Replace 替换函数
func (s *StringBuilderBody) Replace(old, new string) *StringBuilderBody {
	s.origin = strings.Replace(s.origin, old, new, -1)
	return s
}

// Build 返回字符串
func (s *StringBuilderBody) Build() string {
	return s.origin
}

// GetRandomString 获取随机字符串
func GetRandomString(length uint64) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	for i := uint64(0); i < length; i++ {
		result = append(result, bytes[r.Intn(int(len(bytes)))])
	}
	return string(result)
}

// Join 返回包含引号("")的字符串
func Join(a []string, sep string) string {
	switch len(a) {
	case 0:
		return ""
	case 1:
		return "\"" + a[0] + "\""
	}
	n := len(sep) * (len(a) - 1)
	for i := 0; i < len(a); i++ {
		n += len(a[i]) + 2
	}

	var b strings.Builder
	b.Grow(n)
	b.WriteString("\"" + a[0] + "\"")
	for _, s := range a[1:] {
		b.WriteString(sep)
		b.WriteString("\"" + s + "\"")
	}

	return b.String()
}

func ToHex(src []byte) string {
	return hex.EncodeToString(src)
}

func ToHex2(src string) string {
	return fmt.Sprintf("%x", src)
}

func Base64Encode(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

func Base64Decode(str string) string {
	dist, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		log.Infof("%s\n", err.Error())
	}
	return string(dist)
}

// https://juejin.cn/post/6844903648045039624
func IsChinesePhone(phone string) bool {
	reg1 := regexp.MustCompile(`^1(?:3[0-9]|4[5-9]|5[0-9]|6[12456]|7[0-8]|8[0-9]|9[0-9])[0-9]{8}$`)
	if reg1 == nil {
		return false
	}
	//根据规则提取关键信息
	if len(reg1.FindAllStringSubmatch(phone, 1)) > 0 {
		return true
	}

	return false
}
