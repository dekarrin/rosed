package gem

import (
	"testing"
	"unicode"

	"github.com/stretchr/testify/assert"
)

func intSliceRef(to ...int) *[]int {
	return &to
}

func Test_New(t *testing.T) {
	testCases := []struct {
		name   string
		input  string
		expect String
	}{
		{"empty string", "", String{r: []rune{}}},
		{"one-char string", "1", String{r: []rune{'1'}}},
		{"two-char string", "12", String{r: []rune{'1', '2'}}},
		{"three-char string", "123", String{r: []rune{'1', '2', '3'}}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert := assert.New(t)
			actual := New(tc.input)
			assert.True(tc.expect.Equal(actual))
		})
	}
}

func Test_String_Equal(t *testing.T) {
	emptyStr := New("")
	testStr := New("test")

	testCases := []struct {
		name    string
		input   String
		compare interface{}
		expect  bool
	}{
		{"Zero == Zero", Zero, Zero, true},
		{"Zero != &Zero", Zero, &Zero, false},
		{"Zero != nil", Zero, nil, false},
		{"Zero == New empty", Zero, emptyStr, true},
		{"Zero != &New empty", Zero, &emptyStr, false},
		{"New empty == Zero", emptyStr, Zero, true},
		{"New empty != &Zero", emptyStr, &Zero, false},
		{"New empty == New empty", emptyStr, emptyStr, true},
		{"New empty != &New empty", emptyStr, &emptyStr, false},
		{"filled == filled", testStr, testStr, true},
		{"filled != &filled", testStr, &testStr, false},
		{"filled != Zero", testStr, Zero, false},
		{"filled != &Zero", testStr, &Zero, false},
		{"Zero != filled", Zero, testStr, false},
		{"Zero != &filled", Zero, &testStr, false},
		{"filled != New empty", testStr, emptyStr, false},
		{"filled != &New empty", testStr, &emptyStr, false},
		{"New empty != filled", emptyStr, testStr, false},
		{"New empty != &filled", emptyStr, &testStr, false},
		{"filled != nil", testStr, nil, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert := assert.New(t)
			actual := tc.input.Equal(tc.compare)
			assert.Equal(tc.expect, actual)
		})
	}
}

func Test_String_Equal_Ptr(t *testing.T) {
	emptyStr := New("")
	testStr := New("test")

	testCases := []struct {
		name    string
		input   *String
		compare interface{}
		expect  bool
	}{
		{"&Zero == Zero", &Zero, Zero, true},
		{"&Zero != &Zero", &Zero, &Zero, false},
		{"&Zero == nil", &Zero, nil, false},
		{"nil == Zero", nil, Zero, false},
		{"&Zero == New empty", &Zero, emptyStr, true},
		{"&Zero != &New empty", &Zero, &emptyStr, false},
		{"&New empty == Zero", &emptyStr, Zero, true},
		{"&New empty != &Zero", &emptyStr, &Zero, false},
		{"&New empty == New empty", &emptyStr, emptyStr, true},
		{"&New empty != &New empty", &emptyStr, &emptyStr, false},
		{"&filled == filled", &testStr, testStr, true},
		{"&filled != &filled", &testStr, &testStr, false},
		{"&filled != Zero", &testStr, Zero, false},
		{"&filled != &Zero", &testStr, &Zero, false},
		{"&Zero != filled", &Zero, testStr, false},
		{"&Zero != &filled", &Zero, &testStr, false},
		{"&filled != New empty", &testStr, emptyStr, false},
		{"&filled != &New empty", &testStr, &emptyStr, false},
		{"&New empty != filled", &emptyStr, testStr, false},
		{"&New empty != &filled", &emptyStr, &testStr, false},
		{"&filled != nil", &testStr, nil, false},
		{"nil != filled", nil, testStr, false},
		{"nil != &filled", nil, &testStr, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert := assert.New(t)
			var actual bool
			if tc.input == nil {
				actual = tc.compare == nil
			} else {
				actual = tc.input.Equal(tc.compare)
			}
			assert.Equal(tc.expect, actual)
		})
	}
}

func Test_String_Len(t *testing.T) {
	testCases := []struct {
		name   string
		input  String
		expect int
	}{
		{"empty string", Zero, 0},
		{"one-char string", New("1"), 1},
		{"two-char string", New("12"), 2},
		{"three-char string", New("123"), 3},
		{"c-with-cedilla is 1 char", New("Ç"), 1},
		{"c followed by combining-char-cedilla is 1 char (UAX29 analysis)", New("C\u0327"), 1},
		{"zwj emoji sequence with several modifiers is 1 char (UAX29 GB11)", New("\U0001F926\U0001F3FC\u200D\u2642\uFE0F"), 1},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert := assert.New(t)

			actual := tc.input.Len()

			assert.Equal(tc.expect, actual)
		})
	}
}

func Test_String_Add(t *testing.T) {
	testCases := []struct {
		name   string
		input1 String
		input2 String
		expect String
	}{
		{"empty string + empty string", Zero, Zero, Zero},
		{"empty string + non-empty string", Zero, New("test"), New("test")},
		{"non-empty string + empty string", New("test"), Zero, New("test")},
		{"2 non-empty strings", New("test1 + "), New("test2"), New("test1 + test2")},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert := assert.New(t)

			actual := tc.input1.Add(tc.input2)

			assert.True(tc.expect.Equal(actual))
		})
	}
}

func Test_String_CharAt(t *testing.T) {
	testCases := []struct {
		name        string
		str         String
		input       int
		expect      []rune
		expectPanic bool
	}{
		{"empty string panics at idx 0", Zero, 0, nil, true},
		{"empty string panics at idx 1", Zero, 1, nil, true},
		{"empty string panics at idx -1", Zero, -1, nil, true},
		{"empty string panics at high idx", Zero, 382834, nil, true},
		{"empty string panics at low idx", Zero, -382834, nil, true},
		{"1-char string at idx 0", New("1"), 0, []rune{'1'}, false},
		{"1-char string panics at idx 1", New("1"), 1, nil, true},
		{"1-char string panics at idx -1", New("1"), -1, nil, true},
		{"multichar string", New("test"), 2, []rune{'s'}, false},
		{"multichar string with combining mark, after mark", New("test C\u0327 test"), 7, []rune{'t'}, false},
		{"multichar string with combining mark, before mark", New("test C\u0327 test"), 1, []rune{'e'}, false},
		{"multichar string with combining mark, on mark", New("test C\u0327 test"), 5, []rune{'C', '\u0327'}, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert := assert.New(t)

			if tc.expectPanic {
				assert.Panics(func() {
					tc.str.CharAt(tc.input)
				})
			} else {
				actual := tc.str.CharAt(tc.input)
				assert.Equal(tc.expect, actual)
			}
		})
	}
}

func Test_String_SetCharAt(t *testing.T) {
	testCases := []struct {
		name        string
		str         String
		idx         int
		setTo       []rune
		expect      String
		expectPanic bool
	}{
		{"empty string panics at idx 0", Zero, 0, []rune{'1'}, Zero, true},
		{"empty string panics at idx 1", Zero, 1, []rune{'1'}, Zero, true},
		{"empty string panics at idx -1", Zero, -1, []rune{'1'}, Zero, true},
		{"empty string panics at high idx", Zero, 382834, []rune{'1'}, Zero, true},
		{"empty string panics at low idx", Zero, -382834, []rune{'1'}, Zero, true},
		{"1-char string at idx 0", New("1"), 0, []rune{'8'}, New("8"), false},
		{"1-char string panics at idx 1", New("1"), 1, []rune{'8'}, Zero, true},
		{"1-char string panics at idx -1", New("1"), -1, []rune{'8'}, Zero, true},
		{"multichar string", New("test"), 2, []rune{'8'}, New("te8t"), false},
		{"multichar string with combining mark, after mark", New("test C\u0327 test"), 7, []rune{'8'}, New("test C\u0327 8est"), false},
		{"multichar string with combining mark, before mark", New("test C\u0327 test"), 1, []rune{'8'}, New("t8st C\u0327 test"), false},
		{"multichar string with combining mark, on mark", New("test C\u0327 test"), 5, []rune{'8'}, New("test 8 test"), false},
		{"add combining mark to multichar string", New("test"), 2, []rune{'C', '\u0327'}, New("teC\u0327t"), false},
		{"empty replacement runes panics", New("test"), 2, []rune{}, Zero, true},
		{"nil replacement runes panics", New("test"), 2, nil, Zero, true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert := assert.New(t)

			if tc.expectPanic {
				assert.Panics(func() {
					tc.str.SetCharAt(tc.idx, tc.setTo)
				})
			} else {
				actual := tc.str.SetCharAt(tc.idx, tc.setTo)
				isEqual := actual.Equal(tc.expect)
				assert.Equal(true, isEqual)
			}
		})
	}
}

func Test_String_Reverse(t *testing.T) {
	testCases := []struct {
		name   string
		str    String
		expect String
	}{
		{"empty string", Zero, Zero},
		{"single-rune string", New("1"), String{
			r:  []rune{'1'},
			gc: intSliceRef(1),
		}},
		{"single-grapheme, multi-rune string", New("C\u0327"), String{
			r:  []rune{'C', '\u0327'},
			gc: intSliceRef(2),
		}},
		{"multi-grapheme string, odd count", New("kcutsemoh"), String{
			r:  []rune{'h', 'o', 'm', 'e', 's', 't', 'u', 'c', 'k'},
			gc: intSliceRef(1, 2, 3, 4, 5, 6, 7, 8, 9),
		}},
		{"multi-grapheme string, even count", New("bulg"), String{
			r:  []rune{'g', 'l', 'u', 'b'},
			gc: intSliceRef(1, 2, 3, 4),
		}},
		{"anagram check, odd count", New("RACECAR"), String{
			r:  []rune{'R', 'A', 'C', 'E', 'C', 'A', 'R'},
			gc: intSliceRef(1, 2, 3, 4, 5, 6, 7),
		}},
		{"anagram check, even count", New("HANNAH"), String{
			r:  []rune{'H', 'A', 'N', 'N', 'A', 'H'},
			gc: intSliceRef(1, 2, 3, 4, 5, 6),
		}},
		{"multiple multi-rune graphemes", New("Zero:\u0030\uFE0F\u20E3, Num:\u0023\uFE0F\u20E3"), String{
			r:  []rune{'\u0023', '\uFE0F', '\u20E3', ':', 'm', 'u', 'N', ' ', ',', '\u0030', '\uFE0F', '\u20E3', ':', 'o', 'r', 'e', 'Z'},
			gc: intSliceRef(3, 4, 5, 6, 7, 8, 9, 12, 13, 14, 15, 16, 17),
		}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert := assert.New(t)

			actual := tc.str.Reverse()

			assert.Equal(tc.expect, actual)
		})
	}
}

func Test_String_Runes(t *testing.T) {
	testCases := []struct {
		name   string
		str    String
		expect []rune
	}{
		{"empty string", Zero, []rune{}},
		{"one-char string", New("1"), []rune{'1'}},
		{"multi-char string", New("test"), []rune{'t', 'e', 's', 't'}},
		{"multi-char string with combining mark", New("FRANC\u0327AIS"), []rune{'F', 'R', 'A', 'N', 'C', '\u0327', 'A', 'I', 'S'}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert := assert.New(t)

			actual := tc.str.Runes()

			assert.Equal(tc.expect, actual)
		})
	}
}

func Test_String_Less(t *testing.T) {
	testCases := []struct {
		name     string
		leftStr  String
		rightStr String
		expect   bool
	}{
		{"empty string !< empty string", Zero, Zero, false},
		{"empty string < 1-char string", Zero, New("1"), true},
		{"empty string < 2-char string", Zero, New("12"), true},
		{"empty string < 3-char string", Zero, New("123"), true},
		{"1-char string !< empty string", New("1"), Zero, false},
		{"2-char string !< empty string", New("12"), Zero, false},
		{"3-char string !< empty string", New("123"), Zero, false},
		{"large string < small string", New("aadvark"), New("abby"), true},
		{"small string !< large string", New("abby"), New("aardvark"), false},
		{"large string !< small string", New("testing"), New("glub"), false},
		{"small string < large string", New("glub"), New("testing"), true},
		{"string < equal length string", New("glub"), New("test"), true},
		{"string !< equal length string", New("test"), New("glub"), false},
		{"two strings same except for last char, left < right", New("test"), New("tesu"), true},
		{"two strings same except for last char, left !< right", New("tesu"), New("test"), false},
		{"string !< itself", New("test"), New("test"), false},
		{"string < itself + suffix", New("test"), New("test1"), true},
		{"string !< itself - suffix", New("test1"), New("test"), false},

		// library does not support full collation at this time so a c-with-cedilla
		// and c followed by c with cedilla will not be sorted as the same
		{"combining sequence is less than precomposed", New("comment c\u0327a vaSUFFIX"), New("comment ça va"), true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert := assert.New(t)

			actual := tc.leftStr.Less(tc.rightStr)

			assert.Equal(tc.expect, actual)
		})
	}
}

func Test_String_IsEmpty(t *testing.T) {
	testCases := []struct {
		name   string
		str    String
		expect bool
	}{
		{"zero", Zero, true},
		{"empty, manually created", New(""), true},
		{"one char", New("1"), false},
		{"many chars", New("test test test"), false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert := assert.New(t)

			actual := tc.str.IsEmpty()

			assert.Equal(tc.expect, actual)
		})
	}
}

func Test_String_Sub(t *testing.T) {
	testCases := []struct {
		name   string
		str    String
		start  int
		end    int
		expect String
	}{
		{"empty string, 0:0 is allowed", Zero, 0, 0, Zero},
		{"empty string, 0:1 == 0:0", Zero, 0, 1, Zero},
		{"empty string, 1:0 == 0:0", Zero, 1, 0, Zero},
		{"empty string, -1:0 == 0:0", Zero, -1, 0, Zero},
		{"empty string, 0:-1 == 0:0", Zero, 0, -1, Zero},

		{"1-char string, 0:0 is allowed", New("1"), 0, 0, Zero},
		{"1-char string, 1:1 is allowed", New("1"), 1, 1, Zero},
		{"1-char string, 0:1 gives back string", New("1"), 0, 1, New("1")},
		{"1-char string, -1:0 is allowed (and is same as 0:0)", New("1"), -1, 0, Zero},
		{"1-char string, 0:-1 is allowed (and is same as 0:0)", New("1"), 0, -1, Zero},
		{"1-char string, -2:1 == 0:1", New("1"), -2, 1, New("1")},
		{"1-char string, 0:-2 == 0:0", New("1"), 0, -2, Zero},
		{"1-grapheme string, decomposed sequence is preserved", New("C\u0327"), 0, 1, New("C\u0327")},

		{"8-char string, 0:8 gives back string", New("Test1234"), 0, 8, New("Test1234")},
		{"8-char string, 1:8", New("Test1234"), 1, 8, New("est1234")},
		{"8-char string, both negative indexes", New("Test1234"), -5, -2, New("t12")},
		{"8-grapheme string, decomposed sequence is preserved at end", New("Test-C\u0327as"), 0, 6, New("Test-C\u0327")},
		{"8-grapheme string, decomposed sequence is preserved at start", New("Test-C\u0327as"), 5, 7, New("C\u0327a")},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert := assert.New(t)

			actual := tc.str.Sub(tc.start, tc.end)
			isEqual := actual.Equal(tc.expect)

			assert.Equal(true, isEqual)
		})
	}
}

func Test_String_Index(t *testing.T) {
	testCases := []struct {
		name   string
		str    String
		search String
		expect int
	}{
		{"empty string in empty string", Zero, Zero, -1},
		{"empty string in non-empty string", New("TEST"), Zero, 0},
		{"long string in short string", New("GLUB"), New("GLUUUUUUUB"), -1},
		{"not present", New("Some long string"), New("GLUB"), -1},
		{"exact match", New("GLUB"), New("GLUB"), 0},
		{"single-char search, is at start", New("GLUB"), New("G"), 0},
		{"multi-char search, is at start", New("GLUB"), New("GLU"), 0},
		{"single-char search, is in middle", New("Just go glub at it! glub glub!"), New("g"), 5},
		{"multi-char search, is in middle", New("Just go glub at it! glub glub!"), New("glub"), 8},
		{"false match, is in middle", New("Just glue glub on it! glub glub!"), New("glub"), 10},
		{"single-char search, is at end", New("say GLUB"), New("B"), 7},
		{"multi-char search, is at end", New("say GLUB"), New("GLUB"), 4},
		{"search has multi-rune grapheme", New("I said that \u0023\uFE0F\u20E3 is # but in emote form"), New("that \u0023\uFE0F\u20E3 is #"), 7},
		{"text has multi-rune graphemes before match", New("\u0023\uFE0F\u20E3 is an emoji"), New("is"), 2},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert := assert.New(t)

			actual := tc.str.Index(tc.search)

			assert.Equal(tc.expect, actual)
		})
	}
}

func Test_String_IndexFunc(t *testing.T) {
	testCases := []struct {
		name   string
		str    String
		f      func([]rune) bool
		expect int
	}{
		{
			name:   "empty string always returns -1",
			str:    Zero,
			f:      func([]rune) bool { return true },
			expect: -1,
		},
		{
			name:   "matching func",
			str:    New("gallowsCalibrator arachnidsGrip"),
			f:      func(gc []rune) bool { return unicode.IsUpper(gc[0]) },
			expect: 7,
		},
		{
			name:   "non-matching func",
			str:    New("several letters but an upper-case 'z' is not one"),
			f:      func(gc []rune) bool { return string(gc) == "Z" },
			expect: -1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert := assert.New(t)

			actual := tc.str.IndexFunc(tc.f)

			assert.Equal(tc.expect, actual)
		})
	}
}

func Test_String_LastIndex(t *testing.T) {
	testCases := []struct {
		name   string
		str    String
		search String
		expect int
	}{
		{"empty string in empty string", Zero, Zero, -1},
		{"empty string in non-empty string", New("TEST"), Zero, 4},
		{"long string in short string", New("GLUB"), New("GLUUUUUUUB"), -1},
		{"not present", New("Some long string"), New("GLUB"), -1},
		{"exact match", New("GLUB"), New("GLUB"), 0},
		{"single-char search, is at start", New("GLUB"), New("G"), 0},
		{"multi-char search, is at start", New("GLUB"), New("GLU"), 0},
		{"single-char search, is in middle", New("Just go glub at it! glub yeah!"), New("g"), 20},
		{"multi-char search, is in middle", New("Just go glub at it! glub yeah!"), New("glub"), 20},
		{"false match, is in middle", New("Just glue glub on it! glub elub!"), New("glub"), 22},
		{"single-char search, is at end", New("say GLUB"), New("B"), 7},
		{"multi-char search, is at end", New("say GLUB"), New("GLUB"), 4},
		{"search has multi-rune grapheme", New("I said that \u0023\uFE0F\u20E3 is # but in emote form"), New("that \u0023\uFE0F\u20E3 is #"), 7},
		{"text has multi-rune graphemes before match", New("\u0023\uFE0F\u20E3 is an emoji"), New("is"), 2},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert := assert.New(t)

			actual := tc.str.LastIndex(tc.search)

			assert.Equal(tc.expect, actual)
		})
	}
}

func Test_String_LastIndexFunc(t *testing.T) {
	testCases := []struct {
		name   string
		str    String
		f      func([]rune) bool
		expect int
	}{
		{
			name:   "empty string always returns -1",
			str:    Zero,
			f:      func([]rune) bool { return true },
			expect: -1,
		},
		{
			name:   "matching func",
			str:    New("gallowsCalibrator arachnidsGrip"),
			f:      func(gc []rune) bool { return unicode.IsUpper(gc[0]) },
			expect: 27,
		},
		{
			name:   "non-matching func",
			str:    New("several letters but an upper-case 'z' is not one"),
			f:      func(gc []rune) bool { return string(gc) == "Z" },
			expect: -1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert := assert.New(t)

			actual := tc.str.LastIndexFunc(tc.f)

			assert.Equal(tc.expect, actual)
		})
	}
}
