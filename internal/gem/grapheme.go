// Package gem provides operations for individual user-perceived
// characters and implements UAX #29 by Unicode for grapheme boundary finding,
// and UAX #15 (Unicode 3.11) for calculating the NFC of a string for certain
// comparison operations.
//
// It can be used to show the number of characters as a user would perceive one
// character to be, implementing the rules specified by Unicode to be safe
// across multiple character ranges.
package gem

import (
	"fmt"
	"github.com/dekarrin/rosed/internal/util"
)

// String is a series of user-perceived characters. The contents are immutable;
// operations on the String produce a new String. However, some transient state
// does exist on the String, though an effort is made to limit access and use
// in a somewhat threadsafe way, no guarantees are provided for use in threads.
//
// If such guarantees are needed, raw strings can be used primarily with
// conversion to and from gem.String as needed.
//
// The zero value is an empty String.
//
// String.Equal can be used to test against raw strings.
type String struct {
	r  []rune
	gc []int
}

var (
	// Zero is a String of zero length.
	Zero String = String{r: []rune{}, gc: []int{}}
)

// Sub returns the substring given between the two indexes. The returned String
// will be a copy with its contents set to the characters at indexes in the
// range [start, end).
//
// If start or end is less than 0 it is assumed to be that many away from the
// actual end of the string; e.g. -1 would be Len()-1, -2 would be Len()-2, etc.
// If end or start are greater than Len, they are assumed to be Len. If start or
// end are negative and point to an index less than 0 after calculating, it is
// assumed that they are pointing to 0.
func (str String) Sub(start, end int) String {
	start, end = util.RangeToIndexes(str.Len(), start, end)
	
	if start == end {
		return Zero
	}
	
	if str.gc == nil {
		// we need a split operation on the graphemes
		str.gc = Split(str.r)
	}
	
	copy := str.clone()
	
	var runesStart int
	if start > 0 {
		runesStart = copy.gc[start-1]
	}
	runesEnd := copy.gc[end-1]
	
	copy.r = copy.r[runesStart:runesEnd]
	
	// ANY further ops require resplitting
	// TODO: shouldnt we be able to calc it? I mean we know what was removed from the
	// rune slice
	copy.gc = nil
	return copy
}

// IsEmpty return whether the String is the empty string "".
func (str String) IsEmpty() bool {
	return len(str.r) == 0
}

// Less returns whether one String is lexigraphically less than another.
func (str String) Less(s String) bool {
	sR := s.Runes()
	minLen := len(sR)
	if minLen > len(str.r) {
		minLen = len(str.r)
	}
	
	for i := 0; i < minLen; i++ {
		if str.r[i] < sR[i] {
			return true
		}
		
		if str.r[i] > sR[i] {
			return false
		}
	}

	// if we get here, they are exactly the same up to minLen
	if minLen == len(sR) && minLen == len(str.r) {
		// exactly the same, not less
		return false
	}

	// if it is shorter, then it is less
	return minLen == len(str.r)
}

// Slice turns the from slice into a slice of String objects.
func Slice(from []string) []String {
	str := make([]String, len(from))
	for i := range from {
		str[i] = New(from[i])
	}
	return str
}

// Strings turns the from slice into a slice of plain string objects.
func Strings(from []String) []string {
	str := make([]string, len(from))
	for i := range from {
		str[i] = from[i].String()
	}
	return str
}

// Equal returns whether one String is equal to another object. If the object is
// another String struct, their resulting strings are compared. If the object is
// a raw string object, it is compared to the output of calling String() on the
// gem.String. Otherwise, false is returned.
func (str String) Equal(other interface{}) bool {
	if other == nil {
		return false
	}
	otherStr, otherIsRawStr := other.(string)
	if otherIsRawStr {
		return str.String() == otherStr
	}

	s2, otherIsString := other.(String)
	if !otherIsString {
		return false
	}

	if len(s2.r) != len(str.r) {
		return false
	}
	for i := range str.r {
		if s2.r[i] != str.r[i] {
			return false
		}
	}
	return true
}

// Runes returns the string's raw Runes. Modifying the returned slice has no
// effect on the String.
func (str String) Runes() []rune {
	r := make([]rune, len(str.r))
	for i := range str.r {
		r[i] = str.r[i]
	}
	return r
}

// GraphemeIndexes returns a slice of rune start and end indexes within the
// string of the rune(s) that make up each grapheme. Note that these cannot
// be used by callers for other calls into the String (e.g. Sub() takes
// grapheme indexes, not rune indexes) and this function is purely to query
// a gem String.
//
// The end indexes will be exclusive; e.g. a gem.String with contents "test"
// would produce [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}}.
func (str String) GraphemeIndexes() [][]int {
	gc := str.gc
	if gc == nil {
		gc = Split(str.r)
		str.gc = gc
	}
	
	indexes := make([][]int, len(gc))
	prevEnd := 0
	for i := range gc {
		grapheme := make([]int, 2)
		grapheme[0] = prevEnd
		grapheme[1] = gc[i]
		indexes[i] = grapheme
		
		prevEnd = gc[i]
	}
	
	return indexes
}

// SetCharAt sets the character at the given index to the given value and
// returns the resulting String. The original String is not modified.
func (str String) SetCharAt(idx int, r []rune) String {
	if len(r) == 0 {
		panic("SetCharAt received empty or nil replacement runes slice r")
	}
	
	copy := str.clone()

	if copy.gc == nil {
		copy.gc = Split(copy.r)
	}

	var startIdx int
	if idx > 0 {
		startIdx = copy.gc[idx-1]
	}
	curEndIdx := copy.gc[idx]
	copy.r = append(copy.r[:startIdx], append(r, copy.r[curEndIdx:]...)...)
	copy.gc = nil
	return copy
}

// Format formats the String for printing.
func (str *String) Format(f fmt.State, verb rune) {
	if verb == 'q' {
		if str == nil {
			f.Write([]byte("<nil>"))
		}
		f.Write([]byte(fmt.Sprintf("%q", str.String())))
	} else {
		f.Write([]byte(fmt.Sprintf("%s", str.String())))
	}
}

// String gets the contents as the built-in string type.
func (str String) String() string {
	return string(str.r)
}

// CharAt returns the runes that make up the user-perceived character (grapheme
// cluster) of the given index. Modifying the returned slice will not modify the
// String.
func (str String) CharAt(idx int) []rune {
	gc := str.gc
	if gc == nil {
		gc = Split(str.r)
		str.gc = gc
	}

	var startIdx = 0
	if idx > 0 {
		startIdx = gc[idx-1]
	}
	length := gc[idx] - startIdx
	cluster := make([]rune, length)
	for i := 0; i < length; i++ {
		cluster[i] = str.r[startIdx+i]
	}
	return cluster
}

// Add adds two strings together and returns the result. The original String is
// not modified.
func (str String) Add(s2 String) String {
	r2 := str.clone()
	r2.gc = nil
	r2.r = append(r2.r, s2.Runes()...)
	return r2
}

// Len returns the number of grapheme clusters (user-perceivable characters)
// that are in the String.
//
// This function may trigger UAX29 analysis on the String if it hasn't yet
// occured.
func (str String) Len() int {
	gc := str.gc
	if gc == nil {
		if len(str.r) == 0 {
			return 0
		}
		gc = Split(str.r)
		str.gc = gc
	}
	return len(gc)
}

// makes an exact duplicate by copying underlying slices.
//
// modifications to any members of the returned string are guaranteed not to
// modify the original. calling this is not needed unless a modification is
// about to occur, even though passing String by value does pass pointers (via
// slice-type members)
//
// TODO: gem.String is generally passed by value now and immutable. Is this still
// needed?
func (str String) clone() String {
	gc := str.gc
	clone := String{
		r: make([]rune, len(str.r)),
	}
	for i := range str.r {
		clone.r[i] = str.r[i]
	}
	if gc != nil {
		clone.gc = make([]int, len(gc))
		for i := range gc {
			clone.gc[i] = gc[i]
		}
	}
	return clone
}

// New takes the given string and converts it into a graphemes.String object for
// use with grapheme-aware functions. UAX-29 analysis is performed on a lazy
// basis; the contents of s are not scanned for grapheme clusters until an
// operation requires it.
func New(s string) String {
	return String{r: []rune(s)}
}

// Char creates a String from single user-perceived character made up of the
// given runes.
func Char(ch []rune) String {
	return New(string(ch))
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
