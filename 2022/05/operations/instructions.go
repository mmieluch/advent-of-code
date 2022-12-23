package operations

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Instruction struct {
	NumOps int
	Source int
	Target int
}

func NewInstructionFromRegexpMatches(numOps string, source string, target string) (Instruction, error) {
	numOpsI, err := strconv.Atoi(numOps)
	if err != nil {
		return Instruction{}, err
	}

	sourceI, err := strconv.Atoi(source)
	if err != nil {
		return Instruction{}, err
	}

	targetI, err := strconv.Atoi(target)
	if err != nil {
		return Instruction{}, err
	}

	return Instruction{
		NumOps: numOpsI,
		Source: sourceI,
		Target: targetI,
	}, nil
}

func ParseInstructions(input string) ([]Instruction, error) {
	// Extract all lines that include operating instructions.
	// Each instruction line has the following format:
	// move [numContainers} from {stackX} to {stackY}

	ll := strings.Split(input, "\n")
	var iis []string
	for _, l := range ll {
		if strings.HasPrefix(l, "move ") {
			iis = append(iis, l)
		}
	}

	var ii []Instruction
	re := regexp.MustCompile(`move (?P<numOps>\d+) from (?P<source>\d+) to (?P<target>\d+)`)

	for _, l := range iis {
		matches := re.FindStringSubmatch(l)
		if len(matches) != 4 {
			return []Instruction{}, errors.New(fmt.Sprintf("couldn't parse instructions line: %s", l))
		}

		ins, err := NewInstructionFromRegexpMatches(matches[1], matches[2], matches[3])
		if err != nil {
			return []Instruction{}, errors.New(fmt.Sprintf("couldn't create instructions from matches: %v", matches[1:]))
		}

		ii = append(ii, ins)
	}

	return ii, nil
}
