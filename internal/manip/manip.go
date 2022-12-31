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
		text = gem.New(strings.ReplaceAll(text.String(), lineSep.String(), " "))
	}
	for i := 0; i < text.Len(); i++ {
		if unicode.IsSpace(text.CharAt(i)[0]) {
			text = text.SetCharAt(i, []rune{' '}) // set it to actual space char
		}
	}
	collapsed := spaceCollapser.ReplaceAllString(text.String(), " ")
	return gem.New(collapsed)
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

		combined.Append(gem.New(fmt.Sprintf("%s%s%s", leftLine, midSpacer, rightLine)))
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
	text = CollapseSpace(text, gem.New("\n")) // doing \n which would be whitespace-collapsed anyways

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
		fullList = append(fullList, gem.New(word))
		if idx+1 < len(splitWords) {
			fullList = append(fullList, gem.New(" "))
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
		fullList[spaceWordIdx] = fullList[spaceWordIdx].Add(gem.New(" "))
		fromRight = !fromRight
		spaceIdx++
		if spaceIdx >= numGaps {
			spaceIdx = 0
		}
	}

	finishedWord := strings.Join(gem.Strings(fullList), "")
	return gem.New(finishedWord)
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
			curLine = appendWordToWrappedLine(&lines, curWord, curLine, width)
			curWord = gem.Zero
		} else {
			curWord = curWord.Add(gem.New(string(ch)))
		}
	}

	if !curWord.IsEmpty() {
		curLine = appendWordToWrappedLine(&lines, curWord, curLine, width)
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
	trailingSpace := gem.RepeatStr(" ", spaceLen)
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
	leadingSpace := gem.RepeatStr(" ", spaceLen)
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

	leftSpace := gem.RepeatStr(" ", leftSpaceNeeded)
	rightSpace := gem.RepeatStr(" ", rightSpaceNeeded)

	text = leftSpace.Add(midText).Add(rightSpace)

	return text
}

// CountLeadingWhitespace counts the number of ws characters at the front of the
// string.
func CountLeadingWhitespace(text gem.String) int {
	firstIndex := text.IndexFunc(func(gc []rune) bool {
		return !unicode.IsSpace(gc[0])
	})
	if firstIndex == -1 {
		return text.Len()
	}
	return firstIndex
}

// CountTrailingWhitespace counts the number of ws characters at the end of the
// string.
func CountTrailingWhitespace(text gem.String) int {
	lastIndex := text.LastIndexFunc(func(gc []rune) bool {
		return !unicode.IsSpace(gc[0])
	})
	return text.Len() - lastIndex - 1
}

// appendWordToLine adds the current word to the the current line unless it
// would make the width of curLine too long, in which case curLine is put into
// the lines, curLine is reset, curWord is added to the curLine. Lines will be
// modified to add the appended line if curLine is full.
func appendWordToWrappedLine(lines *tb.Block, curWord gem.String, curLine gem.String, width int) (newCurLine gem.String) {
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
				curLine = curLine.Add(gem.New(" "))
			}

			curLine = curLine.Add(curWord)
			lines.Append(curLine)
			curLine = gem.Zero
			curWord = gem.Zero
		} else if curLine.Len()+addedChars > width {
			if curLine.Len() == 0 {
				curLine = curLine.Add(curWord.Sub(0, width-1))
				curLine = curLine.Add(gem.New("-"))
				curWord = curWord.Sub(width-1, curWord.Len())
			}
			lines.Append(curLine)
			curLine = gem.Zero
		} else {
			if curLine.Len() != 0 {
				curLine = curLine.Add(gem.New(" "))
			}
			curLine = curLine.Add(curWord)
			curWord = gem.Zero
		}
	}
	return curLine
}
