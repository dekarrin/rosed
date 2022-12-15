package rosed

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Edit(t *testing.T) {
	testCases := []struct {
		name   string
		input  string
		expect Editor
	}{
		{"empty string", "", Editor{Text: ""}},
		{"non-empty string", "test", Editor{Text: "test"}},
		{"many breaks in string", "test \n testtest\n\ntest", Editor{Text: "test \n testtest\n\ntest"}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert := assert.New(t)

			actual := Edit(tc.input)

			assert.Equal(tc.expect, actual)
		})
	}
}

func Test_Editor_CharCount(t *testing.T) {
	testCases := []struct {
		name   string
		input  string
		expect int
	}{
		{"empty string", "", 0},
		{"1-char", "1", 1},
		{"multi-char", "test", 4},
		{"decomposed sequence", "C\u0327", 1},
		{"one emoji with two codepoints: ❤️", "\u2764\uFE0F", 1},
		{"several codepoint emoji woman kiss woman 👩‍❤️‍💋‍👩", "\U0001F469\u200D\u2764\uFE0F\u200D\U0001F48B\u200D\U0001F469", 1},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert := assert.New(t)

			actual := Edit(tc.input).CharCount()

			assert.Equal(tc.expect, actual)
		})
	}
}

func Test_Editor_WithOptions(t *testing.T) {
	testCases := []struct {
		name   string
		ed     Editor
		opts   Options
		expect Editor
	}{
		{
			name: "no prior options",
			ed: Editor{
				Text:    "",
				Options: Options{},
			},
			opts: Options{
				LineSeparator: "x",
			},
			expect: Editor{
				Text: "",
				Options: Options{
					LineSeparator: "x",
				},
			},
		},
		{
			name: "all options are replaced by default",
			ed: Editor{
				Text: "test",
				Options: Options{
					ParagraphSeparator:       "PP",
					LineSeparator:            "x",
					IndentStr:                "             ",
					NoTrailingLineSeparators: true,
					PreserveParagraphs:       true,
				},
			},
			opts: Options{
				LineSeparator: "\n",
			},
			expect: Editor{
				Text: "test",
				Options: Options{
					LineSeparator: "\n",
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert := assert.New(t)

			actual := tc.ed.WithOptions(tc.opts)

			assert.Equal(tc.expect, actual)
		})
	}
}

func Test_Editor_LineCount(t *testing.T) {
	testCases := []struct {
		name   string
		ed     Editor
		expect int
	}{
		{
			name: "empty string gives 0 with default linesep and trailingseps",
			ed: Editor{
				Text: "",
			},
			expect: 0,
		},
		{
			name: "empty string gives 0 with non-default linesep and default trailingseps",
			ed: Editor{
				Text: "",
				Options: Options{
					LineSeparator: "<P>",
				},
			},
			expect: 0,
		},
		{
			name: "empty string gives 1 with non-default linesep and trailingseps",
			ed: Editor{
				Text: "",
				Options: Options{
					LineSeparator:            "<P>",
					NoTrailingLineSeparators: true,
				},
			},
			expect: 1,
		},
		{
			name: "empty string gives 1 with default linesep and non-default trailingseps",
			ed: Editor{
				Text: "",
				Options: Options{
					NoTrailingLineSeparators: true,
				},
			},
			expect: 1,
		},
		{
			name: "1-line string (no trailing line sep) gives 1",
			ed: Editor{
				Text: "test",
			},
			expect: 1,
		},
		{
			name: "1-line string (with default trailing line sep) gives 1",
			ed: Editor{
				Text: "test" + DefaultLineSeparator,
			},
			expect: 1,
		},
		{
			name: "1-line string (with non-default trailing line sep) gives 1",
			ed: Editor{
				Text: "test<P>",
				Options: Options{
					LineSeparator: "<P>",
				},
			},
			expect: 1,
		},
		{
			name: "1-line string (with NoTrailing on) gives 1",
			ed: Editor{
				Text: "test",
			},
			expect: 1,
		},
		{
			name: "multi-line string",
			ed: Editor{
				Text: "test1\ntest2\ntest3\ntest4",
			},
			expect: 4,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert := assert.New(t)

			actual := tc.ed.LineCount()

			assert.Equal(tc.expect, actual)
		})
	}
}
