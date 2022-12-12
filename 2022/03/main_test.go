package main

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/mmieluch/advent-of-code/2022/03/storage"
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
