/*
package serializer

import (
	"github.com/KnowledgeSense/gobitsegment"
	"github.com/KnowledgeSense/goposrot"
)

// SampleStructure はテストと自動生成用に作成している
type SampleStructure struct {
	Age    uint32
	Pos    goposrot.Vector3
	Rot    goposrot.Quaternion
	Power  float32
	Buffer []byte
	BufLen int
	conv   Conversion
}

func NewSampleStructure() *SampleStructure {
	s := &SampleStructure{}
	comp := NewCompositeConversion()
	// age
	{
		// 0-128
		conv := NewBitUint32Conversion(7)
		comp.Add(NewWrappedUint32Conversion(
			&s.Age,
			conv,
		))
	}
	// pos
	{
		conv := NewClampVector3Conversion(
			18, -100, 2100,
			16, -50, 300,
			18, -100, 2100,
		)
		comp.Add(NewWrappedVector3Conversion(
			&s.Pos,
			conv,
		))
	}
	// rot
	{
		conv := NewCompressedQuaternionConversion(9)
		comp.Add(NewWrappedQuaternionConversion(
			&s.Rot,
			conv,
		))
	}
	// power
	{
		conv := NewClampFloatConversion(6, -30, 30)
		comp.Add(NewWrappedFloat32Conversion(
			&s.Power,
			conv,
		))
	}
	// buffer
	{
		s.Buffer = make([]byte, 256)
		conv := NewSizedBufferConversion(8)
		comp.Add(NewWrappedBufferConversion(
			s.Buffer,
			&s.BufLen,
			conv,
		))
	}
	s.conv = comp
	return s
}

func (s *SampleStructure) Read(segment gobitsegment.Segment) error {
	return s.conv.Read(segment)
}

func (s *SampleStructure) Write(segment gobitsegment.Segment) error {
	return s.conv.Write(segment)
}
*/