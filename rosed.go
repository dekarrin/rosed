package rosed

import "github.com/dekarrin/rosed/internal/gem"

func _g(s string) gem.String {
	return gem.New(s)
}

// Alignment is the type of alignment to apply to text. It is used in the
// [Editor.Align] function.
type Alignment int

const (
	// None is no alignment and is the zero value of an Alignment.
	None Alignment = iota

	// Left is alignment to the left side of the text.
	Left

	// Right is alignment to the right side of the text.
	Right

	// Center is alignment to the center of the text.
	Center
)

// LineOperation is a function that accepts a zero-indexed line number and the
// contents of that line and performs some operation to produce zero or more new
// lines to replace the contents of the line with.
//
// The return value for a LineOperation is a slice of lines to insert at the
// old line position. This can be used to delete the line or insert additional
// new ones; to insert, include the new lines in the returned slice in the
// proper position relative to the old line in the slice, and to delete the
// original line, a slice with len < 1 can be returned.
//
// The parameter idx will always be the index of the line before any
// transformations were applied; e.g. if used in [Editor.Apply], a call to a
// LineOperation with idx = 4 will always be after a call with idx = 3,
// regardless of the size of the returned slice in the prior call.
type LineOperation func(idx int, line string) []string

// ParagraphOperation is a function that accepts a zero-indexed paragraph number
// and the contents of that paragraph and performs some operation to produce
// zero or more new paragraphs to replace the contents of the paragraph with.
//
// The return value for a ParagraphOperation is a slice of paragraphs to insert
// at the old paragraph position. This can be used to delete the paragraph or
// insert additional new ones; to insert, include the new paragraph in the
// returned slice in the proper position relative to the old paragraph in the
// slice; to delete the original paragraph, a slice with len < 1 can be
// returned.
//
// The parameter idx will always be the index of the paragraph before any
// transformations were applied; e.g. if used in [Editor.ApplyParagraphs], a
// call to a ParagraphOperation with idx = 4 will always be after a call with
// idx = 3, regardless of the size of the returned slice in the prior call.
//
// The paragraphs may have additional contents at the beginning and end as part
// of the currently defined ParagraphSeparator. In this case, such content that
// would come at the start of the paragraph is provided in sepPrefix, and such
// content that would come at the end of the paragraph is provied in sepSuffix.
// These values are provided for reference within a ParagraphOperation, but a
// ParagraphOperation should assume the caller of it will automatically add the
// separators (which will include the affixes) as needed to the returned
// paragraph(s).
type ParagraphOperation func(idx int, para, sepPrefix, sepSuffix string) []string

type gParagraphOperation func(idx int, para, sepPrefix, sepSuffix gem.String) []gem.String
