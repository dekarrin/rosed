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
				LineSeparator:            "---",
				IndentStr:                " ",
				NoTrailingLineSeparators: true,
			},
			expected: Options{
				LineSeparator:            "---",
				IndentStr:                " ",
				NoTrailingLineSeparators: true,
			},
		},
		{
			name: "missing line separator",
			input: Options{
				IndentStr:                " ",
				NoTrailingLineSeparators: false,
			},
			expected: Options{
				LineSeparator:            DefaultLineSeparator,
				IndentStr:                " ",
				NoTrailingLineSeparators: false,
			},
		},
		{
			name: "missing indent string",
			input: Options{
				LineSeparator:            "---",
				NoTrailingLineSeparators: true,
			},
			expected: Options{
				LineSeparator:            "---",
				IndentStr:                DefaultIndentString,
				NoTrailingLineSeparators: true,
			},
		},
		{
			name:  "zero-value defaults",
			input: Options{},
			expected: Options{
				LineSeparator:            DefaultLineSeparator,
				IndentStr:                DefaultIndentString,
				NoTrailingLineSeparators: false,
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
