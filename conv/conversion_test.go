/*
package serializer_test

import (
	"math"
	"testing"

	"github.com/KnowledgeSense/goposrot"

	"github.com/KnowledgeSense/GenericSerializer/go/serializer"
	"github.com/KnowledgeSense/gobitsegment"
	"github.com/stretchr/testify/assert"
)

func TestOneBoolConversion(t *testing.T) {
	// setup
	var err error
	c := serializer.NewOneBoolConversion()
	segment := gobitsegment.NewSimpleSegment(1)
	input := true
	var output bool
	//when
	err = c.Write(segment, &input)

	//then
	assert.Nil(t, err)
	assert.Equal(t, 1, segment.Len())

	//when
	err = c.Read(segment, &output)

	//then
	assert.Nil(t, err)
	assert.Equal(t, 0, segment.Len())
	assert.Equal(t, true, output)
}

func TestSizeBufferConversion(t *testing.T) {
	// setup
	var size int
	var err error
	c := serializer.NewSizedBufferConversion(2)
	segment := gobitsegment.NewSimpleSegment(16)
	shortSegment := gobitsegment.NewSimpleSegment(1)
	input := []byte{1, 2, 3, 4}
	output := make([]byte, 4)
	//when
	err = c.Write(segment, input)

	//then
	assert.NotNil(t, err)
	assert.Equal(t, 0, segment.Len())

	//when
	c.Write(shortSegment, input)
	err = c.Write(shortSegment, input[1:])

	//then
	assert.NotNil(t, err)

	//when
	err = c.Write(segment, input[1:])

	//then
	assert.Nil(t, err)
	assert.Equal(t, 26, segment.Len())

	size, err = c.Read(segment, output)

	//then
	assert.Nil(t, err)
	assert.Equal(t, 0, segment.Len())
	assert.Equal(t, 3, size)
	assert.Equal(t, byte(2), output[0])
	assert.Equal(t, byte(3), output[1])
	assert.Equal(t, byte(4), output[2])
	assert.Equal(t, byte(0), output[3])
}

func TestBitUint32Conversion(t *testing.T) {
	var err error
	c := serializer.NewBitUint32Conversion(5)
	segment := gobitsegment.NewSimpleSegment(10)
	var input uint32 = 29
	var output uint32
	//when
	err = c.Write(segment, &input)

	//then
	assert.Nil(t, err)
	assert.Equal(t, 5, segment.Len())

	//when
	c.Write(segment, &input)
	c.Write(segment, &input)
	err = c.Read(segment, &output)

	//then
	assert.Nil(t, err)
	assert.Equal(t, 10, segment.Len())
	assert.Equal(t, input, output)

	//when
	err = c.Read(segment, &output)

	//then
	assert.Nil(t, err)
	assert.Equal(t, 5, segment.Len())
	assert.Equal(t, input, output)

	//when
	err = c.Read(segment, &output)

	//then
	assert.Nil(t, err)
	assert.Equal(t, 0, segment.Len())
	assert.Equal(t, input, output)
}

func TestBitInt32Conversion(t *testing.T) {
	var err error
	c := serializer.NewBitInt32Conversion(32)
	segment := gobitsegment.NewSimpleSegment(10)
	var input int32 = math.MaxInt32
	var output int32

	//when
	err = c.Write(segment, &input)

	//then
	assert.Nil(t, err)
	assert.Equal(t, 32, segment.Len())

	//when
	input = math.MinInt32
	c.Write(segment, &input)
	err = c.Read(segment, &output)

	//then
	assert.Nil(t, err)
	assert.Equal(t, 32, segment.Len())
	assert.Equal(t, int32(math.MaxInt32), output)

	//when
	err = c.Read(segment, &output)

	//then
	assert.Nil(t, err)
	assert.Equal(t, 0, segment.Len())
	assert.Equal(t, int32(math.MinInt32), output)
}

func TestMultiFloatConversion(t *testing.T) {
	//setup
	var err error
	c := serializer.NewMultiFloatConversion(16, 128)
	segment := gobitsegment.NewSimpleSegment(10)
	var input float32 = -1.1345078
	var output float32

	//when
	err = c.Write(segment, &input)

	//then
	assert.Nil(t, err)
	assert.Equal(t, 16, segment.Len())

	//when
	err = c.Read(segment, &output)

	//then
	assert.Nil(t, err)
	assert.Equal(t, 0, segment.Len())
	assert.InDelta(t, input, output, 1./256.)
}

func TestClampFloatConversion(t *testing.T) {
	//setup
	var err error
	c := serializer.NewClampFloatConversion(8, -1, 1)
	segment := gobitsegment.NewSimpleSegment(10)
	var input float32 = -0.456789
	var output float32

	//when
	err = c.Write(segment, &input)

	//then
	assert.Nil(t, err)
	assert.Equal(t, 8, segment.Len())

	//when
	err = c.Read(segment, &output)

	//then
	assert.Nil(t, err)
	assert.Equal(t, 0, segment.Len())
	assert.InDelta(t, input, output, 1./256.)

	//when
	input = 1.23
	err = c.Write(segment, &input)
	err = c.Read(segment, &output)

	//then
	assert.Nil(t, err)
	assert.Equal(t, 0, segment.Len())
	assert.InDelta(t, float32(1), output, 1./256.)
}

func TestClampVector3Conversion(t *testing.T) {
	//setup
	var err error
	c := serializer.NewClampVector3Conversion(
		8, -1., 1.,
		4, -1., 1.,
		8, -1., 1.,
	)
	segment := gobitsegment.NewSimpleSegment(10)
	input := goposrot.Vector3{
		X: 0.5,
		Y: 0.2,
		Z: -0.7,
	}
	var output goposrot.Vector3

	//when
	err = c.Write(segment, &input)
	//then
	assert.Nil(t, err)
	assert.Equal(t, 20, segment.Len())

	//when
	err = c.Read(segment, &output)
	//then
	assert.Nil(t, err)
	assert.Equal(t, 0, segment.Len())
	assert.InDelta(t, 0.5, output.X, 1./256.)
	assert.InDelta(t, 0.2, output.Y, 1./256.)
	assert.InDelta(t, -0.7, output.Z, 1./256.)
}

func TestClampQuaternionConversion(t *testing.T) {
	//setup
	var err error
	c := serializer.NewClampQuaternionConversion(16, 16, 16, 16)
	segment := gobitsegment.NewSimpleSegment(10)
	input := goposrot.Quaternion{
		X: 0.5123,
		Y: 0.15477,
		Z: -0.78931,
		W: 0.3000972,
	}
	var output goposrot.Quaternion

	//when
	err = c.Write(segment, &input)
	//then
	assert.Nil(t, err)
	assert.Equal(t, 64, segment.Len())

	//when
	err = c.Read(segment, &output)
	//then
	assert.Nil(t, err)
	assert.Equal(t, 0, segment.Len())
	assert.InDelta(t, input.X, output.X, 1./256./256.)
	assert.InDelta(t, input.Y, output.Y, 1./256./256.)
	assert.InDelta(t, input.Z, output.Z, 1./256./256.)
	assert.InDelta(t, input.W, output.W, 1./256./256.)
}

func TestCompressedQuaternionconversion(t *testing.T) {
	//setup
	var err error
	c := serializer.NewCompressedQuaternionConversion(9)
	segment := gobitsegment.NewSimpleSegment(10)
	input := goposrot.Quaternion{
		X: 0.5123,
		Y: 0.15477,
		Z: -0.78931,
		W: 0.3000972,
	}
	var output goposrot.Quaternion

	//when
	err = c.Write(segment, &input)
	//then
	assert.Nil(t, err)
	assert.Equal(t, 29, segment.Len())

	//when
	err = c.Read(segment, &output)
	//then
	assert.Nil(t, err)
	assert.Equal(t, 0, segment.Len())
	assert.InDelta(t, -input.X, output.X, 1./256.)
	assert.InDelta(t, -input.Y, output.Y, 1./256.)
	assert.InDelta(t, -input.Z, output.Z, 1./256.)
	assert.InDelta(t, -input.W, output.W, 1./256.)
}

func TestSampleStructure(t *testing.T) {
	var err error
	s := serializer.NewSampleStructure()
	tr := serializer.NewSampleStructure()
	segment := gobitsegment.NewSimpleSegment(512)
	s.Age = 56
	s.Pos.X = 512.31
	s.Pos.Y = 0.3
	s.Pos.Z = -23.3
	s.Rot.X = 0.513
	s.Rot.Y = 0.15477
	s.Rot.Z = -0.78931
	s.Rot.W = 0.3000972
	s.Power = 10.23
	s.Buffer[2] = 203
	s.BufLen = 38

	err = s.Write(segment)
	assert.Nil(t, err)

	err = tr.Read(segment)
	assert.Nil(t, err)

	assert.Equal(t, s.Age, tr.Age)
	assert.InDelta(t, s.Pos.X, tr.Pos.X, 1./256.)
	assert.InDelta(t, s.Pos.Y, tr.Pos.Y, 1./256.)
	assert.InDelta(t, s.Pos.Z, tr.Pos.Z, 1./256.)
	assert.InDelta(t, -s.Rot.X, tr.Rot.X, 1./128.)
	assert.InDelta(t, -s.Rot.Y, tr.Rot.Y, 1./128.)
	assert.InDelta(t, -s.Rot.Z, tr.Rot.Z, 1./128.)
	assert.InDelta(t, -s.Rot.W, tr.Rot.W, 1./128.)
	assert.InDelta(t, s.Power, tr.Power, 1./2.)
	assert.Equal(t, s.BufLen, tr.BufLen)
	assert.Equal(t, s.Buffer[0], tr.Buffer[0])
	assert.Equal(t, s.Buffer[1], tr.Buffer[1])
	assert.Equal(t, s.Buffer[2], tr.Buffer[2])
	assert.Equal(t, s.Buffer[3], tr.Buffer[3])
}

// Benchmark結果の目安
// Read:  349ns
// Write: 452ns
// Writeは書き込む内容をまとめてから書き込むようにしているため余計なコピーが発生している
func BenchmarkSampleStructureWrite(b *testing.B) {
	s := serializer.NewSampleStructure()
	segment := gobitsegment.NewSimpleSegment(512)
	s.Age = 56
	s.Pos.X = 512.31
	s.Pos.Y = 0.3
	s.Pos.Z = -23.3
	s.Rot.X = 0.513
	s.Rot.Y = 0.15477
	s.Rot.Z = -0.78931
	s.Rot.W = 0.3000972
	s.Power = 10.23
	s.Buffer[2] = 203
	s.BufLen = 38
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		segment.Reset()
		s.Write(segment)
	}
}
func BenchmarkSampleStructureWriteAndRead(b *testing.B) {
	s := serializer.NewSampleStructure()
	t := serializer.NewSampleStructure()
	segment := gobitsegment.NewSimpleSegment(512)
	s.Age = 56
	s.Pos.X = 512.31
	s.Pos.Y = 0.3
	s.Pos.Z = -23.3
	s.Rot.X = 0.513
	s.Rot.Y = 0.15477
	s.Rot.Z = -0.78931
	s.Rot.W = 0.3000972
	s.Power = 10.23
	s.Buffer[2] = 203
	s.BufLen = 38
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		segment.Reset()
		s.Write(segment)
		t.Read(segment)
	}
}
*/