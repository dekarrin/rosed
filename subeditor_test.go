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
		expectPanic bool
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
	}
	
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert := assert.New(t)
			
			if tc.expectPanic {
				assert.Panics(func() {
					tc.ed.Chars(tc.start, tc.end)
				})
			} else {
				actual := tc.ed.Chars(tc.start, tc.end)
				
				// don't do a full Equal as that will compare unexported
				// fields; instead just check the ones we care about
				
				assert.Equal(tc.expect.Options, actual.Options)
				assert.Equal(tc.expect.Text, actual.Text)
			}
		})
	}
}
