package tool

import (
	"math/rand"
	"sync/atomic"
	"time"
)

var r *rand.Rand
var counter uint64

func init() {
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
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

//SetCounter 设置自增起始点
func SetCounter(base uint64) {
	counter = base
}

//AutoIncrementing 自动自增
func AutoIncrementing() uint64 {
	return atomic.AddUint64(&counter, 1)
}
