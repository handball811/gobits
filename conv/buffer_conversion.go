/*
package serializer

import "github.com/KnowledgeSense/gobitsegment"

type WrappedBufferConversion struct {
	// 読み書き用のBuffer
	buffer []byte
	// 受け取ったバッファのサイズや送るバッファのサイズを指定するのに使用します
	bufferSize *int
	conversion BufferConversion
}

func NewWrappedBufferConversion(
	buffer []byte,
	bufferSize *int,
	conversion BufferConversion,
) *WrappedBufferConversion {
	return &WrappedBufferConversion{
		buffer:     buffer,
		bufferSize: bufferSize,
		conversion: conversion,
	}
}

func (c *WrappedBufferConversion) Read(seg gobitsegment.Segment) error {
	n, err := c.conversion.Read(seg, c.buffer)
	if err != nil {
		return err
	}
	*c.bufferSize = n
	return nil
}

func (c *WrappedBufferConversion) Write(seg gobitsegment.Segment) error {
	if err := c.conversion.Write(seg, c.buffer[:*c.bufferSize]); err != nil {
		return err
	}
	return nil
}

type BufferConversion interface {
	Read(gobitsegment.Segment, []byte) (int, error)
	Write(gobitsegment.Segment, []byte) error
}

type SizedBufferConversion struct {
	maxsize    int
	conversion *BitUint32Conversion
	segment    gobitsegment.Segment
}

func NewSizedBufferConversion(n int) *SizedBufferConversion {
	if n < 1 {
		n = 1
	}
	if n > 16 {
		n = 16
	}
	return &SizedBufferConversion{
		maxsize:    (1 << n) - 1,
		conversion: NewBitUint32Conversion(n),
		segment:    gobitsegment.NewSimpleSegment(((1 << n) + 6)),
	}
}

func (c *SizedBufferConversion) Read(seg gobitsegment.Segment, output []byte) (int, error) {
	var size uint32
	if err := c.conversion.Read(seg, &size); err != nil {
		return 0, err
	}
	if err := seg.ReadWithSize(output, 0, int(size)<<3); err != nil {
		return 0, err
	}
	return int(size), nil
}

func (c *SizedBufferConversion) Write(seg gobitsegment.Segment, input []byte) error {
	if len(input) > c.maxsize {
		return gobitsegment.ErrOutOfBound
	}
	var size uint32 = uint32(len(input))
	// 大きさだけ入力できてしまうと困るので一度まとめてから、Writeを試す
	c.segment.Reset()
	if err := c.conversion.Write(c.segment, &size); err != nil {
		return err
	}
	if err := c.segment.WriteWithSize(input, 0, int(size)<<3); err != nil {
		return err
	}
	return seg.WriteSegment(c.segment)
}
*/