package rosed

import (
	"reflect"
	"testing"
)

func Test_Manip_Wrap(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		width    int
		sep      string
		expected []string
	}{
		{
			name:  "empty input",
			input: "",
			width: 80,
			sep:   "\n",
			expected: []string{
				"",
			},
		},
		{
			name:  "not enough to wrap",
			input: "a test string",
			width: 80,
			sep:   "\n",
			expected: []string{
				"a test string",
			},
		},
		{
			name:  "2 line wrap",
			input: "a string long enough to be wrapped",
			width: 20,
			sep:   "\n",
			expected: []string{
				"a string long enough",
				"to be wrapped",
			},
		},
		{
			name:  "multi line wrap",
			input: "a string long enough to be wrapped more than once",
			width: 20,
			sep:   "\n",
			expected: []string{
				"a string long enough",
				"to be wrapped more",
				"than once",
			},
		},
		{
			name:  "invalid width of -1 is interpreted as 2",
			input: "test",
			width: -1,
			sep:   "\n",
			expected: []string{
				"t-",
				"e-",
				"st",
			},
		},
		{
			name:  "invalid width of 0 is interpreted as 2",
			input: "test",
			width: 0,
			sep:   "\n",
			expected: []string{
				"t-",
				"e-",
				"st",
			},
		},
		{
			name:  "invalid width of 1 is interpreted as 2",
			input: "test",
			width: 1,
			sep:   "\n",
			expected: []string{
				"t-",
				"e-",
				"st",
			},
		},
		{
			name:  "valid width of 2",
			input: "test",
			width: 2,
			sep:   "\n",
			expected: []string{
				"t-",
				"e-",
				"st",
			},
		},
		{
			name:  "valid width of 3",
			input: "test",
			width: 3,
			sep:   "\n",
			expected: []string{
				"te-",
				"st",
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := wrap(tc.input, tc.width, tc.sep).Lines

			if !reflect.DeepEqual(actual, tc.input) {
				t.Fatalf("expected result to be:\n%q\nbut was:\n%q", tc.expected, actual)
			}
		})
	}
}
