package gobits_test

import (
	"errors"
	"testing"

	. "github.com/handball811/gobits"
	"github.com/stretchr/testify/assert"
)

func TestCopy0(t *testing.T) {
	source := []byte{0, 1, 2, 3, 4, 5}
	destination := make([]byte, 6)
	dest := make([]int, 6)
	err := Copy(destination, 15, source, 7, 16)
	for i, d := range destination {
		dest[i] = int(d)
	}
	assert.Nil(t, err)
	assert.Equal(t, 0, dest[0])
	assert.Equal(t, 0, dest[1])
	assert.Equal(t, 1, dest[2])
	assert.Equal(t, 2, dest[3])
	assert.Equal(t, 0, dest[4])
	assert.Equal(t, 0, dest[5])
}

func TestCopy1(t *testing.T) {
	source := []byte{0, 1, 2, 3, 4, 5}
	destination := make([]byte, 6)
	dest := make([]int, 6)
	err := Copy(destination, 16, source, 8, 16)
	for i, d := range destination {
		dest[i] = int(d)
	}
	assert.Nil(t, err)
	assert.Equal(t, 0, dest[0])
	assert.Equal(t, 0, dest[1])
	assert.Equal(t, 1, dest[2])
	assert.Equal(t, 2, dest[3])
	assert.Equal(t, 0, dest[4])
	assert.Equal(t, 0, dest[5])
}

func TestCopy2(t *testing.T) {
	source := []byte{0, 1, 2, 3, 4, 5}
	destination := make([]byte, 6)
	dest := make([]int, 6)
	err := Copy(destination, 4, source, 1, 24)
	for i, d := range destination {
		dest[i] = int(d)
	}
	assert.Nil(t, err)
	assert.Equal(t, 0, dest[0])
	assert.Equal(t, 8, dest[1])
	assert.Equal(t, 16, dest[2])
	assert.Equal(t, 8, dest[3])
	assert.Equal(t, 0, dest[4])
	assert.Equal(t, 0, dest[5])
}

func TestCopy3(t *testing.T) {
	source := []byte{0, 1, 2, 3, 4, 5}
	destination := make([]byte, 6)
	dest := make([]int, 6)
	err := Copy(destination, 1, source, 4, 21)
	for i, d := range destination {
		dest[i] = int(d)
	}
	assert.Nil(t, err)
	assert.Equal(t, 32, dest[0])
	assert.Equal(t, 64, dest[1])
	assert.Equal(t, 32, dest[2])
	assert.Equal(t, 0, dest[3])
	assert.Equal(t, 0, dest[4])
	assert.Equal(t, 0, dest[5])
}

func TestCopy4(t *testing.T) {
	source := []byte{0, 1, 2, 3, 4, 5}
	destination := make([]byte, 6)
	dest := make([]int, 6)
	err := Copy(destination, 0, source, 0, 20)
	for i, d := range destination {
		dest[i] = int(d)
	}
	assert.Nil(t, err)
	assert.Equal(t, 0, dest[0])
	assert.Equal(t, 1, dest[1])
	assert.Equal(t, 2, dest[2])
	assert.Equal(t, 0, dest[3])
	assert.Equal(t, 0, dest[4])
	assert.Equal(t, 0, dest[5])
}

func TestCopy5(t *testing.T) {
	source := []byte{0, 255, 2, 3, 4, 5}
	destination := make([]byte, 6)
	dest := make([]int, 6)
	err := Copy(destination, 1, source, 1, 9)
	for i, d := range destination {
		dest[i] = int(d)
	}
	assert.Nil(t, err)
	assert.Equal(t, 0, dest[0])
	assert.Equal(t, 3, dest[1])
	assert.Equal(t, 0, dest[2])
	assert.Equal(t, 0, dest[3])
	assert.Equal(t, 0, dest[4])
	assert.Equal(t, 0, dest[5])
}

func TestCopyError(t *testing.T) {
	source := []byte{0, 1, 2, 3, 4, 5}
	destination := make([]byte, 6)
	err := Copy(destination, 1, source, 4, 44)
	assert.Nil(t, err)
	err = Copy(destination, 1, source, 4, 45)
	assert.NotNil(t, err)
	err = Copy(destination, 4, source, 1, 44)
	assert.Nil(t, err)
	err = Copy(destination, 4, source, 1, 45)
	assert.NotNil(t, err)
}

func TestSegmentWrite(t *testing.T) {
	var err error
	var size int
	segment := NewSegmentWithSize(8)
	slice0 := NewSliceWithBuffer([]byte{1, 2, 3})
	slice1 := NewSliceWithBuffer([]byte{4, 5, 6})
	slice2 := NewSliceWithBuffer([]byte{7, 8, 9})

	size, err = segment.Write(slice0)
	assert.Nil(t, err)
	assert.Equal(t, 24, size)
	assert.Equal(t, 24, segment.Len())

	size, err = segment.Write(slice1)
	assert.Nil(t, err)
	assert.Equal(t, 24, size)
	assert.Equal(t, 48, segment.Len())

	size, err = segment.Write(slice2)
	assert.Nil(t, err)
	assert.Equal(t, 16, size)
	assert.Equal(t, 64, segment.Len())
}

func TestSegmentRead(t *testing.T) {
	var err error
	var size int
	segment := NewSegmentWithSize(6)
	slice0 := NewSliceWithBuffer([]byte{1, 2, 3, 4, 5, 6, 7, 8, 9})
	segment.Write(slice0)
	readBuffer := make([]byte, 64)
	readSlice := NewSlice(readBuffer, 8, 40, 40)

	size, err = segment.Read(readSlice)
	assert.Nil(t, err)
	assert.Equal(t, 16, segment.Len())
	assert.Equal(t, byte(0), readBuffer[0])
	assert.Equal(t, byte(1), readBuffer[1])
	assert.Equal(t, byte(2), readBuffer[2])
	assert.Equal(t, byte(3), readBuffer[3])
	assert.Equal(t, byte(4), readBuffer[4])
	assert.Equal(t, byte(0), readBuffer[5])

	size, err = segment.Read(readSlice)
	assert.Nil(t, err)
	assert.Equal(t, 0, segment.Len())
	assert.Equal(t, 16, size)
	assert.Equal(t, byte(5), readBuffer[1])
	assert.Equal(t, byte(6), readBuffer[2])
	assert.Equal(t, byte(3), readBuffer[3])
}

///
/// ReadFromTest
///
type mockReader struct {
	ret int
	err error
}

func (m *mockReader) Read(b *Slice) (int, error) {
	return m.ret, m.err
}

func TestSegmentReadFromNormal(t *testing.T) {
	var err error
	segment := NewSegmentWithSize(8)
	mock := &mockReader{
		ret: 23,
		err: nil,
	}
	err = segment.ReadFrom(mock)
	assert.Equal(t, 23, segment.Len())
	assert.Nil(t, err)
}

func TestSegmentReadFromWithError(t *testing.T) {
	var err error
	segment := NewSegmentWithSize(8)
	mock := &mockReader{
		ret: 0,
		err: errors.New("Read Error"),
	}
	err = segment.ReadFrom(mock)
	assert.Equal(t, 0, segment.Len())
	assert.Equal(t, mock.err, err)
}

func TestSegmentReadFromWithErrorAndSize(t *testing.T) {
	var err error
	segment := NewSegmentWithSize(8)
	mock := &mockReader{
		ret: 23,
		err: errors.New("Read Error"),
	}
	err = segment.ReadFrom(mock)
	assert.Equal(t, 23, segment.Len())
	assert.Equal(t, mock.err, err)
}

///
/// WriteToTest
///
type mockWriter struct {
	ret int
	err error
}

func (m *mockWriter) Write(b *Slice) (int, error) {
	return m.ret, m.err
}

func TestSegmentWriteToNormal(t *testing.T) {
	var err error
	segment := NewSegment(NewSliceWithBuffer([]byte{1, 2, 3, 4, 5, 6, 7, 8, 9}))
	mock := &mockWriter{
		ret: 23,
		err: nil,
	}
	err = segment.WriteTo(mock)
	assert.Equal(t, 49, segment.Len())
	assert.Nil(t, err)
}

func TestSegmentWriteWithError(t *testing.T) {
	var err error
	segment := NewSegment(NewSliceWithBuffer([]byte{1, 2, 3, 4, 5, 6, 7, 8, 9}))
	mock := &mockWriter{
		ret: 0,
		err: errors.New("Write Error"),
	}
	err = segment.WriteTo(mock)
	assert.Equal(t, 72, segment.Len())
	assert.Equal(t, mock.err, err)
}

func TestSegmentWriteWithErrorAndSize(t *testing.T) {
	var err error
	segment := NewSegment(NewSliceWithBuffer([]byte{1, 2, 3, 4, 5, 6, 7, 8, 9}))
	mock := &mockWriter{
		ret: 23,
		err: errors.New("Write Error"),
	}
	err = segment.WriteTo(mock)
	assert.Equal(t, 49, segment.Len())
	assert.Equal(t, mock.err, err)
}

func BenchmarkCopyOrdered1024Byte(b *testing.B) {
	source := make([]byte, 4096)
	destination := make([]byte, 4096)
	copySize := 1024 * 8
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Copy(destination, 0, source, 0, copySize)
	}
}

func BenchmarkCopyOrdered16Byte(b *testing.B) {
	source := make([]byte, 4096)
	destination := make([]byte, 4096)
	copySize := 16 * 8
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Copy(destination, 0, source, 0, copySize)
	}
}

func BenchmarkCopyPartial1024Byte(b *testing.B) {
	source := make([]byte, 4096)
	destination := make([]byte, 4096)
	copySize := 1024 * 8
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Copy(destination, 0, source, 5, copySize)
	}
}

func BenchmarkCopyPartial16Byte(b *testing.B) {
	source := make([]byte, 4096)
	destination := make([]byte, 4096)
	copySize := 16 * 8
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Copy(destination, 0, source, 5, copySize)
	}
}

// Collect 127bit*64 -> readBuffer
// 2453 ns/op	       0 B/op	       0 allocs/op
func BenchmarkSegmentCollection(b *testing.B) {
	writeBuffers := make([]*Slice, 64)
	for i := range writeBuffers {
		writeBuffers[i] = NewSlice(make([]byte, 16), 0, 127, 128)
	}
	segment := NewSegmentWithSize(1300)
	readSegment := NewSegmentWithSize(1300)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		segment.Reset()
		for _, buffer := range writeBuffers {
			segment.Write(buffer)
		}
		segment.WriteTo(readSegment)
	}
}

// Collect 2047bit*4 -> readBuffer
// 1501 ns/op	       0 B/op	       0 allocs/op
func BenchmarkSegmentCollection2(b *testing.B) {
	writeBuffers := make([]*Slice, 4)
	for i := range writeBuffers {
		writeBuffers[i] = NewSlice(make([]byte, 256), 0, 2047, 2048)
	}
	segment := NewSegmentWithSize(1300)
	readSegment := NewSegmentWithSize(1300)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		segment.Reset()
		for _, buffer := range writeBuffers {
			segment.Write(buffer)
		}
		segment.WriteTo(readSegment)
	}
}
