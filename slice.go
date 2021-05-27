package gobits

type Slice struct {
	buffer []byte
	top    int
	bot    int
	cap    int
}

func NewSlice(
	buffer []byte,
	top int,
	bot int,
	cap int,
) *Slice {
	return &Slice{
		buffer: buffer,
		top:    top,
		bot:    bot,
		cap:    cap,
	}
}

func NewSliceWithBuffer(
	buffer []byte,
) *Slice {
	return &Slice{
		buffer: buffer,
		top:    0,
		bot:    len(buffer) << 3,
		cap:    len(buffer) << 3,
	}
}

func (b *Slice) Len() int {
	return b.bot - b.top
}

func (b *Slice) RemainLen() int {
	return b.cap - b.bot
}

// must be
// top >= 0 && bot >= top
func (b *Slice) Sub(top, bot int) *Slice {
	if top < 0 || bot < top {
		return nil
	}
	return &Slice{
		buffer: b.buffer,
		top:    b.top + top,
		bot:    b.top + bot,
		cap:    b.cap,
	}
}

func (b *Slice) Move(top, bot int) bool {
	if top < 0 || bot < top {
		return false
	}
	b.bot = b.top + bot
	b.top += top
	return true
}

func (b *Slice) Reset() {
	b.top = 0
	b.bot = 0
}
