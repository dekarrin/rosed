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

// Chars produces an Editor to operate on a subset of the characters in the
// Editor's text. The returned Editor operates on text from the nth character up
// to (but not including) the ith character, where n is start and i is end.
//
// The start or end parameter may be negative, in which case it will be relative
// to the end of the string; -1 would be the index of the last character, -2
// would be the index of the second-to-last character, etc.
//
// If one of the parameters specifies an index that is past the end of the
// string, that index is assumed to be the end of the string. If either specify
// an index that is before the start of the string, it is assumed to be 0.
//
// If end is less than start, it is assumed to be equal to start.
//
// This function is grapheme-aware and indexes text by human-readable
// characters, not by the bytes or runes that make it up. See the note on
// Grapheme-Awareness in the [rosed] package docs for more info.
//
// This is a Sub-Editor function. See the note on [Editor] for more info.
func (ed Editor) Chars(start, end int) Editor {
	// ask gem string for the grapheme-based char positions
	indexes := gem.New(ed.Text).GraphemeIndexes()

	if start == End {
		start = len(indexes)
	}
	if end == End {
		end = len(indexes)
	}

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

			// special case for when thhe runeEnd is known to be out of range,
			// immediately stop looping
			if runeEnd >= len(ed.Text) {
				break
			}
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

// CharsFrom produces an Editor to operate on a subset of the characters in the
// Editor's text. The returned Editor operates on text from the nth character up
// to the end of the text, where n is start.
//
// Calling this function is identical to calling [Editor.Chars] with the given
// start and with end set to the end of the text.
func (ed Editor) CharsFrom(start int) Editor {
	return ed.Chars(start, len(ed.Text))
}

// CharsTo produces an Editor to operate on a subset of the characters in the
// Editor's text. The returned Editor operates on text from the first character
// up to but not including the nth character, where n is end.
//
// Calling this function is identical to calling [Editor.Chars] with the given
// end and with start set to the start of the text.
func (ed Editor) CharsTo(end int) Editor {
	return ed.Chars(0, end)
}

// Commit takes the substring that a sub-editor is operating on and merges it
// with its parent. It returns an Editor which is a copy of the current one but
// with its text set to the merged string.
//
// If the Editor is already a full-text Editor, the merge operation simply
// copies the current text since there is nothing to merge with, so calling
// Commit returns an identical copy of the Editor.
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
// with its parent recursively. Returns an Editor which is a copy of the current
// one but with its text set to the result of merging every sub-editor from the
// one CommitAll is called on up to the root Editor.
//
// If the Editor is already a full-text Editor, the merge operation simply
// copies the current text since there is nothing to merge with, so calling
// CommitAll returns an identical copy of the Editor.
func (ed Editor) CommitAll() Editor {
	// recursively do commit until we get to a full-text Editor.
	for ed.IsSubEditor() {
		ed = ed.Commit()
	}
	return ed
}

// IsSubEditor returns whether the Editor was created to edit a sub-set of the
// text in some parent editor. Calls to [Editor.Lines], [Editor.LinesFrom],
// [Editor.LinesTo], [Editor.Chars], [Editor.CharsFrom], and [Editor.CharsTo]
// will result in such an Editor.
//
// If IsSubEditor returns true, then Editor.Text may be set to an incomplete
// subset of the original text. To get the full text from a sub-editor, use
// [Editor.CommitAll] to get the root parent editor with all sub-editor changes
// applied, including the ones in this sub-editor.
//
// This is a Sub-Editor function. See the note on [Editor] for more info.
func (ed Editor) IsSubEditor() bool {
	return ed.ref != nil
}

// Lines produces an Editor to operate on a subset of the lines in the Editor's
// text. The returned Editor operates on text from the nth line up to (but not
// including) the ith line, where n is start and i is end.
//
// The start or end parameter may be negative, in which case it will be relative
// to the end of the text; -1 would be the index of the last line, -2 would be
// the index of the second-to-last line, etc.
//
// If one of the parameters specifies an index that is past the end of the
// string, that index is assumed to be the end of the string. If either specify
// an index that is before the start of the string, it is assumed to be 0.
//
// If end is less than start, it is assumed to be equal to start.
//
// This function is grapheme-aware and indexes text by human-readable
// characters, not by the bytes or runes that make it up. See the note on
// Grapheme-Awareness in the [rosed] package docs for more info.
//
// This is a Sub-Editor function. See the note on [Editor] for more info.
//
// This function is affected by the following [Options]:
//
//   - LineSeparator specifies what string should be used to delimit lines.
//   - NoTrailingLineSeparators specifies whether the function should consider a
//     trailing instance of LineSeparator to end the prior line, or to start a
//     new line. If NoTrailingLineSeparators is true, a trailing LineSeparator
//     is considered to start a new (empty) line.
func (ed Editor) Lines(start, end int) Editor {
	if ed.Text == "" {
		return ed.subEd(0, 0)
	}

	lc := ed.LineCount()

	if start == End {
		start = lc
	}
	if end == End {
		end = lc
	}

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

// LinesFrom produces an Editor to operate on a subset of the lines in the
// Editor's text. The returned Editor operates on text from the nth line up to
// the end of the text, where n is start.
//
// Calling this function is identical to calling [Editor.Lines] with the given
// start and with end set to the end of the text.
func (ed Editor) LinesFrom(start int) Editor {
	return ed.Lines(start, ed.LineCount())
}

// LinesTo produces an Editor to operate on a subset of the characters in the
// Editor's text. The returned Editor operates on text from the first line up to
// but not including the nth line, where n is end.
//
// Calling this function is identical to calling [Editor.Lines] with the given
// end and with start set to the start of the text.
func (ed Editor) LinesTo(end int) Editor {
	return ed.Lines(0, end)
}

// String returns the finished, fully edited string. If the Editor is a
// sub-editor, CommitAll() is called first and Editor.Text from the resulting
// editor is returned; otherwise, the current Editor's Text is returned.
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
