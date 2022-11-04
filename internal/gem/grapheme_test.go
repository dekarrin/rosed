package gem

import (
	"testing"

	"github.com/dekarrin/assertion"
)

func Test_New(t *testing.T) {
	testCases := []struct {
		name string
		input string
		expect String
	}{
		{"empty string", "", String{r: []rune{}}},
		{"one-char string", "1", String{r: []rune{'1'}}},
		{"two-char string", "12", String{r: []rune{'1', '2'}}},
		{"two-char string", "123", String{r: []rune{'1', '2', '3'}}},
	}
	
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert := assertion.New(t)
			actual := New(tc.input)
			assert.Equal(tc.expect, actual)
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
			assert := assertion.New(t)
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
			assert := assertion.New(t)
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
