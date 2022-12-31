package gem

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Repeat(t *testing.T) {
	testCases := []struct {
		name   string
		s      String
		count  int
		expect String
	}{
		{"repeat empty string -1 times", Zero, -1, Zero},
		{"repeat empty string 0 times", Zero, 0, Zero},
		{"repeat empty string 1 time", Zero, 1, Zero},
		{"repeat empty string 2 times", Zero, 2, Zero},
		{"repeat empty string 3 times", Zero, 3, Zero},
		{"repeat single char string -1 times", New("8"), -1, Zero},
		{"repeat single char string 0 times", New("8"), 0, Zero},
		{"repeat single char string 1 time", New("8"), 1, New("8")},
		{"repeat single char string 2 times", New("8"), 2, New("88")},
		{"repeat single char string 3 times", New("8"), 3, New("888")},
		{"repeat multi char string -1 times", New("BLUH"), -1, Zero},
		{"repeat multi char string 0 times", New("BLUH"), 0, Zero},
		{"repeat multi char string 1 time", New("BLUH"), 1, New("BLUH")},
		{"repeat multi char string 2 times", New("BLUH"), 2, New("BLUHBLUH")},
		{"repeat multi char string 3 times", New("BLUH"), 3, New("BLUHBLUHBLUH")},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert := assert.New(t)

			actual := Repeat(tc.s, tc.count)

			assert.True(tc.expect.Equal(actual))
			assert.Equal(tc.expect.String(), actual.String())
		})
	}
}

func Test_Strings(t *testing.T) {
	testCases := []struct {
		name   string
		input  []String
		expect []string
	}{
		{"empty slices", []String{}, []string{}},
		{"1 empty string", []String{Zero}, []string{""}},
		{"2 strings", []String{New("hello"), New("world")}, []string{"hello", "world"}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert := assert.New(t)

			actual := Strings(tc.input)

			assert.Equal(tc.expect, actual)
		})
	}
}

func Test_Slice(t *testing.T) {
	testCases := []struct {
		name   string
		input  []string
		expect []String
	}{
		{"empty slices", []string{}, []String{}},
		{"1 empty string", []string{""}, []String{Zero}},
		{"2 strings", []string{"hello", "world"}, []String{New("hello"), New("world")}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert := assert.New(t)

			actual := Slice(tc.input)

			// we have custom equality on gem strings so need to check each one manually
			assert.Len(actual, len(tc.expect))

			for idx := range actual {
				assert.True(tc.expect[idx].Equal(actual[idx]))
			}
		})
	}
}

func Test_Split(t *testing.T) {
	testCases := []struct {
		name   string
		input  []rune
		expect []int
	}{
		{"empty string", []rune{}, []int{}},
		{"normal string", []rune{'t', 'e', 's', 't'}, []int{1, 2, 3, 4}},
		{"string with decomposed sequence", []rune{'C', '\u0327', 'e', 's', 't'}, []int{2, 3, 4, 5}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert := assert.New(t)

			actual := Split(tc.input)

			assert.Equal(tc.expect, actual)
		})
	}
}
