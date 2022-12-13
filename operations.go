package rosed

// this file contains operations performed by Editors.

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/dekarrin/rosed/internal/gem"
)

var (
	spaceCollapser = regexp.MustCompile(" +")
)

// LineOperation is a function that accpets a zero-indexed line number and the
// contents of that line and performs some operation to produce zero or more new
// lines to replace the contents of the line with.
//
// The return value for a LineOperation is a slice of lines to insert at the
// old line position. This can be used to delete the line or insert additional
// new ones; to insert, include the new lines in the returned slice in the
// proper position relative to the old line in the slice, and to delete the
// original line, a slice with len < 1 can be returned.
//
// `idx` will always be the index of the line before any transformations were
// applied; e.g. if used in [Editor.Apply], a call to a LineOperation with
// idx = 4 will always be after a call with idx = 3, regardless of the size of
// the returned slice in the prior call.
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
// `idx` will always be the index of the paragraph before any transformations
// were applied; e.g. if used in [Editor.ApplyParagraphs], a call to a
// ParagraphOperation with idx = 4 will always be after a call with idx = 3,
// regardless of the size of the returned slice in the prior call.
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

// Apply applies the given LineOperation to each line in the text. Line
// termination at the last line is transparently handled as per the options set
// on the Editor.
//
// The LineOperation should assume it will receive each line without its line
// terminator, and must assume that anything it returns will have re-adding the
// separator to it handled by the caller.
//
// This function is affected by the following [Options]:
//
//   - LineSeparator specifies what string in the source text should be used to
//     delimit lines to be passed to the LineOperation.
//   - NoTrailingLineSeparators specifies whether the function should consider a
//     final instance of LineSeparator to be ending the prior line or giving the
//     start of a new line. If NoTrailingLineSeparators is true, a trailing
//     LineSeparator is considered to start a new (empty) line; additionally,
//     the LineOperation will be called at least once for an empty string. If
//     NoTrailingLineSeparators is set to false and the Editor text is set to an
//     empty string, the LineOperation will not be called.
func (ed Editor) Apply(op LineOperation) Editor {
	ed = ed.ApplyOpts(op, ed.Options)
	return ed
}

// ApplyOpts applies the given LineOperation to each line in the text, using the
// provided options.
//
// This is identical to [Editor.Apply] but provides the ability to set Options
// for the invocation.
func (ed Editor) ApplyOpts(op LineOperation, opts Options) Editor {
	opts = opts.WithDefaults()
	lines := ed.WithOptions(opts).linesSep(opts.LineSeparator)

	applied := make([]string, 0, len(lines))

	for idx, line := range lines {
		newLines := op(idx, line)
		if len(newLines) > 0 {
			applied = append(applied, newLines...)
		}
	}

	// make sure to preserve the last line sep if it exists; it will have been
	// clobbered in call to lines() if it was.
	if !opts.NoTrailingLineSeparators && strings.HasSuffix(ed.Text, opts.LineSeparator) {
		applied = append(applied, "")
	}

	ed.Text = strings.Join(applied, opts.LineSeparator)
	return ed
}

// ApplyParagraphs applies the given ParagraphOperation to each paragraph in
// the text of the Editor.
//
// The ParagraphOperation should assume it will receive each paragraph without
// its paragraph separator, and must assume that anything it returns will have
// re-adding the separator to it handled by the caller.
//
// When the ParagraphSeparator of the Editor's options is set to a sequence that
// includes visible characters that take up horizontal space, the
// ParagraphOperation will receive the prefix and suffix of the paragraph that
// would be in the joined string due to the separator, with variables sepPrefix
// and sepSuffix. This is not intended to allow the operation to add them back
// in manually, as that is handled by the caller, but for it to perform
// book-keeping and length checks and act accordingly, such when attempting to
// output something that is intended to be aligned.
//
// Unlike with LineSeparator, a ParagraphSeparator is always considered a
// separator, not a terminator, so the affixes may vary per paragraph if the
// ParagraphSeparator has line breaks in it; in that case the first paragraph
// will have an empty prefix, the last paragraph will have an empty suffix, and
// all other paragraphs will have non-empty prefixes and suffixes.
//
// For an example of a ParagraphOperator that uses sepPrefix and sepSuffix and
// a custom ParagraphSeparator that makes them non-empty, see the example for
// [Editor.ApplyParagraphsOpts].
//
// Note that treating the paragraph separator as a splitter and not a terminator
// means that the ParagraphOperation is always called at least once, even for an
// empty editor.
//
// This function is affected by the following options:
//
//   - ParagraphSeparator specifies the string that paragraphs are split by.
func (ed Editor) ApplyParagraphs(op ParagraphOperation) Editor {
	ed = ed.ApplyParagraphsOpts(op, ed.Options)
	return ed
}

// ApplyParagraphsOpts applies the given ParagraphOperation to each paragraph in
// the text of the Editor, using the provided options.
//
// This is identical to [Editor.ApplyParagraphs] but provides the ability to set
// Options for the invocation.
func (ed Editor) ApplyParagraphsOpts(op ParagraphOperation, opts Options) Editor {
	return ed.applyGParagraphsOpts(func(idx int, para, sepPrefix, sepSuffix gem.String) []gem.String {
		return gem.Slice(op(idx, para.String(), sepPrefix.String(), sepSuffix.String()))
	}, opts)
}

// CollapseSpace converts all consecutive whitespace characters to a single
// space character.
//
// This function is affected by the following [Options]:
//
//   - LineSeparator is always considered whitespace, and will be collapsed into
//     a space regardless of the classification of the characters within it.
func (ed Editor) CollapseSpace() Editor {
	return ed.CollapseSpaceOpts(ed.Options)
}

// CollapseSpaceOpts converts all consecutive whitespace characters to a single
// space character using the provided options.
//
// This is identical to [Editor.CollapseSpace] but provides the ability to set
// Options for the invocation.
func (ed Editor) CollapseSpaceOpts(opts Options) Editor {
	opts = opts.WithDefaults()
	ed.Text = collapseSpace(_g(ed.Text), _g(opts.LineSeparator)).String()
	return ed
}

// Delete removes text from the Editor. All text after the deleted sequence is
// moved left to the starting position of the deleted sequence.
//
// This function is grapheme-aware and indexes text by human-readable
// characters, not by the bytes or runes that make it up. See the note on
// Grapheme-Awareness in the [rosed] package docs for more info.
func (ed Editor) Delete(start, end int) Editor {
	if start >= end {
		return ed
	}

	before := ed.CharsTo(start).Text
	after := ed.CharsFrom(end).Text

	ed.Text = before + after
	return ed
}

// Indent adds an indent string at the start of each line in the Editor. The
// string used for a single level of indent is determined by Editor options and
// will be applied level times. If level is 0 or less, the text will be
// unchanged.
//
// With default options set, this operation has no effect on an empty editor.
//
// This function is affected by the following [Options]:
//
//   - IndentStr is the sequence to use to indent a single level.
//   - LineSeparator is the separator that determines what each line is.
//   - NoTrailingLineSeparators alters whether LineSeparator is expected to be
//     at the end of a complete line. If this is set to true, then a
//     LineSeparator does not need to be present at the end of a complete line.
//     Any trailing line separator for a non-empty editor is then considered to
//     split the last line from a new, empty line, which will be indented. In
//     addition, the empty editor will be considered to have a single line,
//     which will be indented.
//   - ParagraphSeparator is the separator used to split paragraphs. It will
//     only have effect if PreserveParagraphs is set to true.
//   - PreserveParagraphs gives whether to respect paragraphs instead of
//     treating paragraph breaks as normal text. If set to true, the text is
//     first split into paragraphs by ParagraphSeparator, then the indent is
//     applied to each paragraph.
//
// Note that there is a behavior that occurs when NoTrailingLineSeparators is
// set to true, PreserveParagraphs is set to true, and the LineSeparator could
// be misinterpreted as part of ParagraphSeparator. In this case, the paragraph
// separation will be prioritized over the line separator, possibly in an
// unintended fashion. For instance, if ParagraphSeparator is set to "\n\n" (the
// default) and LineSeparator is set to "\n" (the default), a sequence of
// "\n\n\n" would be interpreted as the paragraph separator followed by the line
// separator. This may be fixed in a later version of this library.
func (ed Editor) Indent(level int) Editor {
	return ed.IndentOpts(level, ed.Options)
}

// IndentOpts adds an indent string at the start of each line in the Editor
// using the provided options.
//
// This is identical to [Editor.Indent] but provides the ability to set Options
// for the invocation.
func (ed Editor) IndentOpts(level int, opts Options) Editor {
	if level < 1 {
		// caller wants fewer than 1 indent. Okay, that is zero; return
		// unchanged
		return ed
	}

	indent := strings.Repeat(opts.WithDefaults().IndentStr, level)

	doIndent := func(_ int, line string) []string {
		newLine := indent + line

		// only have the one line, returne that
		return []string{newLine}
	}

	if opts.WithDefaults().PreserveParagraphs {
		doIndentPara := func(_ int, para, _, _ string) []string {
			output := Edit(para).WithOptions(opts).ApplyOpts(doIndent, opts).String()
			return []string{output}
		}
		return ed.ApplyParagraphsOpts(doIndentPara, opts)
	} else {
		return ed.ApplyOpts(doIndent, opts)
	}
}

// Insert adds a string to the text at the given position. The position is
// zero-indexed and refers to the visible characters in the text. At whatever
// position is given, the existing text is moved forward to make room for the
// new text.
//
// This function is grapheme-aware and indexes text by human-readable
// characters, not by the bytes or runes that make it up. See the note on
// Grapheme-Awareness in the [rosed] package docs for more info.
func (ed Editor) Insert(charPos int, text string) Editor {
	before := ed.CharsTo(charPos).Text
	after := ed.CharsFrom(charPos).Text

	ed.Text = before + text + after
	return ed
}

// InsertDefinitionsTable creates a table of term definitions and inserts it
// into the text of the Editor. A definitions table is a two-column table that
// puts the terms being defined on the left and their definitions on the right.
// The terms are indented by two space characters.
//
//	A sample definitions table:
//
//	  John  - Has a passion for REALLY TERRIBLE MOVIES. Likes to program
//	          computers but is NOT VERY GOOD AT IT.
//
//	  Rose  - Has a passion for RATHER OBSCURE LITERATURE. Enjoys creative
//	          writing and is SOMEWHAT SECRETIVE ABOUT IT.
//
//	  Dave  - Has a penchant for spinning out UNBELIEVABLY ILL JAMS with his
//	          TURNTABLES AND MIXING GEAR. Likes to rave about BANDS NO ONE'S
//	          EVER HEARD OF BUT HIM.
//
//	  Jade  - Has so many INTERESTS, she has trouble keeping track of them all,
//	          even with an assortment of COLORFUL REMINDERS on her fingers to
//	          help sort out everything on her mind.
//
// The character position to insert the table at is given by the pos argument.
// The definitions themselves are given as a slice of 2-tuples of strings, where
// the first item in each tuple is the term and the second item is the
// definition. If no definitions are given, or an empty slice is passed in,
// there will be no output.
//
// The complete maximum width of the table to output including the leading
// indent is given by the width argument. Note that not every line will be this
// long; wrapping will often cause them to be shorter.
//
// This function is grapheme-aware and indexes text by human-readable
// characters, not by the bytes or runes that make it up. See the note on
// Grapheme-Awareness in the [rosed] package docs for more info.
//
// This function is affected by the following [Options]:
//
//   - LineSeparator is used to separate each line of the table output.
//   - ParagraphSeparator is used to separate each term/definition pair from the
//     other definitions.
//   - NoTrailingLineSeparators sets whether to include a trailing
//     LineSeparator at the end of the table. If set to true, it will be omited,
//     otherwise the table will end with a LineSeparator.
func (ed Editor) InsertDefinitionsTable(pos int, definitions [][2]string, width int) Editor {
	return ed.InsertDefinitionsTableOpts(pos, definitions, width, ed.Options)
}

// InsertDefinitionsTableOpts creates a table of term definitions using the
// provided options and inserts it into the text of the Editor.
//
// This is identical to [Editor.InsertDefinitionsTable] but provides the ability
// to set Options for the invocation.
func (ed Editor) InsertDefinitionsTableOpts(pos int, definitions [][2]string, width int, opts Options) Editor {
	opts = opts.WithDefaults()

	const (
		termLeftTabWidth = 2
		minBetween       = 2
		definitionStart  = "- "
	)

	// first find the longest term
	longestTermLen := -1
	for _, t := range definitions {
		strLen := len([]rune(t[0]))
		if strLen > longestTermLen {
			longestTermLen = strLen
		}
	}

	leftWidth := longestTermLen + termLeftTabWidth
	rightWidth := width - leftWidth - minBetween

	fullTable := block{
		LineSeparator:     _g(opts.LineSeparator),
		TrailingSeparator: !opts.NoTrailingLineSeparators,
	}

	for _, item := range definitions {
		term := item[0]
		def := item[1]
		rightPadding := ""
		if len([]rune(term)) < longestTermLen {
			rightPadding = strings.Repeat(" ", longestTermLen-len([]rune(term)))
		}
		leftTab := strings.Repeat(" ", termLeftTabWidth)
		leftCol := block{
			Lines: []gem.String{_g(fmt.Sprintf("%s%s%s", leftTab, term, rightPadding))},
		}
		// subtract 2 from width so we can put in a left margin of "  "
		rightCol := wrap(_g(def), rightWidth-2, _g(opts.LineSeparator))
		rightCol.Apply(func(idx int, line string) []string {
			if idx == 0 {
				return []string{"- " + line}
			}
			return []string{"  " + line}
		})
		combined := combineColumnBlocks(leftCol, rightCol, minBetween)

		if fullTable.Len() > 0 && combined.Len() > 0 {
			// grab the first line and append it to last line first.
			lastLineIdx := fullTable.Len() - 1
			lastLine := fullTable.Line(lastLineIdx)
			lastLine = lastLine.Add(gem.New(opts.ParagraphSeparator)).Add(combined.Line(0))
			fullTable.Set(lastLineIdx, lastLine)
			combined.Remove(0)
		}

		if combined.Len() > 0 {
			fullTable.AppendBlock(combined)
		}
	}

	if fullTable.Len() > 0 {
		return ed.Insert(pos, fullTable.Join().String())
	} else {
		return ed
	}
}

// InsertTwoColumns builds a two-column layout of side-by-side text from two
// sequences of text and inserts it into the text of the Editor. The leftText
// and the rightText do not need any special preparation to be used as the body
// of each column, as they will be automatically wrapped to fit.
//
// This function has several parameters:
//   - pos gives the position to insert the columns at within the Editor.
//   - leftText is the text to put into the left column.
//   - rightText is the text to put into the right column.
//   - minSpaceBetween is the amount of space between the two columns at the
//     left column's widest possible length.
//   - width is how much horizontal space the two columns along with the space
//     between them should be wrapped to.
//   - leftColPercent is a number from 0.0 to 1.0 that gives how much of the
//     available width (width - minSpaceBetween) the left column should take up.
//     The right column will infer its width from what remains. If
//     leftColPercent is less than 0.0, it will be assumed to be 0.0. If greater
//     than 1.0, it will be assumed to be 1.0.
//
// The minimum width that a column can be is always 2 characters wide.
//
// If the left column ends up taking more vertical space than the right column,
// the left column will have spaces added on subsequent lines to meet with where
// the right column would have started if it had had more lines.
//
// This function is grapheme-aware and indexes text by human-readable
// characters, not by the bytes or runes that make it up. See the note on
// Grapheme-Awareness in the [rosed] package docs for more info.
//
// This function is affected by the following [Options]:
//
//   - LineSeparator is used to separate each line of the output.
//   - NoTrailingLineSeparators sets whether to include a trailing LineSeparator
//     at the end of the generated columns. If set to true, it will be omited,
//     otherwise the columns will end with a LineSeparator.
func (ed Editor) InsertTwoColumns(pos int, leftText string, rightText string, minSpaceBetween int, width int, leftColPercent float64) Editor {
	return ed.InsertTwoColumnsOpts(pos, leftText, rightText, minSpaceBetween, width, leftColPercent, ed.Options)
}

// InsertTwoColumnsOpts builds a two-column layout of side-by-side text from two
// sequences of text using the options provided and inserts it into the text of
// the Editor.
//
// This is identical to [Editor.InsertTwoColumns] but provides the ability to
// set Options for the invocation.
func (ed Editor) InsertTwoColumnsOpts(pos int, leftText string, rightText string, minSpaceBetween int, width int, leftColPercent float64, opts Options) Editor {
	if leftText == "" && rightText == "" {
		return ed
	}

	if leftColPercent <= 0.0 {
		leftColPercent = 0.0
	}
	if leftColPercent > 1.0 {
		leftColPercent = 1.0
	}

	// it unreasonable to wrap each column to anything less than 2;
	// need at least 1 char for the next in a word and 1 for a continuation
	// dash. In addition, there must be enough space for the minSpaceBetween, so
	// maxWidthTarget must be at least the sum of these lengths otherwise we
	// cannot respect it.
	minLeftColWidth := 2
	minRightColWidth := 2
	minWidth := minSpaceBetween + minLeftColWidth + minRightColWidth
	if width < minWidth {
		width = minWidth
	}

	leftColWidth := int(float64(width-minSpaceBetween) * leftColPercent)
	if leftColWidth < minLeftColWidth {
		leftColWidth = minLeftColWidth
	}
	// make sure there is still space for right col
	if (width-minSpaceBetween)-leftColWidth < minRightColWidth {
		minLeftColWidth = (width - minSpaceBetween) - minRightColWidth
	}

	// difference instead of /2 here in case leftColWidth had int truncation
	// happen during its calculation.
	rightColWidth := (width - minSpaceBetween) - leftColWidth
	if rightColWidth < minRightColWidth {
		// should never happen since minWidths are used to calc max width
		panic("rightColWidth < minRightColWidth even though it should have been accounted for in call to CombineColumns")
	}

	opts = opts.WithDefaults()
	leftColBlock := wrap(_g(leftText), leftColWidth, _g(opts.LineSeparator))
	rightColBlock := wrap(_g(rightText), rightColWidth, _g(opts.LineSeparator))

	// need to get longest left-hand line and make the space between make up for the
	// difference
	maxLeftColLineLen := 0
	for i := 0; i < len(leftColBlock.Lines); i++ {
		lineLen := leftColBlock.Line(i).Len()
		if lineLen > maxLeftColLineLen {
			maxLeftColLineLen = lineLen
		}
	}
	// if left col isnt the size it should be, add space so it is
	spaceBetween := minSpaceBetween + (leftColWidth - maxLeftColLineLen)

	combinedBlock := combineColumnBlocks(leftColBlock, rightColBlock, spaceBetween)
	combinedBlock.LineSeparator = _g(opts.LineSeparator)
	combinedBlock.TrailingSeparator = !opts.NoTrailingLineSeparators

	return ed.Insert(pos, combinedBlock.Join().String())
}

// Justify edits the whitespace in each line of the Editor's text such that all
// words are spaced approximately equally and the line as a whole spans the
// given width.
//
// This function is grapheme-aware and indexes text by human-readable
// characters, not by the bytes or runes that make it up. See the note on
// Grapheme-Awareness in the [rosed] package docs for more info.
//
// This function is affected by the following [Options]:
//
//   - LineSeparator is what sequence separates lines of input.
//   - ParagraphSeparator is the separator used to split paragraphs. It will
//     only have effect if PreserveParagraphs is set to true.
//   - PreserveParagraphs gives whether to respect paragraphs instead of simply
//     considering them text to be justified. If set to true, the text is split
//     into paragraphs by ParagraphSeparator, then the justify is applied to
//     each paragraph.
func (ed Editor) Justify(width int) Editor {
	return ed.JustifyOpts(width, ed.Options)
}

// JustifyOpts edits the whitespace in each line of the Editor's text such that
// all words are spaced approximately equally and the line as a whole spans the
// given width using the provided options.
//
// This is identical to [Editor.Justify] but provides the ability to set Options
// for the invocation.
func (ed Editor) JustifyOpts(width int, opts Options) Editor {
	opts = opts.WithDefaults()

	if opts.PreserveParagraphs {
		return ed.applyGParagraphsOpts(func(idx int, para, pre, suf gem.String) []gem.String {
			sepStart := _g(strings.Repeat("A", pre.Len()))
			sepEnd := _g(strings.Repeat("A", suf.Len()))

			bl := newBlock(sepStart.Add(para).Add(sepEnd), _g(opts.LineSeparator))
			bl.Apply(func(idx int, line string) []string {
				return []string{justifyLine(_g(line), width).String()}
			})
			text := bl.Join()

			// remove separator (if any)
			if sepEnd.Len() > 0 {
				para = text.Sub(sepStart.Len(), -sepEnd.Len())
			} else {
				para = text.Sub(sepStart.Len(), text.Len())
			}
			
			return []gem.String{para}
		}, opts)
	}

	return ed.ApplyOpts(func(idx int, line string) []string {
		return []string{justifyLine(_g(line), width).String()}
	}, opts)
}

// Overtype adds characters at the given position, writing over any that already
// exist. If the added text would extend beyond the current end of the Editor
// text, the Editor text is extended to the new size.
//
// This function is grapheme-aware and indexes text by human-readable
// characters, not by the bytes or runes that make it up. See the note on
// Grapheme-Awareness in the [rosed] package docs for more info.
func (ed Editor) Overtype(charPos int, text string) Editor {
	inboundText := gem.New(text)

	before := ed.CharsTo(charPos).Text
	after := ed.CharsFrom(charPos + inboundText.Len()).Text

	ed.Text = before + inboundText.String() + after

	return ed
}

// Wrap wraps the Editor Text to the given width. The text will have whitespace
// collapsed prior to being wrapped.
//
// If width is less than 2, it is assumed to be 2 because no meaningful wrap
// algorithm can be applied to anything smaller.
//
// This function is grapheme-aware and indexes text by human-readable
// characters, not by the bytes or runes that make it up. See the note on
// Grapheme-Awareness in the [rosed] package docs for more info.
//
// This function is affected by the following options:
//
//   - `LineSeparator` is placed at the end of each line. If any sequence of
//     LineSeparator exists in the Text prior to calling this function, it will
//     first be collapsed into a single space character as part of whitespace
//     collapsing.
//   - `ParagraphSeparator` is the separator used to split paragraphs. It will
//     only have effect if PreserveParagraphs is set to true.
//   - `PreserveParagraphs` gives whether to respect paragraphs instead of
//     simply considering them text to be wrapped. If set to true, the Text is
//     first split into paragraphs by ParagraphSeparator, then the wrap is
//     applied to each paragraph.
func (ed Editor) Wrap(width int) Editor {
	return ed.WrapOpts(width, ed.Options)
}

// WrapOpts wraps the Editor Text to the given width using the supplied options.
// The text will have whitespace collapsed prior to being wrapped.
//
// If width is less than 2, it is assumed to be 2 because no meaningful wrap
// algorithm can be applied to anything smaller.
//
// This function is grapheme-aware and indexes text by human-readable
// characters, not by the bytes or runes that make it up. See the note on
// Grapheme-Awareness in the [rosed] package docs for more info.
//
// This function is affected by the following options:
//
//   - `LineSeparator` is placed at the end of each line. If any sequence of
//     LineSeparator exists in the Text prior to calling this function, it will
//     first be collapsed into a single space character as part of whitespace
//     collapsing.
//   - `ParagraphSeparator` is the separator used to split paragraphs. It will
//     only have effect if PreserveParagraphs is set to true.
//   - `PreserveParagraphs` gives whether to respect paragraphs instead of
//     simply considering them text to be wrapped. If set to true, the Text is
//     first split into paragraphs by ParagraphSeparator, then the wrap is
//     applied to each paragraph.
func (ed Editor) WrapOpts(width int, opts Options) Editor {
	opts = opts.WithDefaults()

	if width < 2 {
		width = 2
	}

	if opts.PreserveParagraphs {
		edi := ed.applyGParagraphsOpts(func(idx int, para, sepPrefix, sepSuffix gem.String) []gem.String {
			// need to include the separator prefix/suffix if any
			sepStart := _g(strings.Repeat("", sepPrefix.Len()))
			sepEnd := _g(strings.Repeat("", sepSuffix.Len()))
			textBlock := wrap(sepStart.Add(para).Add(sepEnd), width, _g(opts.LineSeparator))
			text := textBlock.Join()
			return []gem.String{text}
		}, opts)
		return edi
	}

	textBlock := wrap(_g(ed.Text), width, _g(opts.LineSeparator))
	text := textBlock.Join()
	if strings.HasSuffix(ed.Text, opts.LineSeparator) {
		text = text.Add(_g(opts.LineSeparator))
	}

	ed.Text = text.String()
	return ed
}

func (ed Editor) applyGParagraphsOpts(op gParagraphOperation, opts Options) Editor {
	opts = opts.WithDefaults()

	// split the paragraph separator about its line separators so we can see any
	// extra chars that will be chopped off while in a preserve-mode operation
	// that messes with line separators
	var paraSepPrevSuffix, paraSepNextPrefix gem.String
	parts := strings.Split(opts.ParagraphSeparator, opts.LineSeparator)

	paraSepPrevSuffix = _g(parts[0])
	if len(parts) > 1 {
		paraSepNextPrefix = _g(parts[len(parts)-1])
	}

	paragraphs := strings.Split(ed.Text, opts.ParagraphSeparator)
	transformed := make([]string, 0, len(paragraphs))
	for idx, para := range paragraphs {
		// the first one will not have the prev
		var paraPre, paraSuf gem.String
		if idx != 0 {
			paraPre = paraSepNextPrefix
		}
		if idx != len(paragraphs)-1 {
			paraSuf = paraSepPrevSuffix
		}

		nextParas := op(idx, _g(para), paraPre, paraSuf)

		if len(nextParas) > 0 {
			transformed = append(transformed, gem.Strings(nextParas)...)
		}
	}

	ed.Text = strings.Join(transformed, opts.ParagraphSeparator)

	return ed
}
