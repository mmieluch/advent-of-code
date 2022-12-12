package storage

import "errors"

type Rucksack struct {
	Left  Compartment
	Right Compartment
}

type intersection map[rune]interface{}

func (is intersection) ToSlice() []rune {
	if len(is) == 0 {
		return []rune{}
	}

	var out []rune
	for r, _ := range is {
		out = append(out, r)
	}
	return out
}

func (rs Rucksack) CommonItems() []rune {
	is := make(intersection)

	// Extract intersection to a map, to eliminate duplicates.
	for _, r := range rs.Right {
		if rs.Left.Has(r) {
			// Add a rune to the intersection if it doesn't already exist.
			if _, ok := is[r]; !ok {
				is[r] = nil
			}
		}
	}

	return is.ToSlice()
}

func (rs Rucksack) AllItems() []rune {
	var out []rune
	out = append(out, rs.Left...)
	out = append(out, rs.Right...)
	return out
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

func SharedByAll(rr ...Rucksack) []rune {
	if len(rr) < 2 {
		// Nothing to compare to.
		return []rune{}
	}
	// Determine the first intersection
	common := runeSliceIntersection(rr[0].AllItems(), rr[1].AllItems())

	// No matches? Return immediately, nothing more to do here.
	if len(common) == 0 {
		return []rune{}
	}

	for _, r := range rr[2:] {
		common = runeSliceIntersection(common, r.AllItems())
		if len(common) == 0 {
			return []rune{}
		}
	}

	return common
}

func runeSliceIntersection(rr1, rr2 []rune) []rune {
	is := make(intersection)

	for _, r := range rr2 {
		if sliceHasRune(r, rr1) {
			if _, ok := is[r]; !ok {
				is[r] = nil
			}
		}
	}

	return is.ToSlice()

}

func removeRune(needle rune, haystack []rune) []rune {
	var out []rune
	for _, r := range haystack {
		if r != needle {
			out = append(out, r)
		}
	}
	return out
}

func sliceHasRune(needle rune, haystack []rune) bool {
	for _, r := range haystack {
		if r == needle {
			return true
		}
	}
	return false
}
