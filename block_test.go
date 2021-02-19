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
		{
			name: "multi line - no trailing separator - alternate separator",
			text: "john\negbert\trose\nlalonde\tdave\nstrider\tjade\nharley",
			sep:  "\t",
			expected: Block{
				Lines: []string{
					"john\negbert",
					"rose\nlalonde",
					"dave\nstrider",
					"jade\nharley",
				},
				LineSeparator:     "\t",
				TrailingSeparator: false,
			},
		},
		{
			name: "multi line - trailing separator - alternate separator",
			text: "john\negbert\trose\nlalonde\tdave\nstrider\tjade\nharley\t",
			sep:  "\t",
			expected: Block{
				Lines: []string{
					"john\negbert",
					"rose\nlalonde",
					"dave\nstrider",
					"jade\nharley",
				},
				LineSeparator:     "\t",
				TrailingSeparator: true,
			},
		},
	}

	assert := assertion.New(t)
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Reset(t)

			actual := NewBlock(tc.text, tc.sep)

			assert.Var("block").Equal(tc.expected, actual)
		})
	}
}

func Test_Block_Equal(t *testing.T) {
	testCases := []struct {
		name     string
		b1       Block
		b2       Block
		expected bool
	}{
		{
			name:     "nil Lines == empty Lines",
			b1:       Block{Lines: []string{}},
			b2:       Block{Lines: nil},
			expected: true,
		},
		{
			name:     "nil Lines == nil Lines",
			b1:       Block{Lines: nil},
			b2:       Block{Lines: nil},
			expected: true,
		},
		{
			name:     "empty Lines == empty Lines",
			b1:       Block{Lines: []string{}},
			b2:       Block{Lines: []string{}},
			expected: true,
		},
		{
			name:     "default == empty Lines",
			b1:       Block{},
			b2:       Block{Lines: []string{}},
			expected: true,
		},
		{
			name:     "default == nil Lines",
			b1:       Block{},
			b2:       Block{Lines: nil},
			expected: true,
		},
		{
			name:     "default == default",
			b1:       Block{},
			b2:       Block{},
			expected: true,
		},
		{
			name:     "empty Lines != 1 empty line",
			b1:       Block{Lines: []string{}},
			b2:       Block{Lines: []string{""}},
			expected: false,
		},
		{
			name:     "nil Lines != 1 empty line",
			b1:       Block{Lines: nil},
			b2:       Block{Lines: []string{""}},
			expected: false,
		},
		{
			name:     "default != 1 empty line",
			b1:       Block{},
			b2:       Block{Lines: []string{""}},
			expected: false,
		},
		{
			name:     "1 empty line == 1 empty line",
			b1:       Block{Lines: []string{""}},
			b2:       Block{Lines: []string{""}},
			expected: true,
		},
		{
			name:     "1 filled line == same filled line",
			b1:       Block{Lines: []string{"test"}},
			b2:       Block{Lines: []string{"test"}},
			expected: true,
		},
		{
			name:     "1 filled line != different filled line",
			b1:       Block{Lines: []string{"test1"}},
			b2:       Block{Lines: []string{"test2"}},
			expected: false,
		},
		{
			name:     "3 empty lines != 4 empty lines",
			b1:       Block{Lines: []string{"", "", ""}},
			b2:       Block{Lines: []string{"", "", "", ""}},
			expected: false,
		},
		{
			name:     "3 empty lines == 3 empty lines",
			b1:       Block{Lines: []string{"", "", ""}},
			b2:       Block{Lines: []string{"", "", ""}},
			expected: true,
		},
		{
			name:     "3 filled lines == same 3 filled lines",
			b1:       Block{Lines: []string{"a", "ab", "abc"}},
			b2:       Block{Lines: []string{"a", "ab", "abc"}},
			expected: true,
		},
		{
			name:     "3 filled lines != 4 filled lines",
			b1:       Block{Lines: []string{"a", "ab", "abc"}},
			b2:       Block{Lines: []string{"a", "ab", "abc", "abcd"}},
			expected: false,
		},
		{
			name:     "different separators",
			b1:       Block{LineSeparator: "\n"},
			b2:       Block{LineSeparator: "\t"},
			expected: false,
		},
		{
			name:     "different newline behavior",
			b1:       Block{TrailingSeparator: true},
			b2:       Block{TrailingSeparator: false},
			expected: false,
		},
	}

	assert := assertion.New(t)
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Reset(t)
			assert.Var("b1.Equal(b2)").Equal(tc.expected, tc.b1.Equal(tc.b2))
			assert.Var("b2.Equal(b1)").Equal(tc.expected, tc.b2.Equal(tc.b1))
		})
	}
}
