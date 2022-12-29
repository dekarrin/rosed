package rosed

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Options_WithLineSeparator(t *testing.T) {
	testCases := []struct {
		name       string
		input      Options
		newLineSep string
		expected   Options
	}{
		{
			name: "from defaults",
			input: Options{
				ParagraphSeparator:       DefaultParagraphSeparator,
				LineSeparator:            DefaultLineSeparator,
				IndentStr:                DefaultIndentString,
				NoTrailingLineSeparators: false,
				PreserveParagraphs:       false,
			},
			newLineSep: "1234",
			expected: Options{
				ParagraphSeparator:       DefaultParagraphSeparator,
				LineSeparator:            "1234",
				IndentStr:                DefaultIndentString,
				NoTrailingLineSeparators: false,
				PreserveParagraphs:       false,
			},
		},
		{
			name: "from empty",
			input: Options{
				ParagraphSeparator:       "",
				LineSeparator:            "",
				IndentStr:                "",
				NoTrailingLineSeparators: false,
				PreserveParagraphs:       false,
			},
			newLineSep: "88888888",
			expected: Options{
				ParagraphSeparator:       "",
				LineSeparator:            "88888888",
				IndentStr:                "",
				NoTrailingLineSeparators: false,
				PreserveParagraphs:       false,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert := assert.New(t)

			actual := tc.input.WithLineSeparator(tc.newLineSep)

			assert.Equal(tc.expected, actual)
		})
	}
}

func Test_Options_WithIndentStr(t *testing.T) {
	testCases := []struct {
		name         string
		input        Options
		newIndentStr string
		expected     Options
	}{
		{
			name: "from defaults",
			input: Options{
				ParagraphSeparator:       DefaultParagraphSeparator,
				LineSeparator:            DefaultLineSeparator,
				IndentStr:                DefaultIndentString,
				NoTrailingLineSeparators: false,
				PreserveParagraphs:       false,
			},
			newIndentStr: "1234",
			expected: Options{
				ParagraphSeparator:       DefaultParagraphSeparator,
				LineSeparator:            DefaultLineSeparator,
				IndentStr:                "1234",
				NoTrailingLineSeparators: false,
				PreserveParagraphs:       false,
			},
		},
		{
			name: "from empty",
			input: Options{
				ParagraphSeparator:       "",
				LineSeparator:            "",
				IndentStr:                "",
				NoTrailingLineSeparators: false,
				PreserveParagraphs:       false,
			},
			newIndentStr: "88888888",
			expected: Options{
				ParagraphSeparator:       "",
				LineSeparator:            "",
				IndentStr:                "88888888",
				NoTrailingLineSeparators: false,
				PreserveParagraphs:       false,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert := assert.New(t)

			actual := tc.input.WithIndentStr(tc.newIndentStr)

			assert.Equal(tc.expected, actual)
		})
	}
}

func Test_Options_WithJustifyLastLine(t *testing.T) {
	testCases := []struct {
		name       string
		input      Options
		newJustify bool
		expected   Options
	}{
		{
			name: "from defaults",
			input: Options{
				ParagraphSeparator:       DefaultParagraphSeparator,
				LineSeparator:            DefaultLineSeparator,
				IndentStr:                DefaultIndentString,
				NoTrailingLineSeparators: false,
				PreserveParagraphs:       false,
				JustifyLastLine:          false,
			},
			newJustify: true,
			expected: Options{
				ParagraphSeparator:       DefaultParagraphSeparator,
				LineSeparator:            DefaultLineSeparator,
				IndentStr:                DefaultIndentString,
				NoTrailingLineSeparators: false,
				PreserveParagraphs:       false,
				JustifyLastLine:          true,
			},
		},
		{
			name: "from empty",
			input: Options{
				ParagraphSeparator:       "",
				LineSeparator:            "",
				IndentStr:                "",
				NoTrailingLineSeparators: false,
				PreserveParagraphs:       false,
				JustifyLastLine:          false,
			},
			newJustify: true,
			expected: Options{
				ParagraphSeparator:       "",
				LineSeparator:            "",
				IndentStr:                "",
				NoTrailingLineSeparators: false,
				PreserveParagraphs:       false,
				JustifyLastLine:          true,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert := assert.New(t)

			actual := tc.input.WithJustifyLastLine(tc.newJustify)

			assert.Equal(tc.expected, actual)
		})
	}
}

func Test_Options_WithParagraphSeparator(t *testing.T) {
	testCases := []struct {
		name       string
		input      Options
		newParaSep string
		expected   Options
	}{
		{
			name: "from defaults",
			input: Options{
				ParagraphSeparator:       DefaultParagraphSeparator,
				LineSeparator:            DefaultLineSeparator,
				IndentStr:                DefaultIndentString,
				NoTrailingLineSeparators: false,
				PreserveParagraphs:       false,
			},
			newParaSep: "1234",
			expected: Options{
				ParagraphSeparator:       "1234",
				LineSeparator:            DefaultLineSeparator,
				IndentStr:                DefaultIndentString,
				NoTrailingLineSeparators: false,
				PreserveParagraphs:       false,
			},
		},
		{
			name: "from empty",
			input: Options{
				ParagraphSeparator:       "",
				LineSeparator:            "",
				IndentStr:                "",
				NoTrailingLineSeparators: false,
				PreserveParagraphs:       false,
			},
			newParaSep: "88888888",
			expected: Options{
				ParagraphSeparator:       "88888888",
				LineSeparator:            "",
				IndentStr:                "",
				NoTrailingLineSeparators: false,
				PreserveParagraphs:       false,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert := assert.New(t)

			actual := tc.input.WithParagraphSeparator(tc.newParaSep)

			assert.Equal(tc.expected, actual)
		})
	}
}

func Test_Options_WithNoTrailingLineSeparators(t *testing.T) {
	testCases := []struct {
		name        string
		input       Options
		newTrailing bool
		expected    Options
	}{
		{
			name: "from defaults",
			input: Options{
				ParagraphSeparator:       DefaultParagraphSeparator,
				LineSeparator:            DefaultLineSeparator,
				IndentStr:                DefaultIndentString,
				NoTrailingLineSeparators: false,
				PreserveParagraphs:       false,
			},
			newTrailing: true,
			expected: Options{
				ParagraphSeparator:       DefaultParagraphSeparator,
				LineSeparator:            DefaultLineSeparator,
				IndentStr:                DefaultIndentString,
				NoTrailingLineSeparators: true,
				PreserveParagraphs:       false,
			},
		},
		{
			name: "from empty",
			input: Options{
				ParagraphSeparator:       "",
				LineSeparator:            "",
				IndentStr:                "",
				NoTrailingLineSeparators: false,
				PreserveParagraphs:       false,
			},
			newTrailing: true,
			expected: Options{
				ParagraphSeparator:       "",
				LineSeparator:            "",
				IndentStr:                "",
				NoTrailingLineSeparators: true,
				PreserveParagraphs:       false,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert := assert.New(t)

			actual := tc.input.WithNoTrailingLineSeparators(tc.newTrailing)

			assert.Equal(tc.expected, actual)
		})
	}
}

func Test_Options_WithPreserveParagraphs(t *testing.T) {
	testCases := []struct {
		name        string
		input       Options
		newPreserve bool
		expected    Options
	}{
		{
			name: "from defaults",
			input: Options{
				ParagraphSeparator:       DefaultParagraphSeparator,
				LineSeparator:            DefaultLineSeparator,
				IndentStr:                DefaultIndentString,
				NoTrailingLineSeparators: false,
				PreserveParagraphs:       false,
			},
			newPreserve: true,
			expected: Options{
				ParagraphSeparator:       DefaultParagraphSeparator,
				LineSeparator:            DefaultLineSeparator,
				IndentStr:                DefaultIndentString,
				NoTrailingLineSeparators: false,
				PreserveParagraphs:       true,
			},
		},
		{
			name: "from empty",
			input: Options{
				ParagraphSeparator:       "",
				LineSeparator:            "",
				IndentStr:                "",
				NoTrailingLineSeparators: false,
				PreserveParagraphs:       false,
			},
			newPreserve: true,
			expected: Options{
				ParagraphSeparator:       "",
				LineSeparator:            "",
				IndentStr:                "",
				NoTrailingLineSeparators: false,
				PreserveParagraphs:       true,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert := assert.New(t)

			actual := tc.input.WithPreserveParagraphs(tc.newPreserve)

			assert.Equal(tc.expected, actual)
		})
	}
}

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
				JustifyLastLine:          true,
				TableBorders:             true,
				TableHeaders:             true,
				TableCharSet:             "#!=",
			},
			expected: Options{
				ParagraphSeparator:       "\n\n--\n\n",
				LineSeparator:            "---",
				IndentStr:                " ",
				NoTrailingLineSeparators: true,
				PreserveParagraphs:       true,
				JustifyLastLine:          true,
				TableBorders:             true,
				TableHeaders:             true,
				TableCharSet:             "#!=",
			},
		},
		{
			name: "missing line separator",
			input: Options{
				ParagraphSeparator:       "\n\n--\n\n",
				IndentStr:                " ",
				NoTrailingLineSeparators: false,
				PreserveParagraphs:       true,
				JustifyLastLine:          true,
				TableBorders:             true,
				TableHeaders:             true,
				TableCharSet:             " ",
			},
			expected: Options{
				ParagraphSeparator:       "\n\n--\n\n",
				LineSeparator:            DefaultLineSeparator,
				IndentStr:                " ",
				NoTrailingLineSeparators: false,
				PreserveParagraphs:       true,
				JustifyLastLine:          true,
				TableBorders:             true,
				TableHeaders:             true,
				TableCharSet:             " |-",
			},
		},
		{
			name: "missing indent string",
			input: Options{
				ParagraphSeparator:       "\n\n--\n\n",
				LineSeparator:            "---",
				NoTrailingLineSeparators: true,
				PreserveParagraphs:       false,
				JustifyLastLine:          true,
				TableBorders:             true,
				TableHeaders:             true,
				TableCharSet:             "XY",
			},
			expected: Options{
				ParagraphSeparator:       "\n\n--\n\n",
				LineSeparator:            "---",
				IndentStr:                DefaultIndentString,
				NoTrailingLineSeparators: true,
				PreserveParagraphs:       false,
				JustifyLastLine:          true,
				TableBorders:             true,
				TableHeaders:             true,
				TableCharSet:             "XY-",
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
				JustifyLastLine:          false,
				TableBorders:             false,
				TableHeaders:             false,
				TableCharSet:             DefaultTableCharSet,
			},
		},
		{
			name: "dont unset JustifyLastLine",
			input: Options{
				JustifyLastLine: true,
			},
			expected: Options{
				ParagraphSeparator:       DefaultParagraphSeparator,
				LineSeparator:            DefaultLineSeparator,
				IndentStr:                DefaultIndentString,
				NoTrailingLineSeparators: false,
				PreserveParagraphs:       false,
				JustifyLastLine:          true,
				TableCharSet:             DefaultTableCharSet,
			},
		},
		{
			name: "dont unset TableBorders",
			input: Options{
				TableBorders: true,
			},
			expected: Options{
				ParagraphSeparator:       DefaultParagraphSeparator,
				LineSeparator:            DefaultLineSeparator,
				IndentStr:                DefaultIndentString,
				NoTrailingLineSeparators: false,
				PreserveParagraphs:       false,
				TableBorders:             true,
				TableCharSet:             DefaultTableCharSet,
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

func Test_Options_WithTableBorders(t *testing.T) {
	testCases := []struct {
		name       string
		input      Options
		newBorders bool
		expected   Options
	}{
		{
			name: "from defaults",
			input: Options{
				ParagraphSeparator:       DefaultParagraphSeparator,
				LineSeparator:            DefaultLineSeparator,
				IndentStr:                DefaultIndentString,
				NoTrailingLineSeparators: false,
				PreserveParagraphs:       false,
				JustifyLastLine:          false,
				TableBorders:             false,
				TableHeaders:             false,
				TableCharSet:             DefaultTableCharSet,
			},
			newBorders: true,
			expected: Options{
				ParagraphSeparator:       DefaultParagraphSeparator,
				LineSeparator:            DefaultLineSeparator,
				IndentStr:                DefaultIndentString,
				NoTrailingLineSeparators: false,
				PreserveParagraphs:       false,
				JustifyLastLine:          false,
				TableBorders:             true,
				TableHeaders:             false,
				TableCharSet:             DefaultTableCharSet,
			},
		},
		{
			name: "from empty",
			input: Options{
				ParagraphSeparator:       "",
				LineSeparator:            "",
				IndentStr:                "",
				NoTrailingLineSeparators: false,
				PreserveParagraphs:       false,
				JustifyLastLine:          false,
				TableBorders:             false,
				TableHeaders:             false,
				TableCharSet:             "",
			},
			newBorders: true,
			expected: Options{
				ParagraphSeparator:       "",
				LineSeparator:            "",
				IndentStr:                "",
				NoTrailingLineSeparators: false,
				PreserveParagraphs:       false,
				JustifyLastLine:          false,
				TableBorders:             true,
				TableHeaders:             false,
				TableCharSet:             "",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert := assert.New(t)

			actual := tc.input.WithTableBorders(tc.newBorders)

			assert.Equal(tc.expected, actual)
		})
	}
}

func Test_Options_WithTableHeaders(t *testing.T) {
	testCases := []struct {
		name       string
		input      Options
		newHeaders bool
		expected   Options
	}{
		{
			name: "from defaults",
			input: Options{
				ParagraphSeparator:       DefaultParagraphSeparator,
				LineSeparator:            DefaultLineSeparator,
				IndentStr:                DefaultIndentString,
				NoTrailingLineSeparators: false,
				PreserveParagraphs:       false,
				JustifyLastLine:          false,
				TableBorders:             false,
				TableHeaders:             false,
				TableCharSet:             DefaultTableCharSet,
			},
			newHeaders: true,
			expected: Options{
				ParagraphSeparator:       DefaultParagraphSeparator,
				LineSeparator:            DefaultLineSeparator,
				IndentStr:                DefaultIndentString,
				NoTrailingLineSeparators: false,
				PreserveParagraphs:       false,
				JustifyLastLine:          false,
				TableBorders:             false,
				TableHeaders:             true,
				TableCharSet:             DefaultTableCharSet,
			},
		},
		{
			name: "from empty",
			input: Options{
				ParagraphSeparator:       "",
				LineSeparator:            "",
				IndentStr:                "",
				NoTrailingLineSeparators: false,
				PreserveParagraphs:       false,
				JustifyLastLine:          false,
				TableBorders:             false,
				TableHeaders:             false,
				TableCharSet:             "",
			},
			newHeaders: true,
			expected: Options{
				ParagraphSeparator:       "",
				LineSeparator:            "",
				IndentStr:                "",
				NoTrailingLineSeparators: false,
				PreserveParagraphs:       false,
				JustifyLastLine:          false,
				TableBorders:             false,
				TableHeaders:             true,
				TableCharSet:             "",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert := assert.New(t)

			actual := tc.input.WithTableHeaders(tc.newHeaders)

			assert.Equal(tc.expected, actual)
		})
	}
}

func Test_Options_WithTableCharSet(t *testing.T) {
	testCases := []struct {
		name       string
		input      Options
		newCharSet string
		expected   Options
	}{
		{
			name: "from defaults",
			input: Options{
				ParagraphSeparator:       DefaultParagraphSeparator,
				LineSeparator:            DefaultLineSeparator,
				IndentStr:                DefaultIndentString,
				NoTrailingLineSeparators: false,
				PreserveParagraphs:       false,
				JustifyLastLine:          false,
				TableBorders:             false,
				TableHeaders:             false,
				TableCharSet:             DefaultTableCharSet,
			},
			newCharSet: "#!^",
			expected: Options{
				ParagraphSeparator:       DefaultParagraphSeparator,
				LineSeparator:            DefaultLineSeparator,
				IndentStr:                DefaultIndentString,
				NoTrailingLineSeparators: false,
				PreserveParagraphs:       false,
				JustifyLastLine:          false,
				TableBorders:             false,
				TableHeaders:             false,
				TableCharSet:             "#!^",
			},
		},
		{
			name: "from empty",
			input: Options{
				ParagraphSeparator:       "",
				LineSeparator:            "",
				IndentStr:                "",
				NoTrailingLineSeparators: false,
				PreserveParagraphs:       false,
				JustifyLastLine:          false,
				TableBorders:             false,
				TableHeaders:             false,
				TableCharSet:             "",
			},
			newCharSet: "XIL",
			expected: Options{
				ParagraphSeparator:       "",
				LineSeparator:            "",
				IndentStr:                "",
				NoTrailingLineSeparators: false,
				PreserveParagraphs:       false,
				JustifyLastLine:          false,
				TableBorders:             false,
				TableHeaders:             false,
				TableCharSet:             "XIL",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert := assert.New(t)

			actual := tc.input.WithTableCharSet(tc.newCharSet)

			assert.Equal(tc.expected, actual)
		})
	}
}
