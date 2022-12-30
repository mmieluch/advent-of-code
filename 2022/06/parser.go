package main

import (
	"errors"
	"fmt"
)

const sopMarkerLen = 4

const msgMarkerLen = 14

func findIndex(input []byte, seqlen uint) (uint, error) {
	if seqlen > uint(len(input)) {
		return uint(0), fmt.Errorf("input too short: \"%s\"", string(input))
	}

	start := uint(0)
	end := start + seqlen
	lastidx := uint(len(input) - 1)

	for {
		buf := input[start:end]
		if isUniqueSequence(buf, seqlen) {
			return end, nil
		}

		start++
		end++
		if start+seqlen > lastidx {
			break
		}
	}

	return uint(0), errors.New("index not found")
}

// SOPMarkerIdx returns a uint representing the index of the last character
// identifying the correct start-of-packet marker in the given input.
// A start-of-packet (SOP) marker consists of 4 unique bytes.
func SOPMarkerIdx(input []byte) (uint, error) {
	idx, err := findIndex(input, sopMarkerLen)
	if err != nil {
		return uint(0), fmt.Errorf("input didn't contain a start-of-packet marker: \"%s\"", string(input))
	}

	return idx, nil
}

func MsgMarkerIdx(input []byte) (uint, error) {
	idx, err := findIndex(input, msgMarkerLen)
	if err != nil {
		return uint(0), fmt.Errorf("input didn't contain a start-of-packet marker: \"%s\"", string(input))
	}

	return idx, nil
}

// isUniqueSequence returns true when all bytes in the input slice are unique,
// so none of the bytes is repeated within the slice. Returns false otherwise.
func isUniqueSequence(input []byte, seqlen uint) bool {
	mm := make(map[byte]interface{})
	for _, b := range input {
		mm[b] = nil
	}

	return uint(len(mm)) == seqlen
}
