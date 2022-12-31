package manip

import (
	"testing"

	"github.com/dekarrin/rosed/internal/gem"
	"github.com/dekarrin/rosed/internal/tb"

	"github.com/stretchr/testify/assert"
)

func Test_CollapseSpace(t *testing.T) {
	testCases := []struct {
		name   string
		input  gem.String
		sep    gem.String
		expect gem.String
	}{
		{"empty input", gem.Zero, gem.Zero, gem.Zero},
		{"no space to collapse", gem.New("no_spaces"), gem.Zero, gem.New("no_spaces")},
		{"one space mid-text is not collapsed (odd runecount)", gem.New("testA testB"), gem.Zero, gem.New("testA testB")},
		{"one space mid-text is not collapsed (even runecount)", gem.New("test testB"), gem.Zero, gem.New("test testB")},
		{"one space at start is not collapsed (odd runecount)", gem.New(" test"), gem.Zero, gem.New(" test")},
		{"one space at start is not collapsed (even runecount)", gem.New(" testA"), gem.Zero, gem.New(" testA")},
		{"one space at end is not collapsed (odd runecount)", gem.New("test "), gem.Zero, gem.New("test ")},
		{"one space at end is not collapsed (even runecount)", gem.New("testA "), gem.Zero, gem.New("testA ")},
		{"one space everywhere is not collapsed (odd runecount)", gem.New(" testA testB "), gem.Zero, gem.New(" testA testB ")},
		{"one space everywhere is not collapsed (even runecount)", gem.New(" test testB "), gem.Zero, gem.New(" test testB ")},
		{"non-spacechar whitespace is converted to space (odd runecount)", gem.New("testA\u0085testB"), gem.Zero, gem.New("testA testB")},
		{"non-spacechar whitespace is converted to space (even runecount)", gem.New("test\u0085testB"), gem.Zero, gem.New("test testB")},
		{"ws run is collapsed (spacechar)", gem.New("       testA  testB  "), gem.Zero, gem.New(" testA testB ")},
		{"ws run is collapsed (mixed ws)", gem.New("\u205f\u202ftestA\u200a  \t\n testB\t"), gem.Zero, gem.New(" testA testB ")},
		{"non-ws separator", gem.New("testA\n  testB <SEP>\u205f  testC"), gem.New("<SEP>"), gem.New("testA testB testC")},
		{"ws separator", gem.New("testA\n  testB <SEP>\n\n\u205f  testC"), gem.New("\n\n"), gem.New("testA testB <SEP> testC")},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert := assert.New(t)
			actual := CollapseSpace(tc.input, tc.sep)
			assert.True(tc.expect.Equal(actual))
		})
	}
}

func Test_Wrap(t *testing.T) {
	testCases := []struct {
		name     string
		input    gem.String
		width    int
		sep      gem.String
		expected tb.Block
	}{
		{
			name:  "empty input",
			input: gem.Zero,
			width: 80,
			sep:   gem.New("\n"),
			expected: tb.Block{
				Lines: []gem.String{
					gem.Zero,
				},
				LineSeparator: gem.New("\n"),
			},
		},
		{
			name:  "not enough to wrap",
			input: gem.New("a test string"),
			width: 80,
			sep:   gem.New("\n"),
			expected: tb.Block{
				Lines: []gem.String{
					gem.New("a test string"),
				},
				LineSeparator: gem.New("\n"),
			},
		},
		{
			name:  "2 line wrap",
			input: gem.New("a string long enough to be wrapped"),
			width: 20,
			sep:   gem.New("\n"),
			expected: tb.Block{
				Lines: []gem.String{
					gem.New("a string long enough"),
					gem.New("to be wrapped"),
				},
				LineSeparator: gem.New("\n"),
			},
		},
		{
			name:  "multi line wrap",
			input: gem.New("a string long enough to be wrapped more than once"),
			width: 20,
			sep:   gem.New("\n"),
			expected: tb.Block{
				Lines: []gem.String{
					gem.New("a string long enough"),
					gem.New("to be wrapped more"),
					gem.New("than once"),
				},
				LineSeparator: gem.New("\n"),
			},
		},
		{
			name:  "invalid width of -1 is interpreted as 2",
			input: gem.New("test"),
			width: -1,
			sep:   gem.New("\n"),
			expected: tb.Block{
				Lines: []gem.String{
					gem.New("t-"),
					gem.New("e-"),
					gem.New("st"),
				},
				LineSeparator: gem.New("\n"),
			},
		},
		{
			name:  "invalid width of 0 is interpreted as 2",
			input: gem.New("test"),
			width: 0,
			sep:   gem.New("\n"),
			expected: tb.Block{
				Lines: []gem.String{
					gem.New("t-"),
					gem.New("e-"),
					gem.New("st"),
				},
				LineSeparator: gem.New("\n"),
			},
		},
		{
			name:  "invalid width of 1 is interpreted as 2",
			input: gem.New("test"),
			width: 1,
			sep:   gem.New("\n"),
			expected: tb.Block{
				Lines: []gem.String{
					gem.New("t-"),
					gem.New("e-"),
					gem.New("st"),
				},
				LineSeparator: gem.New("\n"),
			},
		},
		{
			name:  "valid width of 2",
			input: gem.New("test"),
			width: 2,
			sep:   gem.New("\n"),
			expected: tb.Block{
				Lines: []gem.String{
					gem.New("t-"),
					gem.New("e-"),
					gem.New("st"),
				},
				LineSeparator: gem.New("\n"),
			},
		},
		{
			name:  "valid width of 3",
			input: gem.New("test"),
			width: 3,
			sep:   gem.New("\n"),
			expected: tb.Block{
				Lines: []gem.String{
					gem.New("te-"),
					gem.New("st"),
				},
				LineSeparator: gem.New("\n"),
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert := assert.New(t)
			actual := Wrap(tc.input, tc.width, tc.sep)
			assert.True(tc.expected.Equal(actual))
		})
	}
}

func Test_JustifyLine(t *testing.T) {
	testCases := []struct {
		name   string
		input  gem.String
		width  int
		expect gem.String
	}{
		{"empty line", gem.New(""), 10, gem.New("")},
		{"no spaces", gem.New("bluh"), 10, gem.New("bluh")},
		{"2 words", gem.New("word1 word2"), 20, gem.New("word1          word2")},
		{"3 words", gem.New("word1 word2 word3"), 20, gem.New("word1   word2  word3")},
		{"3 words with runs of spaces", gem.New("word1        word2  word3"), 20, gem.New("word1   word2  word3")},
		{"line longer than width", gem.New("hello"), 3, gem.New("hello")},
		{"bad width", gem.New("hello"), -1, gem.New("hello")},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert := assert.New(t)

			actual := JustifyLine(tc.input, tc.width)

			assert.True(tc.expect.Equal(actual))
		})
	}
}

func Test_CombineColumnBlocks(t *testing.T) {
	testCases := []struct {
		name     string
		left     tb.Block
		right    tb.Block
		minSpace int
		expect   tb.Block
	}{
		{
			name: "empty lines",
			left: tb.Block{
				Lines: []gem.String{gem.New("")},
			},
			right: tb.Block{
				Lines: []gem.String{gem.New("")},
			},
			minSpace: 0,
			expect: tb.Block{
				Lines: []gem.String{gem.New("")},
			},
		},
		{
			name: "right col bigger",
			left: tb.Block{
				Lines: []gem.String{
					gem.New("This is a test"),
					gem.New("string for the"),
					gem.New("left side")},
			},
			right: tb.Block{
				Lines: []gem.String{
					gem.New("Column number"),
					gem.New("two is right"),
					gem.New("here! And it"),
					gem.New("has a lot of"),
					gem.New("content that"),
					gem.New("will be"),
					gem.New("included"),
				},
			},
			minSpace: 2,
			expect: tb.Block{
				Lines: []gem.String{
					gem.New("This is a test  Column number"),
					gem.New("string for the  two is right"),
					gem.New("left side       here! And it"),
					gem.New("                has a lot of"),
					gem.New("                content that"),
					gem.New("                will be"),
					gem.New("                included"),
				},
			},
		},
		{
			name: "left col bigger",
			left: tb.Block{
				Lines: []gem.String{
					gem.New("Column number"),
					gem.New("one is right"),
					gem.New("here! And it"),
					gem.New("has a lot of"),
					gem.New("content that"),
					gem.New("will be"),
					gem.New("included"),
				},
			},
			right: tb.Block{
				Lines: []gem.String{
					gem.New("This is a test"),
					gem.New("string for the"),
					gem.New("right side")},
			},
			minSpace: 2,
			expect: tb.Block{
				Lines: []gem.String{
					gem.New("Column number  This is a test"),
					gem.New("one is right   string for the"),
					gem.New("here! And it   right side"),
					gem.New("has a lot of   "),
					gem.New("content that   "),
					gem.New("will be        "),
					gem.New("included       "),
				},
			},
		},
		{
			name: "equal size columns",
			left: tb.Block{
				Lines: []gem.String{
					gem.New("Column number"),
					gem.New("one is right"),
					gem.New("here!"),
				},
			},
			right: tb.Block{
				Lines: []gem.String{
					gem.New("This is a test"),
					gem.New("string for the"),
					gem.New("right side")},
			},
			minSpace: 2,
			expect: tb.Block{
				Lines: []gem.String{
					gem.New("Column number  This is a test"),
					gem.New("one is right   string for the"),
					gem.New("here!          right side"),
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert := assert.New(t)

			actual := CombineColumnBlocks(tc.left, tc.right, tc.minSpace)

			assert.Equal(tc.expect, actual)
		})
	}
}

func Test_AlignLineLeft(t *testing.T) {
	testCases := []struct {
		name   string
		input  gem.String
		width  int
		expect gem.String
	}{
		{"empty string to 0", gem.Zero, 0, gem.Zero},
		{"empty string to 10", gem.Zero, 10, gem.New("          ")},
		{"empty string to -1", gem.Zero, 0, gem.Zero},
		{"non-empty, normal", gem.New("rose"), 10, gem.New("rose      ")},
		{"non-empty, already big", gem.New("rose"), 2, gem.New("rose")},
		{"non-empty, to -1", gem.New("rose"), -1, gem.New("rose")},
		{"non-empty, space on left", gem.New(" rose"), 10, gem.New("rose      ")},
		{"non-empty, space on right", gem.New("rose  "), 10, gem.New("rose      ")},
		{"non-empty, space on both sides", gem.New(" rose  "), 10, gem.New("rose      ")},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert := assert.New(t)

			actual := AlignLineLeft(tc.input, tc.width)

			assert.True(tc.expect.Equal(actual))
			assert.Equal(tc.expect.String(), actual.String())
		})
	}
}

func Test_AlignLineRight(t *testing.T) {
	testCases := []struct {
		name   string
		input  gem.String
		width  int
		expect gem.String
	}{
		{"empty string to 0", gem.Zero, 0, gem.Zero},
		{"empty string to 10", gem.Zero, 10, gem.New("          ")},
		{"empty string to -1", gem.Zero, 0, gem.Zero},
		{"non-empty, normal", gem.New("rose  lalonde"), 20, gem.New("       rose  lalonde")},
		{"non-empty, already big", gem.New("rose"), 2, gem.New("rose")},
		{"non-empty, to -1", gem.New("rose"), -1, gem.New("rose")},
		{"non-empty, space on left", gem.New("  rose"), 10, gem.New("      rose")},
		{"non-empty, space on right", gem.New("rose  lalonde  "), 20, gem.New("       rose  lalonde")},
		{"non-empty, space on both sides", gem.New("   rose  "), 10, gem.New("      rose")},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert := assert.New(t)

			actual := AlignLineRight(tc.input, tc.width)

			assert.True(tc.expect.Equal(actual))
			assert.Equal(tc.expect.String(), actual.String())
		})
	}
}

func Test_AlignLineCenter(t *testing.T) {
	testCases := []struct {
		name   string
		input  gem.String
		width  int
		expect gem.String
	}{
		{"empty string to 0", gem.Zero, 0, gem.Zero},
		{"empty string to 10", gem.Zero, 10, gem.New("          ")},
		{"empty string to -1", gem.Zero, 0, gem.Zero},
		{"non-empty, normal", gem.New("john egbert"), 17, gem.New("   john egbert   ")},
		{"non-empty, normal, not even", gem.New("john egbert"), 18, gem.New("    john egbert   ")},
		{"non-empty, already big", gem.New("john egbert"), 2, gem.New("john egbert")},
		{"non-empty, to -1", gem.New("john"), -1, gem.New("john")},
		{"non-empty, space on left", gem.New("  john"), 12, gem.New("    john    ")},
		{"non-empty, space on right", gem.New("john  "), 12, gem.New("    john    ")},
		{"non-empty, space on both sides", gem.New("   john  "), 12, gem.New("    john    ")},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert := assert.New(t)

			actual := AlignLineCenter(tc.input, tc.width)

			assert.True(tc.expect.Equal(actual))
			assert.Equal(tc.expect.String(), actual.String())
		})
	}
}

func Test_CountLeadingWhitespace(t *testing.T) {
	testCases := []struct {
		name   string
		input  gem.String
		expect int
	}{
		{"empty string", gem.Zero, 0},
		{"only space", gem.RepeatStr(" ", 8), 8},
		{"no leading or trailing space", gem.New("feferi peixes"), 0},
		{"no leading space, with trailing space", gem.New("feferi peixes   "), 0},
		{"leading space, no trailing space", gem.New(" feferi peixes"), 1},
		{"leading and trailing space", gem.New("  feferi peixes   "), 2},
		{"multiple kinds of whitespace", gem.New("\f\v\t\n\r\u200a\u3000 feferi peixes"), 8},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert := assert.New(t)

			actual := CountLeadingWhitespace(tc.input)

			assert.Equal(tc.expect, actual)
		})
	}
}

func Test_CountTrailingWhitespace(t *testing.T) {
	testCases := []struct {
		name   string
		input  gem.String
		expect int
	}{
		{"empty string", gem.Zero, 0},
		{"only space", gem.RepeatStr(" ", 8), 8},
		{"no leading or trailing space", gem.New("feferi peixes"), 0},
		{"no leading space, with trailing space", gem.New("feferi peixes   "), 3},
		{"leading space, no trailing space", gem.New(" feferi peixes"), 0},
		{"leading and trailing space", gem.New("  feferi peixes   "), 3},
		{"multiple kinds of whitespace", gem.New("feferi peixes \f\v\t\n\r\u200a\u3000"), 8},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert := assert.New(t)

			actual := CountTrailingWhitespace(tc.input)

			assert.Equal(tc.expect, actual)
		})
	}
}
