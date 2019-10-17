package tool

import (
	"strings"
)

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
