package gobits

import (
	"errors"
)

var (
	LowByteFilter = []byte{
		0,
		0b00000001,
		0b00000011,
		0b00000111,
		0b00001111,
		0b00011111,
		0b00111111,
		0b01111111,
		0b11111111,
	}
	HighByteFilter = []byte{
		0b11111111,
		0b11111110,
		0b11111100,
		0b11111000,
		0b11110000,
		0b11100000,
		0b11000000,
		0b10000000,
		0b00000000,
	}
	InlineByteFilter  = make([]byte, 128)
	OutlineByteFilter = make([]byte, 128)
	ErrOutOfBound     = errors.New("Try to copy out of range")
)

func init() {
	for i := 0; i <= 8; i++ {
		for j := 0; j <= 8-i; j++ {
			InlineByteFilter[i<<3|j] = LowByteFilter[j+i] & HighByteFilter[j]
			OutlineByteFilter[i<<3|j] = HighByteFilter[j+i] | LowByteFilter[j]
		}
	}
}

type Writer interface {
	// offset, size [bit]
	Write(b *Slice) (int, error)
}

type WriterTo interface {
	WriteTo(s Writer) error
}

type Reader interface {
	// offset, size [bit]
	Read(b *Slice) (int, error)
}

type ReaderFrom interface {
	ReadFrom(s Reader) error
}

type ReadWriter interface {
	Reader
	Writer
}
