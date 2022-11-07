package rosed

import (
	"testing"
	
	"github.com/stretchr/testify/assert"
)

func Test_Edit(t *testing.T) {
	testCases := []struct{
		name string
		input string
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

func Test_Editor_WithOptions(t *testing.T) {
	testCases := []struct{
		name string
		ed Editor
		opts Options
		expect Editor
	}{
		{
			name: "no prior options",
			ed: Editor{
				Text: "",
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
