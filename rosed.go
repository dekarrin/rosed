// Package rosed is a quick-and-dirty library for manipulating and laying out
// text.
//
// It assumes that all text is UTF-8 or a format that UTF-8 is compatible with.
//
// All editing starts with an Editor. The zero-value can be used to start with
// blank text, otherwise Edit() can be called to give the text to work with.
// Additionally, if desired, the Editor can have its Text property set to the
// text to operate on.
//
//   ed := Editor{}
//
//   ed := Edit("my text")
//
// Editing Partial Sections:
//
// To edit only a portion of the text in an Editor, a sub-editor can be created
// using Chars(), Lines(), or similar functions. The Editor retured from these
// operations will perform operations only on the section specified until
// Commit() is called, which will produce an Editor consisting of the full text
// prior to selecting a particular section, with all changes made to the
// selection merged in.
//
// If multiple sub-selections have been made, CommitAll() can be used to apply
// all changes in sequence and return to the Editor operating on the full text.
//
// Getting Output
//
// Most functions return a new Editor. To get the modified text, use String().
// The Editor.Text member can be accessed directly, but this might not be the
// complete text if the Editor is a sub-editor; String() always merges all
// changes before returning the text.
package rosed

import "github.com/dekarrin/rosed/internal/gem"

func _g(s string) gem.String {
	return gem.New(s)
}
