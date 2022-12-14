package rosed

import "strings"

// This file contains basic structural elements of Editor as well as functions
// on it that are neither text operations nor sub-editor splitting operations.

// Editor performs transformations on text. It is the primary way to edit text
// using the rosed package.
//
// The zero value is an Editor ready to operate on the empty string.
// Alternatively, [Edit] can be called to produce an Editor ready to operate on
// the passed-in string. There is no functional difference between a zero-value
// Editor having its Text property set manually and calling Edit with a string;
// they will produce identical Editors.
//
// All operations on an Editor treat it as immutable; calling an operation
// returns a new Editor with its Text property set to the result of the
// operation. It is valid to manually set properties of an Editor and this will
// not break the assumptions made by operations called on it.
//
// Editor has an Options member for including an [Options] instance to control
// the behavior of certain functions called on it. If left unset, functions
// called on the Editor will treat it as though it has been set with the
// equivalent of calling [Options.WithDefaults] on an empty Options struct.
//
// # Sub-Editor Functions
//
// Some Editor functions produce a sub-editor, whose Text field will contain
// only the sub-section of text specified. Editing the parent's Text field after
// the sub-editor has been created will have no effect on the sub-editor or any
// Editor produced from it, but note that it may have an affect on the result of
// calling Commit() on the sub-editor and as a result is not the inteded usage
// of sub-editors.
//
// The sub-editor can be merged back into the original Editor by calling
// [Editor.Commit] on the sub-editor. Alternatively, all such sub-editors can be
// merged recursively up to the root Editor by calling [Editor.CommitAll] on the
// sub-editor.
type Editor struct {
	// Text is the string that will be operated on.
	Text string

	// The Options that the Editor will use for the next operation. These can
	// be modified prior to the operation.
	Options Options

	// the following are used only by sub-editors in order to get the full
	// string being edited (since sub-editor's own Text will only have a subset
	// of the original).
	ref *parentRef
}

// Edit creates an Editor with its Text property set to the given string and
// with Options set to their default values.
func Edit(text string) Editor {
	return Editor{
		Text: text,
	}
}

// LineCount returns the number of lines in the Editor's text. Lines are
// considered to be split by the currently set LineSeparator in the Editor's
// Options property; if one has not yet been set, [DefaultLineSeparator] is
// used.
//
// By default, an Editor whose text is set to the empty string will cause this
// function to return 0. This can be altered with the options set on the Editor.
//
// This function is affected by the following [Options]:
//
//   - LineSeparator is used to determine what splits lines to be counted.
//   - NoTrailingLineSeparators sets whether a trailing LineSeparator should be
//     expected in a full line. If set to true, it will consider the empty
//     string to be a non-terminated line as opposed to 0 lines.
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
// To get a set of Options that are identical to the current ones but with a
// single item changed, get the Editor's current Options value and call one of
// the Options.WithX functions on it.
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
