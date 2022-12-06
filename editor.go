package rosed

import "strings"

// This file contains basic structural elements of Editor as well as functions
// on it that are neither text operations nor sub-editor splitting operations.

// Editor contains text that is being operated on and provides several operations
// that can be applied. It is the primary way to edit text using the rosed package.
//
// The zero value is an editor ready to operate on the empty string; alternatively,
// Edit() can be called to produce an Editor ready to operate on the passed-in
// string. There is no functional difference between a zero-value Editor having its
// Text property set manually to a string and calling Edit() with a string; they will
// produce identical Editors.
//
// All operations on an Editor treat the Editor as immutable; calling an operation
// returns a new Editor with its Text property set to the result of the operation.
// It is valid to manually set properties of an Editor and this will not break the
// assumptions made by operations called on it.
type Editor struct {
	// Text is the string that will be operated on.
	Text string

	// The Options that the editor will use for the next operation. These can
	// be modified prior to the operation.
	Options Options

	// the following are used only by sub-editors in order to get the full
	// string being edited (since sub-editor's own Text will only have a subset
	// of the original).
	ref *parentRef
}

// Edit creates an Editor with its Text set to the given string and with Options
// set to their default values.
func Edit(text string) Editor {
	return Editor{
		Text: text,
	}
}

// LineCount returns the number of lines in the Text. Lines are split by the currently
// set LineSeparator in the Editor's Options property; if one has not yet been set,
// DefaultLineSeparator is used.
//
// Note that this will return the number of actual lines. If the Editor has
// NoTrailingLineSeparators set to true in its options, it will consider the empty string
// to in fact be a non-terminated line as opposed to 0 lines.
func (ed Editor) LineCount() int {
	return len(ed.lines())
}

// WithOptions returns an Editor identical to the current one but with its
// Options replaced by the given Options. This does not affect the Editor
// it was called on.
//
// This function has the same effect as manually setting Editor.Options but
// provides a way to do so in a fluent convention.
//
// To get a set of Options that are identical to the current but with a single
// item changed, get the Editor's current options and call one of its WithX
// functions.
//
// Example:
//
// ed = ed.WithOptions(ed.Options.WithIndentStr(" ")))
func (ed Editor) WithOptions(opts Options) Editor {
	ed.Options = opts
	return ed
}

func (ed Editor) lines() []string {
	return ed.linesSep(ed.Options.WithDefaults().LineSeparator)
}

// automatically applies line separator policy and gets all the lines that are
// split by the given separator string.
func (ed Editor) linesSep(sep string) []string {
	lines := strings.Split(ed.Text, sep)

	// unless we have notrailinglineseparators set, consider a final line that is the
	// empty string to not be a line at all (due to no trailing sep), and thus remove
	// it from the returned lines.
	if len(lines) > 0 && !ed.Options.NoTrailingLineSeparators && lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}
	return lines
}
