package storage

type Compartment []rune

func (c Compartment) Has(needle rune) bool {
	for _, r := range c {
		if needle == r {
			return true
		}
	}
	return false
}

func NewCompartment(items []rune) Compartment {
	cmp := make(Compartment, len(items))

	for i, r := range items {
		cmp[i] = r
	}

	return cmp
}
