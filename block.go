package rosed

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

// Block is a "block" of text, that is some number of lines. The line separator
// and Trailing Separator behavior are configurable, but if a constructor
// function is used, they may be automatically detected.
//
// Blocks are mutable and are not thread-safe.
//
// The Zero-value for Block is a block with no lines, with LineSeparator unset
// and TrailingSeparator set to false.
type Block struct {
	Lines             []string
	LineSeparator     string
	TrailingSeparator bool
}

// String returns the string representation of the block.
func (tb Block) String() string {
	return fmt.Sprintf("<Block LineSeparator:%q TrailingSeparator:%v Content:%q>", tb.LineSeparator, tb.TrailingSeparator, tb.Lines)
}

// NewBlock creates a block of text.
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
func NewBlock(text, lineSep string) Block {
	var trailing bool

	// handle special cases of text being empty or text being only the line
	// separator
	if text == "" {
		return Block{
			LineSeparator: lineSep,
		}
	}

	lines := strings.Split(text, lineSep)
	if len(lines) > 1 && lines[len(lines)-1] == "" {
		lines = lines[0 : len(lines)-1]
		trailing = true
	}
	return Block{
		Lines:             lines,
		LineSeparator:     lineSep,
		TrailingSeparator: trailing,
	}
}

// Equal checks whether this Block is equal to another object. Returns whether
// other is also a Block with the same contents as tb.
func (tb Block) Equal(other interface{}) bool {
	b2, ok := other.(Block)
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
		if tb.Lines[i] != b2.Lines[i] {
			return false
		}
	}
	return true
}

// Len gives the number of lines in the block.
func (tb Block) Len() int {
	return len(tb.Lines)
}

// Less reports whether line i should sort before line j.
func (tb Block) Less(i, j int) bool {
	return tb.Line(i) < tb.Line(j)
}

// Swap swaps line i with line j.
func (tb Block) Swap(i, j int) {
	tb.Lines[i], tb.Lines[j] = tb.Lines[j], tb.Lines[i]
}

// Append adds a new line to the block.
func (tb *Block) Append(content string) {
	if len(tb.Lines) < 1 {
		tb.Lines = []string{content}
		return
	}
	tb.Lines = append(tb.Lines, content)
}

// AppendBlock adds all lines in the given block to the end of this one.
func (tb *Block) AppendBlock(b Block) {
	for i := 0; i < b.Len(); i++ {
		tb.Append(b.Line(i))
	}
}

// Set sets a line to new contents.
//
// linePos must be a line that exists.
func (tb *Block) Set(linePos int, content string) {
	tb.Lines[linePos] = content
}

// AddEmpty adds the given number of empty lines to the Block.
func (tb *Block) AddEmpty(count int) {
	for i := 0; i < count; i++ {
		tb.Append("")
	}
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
func (tb *Block) Apply(transform LineOperation) {
	var applied []string

	for idx, line := range tb.Lines {
		applied = append(applied, transform(idx, line)...)
	}

	tb.Lines = applied
}

// Line gives the line at a position. pos is not checked for validity before
// accessing; callers must do so.
func (tb Block) Line(pos int) string {
	return tb.Lines[pos]
}

// RuneCount returns the number of utf8 codepoints in the given line which
// will not include the separator.
func (tb Block) RuneCount(linePos int) int {
	return utf8.RuneCountInString(tb.Line(linePos))
}

// Join converts the block into a single string by appending each line to the
// previous with the line separator given at construction.
func (tb Block) Join() string {
	if tb.Len() < 1 {
		if tb.TrailingSeparator {
			return tb.LineSeparator
		}
		return ""
	}

	full := strings.Join(tb.Lines, tb.LineSeparator)
	if tb.TrailingSeparator {
		full += tb.LineSeparator
	}
	return full
}
