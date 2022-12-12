package efficiency

import (
	"errors"
	"fmt"
	"sync"
)

var (
	mapping map[rune]uint8
	once    sync.Once
)

func buildMapping() map[rune]uint8 {
	runes := []rune{
		'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
		'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z',
	}
	m := make(map[rune]uint8, len(runes))

	for i, r := range runes {
		priority := uint8(i + 1)
		m[r] = priority
	}

	return m
}

// Priority returns a priority number assigned to the given rune. If requested
// rune doesn't exist in the mapping it will return a zero-value and an error.
func Priority(r rune) (uint8, error) {
	// The mapping is static and needs to be built only once per program
	// execution.
	once.Do(func() {
		mapping = buildMapping()
	})

	if p, ok := mapping[r]; ok {
		return p, nil
	}

	return uint8(0), errors.New(fmt.Sprintf("unknown value: %s\n", string(r)))
}
