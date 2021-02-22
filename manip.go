package rosed

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/dekarrin/rosed/internal/gem"
)

// contains functions for manipulation of text. Used by Editor.

// does a wrap without considering any additional lengths. Automatically
// normalizes all runs of space characters to a single space.
//
// The returned value is a Block of all resulting lines. Trailing mode will not
// be set on the Block.
func wrap(text gem.String, width int, lineSep gem.String) block {
	if width < 2 {
		panic(fmt.Sprintf("invalid width: %v", width))
	}

	lines := block{LineSeparator: lineSep}

	// normalize string to convert all whitespace to single space char.
	text = collapseSpace(text, lineSep)
	if text.String() == "" {
		return lines
	}

	toConsume := text
	var curWord, curLine gem.String
	for i := 0; i < toConsume.Len(); i++ {
		ch := toConsume.CharAt(i)
		if ch[0] == ' ' {
			curLine = appendWordToLine(lines, curWord, curLine, width)
			curWord = gem.Zero
		} else {
			curWord = curWord.Add(gem.Char(ch))
		}
	}

	if !curWord.IsEmpty() {
		curLine = appendWordToLine(lines, curWord, curLine, width)
		curWord = gem.Zero
	}

	if !curLine.IsEmpty() {
		lines.Append(curLine)
	}

	return lines
}

// lines will be modified to add the appended line if curLine is full.
func appendWordToLine(lines block, curWord gem.String, curLine gem.String, width int) (newCurLine gem.String) {
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

func collapseSpace(text gem.String, lineSep gem.String) gem.String {
	if text == nil {
		return gem.Zero
	}

	// handle the separator but do not use the empty or nil string.
	text = _g(strings.ReplaceAll(text.String(), lineSep.String(), " "))
	for i := 0; i < text.Len(); i++ {
		if unicode.IsSpace(text.CharAt(i)[0]) {
			text = text.SetCharAt(i, []rune{' '}) // set it to actual space char
		}
	}
	collapsed := spaceCollapser.ReplaceAllString(text.String(), " ")
	collapsed = strings.TrimSpace(collapsed)
	return _g(collapsed)
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
func combineColumnBlocks(left, right block, minSpaceBetween int) block {
	if left.Len() == 0 && right.Len() == 0 {
		return block{}
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

	combined := block{}
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

// justifyLine takes the given text and attempts to justify it. No attempt is
// made to split the given line into multiple lines.
//
// If there are no spaces in the given string, it is returned centered.
// If it is longer than the desired width after collapsing spaces in it, the
// collapsed-space string is returned without further modification.
func justifyLine(text gem.String, width int) gem.String {
	// collapseSpace in a line so that it can be properly laid out
	text = collapseSpace(text, _g("\n")) // doing \n which would be whitespace-collapsed anyways

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
