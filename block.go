package rosed

import (
	"fmt"
	"strings"

	"github.com/dekarrin/rosed/internal/gem"
)

// block is a "block" of text, that is some number of lines. The line separator
// and Trailing Separator behavior are configurable, but if a constructor
// function is used, they may be automatically detected.
//
// Blocks are mutable and are not thread-safe.
//
// The Zero-value for Block is a block with no lines, with LineSeparator unset
// and TrailingSeparator set to false.
type block struct {
	Lines             []gem.String
	LineSeparator     gem.String
	TrailingSeparator bool
}

// String returns the string representation of the block.
func (tb block) String() string {
	return fmt.Sprintf("<Block LineSeparator:%q TrailingSeparator:%v Content:%q>", tb.LineSeparator, tb.TrailingSeparator, tb.Lines)
}

// newBlock creates a block of text.
//
// If text is present, it is split into lines using the provided lineSep. If any
// trailing line separator is present, it will be removed prior to split.
//
// The returned Block will contain the lines that were obtained from text. If
// text was empty, it will have no lines. If text consisted only of lineSep,
// there will be exactly empty line in the resulting block.
//
// The returned Block will have TrailingSeparator set to match whatever mode
// the passed in text had; if it was empty, TrailingSeparator will always be
// false.
func newBlock(text, lineSep gem.String) block {
	var trailing bool

	// handle special cases of text being empty or text being only the line
	// separator
	if text.IsEmpty() {
		return block{
			LineSeparator: lineSep,
		}
	}

	lines := strings.Split(text.String(), lineSep.String())
	if len(lines) > 1 && lines[len(lines)-1] == "" {
		lines = lines[0 : len(lines)-1]
		trailing = true
	}
	bl := block{
		Lines:             make([]gem.String, len(lines)),
		LineSeparator:     lineSep,
		TrailingSeparator: trailing,
	}

	for i := range lines {
		bl.Lines[i] = gem.S(lines[i])
	}

	return bl
}

// Equal checks whether this Block is equal to another object. Returns whether
// other is also a Block with the same contents as tb.
func (tb block) Equal(other interface{}) bool {
	b2, ok := other.(block)
	if !ok {
		return false
	}

	if tb.LineSeparator != b2.LineSeparator {
		return false
	}
	if tb.TrailingSeparator != b2.TrailingSeparator {
		return false
	}

	// don't use deep equal because it will fail if one has nil Lines and
	// another has empty Lines even though that case should compare equal.
	if len(tb.Lines) != len(b2.Lines) {
		return false
	}
	for i := 0; i < len(tb.Lines); i++ {
		if tb.Lines[i].Equal(b2.Lines[i]) {
			return false
		}
	}
	return true
}

// Len gives the number of lines in the block.
func (tb block) Len() int {
	return len(tb.Lines)
}

// Less reports whether line i should sort before line j.
func (tb block) Less(i, j int) bool {
	return tb.Line(i).Less(tb.Line(j))
}

// Swap swaps line i with line j.
func (tb block) Swap(i, j int) {
	tb.Lines[i], tb.Lines[j] = tb.Lines[j], tb.Lines[i]
}

// Append adds a new line to the block.
func (tb *block) Append(content gem.String) {
	if len(tb.Lines) < 1 {
		tb.Lines = []gem.String{content}
		return
	}
	tb.Lines = append(tb.Lines, content)
}

// AppendBlock adds all lines in the given block to the end of this one.
func (tb *block) AppendBlock(b block) {
	for i := 0; i < b.Len(); i++ {
		tb.Append(b.Line(i))
	}
}

// AppendEmpty adds the given number of empty lines to the Block.
func (tb *block) AppendEmpty(count int) {
	for i := 0; i < count; i++ {
		tb.Append(gem.Z)
	}
}

// Set sets a line to new contents.
//
// linePos must be a line that exists.
func (tb *block) Set(linePos int, content gem.String) {
	tb.Lines[linePos] = content
}

// Apply performs the given transformation on each line and applies the results
// to its list of lines. Lines may be both added and removed this way.
//
// It is important to note that the Block will always have at least one line,
// even in the case of a Block with no content added (the line will just be
// empty). In this case, the transform function will be called once with an
// empty string as its argument.
//
// All lines should be assumed to not have line terminators, and none should be
// added.
func (tb *block) Apply(transform LineOperation) {
	var applied []gem.String

	for idx, line := range tb.Lines {
		applied = append(applied, gem.Slice(transform(idx, line.String()))...)
	}

	tb.Lines = applied
}

func (tb *block) applyGFunc(transform LineOperation) {
	var applied []gem.String

	for idx, line := range tb.Lines {
		applied = append(applied, gem.Slice(transform(idx, line.String()))...)
	}

	tb.Lines = applied
}

// Line gives the line at a position. pos is not checked for validity before
// accessing; callers must do so.
func (tb block) Line(pos int) gem.String {
	return tb.Lines[pos]
}

// CharCount returns the number of characters in the given line which will not
// include the separator.
func (tb block) CharCount(linePos int) int {
	return tb.Line(linePos).Len()
}

// Join converts the block into a single string by appending each line to the
// previous with the line separator given at construction.
func (tb block) Join() gem.String {
	if tb.Len() < 1 {
		if tb.TrailingSeparator {
			return tb.LineSeparator
		}
		return gem.Z
	}

	full := strings.Join(gem.Strings(tb.Lines), tb.LineSeparator.String())
	str := gem.S(full)
	if tb.TrailingSeparator {
		str = str.Add(tb.LineSeparator)
	}
	return str
}
