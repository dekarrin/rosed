package rosed

// this file contains operations performed by Editors.

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"
)

var (
	spaceCollapser = regexp.MustCompile(" +")
)

// LineOperation is a function that takes a zero-indexed line number and the
// contents of that line and performs some operation on the string to get a
// new string to replace the contents of the line with.
//
// The return value for a LineOperation is a slice of lines to insert at the
// old line position. This can be used to delete the line or insert additional
// new ones; to insert, include the new lines in the returned slice in the
// proper position relative to the old line in the slice, and to delete the
// original line, a slice with len < 1 can be returned.
//
// The idx will always be the index of the line before any transformations were
// applied; e.g. if used in ForEachLine, a call to a LineOperation with idx = 4
// will always be after a call with idx = 3, regardless of the size of the
// returned slice in the prior call.
type LineOperation func(idx int, line string) []string

// ParagraphOperation is a function that takes a zero-indexed paragraph number
// and the contents of that paragraph and performs some operation on the string
// to get a new string to replace the contents of the paragraph with.
//
// The return value for a ParagraphOperation is a slice of paragraphs to insert
// at the old paragraph position. This can be used to delete the paragraph or
// insert additional new ones; to insert, include the new paragraph in the
// returned slice in the proper position relative to the old paragraph in the
// slice; to delete the original paragraph, a slice with len < 1 can be
// returned.
//
// The idx will always be the index of the paragraph before any transformations
// were applied; e.g. if used in ForEachLine, a call to a ParagraphOperation
// with idx = 4 will always be after a call with idx = 3, regardless of the size
// of the returned slice in the prior call.
//
// The paragraphs may have additional contents at the beginning and end as part
// of a the currently defined ParagraphSeparator. In this case, such content
// that would come at the start of the paragraph is provided in sepPrefix, and
// such content that would come at the end of the paragraph is provied in
// sepSuffix. Callers of the ParagraphOperation will automatically add the
// separators (which will include the affixes) as needed to the returned
// paragraph(s).
type ParagraphOperation func(idx int, para, sepPrefix, sepSuffix string) []string

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

// ApplyParagraphs applies a ParagraphOperation to the text in the Editor.
// TODO: better docs
func (ed Editor) ApplyParagraphs(op ParagraphOperation) Editor {
	return ed.ApplyParagraphsOpts(op, ed.Options)
}

// ApplyParagraphsOpts gets all the paragraphs without any paragraph separators,
// performs some operation on them, and then puts the paragraphs back together.
// TODO: Better docs
func (ed Editor) ApplyParagraphsOpts(op ParagraphOperation, opts Options) Editor {
	opts = opts.WithDefaults()

	// split the paragraph separator about its line separators so we can see any
	// extra chars that will be chopped off while in a preserve-mode wrap that
	// messes with line separators
	var paraSepPrevSuffix, paraSepNextPrefix string
	parts := strings.Split(opts.ParagraphSeparator, opts.LineSeparator)
	paraSepPrevSuffix = parts[0]
	if len(parts) > 1 {
		paraSepNextPrefix = parts[len(parts)-1]
	}

	paragraphs := strings.Split(ed.Text, opts.ParagraphSeparator)
	transformed := make([]string, 0, len(paragraphs))
	for idx, para := range paragraphs {
		// the first one will not have the prev
		var paraPre, paraSuf string
		if idx != 0 {
			paraPre = paraSepNextPrefix
		}
		if idx != len(paragraphs)-1 {
			paraSuf = paraSepPrevSuffix
		}

		nextParas := op(idx, para, paraPre, paraSuf)

		if len(nextParas) > 0 {
			transformed = append(transformed, nextParas...)
		}
	}

	ed.Text = strings.Join(transformed, opts.ParagraphSeparator)
	return ed
}

// Insert adds a string to the text at the following position. The position is
// zero-indexed and is the unicode codepoint index. The text will be inserted
// starting at this index and any content previously there will be moved up to
// make room.
func (ed Editor) Insert(charPos int, text string) Editor {
	before := ed.CharsTo(charPos).Text
	after := ed.CharsFrom(charPos).Text

	ed.Text = before + text + after
	return ed
}

/*
// InsertDefinitionsTable creates a table that gives two-columns; one for words
// on the left and the other for definitions on the right.
func InsertDefinitionsTable(definitions [][2]string, width int) string {
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

	lines := []string{}

	for _, item := range definitions {
		term := item[0]
		def := item[1]
		rightPadding := ""
		if len([]rune(term)) < longestTermLen {
			rightPadding = strings.Repeat(" ", longestTermLen-len([]rune(term)))
		}
		leftTab := strings.Repeat(" ", termLeftTabWidth)
		leftCol := []string{fmt.Sprintf("%s%s%s", leftTab, term, rightPadding)}
		rightCol := Wrap(def, rightWidth, &WrapOptions{Prefix: "  "})
		rightCol[0] = "- " + strings.TrimPrefix(rightCol[0], "  ")
		combined := CombineColumnsBlock(leftCol, rightCol, minBetween)
		lines = append(lines, combined...)
		lines = append(lines, "")
	}

	if len(lines) > 0 {
		lines = lines[:len(lines)-1] // remove trailing newline
	}

	return strings.Join(lines, "\n")
}*/

// InsertTwoColumnsOpts takes two seperate text sequences and puts them into two
// columns. Each column will be properly wrapped to fit.
//
// This function will attempt to align the columns such that the returned text
// is widthTarget large at its widest point. If the left and right columns
// cannot be wrapped such that widthTarget is achieved (for instance due to
// widthTarget being smaller than the line with longest combined length of the
// two columns plus minSpaceBetween), the lowest possible integer greater than
// widthTarget will be used.
//
// The columns will be wrapped such that the the left column will take up
// leftColPercent of the available layout area (width - space between), and the
// right column will take up the rest. If leftColPercent is less than 0.0, it
// will be assumed to be 0.0. If greater than 1.0, it will be assumed to be 1.0.
// The minimum width that a column can be is always 2 characters wide.
func (ed Editor) InsertTwoColumnsOpts(pos int, leftText string, rightText string, minSpaceBetween int, widthTarget int, leftColPercent float64, opts Options) Editor {
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
	width := widthTarget
	minLeftColWidth := 2
	minRightColWidth := 2
	minWidth := minSpaceBetween + minLeftColWidth + minRightColWidth
	if widthTarget < minWidth {
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
	leftColBlock := wrap(leftText, leftColWidth, opts.LineSeparator)
	rightColBlock := wrap(rightText, rightColWidth, opts.LineSeparator)

	combinedBlock := combineColumnBlocks(leftColBlock, rightColBlock, minSpaceBetween)
	combinedBlock.LineSeparator = opts.LineSeparator
	combinedBlock.TrailingSeparator = !opts.NoTrailingLineSeparators

	return ed.Insert(pos, combinedBlock.Join())
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

// CollapseSpace converts all runs of white space characters to a single
// space. A sequence of LineSeparator is considered whitespace regardless of the
// classification of the characters within it for the purposes of this function.
func (ed Editor) CollapseSpace() Editor {
	return ed.CollapseSpaceOpts(ed.Options)
}

// CollapseSpaceOpts converts all runs of white space characters to a single
// space. A sequence of LineSeparator is considered whitespace regardless of the
// classification of the characters within it for the purposes of this function.
func (ed Editor) CollapseSpaceOpts(opts Options) Editor {
	opts = opts.WithDefaults()
	ed.Text = collapseSpace(ed.Text, opts.LineSeparator)
	return ed
}

// Wrap performs a wrap of all text to the given width. If width is less than 2,
// it is assumed to be 2 because no meaningful wrap algorithm can be applied to
// anything smaller.
func (ed Editor) Wrap(width int) Editor {
	return ed.WrapOpts(width, ed.Options)
}

// WrapOpts performs a wrap of all text to the given width. The provided options
// are used instead of the Editor's built-in options. If width is less than 2,
// it is assumed to be 2 because no meaningful wrap algorithm can be applied to
// anything smaller.
func (ed Editor) WrapOpts(width int, opts Options) Editor {
	opts = opts.WithDefaults()

	if width < 2 {
		width = 2
	}

	if opts.PreserveParagraphs {
		return ed.ApplyParagraphsOpts(func(idx int, para, sepPrefix, sepSuffix string) []string {
			// need to include the separator prefix/suffix if any
			sepStart := ""
			sepEnd := ""
			for range sepPrefix {
				sepStart += "A"
			}
			for range sepSuffix {
				sepEnd += "A"
			}
			textBlock := wrap(sepStart+para+sepEnd, width, opts.LineSeparator)
			text := textBlock.Join()
			text = text[len(sepStart) : len(text)-len(sepEnd)]
			return []string{text}
		}, opts)
	}

	textBlock := wrap(ed.Text, width, opts.LineSeparator)
	text := textBlock.Join()
	if strings.HasSuffix(ed.Text, opts.LineSeparator) {
		text += opts.LineSeparator
	}

	ed.Text = text
	return ed
}

// does a wrap without considering any additional lengths. Automatically
// normalizes all runs of space characters to a single space.
//
// The returned value is a Block of all resulting lines. Trailing mode will not
// be set on the Block.
func wrap(text string, width int, lineSep string) Block {
	if width < 2 {
		panic(fmt.Sprintf("invalid width: %v", width))
	}

	lines := Block{LineSeparator: lineSep}

	// normalize string to convert all whitespace to single space char.
	text = collapseSpace(text, lineSep)
	if text == "" {
		return lines
	}

	toConsume := []rune(text)
	curWord := []rune{}
	curLine := []rune{}
	for i := 0; i < len(toConsume); i++ {
		ch := toConsume[i]
		if ch == ' ' {
			curLine = appendWordToLine(lines, curWord, curLine, width)
			curWord = []rune{}
		} else {
			curWord = append(curWord, ch)
		}
	}

	if len(curWord) != 0 {
		curLine = appendWordToLine(lines, curWord, curLine, width)
		curWord = []rune{}
	}

	if len(curLine) != 0 {
		lines.Append(string(curLine))
	}

	return lines
}

// lines will be modified to add the appended line if curLine is full.
func appendWordToLine(lines Block, curWord []rune, curLine []rune, width int) (newCurLine []rune) {
	// any width less than 2 is not possible and will result in an infinite loop,
	// as at least one character is required for next in word, and one character for
	// line continuation.
	if width < 2 {
		panic(fmt.Sprintf("invalid width in call to appendWordToLine: %v", width))
	}
	//originalWord := string(curWord)
	for len(curWord) > 0 {
		addedChars := len(curWord)
		if len(curLine) != 0 {
			addedChars++ // for the space
		}
		if len(curLine)+addedChars == width {
			if len(curLine) != 0 {
				curLine = append(curLine, ' ')
			}
			curLine = append(curLine, curWord...)
			lines.Append(string(curLine))
			curLine = []rune{}
			curWord = []rune{}
		} else if len(curLine)+addedChars > width {
			if len(curLine) == 0 {
				curLine = append([]rune{}, curWord[0:width-1]...)
				curLine = append(curLine, '-')
				curWord = curWord[width-1:]
			}
			lines.Append(string(curLine))
			curLine = []rune{}
		} else {
			if len(curLine) != 0 {
				curLine = append(curLine, ' ')
			}
			curLine = append(curLine, curWord...)
			curWord = []rune{}
		}
	}
	return curLine
}

func collapseSpace(text string, lineSep string) string {
	textRunes := []rune(strings.ReplaceAll(text, lineSep, " "))
	for i := 0; i < len(textRunes); i++ {
		if unicode.IsSpace(textRunes[i]) {
			textRunes[i] = ' ' // set it to actual space char
		}
	}
	text = string(textRunes)
	text = spaceCollapser.ReplaceAllString(text, " ")
	text = strings.TrimSpace(text)
	return text
}

// combineColumnBlocks takes two separate columns and combines them into a
// single block of text. The right column will be left-aligned such that it will
// be separated by minSpaceBetween space characters at minimum from the left
// column.
//
// leftText and rightText are slices where each item is a line. The returned
// slice has a similar format.
//
// The left and right column blocks do not need to be the same length; if one
// has more lines than the other, the returned block will have a number of lines
// equal to the greater number of lines between leftText and rightText. A nil
// slice is considered equivalent to a column of line length 0.
//
// The returned block will not have line terminator behavior set on it; callers
// will need to handle line terminators themselves.
func combineColumnBlocks(left, right Block, minSpaceBetween int) Block {
	if left.Len() == 0 && right.Len() == 0 {
		return Block{}
	}
	numLines := left.Len()
	if numLines < right.Len() {
		numLines = right.Len()
	}

	// to find how far the right column should be shifted, we need to find the
	// maximum width of the left column
	var leftColMaxWidth int

	for i := 0; i < left.Len(); i++ {
		lineLen := left.RuneCount(i)
		if lineLen > leftColMaxWidth {
			leftColMaxWidth = lineLen
		}
	}

	totalCharsOnLeft := leftColMaxWidth + minSpaceBetween

	combined := Block{}
	for i := 0; i < numLines; i++ {
		// first get lines from each column
		var leftLine string
		var leftLineCharCount int
		var rightLine string
		if i < left.Len() {
			leftLine = left.Line(i)
			leftLineCharCount = left.RuneCount(i)
		}
		if i < right.Len() {
			rightLine = right.Line(i)
		}

		charsToAddToLeft := totalCharsOnLeft - leftLineCharCount
		midSpacer := strings.Repeat(" ", charsToAddToLeft)

		combined.Append(fmt.Sprintf("%s%s%s", leftLine, midSpacer, rightLine))
	}

	return combined
}
