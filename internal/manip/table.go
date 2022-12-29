package manip

// This file contains the routines for table layout. It's in its own file
// because it is rather complicated compared to the others, requiring multiple
// calculation steps.

import (
	"strings"

	"github.com/dekarrin/rosed/internal/gem"
	"github.com/dekarrin/rosed/internal/tb"
)

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
