package gobits

/*

Segmentのきのうを踏襲しながら
Readの際に内部の読み込み位置をずらさない機能を提供する

つまりtopの値が移動しないことを意味する

その他の機能はSegmentを参照

*/

type FixedSegment struct {
	*Segment
}

func NewFixedSegment(
	segment *Segment,
) *FixedSegment {
	return &FixedSegment{
		Segment: segment,
	}
}

// 内容をSliceにうつすがTopの位置は移動させない
func (s *FixedSegment) Read(b *Slice) (int, error) {
	size := b.RemainLen()
	if size > s.Len() {
		size = s.Len()
	}
	if err := Copy(b.buffer, b.top, s.b.buffer, s.b.top, size); err != nil {
		return 0, err
	}
	return size, nil
}
