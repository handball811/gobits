/*
package serializer

import (
	"github.com/KnowledgeSense/gobitsegment"
	"github.com/KnowledgeSense/goposrot"
)

type WrappedVector3Conversion struct {
	vec        *goposrot.Vector3
	conversion Vector3Conversion
}

func NewWrappedVector3Conversion(
	vec *goposrot.Vector3,
	conversion Vector3Conversion,
) *WrappedVector3Conversion {
	return &WrappedVector3Conversion{
		vec:        vec,
		conversion: conversion,
	}
}

func (c *WrappedVector3Conversion) Read(seg gobitsegment.Segment) error {
	return c.conversion.Read(seg, c.vec)
}

func (c *WrappedVector3Conversion) Write(seg gobitsegment.Segment) error {
	return c.conversion.Write(seg, c.vec)
}

type Vector3Conversion interface {
	Read(gobitsegment.Segment, *goposrot.Vector3) error
	Write(gobitsegment.Segment, *goposrot.Vector3) error
}

type ClampVector3Conversion struct {
	segment     gobitsegment.Segment
	xConversion *ClampFloatConversion
	yConversion *ClampFloatConversion
	zConversion *ClampFloatConversion
}

func NewClampVector3Conversion(
	nx int, ax, bx float32,
	ny int, ay, by float32,
	nz int, az, bz float32,
) *ClampVector3Conversion {
	return &ClampVector3Conversion{
		segment:     gobitsegment.NewSimpleSegment((nx + ny + nz + 7) / 8),
		xConversion: NewClampFloatConversion(nx, ax, bx),
		yConversion: NewClampFloatConversion(ny, ay, by),
		zConversion: NewClampFloatConversion(nz, az, bz),
	}
}

func (c *ClampVector3Conversion) Read(
	seg gobitsegment.Segment,
	output *goposrot.Vector3,
) error {
	var numx, numy, numz float32
	if err := c.xConversion.Read(seg, &numx); err != nil {
		return err
	}
	if err := c.yConversion.Read(seg, &numy); err != nil {
		return err
	}
	if err := c.zConversion.Read(seg, &numz); err != nil {
		return err
	}
	output.X = numx
	output.Y = numy
	output.Z = numz
	return nil
}

func (c *ClampVector3Conversion) Write(
	seg gobitsegment.Segment,
	input *goposrot.Vector3,
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
	return seg.WriteSegment(c.segment)
}
*/
