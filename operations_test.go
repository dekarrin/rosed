package rosed

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

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
