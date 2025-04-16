package log

import (
	"io"
)

// CustomMultiWriter is a struct that implements the io.Writer interface.
type CustomMultiWriter struct {
	writers []io.Writer
}

// NewCustomMultiWriter creates a new CustomMultiWriter.
func NewCustomMultiWriter(writers ...io.Writer) *CustomMultiWriter {
	return &CustomMultiWriter{writers: writers}
}

// Write writes data to all writers and continues even if one fails.
func (mw *CustomMultiWriter) Write(p []byte) (n int, err error) {
	for _, w := range mw.writers {
		w.Write(p)
	}
	return len(p), nil
}
