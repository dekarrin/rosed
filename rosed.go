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
// # Basic Usage
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
//  // Here's an example of using Overtype to replace some text:
//  ed := rosed.Edit("How are you, Miss Lalonde?")
//  ed = ed.Overtype(4, "goes it")
//  output1 := ed.String()  // will be "How goes it, Miss Lalonde?"
//
//  // Alternatively, all calls can be chained:
//  output2 := rosed.Edit("How are you, Miss Lalonde?").Overtype(4, "goes it").String()
//
// Overtype is not a particularly exciting example of this library's use, but
// the [Editor.Wrap], [Editor.InsertTwoColumns], and
// [Editor.InsertDefinitionsTable] for instance, are some examples of the more
// complex operations rosed can perform.
//
// # Options For Behavior Control
//
// Some aspects of this library are controlled using consistent options set on
// an Editor. These are stored in an [Options] struct, and can be applied to an
// Editor either by assignment directly to [Editor.Options] or by calling
// [Editor.WithOptions] to get an Editor with those options set on it.
//
// The Options struct itself supports a fluent interface for changing its
// values, using its WithX methods.
//
// Editor operations will mention which options they are affected by and how in
// their documentation comments. Alternatively, to set the options specifically
// for a particular call of an Editor operation, call the XOpts version of it.
//
//  input := "More then ever, you feel, what's the word? Housetrapped."
//  opts := Options{PreserveParagraphs: true}
//  ed := rosed.Edit(input)
//  
//  // Options can be set by giving the Editor them prior to calling an operation:
//  output := ed.WithOptions(opts).Wrap(10).String()
//
//  // ...or by calling the Opts version of the operation:
//  output := ed.WrapOpts(10, opts).String()
//
// Note that Editor.WithOptions, like all other Editor functions, treats the
// Editor it operates on as immutable. It returns an Editor that has those
// options set on it, but the Editor it is called on is itself unchanged. To
// permanently set the Options on a particular Editor, you will need to assign
// an Options struct to Editor.Options manually.
//
// # Editing Partial Sections
//
// To edit only a portion of the text in an Editor, a sub-editor can be created
// using [Editor.Chars], [Editor.CharsFrom], [Editor.CharsTo], [Editor.Lines],
// [Editor.LinesFrom], or [Editor.LinesTo]. The Editor retured from these
// operations will perform operations only on the section specified, and any
// positions or lengths in operations called on it will be relative to that
// sub-section's start and ends. Writing past the end of the sub-editor's text
// is allowed and does not affect the text outside of the subsection, nor does
// the writing before the start of the text.
//
// These changes can be rolled up to the parent text by calling [Editor.Commit].
// This will produce an Editor that consists of the full text prior to selecting
// a particular section, with all changes made to the sub-section merged in.
//
// If multiple sub-editors have been made, [Editor.CommitAll] can be used to
// apply all changes recursively and return to the Editor operating on the full
// text. Calling [Editor.String] on a sub-editor will also cause CommitAll to be
// called, making it an excellent choice to get [Editor.Text] without needing to
// be concerned with whether they are operating on a sub-editor.
//
// Note that it is possible to create a sub-editor, and then create another
// sub-editor off of the same original Editor. This may have unexpected results
// and is not intended use of sub-editors.
//
// # Negative Indexing
//
// Some Editor functions accept one or more indexes, either of characters
// (graphemes specifically, in the context of this library), lines, or other
// textual elements. The rosed library has implemented the use of negative
// indexes to make it a little easier to reference things releative to the end
// of whatever collection they are in.
//
// What this means is that for a function such as [Editor.Insert] can receive a
// negative index for its position, and it will interpret it as relative to the
// end of the text. For instance, -1 would be the last character, -2 would be
// the second-to-last character, etc. If a negative index would specify someting
// before the start of the collection, it is automatically assumed to be 0.
//
// Many functions that use this will mention this explicitly, but all functions
// that accept positions can interpret negative positions.
package rosed

import "github.com/dekarrin/rosed/internal/gem"

func _g(s string) gem.String {
	return gem.New(s)
}
