// Package rosed is a library for manipulating and laying out fixed-width text.
// It assumes that all text is encoded in UTF-8, the default encoding of source
// code string literals in golang.
//
// "rosed" is pronounced as "Rose-Ed" but any pronunciation that someone could
// use that can be understood to mean this library is accepted.
//
// # Grapheme-Awareness
//
// Functions that operate on text in a way that indexes characters or otherwise
// needs to count them (say, for wrapping purposes, for instance) in rosed do
// so by their grapheme clusters, not their Unicode codepoints, runes, or bytes.
// A grapheme is a single graphical unit that a human viewer would call a single
// character, but may in fact consist of several Unicode codepoints and their
// constituent runes and bytes; one or more codepoints that make up a single
// grapheme is known as a grapheme cluster.
//
// For instance, the sequence "é" is a single grapheme, as is "e". But in the
// case of "é", it may be represented in Unicode either by the single codepoint
// U+00E9 ("Latin Small Letter E With Acute"), or by the codepoint sequence
// U+0065, U+0301, ("e" followed by "Combining Acute Accent"). In golang code,
// the later would be represented in UTF-8 strings as two runes (and whatever
// bytes it would take to make up those runes), meaning that both iteration over
// the string and `utf8.RuneCountInString` would return a higher number of runes
// than perhaps would be expected. See [UAX #15: Unicode Normalization Forms]
// for more info on the forms that a Unicode string can take based on how it
// represents graphemes.
//
// This library implements the algorithms described in [UAX #29: Unicode Text
// Segmentation] to recognize where in text the boundaries of each grapheme
// cluster are and correctly index by grapheme. This means it transparently
// handles all ways that a single "human-readable" character could be
// represented.
//
// Note that this library does not handle Unicode-normalized collation; that
// may be covered at a later time but for now it was deemed too much to
// implement for version 1.0, given the large amount of data that must be read
// and dependence on possible locale settings that would be needed.
//
// [UAX #15: Unicode Normalization Forms]: https://unicode.org/reports/tr15/
// [UAX #29: Unicode Text Segmentation]: https://unicode.org/reports/tr29/
//
// # Editing Text
//
// All editing starts with an [Editor]. The zero-value can be used to start with
// blank text, otherwise Edit() can be called to give the text to work with.
// Additionally, if desired, the Editor can have its Text property set to the
// text to operate on.
//
//	ed := rosed.Editor{}
//	ed := rosed.Edit("my text")
//
// From that point, a function can be called to modify the text in the Editor.
// Editors are considered to be immutable by all operations on them, and so
// each one will return an Editor with its Text set to the result of an
// operation.
//
//  ed := rosed.Edit("How are you, Miss Lalonde?")
//  ed.
//
//
//
// # Editing Partial Sections
//
// To edit only a portion of the text in an Editor, a sub-editor can be created
// using [Editor.Chars], [Editor.CharsFrom], [Editor.CharsTo], [Editor.Lines],
// [Editor.LinesFrom], or [Editor.LinesTo]. The Editor retured from these
// operations will perform operations only on the section specified until
// [Editor.Commit] is called on it, which will produce an Editor consisting of
// the full text prior to selecting a particular section, with all changes made
// to the selection merged in.
//
// If multiple sub-selections have been made, CommitAll() can be used to apply
// all changes in sequence and return to the Editor operating on the full text.
//
// # Getting Output
//
// Most functions return a new Editor. To get the modified text, use String().
// The Editor.Text member can be accessed directly, but this might not be the
// complete text if the Editor is a sub-editor; String() always merges all
// changes before returning the text.
//
// TODO: general note on python slic-ish indexing.
package rosed

import "github.com/dekarrin/rosed/internal/gem"

func _g(s string) gem.String {
	return gem.New(s)
}
