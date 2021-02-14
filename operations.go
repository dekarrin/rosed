package rosed

// this file contains operations performed by Editors.

import "strings"

// LineOperation is a function that takes a zero-indexed line number and the
// contents of that line and performs some operation on the string to get a
// new string to replace the contents of the line with.
//
// The return value for a LineOperation is a slice of lines to insert at the
// old line position. This can be used to delete the line or insert additional
// new ones; to insert, just include the new lines in the returned slice in the
// proper position relative to the old line in the slice, and to delete the
// original line, a slice with len < 1 can be returned.
//
// The idx will always be the index of the line before any transformations were
// applied; i.e. if used in ForEachLine, a call to a LineOperation with idx = 4
// will always be after a call with idx = 3, regardless of the size of the
// returned slice in the prior call.
type LineOperation func(idx int, line string) []string

// Apply applies the given LineOperation for each line in the text. Line
// termination at the last line is transparently handled as per the options
// currently set on the Editor.
func (ed Editor) Apply(op LineOperation) Editor {
	return ed.ApplyOpts(op, ed.Options)
}

// ApplyOpts applies the given LineOperation for each line in the text. Line
// termination at the last line is transparently handled as per the provided
// options.
func (ed Editor) ApplyOpts(op LineOperation, opts Options) Editor {
	opts = opts.WithDefaults()

	lines := ed.linesSep(opts.LineSeparator)
	applied := make([]string, 0, len(lines))

	for idx, line := range lines {
		newLines := op(idx, line)
		if len(newLines) > 0 {
			applied = append(applied, newLines...)
		}
	}

	// make sure to preserve the last line sep if it exists; it will have been
	// clobbered in call to lines() if it was.
	if strings.HasSuffix(ed.Text, opts.LineSeparator) {
		applied = append(applied, "")
	}

	ed.Text = strings.Join(applied, opts.LineSeparator)
	return ed
}

// Indent adds the currently configured indent string level times at the start
// of each line in the Editor. If level is 0 or less, the text is unchanged.
func (ed Editor) Indent(level int) Editor {
	return ed.IndentOpts(level, ed.Options)
}

// IndentOpts adds an indent string level times at the start of each line in the
// Editor. If level is 0 or less, the text is unchanged.
//
// The provided Options object is used to override the options currently set on
// the Editor for the indent. LineSeparator, IndentStr, and
// NoTrailingLineSeparators are read from the provided Options obejct.
func (ed Editor) IndentOpts(level int, opts Options) Editor {
	if level < 1 {
		// caller wants fewer than 1 indent. Okay, that is zero; return
		// unchanged
		return ed
	}

	indent := opts.WithDefaults().IndentStr
	doIndent := func(_ int, line string) []string {
		newLine := indent + line

		// only have the one line, returne that
		return []string{newLine}
	}

	return ed.ApplyOpts(doIndent, opts)
}
