package rosed

import (
	"testing"
	
	"github.com/stretchr/testify/assert"
)

func Test_Editor_Chars(t *testing.T) {
	testCases := []struct{
		name string
		ed Editor
		start int
		end int
		expect Editor
	}{
		{
			name: "typical subsection",
			ed: Editor{
				Text: "some testing text to edit",
			},
			start: 5,
			end: 17,
			expect: Editor{
				Text: "testing text",
			},
		},
		{
			name: "end > end of the string",
			ed: Editor{
				Text: "test",
			},
			start: 1,
			end: 20,
			expect: Editor{
				Text: "est",
			},
		},
		{
			name: "start < 0",
			ed: Editor{
				Text: "test",
			},
			start: -3,
			end: 2,
			expect: Editor{
				Text: "e",
			},
		},
		{
			name: "start > end",
			ed: Editor{
				Text: "test",
			},
			start: 2,
			end: 1,
			expect: Editor{
				Text: "",
			},
		},
		{
			name: "start = end",
			ed: Editor{
				Text: "test",
			},
			start: 2,
			end: 2,
			expect: Editor{
				Text: "",
			},
		},
		{
			name: "both < 0",
			ed: Editor{
				Text: "test",
			},
			start: 2,
			end: 2,
			expect: Editor{
				Text: "",
			},
		},
		{
			name: "get before decomposed grapheme",
			ed: Editor{
				Text: "Franc\u0327ais",
			},
			start: 1,
			end: 4,
			expect: Editor{
				Text: "ran",
			},
		},
		{
			name: "get after decomposed grapheme",
			ed: Editor{
				Text: "Franc\u0327ais",
			},
			start: 5,
			end: 7,
			expect: Editor{
				Text: "ai",
			},
		},
		{
			name: "get across decomposed grapheme",
			ed: Editor{
				Text: "Franc\u0327ais",
			},
			start: 3,
			end: 7,
			expect: Editor{
				Text: "nc\u0327ai",
			},
		},
	}
	
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert := assert.New(t)
			actual := tc.ed.Chars(tc.start, tc.end)
			
			// don't do a full Equal as that will compare unexported
			// fields; instead just check the ones we care about
			
			assert.Equal(tc.expect.Options, actual.Options)
			assert.Equal(tc.expect.Text, actual.Text)
		})
	}
}

