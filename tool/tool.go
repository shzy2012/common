package tool

import (
	"sync/atomic"
)

//Tool 结构体
type Tool struct {
	counter uint64
}

//T 创建Tool
func T() *Tool {
	return &Tool{}
}

//SetCounter 设置自增起始点
func (t *Tool) SetCounter(base uint64) {
	t.counter = base
}

//AutoIncrementing 自动自增
func (t *Tool) AutoIncrementing() uint64 {
	return atomic.AddUint64(&t.counter, 1)
}
