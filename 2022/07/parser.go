package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Op struct {
	Type   string
	Target string
	Output []string
	Raw    string
}

func (op *Op) AddOutputLine(line string) {
	op.Output = append(op.Output, line)
}

type Ops []Op

func (oo *Ops) Add(op Op) {
	*oo = append(*oo, op)
}

func (oo *Ops) Last() *Op {
	ll := len(*oo)
	if ll == 0 {
		return nil
	}

	return &(*oo)[ll-1]
}

func parseInput(input string) (Ops, error) {
	var ops Ops

	for _, line := range strings.Split(input, "\n") {
		err := parseLine(line, &ops)
		if err != nil {
			return Ops{}, err
		}
	}

	return ops, nil
}

type ErrUnknownCommand struct {
	command string
}

func (e *ErrUnknownCommand) Error() string {
	return fmt.Sprintf("unknown command: \"%s\"", e.command)
}

type ErrInvalidCommandSyntax struct {
	command string
}

func (e *ErrInvalidCommandSyntax) Error() string {
	return fmt.Sprintf("invalid command syntax: %s", e.command)
}

type ErrUnexpectedOutput struct {
	line string
}

func (e *ErrUnexpectedOutput) Error() string {
	return fmt.Sprintf("couldn't find a command to attach the output: %s", e.line)
}

func parseLine(line string, oo *Ops) error {
	// Parse command
	if isCommand(line) {
		// `$ cd`
		if isChangeDirCommand(line) {
			target, err := cdTarget(line)
			if err != nil {
				return &ErrInvalidCommandSyntax{strings.Trim(line, "$ ")}
			}
			oo.Add(Op{
				Type:   "cd",
				Target: target,
				Raw:    line,
			})
			return nil
		}
		// `$ ls`
		if isListCommand(line) {
			oo.Add(Op{
				Type: "ls",
				Raw:  line,
			})
			return nil
		}

		fields := strings.Fields(line)
		return &ErrUnknownCommand{fields[1]}
	}

	// If the current line is not a command, then it must be a part of the last
	// command's output. Add it to the operation's output log.
	lastOp := oo.Last()
	if lastOp == nil {
		return &ErrUnexpectedOutput{
			line: line,
		}
	}
	lastOp.AddOutputLine(line)

	return nil
}

func isCommand(line string) bool {
	return strings.HasPrefix(line, "$ ")
}

func isChangeDirCommand(line string) bool {
	return strings.HasPrefix(line, "$ cd ")
}

func isListCommand(line string) bool {
	return strings.HasPrefix(line, "$ ls")
}

var CmdSyntaxError = errors.New("invalid command syntax")

func cdTarget(line string) (string, error) {
	fields := strings.Fields(line)

	if len(fields) < 3 {
		return "", CmdSyntaxError
	}

	return fields[2], nil
}

// parseOperations takes a slice of operations (aliased as the Ops type) and
// builds a tree-like structure of directories and files, based on the ops
// and their respective outputs (if provided).
func parseOperations(oo Ops) (Dir, error) {
	// The very first operation in the input will always be a `cd /`, to
	// determine the root - it's impossible to traverse the tree in a meaningful
	// way without knowing the children of the main node.
	root := Dir{
		name:   "/",
		parent: nil,
	}
	currentNode := &root

	// Next, iterate over all subsequent operations and process them to determine
	// the full shape of the tree.
	for _, op := range oo[1:] {
		switch op.Type {
		// For the `list` command, we need to parse output and create child
		// nodes to allow further traversal.
		case "ls":
			ls(currentNode, op.Output)
		case "cd":
			currentNode = cd(currentNode, op.Target)
		}
	}

	return root, nil
}

func cd(parent *Dir, childName string) *Dir {
	switch childName {
	case "..":
		return parent.Parent().(*Dir)
	default:
		for _, c := range parent.Children() {
			if c.IsDir() && c.Name() == childName {
				return c.(*Dir)
			}
		}
		return &Dir{}
	}
}

func ls(parent Parent, output []string) {
	for _, line := range output {
		fields := strings.Fields(line)
		if fields[0] == "dir" {
			parent.AddChild(&Dir{
				name:   fields[1],
				parent: parent,
			})
			continue
		}

		size, _ := strconv.Atoi(fields[0])
		parent.AddChild(&File{
			name:   fields[1],
			size:   uint(size),
			parent: parent,
		})
	}
}
