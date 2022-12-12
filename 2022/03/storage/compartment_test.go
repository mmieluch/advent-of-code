package storage

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_NewCompartment(t *testing.T) {
	input := [][]rune{
		[]rune("vJrwpWtwJgWr"),
		[]rune("hcsFMMfFFhFp"),
		[]rune("jqHRNqRjqzjGDLGL"),
		[]rune("rsFMfFZSrLrFZsSL"),
		[]rune("PmmdzqPrV"),
		[]rune("vPwwTWBwg"),
		[]rune("wMqvLMZHhHMvwLH"),
		[]rune("jbvcjnnSBnvTQFn"),
		[]rune("ttgJtRGJ"),
		[]rune("QctTZtZT"),
		[]rune("CrZsJsPPZsGz"),
		[]rune("wwsLwLmpwMDw"),
	}
	expectedCompartments := []Compartment{
		Compartment("vJrwpWtwJgWr"),
		Compartment("hcsFMMfFFhFp"),
		Compartment("jqHRNqRjqzjGDLGL"),
		Compartment("rsFMfFZSrLrFZsSL"),
		Compartment("PmmdzqPrV"),
		Compartment("vPwwTWBwg"),
		Compartment("wMqvLMZHhHMvwLH"),
		Compartment("jbvcjnnSBnvTQFn"),
		Compartment("ttgJtRGJ"),
		Compartment("QctTZtZT"),
		Compartment("CrZsJsPPZsGz"),
		Compartment("wwsLwLmpwMDw"),
	}

	for i, expected := range expectedCompartments {
		actual := NewCompartment(input[i])
		if reflect.DeepEqual(expected, actual) != true {
			t.Errorf(
				"expected compartment created using set %d (%s) to be %s, received %s instead",
				i,
				fmt.Sprintf("%v", input[i]),
				fmt.Sprintf("%v", expected),
				fmt.Sprintf("%v", actual),
			)
		}
	}
}

func Test_Compartment_Has(t *testing.T) {
	type testItem struct {
		c        Compartment
		testval  rune
		expected bool
	}
	testItems := []testItem{
		{
			c:        Compartment("vJrwpWtwJgWr"),
			testval:  'p',
			expected: true,
		},
		{
			c:        Compartment("vJrwpWtwJgWr"),
			testval:  'Z',
			expected: false,
		},
		{
			c:        Compartment("hcsFMMfFFhFp"),
			testval:  'f',
			expected: true,
		},
		{
			c:        Compartment("hcsFMMfFFhFp"),
			testval:  'q',
			expected: false,
		},
	}

	for _, ti := range testItems {
		actual := ti.c.Has(ti.testval)

		if ti.expected != actual {
			verb := "have"
			if ti.expected == false {
				verb = "not have"
			}

			t.Errorf("expected compartment to %s value %s; compartment: %v", verb, string(ti.testval), ti.c)
		}
	}
}
