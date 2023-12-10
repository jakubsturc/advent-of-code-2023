package day05

type interval struct{ start, length uint64 }

func (this interval) intersection(other interval) (bool, interval) {
	if this.start > other.start {
		return other.intersection(this)
	}
	if this.start+this.length <= other.start {
		return false, interval{}
	}
	return true, interval{
		max2(this.start, other.start),
		min2(this.start+this.length, other.start+other.length) - max2(this.start, other.start)}
}

func (this interval) subtract(other interval) []interval {
	found, overlap := this.intersection(other)
	if !found {
		return []interval{this}
	}

	if overlap.start == this.start {
		if overlap.length == this.length {
			return nil
		}
		return []interval{interval{this.start + overlap.length, this.length - overlap.length}}
	}

	if overlap.start+overlap.length == this.start+this.length {
		return []interval{interval{this.start, overlap.start - this.start}}
	}

	return []interval{
		interval{this.start, overlap.start - this.start},
		interval{overlap.start + overlap.length, this.start + this.length - overlap.start - overlap.length}}
}

func min2(a, b uint64) uint64 {
	if a < b {
		return a
	}
	return b
}

func max2(a, b uint64) uint64 {
	if a > b {
		return a
	}
	return b
}
