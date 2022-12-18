package rosed

import (
	"testing"

	"github.com/dekarrin/rosed/internal/gem"

	"github.com/stretchr/testify/assert"
)

func Test_Manip_collapseSpace(t *testing.T) {
	testCases := []struct {
		name   string
		input  gem.String
		sep    gem.String
		expect gem.String
	}{
		{"empty input", gem.Zero, gem.Zero, gem.Zero},
		{"no space to collapse", _g("no_spaces"), gem.Zero, _g("no_spaces")},
		{"one space mid-text is not collapsed (odd runecount)", _g("testA testB"), gem.Zero, _g("testA testB")},
		{"one space mid-text is not collapsed (even runecount)", _g("test testB"), gem.Zero, _g("test testB")},
		{"one space at start is not collapsed (odd runecount)", _g(" test"), gem.Zero, _g(" test")},
		{"one space at start is not collapsed (even runecount)", _g(" testA"), gem.Zero, _g(" testA")},
		{"one space at end is not collapsed (odd runecount)", _g("test "), gem.Zero, _g("test ")},
		{"one space at end is not collapsed (even runecount)", _g("testA "), gem.Zero, _g("testA ")},
		{"one space everywhere is not collapsed (odd runecount)", _g(" testA testB "), gem.Zero, _g(" testA testB ")},
		{"one space everywhere is not collapsed (even runecount)", _g(" test testB "), gem.Zero, _g(" test testB ")},
		{"non-spacechar whitespace is converted to space (odd runecount)", _g("testA\u0085testB"), gem.Zero, _g("testA testB")},
		{"non-spacechar whitespace is converted to space (even runecount)", _g("test\u0085testB"), gem.Zero, _g("test testB")},
		{"ws run is collapsed (spacechar)", _g("       testA  testB  "), gem.Zero, _g(" testA testB ")},
		{"ws run is collapsed (mixed ws)", _g("\u205f\u202ftestA\u200a  \t\n testB\t"), gem.Zero, _g(" testA testB ")},
		{"non-ws separator", _g("testA\n  testB <SEP>\u205f  testC"), _g("<SEP>"), _g("testA testB testC")},
		{"ws separator", _g("testA\n  testB <SEP>\n\n\u205f  testC"), _g("\n\n"), _g("testA testB <SEP> testC")},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert := assert.New(t)
			actual := collapseSpace(tc.input, tc.sep)
			assert.True(tc.expect.Equal(actual))
		})
	}
}

func Test_Manip_wrap(t *testing.T) {
	testCases := []struct {
		name     string
		input    gem.String
		width    int
		sep      gem.String
		expected block
	}{
		{
			name:  "empty input",
			input: gem.Zero,
			width: 80,
			sep:   _g("\n"),
			expected: block{
				Lines: []gem.String{
					gem.Zero,
				},
				LineSeparator: _g("\n"),
			},
		},
		{
			name:  "not enough to wrap",
			input: _g("a test string"),
			width: 80,
			sep:   _g("\n"),
			expected: block{
				Lines: []gem.String{
					_g("a test string"),
				},
				LineSeparator: _g("\n"),
			},
		},
		{
			name:  "2 line wrap",
			input: _g("a string long enough to be wrapped"),
			width: 20,
			sep:   _g("\n"),
			expected: block{
				Lines: []gem.String{
					_g("a string long enough"),
					_g("to be wrapped"),
				},
				LineSeparator: _g("\n"),
			},
		},
		{
			name:  "multi line wrap",
			input: _g("a string long enough to be wrapped more than once"),
			width: 20,
			sep:   _g("\n"),
			expected: block{
				Lines: []gem.String{
					_g("a string long enough"),
					_g("to be wrapped more"),
					_g("than once"),
				},
				LineSeparator: _g("\n"),
			},
		},
		{
			name:  "invalid width of -1 is interpreted as 2",
			input: _g("test"),
			width: -1,
			sep:   _g("\n"),
			expected: block{
				Lines: []gem.String{
					_g("t-"),
					_g("e-"),
					_g("st"),
				},
				LineSeparator: _g("\n"),
			},
		},
		{
			name:  "invalid width of 0 is interpreted as 2",
			input: _g("test"),
			width: 0,
			sep:   _g("\n"),
			expected: block{
				Lines: []gem.String{
					_g("t-"),
					_g("e-"),
					_g("st"),
				},
				LineSeparator: _g("\n"),
			},
		},
		{
			name:  "invalid width of 1 is interpreted as 2",
			input: _g("test"),
			width: 1,
			sep:   _g("\n"),
			expected: block{
				Lines: []gem.String{
					_g("t-"),
					_g("e-"),
					_g("st"),
				},
				LineSeparator: _g("\n"),
			},
		},
		{
			name:  "valid width of 2",
			input: _g("test"),
			width: 2,
			sep:   _g("\n"),
			expected: block{
				Lines: []gem.String{
					_g("t-"),
					_g("e-"),
					_g("st"),
				},
				LineSeparator: _g("\n"),
			},
		},
		{
			name:  "valid width of 3",
			input: _g("test"),
			width: 3,
			sep:   _g("\n"),
			expected: block{
				Lines: []gem.String{
					_g("te-"),
					_g("st"),
				},
				LineSeparator: _g("\n"),
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert := assert.New(t)
			actual := wrap(tc.input, tc.width, tc.sep)
			assert.Equal(tc.expected, actual)
		})
	}
}

func Test_Manip_justifyLine(t *testing.T) {
	testCases := []struct {
		name   string
		input  gem.String
		width  int
		expect gem.String
	}{
		{"empty line", _g(""), 10, _g("")},
		{"no spaces", _g("bluh"), 10, _g("bluh")},
		{"2 words", _g("word1 word2"), 20, _g("word1          word2")},
		{"3 words", _g("word1 word2 word3"), 20, _g("word1   word2  word3")},
		{"3 words with runs of spaces", _g("word1        word2  word3"), 20, _g("word1   word2  word3")},
		{"line longer than width", _g("hello"), 3, _g("hello")},
		{"bad width", _g("hello"), -1, _g("hello")},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert := assert.New(t)

			actual := justifyLine(tc.input, tc.width)

			assert.True(tc.expect.Equal(actual))
		})
	}
}

func Test_Manip_combineColumnBlocks(t *testing.T) {
	testCases := []struct {
		name     string
		left     block
		right    block
		minSpace int
		expect   block
	}{
		{
			name: "empty lines",
			left: block{
				Lines: []gem.String{_g("")},
			},
			right: block{
				Lines: []gem.String{_g("")},
			},
			minSpace: 0,
			expect: block{
				Lines: []gem.String{_g("")},
			},
		},
		{
			name: "right col bigger",
			left: block{
				Lines: []gem.String{
					_g("This is a test"),
					_g("string for the"),
					_g("left side")},
			},
			right: block{
				Lines: []gem.String{
					_g("Column number"),
					_g("two is right"),
					_g("here! And it"),
					_g("has a lot of"),
					_g("content that"),
					_g("will be"),
					_g("included"),
				},
			},
			minSpace: 2,
			expect: block{
				Lines: []gem.String{
					_g("This is a test  Column number"),
					_g("string for the  two is right"),
					_g("left side       here! And it"),
					_g("                has a lot of"),
					_g("                content that"),
					_g("                will be"),
					_g("                included"),
				},
			},
		},
		{
			name: "left col bigger",
			left: block{
				Lines: []gem.String{
					_g("Column number"),
					_g("one is right"),
					_g("here! And it"),
					_g("has a lot of"),
					_g("content that"),
					_g("will be"),
					_g("included"),
				},
			},
			right: block{
				Lines: []gem.String{
					_g("This is a test"),
					_g("string for the"),
					_g("right side")},
			},
			minSpace: 2,
			expect: block{
				Lines: []gem.String{
					_g("Column number  This is a test"),
					_g("one is right   string for the"),
					_g("here! And it   right side"),
					_g("has a lot of   "),
					_g("content that   "),
					_g("will be        "),
					_g("included       "),
				},
			},
		},
		{
			name: "equal size columns",
			left: block{
				Lines: []gem.String{
					_g("Column number"),
					_g("one is right"),
					_g("here!"),
				},
			},
			right: block{
				Lines: []gem.String{
					_g("This is a test"),
					_g("string for the"),
					_g("right side")},
			},
			minSpace: 2,
			expect: block{
				Lines: []gem.String{
					_g("Column number  This is a test"),
					_g("one is right   string for the"),
					_g("here!          right side"),
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert := assert.New(t)

			actual := combineColumnBlocks(tc.left, tc.right, tc.minSpace)

			assert.Equal(tc.expect, actual)
		})
	}
}

func Test_Manip_alignLineLeft(t *testing.T) {
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

			actual := alignLineLeft(tc.input, tc.width)

			assert.True(tc.expect.Equal(actual))
			assert.Equal(tc.expect.String(), actual.String())
		})
	}
}

func Test_Manip_alignLineRight(t *testing.T) {
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

			actual := alignLineRight(tc.input, tc.width)

			assert.True(tc.expect.Equal(actual))
			assert.Equal(tc.expect.String(), actual.String())
		})
	}
}

func Test_Manip_alignLineCenter(t *testing.T) {
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

			actual := alignLineCenter(tc.input, tc.width)

			assert.True(tc.expect.Equal(actual))
			assert.Equal(tc.expect.String(), actual.String())
		})
	}
}

func Test_Manip_countLeadingWhitespace(t *testing.T) {
	testCases := []struct {
		name   string
		input  gem.String
		expect int
	}{
		{"empty string", gem.Zero, 0},
		{"no leading or trailing space", gem.New("feferi peixes"), 0},
		{"no leading space, with trailing space", gem.New("feferi peixes   "), 0},
		{"leading space, no trailing space", gem.New(" feferi peixes"), 1},
		{"leading and trailing space", gem.New("  feferi peixes   "), 2},
		{"multiple kinds of whitespace", gem.New("\f\v\t\n\r\u200a\u3000 feferi peixes"), 8},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert := assert.New(t)

			actual := countLeadingWhitespace(tc.input)

			assert.Equal(tc.expect, actual)
		})
	}
}

func Test_Manip_countTrailingWhitespace(t *testing.T) {
	testCases := []struct {
		name   string
		input  gem.String
		expect int
	}{
		{"empty string", gem.Zero, 0},
		{"no leading or trailing space", gem.New("feferi peixes"), 0},
		{"no leading space, with trailing space", gem.New("feferi peixes   "), 3},
		{"leading space, no trailing space", gem.New(" feferi peixes"), 0},
		{"leading and trailing space", gem.New("  feferi peixes   "), 3},
		{"multiple kinds of whitespace", gem.New("feferi peixes \f\v\t\n\r\u200a\u3000"), 8},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert := assert.New(t)

			actual := countTrailingWhitespace(tc.input)

			assert.Equal(tc.expect, actual)
		})
	}
}
