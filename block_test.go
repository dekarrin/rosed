package rosed

import (
	"sort"
	"testing"

	"github.com/dekarrin/assertion"
	"github.com/dekarrin/rosed/internal/gem"

	"github.com/stretchr/testify/assert"
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
			text: gem.Zero,
			sep:  _g("\n"),
			expected: block{
				Lines:             []gem.String{},
				LineSeparator:     _g("\n"),
				TrailingSeparator: false,
			},
		},
		{
			name: "no lines - trailing newline",
			text: _g("\n"),
			sep:  _g("\n"),
			expected: block{
				Lines: []gem.String{
					gem.Zero,
				},
				LineSeparator:     _g("\n"),
				TrailingSeparator: true,
			},
		},
		{
			name: "one line - no trailing newline",
			text: _g("hello"),
			sep:  _g("\n"),
			expected: block{
				Lines: []gem.String{
					_g("hello"),
				},
				LineSeparator:     _g("\n"),
				TrailingSeparator: false,
			},
		},
		{
			name: "one line - trailing newline",
			text: _g("hello\n"),
			sep:  _g("\n"),
			expected: block{
				Lines: []gem.String{
					_g("hello"),
				},
				LineSeparator:     _g("\n"),
				TrailingSeparator: true,
			},
		},
		{
			name: "multi line - no trailing newline",
			text: _g("hello\nthere\ntest"),
			sep:  _g("\n"),
			expected: block{
				Lines: []gem.String{
					_g("hello"),
					_g("there"),
					_g("test"),
				},
				LineSeparator:     _g("\n"),
				TrailingSeparator: false,
			},
		},
		{
			name: "multi line - trailing newline",
			text: _g("hello\nthere\ntest\n"),
			sep:  _g("\n"),
			expected: block{
				Lines: []gem.String{
					_g("hello"),
					_g("there"),
					_g("test"),
				},
				LineSeparator:     _g("\n"),
				TrailingSeparator: true,
			},
		},
		{
			name: "multi line - no trailing separator - alternate separator",
			text: _g("john\negbert\trose\nlalonde\tdave\nstrider\tjade\nharley"),
			sep:  _g("\t"),
			expected: block{
				Lines: []gem.String{
					_g("john\negbert"),
					_g("rose\nlalonde"),
					_g("dave\nstrider"),
					_g("jade\nharley"),
				},
				LineSeparator:     _g("\t"),
				TrailingSeparator: false,
			},
		},
		{
			name: "multi line - trailing separator - alternate separator",
			text: _g("john\negbert\trose\nlalonde\tdave\nstrider\tjade\nharley\t"),
			sep:  _g("\t"),
			expected: block{
				Lines: []gem.String{
					_g("john\negbert"),
					_g("rose\nlalonde"),
					_g("dave\nstrider"),
					_g("jade\nharley"),
				},
				LineSeparator:     _g("\t"),
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
			b2:       block{Lines: []gem.String{gem.Zero}},
			expected: false,
		},
		{
			name:     "nil Lines != 1 empty line",
			b1:       block{Lines: nil},
			b2:       block{Lines: []gem.String{gem.Zero}},
			expected: false,
		},
		{
			name:     "default != 1 empty line",
			b1:       block{},
			b2:       block{Lines: []gem.String{gem.Zero}},
			expected: false,
		},
		{
			name:     "1 empty line == 1 empty line",
			b1:       block{Lines: []gem.String{gem.Zero}},
			b2:       block{Lines: []gem.String{gem.Zero}},
			expected: true,
		},
		{
			name:     "1 filled line == same filled line",
			b1:       block{Lines: []gem.String{_g("test")}},
			b2:       block{Lines: []gem.String{_g("test")}},
			expected: true,
		},
		{
			name:     "1 filled line != different filled line",
			b1:       block{Lines: []gem.String{_g("test1")}},
			b2:       block{Lines: []gem.String{_g("test2")}},
			expected: false,
		},
		{
			name:     "3 empty lines != 4 empty lines",
			b1:       block{Lines: []gem.String{gem.Zero, gem.Zero, gem.Zero}},
			b2:       block{Lines: []gem.String{gem.Zero, gem.Zero, gem.Zero, gem.Zero}},
			expected: false,
		},
		{
			name:     "3 empty lines == 3 empty lines",
			b1:       block{Lines: []gem.String{gem.Zero, gem.Zero, gem.Zero}},
			b2:       block{Lines: []gem.String{gem.Zero, gem.Zero, gem.Zero}},
			expected: true,
		},
		{
			name:     "3 filled lines == same 3 filled lines",
			b1:       block{Lines: []gem.String{_g("a"), _g("ab"), _g("abc")}},
			b2:       block{Lines: []gem.String{_g("a"), _g("ab"), _g("abc")}},
			expected: true,
		},
		{
			name:     "3 filled lines != 4 filled lines",
			b1:       block{Lines: []gem.String{_g("a"), _g("ab"), _g("abc")}},
			b2:       block{Lines: []gem.String{_g("a"), _g("ab"), _g("abc"), _g("abcd")}},
			expected: false,
		},
		{
			name:     "different separators",
			b1:       block{LineSeparator: _g("\n")},
			b2:       block{LineSeparator: _g("\t")},
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
		{"1 empty line", block{Lines: []gem.String{gem.Zero}}, 1},
		{"1 filled line", block{Lines: []gem.String{_g("test")}}, 1},
		{"3 empty lines", block{Lines: []gem.String{gem.Zero, gem.Zero, gem.Zero}}, 3},
		{"3 filled lines", block{Lines: []gem.String{_g("a"), _g("b"), _g("c")}}, 3},
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
			input:    block{Lines: []gem.String{_g("test")}},
			expected: block{Lines: []gem.String{_g("test")}},
		},
		{
			name:     "3 lines already sorted does not change",
			input:    block{Lines: []gem.String{_g("testA"), _g("testB"), _g("testC")}},
			expected: block{Lines: []gem.String{_g("testA"), _g("testB"), _g("testC")}},
		},
		{
			name:     "3 lines",
			input:    block{Lines: []gem.String{_g("testC"), _g("testA"), _g("testB")}},
			expected: block{Lines: []gem.String{_g("testA"), _g("testB"), _g("testC")}},
		},
		{
			name:     "other properties are not touched",
			input:    block{Lines: nil, LineSeparator: _g("\t"), TrailingSeparator: true},
			expected: block{Lines: nil, LineSeparator: _g("\t"), TrailingSeparator: true},
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
			append: gem.Zero,
			input:  block{Lines: nil},
			expect: block{Lines: []gem.String{gem.Zero}},
		},
		{
			name:   "append empty line to empty Lines",
			append: gem.Zero,
			input:  block{Lines: []gem.String{}},
			expect: block{Lines: []gem.String{gem.Zero}},
		},
		{
			name:   "append empty line to default",
			append: gem.Zero,
			input:  block{},
			expect: block{Lines: []gem.String{gem.Zero}},
		},
		{
			name:   "append filled line to default",
			append: _g("vriska"),
			input:  block{},
			expect: block{Lines: []gem.String{_g("vriska")}},
		},
		{
			name:   "append line with separator to default",
			append: _g("vriska\nserket\n"),
			input:  block{LineSeparator: _g("\n")},
			expect: block{LineSeparator: _g("\n"), Lines: []gem.String{_g("vriska\nserket\n")}},
		},
		{
			name:   "append line to multiple Lines",
			append: _g("terezi"),
			input:  block{Lines: []gem.String{_g("vriska"), _g("roxy"), _g("latula")}},
			expect: block{Lines: []gem.String{_g("vriska"), _g("roxy"), _g("latula"), _g("terezi")}},
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
			input:  block{Lines: []gem.String{_g("karkat")}},
			expect: block{Lines: []gem.String{_g("karkat")}},
		},
		{
			name:   "append nil to multi",
			append: block{Lines: nil},
			input:  block{Lines: []gem.String{_g("karkat"), _g("kanaya"), _g("gamzee")}},
			expect: block{Lines: []gem.String{_g("karkat"), _g("kanaya"), _g("gamzee")}},
		},
		{
			name:   "append 1 to nil",
			append: block{Lines: []gem.String{_g("vriska")}},
			input:  block{Lines: nil},
			expect: block{Lines: []gem.String{_g("vriska")}},
		},
		{
			name:   "append 1 to 1",
			append: block{Lines: []gem.String{_g("vriska")}},
			input:  block{Lines: []gem.String{_g("terezi")}},
			expect: block{Lines: []gem.String{_g("terezi"), _g("vriska")}},
		},
		{
			name:   "append 1 to multi",
			append: block{Lines: []gem.String{_g("tavros")}},
			input:  block{Lines: []gem.String{_g("aradia"), _g("sollux")}},
			expect: block{Lines: []gem.String{_g("aradia"), _g("sollux"), _g("tavros")}},
		},
		{
			name:   "append multi to nil",
			append: block{Lines: []gem.String{_g("equius"), _g("nepeta"), _g("eridan")}},
			input:  block{Lines: nil},
			expect: block{Lines: []gem.String{_g("equius"), _g("nepeta"), _g("eridan")}},
		},
		{
			name:   "append multi to 1",
			append: block{Lines: []gem.String{_g("equius"), _g("nepeta"), _g("vriska")}},
			input:  block{Lines: []gem.String{_g("feferi")}},
			expect: block{Lines: []gem.String{_g("feferi"), _g("equius"), _g("nepeta"), _g("vriska")}},
		},
		{
			name:   "append multi to multi",
			append: block{Lines: []gem.String{_g("nepeta"), _g("kanaya"), _g("aradia")}},
			input:  block{Lines: []gem.String{_g("feferi"), _g("eridan"), _g("equius")}},
			expect: block{Lines: []gem.String{_g("feferi"), _g("eridan"), _g("equius"), _g("nepeta"), _g("kanaya"), _g("aradia")}},
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
		{"append 1 to nil", 1, block{Lines: nil}, block{Lines: []gem.String{gem.Zero}}},
		{"append 1 to empty", 1, block{Lines: []gem.String{}}, block{Lines: []gem.String{gem.Zero}}},
		{"append 1 to default", 1, block{}, block{Lines: []gem.String{gem.Zero}}},
		{"append 1 to 1", 1, block{Lines: []gem.String{_g("vriska")}}, block{Lines: []gem.String{_g("vriska"), gem.Zero}}},
		{"append 1 to many", 1, block{Lines: []gem.String{_g("vriska"), _g("terezi")}}, block{Lines: []gem.String{_g("vriska"), _g("terezi"), gem.Zero}}},
		{"append 3 to 1", 3, block{Lines: []gem.String{_g("vriska")}}, block{Lines: []gem.String{_g("vriska"), gem.Zero, gem.Zero, gem.Zero}}},
		{"append 3 to many", 3, block{Lines: []gem.String{_g("vriska"), _g("terezi")}}, block{Lines: []gem.String{_g("vriska"), _g("terezi"), gem.Zero, gem.Zero, gem.Zero}}},
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
			args:   args{-1, gem.Zero},
			input:  block{},
			panics: true,
		},
		{
			name:   "index too high causes panic",
			args:   args{20, gem.Zero},
			input:  block{Lines: []gem.String{gem.Zero, gem.Zero}},
			panics: true,
		},
		{
			name:   "set 0th line",
			args:   args{0, _g("new text")},
			input:  block{Lines: []gem.String{_g("old text"), _g("test")}},
			expect: block{Lines: []gem.String{_g("new text"), _g("test")}},
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
		expect gem.String
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
			input:  block{Lines: []gem.String{gem.Zero, gem.Zero}},
			panics: true,
		},
		{
			name:   "get 0th line",
			pos:    0,
			input:  block{Lines: []gem.String{_g("old text"), _g("test")}},
			expect: _g("old text"),
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

func Test_Block_Remove(t *testing.T) {
	testCases := []struct {
		name string
		pos int
		input block
		expect block
	}{
		{
			name: "remove from empty has no effect",
			pos: 0,
			input: block{},
			expect: block{},
		},
		{
			name: "too-high pos has no effect",
			pos: 7,
			input: block{Lines: []gem.String{
				_g("vriska"),
				_g("terezi"),
				_g("kanaya"),
				_g("wind"),
				_g("mari"),
				_g("deka"),
				_g("equius"),
			}},
			expect: block{Lines: []gem.String{
				_g("vriska"),
				_g("terezi"),
				_g("kanaya"),
				_g("wind"),
				_g("mari"),
				_g("deka"),
				_g("equius"),
			}},
		},
		{
			name: "negative pos has no effect",
			pos: -1,
			input: block{Lines: []gem.String{
				_g("vriska"),
				_g("terezi"),
				_g("kanaya"),
				_g("wind"),
				_g("mari"),
				_g("deka"),
				_g("equius"),
			}},
			expect: block{Lines: []gem.String{
				_g("vriska"),
				_g("terezi"),
				_g("kanaya"),
				_g("wind"),
				_g("mari"),
				_g("deka"),
				_g("equius"),
			}},
		},
		{
			name: "remove the only line that exists",
			pos: 0,
			input: block{Lines: []gem.String{
				_g("line"),
			}},
			expect: block{Lines: []gem.String{}},
		},
		{
			name: "remove from start",
			pos: 0,
			input: block{Lines: []gem.String{
				_g("line1"),
				_g("line2"),
				_g("line3"),
				_g("line4"),
				_g("line5"),
			}},
			expect: block{Lines: []gem.String{
				_g("line2"),
				_g("line3"),
				_g("line4"),
				_g("line5"),
			}},
		},
		{
			name: "remove from end",
			pos: 4,
			input: block{Lines: []gem.String{
				_g("line1"),
				_g("line2"),
				_g("line3"),
				_g("line4"),
				_g("line5"),
			}},
			expect: block{Lines: []gem.String{
				_g("line1"),
				_g("line2"),
				_g("line3"),
				_g("line4"),
			}},
		},
		{
			name: "remove from middle",
			pos: 2,
			input: block{Lines: []gem.String{
				_g("line1"),
				_g("line2"),
				_g("line3"),
				_g("line4"),
				_g("line5"),
			}},
			expect: block{Lines: []gem.String{
				_g("line1"),
				_g("line2"),
				_g("line4"),
				_g("line5"),
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
			input:  block{Lines: []gem.String{gem.Zero, gem.Zero}},
			panics: true,
		},
		{
			name:   "empty string",
			pos:    0,
			input:  block{Lines: []gem.String{gem.Zero}},
			expect: 0,
		},
		{
			name:   "latin-1",
			pos:    0,
			input:  block{Lines: []gem.String{_g("test")}},
			expect: 4,
		},
		{
			name:   "japanese",
			pos:    0,
			input:  block{Lines: []gem.String{_g("こんにちは世界")}},
			expect: 7,
		},
		{
			name:   "russian",
			pos:    0,
			input:  block{Lines: []gem.String{_g("При́пять")}},
			expect: 7,
		},
		{
			name:   "arabic",
			pos:    0,
			input:  block{Lines: []gem.String{_g("الخوارزمية")}},
			expect: 10,
		},
		{
			name:   "emoji",
			pos:    0,
			input:  block{Lines: []gem.String{_g("😍😎😑😐😏")}},
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
		expect gem.String
	}{
		{
			name: "join with some empty, trailing separator",
			input: block{
				Lines: []gem.String{
					_g("test1"),
					gem.Zero,
					_g("test2"),
					gem.Zero,
					gem.Zero,
					_g("test3"),
				},
				LineSeparator:     _g("\n"),
				TrailingSeparator: true,
			},
			expect: _g("test1\n\ntest2\n\n\ntest3\n"),
		},
		{
			name: "join with some empty, no trailing separator",
			input: block{
				Lines: []gem.String{
					_g("test1"),
					gem.Zero,
					_g("test2"),
					gem.Zero,
					gem.Zero,
					_g("test3"),
				},
				LineSeparator: _g("\n"),
			},
			expect: _g("test1\n\ntest2\n\n\ntest3"),
		},
		{
			name: "join 3 lines, trailing separator",
			input: block{
				Lines: []gem.String{
					_g("test1"),
					_g("test2"),
					_g("test3"),
				},
				LineSeparator:     _g("\n"),
				TrailingSeparator: true,
			},
			expect: _g("test1\ntest2\ntest3\n"),
		},
		{
			name: "join 3 lines, no trailing separator",
			input: block{
				Lines: []gem.String{
					_g("test1"),
					_g("test2"),
					_g("test3"),
				},
				LineSeparator:     _g("\n"),
				TrailingSeparator: false,
			},
			expect: _g("test1\ntest2\ntest3"),
		},
		{
			name: "join 3 lines, alternate separator",
			input: block{
				Lines: []gem.String{
					_g("test1"),
					_g("test2"),
					_g("test3"),
				},
				LineSeparator:     _g("_"),
				TrailingSeparator: false,
			},
			expect: _g("test1_test2_test3"),
		},
		{
			name: "join nil lines",
			input: block{
				Lines:         nil,
				LineSeparator: _g("\n"),
			},
			expect: gem.Zero,
		},
		{
			name: "join nil lines, line terminator on",
			input: block{
				Lines:             nil,
				LineSeparator:     _g("\n"),
				TrailingSeparator: true,
			},
			expect: _g("\n"),
		},
		{
			name: "join empty lines",
			input: block{
				Lines:         []gem.String{},
				LineSeparator: _g("\n"),
			},
			expect: gem.Zero,
		},
		{
			name: "join empty lines, line terminator on",
			input: block{
				Lines:             []gem.String{},
				LineSeparator:     _g("\n"),
				TrailingSeparator: true,
			},
			expect: _g("\n"),
		},
		{
			name: "join 1 empty line",
			input: block{
				Lines:         []gem.String{gem.Zero},
				LineSeparator: _g("\n"),
			},
			expect: gem.Zero,
		},
		{
			name: "join 1 empty line, line terminator on",
			input: block{
				Lines:             []gem.String{gem.Zero},
				LineSeparator:     _g("\n"),
				TrailingSeparator: true,
			},
			expect: _g("\n"),
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
