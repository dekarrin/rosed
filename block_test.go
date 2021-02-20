package rosed

import (
	"sort"
	"testing"

	"github.com/dekarrin/assertion"
	"github.com/dekarrin/rosed/internal/gem"
)

func Test_NewBlock(t *testing.T) {
	testCases := []struct {
		name     string
		text     gem.String
		sep      gem.String
		expected block
	}{
		{
			name: "no lines - no trailing newline",
			text: gem.Empty,
			sep:  gem.New("\n"),
			expected: block{
				Lines:             []gem.String{},
				LineSeparator:     gem.New("\n"),
				TrailingSeparator: false,
			},
		},
		{
			name: "no lines - trailing newline",
			text: gem.New("\n"),
			sep:  gem.New("\n"),
			expected: block{
				Lines: []gem.String{
					gem.Empty,
				},
				LineSeparator:     gem.New("\n"),
				TrailingSeparator: true,
			},
		},
		{
			name: "one line - no trailing newline",
			text: gem.New("hello"),
			sep:  gem.New("\n"),
			expected: block{
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
			expected: block{
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
			expected: block{
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
			expected: block{
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
			expected: block{
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
			expected: block{
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

	assert := assertion.New(t)
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Reset(t)

			actual := newBlock(tc.text, tc.sep)

			assert.Var("block").Equal(tc.expected, actual)
		})
	}
}

func Test_Block_Equal(t *testing.T) {
	testCases := []struct {
		name     string
		b1       block
		b2       block
		expected bool
	}{
		{
			name:     "nil Lines == empty Lines",
			b1:       block{Lines: []gem.String{}},
			b2:       block{Lines: nil},
			expected: true,
		},
		{
			name:     "nil Lines == nil Lines",
			b1:       block{Lines: nil},
			b2:       block{Lines: nil},
			expected: true,
		},
		{
			name:     "empty Lines == empty Lines",
			b1:       block{Lines: []gem.String{}},
			b2:       block{Lines: []gem.String{}},
			expected: true,
		},
		{
			name:     "default == empty Lines",
			b1:       block{},
			b2:       block{Lines: []gem.String{}},
			expected: true,
		},
		{
			name:     "default == nil Lines",
			b1:       block{},
			b2:       block{Lines: nil},
			expected: true,
		},
		{
			name:     "default == default",
			b1:       block{},
			b2:       block{},
			expected: true,
		},
		{
			name:     "empty Lines != 1 empty line",
			b1:       block{Lines: []gem.String{}},
			b2:       block{Lines: []gem.String{gem.Empty}},
			expected: false,
		},
		{
			name:     "nil Lines != 1 empty line",
			b1:       block{Lines: nil},
			b2:       block{Lines: []gem.String{gem.Empty}},
			expected: false,
		},
		{
			name:     "default != 1 empty line",
			b1:       block{},
			b2:       block{Lines: []gem.String{gem.Empty}},
			expected: false,
		},
		{
			name:     "1 empty line == 1 empty line",
			b1:       block{Lines: []gem.String{gem.Empty}},
			b2:       block{Lines: []gem.String{gem.Empty}},
			expected: true,
		},
		{
			name:     "1 filled line == same filled line",
			b1:       block{Lines: []gem.String{gem.New("test")}},
			b2:       block{Lines: []gem.String{gem.New("test")}},
			expected: true,
		},
		{
			name:     "1 filled line != different filled line",
			b1:       block{Lines: []gem.String{gem.New("test1")}},
			b2:       block{Lines: []gem.String{gem.New("test2")}},
			expected: false,
		},
		{
			name:     "3 empty lines != 4 empty lines",
			b1:       block{Lines: []gem.String{gem.Empty, gem.Empty, gem.Empty}},
			b2:       block{Lines: []gem.String{gem.Empty, gem.Empty, gem.Empty, gem.Empty}},
			expected: false,
		},
		{
			name:     "3 empty lines == 3 empty lines",
			b1:       block{Lines: []gem.String{gem.Empty, gem.Empty, gem.Empty}},
			b2:       block{Lines: []gem.String{gem.Empty, gem.Empty, gem.Empty}},
			expected: true,
		},
		{
			name:     "3 filled lines == same 3 filled lines",
			b1:       block{Lines: []gem.String{gem.New("a"), gem.New("ab"), gem.New("abc")}},
			b2:       block{Lines: []gem.String{gem.New("a"), gem.New("ab"), gem.New("abc")}},
			expected: true,
		},
		{
			name:     "3 filled lines != 4 filled lines",
			b1:       block{Lines: []gem.String{gem.New("a"), gem.New("ab"), gem.New("abc")}},
			b2:       block{Lines: []gem.String{gem.New("a"), gem.New("ab"), gem.New("abc"), gem.New("abcd")}},
			expected: false,
		},
		{
			name:     "different separators",
			b1:       block{LineSeparator: gem.New("\n")},
			b2:       block{LineSeparator: gem.New("\t")},
			expected: false,
		},
		{
			name:     "different newline behavior",
			b1:       block{TrailingSeparator: true},
			b2:       block{TrailingSeparator: false},
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
		input    block
		expected int
	}{
		{"nil Lines", block{Lines: nil}, 0},
		{"empty Lines", block{Lines: []gem.String{}}, 0},
		{"default Lines", block{}, 0},
		{"1 empty line", block{Lines: []gem.String{gem.Empty}}, 1},
		{"1 filled line", block{Lines: []gem.String{gem.New("test")}}, 1},
		{"3 empty lines", block{Lines: []gem.String{gem.Empty, gem.Empty, gem.Empty}}, 3},
		{"3 filled lines", block{Lines: []gem.String{gem.New("a"), gem.New("b"), gem.New("c")}}, 3},
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
		input    block
		expected block
	}{
		{
			name:     "nil Lines does not change",
			input:    block{Lines: nil},
			expected: block{Lines: nil},
		},
		{
			name:     "empty Lines does not change",
			input:    block{Lines: []gem.String{}},
			expected: block{Lines: []gem.String{}},
		},
		{
			name:     "default does not change",
			input:    block{},
			expected: block{},
		},
		{
			name:     "1 line does not change",
			input:    block{Lines: []gem.String{gem.New("test")}},
			expected: block{Lines: []gem.String{gem.New("test")}},
		},
		{
			name:     "3 lines already sorted does not change",
			input:    block{Lines: []gem.String{gem.New("testA"), gem.New("testB"), gem.New("testC")}},
			expected: block{Lines: []gem.String{gem.New("testA"), gem.New("testB"), gem.New("testC")}},
		},
		{
			name:     "3 lines",
			input:    block{Lines: []gem.String{gem.New("testC"), gem.New("testA"), gem.New("testB")}},
			expected: block{Lines: []gem.String{gem.New("testA"), gem.New("testB"), gem.New("testC")}},
		},
		{
			name:     "other properties are not touched",
			input:    block{Lines: nil, LineSeparator: gem.New("\t"), TrailingSeparator: true},
			expected: block{Lines: nil, LineSeparator: gem.New("\t"), TrailingSeparator: true},
		},
	}

	assert := assertion.New(t)
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Reset(t)

			actual := tc.input
			sort.Sort(actual)

			assert.Var("sorted block").Equal(tc.expected, actual)
		})
	}
}

func Test_Block_Append(t *testing.T) {
	testCases := []struct {
		name   string
		append gem.String
		input  block
		expect block
	}{
		{
			name:   "append empty line to nil",
			append: gem.Empty,
			input:  block{Lines: nil},
			expect: block{Lines: []gem.String{gem.Empty}},
		},
		{
			name:   "append empty line to empty Lines",
			append: gem.Empty,
			input:  block{Lines: []gem.String{}},
			expect: block{Lines: []gem.String{gem.Empty}},
		},
		{
			name:   "append empty line to default",
			append: gem.Empty,
			input:  block{},
			expect: block{Lines: []gem.String{gem.Empty}},
		},
		{
			name:   "append filled line to default",
			append: gem.New("vriska"),
			input:  block{},
			expect: block{Lines: []gem.String{gem.New("vriska")}},
		},
		{
			name:   "append line with separator to default",
			append: gem.New("vriska\nserket\n"),
			input:  block{LineSeparator: gem.New("\n")},
			expect: block{LineSeparator: gem.New("\n"), Lines: []gem.String{gem.New("vriska\nserket\n")}},
		},
		{
			name:   "append line to multiple Lines",
			append: gem.New("terezi"),
			input:  block{Lines: []gem.String{gem.New("vriska"), gem.New("roxy"), gem.New("latula")}},
			expect: block{Lines: []gem.String{gem.New("vriska"), gem.New("roxy"), gem.New("latula"), gem.New("terezi")}},
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
		append block
		input  block
		expect block
	}{
		{
			name:   "append nil to nil",
			append: block{Lines: nil},
			input:  block{Lines: nil},
			expect: block{},
		},
		{
			name:   "append nil to empty",
			append: block{Lines: nil},
			input:  block{Lines: []gem.String{}},
			expect: block{},
		},
		{
			name:   "append nil to default",
			append: block{Lines: nil},
			input:  block{},
			expect: block{},
		},
		{
			name:   "append nil to 1",
			append: block{Lines: nil},
			input:  block{Lines: []gem.String{gem.New("karkat")}},
			expect: block{Lines: []gem.String{gem.New("karkat")}},
		},
		{
			name:   "append nil to multi",
			append: block{Lines: nil},
			input:  block{Lines: []gem.String{gem.New("karkat"), gem.New("kanaya"), gem.New("gamzee")}},
			expect: block{Lines: []gem.String{gem.New("karkat"), gem.New("kanaya"), gem.New("gamzee")}},
		},
		{
			name:   "append 1 to nil",
			append: block{Lines: []gem.String{gem.New("vriska")}},
			input:  block{Lines: nil},
			expect: block{Lines: []gem.String{gem.New("vriska")}},
		},
		{
			name:   "append 1 to 1",
			append: block{Lines: []gem.String{gem.New("vriska")}},
			input:  block{Lines: []gem.String{gem.New("terezi")}},
			expect: block{Lines: []gem.String{gem.New("terezi"), gem.New("vriska")}},
		},
		{
			name:   "append 1 to multi",
			append: block{Lines: []gem.String{gem.New("tavros")}},
			input:  block{Lines: []gem.String{gem.New("aradia"), gem.New("sollux")}},
			expect: block{Lines: []gem.String{gem.New("aradia"), gem.New("sollux"), gem.New("tavros")}},
		},
		{
			name:   "append multi to nil",
			append: block{Lines: []gem.String{gem.New("equius"), gem.New("nepeta"), gem.New("eridan")}},
			input:  block{Lines: nil},
			expect: block{Lines: []gem.String{gem.New("equius"), gem.New("nepeta"), gem.New("eridan")}},
		},
		{
			name:   "append multi to 1",
			append: block{Lines: []gem.String{gem.New("equius"), gem.New("nepeta"), gem.New("vriska")}},
			input:  block{Lines: []gem.String{gem.New("feferi")}},
			expect: block{Lines: []gem.String{gem.New("feferi"), gem.New("equius"), gem.New("nepeta"), gem.New("vriska")}},
		},
		{
			name:   "append multi to multi",
			append: block{Lines: []gem.String{gem.New("nepeta"), gem.New("kanaya"), gem.New("aradia")}},
			input:  block{Lines: []gem.String{gem.New("feferi"), gem.New("eridan"), gem.New("equius")}},
			expect: block{Lines: []gem.String{gem.New("feferi"), gem.New("eridan"), gem.New("equius"), gem.New("nepeta"), gem.New("kanaya"), gem.New("aradia")}},
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

func Test_Block_AppendEmpty(t *testing.T) {
	testCases := []struct {
		name   string
		count  int
		input  block
		expect block
	}{
		{"append 0 to nil", 0, block{Lines: nil}, block{}},
		{"append 0 to empty", 0, block{Lines: []gem.String{}}, block{}},
		{"append 0 to default", 0, block{}, block{}},
		{"append 1 to nil", 1, block{Lines: nil}, block{Lines: []gem.String{gem.Empty}}},
		{"append 1 to empty", 1, block{Lines: []gem.String{}}, block{Lines: []gem.String{gem.Empty}}},
		{"append 1 to default", 1, block{}, block{Lines: []gem.String{gem.Empty}}},
		{"append 1 to 1", 1, block{Lines: []gem.String{gem.New("vriska")}}, block{Lines: []gem.String{gem.New("vriska"), gem.Empty}}},
		{"append 1 to many", 1, block{Lines: []gem.String{gem.New("vriska"), gem.New("terezi")}}, block{Lines: []gem.String{gem.New("vriska"), gem.New("terezi"), gem.Empty}}},
		{"append 3 to 1", 3, block{Lines: []gem.String{gem.New("vriska")}}, block{Lines: []gem.String{gem.New("vriska"), gem.Empty, gem.Empty, gem.Empty}}},
		{"append 3 to many", 3, block{Lines: []gem.String{gem.New("vriska"), gem.New("terezi")}}, block{Lines: []gem.String{gem.New("vriska"), gem.New("terezi"), gem.Empty, gem.Empty, gem.Empty}}},
		{"append -1 to default", -1, block{}, block{}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert := assertion.New(t)

			actual := tc.input
			actual.AppendEmpty(tc.count)

			assert.Equal(tc.expect, actual)
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
		input  block
		expect block
		panics bool
	}{
		{
			name:   "index too low causes panic",
			args:   args{-1, gem.Empty},
			input:  block{},
			panics: true,
		},
		{
			name:   "index too high causes panic",
			args:   args{20, gem.Empty},
			input:  block{Lines: []gem.String{gem.Empty, gem.Empty}},
			panics: true,
		},
		{
			name:   "set 0th line",
			args:   args{0, gem.New("new text")},
			input:  block{Lines: []gem.String{gem.New("old text"), gem.New("test")}},
			expect: block{Lines: []gem.String{gem.New("new text"), gem.New("test")}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert := assertion.New(t)

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
		input  block
		expect string
		panics bool
	}{
		{
			name:   "index too low causes panic",
			pos:    0,
			input:  block{},
			panics: true,
		},
		{
			name:   "index too high causes panic",
			pos:    2,
			input:  block{Lines: []gem.String{gem.Empty, gem.Empty}},
			panics: true,
		},
		{
			name:   "get 0th line",
			pos:    0,
			input:  block{Lines: []gem.String{gem.New("old text"), gem.New("test")}},
			expect: "old text",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert := assertion.New(t)

			if tc.panics {
				assert.Panics(func() { tc.input.Line(tc.pos) })
			} else {
				actual := tc.input.Line(tc.pos)
				assert.Equal(tc.expect, actual)
			}
		})
	}
}

func Test_Block_CharCount(t *testing.T) {
	testCases := []struct {
		name   string
		pos    int
		input  block
		expect int
		panics bool
	}{
		{
			name:   "index too low causes panic",
			pos:    0,
			input:  block{},
			panics: true,
		},
		{
			name:   "index too high causes panic",
			pos:    2,
			input:  block{Lines: []gem.String{gem.Empty, gem.Empty}},
			panics: true,
		},
		{
			name:   "empty string",
			pos:    0,
			input:  block{Lines: []gem.String{gem.Empty}},
			expect: 0,
		},
		{
			name:   "latin-1",
			pos:    0,
			input:  block{Lines: []gem.String{gem.New("test")}},
			expect: 4,
		},
		{
			name:   "japanese",
			pos:    0,
			input:  block{Lines: []gem.String{gem.New("こんにちは世界")}},
			expect: 7,
		},
		{
			name:   "russian",
			pos:    0,
			input:  block{Lines: []gem.String{gem.New("При́пять")}},
			expect: 7,
		},
		{
			name:   "arabic",
			pos:    0,
			input:  block{Lines: []gem.String{gem.New("الخوارزمية")}},
			expect: 10,
		},
		{
			name:   "emoji",
			pos:    0,
			input:  block{Lines: []gem.String{gem.New("😍😎😑😐😏")}},
			expect: 5,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert := assertion.New(t)

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
		input  block
		expect string
	}{
		{
			name: "join with some empty, trailing separator",
			input: block{
				Lines: []gem.String{
					gem.New("test1"),
					gem.Empty,
					gem.New("test2"),
					gem.Empty,
					gem.Empty,
					gem.New("test3"),
				},
				LineSeparator:     gem.New("\n"),
				TrailingSeparator: true,
			},
			expect: "test1\n\ntest2\n\n\ntest3\n",
		},
		{
			name: "join with some empty, no trailing separator",
			input: block{
				Lines: []gem.String{
					gem.New("test1"),
					gem.Empty,
					gem.New("test2"),
					gem.Empty,
					gem.Empty,
					gem.New("test3"),
				},
				LineSeparator: gem.New("\n"),
			},
			expect: "test1\n\ntest2\n\n\ntest3",
		},
		{
			name: "join 3 lines, trailing separator",
			input: block{
				Lines: []gem.String{
					gem.New("test1"),
					gem.New("test2"),
					gem.New("test3"),
				},
				LineSeparator:     gem.New("\n"),
				TrailingSeparator: true,
			},
			expect: "test1\ntest2\ntest3\n",
		},
		{
			name: "join 3 lines, no trailing separator",
			input: block{
				Lines: []gem.String{
					gem.New("test1"),
					gem.New("test2"),
					gem.New("test3"),
				},
				LineSeparator:     gem.New("\n"),
				TrailingSeparator: false,
			},
			expect: "test1\ntest2\ntest3",
		},
		{
			name: "join 3 lines, alternate separator",
			input: block{
				Lines: []gem.String{
					gem.New("test1"),
					gem.New("test2"),
					gem.New("test3"),
				},
				LineSeparator:     gem.New("_"),
				TrailingSeparator: false,
			},
			expect: "test1_test2_test3",
		},
		{
			name: "join nil lines",
			input: block{
				Lines:         nil,
				LineSeparator: gem.New("\n"),
			},
			expect: "",
		},
		{
			name: "join nil lines, line terminator on",
			input: block{
				Lines:             nil,
				LineSeparator:     gem.New("\n"),
				TrailingSeparator: true,
			},
			expect: "\n",
		},
		{
			name: "join empty lines",
			input: block{
				Lines:         []gem.String{},
				LineSeparator: gem.New("\n"),
			},
			expect: "",
		},
		{
			name: "join empty lines, line terminator on",
			input: block{
				Lines:             []gem.String{},
				LineSeparator:     gem.New("\n"),
				TrailingSeparator: true,
			},
			expect: "\n",
		},
		{
			name: "join 1 empty line",
			input: block{
				Lines:         []gem.String{gem.Empty},
				LineSeparator: gem.New("\n"),
			},
			expect: "",
		},
		{
			name: "join 1 empty line, line terminator on",
			input: block{
				Lines:             []gem.String{gem.Empty},
				LineSeparator:     gem.New("\n"),
				TrailingSeparator: true,
			},
			expect: "\n",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert := assertion.New(t)
			actual := tc.input.Join()
			assert.Equal(tc.expect, actual)
		})
	}
}
