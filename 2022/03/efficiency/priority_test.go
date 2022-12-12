package efficiency

import "testing"

func Test_buildMapping(t *testing.T) {
	m := buildMapping()

	expectedLen := 52
	if len(m) != expectedLen {
		t.Errorf("expected returned mapping to include exactly %d entries, got %d instead", expectedLen, len(m))
	}

	// This set of test expectations come straight from the puzzle.
	testdata := map[rune]uint8{
		'p': 16,
		'L': 38,
		'P': 42,
		'v': 22,
		't': 20,
		's': 19,
	}

	for r, expected := range testdata {
		actual, ok := m[r]
		if !ok {
			t.Errorf("expected rune '%s' to be present in the mapping", r)
		}
		if expected != actual {
			t.Errorf("expected priority assigned to a rune '%s' to be %d, but received %d", r, expected, actual)
		}
	}

	_, ok := m['ś'] // This is an exclusively Polish character and shouldn't exist in the mapping
	if ok == true {
		t.Errorf("expected rune 'ś' to not be present in the mapping")
	}
}

func Test_Priority(t *testing.T) {
	// This set of test expectations come straight from the puzzle.
	testdata := map[rune]uint8{
		'p': 16,
		'L': 38,
		'P': 42,
		'v': 22,
		't': 20,
		's': 19,
	}

	for r, expected := range testdata {
		actual, err := Priority(r)
		if err != nil {
			t.Errorf("expected Priority to succeed, but received error: %s\n", err)
		}
		if expected != actual {
			t.Errorf("expected priority assigned to a rune '%s' to be %d, but received %d", r, expected, actual)
		}
	}

	_, err := Priority('ś')
	if err == nil {
		t.Errorf("expected a call to Priority with rune 'ś' to cause an error, but none was returned")
	}
}
