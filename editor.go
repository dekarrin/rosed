package rosed

import "strings"

// This file contains basic structural elements of Editor as well as functions
// on it that are neither text operations nor sub-editor splitting operations.

// Editor works with text. The zero value is an editor ready to operate on the
// empty string; Edit can also be called which gives an Editor ready to operate
// on the passed-in string.
//
// All Editors are immutable; calling an operation returns a new Editor with a
// Text set to the result of the operation.
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

// WithOptions returns an Editor identical to the current one but with its
// Options replaced by the given Options.
//
// This function is no different than just setting ed.Options but provides a way
// to do so in a fluent convention.
//
// To get a set of Options that are identical to the current but with a single
// item changed, get the Editor's current options and call one of its WithX
// functions.
//
// Example:
//
// ed.WithOptions(ed.Options.WithIndentStr(" ")))
func (ed Editor) WithOptions(opts Options) Editor {
	ed.Options = opts
	return ed
}

// LineCount returns the number of lines in the Text as per the line separator.
func (ed Editor) LineCount() int {
	return len(ed.lines())
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
