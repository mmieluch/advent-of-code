package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_filterDirs(t *testing.T) {
	testInput := map[string]uint{
		"foo": 1000,
		"bar": 2000,
		"baz": 3000,
	}
	expected := map[string]uint{
		"bar": 2000,
		"baz": 3000,
	}
	assert.Equal(t, expected, filterDirs(testInput, 2000))
}

func Test_dirSizes(t *testing.T) {
	testInput := []Node{
		&Dir{"a", nil, []Node{
			&Dir{"e", nil, []Node{
				&File{"i", nil, 584},
			}},
			&File{"f", nil, 29116},
			&File{"g", nil, 2557},
			&File{"h.lst", nil, 62596},
		}},
		&File{"b.txt", nil, 14848514},
		&File{"c.dat", nil, 8504156},
		&Dir{"d", nil, []Node{
			&File{"j", nil, 4060174},
			&File{"d.log", nil, 8033020},
			&File{"d.ext", nil, 5626152},
			&File{"d", nil, 7214296},
		}},
	}
	expected := map[string]uint{
		"a": 94853,
		"d": 24933642,
		"e": 584,
	}
	assert.Equal(t, expected, dirSizes(testInput))
}

func Test_smallDirs(t *testing.T) {
	testInput := []Node{
		&Dir{"a", nil, []Node{
			&Dir{"e", nil, []Node{
				&File{"i", nil, 100},
			}},
			&File{"f", nil, 200},
			&File{"g", nil, 300},
			&File{"h.lst", nil, 400},
		}},
		&File{"b.txt", nil, 1},
		&File{"c.dat", nil, 1},
		&Dir{"d", nil, []Node{
			&File{"j", nil, 500},
			&File{"d.log", nil, 600},
			&File{"d.ext", nil, 700},
			&File{"d", nil, 800},
		}},
	}
	expected := []Node{
		&Dir{"a", nil, []Node{
			&Dir{"e", nil, []Node{
				&File{"i", nil, 100},
			}},
			&File{"f", nil, 200},
			&File{"g", nil, 300},
			&File{"h.lst", nil, 400},
		}},
		&Dir{"e", nil, []Node{
			&File{"i", nil, 100},
		}},
	}
	assert.Equal(t, expected, smallDirs(testInput, 1000))
}
