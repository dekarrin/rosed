package rosed

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Editor_IsSubEditor(t *testing.T) {
	testCases := []struct {
		name   string
		ed     Editor
		expect bool
	}{
		{
			name: "not a subeditor",
			ed: Editor{
				ref: nil,
			},
			expect: false,
		},
		{
			name: "is a subeditor",
			ed: Editor{
				ref: &parentRef{
					parent: &Editor{},
				},
			},
			expect: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert := assert.New(t)

			actual := tc.ed.IsSubEditor()

			assert.Equal(tc.expect, actual)
		})
	}
}

func Test_Editor_Commit(t *testing.T) {
	testCases := []struct {
		name   string
		ed     Editor
		expect Editor
	}{
		{
			name: "empty string",
			ed: Editor{
				Text: "",
				ref: &parentRef{
					start: 0,
					end:   0,
					parent: &Editor{
						Text: "",
					},
				},
			},
			expect: Editor{
				Text: "",
			},
		},
		{
			name: "non sub-editor returns self",
			ed: Editor{
				Text: "test",
				ref:  nil,
			},
			expect: Editor{
				Text: "test",
			},
		},
		{
			name: "replace middle",
			ed: Editor{
				Text: "test",
				ref: &parentRef{
					start: 2,
					end:   8,
					parent: &Editor{
						Text: "a useful execution",
					},
				},
			},
			expect: Editor{
				Text: "a test execution",
			},
		},
		{
			name: "insert middle",
			ed: Editor{
				Text: " test",
				ref: &parentRef{
					start: 3,
					end:   3,
					parent: &Editor{
						Text: "one execution",
					},
				},
			},
			expect: Editor{
				Text: "one test execution",
			},
		},
		{
			name: "replace end",
			ed: Editor{
				Text: "test case",
				ref: &parentRef{
					start: 4,
					end:   13,
					parent: &Editor{
						Text: "one execution",
					},
				},
			},
			expect: Editor{
				Text: "one test case",
			},
		},
		{
			name: "insert end",
			ed: Editor{
				Text: " case",
				ref: &parentRef{
					start: 8,
					end:   8,
					parent: &Editor{
						Text: "one test",
					},
				},
			},
			expect: Editor{
				Text: "one test case",
			},
		},
		{
			name: "replace start",
			ed: Editor{
				Text: "a single",
				ref: &parentRef{
					start: 0,
					end:   3,
					parent: &Editor{
						Text: "one test",
					},
				},
			},
			expect: Editor{
				Text: "a single test",
			},
		},
		{
			name: "insert start",
			ed: Editor{
				Text: "one test ",
				ref: &parentRef{
					start: 0,
					end:   0,
					parent: &Editor{
						Text: "case",
					},
				},
			},
			expect: Editor{
				Text: "one test case",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert := assert.New(t)
			actual := tc.ed.Commit()

			// do a full equal assertion because we care about all fields of
			// the returned editor.

			assert.Equal(tc.expect, actual)
		})
	}
}

func Test_Editor_CommitAll(t *testing.T) {
	testCases := []struct {
		name   string
		ed     Editor
		expect Editor
	}{
		{
			name: "non sub-editor returns self",
			ed: Editor{
				Text: "test",
				ref:  nil,
			},
			expect: Editor{
				Text: "test",
			},
		},
		{
			name: "1 sub-Editor chain",
			ed: Editor{
				Text: " full",
				ref: &parentRef{
					start: 1,
					end:   1,
					parent: &Editor{
						Text: "a string",
					},
				},
			},
			expect: Editor{
				Text: "a full string",
			},
		},
		{
			name: "2 sub-Editor chain",
			ed: Editor{
				Text: "and grand",
				ref: &parentRef{
					start: 7,
					end:   7,
					parent: &Editor{
						Text: "middle child",
						ref: &parentRef{
							start: 0,
							end:   6,
							parent: &Editor{
								Text: "parent text",
							},
						},
					},
				},
			},
			expect: Editor{
				Text: "middle and grandchild text",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert := assert.New(t)
			actual := tc.ed.CommitAll()

			// do a full equal assertion because we care about all fields of
			// the returned editor.

			assert.Equal(tc.expect, actual)
		})
	}
}

func Test_Editor_Chars(t *testing.T) {
	testCases := []struct {
		name   string
		ed     Editor
		start  int
		end    int
		expect Editor
	}{
		{
			name: "typical subsection",
			ed: Editor{
				Text: "some testing text to edit",
			},
			start: 5,
			end:   17,
			expect: Editor{
				Text: "testing text",
			},
		},
		{
			name: "end > end of the string",
			ed: Editor{
				Text: "test",
			},
			start: 1,
			end:   20,
			expect: Editor{
				Text: "est",
			},
		},
		{
			name: "start < 0",
			ed: Editor{
				Text: "test",
			},
			start: -3,
			end:   2,
			expect: Editor{
				Text: "e",
			},
		},
		{
			name: "start > end",
			ed: Editor{
				Text: "test",
			},
			start: 2,
			end:   1,
			expect: Editor{
				Text: "",
			},
		},
		{
			name: "start = end",
			ed: Editor{
				Text: "test",
			},
			start: 2,
			end:   2,
			expect: Editor{
				Text: "",
			},
		},
		{
			name: "both < 0",
			ed: Editor{
				Text: "test",
			},
			start: -4,
			end:   -2,
			expect: Editor{
				Text: "te",
			},
		},
		{
			name: "get before decomposed grapheme",
			ed: Editor{
				Text: "Franc\u0327ais",
			},
			start: 1,
			end:   4,
			expect: Editor{
				Text: "ran",
			},
		},
		{
			name: "get after decomposed grapheme",
			ed: Editor{
				Text: "Franc\u0327ais",
			},
			start: 5,
			end:   7,
			expect: Editor{
				Text: "ai",
			},
		},
		{
			name: "get across decomposed grapheme",
			ed: Editor{
				Text: "Franc\u0327ais",
			},
			start: 3,
			end:   7,
			expect: Editor{
				Text: "nc\u0327ai",
			},
		},
		{
			name: "preserve options",
			ed: Editor{
				Text: "test",
				Options: Options{
					LineSeparator: DefaultLineSeparator + DefaultLineSeparator,
				},
			},
			start: 0,
			end:   4,
			expect: Editor{
				Text: "test",
				Options: Options{
					LineSeparator: DefaultLineSeparator + DefaultLineSeparator,
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert := assert.New(t)
			actual := tc.ed.Chars(tc.start, tc.end)

			// don't do a full Equal as that will compare unexported
			// fields; instead just check the ones we care about

			assert.Equal(tc.expect.Options, actual.Options)
			assert.Equal(tc.expect.Text, actual.Text)
		})
	}
}

func Test_Editor_CharsFrom(t *testing.T) {
	testCases := []struct {
		name   string
		ed     Editor
		start  int
		expect Editor
	}{
		{
			name: "from middle",
			ed: Editor{
				Text: "some testing text to edit",
			},
			start: 5,
			expect: Editor{
				Text: "testing text to edit",
			},
		},
		{
			name: "entire string",
			ed: Editor{
				Text: "test",
			},
			start: 0,
			expect: Editor{
				Text: "test",
			},
		},
		{
			name: "empty string",
			ed: Editor{
				Text: "",
			},
			start: 0,
			expect: Editor{
				Text: "",
			},
		},
		{
			name: "start < 0",
			ed: Editor{
				Text: "test",
			},
			start: -3,
			expect: Editor{
				Text: "est",
			},
		},
		{
			name: "start is at end",
			ed: Editor{
				Text: "test",
			},
			start: 4,
			expect: Editor{
				Text: "",
			},
		},
		{
			name: "get after decomposed grapheme",
			ed: Editor{
				Text: "Franc\u0327ais",
			},
			start: 5,
			expect: Editor{
				Text: "ais",
			},
		},
		{
			name: "get across decomposed grapheme",
			ed: Editor{
				Text: "Franc\u0327ais",
			},
			start: 3,
			expect: Editor{
				Text: "nc\u0327ais",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert := assert.New(t)
			actual := tc.ed.CharsFrom(tc.start)

			// don't do a full Equal as that will compare unexported
			// fields; instead just check the ones we care about

			assert.Equal(tc.expect.Options, actual.Options)
			assert.Equal(tc.expect.Text, actual.Text)
		})
	}
}

func Test_Editor_CharsTo(t *testing.T) {
	testCases := []struct {
		name   string
		ed     Editor
		end    int
		expect Editor
	}{
		{
			name: "to middle",
			ed: Editor{
				Text: "some testing text to edit",
			},
			end: 17,
			expect: Editor{
				Text: "some testing text",
			},
		},
		{
			name: "entire string",
			ed: Editor{
				Text: "test",
			},
			end: 4,
			expect: Editor{
				Text: "test",
			},
		},
		{
			name: "to zero of non-empty",
			ed: Editor{
				Text: "toZero",
			},
			end: 0,
			expect: Editor{
				Text: "",
			},
		},
		{
			name: "empty string",
			ed: Editor{
				Text: "",
			},
			end: 0,
			expect: Editor{
				Text: "",
			},
		},
		{
			name: "end > end of the string",
			ed: Editor{
				Text: "test",
			},
			end: 20,
			expect: Editor{
				Text: "test",
			},
		},
		{
			name: "end < 0",
			ed: Editor{
				Text: "test",
			},
			end: -2,
			expect: Editor{
				Text: "te",
			},
		},
		{
			name: "get before decomposed grapheme",
			ed: Editor{
				Text: "Franc\u0327ais",
			},
			end: 3,
			expect: Editor{
				Text: "Fra",
			},
		},
		{
			name: "get across decomposed grapheme",
			ed: Editor{
				Text: "Franc\u0327ais",
			},
			end: 6,
			expect: Editor{
				Text: "Franc\u0327a",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert := assert.New(t)
			actual := tc.ed.CharsTo(tc.end)

			// don't do a full Equal as that will compare unexported
			// fields; instead just check the ones we care about

			assert.Equal(tc.expect.Options, actual.Options)
			assert.Equal(tc.expect.Text, actual.Text)
		})
	}
}

func Test_Editor_Lines(t *testing.T) {
	testCases := []struct {
		name   string
		ed     Editor
		start  int
		end    int
		expect Editor
	}{
		{
			name: "empty string",
			ed: Editor{
				Text: "",
			},
			start: 0,
			end:   0,
			expect: Editor{
				Text: "",
			},
		},
		{
			name: "in middle",
			ed: Editor{
				Text: "line0" + DefaultLineSeparator +
					"line1" + DefaultLineSeparator +
					"line2" + DefaultLineSeparator +
					"line3" + DefaultLineSeparator +
					"line4" + DefaultLineSeparator,
			},
			start: 1,
			end:   3,
			expect: Editor{
				Text: "line1" + DefaultLineSeparator +
					"line2" + DefaultLineSeparator,
			},
		},
		{
			name: "entire string",
			ed: Editor{
				Text: "line0" + DefaultLineSeparator +
					"line1" + DefaultLineSeparator +
					"line2" + DefaultLineSeparator +
					"line3" + DefaultLineSeparator +
					"line4" + DefaultLineSeparator,
			},
			start: 0,
			end:   5,
			expect: Editor{
				Text: "line0" + DefaultLineSeparator +
					"line1" + DefaultLineSeparator +
					"line2" + DefaultLineSeparator +
					"line3" + DefaultLineSeparator +
					"line4" + DefaultLineSeparator,
			},
		},
		{
			name: "end > end of the string",
			ed: Editor{
				Text: "line0" + DefaultLineSeparator +
					"line1" + DefaultLineSeparator +
					"line2" + DefaultLineSeparator,
			},
			start: 1,
			end:   20,
			expect: Editor{
				Text: "line1" + DefaultLineSeparator +
					"line2" + DefaultLineSeparator,
			},
		},
		{
			name: "start < 0",
			ed: Editor{
				Text: "line0" + DefaultLineSeparator +
					"line1" + DefaultLineSeparator +
					"line2" + DefaultLineSeparator +
					"line3" + DefaultLineSeparator +
					"line4" + DefaultLineSeparator,
			},
			start: -3,
			end:   4,
			expect: Editor{
				Text: "line2" + DefaultLineSeparator +
					"line3" + DefaultLineSeparator,
			},
		},
		{
			name: "start > end",
			ed: Editor{
				Text: "line0" + DefaultLineSeparator +
					"line1" + DefaultLineSeparator +
					"line2" + DefaultLineSeparator +
					"line3" + DefaultLineSeparator +
					"line4" + DefaultLineSeparator,
			},
			start: 2,
			end:   1,
			expect: Editor{
				Text: "",
			},
		},
		{
			name: "start = end",
			ed: Editor{
				Text: "line0" + DefaultLineSeparator +
					"line1" + DefaultLineSeparator +
					"line2" + DefaultLineSeparator +
					"line3" + DefaultLineSeparator +
					"line4" + DefaultLineSeparator,
			},
			start: 2,
			end:   2,
			expect: Editor{
				Text: "",
			},
		},
		{
			name: "both < 0",
			ed: Editor{
				Text: "line0" + DefaultLineSeparator +
					"line1" + DefaultLineSeparator +
					"line2" + DefaultLineSeparator +
					"line3" + DefaultLineSeparator +
					"line4" + DefaultLineSeparator,
			},
			start: -4,
			end:   -2,
			expect: Editor{
				Text: "line1" + DefaultLineSeparator +
					"line2" + DefaultLineSeparator,
			},
		},
		{
			name: "preserve options",
			ed: Editor{
				Text: "line0" + DefaultLineSeparator,
				Options: Options{
					IndentStr: DefaultIndentString,
				},
			},
			start: 0,
			end:   1,
			expect: Editor{
				Text: "line0" + DefaultLineSeparator,
				Options: Options{
					IndentStr: DefaultIndentString,
				},
			},
		},
		{
			name: "no trailing line sep, default options",
			ed: Editor{
				Text: "line0" + DefaultLineSeparator +
					"line1",
			},
			start: 0,
			end:   2,
			expect: Editor{
				Text: "line0" + DefaultLineSeparator +
					"line1",
			},
		},
		{
			name: "no trailing line sep, specified in options",
			ed: Editor{
				Text: "line0" + DefaultLineSeparator +
					"line1",
				Options: Options{
					NoTrailingLineSeparators: true,
				},
			},
			start: 0,
			end:   2,
			expect: Editor{
				Text: "line0" + DefaultLineSeparator +
					"line1",
				Options: Options{
					NoTrailingLineSeparators: true,
				},
			},
		},
		{
			name: "non-default line separator",
			ed: Editor{
				Text: "line0<P>" +
					"line1<P>" +
					"line2<P>",
				Options: Options{
					LineSeparator: "<P>",
				},
			},
			start: 0,
			end:   1,
			expect: Editor{
				Text: "line0<P>",
				Options: Options{
					LineSeparator: "<P>",
				},
			},
		},
		{
			name: "decomposed grapheme line sep",
			ed: Editor{
				Text: "line0c\u0327" +
					"line1c\u0327" +
					"line2c\u0327",
				Options: Options{
					LineSeparator: "c\u0327",
				},
			},
			start: 0,
			end:   2,
			expect: Editor{
				Text: "line0c\u0327" +
					"line1c\u0327",
				Options: Options{
					LineSeparator: "c\u0327",
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert := assert.New(t)
			actual := tc.ed.Lines(tc.start, tc.end)

			// don't do a full Equal as that will compare unexported
			// fields; instead just check the ones we care about

			assert.Equal(tc.expect.Options, actual.Options)
			assert.Equal(tc.expect.Text, actual.Text)
		})
	}
}

func Test_Editor_LinesFrom(t *testing.T) {
	testCases := []struct {
		name   string
		ed     Editor
		start  int
		expect Editor
	}{
		{
			name: "empty string",
			ed: Editor{
				Text: "",
			},
			start: 0,
			expect: Editor{
				Text: "",
			},
		},
		{
			name: "in middle",
			ed: Editor{
				Text: "line0" + DefaultLineSeparator +
					"line1" + DefaultLineSeparator +
					"line2" + DefaultLineSeparator +
					"line3" + DefaultLineSeparator +
					"line4" + DefaultLineSeparator,
			},
			start: 1,
			expect: Editor{
				Text: "line1" + DefaultLineSeparator +
					"line2" + DefaultLineSeparator +
					"line3" + DefaultLineSeparator +
					"line4" + DefaultLineSeparator,
			},
		},
		{
			name: "entire string",
			ed: Editor{
				Text: "line0" + DefaultLineSeparator +
					"line1" + DefaultLineSeparator +
					"line2" + DefaultLineSeparator +
					"line3" + DefaultLineSeparator +
					"line4" + DefaultLineSeparator,
			},
			start: 0,
			expect: Editor{
				Text: "line0" + DefaultLineSeparator +
					"line1" + DefaultLineSeparator +
					"line2" + DefaultLineSeparator +
					"line3" + DefaultLineSeparator +
					"line4" + DefaultLineSeparator,
			},
		},
		{
			name: "start < 0",
			ed: Editor{
				Text: "line0" + DefaultLineSeparator +
					"line1" + DefaultLineSeparator +
					"line2" + DefaultLineSeparator +
					"line3" + DefaultLineSeparator +
					"line4" + DefaultLineSeparator,
			},
			start: -3,
			expect: Editor{
				Text: "line2" + DefaultLineSeparator +
					"line3" + DefaultLineSeparator +
					"line4" + DefaultLineSeparator,
			},
		},
		{
			name: "start past end of string",
			ed: Editor{
				Text: "line0" + DefaultLineSeparator +
					"line1" + DefaultLineSeparator +
					"line2" + DefaultLineSeparator +
					"line3" + DefaultLineSeparator +
					"line4" + DefaultLineSeparator,
			},
			start: 20,
			expect: Editor{
				Text: "",
			},
		},
		{
			name: "start at end of string",
			ed: Editor{
				Text: "line0" + DefaultLineSeparator +
					"line1" + DefaultLineSeparator +
					"line2" + DefaultLineSeparator +
					"line3" + DefaultLineSeparator +
					"line4" + DefaultLineSeparator,
			},
			start: 5,
			expect: Editor{
				Text: "",
			},
		},
		{
			name: "preserve options",
			ed: Editor{
				Text: "line0" + DefaultLineSeparator,
				Options: Options{
					IndentStr: DefaultIndentString,
				},
			},
			start: 0,
			expect: Editor{
				Text: "line0" + DefaultLineSeparator,
				Options: Options{
					IndentStr: DefaultIndentString,
				},
			},
		},
		{
			name: "no trailing line sep, default options",
			ed: Editor{
				Text: "line0" + DefaultLineSeparator +
					"line1",
			},
			start: 0,
			expect: Editor{
				Text: "line0" + DefaultLineSeparator +
					"line1",
			},
		},
		{
			name: "no trailing line sep, specified in options",
			ed: Editor{
				Text: "line0" + DefaultLineSeparator +
					"line1",
				Options: Options{
					NoTrailingLineSeparators: true,
				},
			},
			start: 0,
			expect: Editor{
				Text: "line0" + DefaultLineSeparator +
					"line1",
				Options: Options{
					NoTrailingLineSeparators: true,
				},
			},
		},
		{
			name: "non-default line separator",
			ed: Editor{
				Text: "line0<P>" +
					"line1<P>" +
					"line2<P>",
				Options: Options{
					LineSeparator: "<P>",
				},
			},
			start: 1,
			expect: Editor{
				Text: "line1<P>" +
					"line2<P>",
				Options: Options{
					LineSeparator: "<P>",
				},
			},
		},
		{
			name: "decomposed grapheme line sep",
			ed: Editor{
				Text: "line0c\u0327" +
					"line1c\u0327" +
					"line2c\u0327",
				Options: Options{
					LineSeparator: "c\u0327",
				},
			},
			start: 1,
			expect: Editor{
				Text: "line1c\u0327" +
					"line2c\u0327",
				Options: Options{
					LineSeparator: "c\u0327",
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert := assert.New(t)
			actual := tc.ed.LinesFrom(tc.start)

			// don't do a full Equal as that will compare unexported
			// fields; instead just check the ones we care about

			assert.Equal(tc.expect.Options, actual.Options)
			assert.Equal(tc.expect.Text, actual.Text)
		})
	}
}

func Test_Editor_LinesTo(t *testing.T) {
	testCases := []struct {
		name   string
		ed     Editor
		end    int
		expect Editor
	}{
		{
			name: "empty string",
			ed: Editor{
				Text: "",
			},
			end: 0,
			expect: Editor{
				Text: "",
			},
		},
		{
			name: "in middle",
			ed: Editor{
				Text: "line0" + DefaultLineSeparator +
					"line1" + DefaultLineSeparator +
					"line2" + DefaultLineSeparator +
					"line3" + DefaultLineSeparator +
					"line4" + DefaultLineSeparator,
			},
			end: 3,
			expect: Editor{
				Text: "line0" + DefaultLineSeparator +
					"line1" + DefaultLineSeparator +
					"line2" + DefaultLineSeparator,
			},
		},
		{
			name: "entire string",
			ed: Editor{
				Text: "line0" + DefaultLineSeparator +
					"line1" + DefaultLineSeparator +
					"line2" + DefaultLineSeparator +
					"line3" + DefaultLineSeparator +
					"line4" + DefaultLineSeparator,
			},
			end: 5,
			expect: Editor{
				Text: "line0" + DefaultLineSeparator +
					"line1" + DefaultLineSeparator +
					"line2" + DefaultLineSeparator +
					"line3" + DefaultLineSeparator +
					"line4" + DefaultLineSeparator,
			},
		},
		{
			name: "end > end of the string",
			ed: Editor{
				Text: "line0" + DefaultLineSeparator +
					"line1" + DefaultLineSeparator +
					"line2" + DefaultLineSeparator,
			},
			end: 20,
			expect: Editor{
				Text: "line0" + DefaultLineSeparator +
					"line1" + DefaultLineSeparator +
					"line2" + DefaultLineSeparator,
			},
		},
		{
			name: "end < 0",
			ed: Editor{
				Text: "line0" + DefaultLineSeparator +
					"line1" + DefaultLineSeparator +
					"line2" + DefaultLineSeparator +
					"line3" + DefaultLineSeparator +
					"line4" + DefaultLineSeparator,
			},
			end: -2,
			expect: Editor{
				Text: "line0" + DefaultLineSeparator +
					"line1" + DefaultLineSeparator +
					"line2" + DefaultLineSeparator,
			},
		},
		{
			name: "preserve options",
			ed: Editor{
				Text: "line0" + DefaultLineSeparator,
				Options: Options{
					IndentStr: DefaultIndentString,
				},
			},
			end: 1,
			expect: Editor{
				Text: "line0" + DefaultLineSeparator,
				Options: Options{
					IndentStr: DefaultIndentString,
				},
			},
		},
		{
			name: "no trailing line sep, default options",
			ed: Editor{
				Text: "line0" + DefaultLineSeparator +
					"line1",
			},
			end: 2,
			expect: Editor{
				Text: "line0" + DefaultLineSeparator +
					"line1",
			},
		},
		{
			name: "no trailing line sep, specified in options",
			ed: Editor{
				Text: "line0" + DefaultLineSeparator +
					"line1",
				Options: Options{
					NoTrailingLineSeparators: true,
				},
			},
			end: 2,
			expect: Editor{
				Text: "line0" + DefaultLineSeparator +
					"line1",
				Options: Options{
					NoTrailingLineSeparators: true,
				},
			},
		},
		{
			name: "non-default line separator",
			ed: Editor{
				Text: "line0<P>" +
					"line1<P>" +
					"line2<P>",
				Options: Options{
					LineSeparator: "<P>",
				},
			},
			end: 2,
			expect: Editor{
				Text: "line0<P>" +
					"line1<P>",
				Options: Options{
					LineSeparator: "<P>",
				},
			},
		},
		{
			name: "decomposed grapheme line sep",
			ed: Editor{
				Text: "line0c\u0327" +
					"line1c\u0327" +
					"line2c\u0327",
				Options: Options{
					LineSeparator: "c\u0327",
				},
			},
			end: 1,
			expect: Editor{
				Text: "line0c\u0327",
				Options: Options{
					LineSeparator: "c\u0327",
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert := assert.New(t)
			actual := tc.ed.LinesTo(tc.end)

			// don't do a full Equal as that will compare unexported
			// fields; instead just check the ones we care about

			assert.Equal(tc.expect.Options, actual.Options)
			assert.Equal(tc.expect.Text, actual.Text)
		})
	}
}
