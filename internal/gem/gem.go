// Package gem provides operations for individual user-perceived
// characters and implements UAX #29 by Unicode for grapheme boundary finding.
//
// It can be used to show the number of characters as a user would perceive one
// character to be, implementing the rules specified by Unicode to be safe
// across multiple character ranges.
package gem

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

// returns whether two grapheme clusters returned by CharAt and other functions
// are equal.
func graphemesEqual(gc1, gc2 []rune) bool {
	if len(gc1) != len(gc2) {
		return false
	}

	for i := range gc1 {
		if gc1[i] != gc2[i] {
			return false
		}
	}

	return true
}
