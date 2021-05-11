package gobits

type Segment struct {
	ReadWriter
	WriterTo
	ReaderFrom
	b *Slice
}

func NewSegmentWithSize(bufferSize int) *Segment {
	return NewSegment(
		NewSlice(
			make([]byte, bufferSize), 0, 0))
}

func NewSegment(b *Slice) *Segment {
	return &Segment{
		b: b,
	}
}

func (s *Segment) Write(b *Slice) (int, error) {
	size := b.Len()
	if size > s.RemainLen() {
		size = s.RemainLen()
	}
	if err := Copy(s.b.buffer, s.b.bot, b.buffer, b.top, size); err != nil {
		return 0, err
	}
	s.b.bot += size
	return size, nil
}

func (s *Segment) ReadFrom(r Reader) error {
	size, err := r.Read(s.b)
	s.b.bot += size
	return err
}

func (s *Segment) Read(b *Slice) (int, error) {
	size := b.Len()
	if size > s.Len() {
		size = s.Len()
	}
	if err := Copy(b.buffer, b.top, s.b.buffer, s.b.top, size); err != nil {
		return 0, err
	}
	s.b.top += size
	return size, nil
}

func (s *Segment) WriteTo(w Writer) error {
	size, err := w.Write(s.b)
	s.b.top += size
	return err
}

func (s *Segment) Len() int {
	return s.b.bot - s.b.top
}

func (s *Segment) RemainLen() int {
	return (len(s.b.buffer) << 3) - s.b.bot
}

func (s *Segment) Reset() {
	s.b.top = 0
	s.b.bot = 0
}
