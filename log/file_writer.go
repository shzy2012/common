package log

import (
	"fmt"
	"os"
	"path"
	"sync"
	"time"
)

type FileWriter struct {
	mu       sync.Mutex
	file     *os.File
	rootPath string
}

func NewFileWriter(rootPath string) (*FileWriter, error) {
	return &FileWriter{
		rootPath: rootPath,
	}, nil
}

// Write 实现 io.Writer 接口
func (w *FileWriter) Write(p []byte) (n int, err error) {
	w.mu.Lock()
	defer w.mu.Unlock()

	filePath := path.Join(w.rootPath, fmt.Sprintf("%s.log", time.Now().Format("20060102")))
	_, err = os.Stat(filePath)
	if os.IsNotExist(err) || w.file == nil {
		if w.file != nil {
			w.file.Close()
		}
		w.file, err = os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			return 0, fmt.Errorf("failed to open log file: %v", err)
		}
	}

	return w.file.Write(p)
}

// Close 关闭文件
func (w *FileWriter) Close() error {
	w.mu.Lock()
	defer w.mu.Unlock()

	if w.file != nil {
		return w.file.Close()
	}
	return nil
}
