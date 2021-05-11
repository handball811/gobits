package gobits

//TODO: Autoscaling
//TODO: AsInterface
//TODO: RingBuffer

type Slice struct {
	buffer []byte
	top    int
	bot    int
}

func NewSlice(
	buffer []byte,
	top int,
	bot int,
) *Slice {
	return &Slice{
		buffer: buffer,
		top:    top,
		bot:    bot,
	}
}

func NewSliceWithBuffer(
	buffer []byte,
) *Slice {
	return &Slice{
		buffer: buffer,
		top:    0,
		bot:    len(buffer) << 3,
	}
}

func (b *Slice) Len() int {
	return b.bot - b.top
}

func (b *Slice) Cap() int {
	return len(b.buffer) << 3
}

func (b *Slice) Sub(top, bot int) *Slice {
	// TODO: Rescale
	return &Slice{
		buffer: b.buffer,
		top:    b.top + top,
		bot:    b.bot + bot,
	}
}
