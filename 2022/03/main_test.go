package main

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/mmieluch/advent-of-code/2022/03/storage"
	"github.com/stretchr/testify/assert"
)

func Test_parseManifest(t *testing.T) {
	input := `vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw`
	m := parseManifest(input)

	if len(m) != 6 {
		t.Errorf("expected 6 items in the result slice, got %d", len(m))
	}

	testdata := [][]rune{
		[]rune("vJrwpWtwJgWrhcsFMMfFFhFp"),
		[]rune("jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL"),
		[]rune("PmmdzqPrVvPwwTWBwg"),
		[]rune("wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn"),
		[]rune("ttgJtRGJQctTZtZT"),
		[]rune("CrZsJsPPZsGzwwsLwLmpwMDw"),
	}

	for i, expected := range testdata {
		actual := m[i]
		if reflect.DeepEqual(expected, actual) != true {
			t.Errorf(
				"expected manifest item %d to be %s, got %s instead",
				i,
				fmt.Sprintf("%v", expected),
				fmt.Sprintf("%v", actual),
			)
		}
	}
}

func Test_pack(t *testing.T) {
	input := [][]rune{
		[]rune("vJrwpWtwJgWrhcsFMMfFFhFp"),
		[]rune("jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL"),
		[]rune("PmmdzqPrVvPwwTWBwg"),
		[]rune("wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn"),
		[]rune("ttgJtRGJQctTZtZT"),
		[]rune("CrZsJsPPZsGzwwsLwLmpwMDw"),
	}
	expected := []storage.Rucksack{
		storage.Rucksack{
			Left:  storage.Compartment("vJrwpWtwJgWr"),
			Right: storage.Compartment("hcsFMMfFFhFp"),
		},
		storage.Rucksack{
			Left:  storage.Compartment("jqHRNqRjqzjGDLGL"),
			Right: storage.Compartment("rsFMfFZSrLrFZsSL"),
		},
		storage.Rucksack{
			Left:  storage.Compartment("PmmdzqPrV"),
			Right: storage.Compartment("vPwwTWBwg"),
		},
		storage.Rucksack{
			Left:  storage.Compartment("wMqvLMZHhHMvwLH"),
			Right: storage.Compartment("jbvcjnnSBnvTQFn"),
		},
		storage.Rucksack{
			Left:  storage.Compartment("ttgJtRGJ"),
			Right: storage.Compartment("QctTZtZT"),
		},
		storage.Rucksack{
			Left:  storage.Compartment("CrZsJsPPZsGz"),
			Right: storage.Compartment("wwsLwLmpwMDw"),
		},
	}

	rr, err := pack(input)
	if err != nil {
		t.Errorf("expected pack to succeed, got error instead: %s\n", err)
	}
	for i, r := range expected {
		if reflect.DeepEqual(r, rr[i]) != true {
			t.Errorf(
				"expected rucksack %d to be %s, got %s\n",
				i,
				fmt.Sprintf("%v", r),
				fmt.Sprintf("%v", rr[i]),
			)
		}
	}
}

func Test_asTriplets(t *testing.T) {
	input := []storage.Rucksack{
		{storage.Compartment("abc"), storage.Compartment("def")},
		{storage.Compartment("ghi"), storage.Compartment("jkl")},
		{storage.Compartment("mno"), storage.Compartment("pqr")},
		{storage.Compartment("stu"), storage.Compartment("vwx")},
		{storage.Compartment("yzA"), storage.Compartment("BCD")},
		{storage.Compartment("EFG"), storage.Compartment("HIJ")},
	}

	result, err := asTriplets(input)
	assert.Nil(t, err)
	assert.Len(t, result, 2)

	assert.Equal(t, input[0], result[0][0])
	assert.Equal(t, input[1], result[0][1])
	assert.Equal(t, input[2], result[0][2])
	assert.Equal(t, input[3], result[1][0])
	assert.Equal(t, input[4], result[1][1])
	assert.Equal(t, input[5], result[1][2])
}

func Test_asTriplets_errorWhenIncorrectNumberOfRucksacks(t *testing.T) {
	var input []storage.Rucksack
	for i := 0; i < 5; i++ {
		input = append(input, storage.Rucksack{})
	}

	result, err := asTriplets(input)
	assert.Error(t, err, "asTriplets should fail when given a slice of rucksacks that is not a multiple of 3")

	if len(result) > 0 {
		t.Errorf("expected the result to be empty")
	}
}
