package rosed

import "testing"

func Test_Options_WithDefaults(t *testing.T) {
	testCases := []struct {
		name     string
		input    Options
		expected Options
	}{
		{
			name: "no empties",
			input: Options{
				ParagraphSeparator:       "\n\n--\n\n",
				LineSeparator:            "---",
				IndentStr:                " ",
				NoTrailingLineSeparators: true,
				PreserveParagraphs:       true,
			},
			expected: Options{
				ParagraphSeparator:       "\n\n--\n\n",
				LineSeparator:            "---",
				IndentStr:                " ",
				NoTrailingLineSeparators: true,
				PreserveParagraphs:       true,
			},
		},
		{
			name: "missing line separator",
			input: Options{
				ParagraphSeparator:       "\n\n--\n\n",
				IndentStr:                " ",
				NoTrailingLineSeparators: false,
				PreserveParagraphs:       true,
			},
			expected: Options{
				ParagraphSeparator:       "\n\n--\n\n",
				LineSeparator:            DefaultLineSeparator,
				IndentStr:                " ",
				NoTrailingLineSeparators: false,
				PreserveParagraphs:       true,
			},
		},
		{
			name: "missing indent string",
			input: Options{
				ParagraphSeparator:       "\n\n--\n\n",
				LineSeparator:            "---",
				NoTrailingLineSeparators: true,
				PreserveParagraphs:       false,
			},
			expected: Options{
				ParagraphSeparator:       "\n\n--\n\n",
				LineSeparator:            "---",
				IndentStr:                DefaultIndentString,
				NoTrailingLineSeparators: true,
				PreserveParagraphs:       false,
			},
		},
		{
			name:  "zero-value defaults",
			input: Options{},
			expected: Options{
				ParagraphSeparator:       DefaultParagraphSeparator,
				LineSeparator:            DefaultLineSeparator,
				IndentStr:                DefaultIndentString,
				NoTrailingLineSeparators: false,
				PreserveParagraphs:       false,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			actual := tc.input.WithDefaults()

			if actual != tc.expected {
				t.Fatalf("expected %v but was %v", tc.expected, actual)
			}
		})
	}
}
