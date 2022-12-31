package tb

import (
	"sort"
	"testing"

	"github.com/dekarrin/rosed/internal/gem"

	"github.com/stretchr/testify/assert"
)

func Test_NewBlock(t *testing.T) {
	testCases := []struct {
		name     string
		text     gem.String
		sep      gem.String
		expected Block
	}{
		{
			name: "no lines - no trailing newline",
			text: gem.Zero,
			sep:  gem.New("\n"),
			expected: Block{
				Lines:             []gem.String{},
				LineSeparator:     gem.New("\n"),
				TrailingSeparator: false,
			},
		},
		{
			name: "no lines - trailing newline",
			text: gem.New("\n"),
			sep:  gem.New("\n"),
			expected: Block{
				Lines: []gem.String{
					gem.Zero,
				},
				LineSeparator:     gem.New("\n"),
				TrailingSeparator: true,
			},
		},
		{
			name: "one line - no trailing newline",
			text: gem.New("hello"),
			sep:  gem.New("\n"),
			expected: Block{
				Lines: []gem.String{
					gem.New("hello"),
				},
				LineSeparator:     gem.New("\n"),
				TrailingSeparator: false,
			},
		},
		{
			name: "one line - trailing newline",
			text: gem.New("hello\n"),
			sep:  gem.New("\n"),
			expected: Block{
				Lines: []gem.String{
					gem.New("hello"),
				},
				LineSeparator:     gem.New("\n"),
				TrailingSeparator: true,
			},
		},
		{
			name: "multi line - no trailing newline",
			text: gem.New("hello\nthere\ntest"),
			sep:  gem.New("\n"),
			expected: Block{
				Lines: []gem.String{
					gem.New("hello"),
					gem.New("there"),
					gem.New("test"),
				},
				LineSeparator:     gem.New("\n"),
				TrailingSeparator: false,
			},
		},
		{
			name: "multi line - trailing newline",
			text: gem.New("hello\nthere\ntest\n"),
			sep:  gem.New("\n"),
			expected: Block{
				Lines: []gem.String{
					gem.New("hello"),
					gem.New("there"),
					gem.New("test"),
				},
				LineSeparator:     gem.New("\n"),
				TrailingSeparator: true,
			},
		},
		{
			name: "multi line - no trailing separator - alternate separator",
			text: gem.New("john\negbert\trose\nlalonde\tdave\nstrider\tjade\nharley"),
			sep:  gem.New("\t"),
			expected: Block{
				Lines: []gem.String{
					gem.New("john\negbert"),
					gem.New("rose\nlalonde"),
					gem.New("dave\nstrider"),
					gem.New("jade\nharley"),
				},
				LineSeparator:     gem.New("\t"),
				TrailingSeparator: false,
			},
		},
		{
			name: "multi line - trailing separator - alternate separator",
			text: gem.New("john\negbert\trose\nlalonde\tdave\nstrider\tjade\nharley\t"),
			sep:  gem.New("\t"),
			expected: Block{
				Lines: []gem.String{
					gem.New("john\negbert"),
					gem.New("rose\nlalonde"),
					gem.New("dave\nstrider"),
					gem.New("jade\nharley"),
				},
				LineSeparator:     gem.New("\t"),
				TrailingSeparator: true,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert := assert.New(t)

			actual := New(tc.text, tc.sep)

			assert.True(tc.expected.Equal(actual))
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
			b1:       Block{Lines: []gem.String{}},
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
			b1:       Block{Lines: []gem.String{}},
			b2:       Block{Lines: []gem.String{}},
			expected: true,
		},
		{
			name:     "default == empty Lines",
			b1:       Block{},
			b2:       Block{Lines: []gem.String{}},
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
			b1:       Block{Lines: []gem.String{}},
			b2:       Block{Lines: []gem.String{gem.Zero}},
			expected: false,
		},
		{
			name:     "nil Lines != 1 empty line",
			b1:       Block{Lines: nil},
			b2:       Block{Lines: []gem.String{gem.Zero}},
			expected: false,
		},
		{
			name:     "default != 1 empty line",
			b1:       Block{},
			b2:       Block{Lines: []gem.String{gem.Zero}},
			expected: false,
		},
		{
			name:     "1 empty line == 1 empty line",
			b1:       Block{Lines: []gem.String{gem.Zero}},
			b2:       Block{Lines: []gem.String{gem.Zero}},
			expected: true,
		},
		{
			name:     "1 filled line == same filled line",
			b1:       Block{Lines: []gem.String{gem.New("test")}},
			b2:       Block{Lines: []gem.String{gem.New("test")}},
			expected: true,
		},
		{
			name:     "1 filled line != different filled line",
			b1:       Block{Lines: []gem.String{gem.New("test1")}},
			b2:       Block{Lines: []gem.String{gem.New("test2")}},
			expected: false,
		},
		{
			name:     "3 empty lines != 4 empty lines",
			b1:       Block{Lines: []gem.String{gem.Zero, gem.Zero, gem.Zero}},
			b2:       Block{Lines: []gem.String{gem.Zero, gem.Zero, gem.Zero, gem.Zero}},
			expected: false,
		},
		{
			name:     "3 empty lines == 3 empty lines",
			b1:       Block{Lines: []gem.String{gem.Zero, gem.Zero, gem.Zero}},
			b2:       Block{Lines: []gem.String{gem.Zero, gem.Zero, gem.Zero}},
			expected: true,
		},
		{
			name:     "3 filled lines == same 3 filled lines",
			b1:       Block{Lines: []gem.String{gem.New("a"), gem.New("ab"), gem.New("abc")}},
			b2:       Block{Lines: []gem.String{gem.New("a"), gem.New("ab"), gem.New("abc")}},
			expected: true,
		},
		{
			name:     "3 filled lines != 4 filled lines",
			b1:       Block{Lines: []gem.String{gem.New("a"), gem.New("ab"), gem.New("abc")}},
			b2:       Block{Lines: []gem.String{gem.New("a"), gem.New("ab"), gem.New("abc"), gem.New("abcd")}},
			expected: false,
		},
		{
			name:     "different separators",
			b1:       Block{LineSeparator: gem.New("\n")},
			b2:       Block{LineSeparator: gem.New("\t")},
			expected: false,
		},
		{
			name:     "different newline behavior",
			b1:       Block{TrailingSeparator: true},
			b2:       Block{TrailingSeparator: false},
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert := assert.New(t)
			assert.Equal(tc.expected, tc.b1.Equal(tc.b2), "b1.Equal(b2) check failed")
			assert.Equal(tc.expected, tc.b2.Equal(tc.b1), "b2.Equal(b1) check failed")
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
		{"empty Lines", Block{Lines: []gem.String{}}, 0},
		{"default Lines", Block{}, 0},
		{"1 empty line", Block{Lines: []gem.String{gem.Zero}}, 1},
		{"1 filled line", Block{Lines: []gem.String{gem.New("test")}}, 1},
		{"3 empty lines", Block{Lines: []gem.String{gem.Zero, gem.Zero, gem.Zero}}, 3},
		{"3 filled lines", Block{Lines: []gem.String{gem.New("a"), gem.New("b"), gem.New("c")}}, 3},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert := assert.New(t)

			actual := tc.input.Len()

			assert.Equal(tc.expected, actual)
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
			input:    Block{Lines: []gem.String{}},
			expected: Block{Lines: []gem.String{}},
		},
		{
			name:     "default does not change",
			input:    Block{},
			expected: Block{},
		},
		{
			name:     "1 line does not change",
			input:    Block{Lines: []gem.String{gem.New("test")}},
			expected: Block{Lines: []gem.String{gem.New("test")}},
		},
		{
			name:     "3 lines already sorted does not change",
			input:    Block{Lines: []gem.String{gem.New("testA"), gem.New("testB"), gem.New("testC")}},
			expected: Block{Lines: []gem.String{gem.New("testA"), gem.New("testB"), gem.New("testC")}},
		},
		{
			name:     "3 lines",
			input:    Block{Lines: []gem.String{gem.New("testC"), gem.New("testA"), gem.New("testB")}},
			expected: Block{Lines: []gem.String{gem.New("testA"), gem.New("testB"), gem.New("testC")}},
		},
		{
			name:     "other properties are not touched",
			input:    Block{Lines: nil, LineSeparator: gem.New("\t"), TrailingSeparator: true},
			expected: Block{Lines: nil, LineSeparator: gem.New("\t"), TrailingSeparator: true},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert := assert.New(t)

			actual := tc.input
			sort.Sort(actual)

			assert.Equal(tc.expected, actual)
		})
	}
}

func Test_Block_Append(t *testing.T) {
	testCases := []struct {
		name   string
		append gem.String
		input  Block
		expect Block
	}{
		{
			name:   "append empty line to nil",
			append: gem.Zero,
			input:  Block{Lines: nil},
			expect: Block{Lines: []gem.String{gem.Zero}},
		},
		{
			name:   "append empty line to empty Lines",
			append: gem.Zero,
			input:  Block{Lines: []gem.String{}},
			expect: Block{Lines: []gem.String{gem.Zero}},
		},
		{
			name:   "append empty line to default",
			append: gem.Zero,
			input:  Block{},
			expect: Block{Lines: []gem.String{gem.Zero}},
		},
		{
			name:   "append filled line to default",
			append: gem.New("vriska"),
			input:  Block{},
			expect: Block{Lines: []gem.String{gem.New("vriska")}},
		},
		{
			name:   "append line with separator to default",
			append: gem.New("vriska\nserket\n"),
			input:  Block{LineSeparator: gem.New("\n")},
			expect: Block{LineSeparator: gem.New("\n"), Lines: []gem.String{gem.New("vriska\nserket\n")}},
		},
		{
			name:   "append line to multiple Lines",
			append: gem.New("terezi"),
			input:  Block{Lines: []gem.String{gem.New("vriska"), gem.New("roxy"), gem.New("latula")}},
			expect: Block{Lines: []gem.String{gem.New("vriska"), gem.New("roxy"), gem.New("latula"), gem.New("terezi")}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert := assert.New(t)

			actual := tc.input
			actual.Append(tc.append)

			assert.Equal(tc.expect, actual)
		})
	}
}

func Test_Block_AppendBlock(t *testing.T) {
	testCases := []struct {
		name   string
		append Block
		input  Block
		expect Block
	}{
		{
			name:   "append nil to nil",
			append: Block{Lines: nil},
			input:  Block{Lines: nil},
			expect: Block{},
		},
		{
			name:   "append nil to empty",
			append: Block{Lines: nil},
			input:  Block{Lines: []gem.String{}},
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
			input:  Block{Lines: []gem.String{gem.New("karkat")}},
			expect: Block{Lines: []gem.String{gem.New("karkat")}},
		},
		{
			name:   "append nil to multi",
			append: Block{Lines: nil},
			input:  Block{Lines: []gem.String{gem.New("karkat"), gem.New("kanaya"), gem.New("gamzee")}},
			expect: Block{Lines: []gem.String{gem.New("karkat"), gem.New("kanaya"), gem.New("gamzee")}},
		},
		{
			name:   "append 1 to nil",
			append: Block{Lines: []gem.String{gem.New("vriska")}},
			input:  Block{Lines: nil},
			expect: Block{Lines: []gem.String{gem.New("vriska")}},
		},
		{
			name:   "append 1 to 1",
			append: Block{Lines: []gem.String{gem.New("vriska")}},
			input:  Block{Lines: []gem.String{gem.New("terezi")}},
			expect: Block{Lines: []gem.String{gem.New("terezi"), gem.New("vriska")}},
		},
		{
			name:   "append 1 to multi",
			append: Block{Lines: []gem.String{gem.New("tavros")}},
			input:  Block{Lines: []gem.String{gem.New("aradia"), gem.New("sollux")}},
			expect: Block{Lines: []gem.String{gem.New("aradia"), gem.New("sollux"), gem.New("tavros")}},
		},
		{
			name:   "append multi to nil",
			append: Block{Lines: []gem.String{gem.New("equius"), gem.New("nepeta"), gem.New("eridan")}},
			input:  Block{Lines: nil},
			expect: Block{Lines: []gem.String{gem.New("equius"), gem.New("nepeta"), gem.New("eridan")}},
		},
		{
			name:   "append multi to 1",
			append: Block{Lines: []gem.String{gem.New("equius"), gem.New("nepeta"), gem.New("vriska")}},
			input:  Block{Lines: []gem.String{gem.New("feferi")}},
			expect: Block{Lines: []gem.String{gem.New("feferi"), gem.New("equius"), gem.New("nepeta"), gem.New("vriska")}},
		},
		{
			name:   "append multi to multi",
			append: Block{Lines: []gem.String{gem.New("nepeta"), gem.New("kanaya"), gem.New("aradia")}},
			input:  Block{Lines: []gem.String{gem.New("feferi"), gem.New("eridan"), gem.New("equius")}},
			expect: Block{Lines: []gem.String{gem.New("feferi"), gem.New("eridan"), gem.New("equius"), gem.New("nepeta"), gem.New("kanaya"), gem.New("aradia")}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert := assert.New(t)

			actual := tc.input
			actual.AppendBlock(tc.append)

			assert.True(tc.expect.Equal(actual))
		})
	}
}

func Test_Block_AppendEmpty(t *testing.T) {
	testCases := []struct {
		name   string
		count  int
		input  Block
		expect Block
	}{
		{"append 0 to nil", 0, Block{Lines: nil}, Block{}},
		{"append 0 to empty", 0, Block{Lines: []gem.String{}}, Block{}},
		{"append 0 to default", 0, Block{}, Block{}},
		{"append 1 to nil", 1, Block{Lines: nil}, Block{Lines: []gem.String{gem.Zero}}},
		{"append 1 to empty", 1, Block{Lines: []gem.String{}}, Block{Lines: []gem.String{gem.Zero}}},
		{"append 1 to default", 1, Block{}, Block{Lines: []gem.String{gem.Zero}}},
		{"append 1 to 1", 1, Block{Lines: []gem.String{gem.New("vriska")}}, Block{Lines: []gem.String{gem.New("vriska"), gem.Zero}}},
		{"append 1 to many", 1, Block{Lines: []gem.String{gem.New("vriska"), gem.New("terezi")}}, Block{Lines: []gem.String{gem.New("vriska"), gem.New("terezi"), gem.Zero}}},
		{"append 3 to 1", 3, Block{Lines: []gem.String{gem.New("vriska")}}, Block{Lines: []gem.String{gem.New("vriska"), gem.Zero, gem.Zero, gem.Zero}}},
		{"append 3 to many", 3, Block{Lines: []gem.String{gem.New("vriska"), gem.New("terezi")}}, Block{Lines: []gem.String{gem.New("vriska"), gem.New("terezi"), gem.Zero, gem.Zero, gem.Zero}}},
		{"append -1 to default", -1, Block{}, Block{}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert := assert.New(t)

			actual := tc.input
			actual.AppendEmpty(tc.count)

			assert.True(tc.expect.Equal(actual))
		})
	}
}

func Test_Block_Set(t *testing.T) {
	type args struct {
		pos  int
		text gem.String
	}

	testCases := []struct {
		name   string
		args   args
		input  Block
		expect Block
		panics bool
	}{
		{
			name:   "index too low causes panic",
			args:   args{-1, gem.Zero},
			input:  Block{},
			panics: true,
		},
		{
			name:   "index too high causes panic",
			args:   args{20, gem.Zero},
			input:  Block{Lines: []gem.String{gem.Zero, gem.Zero}},
			panics: true,
		},
		{
			name:   "set 0th line",
			args:   args{0, gem.New("new text")},
			input:  Block{Lines: []gem.String{gem.New("old text"), gem.New("test")}},
			expect: Block{Lines: []gem.String{gem.New("new text"), gem.New("test")}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert := assert.New(t)

			actual := tc.input
			if tc.panics {
				assert.Panics(func() { actual.Set(tc.args.pos, tc.args.text) })
			} else {
				actual.Set(tc.args.pos, tc.args.text)
				assert.Equal(tc.expect, actual)
			}
		})
	}
}

func Test_Block_Line(t *testing.T) {
	testCases := []struct {
		name   string
		pos    int
		input  Block
		expect gem.String
		panics bool
	}{
		{
			name:   "index too low causes panic",
			pos:    0,
			input:  Block{},
			panics: true,
		},
		{
			name:   "index too high causes panic",
			pos:    2,
			input:  Block{Lines: []gem.String{gem.Zero, gem.Zero}},
			panics: true,
		},
		{
			name:   "get 0th line",
			pos:    0,
			input:  Block{Lines: []gem.String{gem.New("old text"), gem.New("test")}},
			expect: gem.New("old text"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert := assert.New(t)

			if tc.panics {
				assert.Panics(func() { tc.input.Line(tc.pos) })
			} else {
				actual := tc.input.Line(tc.pos)
				assert.Equal(tc.expect, actual)
			}
		})
	}
}

func Test_Block_Remove(t *testing.T) {
	testCases := []struct {
		name   string
		pos    int
		input  Block
		expect Block
	}{
		{
			name:   "remove from empty has no effect",
			pos:    0,
			input:  Block{},
			expect: Block{},
		},
		{
			name: "too-high pos has no effect",
			pos:  7,
			input: Block{Lines: []gem.String{
				gem.New("vriska"),
				gem.New("terezi"),
				gem.New("kanaya"),
				gem.New("wind"),
				gem.New("mari"),
				gem.New("deka"),
				gem.New("equius"),
			}},
			expect: Block{Lines: []gem.String{
				gem.New("vriska"),
				gem.New("terezi"),
				gem.New("kanaya"),
				gem.New("wind"),
				gem.New("mari"),
				gem.New("deka"),
				gem.New("equius"),
			}},
		},
		{
			name: "negative pos has no effect",
			pos:  -1,
			input: Block{Lines: []gem.String{
				gem.New("vriska"),
				gem.New("terezi"),
				gem.New("kanaya"),
				gem.New("wind"),
				gem.New("mari"),
				gem.New("deka"),
				gem.New("equius"),
			}},
			expect: Block{Lines: []gem.String{
				gem.New("vriska"),
				gem.New("terezi"),
				gem.New("kanaya"),
				gem.New("wind"),
				gem.New("mari"),
				gem.New("deka"),
				gem.New("equius"),
			}},
		},
		{
			name: "remove the only line that exists",
			pos:  0,
			input: Block{Lines: []gem.String{
				gem.New("line"),
			}},
			expect: Block{Lines: []gem.String{}},
		},
		{
			name: "remove from start",
			pos:  0,
			input: Block{Lines: []gem.String{
				gem.New("line1"),
				gem.New("line2"),
				gem.New("line3"),
				gem.New("line4"),
				gem.New("line5"),
			}},
			expect: Block{Lines: []gem.String{
				gem.New("line2"),
				gem.New("line3"),
				gem.New("line4"),
				gem.New("line5"),
			}},
		},
		{
			name: "remove from end",
			pos:  4,
			input: Block{Lines: []gem.String{
				gem.New("line1"),
				gem.New("line2"),
				gem.New("line3"),
				gem.New("line4"),
				gem.New("line5"),
			}},
			expect: Block{Lines: []gem.String{
				gem.New("line1"),
				gem.New("line2"),
				gem.New("line3"),
				gem.New("line4"),
			}},
		},
		{
			name: "remove from middle",
			pos:  2,
			input: Block{Lines: []gem.String{
				gem.New("line1"),
				gem.New("line2"),
				gem.New("line3"),
				gem.New("line4"),
				gem.New("line5"),
			}},
			expect: Block{Lines: []gem.String{
				gem.New("line1"),
				gem.New("line2"),
				gem.New("line4"),
				gem.New("line5"),
			}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert := assert.New(t)

			tc.input.Remove(tc.pos)
			actual := tc.input

			assert.Equal(tc.expect, actual)
		})
	}
}

func Test_Block_CharCount(t *testing.T) {
	testCases := []struct {
		name   string
		pos    int
		input  Block
		expect int
		panics bool
	}{
		{
			name:   "index too low causes panic",
			pos:    0,
			input:  Block{},
			panics: true,
		},
		{
			name:   "index too high causes panic",
			pos:    2,
			input:  Block{Lines: []gem.String{gem.Zero, gem.Zero}},
			panics: true,
		},
		{
			name:   "empty string",
			pos:    0,
			input:  Block{Lines: []gem.String{gem.Zero}},
			expect: 0,
		},
		{
			name:   "latin-1",
			pos:    0,
			input:  Block{Lines: []gem.String{gem.New("test")}},
			expect: 4,
		},
		{
			name:   "japanese",
			pos:    0,
			input:  Block{Lines: []gem.String{gem.New("こんにちは世界")}},
			expect: 7,
		},
		{
			name:   "russian",
			pos:    0,
			input:  Block{Lines: []gem.String{gem.New("При́пять")}},
			expect: 7,
		},
		{
			name:   "arabic",
			pos:    0,
			input:  Block{Lines: []gem.String{gem.New("الخوارزمية")}},
			expect: 10,
		},
		{
			name:   "emoji",
			pos:    0,
			input:  Block{Lines: []gem.String{gem.New("😍😎😑😐😏")}},
			expect: 5,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert := assert.New(t)

			if tc.panics {
				assert.Panics(func() { tc.input.CharCount(tc.pos) })
			} else {
				actual := tc.input.CharCount(tc.pos)
				assert.Equal(tc.expect, actual)
			}
		})
	}
}

func Test_Block_Join(t *testing.T) {
	testCases := []struct {
		name   string
		input  Block
		expect gem.String
	}{
		{
			name: "join with some empty, trailing separator",
			input: Block{
				Lines: []gem.String{
					gem.New("test1"),
					gem.Zero,
					gem.New("test2"),
					gem.Zero,
					gem.Zero,
					gem.New("test3"),
				},
				LineSeparator:     gem.New("\n"),
				TrailingSeparator: true,
			},
			expect: gem.New("test1\n\ntest2\n\n\ntest3\n"),
		},
		{
			name: "join with some empty, no trailing separator",
			input: Block{
				Lines: []gem.String{
					gem.New("test1"),
					gem.Zero,
					gem.New("test2"),
					gem.Zero,
					gem.Zero,
					gem.New("test3"),
				},
				LineSeparator: gem.New("\n"),
			},
			expect: gem.New("test1\n\ntest2\n\n\ntest3"),
		},
		{
			name: "join 3 lines, trailing separator",
			input: Block{
				Lines: []gem.String{
					gem.New("test1"),
					gem.New("test2"),
					gem.New("test3"),
				},
				LineSeparator:     gem.New("\n"),
				TrailingSeparator: true,
			},
			expect: gem.New("test1\ntest2\ntest3\n"),
		},
		{
			name: "join 3 lines, no trailing separator",
			input: Block{
				Lines: []gem.String{
					gem.New("test1"),
					gem.New("test2"),
					gem.New("test3"),
				},
				LineSeparator:     gem.New("\n"),
				TrailingSeparator: false,
			},
			expect: gem.New("test1\ntest2\ntest3"),
		},
		{
			name: "join 3 lines, alternate separator",
			input: Block{
				Lines: []gem.String{
					gem.New("test1"),
					gem.New("test2"),
					gem.New("test3"),
				},
				LineSeparator:     gem.New("_"),
				TrailingSeparator: false,
			},
			expect: gem.New("test1_test2_test3"),
		},
		{
			name: "join nil lines",
			input: Block{
				Lines:         nil,
				LineSeparator: gem.New("\n"),
			},
			expect: gem.Zero,
		},
		{
			name: "join nil lines, line terminator on",
			input: Block{
				Lines:             nil,
				LineSeparator:     gem.New("\n"),
				TrailingSeparator: true,
			},
			expect: gem.New("\n"),
		},
		{
			name: "join empty lines",
			input: Block{
				Lines:         []gem.String{},
				LineSeparator: gem.New("\n"),
			},
			expect: gem.Zero,
		},
		{
			name: "join empty lines, line terminator on",
			input: Block{
				Lines:             []gem.String{},
				LineSeparator:     gem.New("\n"),
				TrailingSeparator: true,
			},
			expect: gem.New("\n"),
		},
		{
			name: "join 1 empty line",
			input: Block{
				Lines:         []gem.String{gem.Zero},
				LineSeparator: gem.New("\n"),
			},
			expect: gem.Zero,
		},
		{
			name: "join 1 empty line, line terminator on",
			input: Block{
				Lines:             []gem.String{gem.Zero},
				LineSeparator:     gem.New("\n"),
				TrailingSeparator: true,
			},
			expect: gem.New("\n"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert := assert.New(t)
			actual := tc.input.Join()
			assert.True(tc.expect.Equal(actual))
		})
	}
}
