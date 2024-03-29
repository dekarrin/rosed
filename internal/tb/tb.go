// Package tb holds the primitives for working with blocks of gem.String
// text.
package tb

import (
	"fmt"
	"strings"

	"github.com/dekarrin/rosed/internal/gem"
)

// BlockLineOperation is a function that accepts a zero-indexed line number and
// the contents of that line and performs some operation to produce zero or more
// new lines to replace the contents of the line with.
//
// The return value for a BlockLineOperation is a slice of lines to insert at
// the old line position. This can be used to delete the line or insert
// additional new ones; to insert, include the new lines in the returned slice
// in the proper position relative to the old line in the slice, and to delete
// the original line, a slice with len < 1 can be returned.
//
// The parameter idx will always be the index of the line before any
// transformations were applied; e.g. a call to a BlockLineOperation with idx =
// 4 will always be after a call with idx = 3, regardless of the size of the
// returned slice in the prior call.
type BlockLineOperation func(idx int, line string) []string

// Block holds lines of text in a line-terminator agnostic way and provides a
// way to operate on each individually, independent of what the original
// line terminator behavior is.
//
// The LineSeparator and TrailingSeparator properties can be manually set,
// but if newBlock is used to create a block, they may be automatically detected.
//
// Blocks are mutable and are not thread-safe.
//
// The Zero-value for Block is safe to use and is a block with no lines, with
// LineSeparator unset and TrailingSeparator set to false.
//
// block implements the sort.Interface interface.
type Block struct {
	Lines             []gem.String
	LineSeparator     gem.String
	TrailingSeparator bool
}

// Append adds a new line to the block.
func (tb *Block) Append(content gem.String) {
	if len(tb.Lines) < 1 {
		tb.Lines = []gem.String{content}
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

// AppendEmpty adds the given number of empty lines to the Block.
func (tb *Block) AppendEmpty(count int) {
	for i := 0; i < count; i++ {
		tb.Append(gem.Zero)
	}
}

// Apply performs the given transformation on each line and applies the results
// to its list of lines. Lines may be both added and removed this way.
//
// If called on a block with empty lines, the transform function will be called
// once with an empty string as its argument, allowing a caller to use the
// LineOperation to produce initial text.
//
// All lines received by transform should be assumed to not have line
// terminators, and none should be added by it.
func (tb *Block) Apply(transform BlockLineOperation) {
	var applied []gem.String

	for idx, line := range tb.Lines {
		applied = append(applied, gem.Slice(transform(idx, line.String()))...)
	}

	tb.Lines = applied
}

// CharCount returns the number of characters in the given line which will not
// include the separator.
func (tb Block) CharCount(linePos int) int {
	return tb.Line(linePos).Len()
}

// Equal checks whether this Block is equal to another object. Returns whether
// other is also a block with the same contents and separator behavior as tb.
func (tb Block) Equal(other interface{}) bool {
	b2, ok := other.(Block)
	if !ok {
		return false
	}

	if tb.LineSeparator.IsEmpty() {
		if !b2.LineSeparator.IsEmpty() {
			return false
		}
	} else if b2.LineSeparator.IsEmpty() {
		return false
	} else if !tb.LineSeparator.Equal(b2.LineSeparator) {
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
		if !tb.Lines[i].Equal(b2.Lines[i]) {
			return false
		}
	}
	return true
}

// Join converts the block into a single string by appending each line to the
// previous. Its currently-set LineSeparator is placed between each line and
// if and only if TrailingSeparator is set, the LineSeparator is placed at the
// end of the string as well.
func (tb Block) Join() gem.String {
	if tb.Len() < 1 {
		if tb.TrailingSeparator {
			return tb.LineSeparator
		}
		return gem.Zero
	}

	full := strings.Join(gem.Strings(tb.Lines), tb.LineSeparator.String())
	str := gem.New(full)
	if tb.TrailingSeparator {
		str = str.Add(tb.LineSeparator)
	}
	return str
}

// Len gives the number of lines in the block. This is part of the implementation of
// sort.Interface.
func (tb Block) Len() int {
	return len(tb.Lines)
}

// Less reports whether line i should sort before line j. This is part of the implementation of
// sort.Interface.
func (tb Block) Less(i, j int) bool {
	return tb.Line(i).Less(tb.Line(j))
}

// Line returns the line at the given position. pos is not checked for validity before
// accessing; callers must do so.
func (tb Block) Line(pos int) gem.String {
	return tb.Lines[pos]
}

// Remove removes the line at the given position. If pos does not exist, no action is
// taken.
func (tb *Block) Remove(pos int) {
	if pos >= 0 && len(tb.Lines) > pos {
		tb.Lines = append(tb.Lines[:pos], tb.Lines[pos+1:]...)
	}
}

// Set sets a line to new contents.
//
// linePos must be a line that exists.
func (tb *Block) Set(linePos int, content gem.String) {
	tb.Lines[linePos] = content
}

// String returns the string representation of the block.
func (tb Block) String() string {
	return fmt.Sprintf("<block LineSeparator:%q TrailingSeparator:%v Lines:%q>", tb.LineSeparator, tb.TrailingSeparator, tb.Lines)
}

// Swap swaps line i with line j. This is part of the implementation of
// sort.Interface.
func (tb Block) Swap(i, j int) {
	tb.Lines[i], tb.Lines[j] = tb.Lines[j], tb.Lines[i]
}

// New creates a block of text.
//
// The text is split into lines using the provided lineSep. If any trailing line
// separator is present, it will be removed prior to split.
//
// The returned Block will contain the lines that were obtained from text. If
// text was empty, it will have no lines. If text consisted only of lineSep,
// there will be exactly empty line in the resulting block.
//
// The returned Block will have TrailingSeparator set to match whatever mode
// the passed in text had; if it was empty, TrailingSeparator will always be
// false.
func New(text, lineSep gem.String) Block {
	var trailing bool

	// handle special cases of text being empty or text being only the line
	// separator
	if text.IsEmpty() {
		return Block{
			LineSeparator: lineSep,
		}
	}

	lines := strings.Split(text.String(), lineSep.String())
	if len(lines) > 1 && lines[len(lines)-1] == "" {
		lines = lines[0 : len(lines)-1]
		trailing = true
	}
	bl := Block{
		Lines:             make([]gem.String, len(lines)),
		LineSeparator:     lineSep,
		TrailingSeparator: trailing,
	}

	for i := range lines {
		bl.Lines[i] = gem.New(lines[i])
	}

	return bl
}
