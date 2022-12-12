package storage

import "errors"

type Rucksack struct {
	Left  Compartment
	Right Compartment
}

func (rs Rucksack) CommonItems() []rune {
	intersection := make(map[rune]interface{})

	// Extract intersection to a map, to eliminate duplicates.
	for _, r := range rs.Right {
		if rs.Left.Has(r) {
			// Add a rune to the intersection if it doesn't already exist.
			if _, ok := intersection[r]; !ok {
				intersection[r] = nil
			}
		}
	}

	// Return a slice of runes by putting all intersection map keys in a slice.
	var is []rune
	for r, _ := range intersection {
		is = append(is, r)
	}

	return is
}

func NewRucksack(items []rune) (Rucksack, error) {
	count := len(items)
	if count%2 > 0 {
		return Rucksack{}, errors.New("the number of items to put in a rucksack must be even")
	}

	half := count / 2
	return Rucksack{
		Left:  NewCompartment(items[:half]),
		Right: NewCompartment(items[half:]),
	}, nil
}
