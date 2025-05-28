package log

import (
	"io"
	"log"
)

// CustomMultiWriter 是一个自定义的 io.MultiWriter
type CustomMultiWriter struct {
	writers []io.Writer
}

// NewCustomMultiWriter 创建一个新的 CustomMultiWriter
func NewCustomMultiWriter(writers ...io.Writer) *CustomMultiWriter {
	return &CustomMultiWriter{writers: writers}
}

// Write 实现了 io.Writer 接口，将数据写入所有的 writer
func (mw *CustomMultiWriter) Write(p []byte) (n int, err error) {
	for _, w := range mw.writers {
		n, err = w.Write(p)
		if err != nil {
			log.Println("write error", err)
			// return n, err
			continue
		}
		if n != len(p) {
			log.Println("short write", n, len(p))
			// return n, io.ErrShortWrite
			continue
		}
	}
	return len(p), nil
}
