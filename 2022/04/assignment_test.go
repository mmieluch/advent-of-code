package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Assignment_ContainsSection(t *testing.T) {
	testdata := []struct {
		a        Assignment
		needle   uint
		expected bool
	}{
		{Assignment{2, 3, 4}, 1, false},
		{Assignment{2, 3, 4}, 2, true},
		{Assignment{2, 3, 4}, 3, true},
		{Assignment{2, 3, 4}, 4, true},
		{Assignment{2, 3, 4}, 5, false},
	}

	for _, ti := range testdata {
		assert.Equal(t, ti.expected, ti.a.ContainsSection(ti.needle))
	}
}

func Test_Assignment_FullyContains(t *testing.T) {
	testdata := []struct {
		base       Assignment
		comparison Assignment
		expected   bool
	}{
		{Assignment{2, 3, 4}, Assignment{6, 7, 8}, false},
		{Assignment{2, 3}, Assignment{4, 5}, false},
		{Assignment{5, 6, 7}, Assignment{7, 8, 9}, false},
		{Assignment{2, 3, 4, 5, 6, 7, 8}, Assignment{3, 4, 5, 6, 7}, true},
		{Assignment{6}, Assignment{4, 5, 6}, false},
		{Assignment{4, 5, 6}, Assignment{6}, true},
		{Assignment{2, 3, 4, 5, 6}, Assignment{4, 5, 6, 7, 8}, false},
	}

	for _, testitem := range testdata {
		assert.Equal(t, testitem.expected, testitem.base.FullyContains(testitem.comparison))
	}
}

func Test_Assignment_Overlaps(t *testing.T) {
	testdata := []struct {
		base       Assignment
		comparison Assignment
		expected   bool
	}{
		{Assignment{2, 3, 4}, Assignment{6, 7, 8}, false},
		{Assignment{2, 3}, Assignment{4, 5}, false},
		{Assignment{5, 6, 7}, Assignment{7, 8, 9}, true},
		{Assignment{2, 3, 4, 5, 6, 7, 8}, Assignment{3, 4, 5, 6, 7}, true},
		{Assignment{6}, Assignment{4, 5, 6}, true},
		{Assignment{4, 5, 6}, Assignment{6}, true},
		{Assignment{2, 3, 4, 5, 6}, Assignment{4, 5, 6, 7, 8}, true},
	}

	for _, testitem := range testdata {
		assert.Equal(t, testitem.expected, testitem.base.Overlaps(testitem.comparison))
	}
}

func Test_NewAssignment(t *testing.T) {
	testdata := []struct {
		min      uint
		max      uint
		expected Assignment
	}{
		{2, 4, Assignment{2, 3, 4}},
		{6, 8, Assignment{6, 7, 8}},
		{2, 3, Assignment{2, 3}},
		{4, 5, Assignment{4, 5}},
		{5, 7, Assignment{5, 6, 7}},
		{7, 9, Assignment{7, 8, 9}},
		{2, 8, Assignment{2, 3, 4, 5, 6, 7, 8}},
		{3, 7, Assignment{3, 4, 5, 6, 7}},
		{6, 6, Assignment{6}},
		{4, 6, Assignment{4, 5, 6}},
		{2, 6, Assignment{2, 3, 4, 5, 6}},
		{4, 8, Assignment{4, 5, 6, 7, 8}},
	}

	for _, item := range testdata {
		assert.Equal(t, item.expected, NewAssignment(item.min, item.max))
	}
}

func Test_ParsePlan(t *testing.T) {
	input := `2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8`
	expected := []Pair{
		{Assignment{2, 3, 4}, Assignment{6, 7, 8}},
		{Assignment{2, 3}, Assignment{4, 5}},
		{Assignment{5, 6, 7}, Assignment{7, 8, 9}},
		{Assignment{2, 3, 4, 5, 6, 7, 8}, Assignment{3, 4, 5, 6, 7}},
		{Assignment{6}, Assignment{4, 5, 6}},
		{Assignment{2, 3, 4, 5, 6}, Assignment{4, 5, 6, 7, 8}},
	}

	assert.Equal(t, expected, ParsePlan(input))
}

func Test_extractLimits(t *testing.T) {
	testdata := []struct {
		input       string
		expectedMin uint
		expectedMax uint
	}{
		{"2-4", 2, 4},
		{"6-8", 6, 8},
		{"2-3", 2, 3},
		{"4-5", 4, 5},
		{"5-7", 5, 7},
		{"7-9", 7, 9},
		{"2-8", 2, 8},
		{"3-7", 3, 7},
		{"6-6", 6, 6},
		{"4-6", 4, 6},
		{"2-6", 2, 6},
		{"4-8", 4, 8},
		{"0-10", 0, 10},
	}

	for _, testitem := range testdata {
		min, max := extractLimits(testitem.input)
		assert.Equal(t, testitem.expectedMin, min)
		assert.Equal(t, testitem.expectedMax, max)
	}
}
