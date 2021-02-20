// Package grapheme provides operations for individual user-perceived
// characters and implements UAX #29 by Unicode for grapheme boundary finding.
//
// It can be used to show the number of "characters on screen".
package grapheme

// String is a series of user-perceived characters. It is immutable; operations
// on the String produce a new String.
//
// The zero value is an empty String.
type String interface {
	// CharAt returns the runes that make up the user-perceived character
	// (grapheme cluster) of the given index. Modifying the returned slice will
	// not modify the String.
	CharAt(idx int) []rune

	// Len returns the number of grapheme clusters (user-perceivable characters)
	// that are in the String.
	//
	// This function may trigger UAX29 analysis on the String if it hasn't yet
	// occured.
	Len() int

	// Add adds to strings together and returns the result
	Add(s String) String

	// String gets the contents as the built-in string type.
	String() string

	// Runes returns the string's raw Runes. Modifying the returned slice has no
	// effect on the String.
	Runes() []rune
}

type runeString struct {
	r  []rune
	gc []int
}

func (runes *runeString) Runes() []rune {
	r := make([]rune, len(runes.r))
	for i := range runes.r {
		r[i] = runes.r[i]
	}
	return r
}

func (runes *runeString) String() string {
	return string(runes.r)
}

func (runes *runeString) CharAt(idx int) []rune {
	gc := runes.gc
	if gc == nil {
		gc = Split(runes.r)
		runes.gc = gc
	}

	var startIdx = 0
	if idx > 0 {
		startIdx = gc[idx-1]
	}
	length := gc[idx] - startIdx
	cluster := make([]rune, length)
	for i := 0; i < length; i++ {
		cluster[i] = runes.r[startIdx+i]
	}
	return cluster
}

// Add performs the Add operation of a String.
func (runes *runeString) Add(s2 String) String {
	r2 := runes.clone()
	r2.gc = nil
	r2.r = append(r2.r, s2.Runes()...)
	return r2
}

// Len returns the number of grapheme clusters (user-perceivable characters) are
// in the String.
//
// This function will trigger UAX29 analysis on the String if it hasn't yet
// occured.
func (runes *runeString) Len() int {
	gc := runes.gc
	if gc == nil {
		if len(runes.r) == 0 {
			return 0
		}
		gc = Split(runes.r)
		runes.gc = gc
	}
	return len(gc)
}

// makes an exact duplicate by copying underlying slices.
//
// modifications to any members of the returned string are guaranteed not to
// modify the original. calling this is not needed unless a modification is
// about to occur, even though passing String by value does pass pointers (via
// slice-type members)
func (runes *runeString) clone() *runeString {
	gc := runes.gc
	clone := runeString{
		r: make([]rune, len(runes.r)),
	}
	for i := range runes.r {
		clone.r[i] = runes.r[i]
	}
	if gc != nil {
		clone.gc = make([]int, len(gc))
		for i := range gc {
			clone.gc[i] = gc[i]
		}
	}
	return &clone
}

// New takes the given string and converts it into a graphemes.String object for
// use with grapheme-aware functions. UAX-29 analysis is performed on a lazy
// basis; the contents of s are not scanned for grapheme clusters until an
// operation requires it.
func New(s string) String {
	return &runeString{r: []rune(s)}
}

// Split splits the given runes into a series of grapheme clusters. The
// returned slice contains slices of indexes that give the exclusive ending
// index of runes that make up each grapheme at that position; e.g. the returned
// slice for "test" would be []int{1, 2, 3, 4} but one for two rune glyphs such
// as (и́) would be []int{2} (just the one) despite it containing two runes.
//
// The inclusive start index of each cluster is the last index. For the first
// item, it is 0
func Split(r []rune) []int {
	done := make([]int, 0)
	for i := range r {
		if shouldBreakAfter(r[i], r, i) {
			done = append(done, i+1)
		}
	}

	return done
}

func shouldBreakAfter(r rune, chars []rune, i int) bool {
	// GB1 - Break at the start of the text, implemented when starting

	// GB2 - Break at the end of the text
	if i+1 >= len(chars) {
		return true
	}

	// GB2 guarentees that i+1 is safe to access
	nextR := chars[i+1]

	// GB3 - Do not break between a CR and LF
	if isCbCR(r) && isCbLF(nextR) {
		return false
	}

	// GB4 - Break after controls
	if isCbControl(r) || isCbCR(r) || isCbLF(r) {
		return true
	}

	// GB5 - Break before controls
	if isCbControl(nextR) || isCbCR(nextR) || isCbLF(nextR) {
		return true
	}

	// GB6 - Do not break Hangul syllable sequences (1)
	if isCbL(r) && (isCbL(nextR) || isCbV(nextR) || isCbLV(nextR) || isCbLVT(nextR)) {
		return false
	}

	// GB7 - Do not break Hangul syllable sequences (2)
	if (isCbLV(r) || isCbV(r)) && (isCbV(nextR) || isCbT(nextR)) {
		return false
	}

	// GB8 - Do not break Hangul syllable sequences (3)
	if (isCbLVT(r) || isCbT(r)) && isCbT(nextR) {
		return false
	}

	// GB9 - Do not break before extending characters or ZWJ
	if isCbExtend(nextR) || isCbZWJ(nextR) {
		return false
	}

	// GB9a - (Extended grapheme clusters only) Do not break before spacing marks
	if isCbSpacingMark(nextR) {
		return false
	}

	// GB9b - (Extended grapheme clusters only) Do not break after Prepend characters
	if isCbPrepend(r) {
		return false
	}

	// GB11 - Do not break within emoji modifier sequences or emoji ZWJ sequnces
	if i-2 >= 0 {
		if isCbZWJ(chars[i-1]) {
			for j := i - 2; j >= 0; j-- {
				if isExtPicto(chars[j]) {
					return false
				}
				if !isCbExtend(chars[j]) {
					break
				}
			}
		}
	}

	// GB12 & GB13 - Do not break wihtin emoji flag sequences
	if isCbRegionalIndicator(r) {
		for j := i; j >= 0; j -= 2 {
			// the one we are on chars[j] is always RI
			if j-1 < 0 {
				// odd number of RIs
				return false
			}
			if j-2 < 0 {
				if !isCbRegionalIndicator(chars[j]) {
					// odd number of RIs
					return false
				}
				// even number of RIs
				continue
			}

			// N, _
			// N, N,

			if !isCbRegionalIndicator(chars[j-1]) {
				// odd number of RIs
				return false
			}

			if !isCbRegionalIndicator(chars[j-2]) {
				// even number of RIs
				break
			}
		}
	}

	// GB999 - Otherwise, break anywhere
	return true
}
