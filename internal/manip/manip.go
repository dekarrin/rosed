// Package manip contains manipulation primitives used by Editor functions to
// manipulate text. This is to split the operation on Editor itself from the
// actual manipulation.
//
// All functions in manip operate on block.Blocks and gem.Strings.
package manip

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"

	"github.com/dekarrin/rosed/internal/gem"
	"github.com/dekarrin/rosed/internal/tb"
)

var (
	spaceCollapser = regexp.MustCompile(" +")
)

// CollapseSpace takes all runs of space in a gem String and collapses them into
// a single space. The lineSep is considered whitespace if non-empty.
func CollapseSpace(text gem.String, lineSep gem.String) gem.String {
	// handle the separator but do not use the empty string.
	if !lineSep.IsEmpty() {
		text = _g(strings.ReplaceAll(text.String(), lineSep.String(), " "))
	}
	for i := 0; i < text.Len(); i++ {
		if unicode.IsSpace(text.CharAt(i)[0]) {
			text = text.SetCharAt(i, []rune{' '}) // set it to actual space char
		}
	}
	collapsed := spaceCollapser.ReplaceAllString(text.String(), " ")
	//collapsed = strings.TrimSpace(collapsed)
	return _g(collapsed)
}

// CombineColumnBlocks takes two separate columns and combines them into a
// single block of text. The right column will be left-aligned such that it will
// be separated by minSpaceBetween space characters at minimum from the left
// column.
//
// left and right are blocks where each Line is an already-wrapped line. The
// returned block will have the lines joined and stored in its Lines property.
//
// The left and right column blocks do not need to have the same number of lines;
// if one has more lines than the other, the returned block will have a number of
// lines equal to the greater number of lines between leftText and rightText.
//
// The returned block will not have line terminator behavior set on it; callers
// will need to handle line terminators themselves.
//
// Additionally, if the left column has more lines than the right column, note
// that the last few lines will have the center spacing inserted still. So will
// end where the right column would start if there was more of it.
func CombineColumnBlocks(left, right tb.Block, minSpaceBetween int) tb.Block {
	if left.Len() == 0 && right.Len() == 0 {
		return tb.Block{}
	}
	numLines := left.Len()
	if numLines < right.Len() {
		numLines = right.Len()
	}

	// to find how far the right column should be shifted, we need to find the
	// maximum width of the left column
	var leftColMaxWidth int

	for i := 0; i < left.Len(); i++ {
		lineLen := left.CharCount(i)
		if lineLen > leftColMaxWidth {
			leftColMaxWidth = lineLen
		}
	}

	totalCharsOnLeft := leftColMaxWidth + minSpaceBetween

	combined := tb.Block{}
	for i := 0; i < numLines; i++ {
		// first get lines from each column
		var leftLine gem.String
		var leftLineCharCount int
		var rightLine gem.String
		if i < left.Len() {
			leftLine = left.Line(i)
			leftLineCharCount = left.CharCount(i)
		}
		if i < right.Len() {
			rightLine = right.Line(i)
		}

		charsToAddToLeft := totalCharsOnLeft - leftLineCharCount
		midSpacer := strings.Repeat(" ", charsToAddToLeft)

		combined.Append(_g(fmt.Sprintf("%s%s%s", leftLine, midSpacer, rightLine)))
	}

	return combined
}

// JustifyLine takes the given text and attempts to justify it. No attempt is
// made to split the given line into multiple lines.
//
// If there are no spaces in the given string, it is returned centered.
// If it is longer than the desired width after collapsing spaces in it, the
// collapsed-space string is returned without further modification.
func JustifyLine(text gem.String, width int) gem.String {
	// collapseSpace in a line so that it can be properly laid out
	text = CollapseSpace(text, _g("\n")) // doing \n which would be whitespace-collapsed anyways

	if text.Len() >= width {
		return text
	}

	splitWords := strings.Split(text.String(), " ")
	numGaps := len(splitWords) - 1
	if numGaps < 1 {
		return text
	}
	fullList := []gem.String{}
	for idx, word := range splitWords {
		fullList = append(fullList, _g(word))
		if idx+1 < len(splitWords) {
			fullList = append(fullList, _g(" "))
		}
	}

	spacesToAdd := width - text.Len()
	spaceIdx := 0
	fromRight := false
	oddSubtractor := 1
	if numGaps%2 == 0 {
		oddSubtractor = 0
	}
	for i := 0; i < spacesToAdd; i++ {
		spaceWordIdx := (spaceIdx * 2) + 1
		if fromRight {
			spaceWordIdx = (((numGaps - oddSubtractor) - spaceIdx) * 2) + 1
		}
		fullList[spaceWordIdx] = fullList[spaceWordIdx].Add(_g(" "))
		fromRight = !fromRight
		spaceIdx++
		if spaceIdx >= numGaps {
			spaceIdx = 0
		}
	}

	finishedWord := strings.Join(gem.Strings(fullList), "")
	return _g(finishedWord)
}

// Wrap does a wrap without considering any additional lengths. Automatically
// normalizes all runs of space characters to a single space.
//
// The returned value is a Block of all resulting lines. Trailing mode will not
// be set on the Block.
func Wrap(text gem.String, width int, lineSep gem.String) tb.Block {
	if width < 2 {
		width = 2
	}

	lines := tb.Block{LineSeparator: lineSep}

	// normalize string to convert all whitespace to single space char.
	text = CollapseSpace(text, lineSep)
	if text.String() == "" {
		lines.Append(gem.Zero)
		return lines
	}

	toConsume := text
	var curWord, curLine gem.String
	for i := 0; i < toConsume.Len(); i++ {
		ch := toConsume.CharAt(i)
		if ch[0] == ' ' {
			curLine = appendWordToLine(&lines, curWord, curLine, width)
			curWord = gem.Zero
		} else {
			curWord = curWord.Add(gem.New(string(ch)))
		}
	}

	if !curWord.IsEmpty() {
		curLine = appendWordToLine(&lines, curWord, curLine, width)
	}

	if !curLine.IsEmpty() {
		lines.Append(curLine)
	}

	return lines
}

// AlignLineLeft performs a left-align. Space is added to the right to make the
// line fill the width.
func AlignLineLeft(text gem.String, width int) gem.String {
	// find first instance of non-space grapheme at start.
	startSpaces := CountLeadingWhitespace(text)

	// if there are leading spaces, split the string there
	var endingText gem.String
	if startSpaces > 0 {
		endingText = text.Sub(startSpaces, text.Len())
	} else {
		endingText = text
	}

	// and add filler space, including the extra space needed for full width
	spaceLen := 0
	extraNeeded := width - endingText.Len()
	if extraNeeded > 0 {
		spaceLen = extraNeeded
	}
	trailingSpace := gem.New(strings.Repeat(" ", spaceLen))
	text = endingText.Add(trailingSpace)

	return text
}

// AlignLineRight performs a right-align. Space is added to the left to make the
// line fill the width.
func AlignLineRight(text gem.String, width int) gem.String {
	// find first instance of non-space grapheme at end.
	endSpaces := CountTrailingWhitespace(text)

	// if there are trailing spaces, split the string there
	var startingText gem.String
	//
	// 0 -> sub(0, 1)
	// 1 -> sub(
	if endSpaces > 0 {
		startingText = text.Sub(0, -endSpaces)
	} else {
		startingText = text
	}

	// and add filler space, including the extra space needed for full width
	spaceLen := 0
	extraNeeded := width - startingText.Len()
	if extraNeeded > 0 {
		spaceLen = extraNeeded
	}
	leadingSpace := gem.New(strings.Repeat(" ", spaceLen))
	text = leadingSpace.Add(startingText)

	return text
}

// AlignLineLeft performs a center of the text. Space is added to both sides to
// make the line fill the width.
func AlignLineCenter(text gem.String, width int) gem.String {
	// find first instance of non-space grapheme at start.
	startSpaces := CountLeadingWhitespace(text)
	endSpaces := CountTrailingWhitespace(text)

	// get the text to be centered
	var midText gem.String
	if endSpaces > 0 {
		midText = text.Sub(startSpaces, -endSpaces)
	} else {
		midText = text.Sub(startSpaces, text.Len())
	}

	spaceNeeded := width - midText.Len()

	if spaceNeeded <= 0 {
		// string is already at the length or too long to do further centering
		// so just return it
		return midText
	}

	rightSpaceNeeded := spaceNeeded / 2
	leftSpaceNeeded := spaceNeeded - rightSpaceNeeded

	leftSpace := gem.New(strings.Repeat(" ", leftSpaceNeeded))
	rightSpace := gem.New(strings.Repeat(" ", rightSpaceNeeded))

	text = leftSpace.Add(midText).Add(rightSpace)

	return text
}

// CountLeadingWhitespace counts the number of ws characters at the front of the
// string.
func CountLeadingWhitespace(text gem.String) int {
	for i := 0; i < text.Len(); i++ {
		if !unicode.IsSpace(text.CharAt(i)[0]) {
			return i
		}
	}
	return 0
}

// CountTrailingWhitespace counts the number of ws characters at the end of the
// string.
func CountTrailingWhitespace(text gem.String) int {
	for i := text.Len() - 1; i >= 0; i-- {
		if !unicode.IsSpace(text.CharAt(i)[0]) {
			return text.Len() - (i + 1)
		}
	}
	return 0
}

// LayoutTable creates a table from the given slice of rows, where each row
// is a slice of column content.
//
// width is the width to make the table. If the content and border options make
// it so it already meets or exceeds this, no adjustments to table content are
// made; otherwise, cells are padded to make the full table this wide.
//
// headerBreak is whether to have the first row be offset from the others. It
// will also have centered and upper-case content.
//
// lineSep is used to separate lines of output.
//
// charSet is string with "<CORNER><VERT><HORZ>" where <CORNER> is the char to
// use for corner character, <VERT> is the char to use for the vertical char,
// and <HORZ> is the char to use for the horizontal character.
//
// border is whether to have a border
func LayoutTable(table [][]gem.String, width int, lineSep gem.String, header bool, border bool, charSet gem.String) tb.Block {
	// TODO: clean up, this function is huge and probably could be broken down
	// for readability sake even if constituent parts turns out to not be very
	// re-usable
	const minNonBorderInterColumnPadding = 2

	// sanity check table input
	if len(table) < 1 {
		return tb.New(gem.Zero, lineSep)
	}

	// find how many columns the final table will have
	colCount := 0
	for i := range table {
		if len(table[i]) > colCount {
			colCount = len(table[i])
		}
	}

	if colCount == 0 {
		// there are no columns so no table to create
		return tb.New(gem.Zero, lineSep)
	}

	// if charSet is incomplete, set it to defaults
	if charSet.Len() < 3 {
		defaultSet := gem.New("+|-")
		toAdd := defaultSet.Sub(0, 3-charSet.Len())
		charSet = charSet.Add(toAdd)
	} else if charSet.Len() > 3 {
		charSet = charSet.Sub(0, 3)
	}
	cornerChar := charSet.Sub(0, 1)
	vertChar := charSet.Sub(1, 2)
	horzChar := charSet.Sub(2, 3)

	// need to calc the length of the widest item in each column
	colContentWidths := make([]int, colCount)

	for col := 0; col < colCount; col++ {
		colContentWidths[col] = 0

		for row := range table {
			var content gem.String
			if col < len(table[row]) {
				content = table[row][col]
			}

			strLen := content.Len()

			if strLen >= colContentWidths[col] {
				colContentWidths[col] = strLen
			}
		}
	}

	// add up the column widths with padding to find how much space it takes
	// up
	colContentWithPaddingWidths := make([]int, len(colContentWidths))
	copy(colContentWithPaddingWidths, colContentWidths)

	minTableWidth := 0
	if border {
		// pre-add extra space for each min col padding (2) along with the
		// additional horz border char.
		minTableWidth = horzChar.Len()
	}

	for i := range colContentWidths {
		var minPadding int
		if border {
			minPadding = 2
		} else if i+1 < len(colContentWidths) {
			// all except last column get some padding even at the smallest size
			minPadding = minNonBorderInterColumnPadding
		}
		colContentWithPaddingWidths[i] += minPadding
		minTableWidth += colContentWithPaddingWidths[i]

		if border {
			minTableWidth += horzChar.Len()
		}
	}

	// now calculate actual target column widths (including full padding)
	colWidths := make([]int, colCount)
	// start with the min content padded widths
	copy(colWidths, colContentWithPaddingWidths)

	// find total extra space we need and divide it among all columns, but for
	// cases were it does not divide evenly, go left to right.
	// additionally, final column is excluded because it should not waste space
	// on the right margin.
	spaceToAdd := width - minTableWidth
	if spaceToAdd > 0 {
		// if we are doing border mode, extra space is shared among all columns.
		//
		// if not in border mode, extra space is shared among all columns except
		// for the last so that the right edge of the longest word in last
		// column is at the edge of the width

		numColumnsToSpace := colCount
		if !border {
			numColumnsToSpace--
		}

		spacePerColumn := spaceToAdd / numColumnsToSpace
		remSpace := spaceToAdd % numColumnsToSpace
		for i := range colWidths[:numColumnsToSpace] {
			colWidths[i] += spacePerColumn
			if i < remSpace {
				colWidths[i]++
			}
		}
	} else {
		width = minTableWidth
	}

	// now we have our table widths and can begin building the table
	tableBlock := tb.New(gem.Zero, lineSep)

	// build top border if needed
	var horzBar gem.String
	if border {
		horzBar = cornerChar
		for i := range colWidths {
			for j := 0; j < colWidths[i]; j++ {
				horzBar = horzBar.Add(horzChar)
			}
			horzBar = horzBar.Add(cornerChar)
		}

		tableBlock.Append(horzBar)
	}

	var nonBorderBreakBar gem.String
	if header && !border {
		nonBorderBreakBar = gem.Zero
		for i := 0; i < width; i++ {
			nonBorderBreakBar = nonBorderBreakBar.Add(horzChar)
		}
	}

	// layout all lines
	for row := range table {
		line := gem.Zero
		if border {
			line = vertChar
		}

		var colContent gem.String
		for col := range table[row] {
			if row == 0 && header {
				headerContent := gem.New(strings.ToUpper(table[row][col].String()))
				if border {
					colContent = AlignLineCenter(headerContent, colWidths[col])
					colContent = colContent.Add(vertChar)
				} else {
					colContent = AlignLineLeft(headerContent, colWidths[col])
				}
			} else {
				if border {
					colContent = AlignLineLeft(table[row][col], colWidths[col]-1)
					colContent = gem.New(" ").Add(colContent).Add(vertChar)
				} else {
					colContent = AlignLineLeft(table[row][col], colWidths[col])
				}
			}
			line = line.Add(colContent)
		}

		tableBlock.Append(line)

		if row == 0 && header {
			if border {
				tableBlock.Append(horzBar)
			} else {
				tableBlock.Append(nonBorderBreakBar)
			}
		}
	}

	// build bottom border if needed
	if border {
		tableBlock.Append(horzBar)
	}

	return tableBlock
}

func _g(s string) gem.String {
	return gem.New(s)
}

// appendWordToLine adds the current word to the the current line unless it
// would make the width of curLine too long, in which case curLine is put into
// the lines, curLine is reset, curWord is added to the curLine. Lines will be
// modified to add the appended line if curLine is full.
//
// TODO: this is part of a wrapping operation. Consider renaming because as is
// it does not appear to have anything to do with it.
func appendWordToLine(lines *tb.Block, curWord gem.String, curLine gem.String, width int) (newCurLine gem.String) {
	// any width less than 2 is not possible and will result in an infinite loop,
	// as at least one character is required for next in word, and one character for
	// line continuation.
	if width < 2 {
		panic(fmt.Sprintf("invalid width in call to appendWordToLine: %v", width))
	}
	//originalWord := string(curWord)
	for curWord.Len() > 0 {
		addedChars := curWord.Len()
		if curLine.Len() != 0 {
			addedChars++ // for the space
		}
		if curLine.Len()+addedChars == width {
			if curLine.Len() != 0 {
				curLine = curLine.Add(_g(" "))
			}

			curLine = curLine.Add(curWord)
			lines.Append(curLine)
			curLine = gem.Zero
			curWord = gem.Zero
		} else if curLine.Len()+addedChars > width {
			if curLine.Len() == 0 {
				curLine = curLine.Add(curWord.Sub(0, width-1))
				curLine = curLine.Add(_g("-"))
				curWord = curWord.Sub(width-1, curWord.Len())
			}
			lines.Append(curLine)
			curLine = gem.Zero
		} else {
			if curLine.Len() != 0 {
				curLine = curLine.Add(_g(" "))
			}
			curLine = curLine.Add(curWord)
			curWord = gem.Zero
		}
	}
	return curLine
}
