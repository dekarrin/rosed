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
				"custom para end folded with custom parasep<p><p>" +
				"a third possible para\n\n" +
				"finally a completely<p>" +
				"unique para end<para>\n" +
				"and some text<br/>" +
				"to break things up",
			level: 1,
			options: Options{
				IndentStr: ">",
			},
			
		},
		{
			name: "multi-line, noTrailing",
			input: "here is a custom multi-line with unique break,<br/>normal break\n" +
				"folded with custom parasep break<p>\n" +
				"and a normal para end\n\n" +
				"with some more breaks\n" +
				"and another<p>" +
				"custom para end folded with custom parasep<p><p>" +
				"a third possible para\n\n" +
				"finally a completely<p>" +
				"unique para end<para>\n" +
				"and some text<br/>" +
				"to break things up",
			level: 1,
			options: Options{
				NoTrailingLineSeparators: true,
			},
			expect: ">",
		},
		{
			name: "multi-line, custom paraSep",
			input: "here is a custom multi-line with unique break,<br/>normal break\n" +
				"folded with custom parasep break<p>\n" +
				"and a normal para end\n\n" +
				"with some more breaks\n" +
				"and another<p>" +
				"custom para end folded with custom parasep<p><p>" +
				"a third possible para\n\n" +
				"finally a completely<p>" +
				"unique para end<para>\n" +
				"and some text<br/>" +
				"to break things up",
			level: 1,
			options: Options{
				ParagraphSeparator: "<p><p>",
			},
		},
		{
			name: "multi-line, preserveParas",
			input: "here is a custom multi-line with unique break,<br/>normal break\n" +
				"folded with custom parasep break<p>\n" +
				"and a normal para end\n\n" +
				"with some more breaks\n" +
				"and another<p>" +
				"custom para end folded with custom parasep<p><p>" +
				"a third possible para\n\n" +
				"finally a completely<p>" +
				"unique para end<para>\n" +
				"and some text<br/>" +
				"to break things up",
			level: 1,
			options: Options{
				PreserveParagraphs: true,
			},
		},
		{
			name: "multi-line, custom lineSep",
			input: "here is a custom multi-line with unique break,<br/>normal break\n" +
				"folded with custom parasep break<p>\n" +
				"and a normal para end\n\n" +
				"with some more breaks\n" +
				"and another<p>" +
				"custom para end folded with custom parasep<p><p>" +
				"a third possible para\n\n" +
				"finally a completely<p>" +
				"unique para end<para>\n" +
				"and some text<br/>" +
				"to break things up",
			level: 1,
			options: Options{
				LineSeparator: "<p>",
			},
		},
		{
			name: "multi-line, custom str, noTrailing",
			input: "here is a custom multi-line with unique break,<br/>normal break\n" +
				"folded with custom parasep break<p>\n" +
				"and a normal para end\n\n" +
				"with some more breaks\n" +
				"and another<p>" +
				"custom para end folded with custom parasep<p><p>" +
				"a third possible para\n\n" +
				"finally a completely<p>" +
				"unique para end<para>\n" +
				"and some text<br/>" +
				"to break things up",
			level: 1,
			options: Options{
				IndentStr: ">",
				NoTrailingLineSeparators: true,
			},
		},
		{
			name: "multi-line, custom str, custom paraSep",
			input: "here is a custom multi-line with unique break,<br/>normal break\n" +
				"folded with custom parasep break<p>\n" +
				"and a normal para end\n\n" +
				"with some more breaks\n" +
				"and another<p>" +
				"custom para end folded with custom parasep<p><p>" +
				"a third possible para\n\n" +
				"finally a completely<p>" +
				"unique para end<para>\n" +
				"and some text<br/>" +
				"to break things up",
			level: 1,
			options: Options{
				IndentStr: ">",
				ParagraphSeparator: "<p><p>",
			},
		},
		{
			name: "multi-line, custom str, preserveParas",
			input: "here is a custom multi-line with unique break,<br/>normal break\n" +
				"folded with custom parasep break<p>\n" +
				"and a normal para end\n\n" +
				"with some more breaks\n" +
				"and another<p>" +
				"custom para end folded with custom parasep<p><p>" +
				"a third possible para\n\n" +
				"finally a completely<p>" +
				"unique para end<para>\n" +
				"and some text<br/>" +
				"to break things up",
			level: 1,
			options: Options{
				IndentStr: ">",
				PreserveParagraphs: true,
			},
		},
		{
			name: "multi-line, custom str, custom lineSep",
			input: "here is a custom multi-line with unique break,<br/>normal break\n" +
				"folded with custom parasep break<p>\n" +
				"and a normal para end\n\n" +
				"with some more breaks\n" +
				"and another<p>" +
				"custom para end folded with custom parasep<p><p>" +
				"a third possible para\n\n" +
				"finally a completely<p>" +
				"unique para end<para>\n" +
				"and some text<br/>" +
				"to break things up",
			level: 1,
			options: Options{
				IndentStr: ">",
				LineSeparator: "<p>",
			},
		},
		{
			name: "multi-line, noTrailing, custom paraSep",
			input: "here is a custom multi-line with unique break,<br/>normal break\n" +
				"folded with custom parasep break<p>\n" +
				"and a normal para end\n\n" +
				"with some more breaks\n" +
				"and another<p>" +
				"custom para end folded with custom parasep<p><p>" +
				"a third possible para\n\n" +
				"finally a completely<p>" +
				"unique para end<para>\n" +
				"and some text<br/>" +
				"to break things up",
			level: 1,
			options: Options{
				NoTrailingLineSeparators: true,
				ParagraphSeparator: "<p><p>",
			}, 
		},
		{
			name: "multi-line, noTrailing, preserveParas",
			input: "here is a custom multi-line with unique break,<br/>normal break\n" +
				"folded with custom parasep break<p>\n" +
				"and a normal para end\n\n" +
				"with some more breaks\n" +
				"and another<p>" +
				"custom para end folded with custom parasep<p><p>" +
				"a third possible para\n\n" +
				"finally a completely<p>" +
				"unique para end<para>\n" +
				"and some text<br/>" +
				"to break things up",
			level: 1,
			options: Options{
				NoTrailingLineSeparators: true,
				PreserveParagraphs: true,
			},
		},
		{
			name: "multi-line, noTrailing, custom lineSep",
			input: "here is a custom multi-line with unique break,<br/>normal break\n" +
				"folded with custom parasep break<p>\n" +
				"and a normal para end\n\n" +
				"with some more breaks\n" +
				"and another<p>" +
				"custom para end folded with custom parasep<p><p>" +
				"a third possible para\n\n" +
				"finally a completely<p>" +
				"unique para end<para>\n" +
				"and some text<br/>" +
				"to break things up",
			level: 1,
			options: Options{
				NoTrailingLineSeparators: true,
				LineSeparator: "<p>",
			},
		},
		{
			name: "multi-line, custom paraSep, preserveParas",
			input: "here is a custom multi-line with unique break,<br/>normal break\n" +
				"folded with custom parasep break<p>\n" +
				"and a normal para end\n\n" +
				"with some more breaks\n" +
				"and another<p>" +
				"custom para end folded with custom parasep<p><p>" +
				"a third possible para\n\n" +
				"finally a completely<p>" +
				"unique para end<para>\n" +
				"and some text<br/>" +
				"to break things up",
			level: 1,
			options: Options{
				ParagraphSeparator: "<p><p>",
				PreserveParagraphs: true,
			},
		},
		{
			name: "multi-line, custom paraSep (folded), custom lineSep (folded)",
			input: "here is a custom multi-line with unique break,<br/>normal break\n" +
				"folded with custom parasep break<p>\n" +
				"and a normal para end\n\n" +
				"with some more breaks\n" +
				"and another<p>" +
				"custom para end folded with custom parasep<p><p>" +
				"a third possible para\n\n" +
				"finally a completely<p>" +
				"unique para end<para>\n" +
				"and some text<br/>" +
				"to break things up",
			level: 1,
			options: Options{
				ParagraphSeparator: "<p><p>",
				LineSeparator: "<p>",
			},
		},
		{
			name: "multi-line, custom paraSep (non-folded), custom lineSep",
			input: "here is a custom multi-line with unique break,<br/>normal break\n" +
				"folded with custom parasep break<p>\n" +
				"and a normal para end\n\n" +
				"with some more breaks\n" +
				"and another<p>" +
				"custom para end folded with custom parasep<p><p>" +
				"a third possible para\n\n" +
				"finally a completely<p>" +
				"unique para end<para>\n" +
				"and some text<br/>" +
				"to break things up",
			level: 1,
			options: Options{
				ParagraphSeparator: "<para>",
				LineSeparator: "<p>",
			},
		},
		{
			name: "multi-line, custom paraSep, custom lineSep (non-folded)",
			input: "here is a custom multi-line with unique break,<br/>normal break\n" +
				"folded with custom parasep break<p>\n" +
				"and a normal para end\n\n" +
				"with some more breaks\n" +
				"and another<p>" +
				"custom para end folded with custom parasep<p><p>" +
				"a third possible para\n\n" +
				"finally a completely<p>" +
				"unique para end<para>\n" +
				"and some text<br/>" +
				"to break things up",
			level: 1,
			options: Options{
				ParagraphSeparator: "<p><p>",
				LineSeparator: "<br/>",
			},
		},
		{
			name: "multi-line, preserveParas, custom lineSep",
			input: "here is a custom multi-line with unique break,<br/>normal break\n" +
				"folded with custom parasep break<p>\n" +
				"and a normal para end\n\n" +
				"with some more breaks\n" +
				"and another<p>" +
				"custom para end folded with custom parasep<p><p>" +
				"a third possible para\n\n" +
				"finally a completely<p>" +
				"unique para end<para>\n" +
				"and some text<br/>" +
				"to break things up",
			level: 1,
			options: Options{
				PreserveParagraphs: true,
				LineSeparator: "<p>",
			},
		},
		{
			name: "multi-line, custom str, noTrailing, custom paraSep",
			input: "here is a custom multi-line with unique break,<br/>normal break\n" +
				"folded with custom parasep break<p>\n" +
				"and a normal para end\n\n" +
				"with some more breaks\n" +
				"and another<p>" +
				"custom para end folded with custom parasep<p><p>" +
				"a third possible para\n\n" +
				"finally a completely<p>" +
				"unique para end<para>\n" +
				"and some text<br/>" +
				"to break things up",
			level: 1,
			options: Options{
				IndentStr: ">",
				NoTrailingLineSeparators: true,
				ParagraphSeparator: "<p><p>",
			},
		},
		{
			name: "multi-line, custom str, noTrailing, preserveParas",
			input: "here is a custom multi-line with unique break,<br/>normal break\n" +
				"folded with custom parasep break<p>\n" +
				"and a normal para end\n\n" +
				"with some more breaks\n" +
				"and another<p>" +
				"custom para end folded with custom parasep<p><p>" +
				"a third possible para\n\n" +
				"finally a completely<p>" +
				"unique para end<para>\n" +
				"and some text<br/>" +
				"to break things up",
			level: 1,
			options: Options{
				IndentStr: ">",
				NoTrailingLineSeparators: true,
				PreserveParagraphs: true,
			},
		},
		{
			name: "multi-line, custom str, noTrailing, custom lineSep",
			input: "here is a custom multi-line with unique break,<br/>normal break\n" +
				"folded with custom parasep break<p>\n" +
				"and a normal para end\n\n" +
				"with some more breaks\n" +
				"and another<p>" +
				"custom para end folded with custom parasep<p><p>" +
				"a third possible para\n\n" +
				"finally a completely<p>" +
				"unique para end<para>\n" +
				"and some text<br/>" +
				"to break things up",
			level: 1,
			options: Options{
				IndentStr: ">",
				NoTrailingLineSeparators: true,
				LineSeparator: "<p>",
			},
		},
		{
			name: "multi-line, custom str, custom paraSep, preserveParas",
			input: "here is a custom multi-line with unique break,<br/>normal break\n" +
				"folded with custom parasep break<p>\n" +
				"and a normal para end\n\n" +
				"with some more breaks\n" +
				"and another<p>" +
				"custom para end folded with custom parasep<p><p>" +
				"a third possible para\n\n" +
				"finally a completely<p>" +
				"unique para end<para>\n" +
				"and some text<br/>" +
				"to break things up",
			level: 1,
			options: Options{IndentStr: ">", ParagraphSeparator: "<p><p>", PreserveParagraphs: true}, 
		},
		{
			name: "multi-line, custom str, custom paraSep (folded), custom lineSep (folded)",
			input: "here is a custom multi-line with unique break,<br/>normal break\n" +
				"folded with custom parasep break<p>\n" +
				"and a normal para end\n\n" +
				"with some more breaks\n" +
				"and another<p>" +
				"custom para end folded with custom parasep<p><p>" +
				"a third possible para\n\n" +
				"finally a completely<p>" +
				"unique para end<para>\n" +
				"and some text<br/>" +
				"to break things up",
			level: 1,
			options: Options{
				IndentStr: ">",
				ParagraphSeparator: "<p><p>",
				LineSeparator: "<p>",
			},
		},
		{
			name: "multi-line, custom str, custom paraSep (non-folded), custom lineSep",
			input: "here is a custom multi-line with unique break,<br/>normal break\n" +
				"folded with custom parasep break<p>\n" +
				"and a normal para end\n\n" +
				"with some more breaks\n" +
				"and another<p>" +
				"custom para end folded with custom parasep<p><p>" +
				"a third possible para\n\n" +
				"finally a completely<p>" +
				"unique para end<para>\n" +
				"and some text<br/>" +
				"to break things up",
			level: 1,
			options: Options{
				IndentStr: ">",
				ParagraphSeparator: "<para>",
				LineSeparator: "<p>",
			},
		},
		{
			name: "multi-line, custom str, custom paraSep, custom lineSep (non-folded)",
			input: "here is a custom multi-line with unique break,<br/>normal break\n" +
				"folded with custom parasep break<p>\n" +
				"and a normal para end\n\n" +
				"with some more breaks\n" +
				"and another<p>" +
				"custom para end folded with custom parasep<p><p>" +
				"a third possible para\n\n" +
				"finally a completely<p>" +
				"unique para end<para>\n" +
				"and some text<br/>" +
				"to break things up",
			level: 1,
			options: Options{
				IndentStr: ">",
				ParagraphSeparator: "<p><p>",
				LineSeparator: "<br/>",
			},
		},
		{
			name: "multi-line, custom str, preserveParas, custom lineSep",
			input: "here is a custom multi-line with unique break,<br/>normal break\n" +
				"folded with custom parasep break<p>\n" +
				"and a normal para end\n\n" +
				"with some more breaks\n" +
				"and another<p>" +
				"custom para end folded with custom parasep<p><p>" +
				"a third possible para\n\n" +
				"finally a completely<p>" +
				"unique para end<para>\n" +
				"and some text<br/>" +
				"to break things up",
			level: 1,
			options: Options{
				IndentStr: ">",
				PreserveParagraphs: true,
				LineSeparator: "<p>",
			},
		},
		{
			name: "multi-line, noTrailing, custom paraSep, preserveParas",
			input: "here is a custom multi-line with unique break,<br/>normal break\n" +
				"folded with custom parasep break<p>\n" +
				"and a normal para end\n\n" +
				"with some more breaks\n" +
				"and another<p>" +
				"custom para end folded with custom parasep<p><p>" +
				"a third possible para\n\n" +
				"finally a completely<p>" +
				"unique para end<para>\n" +
				"and some text<br/>" +
				"to break things up",
			level: 1,
			options: Options{
				NoTrailingLineSeparators: true,
				ParagraphSeparator: "<p><p>",
				PreserveParagraphs: true,
			},
		},
		{
			name: "multi-line, noTrailing, custom paraSep (folded), custom lineSep (folded)",
			input: "here is a custom multi-line with unique break,<br/>normal break\n" +
				"folded with custom parasep break<p>\n" +
				"and a normal para end\n\n" +
				"with some more breaks\n" +
				"and another<p>" +
				"custom para end folded with custom parasep<p><p>" +
				"a third possible para\n\n" +
				"finally a completely<p>" +
				"unique para end<para>\n" +
				"and some text<br/>" +
				"to break things up",
			level: 1,
			options: Options{
				NoTrailingLineSeparators: true,
				ParagraphSeparator: "<p><p>",
				LineSeparator: "<p>",
			},
		},
		{
			name: "multi-line, noTrailing, custom paraSep (non-folded), custom lineSep",
			input: "here is a custom multi-line with unique break,<br/>normal break\n" +
				"folded with custom parasep break<p>\n" +
				"and a normal para end\n\n" +
				"with some more breaks\n" +
				"and another<p>" +
				"custom para end folded with custom parasep<p><p>" +
				"a third possible para\n\n" +
				"finally a completely<p>" +
				"unique para end<para>\n" +
				"and some text<br/>" +
				"to break things up",
			level: 1,
			options: Options{
				NoTrailingLineSeparators: true,
				ParagraphSeparator: "<para>",
				LineSeparator: "<p>",
			},
		},
		{
			name: "multi-line, noTrailing, custom paraSep, custom lineSep (non-folded)",
			input: "here is a custom multi-line with unique break,<br/>normal break\n" +
				"folded with custom parasep break<p>\n" +
				"and a normal para end\n\n" +
				"with some more breaks\n" +
				"and another<p>" +
				"custom para end folded with custom parasep<p><p>" +
				"a third possible para\n\n" +
				"finally a completely<p>" +
				"unique para end<para>\n" +
				"and some text<br/>" +
				"to break things up",
			level: 1,
			options: Options{
				NoTrailingLineSeparators: true,
				ParagraphSeparator: "<p><p>",
				LineSeparator: "<br/>",
			},
		},
		{
			name: "multi-line, noTrailing, preserveParas, custom lineSep",
			input: "here is a custom multi-line with unique break,<br/>normal break\n" +
				"folded with custom parasep break<p>\n" +
				"and a normal para end\n\n" +
				"with some more breaks\n" +
				"and another<p>" +
				"custom para end folded with custom parasep<p><p>" +
				"a third possible para\n\n" +
				"finally a completely<p>" +
				"unique para end<para>\n" +
				"and some text<br/>" +
				"to break things up",
			level: 1,
			options: Options{
				NoTrailingLineSeparators: true,
				PreserveParagraphs: true,
				LineSeparator: "<p>",
			},
		},
		{
			name: "multi-line, custom paraSep (folded), preserveParas, custom lineSep (folded)",
			input: "here is a custom multi-line with unique break,<br/>normal break\n" +
				"folded with custom parasep break<p>\n" +
				"and a normal para end\n\n" +
				"with some more breaks\n" +
				"and another<p>" +
				"custom para end folded with custom parasep<p><p>" +
				"a third possible para\n\n" +
				"finally a completely<p>" +
				"unique para end<para>\n" +
				"and some text<br/>" +
				"to break things up",
			level: 1,
			options: Options{
				ParagraphSeparator: "<p><p>",
				PreserveParagraphs: true,
				LineSeparator: "<p>",
			},
		},
		{
			name: "multi-line, custom paraSep (non-folded), preserveParas, custom lineSep",
			input: "here is a custom multi-line with unique break,<br/>normal break\n" +
				"folded with custom parasep break<p>\n" +
				"and a normal para end\n\n" +
				"with some more breaks\n" +
				"and another<p>" +
				"custom para end folded with custom parasep<p><p>" +
				"a third possible para\n\n" +
				"finally a completely<p>" +
				"unique para end<para>\n" +
				"and some text<br/>" +
				"to break things up",
			level: 1,
			options: Options{
				ParagraphSeparator: "<para>",
				PreserveParagraphs: true,
				LineSeparator: "<p>",
			},
		},
		{
			name: "multi-line, custom paraSep, preserveParas, custom lineSep (non-folded)",
			input: "here is a custom multi-line with unique break,<br/>normal break\n" +
				"folded with custom parasep break<p>\n" +
				"and a normal para end\n\n" +
				"with some more breaks\n" +
				"and another<p>" +
				"custom para end folded with custom parasep<p><p>" +
				"a third possible para\n\n" +
				"finally a completely<p>" +
				"unique para end<para>\n" +
				"and some text<br/>" +
				"to break things up",
			level: 1,
			options: Options{
				ParagraphSeparator: "<p><p>",
				PreserveParagraphs: true,
				LineSeparator: "<br/>",
			},
		},
		{
			name: "multi-line, custom str, noTrailing, custom paraSep, preserveParas",
			input: "here is a custom multi-line with unique break,<br/>normal break\n" +
				"folded with custom parasep break<p>\n" +
				"and a normal para end\n\n" +
				"with some more breaks\n" +
				"and another<p>" +
				"custom para end folded with custom parasep<p><p>" +
				"a third possible para\n\n" +
				"finally a completely<p>" +
				"unique para end<para>\n" +
				"and some text<br/>" +
				"to break things up",
			level: 1,
			options: Options{
				IndentStr: ">",
				NoTrailingLineSeparators: true,
				ParagraphSeparator: "<p><p>",
				PreserveParagraphs: true,
			},
		},
		{
			name: "multi-line, custom str, noTrailing, custom paraSep (folded), custom lineSep (folded)",
			input: "here is a custom multi-line with unique break,<br/>normal break\n" +
				"folded with custom parasep break<p>\n" +
				"and a normal para end\n\n" +
				"with some more breaks\n" +
				"and another<p>" +
				"custom para end folded with custom parasep<p><p>" +
				"a third possible para\n\n" +
				"finally a completely<p>" +
				"unique para end<para>\n" +
				"and some text<br/>" +
				"to break things up",
			level: 1,
			options: Options{
				IndentStr: ">",
				NoTrailingLineSeparators: true,
				ParagraphSeparator: "<p><p>",
				LineSeparator: "<p>",
			},
		},
		{
			name: "multi-line, custom str, noTrailing, custom paraSep (non-folded), custom lineSep",
			input: "here is a custom multi-line with unique break,<br/>normal break\n" +
				"folded with custom parasep break<p>\n" +
				"and a normal para end\n\n" +
				"with some more breaks\n" +
				"and another<p>" +
				"custom para end folded with custom parasep<p><p>" +
				"a third possible para\n\n" +
				"finally a completely<p>" +
				"unique para end<para>\n" +
				"and some text<br/>" +
				"to break things up",
			level: 1,
			options: Options{
				IndentStr: ">",
				NoTrailingLineSeparators: true,
				ParagraphSeparator: "<para>",
				LineSeparator: "<p>",
			},
		},
		{
			name: "multi-line, custom str, noTrailing, custom paraSep, custom lineSep (non-folded)",
			input: "here is a custom multi-line with unique break,<br/>normal break\n" +
				"folded with custom parasep break<p>\n" +
				"and a normal para end\n\n" +
				"with some more breaks\n" +
				"and another<p>" +
				"custom para end folded with custom parasep<p><p>" +
				"a third possible para\n\n" +
				"finally a completely<p>" +
				"unique para end<para>\n" +
				"and some text<br/>" +
				"to break things up",
			level: 1,
			options: Options{
				IndentStr: ">",
				NoTrailingLineSeparators: true,
				ParagraphSeparator: "<p><p>",
				LineSeparator: "<br/>",
			},
		},
		{
			name: "multi-line, custom str, noTrailing, preserveParas, custom lineSep",
			input: "here is a custom multi-line with unique break,<br/>normal break\n" +
				"folded with custom parasep break<p>\n" +
				"and a normal para end\n\n" +
				"with some more breaks\n" +
				"and another<p>" +
				"custom para end folded with custom parasep<p><p>" +
				"a third possible para\n\n" +
				"finally a completely<p>" +
				"unique para end<para>\n" +
				"and some text<br/>" +
				"to break things up",
			level: 1,
			options: Options{
				IndentStr: ">",
				NoTrailingLineSeparators: true,
				PreserveParagraphs: true,
				LineSeparator: "<p>",
			},
		},
		{
			name: "multi-line, custom str, custom paraSep (folded), preserveParas, custom lineSep (folded)",
			input: "here is a custom multi-line with unique break,<br/>normal break\n" +
				"folded with custom parasep break<p>\n" +
				"and a normal para end\n\n" +
				"with some more breaks\n" +
				"and another<p>" +
				"custom para end folded with custom parasep<p><p>" +
				"a third possible para\n\n" +
				"finally a completely<p>" +
				"unique para end<para>\n" +
				"and some text<br/>" +
				"to break things up",
			level: 1,
			options: Options{
				IndentStr: ">",
				ParagraphSeparator: "<p><p>",
				PreserveParagraphs: true,
				LineSeparator: "<p>",
			},
		},
		{
			name: "multi-line, custom str, custom paraSep (non-folded), preserveParas, custom lineSep",
			input: "here is a custom multi-line with unique break,<br/>normal break\n" +
				"folded with custom parasep break<p>\n" +
				"and a normal para end\n\n" +
				"with some more breaks\n" +
				"and another<p>" +
				"custom para end folded with custom parasep<p><p>" +
				"a third possible para\n\n" +
				"finally a completely<p>" +
				"unique para end<para>\n" +
				"and some text<br/>" +
				"to break things up",
			level: 1,
			options: Options{
				IndentStr: ">",
				ParagraphSeparator: "<para>",
				PreserveParagraphs: true,
				LineSeparator: "<p>",
			},
		},
		{
			name: "multi-line, custom str, custom paraSep, preserveParas, custom lineSep (non-folded)",
			input: "here is a custom multi-line with unique break,<br/>normal break\n" +
				"folded with custom parasep break<p>\n" +
				"and a normal para end\n\n" +
				"with some more breaks\n" +
				"and another<p>" +
				"custom para end folded with custom parasep<p><p>" +
				"a third possible para\n\n" +
				"finally a completely<p>" +
				"unique para end<para>\n" +
				"and some text<br/>" +
				"to break things up",
			level: 1,
			options: Options{
				IndentStr: ">",
				ParagraphSeparator: "<p><p>",
				PreserveParagraphs: true,
				LineSeparator: "<br/>",
			},
		},
		{
			name: "multi-line, noTrailing, custom paraSep (folded), preserveParas, custom lineSep (folded)",
			input: "here is a custom multi-line with unique break,<br/>normal break\n" +
				"folded with custom parasep break<p>\n" +
				"and a normal para end\n\n" +
				"with some more breaks\n" +
				"and another<p>" +
				"custom para end folded with custom parasep<p><p>" +
				"a third possible para\n\n" +
				"finally a completely<p>" +
				"unique para end<para>\n" +
				"and some text<br/>" +
				"to break things up",
			level: 1,
			options: Options{
				NoTrailingLineSeparators: true,
				ParagraphSeparator: "<p><p>",
				PreserveParagraphs: true,
				LineSeparator: "<p>",
			},
		},
		{
			name: "multi-line, noTrailing, custom paraSep (non-folded), preserveParas, custom lineSep",
			input: "here is a custom multi-line with unique break,<br/>normal break\n" +
				"folded with custom parasep break<p>\n" +
				"and a normal para end\n\n" +
				"with some more breaks\n" +
				"and another<p>" +
				"custom para end folded with custom parasep<p><p>" +
				"a third possible para\n\n" +
				"finally a completely<p>" +
				"unique para end<para>\n" +
				"and some text<br/>" +
				"to break things up",
			level: 1,
			options: Options{
				NoTrailingLineSeparators: true,
				ParagraphSeparator: "<para>",
				PreserveParagraphs: true,
				LineSeparator: "<p>",
			},
		},
		{
			name: "multi-line, noTrailing, custom paraSep, preserveParas, custom lineSep (non-folded)",
			input: "here is a custom multi-line with unique break,<br/>normal break\n" +
				"folded with custom parasep break<p>\n" +
				"and a normal para end\n\n" +
				"with some more breaks\n" +
				"and another<p>" +
				"custom para end folded with custom parasep<p><p>" +
				"a third possible para\n\n" +
				"finally a completely<p>" +
				"unique para end<para>\n" +
				"and some text<br/>" +
				"to break things up",
			level: 1,
			options: Options{
				NoTrailingLineSeparators: true,
				ParagraphSeparator: "<p><p>",
				PreserveParagraphs: true,
				LineSeparator: "<br/>",
			},
		},
		{
			name: "multi-line, custom str, noTrailing, custom paraSep (folded), preserveParas, custom lineSep (folded)",
			input: "here is a custom multi-line with unique break,<br/>normal break\n" +
				"folded with custom parasep break<p>\n" +
				"and a normal para end\n\n" +
				"with some more breaks\n" +
				"and another<p>" +
				"custom para end folded with custom parasep<p><p>" +
				"a third possible para\n\n" +
				"finally a completely<p>" +
				"unique para end<para>\n" +
				"and some text<br/>" +
				"to break things up",
			level: 1,
			options: Options{
				IndentStr: ">",
				NoTrailingLineSeparators: true,
				ParagraphSeparator: "<p><p>",
				PreserveParagraphs: true,
				LineSeparator: "<p>",
			},
		},
		{
			name: "multi-line, custom str, noTrailing, custom paraSep (non-folded), preserveParas, custom lineSep",
			input: "here is a custom multi-line with unique break,<br/>normal break\n" +
				"folded with custom parasep break<p>\n" +
				"and a normal para end\n\n" +
				"with some more breaks\n" +
				"and another<p>" +
				"custom para end folded with custom parasep<p><p>" +
				"a third possible para\n\n" +
				"finally a completely<p>" +
				"unique para end<para>\n" +
				"and some text<br/>" +
				"to break things up",
			level: 1,
			options: Options{
				IndentStr: ">",
				NoTrailingLineSeparators: true,
				ParagraphSeparator: "<para>",
				PreserveParagraphs: true,
				LineSeparator: "<p>",
			},
		},
		{
			name: "multi-line, custom str, noTrailing, custom paraSep, preserveParas, custom lineSep (non-folded)",
			input: "here is a custom multi-line with unique break,<br/>normal break\n" +
				"folded with custom parasep break<p>\n" +
				"and a normal para end\n\n" +
				"with some more breaks\n" +
				"and another<p>" +
				"custom para end folded with custom parasep<p><p>" +
				"a third possible para\n\n" +
				"finally a completely<p>" +
				"unique para end<para>\n" +
				"and some text<br/>" +
				"to break things up",
			level: 1,
			options: Options{
				IndentStr: ">",
				NoTrailingLineSeparators: true,
				ParagraphSeparator: "<p><p>",
				PreserveParagraphs: true,
				LineSeparator: "<br/>",
			},
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

