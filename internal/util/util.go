package util

// RangeToIndexes converts a python-style range specifier that could contain
// negative indexes to their actual literal indexes within the string. The range
// is considered half-open of the form [start, end).
//
// `size` is the length of the collection being indexed into.
//
// If after negative -> positive conversion, a start or end is still negative,
// it will be assumed to be 0. If a start or end is more than the size of the
// target collection being indexed, it will be assumed to be `size`.
//
// Returns the real indexes of the range given.
func RangeToIndexes(size, start, end int) (int, int) {
	if start < 0 {
		start += size
		if start < 0 {
			start = 0
		}
	}
	if end < 0 {
		end += size
		if end < 0 {
			end = 0
		}
	}
	if end > size {
		end = size
	}
	if start > size {
		start = size
	}
	
	return start, end
}
