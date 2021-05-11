/*
package serializer

import "github.com/KnowledgeSense/gobitsegment"

type WrappedUint32Conversion struct {
	num        *uint32
	conversion Uint32Conversion
}

func NewWrappedUint32Conversion(
	num *uint32,
	conversion Uint32Conversion,
) *WrappedUint32Conversion {
	return &WrappedUint32Conversion{
		num:        num,
		conversion: conversion,
	}
}

func (c *WrappedUint32Conversion) Read(seg gobitsegment.Segment) error {
	return c.conversion.Read(seg, c.num)
}

func (c *WrappedUint32Conversion) Write(seg gobitsegment.Segment) error {
	return c.conversion.Write(seg, c.num)
}

type Uint32Conversion interface {
	Read(gobitsegment.Segment, *uint32) error
	Write(gobitsegment.Segment, *uint32) error
}

// BitUint32Conversion is bit parse Conversion
type BitUint32Conversion struct {
	buffer []byte
	size   int
}

func NewBitUint32Conversion(n int) *BitUint32Conversion {
	if n < 1 {
		n = 1
	} else if n > 32 {
		n = 32
	}
	return &BitUint32Conversion{
		buffer: make([]byte, (n+7)>>3),
		size:   n,
	}
}

func (c *BitUint32Conversion) Read(seg gobitsegment.Segment, output *uint32) error {
	c.buffer[len(c.buffer)-1] = 0 // 最後のバイトは余計な値が入らないように0にしておく
	if err := seg.ReadWithSize(c.buffer, 0, c.size); err != nil {
		return err
	}
	var ret uint32 = 0
	for i, b := range c.buffer {
		ret |= uint32(b) << (i << 3)
	}
	*output = ret
	return nil
}

func (c *BitUint32Conversion) Write(seg gobitsegment.Segment, input *uint32) error {
	for i := range c.buffer {
		c.buffer[i] = byte((*input >> (i << 3)) & 255)
	}
	return seg.WriteWithSize(c.buffer, 0, c.size)
}
*/