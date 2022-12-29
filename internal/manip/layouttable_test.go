package manip

import (
	"testing"

	"github.com/dekarrin/rosed/internal/gem"
	"github.com/dekarrin/rosed/internal/tb"
	"github.com/stretchr/testify/assert"
)

func Test_LayoutTable(t *testing.T) {
	testCases := []struct {
		name    string
		table   [][]gem.String
		width   int
		lineSep gem.String
		header  bool
		border  bool
		charSet gem.String
		expect  tb.Block
	}{
		{
			name:    "nil table",
			table:   nil,
			width:   10,
			lineSep: gem.New("\n"),
			header:  false,
			border:  false,
			charSet: gem.Zero,
			expect: tb.Block{
				Lines:         []gem.String{},
				LineSeparator: gem.New("\n"),
			},
		},
		{
			name:    "empty table",
			table:   [][]gem.String{},
			width:   10,
			lineSep: gem.New("\n"),
			header:  false,
			border:  false,
			charSet: gem.Zero,
			expect: tb.Block{
				Lines:         []gem.String{},
				LineSeparator: gem.New("\n"),
			},
		},
		{
			name: "table with one empty row",
			table: [][]gem.String{
				{},
			},
			width:   10,
			lineSep: gem.New("\n"),
			header:  false,
			border:  false,
			charSet: gem.Zero,
			expect: tb.Block{
				Lines:         []gem.String{},
				LineSeparator: gem.New("\n"),
			},
		},
		{
			name: "table with only empty rows",
			table: [][]gem.String{
				{},
				{},
				{},
				{},
			},
			width:   10,
			lineSep: gem.New("\n"),
			header:  false,
			border:  false,
			charSet: gem.Zero,
			expect: tb.Block{
				Lines:         []gem.String{},
				LineSeparator: gem.New("\n"),
			},
		},
		{
			name: "width < min",
			table: [][]gem.String{
				{gem.New("John"), gem.New("Egbert"), gem.New("Heir"), gem.New("Breath")},
				{gem.New("Rose"), gem.New("Lalonde"), gem.New("Seer"), gem.New("Light")},
				{gem.New("Roxy"), gem.New("Lalonde"), gem.New("Rogue"), gem.New("Void")},
				{gem.New("Vriska"), gem.New("Serket"), gem.New("Thief"), gem.New("Light")},
				{gem.New("Jack"), gem.New("Noir"), gem.New("Dog"), gem.New("Woof")},
				{gem.New("Nepeta"), gem.New("Leijon"), gem.New("Rogue"), gem.New("Heart")},
			},
			width:   2,
			lineSep: gem.New("\n"),
			header:  false,
			border:  false,
			charSet: gem.Zero,
			expect: tb.Block{
				Lines: []gem.String{
					gem.New("John    Egbert   Heir   Breath"),
					gem.New("Rose    Lalonde  Seer   Light "),
					gem.New("Roxy    Lalonde  Rogue  Void  "),
					gem.New("Vriska  Serket   Thief  Light "),
					gem.New("Jack    Noir     Dog    Woof  "),
					gem.New("Nepeta  Leijon   Rogue  Heart "),
				},
				LineSeparator: gem.New("\n"),
			},
		},
		{
			name: "width < min, single-column",
			table: [][]gem.String{
				{gem.New("John Egbert")},
				{gem.New("Rose Lalonde")},
				{gem.New("Dave Strider")},
				{gem.New("Jade Harley")},
			},
			width:   2,
			lineSep: gem.New("\n"),
			header:  false,
			border:  false,
			charSet: gem.Zero,
			expect: tb.Block{
				Lines: []gem.String{
					gem.New("John Egbert "),
					gem.New("Rose Lalonde"),
					gem.New("Dave Strider"),
					gem.New("Jade Harley "),
				},
				LineSeparator: gem.New("\n"),
			},
		},
		{
			name: "width == min",
			table: [][]gem.String{
				{gem.New("John"), gem.New("Egbert"), gem.New("Heir"), gem.New("Breath")},
				{gem.New("Rose"), gem.New("Lalonde"), gem.New("Seer"), gem.New("Light")},
				{gem.New("Roxy"), gem.New("Lalonde"), gem.New("Rogue"), gem.New("Void")},
				{gem.New("Vriska"), gem.New("Serket"), gem.New("Thief"), gem.New("Light")},
				{gem.New("Jack"), gem.New("Noir"), gem.New("Dog"), gem.New("Woof")},
				{gem.New("Nepeta"), gem.New("Leijon"), gem.New("Rogue"), gem.New("Heart")},
			},
			width:   30,
			lineSep: gem.New("\n"),
			header:  false,
			border:  false,
			charSet: gem.Zero,
			expect: tb.Block{
				Lines: []gem.String{
					gem.New("John    Egbert   Heir   Breath"),
					gem.New("Rose    Lalonde  Seer   Light "),
					gem.New("Roxy    Lalonde  Rogue  Void  "),
					gem.New("Vriska  Serket   Thief  Light "),
					gem.New("Jack    Noir     Dog    Woof  "),
					gem.New("Nepeta  Leijon   Rogue  Heart "),
				},
				LineSeparator: gem.New("\n"),
			},
		},
		{
			name: "width > min",
			table: [][]gem.String{
				{gem.New("John"), gem.New("Egbert"), gem.New("Heir"), gem.New("Breath")},
				{gem.New("Rose"), gem.New("Lalonde"), gem.New("Seer"), gem.New("Light")},
				{gem.New("Roxy"), gem.New("Lalonde"), gem.New("Rogue"), gem.New("Void")},
				{gem.New("Vriska"), gem.New("Serket"), gem.New("Thief"), gem.New("Light")},
				{gem.New("Jack"), gem.New("Noir"), gem.New("Dog"), gem.New("Woof")},
				{gem.New("Nepeta"), gem.New("Leijon"), gem.New("Rogue"), gem.New("Heart")},
			},
			width:   40,
			lineSep: gem.New("\n"),
			header:  false,
			border:  false,
			charSet: gem.Zero,
			expect: tb.Block{
				Lines: []gem.String{
					gem.New("John        Egbert      Heir      Breath"),
					gem.New("Rose        Lalonde     Seer      Light "),
					gem.New("Roxy        Lalonde     Rogue     Void  "),
					gem.New("Vriska      Serket      Thief     Light "),
					gem.New("Jack        Noir        Dog       Woof  "),
					gem.New("Nepeta      Leijon      Rogue     Heart "),
				},
				LineSeparator: gem.New("\n"),
			},
		},
		{
			name: "width > min, single-column",
			table: [][]gem.String{
				{gem.New("John Egbert")},
				{gem.New("Rose Lalonde")},
				{gem.New("Dave Strider")},
				{gem.New("Jade Harley")},
			},
			width:   15,
			lineSep: gem.New("\n"),
			header:  false,
			border:  false,
			charSet: gem.Zero,
			expect: tb.Block{
				Lines: []gem.String{
					gem.New("John Egbert    "),
					gem.New("Rose Lalonde   "),
					gem.New("Dave Strider   "),
					gem.New("Jade Harley    "),
				},
				LineSeparator: gem.New("\n"),
			},
		},
		{
			name: "with header, width < min",
			table: [][]gem.String{
				{gem.New("Name"), gem.New("Surname"), gem.New("Class"), gem.New("Aspect")},
				{gem.New("John"), gem.New("Egbert"), gem.New("Heir"), gem.New("Breath")},
				{gem.New("Rose"), gem.New("Lalonde"), gem.New("Seer"), gem.New("Light")},
				{gem.New("Roxy"), gem.New("Lalonde"), gem.New("Rogue"), gem.New("Void")},
				{gem.New("Vriska"), gem.New("Serket"), gem.New("Thief"), gem.New("Light")},
				{gem.New("Jack"), gem.New("Noir"), gem.New("Dog"), gem.New("Woof")},
				{gem.New("Nepeta"), gem.New("Leijon"), gem.New("Rogue"), gem.New("Heart")},
			},
			width:   2,
			lineSep: gem.New("\n"),
			header:  true,
			border:  false,
			charSet: gem.Zero,
			expect: tb.Block{
				Lines: []gem.String{
					gem.New("NAME    SURNAME  CLASS  ASPECT"),
					gem.New("------------------------------"),
					gem.New("John    Egbert   Heir   Breath"),
					gem.New("Rose    Lalonde  Seer   Light "),
					gem.New("Roxy    Lalonde  Rogue  Void  "),
					gem.New("Vriska  Serket   Thief  Light "),
					gem.New("Jack    Noir     Dog    Woof  "),
					gem.New("Nepeta  Leijon   Rogue  Heart "),
				},
				LineSeparator: gem.New("\n"),
			},
		},
		{
			name: "with header, width < min, single-column",
			table: [][]gem.String{
				{gem.New("Name")},
				{gem.New("John Egbert")},
				{gem.New("Rose Lalonde")},
				{gem.New("Dave Strider")},
				{gem.New("Jade Harley")},
			},
			width:   2,
			lineSep: gem.New("\n"),
			header:  true,
			border:  false,
			charSet: gem.Zero,
			expect: tb.Block{
				Lines: []gem.String{
					gem.New("NAME        "),
					gem.New("------------"),
					gem.New("John Egbert "),
					gem.New("Rose Lalonde"),
					gem.New("Dave Strider"),
					gem.New("Jade Harley "),
				},
				LineSeparator: gem.New("\n"),
			},
		},
		{
			name: "with header, width == min",
			table: [][]gem.String{
				{gem.New("Name"), gem.New("Surname"), gem.New("Class"), gem.New("Aspect")},
				{gem.New("John"), gem.New("Egbert"), gem.New("Heir"), gem.New("Breath")},
				{gem.New("Rose"), gem.New("Lalonde"), gem.New("Seer"), gem.New("Light")},
				{gem.New("Roxy"), gem.New("Lalonde"), gem.New("Rogue"), gem.New("Void")},
				{gem.New("Vriska"), gem.New("Serket"), gem.New("Thief"), gem.New("Light")},
				{gem.New("Jack"), gem.New("Noir"), gem.New("Dog"), gem.New("Woof")},
				{gem.New("Nepeta"), gem.New("Leijon"), gem.New("Rogue"), gem.New("Heart")},
			},
			width:   30,
			lineSep: gem.New("\n"),
			header:  true,
			border:  false,
			charSet: gem.Zero,
			expect: tb.Block{
				Lines: []gem.String{
					gem.New("NAME    SURNAME  CLASS  ASPECT"),
					gem.New("------------------------------"),
					gem.New("John    Egbert   Heir   Breath"),
					gem.New("Rose    Lalonde  Seer   Light "),
					gem.New("Roxy    Lalonde  Rogue  Void  "),
					gem.New("Vriska  Serket   Thief  Light "),
					gem.New("Jack    Noir     Dog    Woof  "),
					gem.New("Nepeta  Leijon   Rogue  Heart "),
				},
				LineSeparator: gem.New("\n"),
			},
		},
		{
			name: "with header, width > min",
			table: [][]gem.String{
				{gem.New("Name"), gem.New("Surname"), gem.New("Class"), gem.New("Aspect")},
				{gem.New("John"), gem.New("Egbert"), gem.New("Heir"), gem.New("Breath")},
				{gem.New("Rose"), gem.New("Lalonde"), gem.New("Seer"), gem.New("Light")},
				{gem.New("Roxy"), gem.New("Lalonde"), gem.New("Rogue"), gem.New("Void")},
				{gem.New("Vriska"), gem.New("Serket"), gem.New("Thief"), gem.New("Light")},
				{gem.New("Jack"), gem.New("Noir"), gem.New("Dog"), gem.New("Woof")},
				{gem.New("Nepeta"), gem.New("Leijon"), gem.New("Rogue"), gem.New("Heart")},
			},
			width:   50,
			lineSep: gem.New("\n"),
			header:  true,
			border:  false,
			charSet: gem.Zero,
			expect: tb.Block{
				Lines: []gem.String{
					gem.New("NAME           SURNAME         CLASS        ASPECT"),
					gem.New("--------------------------------------------------"),
					gem.New("John           Egbert          Heir         Breath"),
					gem.New("Rose           Lalonde         Seer         Light "),
					gem.New("Roxy           Lalonde         Rogue        Void  "),
					gem.New("Vriska         Serket          Thief        Light "),
					gem.New("Jack           Noir            Dog          Woof  "),
					gem.New("Nepeta         Leijon          Rogue        Heart "),
				},
				LineSeparator: gem.New("\n"),
			},
		},
		{
			name: "with header, width > min, single-column",
			table: [][]gem.String{
				{gem.New("Name")},
				{gem.New("John Egbert")},
				{gem.New("Rose Lalonde")},
				{gem.New("Dave Strider")},
				{gem.New("Jade Harley")},
			},
			width:   15,
			lineSep: gem.New("\n"),
			header:  true,
			border:  false,
			charSet: gem.Zero,
			expect: tb.Block{
				Lines: []gem.String{
					gem.New("NAME           "),
					gem.New("---------------"),
					gem.New("John Egbert    "),
					gem.New("Rose Lalonde   "),
					gem.New("Dave Strider   "),
					gem.New("Jade Harley    "),
				},
				LineSeparator: gem.New("\n"),
			},
		},
		{
			name: "with border, width < min",
			table: [][]gem.String{
				{gem.New("John"), gem.New("Egbert"), gem.New("Heir"), gem.New("Breath")},
				{gem.New("Rose"), gem.New("Lalonde"), gem.New("Seer"), gem.New("Light")},
				{gem.New("Roxy"), gem.New("Lalonde"), gem.New("Rogue"), gem.New("Void")},
				{gem.New("Vriska"), gem.New("Serket"), gem.New("Thief"), gem.New("Light")},
				{gem.New("Jack"), gem.New("Noir"), gem.New("Dog"), gem.New("Woof")},
				{gem.New("Nepeta"), gem.New("Leijon"), gem.New("Rogue"), gem.New("Heart")},
			},
			width:   2,
			lineSep: gem.New("\n"),
			header:  false,
			border:  true,
			charSet: gem.Zero,
			expect: tb.Block{
				Lines: []gem.String{
					gem.New("+--------+---------+-------+--------+"),
					gem.New("| John   | Egbert  | Heir  | Breath |"),
					gem.New("| Rose   | Lalonde | Seer  | Light  |"),
					gem.New("| Roxy   | Lalonde | Rogue | Void   |"),
					gem.New("| Vriska | Serket  | Thief | Light  |"),
					gem.New("| Jack   | Noir    | Dog   | Woof   |"),
					gem.New("| Nepeta | Leijon  | Rogue | Heart  |"),
					gem.New("+--------+---------+-------+--------+"),
				},
				LineSeparator: gem.New("\n"),
			},
		},
		{
			name: "with border, width < min, single-column",
			table: [][]gem.String{
				{gem.New("John Egbert")},
				{gem.New("Rose Lalonde")},
				{gem.New("Dave Strider")},
				{gem.New("Jade Harley")},
			},
			width:   2,
			lineSep: gem.New("\n"),
			header:  false,
			border:  true,
			charSet: gem.Zero,
			expect: tb.Block{
				Lines: []gem.String{
					gem.New(""),
					gem.New("+--------------+"),
					gem.New("| John Egbert  |"),
					gem.New("| Rose Lalonde |"),
					gem.New("| Dave Strider |"),
					gem.New("| Jade Harley  |"),
					gem.New("+--------------+"),
				},
				LineSeparator: gem.New("\n"),
			},
		},
		{
			name: "with border, width == min",
			table: [][]gem.String{
				{gem.New("John"), gem.New("Egbert"), gem.New("Heir"), gem.New("Breath")},
				{gem.New("Rose"), gem.New("Lalonde"), gem.New("Seer"), gem.New("Light")},
				{gem.New("Roxy"), gem.New("Lalonde"), gem.New("Rogue"), gem.New("Void")},
				{gem.New("Vriska"), gem.New("Serket"), gem.New("Thief"), gem.New("Light")},
				{gem.New("Jack"), gem.New("Noir"), gem.New("Dog"), gem.New("Woof")},
				{gem.New("Nepeta"), gem.New("Leijon"), gem.New("Rogue"), gem.New("Heart")},
			},
			width:   37,
			lineSep: gem.New("\n"),
			header:  false,
			border:  true,
			charSet: gem.Zero,
			expect: tb.Block{
				Lines: []gem.String{
					gem.New("+--------+---------+-------+--------+"),
					gem.New("| John   | Egbert  | Heir  | Breath |"),
					gem.New("| Rose   | Lalonde | Seer  | Light  |"),
					gem.New("| Roxy   | Lalonde | Rogue | Void   |"),
					gem.New("| Vriska | Serket  | Thief | Light  |"),
					gem.New("| Jack   | Noir    | Dog   | Woof   |"),
					gem.New("| Nepeta | Leijon  | Rogue | Heart  |"),
					gem.New("+--------+---------+-------+--------+"),
				},
				LineSeparator: gem.New("\n"),
			},
		},
		{
			name: "with border, width > min (final cell is spaced, unlike non-border)",
			table: [][]gem.String{
				{gem.New("John"), gem.New("Egbert"), gem.New("Heir"), gem.New("Breath")},
				{gem.New("Rose"), gem.New("Lalonde"), gem.New("Seer"), gem.New("Light")},
				{gem.New("Roxy"), gem.New("Lalonde"), gem.New("Rogue"), gem.New("Void")},
				{gem.New("Vriska"), gem.New("Serket"), gem.New("Thief"), gem.New("Light")},
				{gem.New("Jack"), gem.New("Noir"), gem.New("Dog"), gem.New("Woof")},
				{gem.New("Nepeta"), gem.New("Leijon"), gem.New("Rogue"), gem.New("Heart")},
			},
			width:   50,
			lineSep: gem.New("\n"),
			header:  false,
			border:  true,
			charSet: gem.Zero,
			expect: tb.Block{
				Lines: []gem.String{
					gem.New("+------------+------------+----------+-----------+"),
					gem.New("| John       | Egbert     | Heir     | Breath    |"),
					gem.New("| Rose       | Lalonde    | Seer     | Light     |"),
					gem.New("| Roxy       | Lalonde    | Rogue    | Void      |"),
					gem.New("| Vriska     | Serket     | Thief    | Light     |"),
					gem.New("| Jack       | Noir       | Dog      | Woof      |"),
					gem.New("| Nepeta     | Leijon     | Rogue    | Heart     |"),
					gem.New("+------------+------------+----------+-----------+"),
				},
				LineSeparator: gem.New("\n"),
			},
		},
		{
			name: "with border, width > min, single-column",
			table: [][]gem.String{
				{gem.New("John Egbert")},
				{gem.New("Rose Lalonde")},
				{gem.New("Dave Strider")},
				{gem.New("Jade Harley")},
			},
			width:   20,
			lineSep: gem.New("\n"),
			header:  false,
			border:  true,
			charSet: gem.Zero,
			expect: tb.Block{
				Lines: []gem.String{
					gem.New(""),
					gem.New("+------------------+"),
					gem.New("| John Egbert      |"),
					gem.New("| Rose Lalonde     |"),
					gem.New("| Dave Strider     |"),
					gem.New("| Jade Harley      |"),
					gem.New("+------------------+"),
				},
				LineSeparator: gem.New("\n"),
			},
		},
		{
			name: "with border and header, width < min",
			table: [][]gem.String{
				{gem.New("Name"), gem.New("Surname"), gem.New("Class"), gem.New("Aspect")},
				{gem.New("John"), gem.New("Egbert"), gem.New("Heir"), gem.New("Breath")},
				{gem.New("Rose"), gem.New("Lalonde"), gem.New("Seer"), gem.New("Light")},
				{gem.New("Roxy"), gem.New("Lalonde"), gem.New("Rogue"), gem.New("Void")},
				{gem.New("Vriska"), gem.New("Serket"), gem.New("Thief"), gem.New("Light")},
				{gem.New("Jack"), gem.New("Noir"), gem.New("Dog"), gem.New("Woof")},
				{gem.New("Nepeta"), gem.New("Leijon"), gem.New("Rogue"), gem.New("Heart")},
			},
			width:   2,
			lineSep: gem.New("\n"),
			header:  true,
			border:  true,
			charSet: gem.Zero,
			expect: tb.Block{
				Lines: []gem.String{
					gem.New("+--------+---------+-------+--------+"),
					gem.New("|  NAME  | SURNAME | CLASS | ASPECT |"),
					gem.New("+--------+---------+-------+--------+"),
					gem.New("| John   | Egbert  | Heir  | Breath |"),
					gem.New("| Rose   | Lalonde | Seer  | Light  |"),
					gem.New("| Roxy   | Lalonde | Rogue | Void   |"),
					gem.New("| Vriska | Serket  | Thief | Light  |"),
					gem.New("| Jack   | Noir    | Dog   | Woof   |"),
					gem.New("| Nepeta | Leijon  | Rogue | Heart  |"),
					gem.New("+--------+---------+-------+--------+"),
				},
				LineSeparator: gem.New("\n"),
			},
		},
		{
			name: "with border and header, width < min, single-column",
			table: [][]gem.String{
				{gem.New("Name")},
				{gem.New("John Egbert")},
				{gem.New("Rose Lalonde")},
				{gem.New("Dave Strider")},
				{gem.New("Jade Harley")},
			},
			width:   2,
			lineSep: gem.New("\n"),
			header:  true,
			border:  true,
			charSet: gem.Zero,
			expect: tb.Block{
				Lines: []gem.String{
					gem.New("+--------------+"),
					gem.New("|     NAME     |"),
					gem.New("+--------------+"),
					gem.New("| John Egbert  |"),
					gem.New("| Rose Lalonde |"),
					gem.New("| Dave Strider |"),
					gem.New("| Jade Harley  |"),
					gem.New("+--------------+"),
				},
				LineSeparator: gem.New("\n"),
			},
		},
		{
			name: "with border and header, width == min",
			table: [][]gem.String{
				{gem.New("Name"), gem.New("Surname"), gem.New("Class"), gem.New("Aspect")},
				{gem.New("John"), gem.New("Egbert"), gem.New("Heir"), gem.New("Breath")},
				{gem.New("Rose"), gem.New("Lalonde"), gem.New("Seer"), gem.New("Light")},
				{gem.New("Roxy"), gem.New("Lalonde"), gem.New("Rogue"), gem.New("Void")},
				{gem.New("Vriska"), gem.New("Serket"), gem.New("Thief"), gem.New("Light")},
				{gem.New("Jack"), gem.New("Noir"), gem.New("Dog"), gem.New("Woof")},
				{gem.New("Nepeta"), gem.New("Leijon"), gem.New("Rogue"), gem.New("Heart")},
			},
			width:   37,
			lineSep: gem.New("\n"),
			header:  true,
			border:  true,
			charSet: gem.Zero,
			expect: tb.Block{
				Lines: []gem.String{
					gem.New("+--------+---------+-------+--------+"),
					gem.New("|  NAME  | SURNAME | CLASS | ASPECT |"),
					gem.New("+--------+---------+-------+--------+"),
					gem.New("| John   | Egbert  | Heir  | Breath |"),
					gem.New("| Rose   | Lalonde | Seer  | Light  |"),
					gem.New("| Roxy   | Lalonde | Rogue | Void   |"),
					gem.New("| Vriska | Serket  | Thief | Light  |"),
					gem.New("| Jack   | Noir    | Dog   | Woof   |"),
					gem.New("| Nepeta | Leijon  | Rogue | Heart  |"),
					gem.New("+--------+---------+-------+--------+"),
				},
				LineSeparator: gem.New("\n"),
			},
		},
		{
			name: "with border and header, width > min",
			table: [][]gem.String{
				{gem.New("Name"), gem.New("Surname"), gem.New("Class"), gem.New("Aspect")},
				{gem.New("John"), gem.New("Egbert"), gem.New("Heir"), gem.New("Breath")},
				{gem.New("Rose"), gem.New("Lalonde"), gem.New("Seer"), gem.New("Light")},
				{gem.New("Roxy"), gem.New("Lalonde"), gem.New("Rogue"), gem.New("Void")},
				{gem.New("Vriska"), gem.New("Serket"), gem.New("Thief"), gem.New("Light")},
				{gem.New("Jack"), gem.New("Noir"), gem.New("Dog"), gem.New("Woof")},
				{gem.New("Nepeta"), gem.New("Leijon"), gem.New("Rogue"), gem.New("Heart")},
			},
			width:   52,
			lineSep: gem.New("\n"),
			header:  true,
			border:  true,
			charSet: gem.Zero,
			expect: tb.Block{
				Lines: []gem.String{
					gem.New("+------------+-------------+-----------+-----------+"),
					gem.New("|    NAME    |   SURNAME   |   CLASS   |   ASPECT  |"),
					gem.New("+------------+-------------+-----------+-----------+"),
					gem.New("| John       | Egbert      | Heir      | Breath    |"),
					gem.New("| Rose       | Lalonde     | Seer      | Light     |"),
					gem.New("| Roxy       | Lalonde     | Rogue     | Void      |"),
					gem.New("| Vriska     | Serket      | Thief     | Light     |"),
					gem.New("| Jack       | Noir        | Dog       | Woof      |"),
					gem.New("| Nepeta     | Leijon      | Rogue     | Heart     |"),
					gem.New("+------------+-------------+-----------+-----------+"),
				},
				LineSeparator: gem.New("\n"),
			},
		},
		{
			name: "with border and header, width > min, single-column",
			table: [][]gem.String{
				{gem.New("Name")},
				{gem.New("John Egbert")},
				{gem.New("Rose Lalonde")},
				{gem.New("Dave Strider")},
				{gem.New("Jade Harley")},
			},
			width:   20,
			lineSep: gem.New("\n"),
			header:  true,
			border:  true,
			charSet: gem.Zero,
			expect: tb.Block{
				Lines: []gem.String{
					gem.New("+------------------+"),
					gem.New("|       NAME       |"),
					gem.New("+------------------+"),
					gem.New("| John Egbert      |"),
					gem.New("| Rose Lalonde     |"),
					gem.New("| Dave Strider     |"),
					gem.New("| Jade Harley      |"),
					gem.New("+------------------+"),
				},
				LineSeparator: gem.New("\n"),
			},
		},
		{
			name: "missing columns",
			table: [][]gem.String{
				{gem.New("Rose"), gem.New("Lalonde")},
				{gem.New("Calliope")},
				{gem.Zero, gem.New("Strider")},
				{gem.New("Jade"), gem.New("Harley")},
			},
			width:   20,
			lineSep: gem.New("\n"),
			header:  false,
			border:  false,
			charSet: gem.Zero,
			expect: tb.Block{
				Lines: []gem.String{
					gem.New("Rose         Lalonde"),
					gem.New("Calliope            "),
					gem.New("             Strider"),
					gem.New("Jade         Harley "),
				},
				LineSeparator: gem.New("\n"),
			},
		},
		{
			name: "custom lineSep",
			table: [][]gem.String{
				{gem.New("John"), gem.New("Egbert"), gem.New("Heir"), gem.New("Breath")},
				{gem.New("Rose"), gem.New("Lalonde"), gem.New("Seer"), gem.New("Light")},
				{gem.New("Roxy"), gem.New("Lalonde"), gem.New("Rogue"), gem.New("Void")},
				{gem.New("Vriska"), gem.New("Serket"), gem.New("Thief"), gem.New("Light")},
				{gem.New("Jack"), gem.New("Noir"), gem.New("Dog"), gem.New("Woof")},
				{gem.New("Nepeta"), gem.New("Leijon"), gem.New("Rogue"), gem.New("Heart")},
			},
			width:   30,
			lineSep: gem.New("\n<P>\n"),
			header:  false,
			border:  false,
			charSet: gem.Zero,
			expect: tb.Block{
				Lines: []gem.String{
					gem.New("John    Egbert   Heir   Breath"),
					gem.New("Rose    Lalonde  Seer   Light "),
					gem.New("Roxy    Lalonde  Rogue  Void  "),
					gem.New("Vriska  Serket   Thief  Light "),
					gem.New("Jack    Noir     Dog    Woof  "),
					gem.New("Nepeta  Leijon   Rogue  Heart "),
				},
				LineSeparator: gem.New("\n<P>\n"),
			},
		},
		{
			name: "custom charSet, no border",
			table: [][]gem.String{
				{gem.New("Name"), gem.New("Surname"), gem.New("Class"), gem.New("Aspect")},
				{gem.New("John"), gem.New("Egbert"), gem.New("Heir"), gem.New("Breath")},
				{gem.New("Rose"), gem.New("Lalonde"), gem.New("Seer"), gem.New("Light")},
				{gem.New("Roxy"), gem.New("Lalonde"), gem.New("Rogue"), gem.New("Void")},
				{gem.New("Vriska"), gem.New("Serket"), gem.New("Thief"), gem.New("Light")},
				{gem.New("Jack"), gem.New("Noir"), gem.New("Dog"), gem.New("Woof")},
				{gem.New("Nepeta"), gem.New("Leijon"), gem.New("Rogue"), gem.New("Heart")},
			},
			width:   36,
			lineSep: gem.New("\n"),
			header:  true,
			border:  false,
			charSet: gem.New("+|="),
			expect: tb.Block{
				Lines: []gem.String{
					gem.New("NAME      SURNAME    CLASS    ASPECT"),
					gem.New("===================================="),
					gem.New("John      Egbert     Heir     Breath"),
					gem.New("Rose      Lalonde    Seer     Light "),
					gem.New("Roxy      Lalonde    Rogue    Void  "),
					gem.New("Vriska    Serket     Thief    Light "),
					gem.New("Jack      Noir       Dog      Woof  "),
					gem.New("Nepeta    Leijon     Rogue    Heart "),
				},
				LineSeparator: gem.New("\n"),
			},
		},
		{
			name: "custom charSet, border",
			table: [][]gem.String{
				{gem.New("Name"), gem.New("Surname"), gem.New("Class"), gem.New("Aspect")},
				{gem.New("John"), gem.New("Egbert"), gem.New("Heir"), gem.New("Breath")},
				{gem.New("Rose"), gem.New("Lalonde"), gem.New("Seer"), gem.New("Light")},
				{gem.New("Roxy"), gem.New("Lalonde"), gem.New("Rogue"), gem.New("Void")},
				{gem.New("Vriska"), gem.New("Serket"), gem.New("Thief"), gem.New("Light")},
				{gem.New("Jack"), gem.New("Noir"), gem.New("Dog"), gem.New("Woof")},
				{gem.New("Nepeta"), gem.New("Leijon"), gem.New("Rogue"), gem.New("Heart")},
			},
			width:   52,
			lineSep: gem.New("\n"),
			header:  true,
			border:  true,
			charSet: gem.New("X#_"),
			expect: tb.Block{
				Lines: []gem.String{
					gem.New("X____________X_____________X___________X___________X"),
					gem.New("#    NAME    #   SURNAME   #   CLASS   #   ASPECT  #"),
					gem.New("X____________X_____________X___________X___________X"),
					gem.New("# John       # Egbert      # Heir      # Breath    #"),
					gem.New("# Rose       # Lalonde     # Seer      # Light     #"),
					gem.New("# Roxy       # Lalonde     # Rogue     # Void      #"),
					gem.New("# Vriska     # Serket      # Thief     # Light     #"),
					gem.New("# Jack       # Noir        # Dog       # Woof      #"),
					gem.New("# Nepeta     # Leijon      # Rogue     # Heart     #"),
					gem.New("X____________X_____________X___________X___________X"),
				},
				LineSeparator: gem.New("\n"),
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert := assert.New(t)

			actual := LayoutTable(tc.table, tc.width, tc.lineSep, tc.header, tc.border, tc.charSet)

			assert.True(tc.expect.Equal(actual))
			assert.Equal(tc.expect.Join().String(), actual.Join().String())
		})
	}
}
