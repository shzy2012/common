package tools

import (
	"math/rand"
	"time"
)

//Tools 工具集
type Tools struct {
	rand *rand.Rand
}

//NewTools 创建 Tools
func NewTools() *Tools {
	return &Tools{
		rand: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

//GetRandomString 获取随机字符串
func (t *Tools) GetRandomString(length int64) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	for i := int64(0); i < length; i++ {
		result = append(result, bytes[t.rand.Intn(int(len(bytes)))])
	}
	return string(result)
}
