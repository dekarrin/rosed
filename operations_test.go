package rosed

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Wrap(t *testing.T) {
	testCases := []struct {
		name   string
		input  string
		width  int
		expect []string
	}{
		{
			name:  "empty string",
			input: "",
			width: 20,
			expect: []string{
				"",
			},
		},
		{
			name:  "1 line",
			input: "Things will never stop from keep happening constantly.",
			width: 20,
			expect: []string{
				"Things will never",
				"stop from keep",
				"happening",
				"constantly.",
			},
		},
		{
			name:  "2 paragraphs are joined into one",
			input: "this is a line that is split by paragraph in the input.\n\nThis is the second paragraph",
			width: 20,
			expect: []string{
				"this is a line that",
				"is split by",
				"paragraph in the",
				"input. This is the",
				"second paragraph",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert := assert.New(t)

			// try via Edit().WithOptions(), no defaults set
			actualText := Edit(tc.input).Wrap(tc.width).String()

			// String returns one string; turn it into actual lines
			actual := strings.Split(actualText, DefaultLineSeparator)
			assert.Equal(tc.expect, actual)
		})
	}
}

func Test_WrapOpts(t *testing.T) {
	testCases := []struct {
		name    string
		input   string
		width   int
		options Options
		expect  []string
	}{
		{
			name:    "2 paragraphs, not preserved",
			input:   "this is a line that is split by paragraph in the input.\n\nThis is the second paragraph",
			width:   20,
			options: Options{PreserveParagraphs: false},
			expect: []string{
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
			expect: []string{
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
			expect: []string{
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
			expect: []string{
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

			actualDirectText := Edit(tc.input).WrapOpts(tc.width, tc.options).String()
			actualPreOptsText := Edit(tc.input).WithOptions(tc.options).Wrap(tc.width).String()

			actualDirect := strings.Split(actualDirectText, tc.options.WithDefaults().LineSeparator)
			actualPreOpts := strings.Split(actualPreOptsText, tc.options.WithDefaults().LineSeparator)

			assert.Equal(tc.expect, actualDirect, "WrapOpts(opts) check failed")
			assert.Equal(tc.expect, actualPreOpts, "WithOptions(opts).Wrap() check failed")
		})
	}
}

func Test_CollapseSpace(t *testing.T) {
	testCases := []struct {
		name   string
		input  string
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

func Test_CollapseSpaceOpts(t *testing.T) {
	testCases := []struct {
		name    string
		input   string
		options Options
		expect  string
	}{
		{
			name:    "empty line, custom lineSep",
			input:   "",
			options: Options{LineSeparator: "<P>"},
			expect:  "",
		},
		{
			name:    "no spaces, custom lineSep",
			input:   "bluh",
			options: Options{LineSeparator: "<P>"},
			expect:  "bluh",
		},
		{
			name:    "2 words, custom lineSep",
			input:   "word1 word2",
			options: Options{LineSeparator: "<P>"},
			expect:  "word1 word2",
		},
		{
			name:    "3 words, custom lineSep",
			input:   "word1 word2 word3",
			options: Options{LineSeparator: "<P>"},
			expect:  "word1 word2 word3",
		},
		{
			name:    "3 words with runs of spaces, custom lineSep",
			input:   "word1        word2  word3",
			options: Options{LineSeparator: "<P>"},
			expect:  "word1 word2 word3",
		},
		{
			name:    "run of non-uniform spaces",
			input:   "word1 \t\t word2",
			options: Options{LineSeparator: "<P>"},
			expect:  "word1 word2",
		},
		{
			name:    "include lineSep",
			input:   "   <P> word1",
			options: Options{LineSeparator: "<P>"},
			expect:  " word1",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert := assert.New(t)

			actualDirect := Edit(tc.input).CollapseSpaceOpts(tc.options).String()
			actualPreOpts := Edit(tc.input).WithOptions(tc.options).CollapseSpace().String()

			assert.Equal(tc.expect, actualDirect, "CollapseSpaceOpts(opts) check failed")
			assert.Equal(tc.expect, actualPreOpts, "WithOptions(opts).CollapseSpace() check failed")
		})
	}
}

func Test_Insert(t *testing.T) {
	testCases := []struct {
		name   string
		input  string
		pos    int
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
	testCases := []struct {
		name   string
		input  string
		op     LineOperation
		expect string
	}{
		{
			name:  "apply nothing to empty editor",
			input: "",
			op: func(lineNo int, line string) []string {
				return []string{"test"}
			},
			expect: "",
		},
		{
			name:  "replace one terminated empty line",
			input: DefaultLineSeparator,
			op: func(lineNo int, line string) []string {
				return []string{"test"}
			},
			expect: "test" + DefaultLineSeparator,
		},
		{
			name:  "replace two terminated lines",
			input: DefaultLineSeparator + DefaultLineSeparator,
			op: func(lineNo int, line string) []string {
				return []string{"test"}
			},
			expect: "test" + DefaultLineSeparator + "test" + DefaultLineSeparator,
		},
		{
			name:  "replace multiple lines using line number",
			input: DefaultLineSeparator + DefaultLineSeparator + DefaultLineSeparator,
			op: func(lineNo int, line string) []string {
				newLine := fmt.Sprintf("test%d", lineNo)
				return []string{newLine}
			},
			expect: "test0" + DefaultLineSeparator + "test1" + DefaultLineSeparator + "test2" + DefaultLineSeparator,
		},
		{
			name:  "insert extra line at target position",
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
			name:  "delete line at target position",
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
	testCases := []struct {
		name    string
		input   string
		op      LineOperation
		options Options
		expect  string
	}{
		{
			name:  "apply nothing to empty editor, custom lineSep",
			input: "",
			op: func(lineNo int, line string) []string {
				return []string{"test"}
			},
			options: Options{LineSeparator: "<P>"},
			expect:  "",
		},
		{
			name:  "replace one terminated empty line, custom lineSep",
			input: "<P>",
			op: func(lineNo int, line string) []string {
				return []string{"test"}
			},
			options: Options{LineSeparator: "<P>"},
			expect:  "test<P>",
		},
		{
			name:  "replace two terminated lines, custom lineSep",
			input: "<P><P>",
			op: func(lineNo int, line string) []string {
				return []string{"test"}
			},
			options: Options{LineSeparator: "<P>"},
			expect:  "test<P>test<P>",
		},
		{
			name:  "replace multiple lines using line number, custom lineSep",
			input: "<P><P><P>",
			op: func(lineNo int, line string) []string {
				newLine := fmt.Sprintf("test%d", lineNo)
				return []string{newLine}
			},
			options: Options{LineSeparator: "<P>"},
			expect:  "test0<P>test1<P>test2<P>",
		},
		{
			name:  "insert extra line at target position, custom lineSep",
			input: "line0<P>line2<P>",
			op: func(lineNo int, line string) []string {
				if lineNo == 0 {
					return []string{line, "line1"}
				}
				return []string{line}
			},
			options: Options{LineSeparator: "<P>"},
			expect:  "line0<P>line1<P>line2<P>",
		},
		{
			name:  "delete line at target position, custom lineSep",
			input: "line0<P>extra<P>line1<P>",
			op: func(lineNo int, line string) []string {
				if lineNo == 1 {
					return []string{}
				}
				return []string{line}
			},
			options: Options{LineSeparator: "<P>"},
			expect:  "line0<P>line1<P>",
		},
		{
			name:  "apply once for empty editor, noTrailing=true",
			input: "",
			op: func(lineNo int, line string) []string {
				return []string{"test"}
			},
			options: Options{NoTrailingLineSeparators: true},
			expect:  "test",
		},
		{
			name:  "apply n+1 for n lines, noTrailing=true",
			input: DefaultLineSeparator + DefaultLineSeparator,
			op: func(lineNo int, line string) []string {
				return []string{"test"}
			},
			options: Options{NoTrailingLineSeparators: true},
			expect:  "test" + DefaultLineSeparator + "test" + DefaultLineSeparator + "test",
		},
		{
			name:  "apply once for empty editor, noTrailing=true, custom lineSep",
			input: "",
			op: func(lineNo int, line string) []string {
				return []string{"test"}
			},
			options: Options{NoTrailingLineSeparators: true, LineSeparator: "<P>"},
			expect:  "test",
		},
		{
			name:  "apply n+1 for n lines, noTrailing=true, custom lineSep",
			input: "<P><P>",
			op: func(lineNo int, line string) []string {
				return []string{"test"}
			},
			options: Options{NoTrailingLineSeparators: true, LineSeparator: "<P>"},
			expect:  "test<P>test<P>test",
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
	testCases := []struct {
		name   string
		input  string
		op     ParagraphOperation
		expect string
	}{
		{
			name:  "empty string = 1 empty para",
			input: "",
			op: func(idx int, para, sepPrefix, sepSuffix string) []string {
				return []string{"test"}
			},
			expect: "test",
		},
		{
			name:  "replace two empty paras",
			input: DefaultParagraphSeparator,
			op: func(idx int, para, sepPrefix, sepSuffix string) []string {
				return []string{"test"}
			},
			expect: "test" + DefaultParagraphSeparator + "test",
		},
		{
			name:  "replace multiple paras using para number",
			input: DefaultParagraphSeparator + DefaultParagraphSeparator,
			op: func(idx int, para, sepPrefix, sepSuffix string) []string {
				newPara := fmt.Sprintf("test%d", idx)
				return []string{newPara}
			},
			expect: "test0" + DefaultParagraphSeparator + "test1" + DefaultParagraphSeparator + "test2",
		},
		{
			name:  "insert extra para at target position",
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
			name:  "delete para at target position",
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
	testCases := []struct {
		name    string
		input   string
		op      ParagraphOperation
		options Options
		expect  string
	}{
		{
			name:  "apply once to empty editor, custom paraSep",
			input: "",
			op: func(idx int, para, sepPrefix, sepSuffix string) []string {
				return []string{"test"}
			},
			options: Options{ParagraphSeparator: "<P>"},
			expect:  "test",
		},
		{
			name:  "replace two empty paras, custom paraSep",
			input: "<P>",
			op: func(idx int, para, sepPrefix, sepSuffix string) []string {
				return []string{"test"}
			},
			options: Options{ParagraphSeparator: "<P>"},
			expect:  "test<P>test",
		},
		{
			name:  "replace multiple paras using para number, custom paraSep",
			input: "<P><P>",
			op: func(idx int, para, sepPrefix, sepSuffix string) []string {
				newPara := fmt.Sprintf("test%d", idx)
				return []string{newPara}
			},
			options: Options{ParagraphSeparator: "<P>"},
			expect:  "test0<P>test1<P>test2",
		},
		{
			name:  "insert extra para at target position, custom paraSep",
			input: "para0<P>para2",
			op: func(idx int, para, sepPrefix, sepSuffix string) []string {
				if idx == 0 {
					return []string{para, "para1"}
				}
				return []string{para}
			},
			options: Options{ParagraphSeparator: "<P>"},
			expect:  "para0<P>para1<P>para2",
		},
		{
			name:  "delete para at target position, custom paraSep",
			input: "para0<P>extra<P>para1",
			op: func(idx int, para, sepPrefix, sepSuffix string) []string {
				if idx == 1 {
					return []string{}
				}
				return []string{para}
			},
			options: Options{ParagraphSeparator: "<P>"},
			expect:  "para0<P>para1",
		},
		{
			name:  "custom paraSep with a line-broken parasep",
			input: "para1<P1>\n<P2>para2<P1>\n<P2>para3<P1>\n<P2>para4",
			op: func(idx int, para, sepPrefix, sepSuffix string) []string {
				newPara := fmt.Sprintf("(PREFIX=%s,TEXT=%s,SUFFIX=%s)", sepPrefix, para, sepSuffix)

				return []string{newPara}
			},
			options: Options{ParagraphSeparator: "<P1>\n<P2>"},
			expect: "(PREFIX=,TEXT=para1,SUFFIX=<P1>)<P1>\n" +
				"<P2>(PREFIX=<P2>,TEXT=para2,SUFFIX=<P1>)<P1>\n" +
				"<P2>(PREFIX=<P2>,TEXT=para3,SUFFIX=<P1>)<P1>\n" +
				"<P2>(PREFIX=<P2>,TEXT=para4,SUFFIX=)",
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
	testCases := []struct {
		name   string
		input  string
		level  int
		expect string
	}{
		{
			name:   "do nothing to empty editor",
			input:  "",
			level:  20,
			expect: "",
		},
		{
			name:   "empty line, 0 levels",
			input:  DefaultLineSeparator,
			level:  0,
			expect: DefaultLineSeparator,
		},
		{
			name:   "empty line, -1 levels",
			input:  DefaultLineSeparator,
			level:  -1,
			expect: DefaultLineSeparator,
		},
		{
			name:   "empty line, 1 level",
			input:  DefaultLineSeparator,
			level:  1,
			expect: DefaultIndentString + DefaultLineSeparator,
		},
		{
			name:   "empty line, multiple levels",
			input:  DefaultLineSeparator,
			level:  3,
			expect: DefaultIndentString + DefaultIndentString + DefaultIndentString + DefaultLineSeparator,
		},
		{
			name:   "non-empty line, 0 levels",
			input:  "test" + DefaultLineSeparator,
			level:  0,
			expect: "test" + DefaultLineSeparator,
		},
		{
			name:   "non-empty string, -1 levels",
			input:  "test" + DefaultLineSeparator,
			level:  -1,
			expect: "test" + DefaultLineSeparator,
		},
		{
			name:   "non-empty string, 1 level",
			input:  "test" + DefaultLineSeparator,
			level:  1,
			expect: DefaultIndentString + "test" + DefaultLineSeparator,
		},
		{
			name:   "non-empty string, multiple levels",
			input:  "test" + DefaultLineSeparator,
			level:  3,
			expect: DefaultIndentString + DefaultIndentString + DefaultIndentString + "test" + DefaultLineSeparator,
		},
		{
			name:  "multi-line string, 0 levels",
			input: "test0" + DefaultLineSeparator + "test1" + DefaultLineSeparator + "test2" + DefaultLineSeparator,
			level: 0,
			expect: "test0" + DefaultLineSeparator +
				"test1" + DefaultLineSeparator +
				"test2" + DefaultLineSeparator,
		},
		{
			name:  "multi-line string, -1 levels",
			input: "test0" + DefaultLineSeparator + "test1" + DefaultLineSeparator + "test2" + DefaultLineSeparator,
			level: -1,
			expect: "test0" + DefaultLineSeparator +
				"test1" + DefaultLineSeparator +
				"test2" + DefaultLineSeparator,
		},
		{
			name:  "multi-line string, 1 levels",
			input: "test0" + DefaultLineSeparator + "test1" + DefaultLineSeparator + "test2" + DefaultLineSeparator,
			level: 1,
			expect: DefaultIndentString + "test0" + DefaultLineSeparator +
				DefaultIndentString + "test1" + DefaultLineSeparator +
				DefaultIndentString + "test2" + DefaultLineSeparator,
		},
		{
			name:  "multi-line string, multiple levels",
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
	testCases := []struct {
		name    string
		input   string
		level   int
		options Options
		expect  string
	}{
		{
			name:  "empty line, noTrailing",
			input: "",
			level: 1,
			options: Options{
				NoTrailingLineSeparators: true,
			},
			expect: DefaultIndentString,
		},
		{
			name:  "empty line, custom str",
			input: DefaultLineSeparator,
			level: 1,
			options: Options{
				IndentStr: ">",
			},
			expect: ">" + DefaultLineSeparator,
		},
		{
			name:  "empty line, custom str, noTrailing",
			input: "",
			level: 1,
			options: Options{
				IndentStr:                ">",
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
			name:  "multi-line, noTrailing",
			input: "line0\nline1\n",
			level: 1,
			options: Options{
				NoTrailingLineSeparators: true,
			},
			expect: "\tline0\n\tline1\n\t",
		},
		{
			name:  "multi-line, custom paraSep",
			input: "p0,line0\np0,line1\np0,line2<p><p>p1,line0\np1,line1",
			level: 1,
			options: Options{
				ParagraphSeparator: "<p><p>",
			},
			expect: "\tp0,line0\n\tp0,line1\n\tp0,line2<p><p>p1,line0\n\tp1,line1",
		},
		{
			name:  "multi-line, preserveParas",
			input: "p0,line0\np0,line1\n\np1,line0\np1,line1",
			level: 1,
			options: Options{
				PreserveParagraphs: true,
			},
			expect: "\tp0,line0\n\tp0,line1\n\n\tp1,line0\n\tp1,line1",
		},
		{
			name:  "multi-line, custom lineSep",
			input: "line0\nmore line0<p>line1<p>line2\nbreak ignored",
			level: 1,
			options: Options{
				LineSeparator: "<p>",
			},
			expect: "\tline0\nmore line0<p>\tline1<p>\tline2\nbreak ignored",
		},
		{
			name:  "multi-line, custom str, noTrailing",
			input: "line0\nline1\nline2\n",
			level: 1,
			options: Options{
				IndentStr:                ">",
				NoTrailingLineSeparators: true,
			},
			expect: ">line0\n>line1\n>line2\n>",
		},
		{
			name:  "multi-line, custom str, custom paraSep",
			input: "p0,line0\np0,line1\np0,line2<p><p>p1,line0\np1,line1",
			level: 1,
			options: Options{
				IndentStr:          ">",
				ParagraphSeparator: "<p><p>",
			},
			expect: ">p0,line0\n>p0,line1\n>p0,line2<p><p>p1,line0\n>p1,line1",
		},
		{
			name:  "multi-line, custom str, preserveParas",
			input: "p0,line0\np0,line1\n\np1,line0\np1,line1",
			level: 1,
			options: Options{
				IndentStr:          ">",
				PreserveParagraphs: true,
			},
			expect: ">p0,line0\n>p0,line1\n\n>p1,line0\n>p1,line1",
		},
		{
			name:  "multi-line, custom str, custom lineSep",
			input: "line0\nmore line0<p>line1<p>line2\nbreak ignored",
			level: 1,
			options: Options{
				IndentStr:     ">",
				LineSeparator: "<p>",
			},
			expect: ">line0\nmore line0<p>>line1<p>>line2\nbreak ignored",
		},
		{
			name:  "multi-line, custom paraSep, preserveParas",
			input: "p0,A\np0,B\n\np0,C<p><p>p1,A\np1,B<p><p>p2,A\np2,B",
			level: 1,
			options: Options{
				ParagraphSeparator: "<p><p>",
				PreserveParagraphs: true,
			},
			expect: "\tp0,A\n\tp0,B\n\t\n\tp0,C<p><p>\tp1,A\n\tp1,B<p><p>\tp2,A\n\tp2,B",
		},
		{
			name:  "multi-line, preserveParas, custom lineSep",
			input: "p0,l0\nmore l0<p>p0,l1\n\np1,l0<p>p1,l1<p>p1,l2\nmore l2",
			level: 1,
			options: Options{
				PreserveParagraphs: true,
				LineSeparator:      "<p>",
			},
			expect: "\tp0,l0\nmore l0<p>\tp0,l1\n\n\tp1,l0<p>\tp1,l1<p>\tp1,l2\nmore l2",
		},
		{
			name:  "multi-line, custom paraSep (folded), preserveParas, custom lineSep (folded)",
			input: "p0,l0<p>p0,l1\nmore l1<p><p>p1,l0\n\nmore l0<p>line1<p><p><p><p>next para<p>pn,ln",
			level: 1,
			options: Options{
				ParagraphSeparator: "<p><p>",
				PreserveParagraphs: true,
				LineSeparator:      "<p>",
			},
			expect: "\tp0,l0<p>\tp0,l1\nmore l1<p><p>\tp1,l0\n\nmore l0<p>\tline1<p><p><p><p>\tnext para<p>\tpn,ln",
		},
		{
			name:  "multi-line, custom paraSep (non-folded), preserveParas, custom lineSep",
			input: "p0,l0<p>p0,l1\nmore l1<para>p1,l0\n\nmore l0<p>line1<para><para>next para<p>pn,ln",
			level: 1,
			options: Options{
				ParagraphSeparator: "<para>",
				PreserveParagraphs: true,
				LineSeparator:      "<p>",
			},
			expect: "\tp0,l0<p>\tp0,l1\nmore l1<para>\tp1,l0\n\nmore l0<p>\tline1<para><para>\tnext para<p>\tpn,ln",
		},
		{
			name:  "multi-line, custom str, noTrailing, custom paraSep (folded), preserveParas, custom lineSep",
			input: "p0,line0<p>p0,line1<p><p><p>p1,line0<p>p1,line1<p><p>p2,line1<p>p2,line2<p>",
			level: 1,
			options: Options{
				IndentStr:                ">",
				NoTrailingLineSeparators: true,
				ParagraphSeparator:       "<p><p>",
				PreserveParagraphs:       true,
				LineSeparator:            "<p>",
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
	testCases := []struct {
		name   string
		input  string
		width  int
		expect string
	}{
		{
			name:   "empty string",
			input:  "",
			width:  10,
			expect: "",
		},
		{
			name:   "no spaces",
			input:  "bluh",
			width:  10,
			expect: "bluh",
		},
		{
			name:   "2 words",
			input:  "word1 word2",
			width:  20,
			expect: "word1          word2",
		},
		{
			name:   "3 words",
			input:  "word1 word2 word3",
			width:  20,
			expect: "word1   word2  word3",
		},
		{
			name:   "3 words with runs of spaces",
			input:  "word1        word2  word3",
			width:  20,
			expect: "word1   word2  word3",
		},
		{
			name:   "line longer than width",
			input:  "hello",
			width:  3,
			expect: "hello",
		},
		{
			name:   "bad width",
			input:  "bluh",
			width:  -1,
			expect: "bluh",
		},
		{
			name: "multi-line",
			input: "a set of three lines" + DefaultLineSeparator +
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
	testCases := []struct {
		name    string
		input   string
		width   int
		options Options
		expect  string
	}{
		{
			name: "multi-line, noTrailing",
			input: "a set of three lines" + DefaultLineSeparator +
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
			input: "a set of three lines" + DefaultLineSeparator +
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
			input: "a set of three lines" + DefaultLineSeparator +
				"to justify in a" + DefaultLineSeparator +
				"good way" +
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
				"good        way" + // not taking up full length to acct for parasep that takes space
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
		pos         int
		left        string
		right       string
		minBetween  int
		width       int
		leftPercent float64
	}

	testCases := []struct {
		name   string
		args   args
		input  string
		expect string
	}{
		{
			name: "empty lines",
			args: args{
				pos:         0,
				left:        "",
				right:       "",
				minBetween:  0,
				width:       0,
				leftPercent: 0.0,
			},
			input:  "",
			expect: "",
		},
		{
			name: "right col bigger",
			args: args{
				pos:         0,
				left:        "This is a test string for the left side",
				right:       "Column number two is right here! And it has a lot of content that will be wrapped",
				minBetween:  2,
				width:       30,
				leftPercent: 0.5,
			},
			input: "",
			expect: "This is a test  Column number" + DefaultLineSeparator +
				"string for the  two is right" + DefaultLineSeparator +
				"left side       here! And it" + DefaultLineSeparator +
				"                has a lot of" + DefaultLineSeparator +
				"                content that" + DefaultLineSeparator +
				"                will be" + DefaultLineSeparator +
				"                wrapped" + DefaultLineSeparator,
		},
		{
			name: "left col bigger",
			args: args{
				pos:         0,
				left:        "Column number one is right here! And it has a lot of content that will be included",
				right:       "This is a test string for the right side",
				minBetween:  2,
				width:       30,
				leftPercent: 0.5,
			},
			input: "",
			expect: "Column number   This is a test" + DefaultLineSeparator +
				"one is right    string for the" + DefaultLineSeparator +
				"here! And it    right side" + DefaultLineSeparator +
				"has a lot of    " + DefaultLineSeparator +
				"content that    " + DefaultLineSeparator +
				"will be         " + DefaultLineSeparator +
				"included        " + DefaultLineSeparator,
		},
		{
			name: "equal size columns",
			args: args{
				pos:         0,
				left:        "Column number one is right here!",
				right:       "This is a test string for the right side",
				minBetween:  2,
				width:       30,
				leftPercent: 0.5,
			},
			input: "",
			expect: "Column number   This is a test" + DefaultLineSeparator +
				"one is right    string for the" + DefaultLineSeparator +
				"here!           right side" + DefaultLineSeparator,
		},
		{
			name: "insert in middle",
			args: args{
				pos:         2,
				left:        "Column number one is right here!",
				right:       "This is a test string for the right side",
				minBetween:  2,
				width:       30,
				leftPercent: 0.5,
			},
			input: "hello",
			expect: "heColumn number   This is a test" + DefaultLineSeparator +
				"one is right    string for the" + DefaultLineSeparator +
				"here!           right side" + DefaultLineSeparator +
				"llo",
		},
		{
			name: "quarter left col",
			args: args{
				pos:         0,
				left:        "Column number one is right here!",
				right:       "This is a test string for the right side",
				minBetween:  2,
				width:       32,
				leftPercent: 0.25,
			},
			input: "",
			expect: "Column   This is a test string" + DefaultLineSeparator +
				"number   for the right side" + DefaultLineSeparator +
				"one is   " + DefaultLineSeparator +
				"right    " + DefaultLineSeparator +
				"here!    " + DefaultLineSeparator,
		},
		{
			name: "quarter right col",
			args: args{
				pos:         0,
				left:        "Column number one is right here!",
				right:       "This is a test string for the right side",
				minBetween:  2,
				width:       32,
				leftPercent: 0.75,
			},
			input: "",
			expect: "Column number one is    This is" + DefaultLineSeparator +
				"right here!             a test" + DefaultLineSeparator +
				"                        string" + DefaultLineSeparator +
				"                        for the" + DefaultLineSeparator +
				"                        right" + DefaultLineSeparator +
				"                        side" + DefaultLineSeparator,
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

func Test_InsertTwoColumnsOpts(t *testing.T) {
	type args struct {
		pos         int
		left        string
		right       string
		minBetween  int
		width       int
		leftPercent float64
	}

	testCases := []struct {
		name    string
		args    args
		options Options
		input   string
		expect  string
	}{
		{
			name: "noTrailing = true",
			args: args{
				pos:         0,
				left:        "Column number one is right here!",
				right:       "This is a test string for the right side",
				minBetween:  2,
				width:       30,
				leftPercent: 0.5,
			},
			options: Options{
				NoTrailingLineSeparators: true,
			},
			input: "",
			expect: "Column number   This is a test" + DefaultLineSeparator +
				"one is right    string for the" + DefaultLineSeparator +
				"here!           right side",
		},
		{
			name: "custom line separator",
			args: args{
				pos:         0,
				left:        "Column number one is right here!",
				right:       "This is a test string for the right side",
				minBetween:  2,
				width:       30,
				leftPercent: 0.5,
			},
			options: Options{
				LineSeparator: "<br/>",
			},
			input: "",
			expect: "Column number   This is a test<br/>" +
				"one is right    string for the<br/>" +
				"here!           right side<br/>",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert := assert.New(t)

			actualDirect := Edit(tc.input).InsertTwoColumnsOpts(
				tc.args.pos,
				tc.args.left,
				tc.args.right,
				tc.args.minBetween,
				tc.args.width,
				tc.args.leftPercent,
				tc.options,
			).String()

			actualPreOpts := Edit(tc.input).WithOptions(tc.options).InsertTwoColumns(
				tc.args.pos,
				tc.args.left,
				tc.args.right,
				tc.args.minBetween,
				tc.args.width,
				tc.args.leftPercent,
			).String()

			assert.Equal(tc.expect, actualDirect, "InsertTwoColumnsOpts(opts) check failed")
			assert.Equal(tc.expect, actualPreOpts, "WithOptions(opts).InsertTwoColumns() check failed")
		})
	}
}

func Test_InsertDefinitionsTable(t *testing.T) {
	testCases := []struct {
		name   string
		input  string
		pos    int
		defs   [][2]string
		width  int
		expect string
	}{
		{
			name:   "empty lines",
			input:  "",
			pos:    0,
			defs:   [][2]string{},
			width:  80,
			expect: "",
		},
		{
			name:  "one def, too short to wrap",
			input: "",
			pos:   0,
			defs: [][2]string{
				{"def1", "this is the first definition"},
			},
			width:  80,
			expect: "  def1  - this is the first definition" + DefaultLineSeparator,
		},
		{
			name:  "one def, long enough to wrap once",
			input: "",
			pos:   0,
			defs: [][2]string{
				{"def1", "this is the first definition with a rather long outro which will cause it to become wrapped"},
			},
			width: 80,
			expect: "  def1  - this is the first definition with a rather long outro which will cause" + DefaultLineSeparator +
				"          it to become wrapped" + DefaultLineSeparator,
		},
		{
			name:  "one def, long enough to wrap multiple times",
			input: "",
			pos:   0,
			defs: [][2]string{
				{"pumpkin", "Thing that does not exist. If it did, no it didn't. " +
					"In fact, you are quite certain there never was, " +
					"never has been, nor will there ever be a pumpkin, " +
					"either here or anywhere else. Pumpkin? What pumpkin?",
				},
			},
			width: 80,
			expect: "  pumpkin  - Thing that does not exist. If it did, no it didn't. In fact, you" + DefaultLineSeparator +
				"             are quite certain there never was, never has been, nor will there" + DefaultLineSeparator +
				"             ever be a pumpkin, either here or anywhere else. Pumpkin? What" + DefaultLineSeparator +
				"             pumpkin?" + DefaultLineSeparator,
		},
		{
			name:  "two defs, no wrap",
			input: "",
			pos:   0,
			defs: [][2]string{
				{"John", "Has a passion for REALLY TERRIBLE MOVIES."},
				{"Rose", "Has a passion for RATHER OBSCURE LITERATURE."},
			},
			width: 80,
			expect: "  John  - Has a passion for REALLY TERRIBLE MOVIES." + DefaultParagraphSeparator +
				"  Rose  - Has a passion for RATHER OBSCURE LITERATURE." + DefaultLineSeparator,
		},
		{
			name:  "two defs, multi wrap",
			input: "",
			pos:   0,
			defs: [][2]string{
				{"John", "Has a passion for REALLY TERRIBLE MOVIES. " +
					"Likes to program computers but is NOT VERY GOOD AT IT. " +
					"Has a fondness for PARANORMAL LORE, and is an aspiring " +
					"AMATEUR MAGICIAN. Also likes to play GAMES sometimes.",
				},
				{"Rose", "Has a passion for RATHER OBSCURE LITERATURE. " +
					"Enjoys creative writing and is SOMEWHAT SECRETIVE ABOUT IT. " +
					"Has a fondness for the BEASTIALLY STRANGE AND FICTICIOUS, " +
					"and sometimes dabbles in PSYCHOANALYSIS. Also likes to KNIT, " +
					"and on occasion, if just the right one strikes your fancy, " +
					"you like to play VIDEO GAMES with your friends.",
				},
			},
			width: 80,
			expect: "  John  - Has a passion for REALLY TERRIBLE MOVIES. Likes to program computers" + DefaultLineSeparator +
				"          but is NOT VERY GOOD AT IT. Has a fondness for PARANORMAL LORE, and is" + DefaultLineSeparator +
				"          an aspiring AMATEUR MAGICIAN. Also likes to play GAMES sometimes." + DefaultParagraphSeparator +
				"  Rose  - Has a passion for RATHER OBSCURE LITERATURE. Enjoys creative writing" + DefaultLineSeparator +
				"          and is SOMEWHAT SECRETIVE ABOUT IT. Has a fondness for the BEASTIALLY" + DefaultLineSeparator +
				"          STRANGE AND FICTICIOUS, and sometimes dabbles in PSYCHOANALYSIS. Also" + DefaultLineSeparator +
				"          likes to KNIT, and on occasion, if just the right one strikes your" + DefaultLineSeparator +
				"          fancy, you like to play VIDEO GAMES with your friends." + DefaultLineSeparator,
		},
		{
			name:  "in middle of content",
			input: "PRE-EXISTING CONTENT",
			pos:   13,
			defs: [][2]string{
				{"def1", "this is the first definition"},
			},
			width: 80,
			expect: "PRE-EXISTING   def1  - this is the first definition" + DefaultLineSeparator +
				"CONTENT",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert := assert.New(t)
			actual := Edit(tc.input).InsertDefinitionsTable(tc.pos, tc.defs, tc.width).String()
			assert.Equal(tc.expect, actual)
		})
	}

}

func Test_InsertDefinitionsTableOpts(t *testing.T) {
	testCases := []struct {
		name    string
		input   string
		pos     int
		defs    [][2]string
		width   int
		options Options
		expect  string
	}{
		{
			name:    "for empty lines, noTrailing = true does not affect output",
			input:   "",
			pos:     0,
			defs:    [][2]string{},
			width:   80,
			options: Options{NoTrailingLineSeparators: true},
			expect:  "",
		},
		{
			name:  "for non-empty lines, noTrailing = true means no newline at end",
			input: "",
			pos:   0,
			defs: [][2]string{
				{"def1", "this is the first definition"},
			},
			width:   80,
			options: Options{NoTrailingLineSeparators: true},
			expect:  "  def1  - this is the first definition",
		},
		{
			name:  "custom line separator",
			input: "",
			pos:   0,
			defs: [][2]string{
				{"pumpkin", "Thing that does not exist. If it did, no it didn't. " +
					"In fact, you are quite certain there never was, " +
					"never has been, nor will there ever be a pumpkin, " +
					"either here or anywhere else. Pumpkin? What pumpkin?",
				},
			},
			width:   80,
			options: Options{LineSeparator: "<br/>"},
			expect: "  pumpkin  - Thing that does not exist. If it did, no it didn't. In fact, you<br/>" +
				"             are quite certain there never was, never has been, nor will there<br/>" +
				"             ever be a pumpkin, either here or anywhere else. Pumpkin? What<br/>" +
				"             pumpkin?<br/>",
		},
		{
			name:  "custom paragraph separator",
			input: "",
			pos:   0,
			defs: [][2]string{
				{"Dave", "Has a penchant for spinning out UNBELIEVABLY ILL JAMS with " +
					"his TURNTABLES AND MIXING GEAR. Likes to rave about BANDS " +
					"NO ONE'S EVER HEARD OF BUT HIM. Collects WEIRD DEAD THINGS " +
					"PRESERVED IN VARIOUS WAYS. Is an AMATEUR PHOTOGRAPHER and " +
					"operates own MAKESHIFT DARKROOM. Maintains a number of IRONICALLY " +
					"HUMOROUS BLOGS, WEBSITES, AND SOCIAL NETWORKING PROFILES. And " +
					"if the inspiration strikes, won't hesitate to drop some PHAT " +
					"RHYMES on a mofo and REPRESENT.",
				},
				{"Jade", "Has so many INTERESTS, she has trouble keeping track of them " +
					"all, even with an assortment of COLORFUL REMINDERS on her " +
					"fingers to help sort out everything on her mind. Nevertheless, " +
					"when she spends time in her GARDEN ATRIUM, the only thing on " +
					"her mind is her deep passion for HORTICULTURE.",
				},
			},
			width:   80,
			options: Options{ParagraphSeparator: "<P>"},
			expect: "  Dave  - Has a penchant for spinning out UNBELIEVABLY ILL JAMS with his" + DefaultLineSeparator +
				"          TURNTABLES AND MIXING GEAR. Likes to rave about BANDS NO ONE'S EVER" + DefaultLineSeparator +
				"          HEARD OF BUT HIM. Collects WEIRD DEAD THINGS PRESERVED IN VARIOUS" + DefaultLineSeparator +
				"          WAYS. Is an AMATEUR PHOTOGRAPHER and operates own MAKESHIFT DARKROOM." + DefaultLineSeparator +
				"          Maintains a number of IRONICALLY HUMOROUS BLOGS, WEBSITES, AND SOCIAL" + DefaultLineSeparator +
				"          NETWORKING PROFILES. And if the inspiration strikes, won't hesitate to" + DefaultLineSeparator +
				"          drop some PHAT RHYMES on a mofo and REPRESENT." +
				"<P>" +
				"  Jade  - Has so many INTERESTS, she has trouble keeping track of them all, even" + DefaultLineSeparator +
				"          with an assortment of COLORFUL REMINDERS on her fingers to help sort" + DefaultLineSeparator +
				"          out everything on her mind. Nevertheless, when she spends time in her" + DefaultLineSeparator +
				"          GARDEN ATRIUM, the only thing on her mind is her deep passion for" + DefaultLineSeparator +
				"          HORTICULTURE." + DefaultLineSeparator,
		},
		{
			name:  "custom line and paragraph separator",
			input: "",
			pos:   0,
			defs: [][2]string{
				{"def1", "The first definition which is long enough to span at least two lines and has the content for it."},
				{"def2", "A second definition to complement the first. This one takes multiple sentences to reach the end."},
				{"def3", "Third definition is the final one. It's a bit more terse than the other two. Slightly."},
			},
			width:   80,
			options: Options{ParagraphSeparator: "<P>", LineSeparator: "<br/>\n"},
			expect: "  def1  - The first definition which is long enough to span at least two lines<br/>\n" +
				"          and has the content for it." +
				"<P>" +
				"  def2  - A second definition to complement the first. This one takes multiple<br/>\n" +
				"          sentences to reach the end." +
				"<P>" +
				"  def3  - Third definition is the final one. It's a bit more terse than the<br/>\n" +
				"          other two. Slightly.<br/>\n",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert := assert.New(t)

			actualDirect := Edit(tc.input).InsertDefinitionsTableOpts(tc.pos, tc.defs, tc.width, tc.options).String()
			actualPreOpts := Edit(tc.input).WithOptions(tc.options).InsertDefinitionsTable(tc.pos, tc.defs, tc.width).String()

			assert.Equal(tc.expect, actualDirect, "InsertDefinitionsTableOpts(opts) check failed")
			assert.Equal(tc.expect, actualPreOpts, "WithOptions(opts).InsertDefinitionsTable() check failed")
		})
	}
}

func Test_Delete(t *testing.T) {
	testCases := []struct {
		name   string
		input  string
		start  int
		end    int
		expect string
	}{
		{
			name:   "delete nothing from empty",
			input:  "",
			start:  0,
			end:    0,
			expect: "",
		},
		{
			name:   "delete 1 char from empty",
			input:  "",
			start:  0,
			end:    1,
			expect: "",
		},
		{
			name:   "delete several chars from empty",
			input:  "",
			start:  0,
			end:    3,
			expect: "",
		},
		{
			name:   "delete past end of empty",
			input:  "",
			start:  1,
			end:    2,
			expect: "",
		},
		{
			name:   "delete invalid range from empty",
			input:  "",
			start:  20,
			end:    2,
			expect: "",
		},
		{
			name:   "delete invalid range from empty",
			input:  "",
			start:  20,
			end:    2,
			expect: "",
		},
		{
			name:   "delete nothing from start",
			input:  "rose",
			start:  0,
			end:    0,
			expect: "rose",
		},
		{
			name:   "delete nothing from middle",
			input:  "rose",
			start:  1,
			end:    1,
			expect: "rose",
		},
		{
			name:   "delete nothing from end",
			input:  "rose",
			start:  3,
			end:    3,
			expect: "rose",
		},
		{
			name:   "delete nothing from past end",
			input:  "rose",
			start:  4,
			end:    4,
			expect: "rose",
		},
		{
			name:   "delete 1 char at start",
			input:  "rose",
			start:  0,
			end:    1,
			expect: "ose",
		},
		{
			name:   "delete 1 char in middle",
			input:  "rose",
			start:  2,
			end:    3,
			expect: "roe",
		},
		{
			name:   "delete 1 char at end",
			input:  "rose",
			start:  3,
			end:    4,
			expect: "ros",
		},
		{
			name:   "delete 1 char past end (for none total)",
			input:  "rose",
			start:  4,
			end:    5,
			expect: "rose",
		},
		{
			name:   "delete multi chars at start",
			input:  "rose and jade",
			start:  0,
			end:    9,
			expect: "jade",
		},
		{
			name:   "delete multi chars in middle",
			input:  "rose and jade",
			start:  4,
			end:    9,
			expect: "rosejade",
		},
		{
			name:   "delete multi chars at end",
			input:  "rose and jade",
			start:  4,
			end:    13,
			expect: "rose",
		},
		{
			name:   "delete multi chars through end",
			input:  "rose and jade",
			start:  4,
			end:    200,
			expect: "rose",
		},
		{
			name:   "delete entire string",
			input:  "rose and jade",
			start:  0,
			end:    13,
			expect: "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert := assert.New(t)

			actual := Edit(tc.input).Delete(tc.start, tc.end).String()

			assert.Equal(tc.expect, actual)
		})
	}
}

func Test_Overtype(t *testing.T) {
	testCases := []struct {
		name   string
		input  string
		pos    int
		text   string
		expect string
	}{
		{
			name:   "add empty string to empty string",
			input:  "",
			pos:    0,
			text:   "",
			expect: "",
		},
		{
			name:   "add non-empty to empty",
			input:  "",
			pos:    0,
			text:   "Hello!",
			expect: "Hello!",
		},
		{
			name:   "add empty to non-empty start",
			input:  "test",
			pos:    0,
			text:   "",
			expect: "test",
		},
		{
			name:   "add empty to non-empty middle",
			input:  "test",
			pos:    1,
			text:   "",
			expect: "test",
		},
		{
			name:   "add empty to non-empty end",
			input:  "test",
			pos:    4,
			text:   "",
			expect: "test",
		},
		{
			name:   "add empty past non-empty end",
			input:  "test",
			pos:    80,
			text:   "",
			expect: "test",
		},
		{
			name:   "overtype at start",
			input:  "test-test-test",
			pos:    0,
			text:   "TEST",
			expect: "TEST-test-test",
		},
		{
			name:   "overtype in middle",
			input:  "test-test-test",
			pos:    5,
			text:   "TEST",
			expect: "test-TEST-test",
		},
		{
			name:   "overtype at end",
			input:  "test-test-test",
			pos:    10,
			text:   "TEST",
			expect: "test-test-TEST",
		},
		{
			name:   "overtype past end",
			input:  "test-test-test",
			pos:    14,
			text:   "TEST",
			expect: "test-test-testTEST",
		},
		{
			name:   "overtype through end, from middle",
			input:  "test-test-test",
			pos:    5,
			text:   "8888-8888-TEST",
			expect: "test-8888-8888-TEST",
		},
		{
			name:   "overtype through end, from start",
			input:  "test",
			pos:    0,
			text:   "TEST-AND-TEST",
			expect: "TEST-AND-TEST",
		},
		{
			name:   "overtype exactly from start to end",
			input:  "test",
			pos:    0,
			text:   "TEST",
			expect: "TEST",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert := assert.New(t)

			actual := Edit(tc.input).Overtype(tc.pos, tc.text).String()

			assert.Equal(tc.expect, actual)
		})
	}
}

func Test_Align(t *testing.T) {
	testCases := []struct {
		name   string
		input  string
		align  Alignment
		width  int
		expect string
	}{
		{
			name:   "left: empty string",
			input:  "",
			align:  Left,
			width:  10,
			expect: "",
		},
		{
			name:   "left: 1 word line",
			input:  "bluh",
			align:  Left,
			width:  10,
			expect: "bluh      ",
		},
		{
			name:   "left: multi word line",
			input:  "john egbert",
			align:  Left,
			width:  15,
			expect: "john egbert    ",
		},
		{
			name: "left: multiple lines, no trailing lineSep",
			input: "  john egbert" + DefaultLineSeparator +
				"        rose lalonde" + DefaultLineSeparator +
				"jade harley    " + DefaultLineSeparator +
				"     dave strider",
			align: Left,
			width: 15,
			expect: "john egbert    " + DefaultLineSeparator +
				"rose lalonde   " + DefaultLineSeparator +
				"jade harley    " + DefaultLineSeparator +
				"dave strider   ",
		},
		{
			name: "left: multiple lines, with trailing lineSep",
			input: "  john egbert" + DefaultLineSeparator +
				"        rose lalonde" + DefaultLineSeparator +
				"jade harley    " + DefaultLineSeparator +
				"     dave strider" + DefaultLineSeparator,
			align: Left,
			width: 15,
			expect: "john egbert    " + DefaultLineSeparator +
				"rose lalonde   " + DefaultLineSeparator +
				"jade harley    " + DefaultLineSeparator +
				"dave strider   " + DefaultLineSeparator,
		},
		{
			name:   "right: empty string",
			input:  "",
			align:  Right,
			width:  10,
			expect: "",
		},
		{
			name:   "right: 1 word line",
			input:  "bluh",
			align:  Right,
			width:  10,
			expect: "      bluh",
		},
		{
			name:   "right: multi word line",
			input:  "john egbert",
			align:  Right,
			width:  15,
			expect: "    john egbert",
		},
		{
			name: "right: multiple lines, no trailing lineSep",
			input: "  john egbert" + DefaultLineSeparator +
				"   rose lalonde  " + DefaultLineSeparator +
				"jade harley    " + DefaultLineSeparator +
				" dave strider ",
			align: Right,
			width: 15,
			expect: "    john egbert" + DefaultLineSeparator +
				"   rose lalonde" + DefaultLineSeparator +
				"    jade harley" + DefaultLineSeparator +
				"   dave strider",
		},
		{
			name: "right: multiple lines, with trailing lineSep",
			input: "  john egbert" + DefaultLineSeparator +
				"   rose lalonde  " + DefaultLineSeparator +
				"jade harley    " + DefaultLineSeparator +
				" dave strider " + DefaultLineSeparator,
			align: Right,
			width: 15,
			expect: "    john egbert" + DefaultLineSeparator +
				"   rose lalonde" + DefaultLineSeparator +
				"    jade harley" + DefaultLineSeparator +
				"   dave strider" + DefaultLineSeparator,
		},
		{
			name:   "center: empty string",
			input:  "",
			align:  Center,
			width:  10,
			expect: "",
		},
		{
			name:   "center: 1 word line",
			input:  "bluh",
			align:  Center,
			width:  10,
			expect: "   bluh   ",
		},
		{
			name:   "center: multi word line",
			input:  "john egbert",
			align:  Center,
			width:  15,
			expect: "  john egbert  ",
		},
		{
			name: "center: multiple lines, no trailing lineSep",
			input: "  john egbert" + DefaultLineSeparator +
				"   rose lalonde  " + DefaultLineSeparator +
				"jade harley    " + DefaultLineSeparator +
				" dave strider ",
			align: Center,
			width: 15,
			expect: "  john egbert  " + DefaultLineSeparator +
				"  rose lalonde " + DefaultLineSeparator +
				"  jade harley  " + DefaultLineSeparator +
				"  dave strider ",
		},
		{
			name: "center: multiple lines, with trailing lineSep",
			input: "  john egbert" + DefaultLineSeparator +
				"   rose lalonde  " + DefaultLineSeparator +
				"jade harley    " + DefaultLineSeparator +
				" dave strider " + DefaultLineSeparator,
			align: Center,
			width: 15,
			expect: "  john egbert  " + DefaultLineSeparator +
				"  rose lalonde " + DefaultLineSeparator +
				"  jade harley  " + DefaultLineSeparator +
				"  dave strider " + DefaultLineSeparator,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert := assert.New(t)
			actual := Edit(tc.input).Align(tc.align, tc.width).String()
			assert.Equal(tc.expect, actual)
		})
	}
}

func Test_AlignOpts(t *testing.T) {
	testCases := []struct {
		name    string
		input   string
		align   Alignment
		width   int
		options Options
		expect  string
	}{
		{
			name:  "left: empty line, noTrailing",
			input: "",
			align: Left,
			width: 10,
			options: Options{
				NoTrailingLineSeparators: true,
			},
			expect: "          ",
		},
		{
			name: "left: multi-line, noTrailing, with final lineSep",
			input: "  This quadrant presides over" + DefaultLineSeparator +
				" MOIRALLEGIENCE, the other" + DefaultLineSeparator +
				" conciliatory relationship. " + DefaultLineSeparator,
			align: Left,
			width: 32,
			options: Options{
				NoTrailingLineSeparators: true,
			},
			expect: "This quadrant presides over     " + DefaultLineSeparator +
				"MOIRALLEGIENCE, the other       " + DefaultLineSeparator +
				"conciliatory relationship.      " + DefaultLineSeparator +
				"                                ",
		},
		{
			name: "left: multi-line, noTrailing, without final lineSep",
			input: "  This quadrant presides over" + DefaultLineSeparator +
				" MOIRALLEGIENCE, the other" + DefaultLineSeparator +
				" conciliatory relationship. ",
			align: Left,
			width: 32,
			options: Options{
				NoTrailingLineSeparators: true,
			},
			expect: "This quadrant presides over     " + DefaultLineSeparator +
				"MOIRALLEGIENCE, the other       " + DefaultLineSeparator +
				"conciliatory relationship.      ",
		},
		{
			name: "left: multi-paragraph, preserved, default parasep",
			input: "Pale Quadrant:" + DefaultLineSeparator +
				"  This quadrant presides over" + DefaultLineSeparator +
				" MOIRALLEGIENCE, the other  " + DefaultLineSeparator +
				" conciliatory relationship. " +
				DefaultParagraphSeparator +
				"  Flush Quadrant: " + DefaultLineSeparator +
				"   When two individuals find" + DefaultLineSeparator +
				"themselves in the flushed  " + DefaultLineSeparator +
				"  quadrant together, they are" + DefaultLineSeparator +
				"said to be MATESPRITS.",
			align: Left,
			width: 32,
			options: Options{
				PreserveParagraphs: true,
			},
			expect: "Pale Quadrant:                  " + DefaultLineSeparator +
				"This quadrant presides over     " + DefaultLineSeparator +
				"MOIRALLEGIENCE, the other       " + DefaultLineSeparator +
				"conciliatory relationship.      " +
				DefaultParagraphSeparator +
				"Flush Quadrant:                 " + DefaultLineSeparator +
				"When two individuals find       " + DefaultLineSeparator +
				"themselves in the flushed       " + DefaultLineSeparator +
				"quandrant together, they are    " + DefaultLineSeparator +
				"said to be MATESPRITS.",
		},
		{
			name: "left: multi-paragraph, preserved, custom parasep",
			input: "Pale Quadrant:" + DefaultLineSeparator +
				"  This quadrant presides over" + DefaultLineSeparator +
				" MOIRALLEGIENCE, the other  " + DefaultLineSeparator +
				" conciliatory relationship. " +
				"<P>---<P>\n" +
				"  Flush Quadrant: " + DefaultLineSeparator +
				"   When two individuals find" + DefaultLineSeparator +
				"themselves in the flushed  " + DefaultLineSeparator +
				"  quadrant together, they are" + DefaultLineSeparator +
				"said to be MATESPRITS.",
			align: Left,
			width: 32,
			options: Options{
				PreserveParagraphs: true,
				ParagraphSeparator: "<P>-\n-<P>",
			},
			expect: "Pale Quadrant:                  " + DefaultLineSeparator +
				"This quadrant presides over     " + DefaultLineSeparator +
				"MOIRALLEGIENCE, the other       " + DefaultLineSeparator +
				"conciliatory relationship.      " +
				"<P>-\n-<P>" +
				"Flush Quadrant:                 " + DefaultLineSeparator +
				"When two individuals find       " + DefaultLineSeparator +
				"themselves in the flushed       " + DefaultLineSeparator +
				"quandrant together, they are    " + DefaultLineSeparator +
				"said to be MATESPRITS.",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert := assert.New(t)

			actualDirect := Edit(tc.input).AlignOpts(tc.align, tc.width, tc.options).String()
			actualPreOpts := Edit(tc.input).WithOptions(tc.options).Align(tc.align, tc.width).String()

			assert.Equal(tc.expect, actualDirect, "AlignOpts(opts) check failed")
			assert.Equal(tc.expect, actualPreOpts, "WithOptions(opts).Align() check failed")

		})
	}
}
