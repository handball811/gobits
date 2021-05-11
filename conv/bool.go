package conv

import (
	"github.com/handball811/gobits"
)

type BoolReadWriter interface {
	BoolReader
	BoolWriter
}

type BoolReader interface {
	Read(r gobits.Reader, b *bool) error
}

type BoolWriter interface {
	Write(w gobits.Writer, b bool) error
}

type OneBoolReader struct {
	BoolReadWriter
	b []byte
	s *gobits.Slice
}

func NewOneBoolReader() *OneBoolReader {
	b := make([]byte, 1)
	return &OneBoolReader{
		b: b,
		s: gobits.NewSliceWithBuffer(b),
	}
}

func (c *OneBoolReader) Read(r gobits.Reader, b *bool) error {
	size, err := r.Read(c.s)
	if err != nil {
		return err
	}
	*b = size == 1 && c.b[0] != 0
	return nil
}

type OneBoolWriter struct {
	BoolReadWriter
	b []byte
	s *gobits.Slice
}

func NewOneBoolWriter() *OneBoolWriter {
	b := make([]byte, 1)
	return &OneBoolWriter{
		b: b,
		s: gobits.NewSliceWithBuffer(b),
	}
}

func (c *OneBoolWriter) Write(w gobits.Writer, b bool) error {
	c.b[0] = 0
	if b {
		c.b[0] = 1
	}
	_, err := w.Write(c.s)
	return err
}
