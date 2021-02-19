package rosed

import (
	"sort"
	"testing"

	"github.com/dekarrin/assertion"
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

func Test_Block_Len(t *testing.T) {
	testCases := []struct {
		name     string
		input    Block
		expected int
	}{
		{"nil Lines", Block{Lines: nil}, 0},
		{"empty Lines", Block{Lines: []string{}}, 0},
		{"default Lines", Block{}, 0},
		{"1 empty line", Block{Lines: []string{""}}, 1},
		{"1 filled line", Block{Lines: []string{"test"}}, 1},
		{"3 empty lines", Block{Lines: []string{"", "", ""}}, 3},
		{"3 filled lines", Block{Lines: []string{"a", "b", "c"}}, 3},
	}

	assert := assertion.New(t)
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Reset(t)

			actual := tc.input.Len()

			assert.Var("Len()").Equal(tc.expected, actual)
		})
	}
}

func Test_sort_Block(t *testing.T) {
	testCases := []struct {
		name     string
		input    Block
		expected Block
	}{
		{
			name:     "nil Lines does not change",
			input:    Block{Lines: nil},
			expected: Block{Lines: nil},
		},
		{
			name:     "empty Lines does not change",
			input:    Block{Lines: []string{}},
			expected: Block{Lines: []string{}},
		},
		{
			name:     "default does not change",
			input:    Block{},
			expected: Block{},
		},
		{
			name:     "1 line does not change",
			input:    Block{Lines: []string{"test"}},
			expected: Block{Lines: []string{"test"}},
		},
		{
			name:     "3 lines already sorted does not change",
			input:    Block{Lines: []string{"testA", "testB", "testC"}},
			expected: Block{Lines: []string{"testA", "testB", "testC"}},
		},
		{
			name:     "3 lines",
			input:    Block{Lines: []string{"testC", "testA", "testB"}},
			expected: Block{Lines: []string{"testA", "testB", "testC"}},
		},
		{
			name:     "other properties are not touched",
			input:    Block{Lines: nil, LineSeparator: "\t", TrailingSeparator: true},
			expected: Block{Lines: nil, LineSeparator: "\t", TrailingSeparator: true},
		},
	}

	assert := assertion.New(t)
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Reset(t)

			actual := tc.input
			sort.Sort(actual)

			assert.Var("sorted Block").Equal(tc.expected, actual)
		})
	}
}

func Test_Block_Append(t *testing.T) {
	testCases := []struct {
		name   string
		append string
		input  Block
		expect Block
	}{
		{
			name:   "append empty line to nil",
			append: "",
			input:  Block{Lines: nil},
			expect: Block{Lines: []string{""}},
		},
		{
			name:   "append empty line to empty Lines",
			append: "",
			input:  Block{Lines: []string{}},
			expect: Block{Lines: []string{""}},
		},
		{
			name:   "append empty line to default",
			append: "",
			input:  Block{},
			expect: Block{Lines: []string{""}},
		},
		{
			name:   "append filled line to default",
			append: "vriska",
			input:  Block{},
			expect: Block{Lines: []string{"vriska"}},
		},
		{
			name:   "append line with separator to default",
			append: "vriska\nserket\n",
			input:  Block{LineSeparator: "\n"},
			expect: Block{LineSeparator: "\n", Lines: []string{"vriska\nserket\n"}},
		},
		{
			name:   "append line to multiple Lines",
			append: "terezi",
			input:  Block{Lines: []string{"vriska", "roxy", "latula"}},
			expect: Block{Lines: []string{"vriska", "roxy", "latula", "terezi"}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert := assertion.New(t)

			actual := tc.input
			actual.Append(tc.append)

			assert.Equal(tc.expect, actual)
		})
	}
}

func Test_Block_AppendBlock(t *testing.T) {
	testCases := []struct {
		name   string
		input  Block
		append Block
		expect Block
	}{
		{
			name:   "append nil to nil",
			input:  Block{Lines: nil},
			append: Block{Lines: nil},
			expect: Block{},
		},
		{
			name:   "append nil to empty",
			append: Block{Lines: nil},
			input:  Block{Lines: []string{}},
			expect: Block{},
		},
		{
			name:   "append nil to default",
			append: Block{Lines: nil},
			input:  Block{},
			expect: Block{},
		},
		{
			name:   "append nil to 1",
			append: Block{Lines: nil},
			input:  Block{Lines: []string{"karkat"}},
			expect: Block{Lines: []string{"karkat"}},
		},
		{
			name:   "append nil to multi",
			append: Block{Lines: nil},
			input:  Block{Lines: []string{"karkat", "kanaya", "gamzee"}},
			expect: Block{Lines: []string{"karkat", "kanaya", "gamzee"}},
		},
		{
			name:   "append 1 to nil",
			append: Block{Lines: []string{"vriska"}},
			input:  Block{Lines: nil},
			expect: Block{Lines: []string{"vriska"}},
		},
		{
			name:   "append 1 to 1",
			append: Block{Lines: []string{"vriska"}},
			input:  Block{Lines: []string{"terezi"}},
			expect: Block{Lines: []string{"terezi", "vriska"}},
		},
		{
			name:   "append 1 to multi",
			append: Block{Lines: []string{"tavros"}},
			input:  Block{Lines: []string{"aradia", "sollux"}},
			expect: Block{Lines: []string{"aradia", "sollux", "tavros"}},
		},
		{
			name:   "append multi to nil",
			append: Block{Lines: []string{"equius", "nepeta", "eridan"}},
			input:  Block{Lines: nil},
			expect: Block{Lines: []string{"equius", "nepeta", "eridan"}},
		},
		{
			name:   "append multi to 1",
			append: Block{Lines: []string{"equius", "nepeta", "vriska"}},
			input:  Block{Lines: []string{"feferi"}},
			expect: Block{Lines: []string{"feferi", "equius", "nepeta", "vriska"}},
		},
		{
			name:   "append multi to multi",
			append: Block{Lines: []string{"nepeta", "kanaya", "aradia"}},
			input:  Block{Lines: []string{"feferi", "eridan", "equius"}},
			expect: Block{Lines: []string{"feferi", "eridan", "equius", "nepeta", "kanaya", "aradia"}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert := assertion.New(t)

			actual := tc.input
			actual.AppendBlock(tc.append)

			assert.Equal(tc.expect, actual)
		})
	}
}
