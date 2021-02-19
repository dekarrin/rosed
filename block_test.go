package rosed

import (
	"testing"

	"github.com/dekarrin/rosed/internal/assertion"
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
				Lines:             []string{},
				LineSeparator:     "\n",
				TrailingSeparator: false,
			},
		},
		{
			name: "one line - no trailing newline",
			text: "hello",
			sep:  "\n",
			expected: Block{
				Lines: []string{
					"hello",
				},
				LineSeparator:     "\n",
				TrailingSeparator: false,
			},
		},
		{
			name: "one line - trailing newline",
			text: "hello\n",
			sep:  "\n",
			expected: Block{
				Lines: []string{
					"hello",
				},
				LineSeparator:     "\n",
				TrailingSeparator: true,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert := assertion.New(t)

			actual := NewBlock(tc.text, tc.sep)

			assert.Var("block").Equal(tc.expected, actual)
		})
	}
}
