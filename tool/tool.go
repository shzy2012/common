package tool

import (
	"math/rand"
	"sync/atomic"
	"time"
)

//Tool 工具集
type Tool struct {
	rand    *rand.Rand
	counter uint64
}

//NewTool 创建 Tool
func NewTool() *Tool {
	return &Tool{
		rand: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

//GetRandomString 获取随机字符串
func (t *Tool) GetRandomString(length uint64) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	for i := uint64(0); i < length; i++ {
		result = append(result, bytes[t.rand.Intn(int(len(bytes)))])
	}
	return string(result)
}

//SetCounter 设置自增起始点
func (t *Tool) SetCounter(base uint64) {
	t.counter = base
}

//AutoIncrementing 自动自增
func (t *Tool) AutoIncrementing() uint64 {
	return atomic.AddUint64(&t.counter, 1)
}
