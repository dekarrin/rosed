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
