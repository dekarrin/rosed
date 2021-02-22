package gem

import (
	"testing"

	"github.com/dekarrin/assertion"
)

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
		{"Zero == &Zero", Zero, &Zero, true},
		{"Zero == nil", Zero, nil, true},
		{"Zero == New empty", Zero, emptyStr, true},
		{"Zero == &New empty", Zero, &emptyStr, true},
		{"New empty == Zero", emptyStr, Zero, true},
		{"New empty == &Zero", emptyStr, &Zero, true},
		{"New empty == New empty", emptyStr, emptyStr, true},
		{"New empty == &New empty", emptyStr, &emptyStr, true},
		{"filled == filled", testStr, testStr, true},
		{"filled == &filled", testStr, &testStr, true},
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
		{"&Zero == &Zero", &Zero, &Zero, true},
		{"&Zero == nil", &Zero, nil, true},
		{"nil == Zero", nil, Zero, true},
		{"&Zero == New empty", &Zero, emptyStr, true},
		{"&Zero == &New empty", &Zero, &emptyStr, true},
		{"&New empty == Zero", &emptyStr, Zero, true},
		{"&New empty == &Zero", &emptyStr, &Zero, true},
		{"&New empty == New empty", &emptyStr, emptyStr, true},
		{"&New empty == &New empty", &emptyStr, &emptyStr, true},
		{"&filled == filled", &testStr, testStr, true},
		{"&filled == &filled", &testStr, &testStr, true},
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
			actual := tc.input.Equal(tc.compare)
			assert.Equal(tc.expect, actual)
		})
	}
}
