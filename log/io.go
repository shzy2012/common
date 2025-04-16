package log

import (
	"fmt"
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
	var errs []error
	for _, w := range mw.writers {
		_, err := w.Write(p)
		if err != nil {
			errs = append(errs, err)
		}
	}
	if len(errs) > 0 {
		return 0, fmt.Errorf("multiple errors: %v", errs)
	}
	return len(p), nil
}
