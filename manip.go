package rosed

import (
	"fmt"
	"strings"
	"unicode"
)

// contains functions for manipulation of text. Used by Editor.

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
		lineLen := left.CharCount(i)
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
			leftLineCharCount = left.CharCount(i)
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

// justifyLine takes the given text and attempts to justify it. No attempt is
// made to split the given line into multiple lines.
//
// If there are no spaces in the given string, it is returned centered.
// If it is longer than the desired width after collapsing spaces in it, the
// collapsed-space string is returned without further modification.
func justifyLine(text string, width int) string {
	// collapseSpace in a line so that it can be properly laid out
	text = collapseSpace(text, "\n") // doing \n which would be whitespace-collapsed anyways

	textRunes := []rune(text)
	curLength := len([]rune(textRunes))
	if curLength >= width {
		return text
	}

	splitWords := strings.Split(text, " ")
	numGaps := len(splitWords) - 1
	if numGaps < 1 {
		return text
	}
	fullList := []string{}
	for idx, word := range splitWords {
		fullList = append(fullList, word)
		if idx+1 < len(splitWords) {
			fullList = append(fullList, " ")
		}
	}

	spacesToAdd := width - curLength
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
		fullList[spaceWordIdx] = fullList[spaceWordIdx] + " "
		fromRight = !fromRight
		spaceIdx++
		if spaceIdx >= numGaps {
			spaceIdx = 0
		}
	}

	finishedWord := strings.Join(fullList, "")
	return finishedWord
}
