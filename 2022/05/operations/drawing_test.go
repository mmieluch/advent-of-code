package operations

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_extractDrawing(t *testing.T) {
	input := `    [D]
[N] [C]
[Z] [M] [P]
 1   2   3

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2`

	expected := `    [D]
[N] [C]
[Z] [M] [P]`
	actual := extractDrawing(input)
	assert.Equal(t, expected, actual)
}

func Test_reverse(t *testing.T) {
	input := []string{"a", "b"}
	expected := []string{"b", "a"}
	reverse(input)
	assert.Equal(t, expected, input)

	input = []string{"foo", "bar", "baz"}
	expected = []string{"baz", "bar", "foo"}
	reverse(input)
	assert.Equal(t, expected, input)
}

func Test_tokeniseDrawing(t *testing.T) {
	input := `    [D]
[N] [C]
[Z] [M] [P]`
	expected := Stacks{
		1: {"Z", "N"},
		2: {"M", "C", "D"},
		3: {"P"},
	}
	assert.Equal(t, expected, tokeniseDrawing(input))
}

func Test_pop(t *testing.T) {
	input := Stack{"a", "b", "c"}
	expectedItem := "c"
	expectedStack := Stack{"a", "b"}
	item, popped := pop(input)

	assert.Equal(t, expectedItem, item)
	assert.Equal(t, expectedStack, popped)
}

func Test_ApplyInstructions(t *testing.T) {
	input := Stacks{
		1: {"Z", "N"},
		2: {"M", "C", "D"},
		3: {"P"},
	}
	instructions := []Instruction{
		{NumOps: 1, Source: 2, Target: 1},
		{NumOps: 3, Source: 1, Target: 3},
		{NumOps: 2, Source: 2, Target: 1},
		{NumOps: 1, Source: 1, Target: 2},
	}
	expected := Stacks{
		1: {"C"},
		2: {"M"},
		3: {"P", "D", "N", "Z"},
	}
	updated, err := Reorder(input, instructions)
	assert.Nil(t, err)
	assert.Equal(t, expected, updated)
}

func Test_Stack_GetTopItem(t *testing.T) {
	testData := []struct {
		input    Stack
		expected string
	}{
		{
			input:    Stack{"a", "b", "c"},
			expected: "c",
		},
		{
			input:    Stack{},
			expected: "",
		},
	}

	for _, testitem := range testData {
		assert.Equal(t, testitem.expected, testitem.input.GetTopItem())
	}
}
