package rosed

import (
	"reflect"
	"testing"

	"github.com/dekarrin/rosed/internal/assert"
)

type compBlock Block

func (cb compBlock) Equal(other interface{}) bool {
	b1 := Block(cb)
	b2, ok := other.(Block)
	if !ok {
		return false
	}

	if b1.LineSeparator != b2.LineSeparator {
		return false
	}
	if b1.TrailingSeparator != b2.TrailingSeparator {
		return false
	}
	return reflect.DeepEqual(b1.Lines, b2.Lines)
}

func Test_NewBlock(t *testing.T) {
	testCases := []struct {
		name     string
		text     string
		sep      string
		expected Block
	}{
		{
			name: "no lines",
			text: "",
			sep:  "\n",
			expected: Block{
				Lines: []string{
					"a",
				},
				LineSeparator:     "\n",
				TrailingSeparator: false,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			asrt := assert.New(t)

			actual := NewBlock(tc.text, tc.sep)

			asrt.Var("block").Equal(compBlock(tc.expected), compBlock(actual))

		})
	}
}
