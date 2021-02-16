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
	lines             []string
	LineSeparator     string
	TrailingSeparator bool
}

// String returns the string representation of the block.
func (tb Block) String() string {
	return fmt.Sprintf("<Block LineSeparator:%q TrailingSeparator:%b Content:%q>", tb.LineSeparator, tb.TrailingSeparator, tb.lines)
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
	if text == lineSep {
		return Block{
			lines:             []string{""},
			LineSeparator:     lineSep,
			TrailingSeparator: true,
		}
	}

	lines := strings.Split(text, lineSep)
	if len(lines) > 1 && lines[len(lines)-1] == "" {
		lines = lines[0 : len(lines)-1]
		trailing = true
	}
	return Block{
		lines:             lines,
		LineSeparator:     lineSep,
		TrailingSeparator: trailing,
	}
}

// Len gives the number of lines in the block.
func (tb Block) Len() int {
	return len(tb.lines)
}

// Less reports whether line i should sort before line j.
func (tb Block) Less(i, j int) bool {
	return tb.Line(i) < tb.Line(j)
}

// Swap swaps line i with line j.
func (tb Block) Swap(i, j int) {
	tb.lines[i], tb.lines[j] = tb.lines[j], tb.lines[i]
}

// Append adds a new line to the block.
func (tb *Block) Append(content string) {
	if len(tb.lines) < 1 {
		tb.lines = []string{content}
		return
	}
	tb.lines = append(tb.lines, content)
}

// Set sets a line to new contents.
//
// linePos must be a line that exists.
func (tb *Block) Set(linePos int, content string) {
	tb.lines[linePos] = content
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

	for idx, line := range tb.lines {
		applied = append(applied, transform(idx, line)...)
	}

	tb.lines = applied
}

// Line gives the line at a position. pos is not checked for validity before
// accessing; callers must do so.
func (tb Block) Line(pos int) string {
	return tb.lines[pos]
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

	full := strings.Join(tb.lines, tb.LineSeparator)
	if tb.TrailingSeparator {
		full += tb.LineSeparator
	}
	return full
}
