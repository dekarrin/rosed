package rosed

import (
	"testing"

	"github.com/dekarrin/assertion"
	"github.com/dekarrin/rosed/internal/gem"
)

func Test_Manip_collapseSpace(t *testing.T) {
	testCases := []struct {
		name   string
		input  gem.String
		sep    gem.String
		expect gem.String
	}{
		{"empty input", gem.Zero, gem.Zero, gem.Zero},
		{"no space to collapse", _g("no_spaces"), gem.Zero, _g("no_spaces")},
		{"one space mid-text is not collapsed (odd runecount)", _g("testA testB"), gem.Zero, _g("testA testB")},
		{"one space mid-text is not collapsed (even runecount)", _g("test testB"), gem.Zero, _g("test testB")},
		{"one space at start is not collapsed (odd runecount)", _g(" test"), gem.Zero, _g(" test")},
		{"one space at start is not collapsed (even runecount)", _g(" testA"), gem.Zero, _g(" testA")},
		{"one space at end is not collapsed (odd runecount)", _g("test "), gem.Zero, _g("test ")},
		{"one space at end is not collapsed (even runecount)", _g("testA "), gem.Zero, _g("testA ")},
		{"one space everywhere is not collapsed (odd runecount)", _g(" testA testB "), gem.Zero, _g(" testA testB ")},
		{"one space everywhere is not collapsed (even runecount)", _g(" test testB "), gem.Zero, _g(" test testB ")},
		{"non-spacechar whitespace is converted to space (odd runecount)", _g("testA\u0085testB"), gem.Zero, _g("testA testB")},
		{"non-spacechar whitespace is converted to space (even runecount)", _g("test\u0085testB"), gem.Zero, _g("test testB")},
		{"ws run is collapsed (spacechar)", _g("       testA  testB  "), gem.Zero, _g(" testA testB ")},
		{"ws run is collapsed (mixed ws)", _g("\u205f\u202ftestA\u200a  \t\n testB\t"), gem.Zero, _g(" testA testB ")},
		{"non-ws separator", _g("testA\n  testB <SEP>\u205f  testC"), _g("<SEP>"), _g("testA testB testC")},
		{"ws separator", _g("testA\n  testB <SEP>\u205f  testC"), _g("<SEP>"), _g("testA testB <SEP> testC")},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert := assertion.New(t)
			actual := collapseSpace(tc.input, tc.sep)
			assert.Equal(tc.expect, actual)
		})
	}
}

func Test_Manip_Wrap(t *testing.T) {
	testCases := []struct {
		name     string
		input    gem.String
		width    int
		sep      gem.String
		expected block
	}{
		{
			name:  "empty input",
			input: gem.Zero,
			width: 80,
			sep:   _g("\n"),
			expected: block{Lines: []gem.String{
				gem.Zero,
			}},
		},
		{
			name:  "not enough to wrap",
			input: _g("a test string"),
			width: 80,
			sep:   _g("\n"),
			expected: block{Lines: []gem.String{
				_g("a test string"),
			}},
		},
		{
			name:  "2 line wrap",
			input: _g("a string long enough to be wrapped"),
			width: 20,
			sep:   _g("\n"),
			expected: block{Lines: []gem.String{
				_g("a string long enough"),
				_g("to be wrapped"),
			}},
		},
		{
			name:  "multi line wrap",
			input: _g("a string long enough to be wrapped more than once"),
			width: 20,
			sep:   _g("\n"),
			expected: block{Lines: []gem.String{
				_g("a string long enough"),
				_g("to be wrapped more"),
				_g("than once"),
			}},
		},
		{
			name:  "invalid width of -1 is interpreted as 2",
			input: _g("test"),
			width: -1,
			sep:   _g("\n"),
			expected: block{Lines: []gem.String{
				_g("t-"),
				_g("e-"),
				_g("st"),
			}},
		},
		{
			name:  "invalid width of 0 is interpreted as 2",
			input: _g("test"),
			width: 0,
			sep:   _g("\n"),
			expected: block{Lines: []gem.String{
				_g("t-"),
				_g("e-"),
				_g("st"),
			}},
		},
		{
			name:  "invalid width of 1 is interpreted as 2",
			input: _g("test"),
			width: 1,
			sep:   _g("\n"),
			expected: block{Lines: []gem.String{
				_g("t-"),
				_g("e-"),
				_g("st"),
			}},
		},
		{
			name:  "valid width of 2",
			input: _g("test"),
			width: 2,
			sep:   _g("\n"),
			expected: block{Lines: []gem.String{
				_g("t-"),
				_g("e-"),
				_g("st"),
			}},
		},
		{
			name:  "valid width of 3",
			input: _g("test"),
			width: 3,
			sep:   _g("\n"),
			expected: block{Lines: []gem.String{
				_g("te-"),
				_g("st"),
			}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert := assertion.New(t)
			actual := wrap(tc.input, tc.width, tc.sep).Lines
			assert.EqualSlices(tc.expected, actual)
		})
	}
}
