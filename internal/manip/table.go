package manip

// This file contains the routines for table layout. It's in its own file
// because it is rather complicated compared to the others, requiring multiple
// calculation steps.

import (
	"strings"

	"github.com/dekarrin/rosed/internal/gem"
	"github.com/dekarrin/rosed/internal/tb"
)

type tableCharSet struct {
	corner gem.String
	vert   gem.String
	horz   gem.String
}

// MakeTable creates a table from the given slice of rows, where each row
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
func MakeTable(data [][]gem.String, width int, lineSep gem.String, header bool, border bool, charSet gem.String) tb.Block {
	const minNonBorderInterColumnPadding = 2

	// sanity check table input
	if len(data) < 1 {
		return tb.New(gem.Zero, lineSep)
	}

	// find how many columns the final table will have
	colCount := 0
	for i := range data {
		if len(data[i]) > colCount {
			colCount = len(data[i])
		}
	}

	if colCount == 0 {
		// there are no columns so no table to create
		return tb.New(gem.Zero, lineSep)
	}

	// if charSet is incomplete, set it to defaults
	tableChars := parseTableCharSet(charSet)

	// need to calc the length of the widest item in each column
	colContentWidths := make([]int, colCount)

	for col := 0; col < colCount; col++ {
		colContentWidths[col] = 0

		for row := range data {
			var content gem.String
			if col < len(data[row]) {
				content = data[row][col]
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
		minTableWidth = tableChars.horz.Len()
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
			minTableWidth += tableChars.horz.Len()
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
		//
		// the exception to the above is the single-column table; in that case,
		// all space is added to the first column.

		numColumnsToSpace := colCount
		if !border && colCount > 1 {
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
	tableBlock := buildTable(data, colWidths, width, lineSep, header, border, tableChars)

	return tableBlock
}

func parseTableCharSet(charSet gem.String) tableCharSet {
	if charSet.Len() < 3 {
		defaultSet := gem.New("+|-")
		toAdd := defaultSet.Sub(0, 3-charSet.Len())
		charSet = charSet.Add(toAdd)
	} else if charSet.Len() > 3 {
		charSet = charSet.Sub(0, 3)
	}
	return tableCharSet{
		corner: charSet.Sub(0, 1),
		vert:   charSet.Sub(1, 2),
		horz:   charSet.Sub(2, 3),
	}
}

func buildTable(data [][]gem.String, colWidths []int, width int, lineSep gem.String, header bool, border bool, chars tableCharSet) tb.Block {
	tableBlock := tb.New(gem.Zero, lineSep)

	// build top border if needed
	var horzBar gem.String
	if border {
		horzBar = chars.corner
		for i := range colWidths {
			for j := 0; j < colWidths[i]; j++ {
				horzBar = horzBar.Add(chars.horz)
			}
			horzBar = horzBar.Add(chars.corner)
		}

		tableBlock.Append(horzBar)
	}

	var nonBorderBreakBar gem.String
	if header && !border {
		nonBorderBreakBar = gem.Zero
		for i := 0; i < width; i++ {
			nonBorderBreakBar = nonBorderBreakBar.Add(chars.horz)
		}
	}

	// layout all lines
	for row := range data {
		line := gem.Zero
		if border {
			line = chars.vert
		}

		var cellContent gem.String
		for col := 0; col < len(colWidths); col++ {
			var cellData gem.String
			if col < len(data[row]) {
				cellData = data[row][col]
			}

			if row == 0 && header {
				headerContent := gem.New(strings.ToUpper(cellData.String()))
				if border {
					cellContent = AlignLineCenter(headerContent, colWidths[col])
					cellContent = cellContent.Add(chars.vert)
				} else {
					cellContent = AlignLineLeft(headerContent, colWidths[col])
				}
			} else {
				if border {
					cellContent = AlignLineLeft(cellData, colWidths[col]-1)
					cellContent = gem.New(" ").Add(cellContent).Add(chars.vert)
				} else {
					cellContent = AlignLineLeft(cellData, colWidths[col])
				}
			}
			line = line.Add(cellContent)
		}

		tableBlock.Append(line)

		if row == 0 && header {
			if border {
				// do this in a bordered table ONLY if there are more elements
				if len(data) > 1 {
					tableBlock.Append(horzBar)
				}
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
