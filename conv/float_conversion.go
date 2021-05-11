/*
package serializer

import (
	"math"

	"github.com/KnowledgeSense/gobitsegment"
)

type WrappedFloat32Conversion struct {
	num        *float32
	conversion Float32Conversion
}

func NewWrappedFloat32Conversion(
	num *float32,
	conversion Float32Conversion,
) *WrappedFloat32Conversion {
	return &WrappedFloat32Conversion{
		num:        num,
		conversion: conversion,
	}
}

func (c *WrappedFloat32Conversion) Read(seg gobitsegment.Segment) error {
	return c.conversion.Read(seg, c.num)
}

func (c *WrappedFloat32Conversion) Write(seg gobitsegment.Segment) error {
	return c.conversion.Write(seg, c.num)
}

type Float32Conversion interface {
	Read(gobitsegment.Segment, *float32) error
	Write(gobitsegment.Segment, *float32) error
}

type MultiFloatConversion struct {
	multi      float64 // Floatを何倍して整数化するかの定義
	conversion *BitInt32Conversion
}

func NewMultiFloatConversion(
	bits int,
	multi int,
) *MultiFloatConversion {
	return &MultiFloatConversion{
		multi:      float64(multi),
		conversion: NewBitInt32Conversion(bits),
	}
}

func (c *MultiFloatConversion) Read(seg gobitsegment.Segment, output *float32) error {
	var num int32
	if err := c.conversion.Read(seg, &num); err != nil {
		return err
	}
	*output = float32(float64(num) / c.multi)
	return nil
}

func (c *MultiFloatConversion) Write(seg gobitsegment.Segment, input *float32) error {
	var num int32 = int32(math.Round(float64(*input) * c.multi))
	return c.conversion.Write(seg, &num)
}

type ClampFloatConversion struct {
	min        float32
	max        float32
	unit       float64
	conversion *BitUint32Conversion
}

func NewClampFloatConversion(
	bits int,
	min float32,
	max float32,
) *ClampFloatConversion {
	var maxsize uint32 = math.MaxUint32
	if bits < 32 {
		maxsize = (uint32(1) << bits) - 1
	}
	return &ClampFloatConversion{
		min:        min,
		max:        max,
		unit:       float64(max-min) / float64(maxsize),
		conversion: NewBitUint32Conversion(bits),
	}
}

func (c *ClampFloatConversion) Read(seg gobitsegment.Segment, output *float32) error {
	var num uint32
	if err := c.conversion.Read(seg, &num); err != nil {
		return err
	}
	*output = float32(c.unit*float64(num)) + c.min
	return nil
}

func (c *ClampFloatConversion) Write(seg gobitsegment.Segment, input *float32) error {
	inum := *input
	if inum < c.min {
		inum = c.min
	} else if inum > c.max {
		inum = c.max
	}
	num := uint32(math.Round(float64(inum-c.min) / c.unit))
	return c.conversion.Write(seg, &num)
}
*/