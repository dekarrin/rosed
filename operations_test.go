package rosed

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TODO: do not have test_wrap do any options checking, instead leave both WithOpts().Wrap() and WrapOpts() to
// Test_WrapOpts
func Test_Wrap(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		width    int
		options  Options
		expected []string
	}{
		{
			name:    "2 paragraphs, not preserved",
			input:   "this is a line that is split by paragraph in the input.\n\nThis is the second paragraph",
			width:   20,
			options: Options{PreserveParagraphs: false},
			expected: []string{
				"this is a line that",
				"is split by",
				"paragraph in the",
				"input. This is the",
				"second paragraph",
			},
		},
		{
			name:    "1 paragraph, preserved",
			input:   "this is a line that is split by paragraph in the input.",
			width:   20,
			options: Options{PreserveParagraphs: true},
			expected: []string{
				"this is a line that",
				"is split by",
				"paragraph in the",
				"input.",
			},
		},
		{
			name:    "2 paragraphs, preserved",
			input:   "this is a line that is split by paragraph in the input.\n\nThis is the second paragraph",
			width:   20,
			options: Options{PreserveParagraphs: true},
			expected: []string{
				"this is a line that",
				"is split by",
				"paragraph in the",
				"input.",
				"",
				"This is the second",
				"paragraph",
			},
		},
		{
			name:    "3 paragraphs, preserved",
			input:   "this is a line that is split by paragraph in the input.\n\nThis is the second paragraph.\n\nAnd this is the third",
			width:   20,
			options: Options{PreserveParagraphs: true},
			expected: []string{
				"this is a line that",
				"is split by",
				"paragraph in the",
				"input.",
				"",
				"This is the second",
				"paragraph.",
				"",
				"And this is the",
				"third",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert := assert.New(t)
			
			// try via Edit().WithOptions(), no defaults set
			actualText := Edit(tc.input).WithOptions(tc.options).Wrap(tc.width).String()

			// Edit().String returns one string; turn it into actual lines
			actual := strings.Split(actualText, tc.options.WithDefaults().LineSeparator)
			assert.Equal(tc.expected, actual)
		})
	}
}

func Test_WrapOpts(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		width    int
		options  Options
		expected []string
	}{
		{
			name:    "2 paragraphs, not preserved",
			input:   "this is a line that is split by paragraph in the input.\n\nThis is the second paragraph",
			width:   20,
			options: Options{PreserveParagraphs: false},
			expected: []string{
				"this is a line that",
				"is split by",
				"paragraph in the",
				"input. This is the",
				"second paragraph",
			},
		},
		{
			name:    "1 paragraph, preserved",
			input:   "this is a line that is split by paragraph in the input.",
			width:   20,
			options: Options{PreserveParagraphs: true},
			expected: []string{
				"this is a line that",
				"is split by",
				"paragraph in the",
				"input.",
			},
		},
		{
			name:    "2 paragraphs, preserved",
			input:   "this is a line that is split by paragraph in the input.\n\nThis is the second paragraph",
			width:   20,
			options: Options{PreserveParagraphs: true},
			expected: []string{
				"this is a line that",
				"is split by",
				"paragraph in the",
				"input.",
				"",
				"This is the second",
				"paragraph",
			},
		},
		{
			name:    "3 paragraphs, preserved",
			input:   "this is a line that is split by paragraph in the input.\n\nThis is the second paragraph.\n\nAnd this is the third",
			width:   20,
			options: Options{PreserveParagraphs: true},
			expected: []string{
				"this is a line that",
				"is split by",
				"paragraph in the",
				"input.",
				"",
				"This is the second",
				"paragraph.",
				"",
				"And this is the",
				"third",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert := assert.New(t)
			actualText := Edit(tc.input).WrapOpts(tc.width, tc.options).String()
			actual := strings.Split(actualText, tc.options.WithDefaults().LineSeparator)
			assert.Equal(tc.expected, actual)
		})
	}
}

func Test_CollapseSpaces(t *testing.T) {
	testCases := []struct{
		name string
		input string
		expect string
	}{
		{"empty line", "", ""},
		{"no spaces", "bluh", "bluh"},
		{"2 words", "word1 word2", "word1 word2"},
		{"3 words", "word1 word2 word3", "word1 word2 word3"},
		{"3 words with runs of spaces", "word1        word2  word3", "word1 word2 word3"},
		{"run of non-uniform spaces", "word1 \t\t word2", "word1 word2"},
		{"include lineSep", "   " + DefaultLineSeparator + " word1", " word1"},
	}
	
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert := assert.New(t)
			actual := Edit(tc.input).CollapseSpace().String()
			assert.Equal(tc.expect, actual)
		})
	}
}

func Test_CollapseSpacesOpts(t *testing.T) {
	testCases := []struct{
		name string
		input string
		options Options
		expect string
	}{
		{
			name: "empty line, custom lineSep",
			input: "",
			options: Options{LineSeparator: "<P>"},
			expect: "",
		},
		{
			name: "no spaces, custom lineSep",
			input: "bluh",
			options: Options{LineSeparator: "<P>"},
			expect: "bluh",
		},
		{
			name: "2 words, custom lineSep",
			input: "word1 word2",
			options: Options{LineSeparator: "<P>"},
			expect: "word1 word2",
		},
		{
			name: "3 words, custom lineSep",
			input: "word1 word2 word3",
			options: Options{LineSeparator: "<P>"},
			expect: "word1 word2 word3",
		},
		{
			name: "3 words with runs of spaces, custom lineSep",
			input: "word1        word2  word3",
			options: Options{LineSeparator: "<P>"},
			expect: "word1 word2 word3",
		},
		{
			name: "run of non-uniform spaces",
			input: "word1 \t\t word2",
			options: Options{LineSeparator: "<P>"},
			expect: "word1 word2",
		},
		{
			name: "include lineSep",
			input: "   <P> word1",
			options: Options{LineSeparator: "<P>"},
			expect: " word1",
		},
	}
	
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert := assert.New(t)
			
			actualDirect := Edit(tc.input).CollapseSpaceOpts(tc.options).String()
			actualPreOpts := Edit(tc.input).WithOptions(tc.options).CollapseSpace().String()
			
			assert.Equal(tc.expect, actualDirect, "CollapseSpace(opts) check failed")
			assert.Equal(tc.expect, actualPreOpts, "WithOptions(opts).CollapseSpace() check failed")
		})
	}
}

func Test_Insert(t *testing.T) {
	testCases := []struct{
		name string
		input string
		pos int
		insert string
		expect string
	}{
		{"empty into empty", "", 0, "", ""},
		{"empty into non-empty", "start", 0, "", "start"},
		{"non-empty into empty", "", 0, "TEST", "TEST"},
		{"at start of non-empty", "start", 0, "TEST", "TESTstart"},
		{"in middle of text", "abcdef", 2, "TEST", "abTESTcdef"},
		{"at end of text", "end", 3, "TEST", "endTEST"},
		{"far past end interpreted as end", "end", 300, "TEST", "endTEST"},
		{"negative index", "start", -3, "TEST", "stTESTart"},
		{"before decomposed grapheme", "franc\u0327ais", 1, "TEST", "fTESTranc\u0327ais"},
		{"after decomposed grapheme", "franc\u0327ais", 6, "TEST", "franc\u0327aTESTis"},
	}
	
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert := assert.New(t)
			
			actual := Edit(tc.input).Insert(tc.pos, tc.insert).String()
			
			assert.Equal(tc.expect, actual)
		})
	}
}

func Test_Apply(t *testing.T) {
	testCases := []struct{
		name string
		input string
		op LineOperation
		expect string
	}{
		{
			name: "apply nothing to empty editor",
			input: "",
			op: func(lineNo int, line string) []string {
				return []string{"test"}
			},
			expect: "",
		},
		{
			name: "replace one terminated empty line",
			input: DefaultLineSeparator,
			op: func(lineNo int, line string) []string {
				return []string{"test"}
			},
			expect: "test" + DefaultLineSeparator,
		},
		{
			name: "replace two terminated lines",
			input: DefaultLineSeparator + DefaultLineSeparator,
			op: func(lineNo int, line string) []string {
				return []string{"test"}
			},
			expect: "test" + DefaultLineSeparator + "test" + DefaultLineSeparator,
		},
		{
			name: "replace multiple lines using line number",
			input: DefaultLineSeparator + DefaultLineSeparator + DefaultLineSeparator,
			op: func(lineNo int, line string) []string {
				newLine := fmt.Sprintf("test%d", lineNo)
				return []string{newLine}
			},
			expect: "test0" + DefaultLineSeparator + "test1" + DefaultLineSeparator + "test2" + DefaultLineSeparator,
		},
		{
			name: "insert extra line at target position",
			input: "line0" + DefaultLineSeparator + "line2" + DefaultLineSeparator,
			op: func(lineNo int, line string) []string {
				if lineNo == 0 {
					return []string{line, "line1"}
				}
				return []string{line}
			},
			expect: "line0" + DefaultLineSeparator + "line1" + DefaultLineSeparator + "line2" + DefaultLineSeparator,
		},
		{
			name: "delete line at target position",
			input: "line0" + DefaultLineSeparator + "extra" + DefaultLineSeparator + "line1" + DefaultLineSeparator,
			op: func(lineNo int, line string) []string {
				if lineNo == 1 {
					return []string{}
				}
				return []string{line}
			},
			expect: "line0" + DefaultLineSeparator + "line1" + DefaultLineSeparator,
		},
	}
	
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert := assert.New(t)
			actual := Edit(tc.input).Apply(tc.op).String()
			assert.Equal(tc.expect, actual)
		})
	}
}

func Test_ApplyOpts(t *testing.T) {
	testCases := []struct{
		name string
		input string
		op LineOperation
		options Options
		expect string
	}{
		{
			name: "apply nothing to empty editor, custom lineSep",
			input: "",
			op: func(lineNo int, line string) []string {
				return []string{"test"}
			},
			options: Options{LineSeparator: "<P>"},
			expect: "",
		},
		{
			name: "replace one terminated empty line, custom lineSep",
			input: "<P>",
			op: func(lineNo int, line string) []string {
				return []string{"test"}
			},
			options: Options{LineSeparator: "<P>"},
			expect: "test<P>",
		},
		{
			name: "replace two terminated lines, custom lineSep",
			input: "<P><P>",
			op: func(lineNo int, line string) []string {
				return []string{"test"}
			},
			options: Options{LineSeparator: "<P>"},
			expect: "test<P>test<P>",
		},
		{
			name: "replace multiple lines using line number, custom lineSep",
			input: "<P><P><P>",
			op: func(lineNo int, line string) []string {
				newLine := fmt.Sprintf("test%d", lineNo)
				return []string{newLine}
			},
			options: Options{LineSeparator: "<P>"},
			expect: "test0<P>test1<P>test2<P>",
		},
		{
			name: "insert extra line at target position, custom lineSep",
			input: "line0<P>line2<P>",
			op: func(lineNo int, line string) []string {
				if lineNo == 0 {
					return []string{line, "line1"}
				}
				return []string{line}
			},
			options: Options{LineSeparator: "<P>"},
			expect: "line0<P>line1<P>line2<P>",
		},
		{
			name: "delete line at target position, custom lineSep",
			input: "line0<P>extra<P>line1<P>",
			op: func(lineNo int, line string) []string {
				if lineNo == 1 {
					return []string{}
				}
				return []string{line}
			},
			options: Options{LineSeparator: "<P>"},
			expect: "line0<P>line1<P>",
		},
		{
			name: "apply once for empty editor, noTrailing=true",
			input: "",
			op: func(lineNo int, line string) []string {
				return []string{"test"}
			},
			options: Options{NoTrailingLineSeparators: true},
			expect: "test",
		},
		{
			name: "apply n+1 for n lines, noTrailing=true",
			input: DefaultLineSeparator + DefaultLineSeparator,
			op: func(lineNo int, line string) []string {
				return []string{"test"}
			},
			options: Options{NoTrailingLineSeparators: true},
			expect: "test" + DefaultLineSeparator + "test" + DefaultLineSeparator + "test",
		},
		{
			name: "apply once for empty editor, noTrailing=true, custom lineSep",
			input: "",
			op: func(lineNo int, line string) []string {
				return []string{"test"}
			},
			options: Options{NoTrailingLineSeparators: true, LineSeparator: "<P>"},
			expect: "test",
		},
		{
			name: "apply n+1 for n lines, noTrailing=true, custom lineSep",
			input: "<P><P>",
			op: func(lineNo int, line string) []string {
				return []string{"test"}
			},
			options: Options{NoTrailingLineSeparators: true, LineSeparator: "<P>"},
			expect: "test<P>test<P>test",
		},
	}
	
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert := assert.New(t)
			
			actualDirect := Edit(tc.input).ApplyOpts(tc.op, tc.options).String()
			actualPreOpts := Edit(tc.input).WithOptions(tc.options).Apply(tc.op).String()
			
			assert.Equal(tc.expect, actualDirect, "ApplyOpts(opts) check failed")
			assert.Equal(tc.expect, actualPreOpts, "WithOptions(opts).Apply() check failed")
		})
	}
}

func Test_ApplyParagraphs(t *testing.T) {
	testCases := []struct{
		name string
		input string
		op ParagraphOperation
		expect string
	}{
		{
			name: "empty string = 1 empty para",
			input: "",
			op: func(idx int, para, sepPrefix, sepSuffix string) []string {
				return []string{"test"}
			},
			expect: "test",
		},
		{
			name: "replace two empty paras",
			input: DefaultParagraphSeparator,
			op: func(idx int, para, sepPrefix, sepSuffix string) []string {
				return []string{"test"}
			},
			expect: "test" + DefaultParagraphSeparator + "test",
		},
		{
			name: "replace multiple paras using para number",
			input: DefaultParagraphSeparator + DefaultParagraphSeparator,
			op: func(idx int, para, sepPrefix, sepSuffix string) []string {
				newPara := fmt.Sprintf("test%d", idx)
				return []string{newPara}
			},
			expect: "test0" + DefaultParagraphSeparator + "test1" + DefaultParagraphSeparator + "test2",
		},
		{
			name: "insert extra para at target position",
			input: "para0" + DefaultParagraphSeparator + "para2",
			op: func(idx int, para, sepPrefix, sepSuffix string) []string {
				if idx == 0 {
					return []string{para, "para1"}
				}
				return []string{para}
			},
			expect: "para0" + DefaultParagraphSeparator + "para1" + DefaultParagraphSeparator + "para2",
		},
		{
			name: "delete para at target position",
			input: "para0" + DefaultParagraphSeparator + "extra" + DefaultParagraphSeparator + "para1",
			op: func(idx int, para, sepPrefix, sepSuffix string) []string {
				if idx == 1 {
					return []string{}
				}
				return []string{para}
			},
			expect: "para0" + DefaultParagraphSeparator + "para1",
		},
	}
	
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert := assert.New(t)
			actual := Edit(tc.input).ApplyParagraphs(tc.op).String()
			assert.Equal(tc.expect, actual)
		})
	}
}

func Test_ApplyParagraphsOpts(t *testing.T) {
	testCases := []struct{
		name string
		input string
		op ParagraphOperation
		options Options
		expect string
	}{
		{
			name: "apply once to empty editor, custom paraSep",
			input: "",
			op: func(idx int, para, sepPrefix, sepSuffix string) []string {
				return []string{"test"}
			},
			options: Options{ParagraphSeparator: "<P>"},
			expect: "test",
		},
		{
			name: "replace two empty paras, custom paraSep",
			input: "<P>",
			op: func(idx int, para, sepPrefix, sepSuffix string) []string {
				return []string{"test"}
			},
			options: Options{ParagraphSeparator: "<P>"},
			expect: "test<P>test",
		},
		{
			name: "replace multiple paras using para number, custom paraSep",
			input: "<P><P>",
			op: func(idx int, para, sepPrefix, sepSuffix string) []string {
				newPara := fmt.Sprintf("test%d", idx)
				return []string{newPara}
			},
			options: Options{ParagraphSeparator: "<P>"},
			expect: "test0<P>test1<P>test2",
		},
		{
			name: "insert extra para at target position, custom paraSep",
			input: "para0<P>para2",
			op: func(idx int, para, sepPrefix, sepSuffix string) []string {
				if idx == 0 {
					return []string{para, "para1"}
				}
				return []string{para}
			},
			options: Options{ParagraphSeparator: "<P>"},
			expect: "para0<P>para1<P>para2",
		},
		{
			name: "delete para at target position, custom paraSep",
			input: "para0<P>extra<P>para1",
			op: func(idx int, para, sepPrefix, sepSuffix string) []string {
				if idx == 1 {
					return []string{}
				}
				return []string{para}
			},
			options: Options{ParagraphSeparator: "<P>"},
			expect: "para0<P>para1",
		},
	}
	
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert := assert.New(t)
			
			actualDirect := Edit(tc.input).ApplyParagraphsOpts(tc.op, tc.options).String()
			actualPreOpts := Edit(tc.input).WithOptions(tc.options).ApplyParagraphs(tc.op).String()
			
			assert.Equal(tc.expect, actualDirect, "ApplyParagraphsOpts(opts) check failed")
			assert.Equal(tc.expect, actualPreOpts, "WithOptions(opts).ApplyParagraphs() check failed")
			
		})
	}
}

func Test_Indent(t *testing.T) {
	testCases := []struct{
		name string
		input string
		level int
		expect string
	}{
		{
			name: "do nothing to empty editor",
			input: "",
			level: 20,
			expect: "",
		},
		{
			name: "empty line, 0 levels",
			input: DefaultLineSeparator,
			level: 0,
			expect: DefaultLineSeparator,
		},
		{
			name: "empty line, -1 levels",
			input: DefaultLineSeparator,
			level: -1,
			expect: DefaultLineSeparator,
		},
		{
			name: "empty line, 1 level",
			input: DefaultLineSeparator,
			level: 1,
			expect: DefaultIndentString + DefaultLineSeparator,
		},
		{
			name: "empty line, multiple levels",
			input: DefaultLineSeparator,
			level: 3,
			expect: DefaultIndentString + DefaultIndentString + DefaultIndentString + DefaultLineSeparator,
		},
		{
			name: "non-empty line, 0 levels",
			input: "test" + DefaultLineSeparator,
			level: 0,
			expect: "test" + DefaultLineSeparator,
		},
		{
			name: "non-empty string, -1 levels",
			input: "test" + DefaultLineSeparator,
			level: -1,
			expect: "test" + DefaultLineSeparator,
		},
		{
			name: "non-empty string, 1 level",
			input: "test" + DefaultLineSeparator,
			level: 1,
			expect: DefaultIndentString + "test" + DefaultLineSeparator,
		},
		{
			name: "non-empty string, multiple levels",
			input: "test" + DefaultLineSeparator,
			level: 3,
			expect: DefaultIndentString + DefaultIndentString + DefaultIndentString + "test" + DefaultLineSeparator,
		},
		{
			name: "multi-line string, 0 levels",
			input: "test0" + DefaultLineSeparator + "test1" + DefaultLineSeparator + "test2" + DefaultLineSeparator,
			level: 0,
			expect: "test0" + DefaultLineSeparator +
			        "test1" + DefaultLineSeparator +
			        "test2" + DefaultLineSeparator,
		},
		{
			name: "multi-line string, -1 levels",
			input: "test0" + DefaultLineSeparator + "test1" + DefaultLineSeparator + "test2" + DefaultLineSeparator,
			level: -1,
			expect: "test0" + DefaultLineSeparator +
			        "test1" + DefaultLineSeparator +
			        "test2" + DefaultLineSeparator,
		},
		{
			name: "multi-line string, 1 levels",
			input: "test0" + DefaultLineSeparator + "test1" + DefaultLineSeparator + "test2" + DefaultLineSeparator,
			level: 1,
			expect: DefaultIndentString + "test0" + DefaultLineSeparator +
			        DefaultIndentString + "test1" + DefaultLineSeparator +
			        DefaultIndentString + "test2" + DefaultLineSeparator,
		},
		{
			name: "multi-line string, multiple levels",
			input: "test0" + DefaultLineSeparator + "test1" + DefaultLineSeparator + "test2" + DefaultLineSeparator,
			level: 3,
			expect: DefaultIndentString + DefaultIndentString + DefaultIndentString + "test0" + DefaultLineSeparator +
			        DefaultIndentString + DefaultIndentString + DefaultIndentString + "test1" + DefaultLineSeparator +
			        DefaultIndentString + DefaultIndentString + DefaultIndentString + "test2" + DefaultLineSeparator,
		},
	}
	
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert := assert.New(t)
			actual := Edit(tc.input).Indent(tc.level).String()
			assert.Equal(tc.expect, actual)
		})
	}
}

	
// NOTE: IndentOpts needs to test ALL permutations of options: preserveparas, custom line sep, custom para sep, AND trailing line behave, as well as
// indentStr ofc.
func Test_IndentOpts(t *testing.T) {
	testCases := []struct{
		name string
		input string
		level int
		options Options
		expect string
	}{
		{
			name: "empty line, noTrailing",
			input: "",
			level: 1,
			options: Options{
				NoTrailingLineSeparators: true,
			},
			expect: DefaultIndentString,
		},
		{
			name: "empty line, custom str",
			input: DefaultLineSeparator,
			level: 1,
			options: Options{
				IndentStr: ">",
			},
			expect: ">" + DefaultLineSeparator,
		},
		{
			name: "empty line, custom str, noTrailing",
			input: "",
			level: 1,
			options: Options{
				IndentStr: ">",
				NoTrailingLineSeparators: true,
			},
			expect: ">",
		},
		{
			name: "multi-line, custom str",
			input: "here is a custom multi-line with unique break,<br/>normal break\n" +
				"folded with custom parasep break<p>\n" +
				"and a normal para end\n\n" +
				"with some more breaks\n" +
				"and another<p>" +
				"custom line end folded with custom parasep<p><p>" +
				"a third possible para\n\n" +
				"finally a completely<p>" +
				"unique para end<para>\n" +
				"and some text<br/>" +
				"to break things up\n",
			level: 1,
			options: Options{
				IndentStr: ">",
			},
			expect: ">here is a custom multi-line with unique break,<br/>normal break\n" +
				">folded with custom parasep break<p>\n" +
				">and a normal para end\n>\n" +
				">with some more breaks\n" +
				">and another<p>" +
				"custom line end folded with custom parasep<p><p>" +
				"a third possible para\n>\n" +
				">finally a completely<p>" +
				"unique para end<para>\n" +
				">and some text<br/>" +
				"to break things up\n",
			
		},
		{
			name: "multi-line, noTrailing",
			input: "line0\nline1\n",
			level: 1,
			options: Options{
				NoTrailingLineSeparators: true,
			},
			expect: "\tline0\n\tline1\n\t",
		},
		{
			name: "multi-line, custom paraSep",
			input: "p0,line0\np0,line1\np0,line2<p><p>p1,line0\np1,line1",
			level: 1,
			options: Options{
				ParagraphSeparator: "<p><p>",
			},
			expect: "\tp0,line0\n\tp0,line1\n\tp0,line2<p><p>p1,line0\n\tp1,line1",
		},
		{
			name: "multi-line, preserveParas",
			input: "p0,line0\np0,line1\n\np1,line0\np1,line1",
			level: 1,
			options: Options{
				PreserveParagraphs: true,
			},
			expect: "\tp0,line0\n\tp0,line1\n\n\tp1,line0\n\tp1,line1",
		},
		{
			name: "multi-line, custom lineSep",
			input: "line0\nmore line0<p>line1<p>line2\nbreak ignored",
			level: 1,
			options: Options{
				LineSeparator: "<p>",
			},
			expect: "\tline0\nmore line0<p>\tline1<p>\tline2\nbreak ignored",
		},
		{
			name: "multi-line, custom str, noTrailing",
			input: "line0\nline1\nline2\n",
			level: 1,
			options: Options{
				IndentStr: ">",
				NoTrailingLineSeparators: true,
			},
			expect: ">line0\n>line1\n>line2\n>",
		},
		{
			name: "multi-line, custom str, custom paraSep",
			input: "p0,line0\np0,line1\np0,line2<p><p>p1,line0\np1,line1",
			level: 1,
			options: Options{
				IndentStr: ">",
				ParagraphSeparator: "<p><p>",
			},
			expect: ">p0,line0\n>p0,line1\n>p0,line2<p><p>p1,line0\n>p1,line1",
		},
		{
			name: "multi-line, custom str, preserveParas",
			input: "p0,line0\np0,line1\n\np1,line0\np1,line1",
			level: 1,
			options: Options{
				IndentStr: ">",
				PreserveParagraphs: true,
			},
			expect: ">p0,line0\n>p0,line1\n\n>p1,line0\n>p1,line1",
		},
		{
			name: "multi-line, custom str, custom lineSep",
			input: "line0\nmore line0<p>line1<p>line2\nbreak ignored",
			level: 1,
			options: Options{
				IndentStr: ">",
				LineSeparator: "<p>",
			},
			expect: ">line0\nmore line0<p>>line1<p>>line2\nbreak ignored",
		},
		{
			name: "multi-line, custom paraSep, preserveParas",
			input: "p0,A\np0,B\n\np0,C<p><p>p1,A\np1,B<p><p>p2,A\np2,B",
			level: 1,
			options: Options{
				ParagraphSeparator: "<p><p>",
				PreserveParagraphs: true,
			},
			expect: "\tp0,A\n\tp0,B\n\t\n\tp0,C<p><p>\tp1,A\n\tp1,B<p><p>\tp2,A\n\tp2,B",
		},
		{
			name: "multi-line, preserveParas, custom lineSep",
			input: "p0,l0\nmore l0<p>p0,l1\n\np1,l0<p>p1,l1<p>p1,l2\nmore l2",
			level: 1,
			options: Options{
				PreserveParagraphs: true,
				LineSeparator: "<p>",
			},
			expect: "\tp0,l0\nmore l0<p>\tp0,l1\n\n\tp1,l0<p>\tp1,l1<p>\tp1,l2\nmore l2",
		},
		{
			name: "multi-line, custom paraSep (folded), preserveParas, custom lineSep (folded)",
			input: "p0,l0<p>p0,l1\nmore l1<p><p>p1,l0\n\nmore l0<p>line1<p><p><p><p>next para<p>pn,ln",
			level: 1,
			options: Options{
				ParagraphSeparator: "<p><p>",
				PreserveParagraphs: true,
				LineSeparator: "<p>",
			},
			expect: "\tp0,l0<p>\tp0,l1\nmore l1<p><p>\tp1,l0\n\nmore l0<p>\tline1<p><p><p><p>\tnext para<p>\tpn,ln",
		},
		{
			name: "multi-line, custom paraSep (non-folded), preserveParas, custom lineSep",
			input: "p0,l0<p>p0,l1\nmore l1<para>p1,l0\n\nmore l0<p>line1<para><para>next para<p>pn,ln",
			level: 1,
			options: Options{
				ParagraphSeparator: "<para>",
				PreserveParagraphs: true,
				LineSeparator: "<p>",
			},
			expect: "\tp0,l0<p>\tp0,l1\nmore l1<para>\tp1,l0\n\nmore l0<p>\tline1<para><para>\tnext para<p>\tpn,ln",
		},
		{
			name: "multi-line, custom str, noTrailing, custom paraSep (folded), preserveParas, custom lineSep",
			input: "p0,line0<p>p0,line1<p><p><p>p1,line0<p>p1,line1<p><p>p2,line1<p>p2,line2<p>",
			level: 1,
			options: Options{
				IndentStr: ">",
				NoTrailingLineSeparators: true,
				ParagraphSeparator: "<p><p>",
				PreserveParagraphs: true,
				LineSeparator: "<p>",
			},
			expect: ">p0,line0<p>>p0,line1<p><p>><p>>p1,line0<p>>p1,line1<p><p>>p2,line1<p>>p2,line2<p>>",
		},
	}
	
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert := assert.New(t)
			
			actualDirect := Edit(tc.input).IndentOpts(tc.level, tc.options).String()
			actualPreOpts := Edit(tc.input).WithOptions(tc.options).Indent(tc.level).String()
			
			assert.Equal(tc.expect, actualDirect, "IndentOpts(opts) check failed")
			assert.Equal(tc.expect, actualPreOpts, "WithOptions(opts).Indent() check failed")
			
		})
	}
}

func Test_Justify(t *testing.T) {
	testCases := []struct{
		name string
		input string
		width int
		expect string
	}{
		{
			name: "empty string",
			input: "",
			width: 10,
			expect: "",
		},
		{
			name: "no spaces",
			input: "bluh",
			width: 10,
			expect: "bluh",
		},
		{
			name: "2 words",
			input: "word1 word2",
			width: 20,
			expect: "word1          word2",
		},
		{
			name: "3 words",
			input: "word1 word2 word3",
			width: 20,
			expect: "word1   word2  word3",
		},
		{
			name: "3 words with runs of spaces",
			input: "word1        word2  word3",
			width: 20,
			expect: "word1   word2  word3",
		},
		{
			name: "line longer than width",
			input: "hello",
			width: 3,
			expect: "hello",
		},
		{
			name: "bad width",
			input: "bluh",
			width: -1,
			expect: "bluh",
		},
		{
			name: "multi-line",
			input:  "a set of three lines" + DefaultLineSeparator +
				"to justify in a" + DefaultLineSeparator +
				"pleasing manner",
			width: 22,
			expect: "a  set of three  lines" + DefaultLineSeparator +
				"to    justify   in   a" + DefaultLineSeparator +
				"pleasing        manner",
		},
	}
	
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert := assert.New(t)
			actual := Edit(tc.input).Justify(tc.width).String()
			assert.Equal(tc.expect, actual)
		})
	}
}

func Test_JustifyOpts(t *testing.T) {
	testCases := []struct{
		name string
		input string
		width int
		options Options
		expect string
	}{
		{
			name: "multi-line, noTrailing",
			input:  "a set of three lines" + DefaultLineSeparator +
				"to justify in a" + DefaultLineSeparator +
				"pleasing manner" + DefaultLineSeparator,
			width: 22,
			options: Options{
				NoTrailingLineSeparators: true,
			},
			expect: "a  set of three  lines" + DefaultLineSeparator +
				"to    justify   in   a" + DefaultLineSeparator +
				"pleasing        manner" + DefaultLineSeparator,
		},
		{
			name: "multi-paragraph, preserved, default parasep",
			input:  "a set of three lines" + DefaultLineSeparator +
				"to justify in a" + DefaultLineSeparator +
				"pleasing manner" +
				DefaultParagraphSeparator +
				"a second paragraph" + DefaultLineSeparator +
				"which should also be" + DefaultLineSeparator +
				"respected",
			width: 22,
			options: Options{
				PreserveParagraphs: true,
			},
			expect: "a  set of three  lines" + DefaultLineSeparator +
				"to    justify   in   a" + DefaultLineSeparator +
				"pleasing        manner" +
				DefaultParagraphSeparator +
				"a   second   paragraph" + DefaultLineSeparator +
				"which  should  also be" + DefaultLineSeparator +
				"respected",
		},
		{
			name: "multi-paragraph, preserved, custom parasep",
			input:  "a set of three lines" + DefaultLineSeparator +
				"to justify in a" + DefaultLineSeparator +
				"pleasing manner" +
				"<P> <P>" +
				"a second paragraph" + DefaultLineSeparator +
				"which should also be" + DefaultLineSeparator +
				"respected",
			width: 22,
			options: Options{
				ParagraphSeparator: "<P> <P>",
				PreserveParagraphs: true,
			},
			expect: "a  set of three  lines" + DefaultLineSeparator +
				"to    justify   in   a" + DefaultLineSeparator +
				"pleasing        manner" +
				"<P> <P>" +
				"a   second   paragraph" + DefaultLineSeparator +
				"which  should  also be" + DefaultLineSeparator +
				"respected",
		},
	}
	
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert := assert.New(t)
			
			actualDirect := Edit(tc.input).JustifyOpts(tc.width, tc.options).String()
			actualPreOpts := Edit(tc.input).WithOptions(tc.options).Justify(tc.width).String()
			
			assert.Equal(tc.expect, actualDirect, "JustifyOpts(opts) check failed")
			assert.Equal(tc.expect, actualPreOpts, "WithOptions(opts).Justify() check failed")
			
		})
	}
}

func Test_InsertTwoColumns(t *testing.T) {
	type args struct {
		pos int
		left string
		right string
		minBetween int
		width int
		leftPercent float64
	}

	testCases := []struct {
		name string
		args args
		input string
		expect string
	}{
		{
			name: "empty lines",
			args: args{
				pos: 0,
				left: "",
				right: "",
				minBetween: 0,
				width: 0,
				leftPercent: 0.0,
			},
			input: "",
			expect: "",
		},
		{
			name: "right col bigger",
			args: args{
				pos: 0,
				left: "This is a test string for the left side",
				right: "Column number two is right here! And it has a lot of content that will be wrapped",
				minBetween: 2,
				width: 30,
				leftPercent: 0.5,
			},
			input: "",
			expect: "This is a test  Column number" + DefaultLineSeparator +
				"string for the  two is right"  + DefaultLineSeparator +
				"left side       here! And it"  + DefaultLineSeparator +
				"                has a lot of"  + DefaultLineSeparator +
				"                content that"  + DefaultLineSeparator +
				"                will be"       + DefaultLineSeparator +
				"                wrapped"       + DefaultLineSeparator,
		},
		{
			name: "left col bigger",
			args: args{
				pos: 0,
				left: "Column number one is right here! And it has a lot of content that will be included",
				right: "This is a test string for the right side",
				minBetween: 2,
				width: 30,
				leftPercent: 0.5,
			},
			input: "",
			expect: "Column number   This is a test" + DefaultLineSeparator +
				"one is right    string for the" + DefaultLineSeparator +
				"here! And it    right side"     + DefaultLineSeparator +
				"has a lot of    "               + DefaultLineSeparator +
				"content that    "               + DefaultLineSeparator +
				"will be         "               + DefaultLineSeparator +
				"included        "               + DefaultLineSeparator,
		},
		{
			name: "equal size columns",
			args: args{
				pos: 0,
				left: "Column number one is right here!",
				right: "This is a test string for the right side",
				minBetween: 2,
				width: 30,
				leftPercent: 0.5,
			},
			input: "",
			expect: "Column number   This is a test" + DefaultLineSeparator +
				"one is right    string for the" + DefaultLineSeparator +
				"here!           right side"     + DefaultLineSeparator,
		},
		{
			name: "insert in middle",
			args: args{
				pos: 2,
				left: "Column number one is right here!",
				right: "This is a test string for the right side",
				minBetween: 2,
				width: 30,
				leftPercent: 0.5,
			},
			input: "hello",
			expect: "heColumn number   This is a test" + DefaultLineSeparator +
				"one is right    string for the" + DefaultLineSeparator +
				"here!           right side"     + DefaultLineSeparator +
				"llo",
		},
		{
			name: "quarter left col",
			args: args{
				pos: 0,
				left: "Column number one is right here!",
				right: "This is a test string for the right side",
				minBetween: 2,
				width: 32,
				leftPercent: 0.25,
			},
			input: "",
			expect: "Column   This is a test string" + DefaultLineSeparator +
				"number   for the right side"    + DefaultLineSeparator +
				"one is   "                      + DefaultLineSeparator +
				"right    "                      + DefaultLineSeparator +
				"here!    "                      + DefaultLineSeparator,
		},
		{
			name: "quarter right col",
			args: args{
				pos: 0,
				left: "Column number one is right here!",
				right: "This is a test string for the right side",
				minBetween: 2,
				width: 32,
				leftPercent: 0.75,
			},
			input: "",
			expect: "Column number one is    This is" + DefaultLineSeparator +
				"right here!             a test"  + DefaultLineSeparator +
				"                        string"  + DefaultLineSeparator +
				"                        for the" + DefaultLineSeparator +
				"                        right"   + DefaultLineSeparator +
				"                        side"    + DefaultLineSeparator,
		},
	}
	
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert := assert.New(t)
			
			ed := Edit(tc.input)
			ed = ed.InsertTwoColumns(
				tc.args.pos,
				tc.args.left,
				tc.args.right,
				tc.args.minBetween,
				tc.args.width,
				tc.args.leftPercent,
			)
			actual := ed.String()
			
			assert.Equal(tc.expect, actual)
		})
	}
}

