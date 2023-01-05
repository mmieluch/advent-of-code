package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Matrix_Row(t *testing.T) {
	input := matrix{
		row{3, 0, 3, 7, 3},
		row{2, 5, 5, 1, 2},
		row{6, 5, 3, 3, 2},
		row{3, 3, 5, 4, 9},
		row{3, 5, 3, 9, 0},
	}
	tests := map[string]struct {
		y        int
		expected row
	}{
		"matrix row [0]": {0, row{3, 0, 3, 7, 3}},
		"matrix row [1]": {1, row{2, 5, 5, 1, 2}},
		"matrix row [2]": {2, row{6, 5, 3, 3, 2}},
		"matrix row [3]": {3, row{3, 3, 5, 4, 9}},
		"matrix row [4]": {4, row{3, 5, 3, 9, 0}},
	}
	for name, test := range tests {
		assert.Equal(t, test.expected, input.Row(test.y), name)
	}
}

func Test_Matrix_Col(t *testing.T) {
	input := matrix{
		row{3, 0, 3, 7, 3},
		row{2, 5, 5, 1, 2},
		row{6, 5, 3, 3, 2},
		row{3, 3, 5, 4, 9},
		row{3, 5, 3, 9, 0},
	}
	tests := map[string]struct {
		x        int
		expected []uint
	}{
		"matrix col [0]": {0, []uint{3, 2, 6, 3, 3}},
		"matrix col [1]": {1, []uint{0, 5, 5, 3, 5}},
		"matrix col [2]": {2, []uint{3, 5, 3, 5, 3}},
		"matrix col [3]": {3, []uint{7, 1, 3, 4, 9}},
		"matrix col [4]": {4, []uint{3, 2, 2, 9, 0}},
	}
	for name, test := range tests {
		assert.Equal(t, test.expected, input.Col(test.x), name)
	}
}

func Test_Matrix_Pos(t *testing.T) {
	input := matrix{
		row{3, 0, 3, 7, 3},
		row{2, 5, 5, 1, 2},
		row{6, 5, 3, 3, 2},
		row{3, 3, 5, 4, 9},
		row{3, 5, 3, 9, 0},
	}
	tests := map[string]struct {
		x        int
		y        int
		expected uint
	}{
		"matrix col [0][0]": {0, 0, 3},
		"matrix col [0][1]": {0, 1, 2},
		"matrix col [0][2]": {0, 2, 6},
		"matrix col [0][3]": {0, 3, 3},
		"matrix col [0][4]": {0, 4, 3},

		"matrix col [1][0]": {1, 0, 0},
		"matrix col [1][1]": {1, 1, 5},
		"matrix col [1][2]": {1, 2, 5},
		"matrix col [1][3]": {1, 3, 3},
		"matrix col [1][4]": {1, 4, 5},

		"matrix col [2][0]": {2, 0, 3},
		"matrix col [2][1]": {2, 1, 5},
		"matrix col [2][2]": {2, 2, 3},
		"matrix col [2][3]": {2, 3, 5},
		"matrix col [2][4]": {2, 4, 3},

		"matrix col [3][0]": {3, 0, 7},
		"matrix col [3][1]": {3, 1, 1},
		"matrix col [3][2]": {3, 2, 3},
		"matrix col [3][3]": {3, 3, 4},
		"matrix col [3][4]": {3, 4, 9},

		"matrix col [4][0]": {4, 0, 3},
		"matrix col [4][1]": {4, 1, 2},
		"matrix col [4][2]": {4, 2, 2},
		"matrix col [4][3]": {4, 3, 9},
		"matrix col [4][4]": {4, 4, 0},
	}
	for name, test := range tests {
		assert.Equal(t, test.expected, input.Pos(test.x, test.y), name)
	}
}

func Test_Matrix_LeftPos(t *testing.T) {
	input := matrix{
		row{3, 0, 3, 7, 3},
		row{2, 5, 5, 1, 2},
		row{6, 5, 3, 3, 2},
		row{3, 3, 5, 4, 9},
		row{3, 5, 3, 9, 0},
	}
	tests := []struct {
		x        int
		y        int
		expected []uint
		msg      string
	}{
		// Row 1 (matrix[0])
		{0, 0, []uint{}, "numbers left to [0][0]"},
		{1, 0, []uint{3}, "numbers left to [1][0]"},
		{2, 0, []uint{3, 0}, "numbers left to [2][0]"},
		{3, 0, []uint{3, 0, 3}, "numbers left to [3][0]"},
		{4, 0, []uint{3, 0, 3, 7}, "numbers left to [4][0]"},

		// Row 2 (matrix[1])
		{0, 1, []uint{}, "numbers left to [0][1]"},
		{1, 1, []uint{2}, "numbers left to [1][1]"},
		{2, 1, []uint{2, 5}, "numbers left to [2][1]"},
		{3, 1, []uint{2, 5, 5}, "numbers left to [3][1]"},
		{4, 1, []uint{2, 5, 5, 1}, "numbers left to [4][1]"},

		// Row 3 (matrix[2])
		{0, 2, []uint{}, "numbers left to [0][2]"},
		{1, 2, []uint{6}, "numbers left to [1][2]"},
		{2, 2, []uint{6, 5}, "numbers left to [2][2]"},
		{3, 2, []uint{6, 5, 3}, "numbers left to [3][2]"},
		{4, 2, []uint{6, 5, 3, 3}, "numbers left to [4][2]"},

		// Row 4 (matrix[3])
		{0, 3, []uint{}, "numbers left to [0][3]"},
		{1, 3, []uint{3}, "numbers left to [1][3]"},
		{2, 3, []uint{3, 3}, "numbers left to [2][3]"},
		{3, 3, []uint{3, 3, 5}, "numbers left to [3][3]"},
		{4, 3, []uint{3, 3, 5, 4}, "numbers left to [4][3]"},

		// Row 5 (matrix[4])
		{0, 4, []uint{}, "numbers left to [0][4]"},
		{1, 4, []uint{3}, "numbers left to [1][4]"},
		{2, 4, []uint{3, 5}, "numbers left to [2][4]"},
		{3, 4, []uint{3, 5, 3}, "numbers left to [3][4]"},
		{4, 4, []uint{3, 5, 3, 9}, "numbers left to [4][4]"},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, input.LeftPos(test.x, test.y), test.msg)
	}
}

func Test_Matrix_RightPos(t *testing.T) {
	input := matrix{
		row{3, 0, 3, 7, 3},
		row{2, 5, 5, 1, 2},
		row{6, 5, 3, 3, 2},
		row{3, 3, 5, 4, 9},
		row{3, 5, 3, 9, 0},
	}
	tests := []struct {
		x        int
		y        int
		expected []uint
		msg      string
	}{
		// Row 1 (matrix[0])
		{0, 0, []uint{0, 3, 7, 3}, "numbers right to [0][0]"},
		{1, 0, []uint{3, 7, 3}, "numbers right to [1][0]"},
		{2, 0, []uint{7, 3}, "numbers right to [2][0]"},
		{3, 0, []uint{3}, "numbers right to [3][0]"},
		{4, 0, []uint{}, "numbers right to [4][0]"},

		// Row 2 (matrix[1])
		{0, 1, []uint{5, 5, 1, 2}, "numbers right to [0][1]"},
		{1, 1, []uint{5, 1, 2}, "numbers right to [1][1]"},
		{2, 1, []uint{1, 2}, "numbers right to [2][1]"},
		{3, 1, []uint{2}, "numbers right to [3][1]"},
		{4, 1, []uint{}, "numbers right to [4][1]"},

		// Row 3 (matrix[2])
		{0, 2, []uint{5, 3, 3, 2}, "numbers right to [0][2]"},
		{1, 2, []uint{3, 3, 2}, "numbers right to [1][2]"},
		{2, 2, []uint{3, 2}, "numbers right to [2][2]"},
		{3, 2, []uint{2}, "numbers right to [3][2]"},
		{4, 2, []uint{}, "numbers right to [4][2]"},

		// Row 4 (matrix[3])
		{0, 3, []uint{3, 5, 4, 9}, "numbers right to [0][3]"},
		{1, 3, []uint{5, 4, 9}, "numbers right to [1][3]"},
		{2, 3, []uint{4, 9}, "numbers right to [2][3]"},
		{3, 3, []uint{9}, "numbers right to [3][3]"},
		{4, 3, []uint{}, "numbers right to [4][3]"},

		// Row 5 (matrix[4])
		{0, 4, []uint{5, 3, 9, 0}, "numbers right to [0][4]"},
		{1, 4, []uint{3, 9, 0}, "numbers right to [1][4]"},
		{2, 4, []uint{9, 0}, "numbers right to [2][4]"},
		{3, 4, []uint{0}, "numbers right to [3][4]"},
		{4, 4, []uint{}, "numbers right to [4][4]"},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, input.RightPos(test.x, test.y), test.msg)
	}
}

func Test_Matrix_AbovePos(t *testing.T) {
	input := matrix{
		row{3, 0, 3, 7, 3},
		row{2, 5, 5, 1, 2},
		row{6, 5, 3, 3, 2},
		row{3, 3, 5, 4, 9},
		row{3, 5, 3, 9, 0},
	}
	tests := []struct {
		x        int
		y        int
		expected []uint
		msg      string
	}{
		// Column 1 (offset 0)
		{0, 0, []uint{}, "above [0][0]"},
		{0, 1, []uint{3}, "above [0][1]"},
		{0, 2, []uint{3, 2}, "above [0][2]"},
		{0, 3, []uint{3, 2, 6}, "above [0][3]"},
		{0, 4, []uint{3, 2, 6, 3}, "above [0][4]"},

		// Column 2 (offset 1)
		{1, 0, []uint{}, "above [1][0]"},
		{1, 1, []uint{0}, "above [1][1]"},
		{1, 2, []uint{0, 5}, "above [1][2]"},
		{1, 3, []uint{0, 5, 5}, "above [1][3]"},
		{1, 4, []uint{0, 5, 5, 3}, "above [1][4]"},

		// Column 3 (offset 2)
		{2, 0, []uint{}, "above [2][0]"},
		{2, 1, []uint{3}, "above [2][1]"},
		{2, 2, []uint{3, 5}, "above [2][2]"},
		{2, 3, []uint{3, 5, 3}, "above [2][3]"},
		{2, 4, []uint{3, 5, 3, 5}, "above [2][4]"},

		// Column 4 (offset 3)
		{3, 0, []uint{}, "above [3][0]"},
		{3, 1, []uint{7}, "above [3][1]"},
		{3, 2, []uint{7, 1}, "above [3][2]"},
		{3, 3, []uint{7, 1, 3}, "above [3][3]"},
		{3, 4, []uint{7, 1, 3, 4}, "above [3][4]"},

		// Column 5 (offset 4)
		{4, 0, []uint{}, "above [4][0]"},
		{4, 1, []uint{3}, "above [4][1]"},
		{4, 2, []uint{3, 2}, "above [4][2]"},
		{4, 3, []uint{3, 2, 2}, "above [4][3]"},
		{4, 4, []uint{3, 2, 2, 9}, "above [4][4]"},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, input.AbovePos(test.x, test.y), test.msg)
	}
}

func Test_Matrix_BelowPos(t *testing.T) {
	input := matrix{
		row{3, 0, 3, 7, 3},
		row{2, 5, 5, 1, 2},
		row{6, 5, 3, 3, 2},
		row{3, 3, 5, 4, 9},
		row{3, 5, 3, 9, 0},
	}
	tests := []struct {
		x        int
		y        int
		expected []uint
		msg      string
	}{
		// Column 1 (offset 0)
		{0, 0, []uint{2, 6, 3, 3}, "below [0][0]"},
		{0, 1, []uint{6, 3, 3}, "below [0][1]"},
		{0, 2, []uint{3, 3}, "below [0][2]"},
		{0, 3, []uint{3}, "below [0][3]"},
		{0, 4, []uint{}, "below [0][4]"},

		// Column 2 (offset 1)
		{1, 0, []uint{5, 5, 3, 5}, "below [1][0]"},
		{1, 1, []uint{5, 3, 5}, "below [1][1]"},
		{1, 2, []uint{3, 5}, "below [1][2]"},
		{1, 3, []uint{5}, "below [1][3]"},
		{1, 4, []uint{}, "below [1][4]"},

		// Column 3 (offset 2)
		{2, 0, []uint{5, 3, 5, 3}, "below [2][0]"},
		{2, 1, []uint{3, 5, 3}, "below [2][1]"},
		{2, 2, []uint{5, 3}, "below [2][2]"},
		{2, 3, []uint{3}, "below [2][3]"},
		{2, 4, []uint{}, "below [2][4]"},

		// Column 4 (offset 3)
		{3, 0, []uint{1, 3, 4, 9}, "below [3][0]"},
		{3, 1, []uint{3, 4, 9}, "below [3][1]"},
		{3, 2, []uint{4, 9}, "below [3][2]"},
		{3, 3, []uint{9}, "below [3][3]"},
		{3, 4, []uint{}, "below [3][4]"},

		// Column 5 (offset 4)
		{4, 0, []uint{2, 2, 9, 0}, "below [4][0]"},
		{4, 1, []uint{2, 9, 0}, "below [4][1]"},
		{4, 2, []uint{9, 0}, "below [4][2]"},
		{4, 3, []uint{0}, "below [4][3]"},
		{4, 4, []uint{}, "below [4][4]"},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, input.BelowPos(test.x, test.y), test.msg)
	}
}

func TestParse(t *testing.T) {
	input := `30373
25512
65332
33549
35390`
	expected := matrix{
		row{3, 0, 3, 7, 3},
		row{2, 5, 5, 1, 2},
		row{6, 5, 3, 3, 2},
		row{3, 3, 5, 4, 9},
		row{3, 5, 3, 9, 0},
	}
	assert.Equal(t, expected, parse(input))
}
