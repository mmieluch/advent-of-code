package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SOPMarkerIdx(t *testing.T) {
	testData := []struct {
		input            []byte
		expectedVal      uint
		wasErrorExpected bool
	}{
		{[]byte("mjqjpqmgbljsphdztnvjfqwrcgsmlb"), 7, false},
		{[]byte("bvwbjplbgvbhsrlpgdmjqwftvncz"), 5, false},
		{[]byte("nppdvjthqldpwncqszvftbrmjlhg"), 6, false},
		{[]byte("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg"), 10, false},
		{[]byte("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"), 11, false},
		{[]byte("abcabcabcabcabcabcabcabcabcaabbccaabbcc"), 0, true},
	}

	for _, testItem := range testData {
		res, err := SOPMarkerIdx(testItem.input)
		assert.Equal(t, testItem.expectedVal, res)
		if testItem.wasErrorExpected {
			assert.Error(t, err)
		}
	}
}

func Test_MsgMarkerIdx(t *testing.T) {
	testData := []struct {
		input            []byte
		expectedVal      uint
		wasErrorExpected bool
	}{
		{[]byte("mjqjpqmgbljsphdztnvjfqwrcgsmlb"), 19, false},
		{[]byte("bvwbjplbgvbhsrlpgdmjqwftvncz"), 23, false},
		{[]byte("nppdvjthqldpwncqszvftbrmjlhg"), 23, false},
		{[]byte("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg"), 29, false},
		{[]byte("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"), 26, false},
		{[]byte("abcabcabcabcabcabcabcabcabcaabbccaabbcc"), 0, true},
		{[]byte("abcd"), 0, true},
	}

	for _, testItem := range testData {
		res, err := MsgMarkerIdx(testItem.input)
		assert.Equal(t, testItem.expectedVal, res)
		if testItem.wasErrorExpected {
			assert.Error(t, err)
		}
	}
}
