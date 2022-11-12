package util


import (
	"testing"
	
	"github.com/stretchr/testify/assert"
)

func Test_RangeToIndexes(t *testing.T) {
	type pair struct {
		start int
		end int
	}

	testCases := []struct{
		name string
		size int
		inputRange pair
		expect pair
	}{
		{
			name: "all zero",
			size: 0,
			inputRange: pair{0, 0},
			expect: pair{0, 0},
		},
		{
			name: "non-zero size, all-zero range",
			size: 1,
			inputRange: pair{0, 0},
			expect: pair{0, 0},
		},
		{
			name: "start > end -> end = start",
			size: 20,
			inputRange: pair{1, 0},
			expect: pair{1, 1},
		},
		{
			name: "negative start",
			size: 10,
			inputRange: pair{-5, 9},
			expect: pair{5, 9},
		},
		{
			name: "negative end",
			size: 10,
			inputRange: pair{5, -2},
			expect: pair{5, 8},
		},
	}
	
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert := assert.New(t)
			
			start, end := RangeToIndexes(tc.size, tc.inputRange.start, tc.inputRange.end)
			actual := pair{start, end}
			
			assert.Equal(tc.expect, actual)
		})	
	}
}
