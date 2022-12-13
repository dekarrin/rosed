// Package rosed is a library for manipulating and laying out fixed-width text.
// It assumes that all text is encoded in UTF-8, the default encoding of source
// code string literals in golang.
//
// "rosed" is pronounced as "Rose-Ed" but any pronunciation that can be
// understood to mean this library is accepted.
//
// # Grapheme-Awareness
//
// Functions in rosed that operate on text in a way that indexes characters or
// otherwise needs to count them (for instance, for wrapping purposes) do so by
// their grapheme clusters, not their Unicode codepoints, runes, or bytes. A
// grapheme is a single graphical unit that a human viewer would call a single
// character, but may in fact consist of several Unicode codepoints and their
// constituent runes and bytes; one or more codepoints that make up a single
// grapheme is known as a grapheme cluster.
//
// For example, the sequence "é" is a single grapheme, as is "e". But in the
// case of "é", it may be represented in Unicode either by the single codepoint
// U+00E9 ("Latin Small Letter E With Acute"), or by the codepoint sequence
// U+0065, U+0301 ("Latin Small Letter E" followed by "Combining Acute Accent").
// In golang source code, the later would be represented in UTF-8 strings as two
// runes (and whatever bytes it would take to make up those runes), meaning that
// both iteration over the string and [unicode/utf8.RuneCountInString] would
// return a higher number of runes than perhaps would be expected.
//
//	import (
//	    "fmt"
//	    "unicode/utf"
//	)
//
//	func UnicodeTest() {
//	    // This word appears to be 7 characters long:
//	    precomposed := "fiancée"
//	    decomposed := "fiance\u0301e"
//
//	    // and in fact, if printed, both show the same sequence to a human
//	    // user:
//	    fmt.Println(precomposed)  // shows "fiancée"
//	    fmt.Println(decomposed)   // ALSO shows "fiancée"
//
//	    fmt.Println(utf8.RuneCountInString(precomposed))  // prints 7
//	    fmt.Println(utf8.RuneCountInString(decomposed))   // prints 8 (?!)
//	}
//
// See [UAX #15: Unicode Normalization Forms] for more info on the forms that a
// Unicode string can take based on how it represents graphemes.
//
// This library implements the algorithms described in [UAX #29: Unicode Text
// Segmentation] to recognize where in text the boundaries of each grapheme
// cluster are and correctly index by grapheme. This means it transparently
// handles all ways that a single "human-readable" character could be
// represented.
//
//	import (
//	    "fmt"
//
//	    "github.com/dekarrin/rosed"
//	)
//
//	wrapped1 := rosed.Edit("My fiancée and I went to the bistro").Wrap(10).String()
//	wrapped2 := rosed.Edit("My fiance\u0301e and I went to the bistro").Wrap(10).String()
//
//	// because rosed is grapheme aware, both representations are wrapped the
//	// same way:
//
//	fmt.Println(wrapped1)
//	// Prints out:
//	//
//	// My fiancée
//	// and I went
//	// to the
//	// bistro.
//
//	fmt.Println(wrapped2)
//	// Prints out:
//	//
//	// My fiancée
//	// and I went
//	// to the
//	// bistro.
//
// Note that this library does not handle Unicode-normalized collation; that
// may be covered at a later time but for now it was deemed too much to
// implement for version 1.0, given the large amount of data from Unicode that
// must be added to the program for it to function and possible dependence on
// locale-specific settings.
//
// # Basic Usage
//
// All editing starts with an [Editor]. The zero-value can be used to start with
// blank text, otherwise Edit() can be called to give the text to work with.
// Additionally, if desired, the Editor can have its Text property set to the
// text to operate on.
//
//	ed1 := rosed.Editor{}
//	ed2 := rosed.Edit("my text")
//	ed3 := rosed.Editor{Text: "my text"}
//
// From that point, Editor functions can be called to modify the text in it.
// Editors are considered to be immutable by all functions that operate on them;
// these functions will return a new Editor with its Text set to the result of
// an operation as opposed to actually modifying the Text of the Editor they are
// called on.
//
//	// Here's an example of using Overtype to replace some text:
//	ed := rosed.Edit("How are you, Miss Lalonde?")
//	overtypedEd := ed.Overtype(4, "goes it")
//
//	// The original Editor is unchanged; the text
//	// will be "How are you, Miss Lalonde?":
//	originalText := ed.String()
//
//	// But the text in the Editor returned by Overtype
//	// will be "How goes it, Miss Lalonde?":
//	output1 := overtypedEd.String()
//
//	// Alternatively, all calls can be chained:
//	output2 := rosed.Edit("How are you, Miss Lalonde?").Overtype(4, "goes it").String()
//
// [Editor.Overtype] is not a particularly exciting example of this library's
// use; [Editor.Wrap], [Editor.InsertTwoColumns], and
// [Editor.InsertDefinitionsTable] are some examples of the more complex
// operations that can be performed on an Editor.
//
// # Options For Behavior Control
//
// Some aspects of this library are controlled using consistent options set on
// an Editor. These are stored in an [Options] struct, and can be applied to an
// Editor either by assignment directly to Editor.Options or by calling
// [Editor.WithOptions] to get an Editor with those options set on it.
//
// The Options struct itself supports a fluent interface for changing its
// values, using its WithX methods.
//
//	// can create Options struct with members set directly:
//	opts1 := rosed.Options{PreserveParagraphs: true}
//
//	// or by taking the zero-value and calling WithX functions:
//	opts2 := rosed.Options{}.WithPreserveParagraphs(true)
//
// Editor operations will mention which options they are affected by and how in
// their documentation comments. Alternatively, to set the options specifically
// for a particular call of an Editor operation, call the XOpts version of it.
//
//	input := "More then ever, you feel, what's the word you're looking for? Of course. Housetrapped."
//	opts := rosed.Options{PreserveParagraphs: true}
//	ed := rosed.Edit(input)
//
//	// Options can be set by giving them to the Editor prior to calling an operation:
//	output := ed.WithOptions(opts).Wrap(10).String()
//
//	// ...or by calling the Opts version of the operation:
//	output := ed.WrapOpts(10, opts).String()
//
// Note that Editor.WithOptions, like all other Editor functions, treats the
// Editor it operates on as immutable. It returns an Editor that has those
// options set on it, but the Editor it is called on is itself unchanged. To
// permanently set the Options on a particular Editor, you can assign an Options
// struct to Editor.Options manually.
//
//	// this is perfectly acceptable:
//	ed := rosed.Editor{
//	    Options: rosed.Options{
//	        LineSeparator: "\r\n",
//	    },
//	}
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
//	ed := rosed.Edit("Hello, World!")
//
//	// edits only the ", " part
//	subEd := ed.Chars(5, 7)
//
//	// Make subEd.Text be "\t, "
//	subEd = subEd.Indent(1)
//
//	// And bring it all back together
//	ed = subEd.Commit()
//
//	// Get the result string, which will be "Hello\t, world!"
//	output1 := ed.String()
//
//	// The same steps using the fluent style:
//	output2 := rosed.Edit("Hello, World!").Chars(5, 7).Indent(1).Commit().String()
//
// If multiple sub-editors have been made, [Editor.CommitAll] can be used to
// apply all changes recursively and return an Editor operating on the full
// text. Calling [Editor.String] on a sub-editor will also cause CommitAll to be
// called, making it an excellent choice to get Editor.Text without needing to
// be concerned with whether they are operating on a sub-editor.
//
//	// Commit() or CommitAll() is not needed before string:
//	output := rosed.Edit("Hello, World!").Chars(5, 7).Indent(1).String()
//
// Note that it is possible to create a sub-editor, and then create another
// sub-editor off of the same original Editor. This may have unexpected results
// and is not the intended use of sub-editors.
//
// # Negative Indexing
//
// Some Editor functions accept one or more indexes, either of characters
// (graphemes specifically, in the context of this library), lines, or other
// textual elements. The rosed library has implemented the use of negative
// indexes to make it a easier to reference positions releative to the end of
// the sequence they are in.
//
// What this means is that a function such as [Editor.Insert] can receive a
// negative index for its position argument, and it will interpret it as
// relative to the end of the text. For instance, -1 would be the last
// character, -2 would be the second-to-last character, etc. If a negative index
// would specify someting before the start of the collection, it is
// automatically assumed to be 0.
//
//	ed := rosed.Edit("John, Dave, and Jade")
//
//	// whoops, we forget to add someone 9 characters before the end!
//	ed = ed.Insert(-9, "Rose, ")
//
//	// this will be the correct "John, Dave, Rose, and Jade"
//	output := ed.String()
//
// Many functions that do this will mention it explicitly, but all functions
// that accept positions can interpret negative positions.
//
// # Custom Operations
//
// The rosed package supports several operations on Editor text out of the box,
// however these are insufficient for all uses. To address cases where custom
// functionality may be needed, two functions are provided. They are
// [Editor.Apply] and [Editor.ApplyParagraphs]. These allow the user to provide
// a custom function to operate on text on a per-line or per-paragraph basis.
//
//	textBody := "John\nRose\nDave\nJade"
//
//	namerFunc := func(lineIdx int, line string) []string {
//	    newStr := fmt.Sprintf("Beta Kid #%d: %s", lineIdx + 1, line)
//	    return []string{newStr}
//	}
//
//	output := rosed.Edit(textBody).Apply(namerFunc).String()
//	// The output will be:
//	//
//	// Beta Kid #1: John
//	// Beta Kid #2: Rose
//	// Beta Kid #3: Dave
//	// Beta Kid #4: Jade
//
// Do note that as the [LineOperation] or [ParagraphOperation] passed to the
// above functions will be user-defined, they will not be grapheme-aware unless
// the user ensures that this is the case.
//
// [UAX #15: Unicode Normalization Forms]: https://unicode.org/reports/tr15/
// [UAX #29: Unicode Text Segmentation]: https://unicode.org/reports/tr29/
package rosed
