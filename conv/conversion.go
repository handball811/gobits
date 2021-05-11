/*
package serializer

import (
	"github.com/KnowledgeSense/gobitsegment"
)

const (
	maxSegmentSize = 1500
)

type Conversion interface {
	Read(gobitsegment.Segment) error
	Write(gobitsegment.Segment) error
}

type CompositeConversion struct {
	segment     gobitsegment.Segment
	conversions []Conversion
}

func NewCompositeConversion() *CompositeConversion {
	return &CompositeConversion{
		segment:     gobitsegment.NewSimpleSegment(maxSegmentSize),
		conversions: make([]Conversion, 0, 8),
	}
}

func (c *CompositeConversion) Add(conversion Conversion) {
	c.conversions = append(c.conversions, conversion)
}

func (c *CompositeConversion) Read(seg gobitsegment.Segment) error {
	for _, conv := range c.conversions {
		if err := conv.Read(seg); err != nil {
			return err
		}
	}
	return nil
}

func (c *CompositeConversion) Write(seg gobitsegment.Segment) error {
	c.segment.Reset()
	for _, conv := range c.conversions {
		if err := conv.Write(c.segment); err != nil {
			return err
		}
	}
	return seg.WriteSegment(c.segment)
}
*/