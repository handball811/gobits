/*
package serializer

import (
	"math"

	"github.com/KnowledgeSense/gobitsegment"
	"github.com/KnowledgeSense/goposrot"
)

var (
	minCompressedQuaternionParam = float32(-math.Sqrt(0.5))
	maxCompressedQuaternionParam = float32(math.Sqrt(0.5))
)

type WrappedQuaternionConversion struct {
	quat       *goposrot.Quaternion
	conversion QuaternionConversion
}

func NewWrappedQuaternionConversion(
	quat *goposrot.Quaternion,
	conversion QuaternionConversion,
) *WrappedQuaternionConversion {
	return &WrappedQuaternionConversion{
		quat:       quat,
		conversion: conversion,
	}
}

func (c *WrappedQuaternionConversion) Read(seg gobitsegment.Segment) error {
	return c.conversion.Read(seg, c.quat)
}

func (c *WrappedQuaternionConversion) Write(seg gobitsegment.Segment) error {
	return c.conversion.Write(seg, c.quat)
}

type QuaternionConversion interface {
	Read(gobitsegment.Segment, *goposrot.Quaternion) error
	Write(gobitsegment.Segment, *goposrot.Quaternion) error
}

type ClampQuaternionConversion struct {
	segment     gobitsegment.Segment
	xConversion *ClampFloatConversion
	yConversion *ClampFloatConversion
	zConversion *ClampFloatConversion
	wConversion *ClampFloatConversion
}

func NewClampQuaternionConversion(
	nx, ny, nz, nw int,
) *ClampQuaternionConversion {
	return &ClampQuaternionConversion{
		segment:     gobitsegment.NewSimpleSegment((nx + ny + nz + nw + 7) / 8),
		xConversion: NewClampFloatConversion(nx, -1., 1.),
		yConversion: NewClampFloatConversion(ny, -1., 1.),
		zConversion: NewClampFloatConversion(nz, -1., 1.),
		wConversion: NewClampFloatConversion(nw, -1., 1.),
	}
}

func (c *ClampQuaternionConversion) Read(
	seg gobitsegment.Segment,
	output *goposrot.Quaternion,
) error {
	var numx, numy, numz, numw float32
	if err := c.xConversion.Read(seg, &numx); err != nil {
		return err
	}
	if err := c.yConversion.Read(seg, &numy); err != nil {
		return err
	}
	if err := c.zConversion.Read(seg, &numz); err != nil {
		return err
	}
	if err := c.wConversion.Read(seg, &numw); err != nil {
		return err
	}
	output.X = numx
	output.Y = numy
	output.Z = numz
	output.W = numw
	return nil
}

func (c *ClampQuaternionConversion) Write(
	seg gobitsegment.Segment,
	input *goposrot.Quaternion,
) error {
	c.segment.Reset()
	if err := c.xConversion.Write(c.segment, &input.X); err != nil {
		return err
	}
	if err := c.yConversion.Write(c.segment, &input.Y); err != nil {
		return err
	}
	if err := c.zConversion.Write(c.segment, &input.Z); err != nil {
		return err
	}
	if err := c.wConversion.Write(c.segment, &input.W); err != nil {
		return err
	}
	return seg.WriteSegment(c.segment)
}

type CompressedQuaternionConversion struct {
	segment          gobitsegment.Segment
	targetConversion *BitUint32Conversion
	conversion       *ClampFloatConversion
}

func NewCompressedQuaternionConversion(
	bits int,
) *CompressedQuaternionConversion {
	return &CompressedQuaternionConversion{
		segment:          gobitsegment.NewSimpleSegment((bits*3 + 2 + 7) / 8),
		targetConversion: NewBitUint32Conversion(2),
		conversion:       NewClampFloatConversion(bits, minCompressedQuaternionParam, maxCompressedQuaternionParam),
	}
}

func (c *CompressedQuaternionConversion) Read(
	seg gobitsegment.Segment,
	output *goposrot.Quaternion,
) error {
	var mv float32 = 1
	var mt uint32
	var x, y, z, w float32
	if err := c.targetConversion.Read(seg, &mt); err != nil {
		return err
	}
	if mt != 0 {
		if err := c.conversion.Read(seg, &x); err != nil {
			return err
		}
		mv -= x * x
	}
	if mt != 1 {
		if err := c.conversion.Read(seg, &y); err != nil {
			return err
		}
		mv -= y * y
	}
	if mt != 2 {
		if err := c.conversion.Read(seg, &z); err != nil {
			return err
		}
		mv -= z * z
	}
	if mt != 3 {
		if err := c.conversion.Read(seg, &w); err != nil {
			return err
		}
		mv -= w * w
	}
	switch mt {
	case 0:
		x = float32(math.Sqrt(math.Max(float64(mv), 0.)))
	case 1:
		y = float32(math.Sqrt(math.Max(float64(mv), 0.)))
	case 2:
		z = float32(math.Sqrt(math.Max(float64(mv), 0.)))
	case 3:
		w = float32(math.Sqrt(math.Max(float64(mv), 0.)))
	}
	output.X = x
	output.Y = y
	output.Z = z
	output.W = w
	return nil
}

func (c *CompressedQuaternionConversion) Write(
	seg gobitsegment.Segment,
	input *goposrot.Quaternion,
) error {
	var mv float32 = input.X * input.X
	var mt uint32 = 0
	var isNeg bool = input.X < 0
	if input.Y*input.Y > mv {
		mt = 1
		mv = input.Y * input.Y
		isNeg = input.Y < 0
	}
	if input.Z*input.Z > mv {
		mt = 2
		mv = input.Z * input.Z
		isNeg = input.Z < 0
	}
	if input.W*input.W > mv {
		mt = 3
		mv = input.W * input.W
		isNeg = input.W < 0
	}
	c.segment.Reset()
	if err := c.targetConversion.Write(c.segment, &mt); err != nil {
		return err
	}
	x := input.X
	y := input.Y
	z := input.Z
	w := input.W
	if isNeg {
		x = -x
		y = -y
		z = -z
		w = -w
	}
	if mt != 0 {
		if err := c.conversion.Write(c.segment, &x); err != nil {
			return err
		}
	}
	if mt != 1 {
		if err := c.conversion.Write(c.segment, &y); err != nil {
			return err
		}
	}
	if mt != 2 {
		if err := c.conversion.Write(c.segment, &z); err != nil {
			return err
		}
	}
	if mt != 3 {
		if err := c.conversion.Write(c.segment, &w); err != nil {
			return err
		}
	}
	return seg.WriteSegment(c.segment)
}
*/