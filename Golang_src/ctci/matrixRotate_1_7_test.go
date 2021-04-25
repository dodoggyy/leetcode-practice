package ctci

import (
	"reflect"
	"testing"
)

func TestMatrixRotate(t *testing.T) {
	cases := []struct {
		input    [][]int
		expected [][]int
	}{
		{
			[][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			[][]int{
				{7, 4, 1},
				{8, 5, 2},
				{9, 6, 3},
			},
		},
		{
			[][]int{},
			[][]int{},
		},
		{
			nil,
			nil,
		},
	}

	for _, c := range cases {
		MatrixRotate(c.input)
		if !reflect.DeepEqual(c.input, c.expected) {
			t.Fatalf("Expected: %v, actual: %v\n", c.expected, c.input)
		}
	}
}
