package log

import (
	"bytes"
	"io"
)

// CustomMultiWriter is a struct that implements the io.Writer interface.
type CustomMultiWriter struct {
	writers []io.Writer
	buffer  *bytes.Buffer
}

// NewCustomMultiWriter creates a new CustomMultiWriter.
func NewCustomMultiWriter(writers ...io.Writer) *CustomMultiWriter {
	return &CustomMultiWriter{
		writers: writers,
		buffer:  bytes.NewBuffer(nil),
	}
}

// Write writes data to all writers immediately.
func (mw *CustomMultiWriter) Write(p []byte) (n int, err error) {
	for _, w := range mw.writers {
		n, err = w.Write(p)
		if err != nil {
			return n, err
		}
	}
	return len(p), nil
}

// Flush writes the buffered data to all writers and clears the buffer.
func (mw *CustomMultiWriter) Flush() error {
	data := mw.buffer.Bytes()
	for _, w := range mw.writers {
		_, err := w.Write(data)
		if err != nil {
			return err
		}
	}
	mw.buffer.Reset()
	return nil
}

// Reset clears the buffer.
func (mw *CustomMultiWriter) Reset() {
	mw.buffer.Reset()
}
