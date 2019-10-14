package tool

import (
	"math/rand"
	"time"
)

//Tool 工具集
type Tool struct {
	rand *rand.Rand
}

//NewTool 创建 Tool
func NewTool() *Tool {
	return &Tool{
		rand: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

//GetRandomString 获取随机字符串
func (t *Tool) GetRandomString(length int64) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	for i := int64(0); i < length; i++ {
		result = append(result, bytes[t.rand.Intn(int(len(bytes)))])
	}
	return string(result)
}
