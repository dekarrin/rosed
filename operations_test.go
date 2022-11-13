package rosed

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TODO: do not have test_wrap do any options checking, instead leave both WithOpts().Wrap() and WrapOpts() to
// Test_WrapOpts
func Test_Wrap(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		width    int
		options  Options
		expected []string
	}{
		{
			name:    "2 paragraphs, not preserved",
			input:   "this is a line that is split by paragraph in the input.\n\nThis is the second paragraph",
			width:   20,
			options: Options{PreserveParagraphs: false},
			expected: []string{
				"this is a line that",
				"is split by",
				"paragraph in the",
				"input. This is the",
				"second paragraph",
			},
		},
		{
			name:    "1 paragraph, preserved",
			input:   "this is a line that is split by paragraph in the input.",
			width:   20,
			options: Options{PreserveParagraphs: true},
			expected: []string{
				"this is a line that",
				"is split by",
				"paragraph in the",
				"input.",
			},
		},
		{
			name:    "2 paragraphs, preserved",
			input:   "this is a line that is split by paragraph in the input.\n\nThis is the second paragraph",
			width:   20,
			options: Options{PreserveParagraphs: true},
			expected: []string{
				"this is a line that",
				"is split by",
				"paragraph in the",
				"input.",
				"",
				"This is the second",
				"paragraph",
			},
		},
		{
			name:    "3 paragraphs, preserved",
			input:   "this is a line that is split by paragraph in the input.\n\nThis is the second paragraph.\n\nAnd this is the third",
			width:   20,
			options: Options{PreserveParagraphs: true},
			expected: []string{
				"this is a line that",
				"is split by",
				"paragraph in the",
				"input.",
				"",
				"This is the second",
				"paragraph.",
				"",
				"And this is the",
				"third",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert := assert.New(t)
			
			// try via Edit().WithOptions(), no defaults set
			actualText := Edit(tc.input).WithOptions(tc.options).Wrap(tc.width).String()

			// Edit().String returns one string; turn it into actual lines
			actual := strings.Split(actualText, tc.options.WithDefaults().LineSeparator)
			assert.Equal(tc.expected, actual)
		})
	}
}

func Test_WrapOpts(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		width    int
		options  Options
		expected []string
	}{
		{
			name:    "2 paragraphs, not preserved",
			input:   "this is a line that is split by paragraph in the input.\n\nThis is the second paragraph",
			width:   20,
			options: Options{PreserveParagraphs: false},
			expected: []string{
				"this is a line that",
				"is split by",
				"paragraph in the",
				"input. This is the",
				"second paragraph",
			},
		},
		{
			name:    "1 paragraph, preserved",
			input:   "this is a line that is split by paragraph in the input.",
			width:   20,
			options: Options{PreserveParagraphs: true},
			expected: []string{
				"this is a line that",
				"is split by",
				"paragraph in the",
				"input.",
			},
		},
		{
			name:    "2 paragraphs, preserved",
			input:   "this is a line that is split by paragraph in the input.\n\nThis is the second paragraph",
			width:   20,
			options: Options{PreserveParagraphs: true},
			expected: []string{
				"this is a line that",
				"is split by",
				"paragraph in the",
				"input.",
				"",
				"This is the second",
				"paragraph",
			},
		},
		{
			name:    "3 paragraphs, preserved",
			input:   "this is a line that is split by paragraph in the input.\n\nThis is the second paragraph.\n\nAnd this is the third",
			width:   20,
			options: Options{PreserveParagraphs: true},
			expected: []string{
				"this is a line that",
				"is split by",
				"paragraph in the",
				"input.",
				"",
				"This is the second",
				"paragraph.",
				"",
				"And this is the",
				"third",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert := assert.New(t)
			actualText := Edit(tc.input).WrapOpts(tc.width, tc.options).String()
			actual := strings.Split(actualText, tc.options.WithDefaults().LineSeparator)
			assert.Equal(tc.expected, actual)
		})
	}
}

func Test_CollapseSpaces(t *testing.T) {
	testCases := []struct{
		name string
		input string
		expect string
	}{
		{"empty line", "", ""},
		{"no spaces", "bluh", "bluh"},
		{"2 words", "word1 word2", "word1 word2"},
		{"3 words", "word1 word2 word3", "word1 word2 word3"},
		{"3 words with runs of spaces", "word1        word2  word3", "word1 word2 word3"},
		{"run of non-uniform spaces", "word1 \t\t word2", "word1 word2"},
		{"include lineSep", "   " + DefaultLineSeparator + " word1", " word1"},
	}
	
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert := assert.New(t)
			actual := Edit(tc.input).CollapseSpace().String()
			assert.Equal(tc.expect, actual)
		})
	}
}

func Test_CollapseSpacesOpts(t *testing.T) {
	testCases := []struct{
		name string
		input string
		options Options
		expect string
	}{
		{
			name: "empty line, custom lineSep",
			input: "",
			options: Options{LineSeparator: "<P>"},
			expect: "",
		},
		{
			name: "no spaces, custom lineSep",
			input: "bluh",
			options: Options{LineSeparator: "<P>"},
			expect: "bluh",
		},
		{
			name: "2 words, custom lineSep",
			input: "word1 word2",
			options: Options{LineSeparator: "<P>"},
			expect: "word1 word2",
		},
		{
			name: "3 words, custom lineSep",
			input: "word1 word2 word3",
			options: Options{LineSeparator: "<P>"},
			expect: "word1 word2 word3",
		},
		{
			name: "3 words with runs of spaces, custom lineSep",
			input: "word1        word2  word3",
			options: Options{LineSeparator: "<P>"},
			expect: "word1 word2 word3",
		},
		{
			name: "run of non-uniform spaces",
			input: "word1 \t\t word2",
			options: Options{LineSeparator: "<P>"},
			expect: "word1 word2",
		},
		{
			name: "include lineSep",
			input: "   <P> word1",
			options: Options{LineSeparator: "<P>"},
			expect: " word1",
		},
	}
	
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert := assert.New(t)
			
			actualDirect := Edit(tc.input).CollapseSpaceOpts(tc.options).String()
			actualPreOpts := Edit(tc.input).WithOptions(tc.options).CollapseSpace().String()
			
			assert.Equal(tc.expect, actualDirect)
			assert.Equal(tc.expect, actualPreOpts)
		})
	}
}

func Test_Insert(t *testing.T) {
	testCases := []struct{
		name string
		input string
		pos int
		insert string
		expect string
	}{
		{"empty into empty", "", 0, "", ""},
		{"empty into non-empty", "start", 0, "", "start"},
		{"non-empty into empty", "", 0, "TEST", "TEST"},
		{"at start of non-empty", "start", 0, "TEST", "TESTstart"},
		{"in middle of text", "abcdef", 2, "TEST", "abTESTcdef"},
		{"at end of text", "end", 3, "TEST", "endTEST"},
		{"far past end interpreted as end", "end", 300, "TEST", "endTEST"},
		{"negative index", "start", -3, "TEST", "stTESTart"},
		{"before decomposed grapheme", "franc\u0327ais", 1, "TEST", "fTESTranc\u0327ais"},
		{"after decomposed grapheme", "franc\u0327ais", 6, "TEST", "franc\u0327aTESTis"},
	}
	
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert := assert.New(t)
			
			actual := Edit(tc.input).Insert(tc.pos, tc.insert).String()
			
			assert.Equal(tc.expect, actual)
		})
	}
}
