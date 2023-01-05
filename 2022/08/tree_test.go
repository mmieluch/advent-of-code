package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTreeFromMatrix(t *testing.T) {
	m := matrix{
		row{3, 0, 3, 7, 3},
		row{2, 5, 5, 1, 2},
		row{6, 5, 3, 3, 2},
		row{3, 3, 5, 4, 9},
		row{3, 5, 3, 9, 0},
	}
	expected := [][]tree{
		{
			{3, []uint{}, []uint{0, 3, 7, 3}, []uint{}, []uint{2, 6, 3, 3}},
			{0, []uint{3}, []uint{3, 7, 3}, []uint{}, []uint{5, 5, 3, 5}},
			{3, []uint{3, 0}, []uint{7, 3}, []uint{}, []uint{5, 3, 5, 3}},
			{7, []uint{3, 0, 3}, []uint{3}, []uint{}, []uint{1, 3, 4, 9}},
			{3, []uint{3, 0, 3, 7}, []uint{}, []uint{}, []uint{2, 2, 9, 0}},
		},
		{
			{2, []uint{}, []uint{5, 5, 1, 2}, []uint{3}, []uint{6, 3, 3}},
			{5, []uint{2}, []uint{5, 1, 2}, []uint{0}, []uint{5, 3, 5}},
			{5, []uint{2, 5}, []uint{1, 2}, []uint{3}, []uint{3, 5, 3}},
			{1, []uint{2, 5, 5}, []uint{2}, []uint{7}, []uint{3, 4, 9}},
			{2, []uint{2, 5, 5, 1}, []uint{}, []uint{3}, []uint{2, 9, 0}},
		},
		{
			{6, []uint{}, []uint{5, 3, 3, 2}, []uint{3, 2}, []uint{3, 3}},
			{5, []uint{6}, []uint{3, 3, 2}, []uint{0, 5}, []uint{3, 5}},
			{3, []uint{6, 5}, []uint{3, 2}, []uint{3, 5}, []uint{5, 3}},
			{3, []uint{6, 5, 3}, []uint{2}, []uint{7, 1}, []uint{4, 9}},
			{2, []uint{6, 5, 3, 3}, []uint{}, []uint{3, 2}, []uint{9, 0}},
		},
		{
			{3, []uint{}, []uint{3, 5, 4, 9}, []uint{3, 2, 6}, []uint{3}},
			{3, []uint{3}, []uint{5, 4, 9}, []uint{0, 5, 5}, []uint{5}},
			{5, []uint{3, 3}, []uint{4, 9}, []uint{3, 5, 3}, []uint{3}},
			{4, []uint{3, 3, 5}, []uint{9}, []uint{7, 1, 3}, []uint{9}},
			{9, []uint{3, 3, 5, 4}, []uint{}, []uint{3, 2, 2}, []uint{0}},
		},
		{
			{3, []uint{}, []uint{5, 3, 9, 0}, []uint{3, 2, 6, 3}, []uint{}},
			{5, []uint{3}, []uint{3, 9, 0}, []uint{0, 5, 5, 3}, []uint{}},
			{3, []uint{3, 5}, []uint{9, 0}, []uint{3, 5, 3, 5}, []uint{}},
			{9, []uint{3, 5, 3}, []uint{0}, []uint{7, 1, 3, 4}, []uint{}},
			{0, []uint{3, 5, 3, 9}, []uint{}, []uint{3, 2, 2, 9}, []uint{}},
		},
	}

	for y, row := range expected {
		for x, tree := range row {
			assert.Equal(t, tree, newTreeFromMatrix(m, x, y), fmt.Sprintf("[%d][%d]", x, y))
		}
	}
}

func Test_Tree_IsVisible(t *testing.T) {
	tests := []struct {
		t        tree
		expected bool
		msg      string
	}{
		// Row 1
		{
			t:        tree{3, []uint{}, []uint{0, 3, 7, 3}, []uint{}, []uint{2, 6, 3, 3}},
			expected: true,
			msg:      "[0][0]",
		},
		{
			t:        tree{0, []uint{3}, []uint{3, 7, 3}, []uint{}, []uint{5, 5, 3, 5}},
			expected: true,
			msg:      "[1][0]",
		},
		{
			t:        tree{3, []uint{3, 0}, []uint{7, 3}, []uint{}, []uint{5, 3, 5, 3}},
			expected: true,
			msg:      "[2][0]",
		},
		{
			t:        tree{7, []uint{3, 0, 3}, []uint{3}, []uint{}, []uint{1, 3, 4, 9}},
			expected: true,
			msg:      "[3][0]",
		},
		{
			t:        tree{3, []uint{3, 0, 3, 7}, []uint{}, []uint{}, []uint{2, 2, 9, 0}},
			expected: true,
			msg:      "[4][0]",
		},

		// Row 2
		{
			t:        tree{2, []uint{}, []uint{5, 5, 1, 2}, []uint{3}, []uint{6, 3, 3}},
			expected: true,
			msg:      "[0][1]",
		},
		{
			t:        tree{5, []uint{2}, []uint{5, 1, 2}, []uint{0}, []uint{5, 3, 5}},
			expected: true,
			msg:      "[1][1]",
		},
		{
			t:        tree{5, []uint{2, 5}, []uint{1, 2}, []uint{3}, []uint{3, 5, 3}},
			expected: true,
			msg:      "[2][1]",
		},
		{
			t:        tree{1, []uint{2, 5, 5}, []uint{2}, []uint{7}, []uint{3, 4, 9}},
			expected: false,
			msg:      "[3][1]",
		},
		{
			t:        tree{2, []uint{2, 5, 5, 1}, []uint{}, []uint{3}, []uint{2, 9, 0}},
			expected: true,
			msg:      "[4][1]",
		},

		// Row 3
		{
			t:        tree{6, []uint{}, []uint{5, 3, 3, 2}, []uint{3, 2}, []uint{3, 3}},
			expected: true,
			msg:      "[0][2]",
		},
		{
			t:        tree{5, []uint{6}, []uint{3, 3, 2}, []uint{0, 5}, []uint{3, 5}},
			expected: true,
			msg:      "[1][2]",
		},
		{
			t:        tree{3, []uint{6, 5}, []uint{3, 2}, []uint{3, 5}, []uint{5, 3}},
			expected: false,
			msg:      "[2][2]",
		},
		{
			t:        tree{3, []uint{6, 5, 3}, []uint{2}, []uint{7, 1}, []uint{4, 9}},
			expected: true,
			msg:      "[3][2]",
		},
		{
			t:        tree{2, []uint{6, 5, 3, 3}, []uint{}, []uint{3, 2}, []uint{9, 0}},
			expected: true,
			msg:      "[4][2]",
		},

		// Row 4
		{
			t:        tree{3, []uint{}, []uint{3, 5, 4, 9}, []uint{3, 2, 6}, []uint{3}},
			expected: true,
			msg:      "[0][3]",
		},
		{
			t:        tree{3, []uint{3}, []uint{5, 4, 9}, []uint{0, 5, 5}, []uint{5}},
			expected: false,
			msg:      "[1][3]",
		},
		{
			t:        tree{5, []uint{3, 3}, []uint{4, 9}, []uint{3, 5, 3}, []uint{3}},
			expected: true,
			msg:      "[2][3]",
		},
		{
			t:        tree{4, []uint{3, 3, 5}, []uint{9}, []uint{7, 1, 3}, []uint{9}},
			expected: false,
			msg:      "[3][3]",
		},
		{
			t:        tree{9, []uint{3, 3, 5, 4}, []uint{}, []uint{3, 2, 2}, []uint{0}},
			expected: true,
			msg:      "[4][3]",
		},

		// Row 5
		{
			t:        tree{3, []uint{}, []uint{5, 3, 9, 0}, []uint{3, 2, 6, 3}, []uint{}},
			expected: true,
			msg:      "[0][4]",
		},
		{
			t:        tree{5, []uint{3}, []uint{3, 9, 0}, []uint{0, 5, 5, 3}, []uint{}},
			expected: true,
			msg:      "[1][4]",
		},
		{
			t:        tree{3, []uint{3, 5}, []uint{9, 0}, []uint{3, 5, 3, 5}, []uint{}},
			expected: true,
			msg:      "[2][4]",
		},
		{
			t:        tree{9, []uint{3, 5, 3}, []uint{0}, []uint{7, 1, 3, 4}, []uint{}},
			expected: true,
			msg:      "[3][4]",
		},
		{
			t:        tree{0, []uint{3, 5, 3, 9}, []uint{}, []uint{3, 2, 2, 9}, []uint{}},
			expected: true,
			msg:      "[4][4]",
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, test.t.IsVisible())
	}
}

func Test_Tree_ScenicScore(t *testing.T) {
	tests := []struct {
		t        tree
		expected uint
		msg      string
	}{
		// Row 1
		{
			t:        tree{3, []uint{}, []uint{0, 3, 7, 3}, []uint{}, []uint{2, 6, 3, 3}},
			expected: 0,
			msg:      "[0][0]",
		},
		{
			t:        tree{0, []uint{3}, []uint{3, 7, 3}, []uint{}, []uint{5, 5, 3, 5}},
			expected: 0,
			msg:      "[1][0]",
		},
		{
			t:        tree{3, []uint{3, 0}, []uint{7, 3}, []uint{}, []uint{5, 3, 5, 3}},
			expected: 0,
			msg:      "[2][0]",
		},
		{
			t:        tree{7, []uint{3, 0, 3}, []uint{3}, []uint{}, []uint{1, 3, 4, 9}},
			expected: 0,
			msg:      "[3][0]",
		},
		{
			t:        tree{3, []uint{3, 0, 3, 7}, []uint{}, []uint{}, []uint{2, 2, 9, 0}},
			expected: 0,
			msg:      "[4][0]",
		},

		// Row 2
		{
			t:        tree{2, []uint{}, []uint{5, 5, 1, 2}, []uint{3}, []uint{6, 3, 3}},
			expected: 0,
			msg:      "[0][1]",
		},
		{
			t: tree{5, []uint{2}, []uint{5, 1, 2}, []uint{0}, []uint{5, 3, 5}},
			// This value is unfortunately enforced by the puzzle's logic error, where a plot with a 0-height tree
			// is still counted as a tree. But when you see a tree that's not there, do you count it as a tree?
			// Didn't think so!
			expected: 1,
			msg:      "[1][1]",
		},
		{
			t:        tree{5, []uint{2, 5}, []uint{1, 2}, []uint{3}, []uint{3, 5, 3}},
			expected: 4,
			msg:      "[2][1]",
		},
		{
			t:        tree{1, []uint{2, 5, 5}, []uint{2}, []uint{7}, []uint{3, 4, 9}},
			expected: 1,
			msg:      "[3][1]",
		},
		{
			t:        tree{2, []uint{2, 5, 5, 1}, []uint{}, []uint{3}, []uint{2, 9, 0}},
			expected: 0,
			msg:      "[4][1]",
		},

		// Row 3
		{
			t:        tree{6, []uint{}, []uint{5, 3, 3, 2}, []uint{3, 2}, []uint{3, 3}},
			expected: 0,
			msg:      "[0][2]",
		},
		{
			t:        tree{5, []uint{6}, []uint{3, 3, 2}, []uint{0, 5}, []uint{3, 5}},
			expected: 6,
			msg:      "[1][2]",
		},
		{
			t:        tree{3, []uint{6, 5}, []uint{3, 2}, []uint{3, 5}, []uint{5, 3}},
			expected: 1,
			msg:      "[2][2]",
		},
		{
			t:        tree{3, []uint{6, 5, 3}, []uint{2}, []uint{7, 1}, []uint{4, 9}},
			expected: 2,
			msg:      "[3][2]",
		},
		{
			t:        tree{2, []uint{6, 5, 3, 3}, []uint{}, []uint{3, 2}, []uint{9, 0}},
			expected: 0,
			msg:      "[4][2]",
		},

		// Row 4
		{
			t:        tree{3, []uint{}, []uint{3, 5, 4, 9}, []uint{3, 2, 6}, []uint{3}},
			expected: 0,
			msg:      "[0][3]",
		},
		{
			t:        tree{3, []uint{3}, []uint{5, 4, 9}, []uint{0, 5, 5}, []uint{5}},
			expected: 1,
			msg:      "[1][3]",
		},
		{
			t:        tree{5, []uint{3, 3}, []uint{4, 9}, []uint{3, 5, 3}, []uint{3}},
			expected: 8,
			msg:      "[2][3]",
		},
		{
			t:        tree{4, []uint{3, 3, 5}, []uint{9}, []uint{7, 1, 3}, []uint{9}},
			expected: 3,
			msg:      "[3][3]",
		},
		{
			t:        tree{9, []uint{3, 3, 5, 4}, []uint{}, []uint{3, 2, 2}, []uint{0}},
			expected: 0,
			msg:      "[4][3]",
		},

		// Row 5
		{
			t:        tree{3, []uint{}, []uint{5, 3, 9, 0}, []uint{3, 2, 6, 3}, []uint{}},
			expected: 0,
			msg:      "[0][4]",
		},
		{
			t:        tree{5, []uint{3}, []uint{3, 9, 0}, []uint{0, 5, 5, 3}, []uint{}},
			expected: 0,
			msg:      "[1][4]",
		},
		{
			t:        tree{3, []uint{3, 5}, []uint{9, 0}, []uint{3, 5, 3, 5}, []uint{}},
			expected: 0,
			msg:      "[2][4]",
		},
		{
			t:        tree{9, []uint{3, 5, 3}, []uint{0}, []uint{7, 1, 3, 4}, []uint{}},
			expected: 0,
			msg:      "[3][4]",
		},
		{
			t:        tree{0, []uint{3, 5, 3, 9}, []uint{}, []uint{3, 2, 2, 9}, []uint{}},
			expected: 0,
			msg:      "[4][4]",
		},
	}

	for _, test := range tests {
		result := test.t.ScenicScore()
		assert.Equal(t, test.expected, result, test.msg)
	}
}
