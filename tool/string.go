package tool

import (
	"math/rand"
	"strings"
	"time"
)

var r *rand.Rand

func init() {
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
}

//StringBuilderBody 对string进行链式处理
type StringBuilderBody struct {
	origin string
}

//StringBuilder new stringbuilder
func StringBuilder(origin string) *StringBuilderBody {
	return &StringBuilderBody{
		origin: origin,
	}
}

//Replace 替换函数
func (s *StringBuilderBody) Replace(old, new string) *StringBuilderBody {
	s.origin = strings.Replace(s.origin, old, new, -1)
	return s
}

//Build 返回字符串
func (s *StringBuilderBody) Build() string {
	return s.origin
}

//GetRandomString 获取随机字符串
func GetRandomString(length uint64) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	for i := uint64(0); i < length; i++ {
		result = append(result, bytes[r.Intn(int(len(bytes)))])
	}
	return string(result)
}
