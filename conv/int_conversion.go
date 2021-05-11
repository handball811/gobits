/*
package serializer

import "github.com/KnowledgeSense/gobitsegment"

type WrappedInt32Conversion struct {
	num        *int32
	conversion Int32Conversion
}

func NewWrappedInt32Conversion(
	num *int32,
	conversion Int32Conversion,
) *WrappedInt32Conversion {
	return &WrappedInt32Conversion{
		num:        num,
		conversion: conversion,
	}
}

func (c *WrappedInt32Conversion) Read(seg gobitsegment.Segment) error {
	return c.conversion.Read(seg, c.num)
}

func (c *WrappedInt32Conversion) Write(seg gobitsegment.Segment) error {
	return c.conversion.Write(seg, c.num)
}

type Int32Conversion interface {
	Read(gobitsegment.Segment, *int32) error
	Write(gobitsegment.Segment, *int32) error
}

type BitInt32Conversion struct {
	conversion *BitUint32Conversion
	sub        uint32
}

func NewBitInt32Conversion(n int) *BitInt32Conversion {
	if n < 1 {
		n = 1
	} else if n > 32 {
		n = 32
	}
	return &BitInt32Conversion{
		conversion: NewBitUint32Conversion(n),
		sub:        uint32(1) << (n - 1),
	}
}

func (c *BitInt32Conversion) Read(seg gobitsegment.Segment, output *int32) error {
	var num uint32
	if err := c.conversion.Read(seg, &num); err != nil {
		return err
	}
	if num >= c.sub {
		*output = int32(num - c.sub)
	} else {
		*output = -int32(c.sub - num)
	}
	return nil
}

func (c *BitInt32Conversion) Write(seg gobitsegment.Segment, input *int32) error {
	var num uint32
	if *input >= 0 {
		num = uint32(*input) + c.sub
	} else {
		num = c.sub - uint32(-(*input))
	}
	return c.conversion.Write(seg, &num)
}
*/