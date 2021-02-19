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
			name: "no lines - no trailing newline",
			text: "",
			sep:  "\n",
			expected: Block{
				Lines:             []string{},
				LineSeparator:     "\n",
				TrailingSeparator: false,
			},
		},
		{
			name: "no lines - trailing newline",
			text: "\n",
			sep:  "\n",
			expected: Block{
				Lines: []string{
					"",
				},
				LineSeparator:     "\n",
				TrailingSeparator: true,
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
		{
			name: "multi line - no trailing newline",
			text: "hello\nthere\ntest",
			sep:  "\n",
			expected: Block{
				Lines: []string{
					"hello",
					"there",
					"test",
				},
				LineSeparator:     "\n",
				TrailingSeparator: false,
			},
		},
		{
			name: "multi line - trailing newline",
			text: "hello\nthere\ntest\n",
			sep:  "\n",
			expected: Block{
				Lines: []string{
					"hello",
					"there",
					"test",
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

func Test_Block_Equal(t *testing.T) {
	testCases := []struct {
		name   string
		b1     Block
		b2     Block
		expect bool
	}{}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

		})
	}
}
