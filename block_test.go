package rosed

import (
	"testing"

	"github.com/dekarrin/rosed/internal/assert"
)

func Test_NewBlock(t *testing.T) {
	testCases := []struct {
		name     string
		text     string
		sep      string
		expected Block
	}{
		{
			name: "no lines",
			text: "",
			sep:  "\n",
			expected: Block{
				Lines: []string{
					"",
				},
				LineSeparator:     "\n",
				TrailingSeparator: false,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			asrt := assert.New(t)

			actual := NewBlock(tc.text, tc.sep)

			asrt.Var("block").Equal(tc.expected, actual)

		})
	}
}
