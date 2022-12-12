package storage

import (
	"reflect"
	"testing"
)

func Test_NewRucksack(t *testing.T) {
	testdata := []struct {
		items    []rune
		expected Rucksack
	}{
		{
			items: []rune("vJrwpWtwJgWrhcsFMMfFFhFp"),
			expected: Rucksack{
				Left:  Compartment("vJrwpWtwJgWr"),
				Right: Compartment("hcsFMMfFFhFp"),
			},
		},
		{
			items: []rune("jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL"),
			expected: Rucksack{
				Left:  Compartment("jqHRNqRjqzjGDLGL"),
				Right: Compartment("rsFMfFZSrLrFZsSL"),
			},
		},
		{
			items: []rune("PmmdzqPrVvPwwTWBwg"),
			expected: Rucksack{
				Left:  Compartment("PmmdzqPrV"),
				Right: Compartment("vPwwTWBwg"),
			},
		},
		{
			items: []rune("wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn"),
			expected: Rucksack{
				Left:  Compartment("wMqvLMZHhHMvwLH"),
				Right: Compartment("jbvcjnnSBnvTQFn"),
			},
		},
		{
			items: []rune("ttgJtRGJQctTZtZT"),
			expected: Rucksack{
				Left:  Compartment("ttgJtRGJ"),
				Right: Compartment("QctTZtZT"),
			},
		},
		{
			items: []rune("CrZsJsPPZsGzwwsLwLmpwMDw"),
			expected: Rucksack{
				Left:  Compartment("CrZsJsPPZsGz"),
				Right: Compartment("wwsLwLmpwMDw"),
			},
		},
	}

	for _, ti := range testdata {
		actual, err := NewRucksack(ti.items)
		if err != nil {
			t.Errorf("expected a rucksack to be created succesfully using %s, received instead: %s\n", string(ti.items), err)
		}
		if reflect.DeepEqual(ti.expected, actual) != true {
			t.Errorf(
				"expected the created rucksack to be identical to %v, instead received: %v",
				ti.expected,
				actual,
			)
		}
	}
}

func Test_NewRucksack_ErrorWhenOddNumberOfItems(t *testing.T) {
	// The first string from the previous test, but with the last character
	// removed, so now it's 23-chars long. Creating a rucksack should fail.
	items := []rune("vJrwpWtwJgWrhcsFMMfFFhF")
	actual, err := NewRucksack(items)
	if err == nil {
		t.Errorf("expected error to not be nil")
	}

	if reflect.DeepEqual(Rucksack{}, actual) != true {
		t.Errorf(
			"expected the failed call to NewRucksack to return a zero-value rucksack, instead received: %v",
			actual,
		)
	}
}

func Test_Rucksack_CommonItems(t *testing.T) {
	testData := []struct {
		r        Rucksack
		expected []rune
	}{
		{
			r:        Rucksack{Compartment("vJrwpWtwJgWr"), Compartment("hcsFMMfFFhFp")},
			expected: []rune{'p'},
		},
		{
			r:        Rucksack{Compartment("jqHRNqRjqzjGDLGL"), Compartment("rsFMfFZSrLrFZsSL")},
			expected: []rune{'L'},
		},
		{
			r:        Rucksack{Compartment("PmmdzqPrV"), Compartment("vPwwTWBwg")},
			expected: []rune{'P'},
		},
		{
			r:        Rucksack{Compartment("wMqvLMZHhHMvwLH"), Compartment("jbvcjnnSBnvTQFn")},
			expected: []rune{'v'},
		},
		{
			r:        Rucksack{Compartment("ttgJtRGJ"), Compartment("QctTZtZT")},
			expected: []rune{'t'},
		},
		{
			r:        Rucksack{Compartment("CrZsJsPPZsGz"), Compartment("wwsLwLmpwMDw")},
			expected: []rune{'s'},
		},
	}

	for i, ti := range testData {
		actual := ti.r.CommonItems()
		if equals(ti.expected, actual) != true {
			t.Errorf(
				"expected the common items from rucksack %d (%v) to be %v, instead received %v",
				i,
				ti.r,
				ti.expected,
				actual,
			)
		}
	}
}

func equals(a, b []rune) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
