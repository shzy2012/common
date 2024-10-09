package tools

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

// 获取随机字符串
func GetRandomString(length uint64) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	for i := uint64(0); i < length; i++ {
		result = append(result, bytes[r.Intn(int(len(bytes)))])
	}
	return string(result)
}

// 获取随机数字
func GetRandomNumber(length uint64) string {
	str := "123456789"
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

// 使用多种符号分割字符串
// text:字符串
// symbols:分隔符.比如匹配逗号、分号: ,;
func Splits(text, symbols string) []string {
	re := regexp.MustCompile(fmt.Sprintf(`[%s]+`, symbols))
	// 使用正则表达式进行分割
	return re.Split(text, -1)
}

/*
s:字符串
fillChar:填充符号
width:目标宽度
*/
func PaddingFromLeft(s string, fillChar string, width int) string {
	if len(s) < width {
		return strings.Repeat(fillChar, width-len(s)) + s
	}

	return s
}

/*
s:字符串
fillChar:填充符号
width:目标宽度
*/
func PaddingFromRight(s string, fillChar string, width int) string {
	// 自定义填充函数
	if len(s) < width {
		return s + strings.Repeat(fillChar, width-len(s))
	}
	return s
}
