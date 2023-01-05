package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Op_AddOutputLine(t *testing.T) {
	op1 := Op{}
	op1.AddOutputLine("foo")

	assert.Len(t, op1.Output, 1)
	assert.Equal(t, "foo", op1.Output[0])

	op2 := Op{
		Output: []string{"foo", "bar"},
	}
	op2.AddOutputLine("baz")

	assert.Len(t, op2.Output, 3)
	assert.Equal(t, "foo", op2.Output[0])
	assert.Equal(t, "bar", op2.Output[1])
	assert.Equal(t, "baz", op2.Output[2])
}

func Test_Ops_Add(t *testing.T) {
	oo := Ops{}
	oo.Add(Op{Type: "foo"})

	assert.Len(t, oo, 1)
	assert.Equal(t, Ops{
		Op{Type: "foo"},
	}, oo)

	oo = Ops{
		Op{Type: "one"},
		Op{Type: "two"},
		Op{Type: "three"},
	}
	oo.Add(Op{Type: "four"})
	assert.Len(t, oo, 4)
	assert.Equal(
		t,
		Ops{
			Op{Type: "one"},
			Op{Type: "two"},
			Op{Type: "three"},
			Op{Type: "four"},
		},
		oo,
	)
}

func Test_Ops_Last(t *testing.T) {
	oo := Ops{
		Op{Type: "one"},
		Op{Type: "two"},
		Op{Type: "three"},
	}
	assert.Equal(t, &Op{Type: "three"}, oo.Last())

	oo = Ops{}
	assert.Nil(t, oo.Last())
}

func Test_parseInput(t *testing.T) {
	input := `$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k`
	expected := Ops{
		Op{Type: "cd", Target: "/", Raw: "$ cd /"},
		Op{Type: "ls", Raw: "$ ls", Output: []string{
			"dir a",
			"14848514 b.txt",
			"8504156 c.dat",
			"dir d",
		}},
		Op{Type: "cd", Target: "a", Raw: "$ cd a"},
		Op{Type: "ls", Raw: "$ ls", Output: []string{
			"dir e",
			"29116 f",
			"2557 g",
			"62596 h.lst",
		}},
		Op{Type: "cd", Target: "e", Raw: "$ cd e"},
		Op{Type: "ls", Raw: "$ ls", Output: []string{
			"584 i",
		}},
		Op{Type: "cd", Target: "..", Raw: "$ cd .."},
		Op{Type: "cd", Target: "..", Raw: "$ cd .."},
		Op{Type: "cd", Target: "d", Raw: "$ cd d"},
		Op{Type: "ls", Raw: "$ ls", Output: []string{
			"4060174 j",
			"8033020 d.log",
			"5626152 d.ext",
			"7214296 k",
		}},
	}
	result, err := parseInput(input)
	assert.Equal(t, expected, result)
	assert.Nil(t, err)
}

func Test_parseInput_errors(t *testing.T) {
	tests := map[string]struct {
		input       string
		expectedErr error
	}{
		// Make sure parsing error is returned when syntax of the `cd` command
		// line is invalid. In the input below the command is missing the
		// required directory name argument.
		"invalid cd command syntax": {
			input:       "$ cd ",
			expectedErr: &ErrInvalidCommandSyntax{},
		},
		// Make sure parsing error is returned when an unknown command is
		// provided in the input.
		"unknown command": {
			input:       "$ foobar",
			expectedErr: &ErrUnknownCommand{},
		},
		// Parsing input should fail if there's a non-command line present in
		// the input that is not preceded with a command. This way, there's
		// clearly something wrong, as it means there's an output line that we
		// cannot match to an operation.
		"loose output line": {
			input:       "123 some_file.txt\n$ cd /\n$ ls",
			expectedErr: &ErrUnexpectedOutput{},
		},
	}

	for name, test := range tests {
		result, err := parseInput(test.input)
		assert.Empty(t, result, name)
		assert.IsType(t, test.expectedErr, err, name)
		assert.NotEmpty(t, err.Error())
	}
}

func Test_isCommand(t *testing.T) {
	testData := []struct {
		input    string
		expected bool
	}{
		{"$ cd /", true},
		{"$ ls", true},
		{"$ cd a", true},
		{"dir a", false},
		{"14848514 b.txt", false},
	}

	for _, item := range testData {
		assert.Equal(t, item.expected, isCommand(item.input), item.input)
	}
}

func Test_isChangeDirCommand(t *testing.T) {
	testData := []struct {
		input    string
		expected bool
	}{
		{"$ cd /", true},
		{"$ ls", false},
		{"$ cd a", true},
		{"dir a", false},
		{"14848514 b.txt", false},
	}

	for _, item := range testData {
		assert.Equal(t, item.expected, isChangeDirCommand(item.input), item.input)
	}
}

func Test_isListCommand(t *testing.T) {
	testData := []struct {
		input    string
		expected bool
	}{
		{"$ cd /", false},
		{"$ ls", true},
		{"$ cd a", false},
		{"dir a", false},
		{"14848514 b.txt", false},
	}

	for _, item := range testData {
		assert.Equal(t, item.expected, isListCommand(item.input), item.input)
	}
}

func Test_cdTarget(t *testing.T) {
	testData := []struct {
		input          string
		expected       string
		isInvalidInput bool
	}{
		{"$ cd /", "/", false},
		{"$ cd ..", "..", false},
		{"$ cd d", "d", false},
		{"foobar", "", true},
	}

	for _, testItem := range testData {
		res, err := cdTarget(testItem.input)
		assert.Equal(t, testItem.expected, res)
		if testItem.isInvalidInput {
			assert.Error(t, err)
		} else {
			assert.Nil(t, err)
		}
	}
}

func Test_parseOperations(t *testing.T) {
	input := Ops{
		{Type: "cd", Target: "/"},
		{Type: "ls", Output: []string{
			"dir a",
			"14848514 b.txt",
			"8504156 c.dat",
			"dir d",
		}},
		{Type: "cd", Target: "a"},
		{Type: "ls", Output: []string{
			"dir e",
			"29116 f",
			"2557 g",
			"62596 h.lst",
		}},
		{Type: "cd", Target: "e"},
		{Type: "ls", Output: []string{
			"584 i",
		}},
		{Type: "cd", Target: ".."},
		{Type: "cd", Target: ".."},
		{Type: "cd", Target: "d"},
		{Type: "ls", Output: []string{
			"4060174 j",
			"8033020 d.log",
			"5626152 d.ext",
			"7214296 k",
		}},
	}

	// /
	root := Dir{name: "/", children: make([]Node, 4)}
	// /a
	root.children[0] = &Dir{name: "a", parent: &root, children: make([]Node, 4)}
	// /a/e
	root.children[0].(*Dir).children[0] = &Dir{name: "e", parent: root.children[0].(*Dir), children: make([]Node, 1)}
	// /a/e/i
	root.children[0].(*Dir).children[0].(*Dir).children[0] = &File{
		name:   "i",
		size:   584,
		parent: root.children[0].(*Dir).children[0].(*Dir),
	}
	// /a/f
	root.children[0].(*Dir).children[1] = &File{name: "f", size: 29116, parent: root.children[0].(*Dir)}
	// /a/g
	root.children[0].(*Dir).children[2] = &File{name: "g", size: 2557, parent: root.children[0].(*Dir)}
	// /a/h.lst
	root.children[0].(*Dir).children[3] = &File{name: "h.lst", size: 62596, parent: root.children[0].(*Dir)}
	// /b.txt
	root.children[1] = &File{name: "b.txt", size: 14848514, parent: &root}
	// /c.dat
	root.children[2] = &File{name: "c.dat", size: 8504156, parent: &root}
	// /d
	root.children[3] = &Dir{name: "d", parent: &root, children: make([]Node, 4)}
	// /d/j
	root.children[3].(*Dir).children[0] = &File{name: "j", size: 4060174, parent: root.children[3].(*Dir)}
	// /d/d.log
	root.children[3].(*Dir).children[1] = &File{name: "d.log", size: 8033020, parent: root.children[3].(*Dir)}
	// /d/d.ext
	root.children[3].(*Dir).children[2] = &File{name: "d.ext", size: 5626152, parent: root.children[3].(*Dir)}
	// /d/k
	root.children[3].(*Dir).children[3] = &File{name: "k", size: 7214296, parent: root.children[3].(*Dir)}

	res, err := parseOperations(input)

	assert.Nil(t, err)
	assert.Equal(t, root, res)
}
