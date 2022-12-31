// Package gem provides operations for individual user-perceived
// characters and implements UAX #29 by Unicode for grapheme boundary finding.
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
// The zero-value is an empty String, ready to use.
//
// String.Equal can be used to test against raw strings.
//
// String implements the fmt.Formatter interface.
type String struct {
	r []rune

	// gc is a pointer but should never be set to nil explicitly. However, zero
	// value creation will result in it being set to nil, so it must still be
	// checked with every access.
	gc *[]int
}

var (
	// Zero is a String of zero length.
	Zero String = String{r: []rune{}, gc: new([]int)}
)

// Add adds two strings together and returns the result. The original String is
// not modified.
func (str String) Add(s2 String) String {
	str = str.initialized()

	r2 := str.clone()
	*r2.gc = nil

	r2.r = append(r2.r, s2.Runes()...)
	return r2
}

// CharAt returns the runes that make up the user-perceived character (grapheme
// cluster) of the given index. Modifying the returned slice will not modify the
// String.
func (str String) CharAt(idx int) []rune {
	str = str.initialized()

	if *str.gc == nil {
		*str.gc = Split(str.r)
	}

	var startIdx = 0
	if idx > 0 {
		startIdx = (*str.gc)[idx-1]
	}
	length := (*str.gc)[idx] - startIdx
	cluster := make([]rune, length)
	for i := 0; i < length; i++ {
		cluster[i] = str.r[startIdx+i]
	}
	return cluster
}

// Equal returns whether one String is equal to another object. If the object is
// another String struct, their resulting strings are compared. If the object is
// a raw string object, it is compared to the output of calling String() on the
// gem.String. Otherwise, false is returned.
func (str String) Equal(other interface{}) bool {
	str = str.initialized()

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

// Format formats the String for printing. This is part of the implementation
// of fmt.Formatter.
func (str String) Format(f fmt.State, verb rune) {
	str = str.initialized()

	if verb == 'q' {
		f.Write([]byte(fmt.Sprintf("%q", str.String())))
	} else {
		f.Write([]byte(str.String()))
	}
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
	str = str.initialized()

	if *str.gc == nil {
		*str.gc = Split(str.r)
	}

	gc := *str.gc
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

// IsEmpty return whether the String is the empty string "".
func (str String) IsEmpty() bool {
	str = str.initialized()

	return len(str.r) == 0
}

// Len returns the number of grapheme clusters (user-perceivable characters)
// that are in the String.
//
// This function may trigger UAX29 analysis on the String if it hasn't yet
// occured.
func (str String) Len() int {
	str = str.initialized()

	if *str.gc == nil {
		if len(str.r) == 0 {
			return 0
		}
		*str.gc = Split(str.r)
	}

	return len(*str.gc)
}

// Less returns whether one String is lexigraphically less than another.
func (str String) Less(s String) bool {
	str = str.initialized()

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

// Runes returns the string's raw Runes. Modifying the returned slice has no
// effect on the String.
func (str String) Runes() []rune {
	str = str.initialized()

	r := make([]rune, len(str.r))
	copy(r, str.r)
	return r
}

// SetCharAt sets the character at the given index to the given value and
// returns the resulting String. The original String is not modified.
func (str String) SetCharAt(idx int, r []rune) String {
	str = str.initialized()

	if len(r) == 0 {
		panic("SetCharAt received empty or nil replacement runes slice r")
	}

	clone := str.clone()

	if *clone.gc == nil {
		*clone.gc = Split(clone.r)
	}

	var startIdx int
	if idx > 0 {
		startIdx = (*clone.gc)[idx-1]
	}
	curEndIdx := (*clone.gc)[idx]
	clone.r = append(clone.r[:startIdx], append(r, clone.r[curEndIdx:]...)...)
	*clone.gc = nil
	return clone
}

// String gets the contents as the built-in string type.
func (str String) String() string {
	str = str.initialized()

	return string(str.r)
}

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
	str = str.initialized()

	start, end = util.RangeToIndexes(str.Len(), start, end)

	if start == end {
		return Zero
	}

	if *str.gc == nil {
		// we need a split operation on the graphemes
		*str.gc = Split(str.r)
	}

	clone := str.clone()

	var runesStart int
	if start > 0 {
		runesStart = (*clone.gc)[start-1]
	}
	runesEnd := (*clone.gc)[end-1]

	clone.r = clone.r[runesStart:runesEnd]
	*clone.gc = (*clone.gc)[start:end]

	// but if we've sub'd from anywhere but the start, because every value in
	// the gc slice is really the difference in runes from the *prior* value,
	// we need to subtract whatever came before the sub'd index.
	if runesStart > 0 {
		for i := range *clone.gc {
			(*clone.gc)[i] -= runesStart
		}
	}

	return clone
}

// New takes the given string and converts it into a graphemes.String object for
// use with grapheme-aware functions. UAX-29 analysis is performed on a lazy
// basis; the contents of s are not scanned for grapheme clusters until an
// operation requires it.
func New(s string) String {
	return String{r: []rune(s), gc: new([]int)}
}

// Repeat returns a new String consisting of count copies of the String s.
func Repeat(s String, count int) String {
	repeated := Zero
	for i := 0; i < count; i++ {
		repeated = repeated.Add(s)
	}
	return repeated
}

// RepeatStr returns a new String consisting of count copies of the string s. It
// is exactly the same as [Repeat] but it allows the user to pass in a regular
// string.
func RepeatStr(s string, count int) String {
	return Repeat(New(s), count)
}

// Slice turns the from slice into a slice of String objects.
func Slice(from []string) []String {
	str := make([]String, len(from))
	for i := range from {
		str[i] = New(from[i])
	}
	return str
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

// Strings turns the from slice into a slice of plain string objects.
func Strings(from []String) []string {
	str := make([]string, len(from))
	for i := range from {
		str[i] = from[i].String()
	}
	return str
}

// makes an exact duplicate by copying underlying slices.
//
// modifications to any members of the returned string are guaranteed not to
// modify the original. calling this is not needed unless a modification is
// about to occur, even though passing String by value does pass pointers (via
// slice-type members)
//
// gem.String is generally passed by value now and immutable, but keeping this
// because it makes it convenient for deep-copying the members of it.
func (str String) clone() String {
	str = str.initialized()

	clone := String{
		r:  make([]rune, len(str.r)),
		gc: new([]int),
	}

	copy(clone.r, str.r)

	if *str.gc != nil {
		newCloneGC := make([]int, len(*str.gc))
		clone.gc = &newCloneGC
		copy(*clone.gc, *str.gc)
	}

	return clone
}

// initialized returns a new String with all properties set to values suitable
// for immediate use. It is used to convert a String created with no setting of
// internal properties to one that is ready for use, with gc explicitly set to a
// non-nil pointer (though it may point to a nil slice), and r set to a non-nil
// slice.
func (str String) initialized() String {
	if str.gc == nil {
		str.gc = new([]int)
	}

	if str.r == nil {
		str.r = []rune{}
	}

	return str
}
