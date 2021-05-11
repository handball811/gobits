package gobits

// Copy copies buffer from src to dst
//
// offset represents bit order
// size represents bit size
// TODO: Speed Up or use uint64 instead
func Copy(
	dst []byte, dstOffset int,
	src []byte, srcOffset int,
	size int,
) error {
	if ((len(dst) << 3) < dstOffset+size) || ((len(src) << 3) < srcOffset+size) {
		return ErrOutOfBound
	}
	offset := dstOffset & 7
	chsize := (8 - offset) & 7
	// 先頭がそろっていた場合はまとめてコピーし時間を短くする
	if offset == (srcOffset&7) && size >= 8+chsize {
		if chsize > 0 {
			dst[dstOffset>>3] = (dst[dstOffset>>3] & OutlineByteFilter[chsize<<3|offset]) | (src[srcOffset>>3] & InlineByteFilter[chsize<<3|offset])
			dstOffset += chsize
			srcOffset += chsize
			size -= chsize
		}
		d := dst[dstOffset>>3:]
		s := src[srcOffset>>3:]
		bufSize := size >> 3
		copy(d[:bufSize], s[:bufSize])
		copyLen := bufSize << 3
		size -= copyLen
		dstOffset += copyLen
		srcOffset += copyLen
	}
	// この部分でバイト長分のコピーが発生している
	for size > 0 {
		dlo := dstOffset & 7
		slo := srcOffset & 7
		dli := dstOffset >> 3
		sli := srcOffset >> 3
		copyLen := size
		if copyLen > 8-dlo {
			copyLen = 8 - dlo
		}
		if copyLen > 8-slo {
			copyLen = 8 - slo
		}
		d := dst[dli]
		s := src[sli]
		if dlo != 0 || size < 8 {
			if dlo == slo {
				dst[dli] = (d & OutlineByteFilter[copyLen<<3|dlo]) | (s & InlineByteFilter[copyLen<<3|slo])
			} else if dlo > slo {
				dst[dli] = (d & OutlineByteFilter[copyLen<<3|dlo]) | ((s & InlineByteFilter[copyLen<<3|slo]) << (dlo - slo))
			} else {
				dst[dli] = (d & OutlineByteFilter[copyLen<<3|dlo]) | ((s & InlineByteFilter[copyLen<<3|slo]) >> (slo - dlo))
			}
			size -= copyLen
			dstOffset += copyLen
			srcOffset += copyLen
		} else {
			lf := InlineByteFilter[copyLen<<3|slo]
			rf := InlineByteFilter[(8-copyLen)<<3]
			for size >= 8 {
				dst[dli] = ((src[sli] & lf) >> slo) | ((src[sli+1] & rf) << copyLen)
				size -= 8
				dstOffset += 8
				srcOffset += 8
				sli++
				dli++
			}
		}
	}
	return nil
}
