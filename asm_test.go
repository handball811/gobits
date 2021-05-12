package gobits_test

import (
	"testing"

	. "github.com/handball811/gobits"
	"github.com/stretchr/testify/assert"
)

func TestAsmSlideCopy(t *testing.T) {
	dst := []byte{0, 0, 0}
	src := []byte{0b01010101, 0b00110011, 0b00001111}
	sz := SlideCopy(dst, src, 0b11100000, 0b11111000, 17)
	assert.Equal(t, uint64(2), sz)
	assert.Equal(t, byte(0b10011010), dst[0])
	assert.Equal(t, byte(0b01111001), dst[1])
	assert.Equal(t, byte(0b00000000), dst[2])
}

func BenchmarkSlideCopy(b *testing.B) {
	dst := make([]byte, 1024)
	src := make([]byte, 1025)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SlideCopy(dst, src, 0b11100000, 0b11111000, 127)
	}
}

func BenchmarkSlide(b *testing.B) {
	dst := make([]byte, 1024)
	src := make([]byte, 1025)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Slide(dst, src, 0b11100000, 0b00011111, 5, 127)
	}
}
