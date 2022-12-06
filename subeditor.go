package rosed

// this file contains functions for splitting an Editor into a sub-Editor.

import (
	"strings"

	"github.com/dekarrin/rosed/internal/gem"
	"github.com/dekarrin/rosed/internal/util"
)

// reference to parent for sub-editors. ed[start:end] gives the Text of the
// sub-editor
type parentRef struct {
	// Parent editor. The sub-editor is operating on a substring of the parent's
	// Text property; parent.Text[start:end] is what is replaced by the
	// sub-editor's results.
	parent *Editor

	// index in parent of start of the substring we are operating on;
	// parent.Text[start:end] is what is replaced by the sub-editor.
	start int

	// index in parent of end of the sub-string we are operating on;
	// parent.Text[start:end] is what is replaced by the sub-editor.
	end int
}

// Chars produces an Editor to operate on a subset of its Text. The returned
// Editor operates on text from the nth character up to (but not including)
// the ith character, where n is `start` and i is `end`.
//
// For instance, to get the 'ell' of "Hello!":
//
//	ed := Edit("Hello!")
//	ed = ed.Chars(1, 4)
//
//	fmt.Printf("%v\n", ed.Text)  // will be "ell".
//
// The `start` or `end` may be negative, in which case it will be relative to
// the end of the string; -1 would be the index of the last character, -2 would
// be the index of the second-to-last character, etc.
//
// If `start` or `end` specifies an index that is past the end of the string,
// that index is assumed to be the end of the string. If either specify an index
// that is before the start of a string, it is assumed to be 0.
//
// If `end` is less than `start`, it is assumed to be equal to `start`.
//
// # Grapheme-Aware Function
//
// The `start` and `end` index the human-readable graphemes in Text, not the
// bytes or runes that make it up; see the note on `grapheme-based indexing` for
// more info on this.
//
// # Sub-editor Function
//
// This function produces a sub-editor, whose Text field will contain only the
// sub-section of text specified. Editing the parent's Text field after the
// sub-editor has been created will have no effect on the sub-editor or any
// Editor produced from it, but note that it may have an affect on the result of
// calling Commit() on the sub-editor.
//
// The sub-editor as well as any sub-editors produced from it can be merged back
// into the original Editor by calling Commit(). Alternatively, all such
// sub-editors can be merged recursively by calling CommitAll().
func (ed Editor) Chars(start, end int) Editor {
	// ask gem string for the grapheme-based char positions
	indexes := gem.New(ed.Text).GraphemeIndexes()
	start, end = util.RangeToIndexes(len(indexes), start, end)

	// interface treats these as python-style slice indexes which means we
	// accept start == the end of the string, but we are about to use them
	// as proper indexes which means that will panic. So if start is past the
	// end, immediately return the subEd of that without further checking
	if start >= len(indexes) {
		return ed.subEd(start, end)
	}

	runeStart := indexes[start][0]

	var runeEnd int
	if end < len(indexes) {
		runeEnd = indexes[end][0]
	} else {
		runeEnd = len(ed.Text)
		// TODO: in this case the below search will go until end of string
		// could make this a bit betta by marking this case and skipping the
		// search for runeEnd
	}

	// now that we have rune indexes we do string analysis to find the byte
	// boundaries of the chars

	chIdx := -1
	byteStart := -1
	byteEnd := -1
	for byteIdx := range ed.Text {
		chIdx++
		if chIdx == runeStart {
			byteStart = byteIdx
		}
		if chIdx == runeEnd {
			byteEnd = byteIdx
			break
		}
	}
	if byteEnd == -1 {
		byteEnd = len(ed.Text)
	}

	return ed.subEd(byteStart, byteEnd)
}

// CharsFrom produces an Editor to operate on a subset of its Text. The returned
// Editor operates on text from the nth character up to the end of the string,
// where n is `start`.
//
// For instance, to get the 'llo!' of "Hello!":
//
//	ed := Edit("Hello!")
//	ed = ed.CharsFrom(2)
//
//	fmt.Printf("%v\n", ed.Text)  // will be "llo!".
//
// `start` may be negative, in which case it will be relative to the end of the
// string; -1 would be the index of the last character, -2 would be the index of
// the second-to-last character, etc.
//
// If `start` specifies an index that is past the end of the string, that index
// is assumed to be the end of the string. If it specifies an index that is
// before the start of a string, it is assumed to be 0.
//
// # Grapheme-Aware Function
//
// The `start` indexes the human-readable graphemes in Text, not the bytes or
// runes that make it up; see the note on `grapheme-based indexing` for more
// info on this.
//
// # Sub-editor Function
//
// This function produces a sub-editor, whose Text field will contain only the
// sub-section of text specified. Editing the parent's Text field after the
// sub-editor has been created will have no effect on the sub-editor or any
// Editor produced from it, but note that it may have an affect on the result of
// calling Commit() on the sub-editor.
//
// The sub-editor as well as any sub-editors produced from it can be merged back
// into the original Editor by calling Commit(). Alternatively, all such
// sub-editors can be merged recursively by calling CommitAll().
func (ed Editor) CharsFrom(start int) Editor {
	return ed.Chars(start, len(ed.Text))
}

// CharsTo produces an Editor to operate on a subset of its Text. The returned
// Editor operates on text from the first character up to but not including the
// nth character, where n is `end`.
//
// For instance, to get the 'He' of "Hello!":
//
//	ed := Edit("Hello!")
//	ed = ed.CharsTo(2)
//
//	fmt.Printf("%v\n", ed.Text)  // will be "He".
//
// `end` may be negative, in which case it will be relative to the end of the
// string; -1 would be the index of the last character, -2 would be the index of
// the second-to-last character, etc.
//
// If `end` specifies an index that is past the end of the string, that index is
// assumed to be the end of the string. If it specifies an index that is before
// the start of a string, it is assumed to be 0.
//
// # Grapheme-Aware Function
//
// The `end` indexes the human-readable graphemes in Text, not the bytes or
// runes that make it up; see the note on `grapheme-based indexing` for more
// info on this.
//
// # Sub-editor Function
//
// This function produces a sub-editor, whose Text field will contain only the
// sub-section of text specified. Editing the parent's Text field after the
// sub-editor has been created will have no effect on the sub-editor or any
// Editor produced from it, but note that it may have an affect on the result of
// calling Commit() on the sub-editor.
//
// The sub-editor as well as any sub-editors produced from it can be merged back
// into the original Editor by calling Commit(). Alternatively, all such
// sub-editors can be merged recursively by calling CommitAll().
func (ed Editor) CharsTo(end int) Editor {
	return ed.Chars(0, end)
}

// Commit takes the substring that a sub-editor is operating on and merges it
// with its parent.
//
// Returns an Editor which is a copy of the current one but with its Text set to
// the merged string.
//
// If the Editor is already full-text Editor, the merge operation is simply to
// copy the current Text since there is no prefix or suffix, so calling Commit
// returns a duplicate of the Editor.
//
// Operations on the returned Editor will be on the parent's string (if any)
// rather than on the subset of it.
func (ed Editor) Commit() Editor {
	if !ed.IsSubEditor() {
		return ed
	}

	parent, subStart, subEnd := ed.ref.parent, ed.ref.start, ed.ref.end

	prefix := parent.Text[:subStart]
	suffix := parent.Text[subEnd:]

	full := prefix + ed.Text + suffix

	// copy via value assignment
	ed = *parent
	ed.Text = full
	return ed
}

// CommitAll takes the substring that a sub-editor is operating on and merges it
// with its parent recursively to get a new complete string.
//
// Returns an Editor which is a copy of the current one but with its Text set to
// the merged complete string.
//
// If the Editor is already a full-text Editor, the merge operation is simply to
// copy the current Text since there is no prefix or suffix, so calling
// CommitAll simply returns a copy of the Editor.
//
// Operations on the returned Editor will be on the complete string rather than
// on a subset of it.
func (ed Editor) CommitAll() Editor {
	// recursively do commit until we get to a full-text Editor.
	for ed.IsSubEditor() {
		ed = ed.Commit()
	}
	return ed
}

// IsSubEditor returns whether the Editor was created to edit a sub-set of the
// text in some parent editor. Calls to Lines(), LinesFrom(), LinesTo(),
// Chars(), CharsFrom(), and CharsTo() will result in such an Editor.
//
// If IsSubEditor returns true, then the Editor's Text could possibly be an
// incomplete subset of the original text. To get the full text from a
// sub-editor, use CommitAll to get the root parent editor with all pending
// changes from sub-editors, including this one, applied.
func (ed Editor) IsSubEditor() bool {
	return ed.ref != nil
}

// Lines produces an Editor to operate on a subset of the lines in the Text. The
// lines are 0-indexed and the `start` and `end` are the same as in slice
// notation. The returned Editor operates on lines from the nth line up to (but
// not including) the ith line, where n is `start` and i is `end`.
//
// For instance, to get the middle two lines of a four-line string:
//
//	ed := Edit("Line #1\nLine #2\nLine #3\nLine #4")
//	ed = ed.Lines(1, 3)
//
//	fmt.Printf("%v\n", ed.Text)  // will be "Line #2\nLine #3\n".
//
// The `start` or `end` may be negative, in which case it will be relative to
// the end of the text; -1 would be the index of the last line, -2 would be the
// index of the second-to-last line, etc.
//
// If `start` or `end` specifies an index that is past the end of the text, that
// index is assumed to be the end of the text (i.e. 1 greater than the index of
// the final line). If either specify an index that is before the start of a
// string, it is assumed to be 0.
//
// If `end` is less than `start`, it is assumed to be equal to `start`.
//
// # Sub-editor Function
//
// This function produces a sub-editor, whose Text field will contain only the
// sub-section of text specified. Editing the parent's Text field after the
// sub-editor has been created will have no effect on the sub-editor or any
// Editor produced from it, but note that it may have an affect on the result of
// calling Commit() on the sub-editor.
//
// The sub-editor as well as any sub-editors produced from it can be merged back
// into the original Editor by calling Commit(). Alternatively, all such
// sub-editors can be merged recursively by calling CommitAll().
//
// # Options
//
// This function is affected by the following options of the Editor it is called
// on:
//
//   - `LineSeparator` specifies what string should be used to delimit lines.
//   - `NoTrailingLineSeparators` specifies whether it should consider a final
//     instance of `LineSeparator` to be ending the prior line, or giving the
//     start of a new line. If NoTrailingLineSeparators is true, a trailing
//     LineSeparator is considered to start a new (empty) line.
func (ed Editor) Lines(start, end int) Editor {
	if ed.Text == "" {
		return ed.subEd(0, 0)
	}

	lc := ed.LineCount()
	start, end = util.RangeToIndexes(lc, start, end)

	// if we know we are about to get past the end of the lines
	// skip the costly search and just get the end
	if start >= lc {
		return ed.subEd(len(ed.Text), len(ed.Text))
	}

	lineSep := ed.Options.WithDefaults().LineSeparator

	lineIdx := 0
	byteStart := 0
	for lineIdx != start {
		lineSepStart := strings.Index(ed.Text[byteStart:], lineSep)
		if lineSepStart == -1 {
			// we are on the last line and haven't gotten to
			// the start, so user is asking for lines beyond the end.
			// should never happen due to above checks but handle it
			// anyways just in case
			return ed.subEd(len(ed.Text), len(ed.Text))
		}

		// byteStart is also the start of the line we are searching from
		// during the for-loop
		byteStart += lineSepStart + len(lineSep)
		lineIdx++
	}

	// byteStart should now be the correct value,
	// now find byteEnd
	byteEnd := byteStart

	for lineIdx != end {
		lineSepStart := strings.Index(ed.Text[byteEnd:], lineSep)
		if lineSepStart == -1 {
			// we are on the last line and there is no trailing newline.
			// we don't actually care about that tho, only relevant info is
			// that we are on the last line and so can stop searching
			return ed.subEd(byteStart, len(ed.Text))
		}

		// byteEnd is also the start of the line we are searching from
		// during the for-loop
		byteEnd += lineSepStart + len(lineSep)
		lineIdx++
	}

	return ed.subEd(byteStart, byteEnd)
}

// LinesFrom produces an Editor to operate on a subset of the lines in the Text.
// The returned Editor operates on lines from the nth line up to the end of the
// text, where n is `start`.
//
// For instance, to get the last two lines of a four-line string:
//
//	ed := Edit("Line #1\nLine #2\nLine #3\nLine #4")
//	ed = ed.LinesFrom(2)
//
//	fmt.Printf("%v\n", ed.Text)  // will be "Line #3\nLine #4".
//
// The `start` may be negative, in which case it will be relative to the end of
// the text; -1 would be the index of the last line, -2 would be the index of
// the second-to-last line, etc.
//
// If `start` specifies an index that is past the end of the text, that index is
// assumed to be the end of the text (i.e. 1 greater than the index of the final
// line). If it specifies an index that is before the start of a string, it is
// assumed to be 0.
//
// # Sub-editor Function
//
// This function produces a sub-editor, whose Text field will contain only the
// sub-section of text specified. Editing the parent's Text field after the
// sub-editor has been created will have no effect on the sub-editor or any
// Editor produced from it, but note that it may have an affect on the result of
// calling Commit() on the sub-editor.
//
// The sub-editor as well as any sub-editors produced from it can be merged back
// into the original Editor by calling Commit(). Alternatively, all such
// sub-editors can be merged recursively by calling CommitAll().
//
// # Options
//
// This function is affected by the following options of the Editor it is called
// on:
//
//   - `LineSeparator` specifies what string should be used to delimit lines.
//   - `NoTrailingLineSeparators` specifies whether it should consider a final
//     instance of `LineSeparator` to be ending the prior line, or giving the
//     start of a new line. If NoTrailingLineSeparators is true, a trailing
//     LineSeparator is considered to start a new (empty) line.
func (ed Editor) LinesFrom(start int) Editor {
	return ed.Lines(start, ed.LineCount())
}

// LinesTo produces an Editor to operate on a subset of the lines in the Text.
// The returned Editor operates on lines from the first line up to but not
// including the nth line, where n is `end`.
//
// For instance, to get the first two lines of a four-line string:
//
//	ed := Edit("Line #1\nLine #2\nLine #3\nLine #4")
//	ed = ed.LinesFrom(2)
//
//	fmt.Printf("%v\n", ed.Text)  // will be "Line #3\nLine #4".
//
// The `start` may be negative, in which case it will be relative to the end of
// the text; -1 would be the index of the last line, -2 would be the index of
// the second-to-last line, etc.
//
// If `end` specifies an index that is past the end of the text, that index is
// assumed to be the end of the text (i.e. 1 greater than the index of the final
// line). If it specifies an index that is before the start of a string, it is
// assumed to be 0.
//
// # Sub-editor Function
//
// This function produces a sub-editor, whose Text field will contain only the
// sub-section of text specified. Editing the parent's Text field after the
// sub-editor has been created will have no effect on the sub-editor or any
// Editor produced from it, but note that it may have an affect on the result of
// calling Commit() on the sub-editor.
//
// The sub-editor as well as any sub-editors produced from it can be merged back
// into the original Editor by calling Commit(). Alternatively, all such
// sub-editors can be merged recursively by calling CommitAll().
//
// # Options
//
// This function is affected by the following options of the Editor it is called
// on:
//
//   - `LineSeparator` specifies what string should be used to delimit lines.
//   - `NoTrailingLineSeparators` specifies whether it should consider a final
//     instance of `LineSeparator` to be ending the prior line, or giving the
//     start of a new line. If NoTrailingLineSeparators is true, a trailing
//     LineSeparator is considered to start a new (empty) line.
func (ed Editor) LinesTo(end int) Editor {
	return ed.Lines(0, end)
}

// String returns the finished, fully edited string. If the Editor is a
// sub-editor, CommitAll() is called first and the Text from the resulting
// editor is returned, else Text is returned.
func (ed Editor) String() string {
	if ed.IsSubEditor() {
		ed = ed.CommitAll()
	}

	return ed.Text
}

func (ed Editor) subEd(start, end int) Editor {
	subEd := ed
	subEd.ref = &parentRef{
		parent: &ed,
		start:  start,
		end:    end,
	}
	subEd.Text = ed.Text[start:end]
	return subEd
}
