package main

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

type Strategy []Round

func (s Strategy) Total() int {
	total := 0

	for _, r := range s {
		total += r.Score()
	}

	return total
}

func extractTokens(input string) ([][2]string, error) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	re := regexp.MustCompile(`(?P<op>[ABC]) (?P<me>[XYZ])`)
	result := make([][2]string, len(lines))

	for idx, line := range lines {
		ss := re.FindStringSubmatch(line)

		if len(ss) != 3 {
			return [][2]string{}, errors.New(fmt.Sprintf("invalid strategy line: %s", line))
		}

		result[idx] = [2]string{ss[1], ss[2]}
	}

	return result, nil
}

func ParseStrictStrategy(input string) (Strategy, error) {
	tokens, err := extractTokens(input)
	if err != nil {
		return Strategy{}, errors.New("parsing strategy failed: " + err.Error())
	}

	s := make(Strategy, len(tokens))

	for idx, tt := range tokens {
		// Opponent's shape
		ops, err := NewShapeFromAlias(tt[0])
		if err != nil {
			return Strategy{}, errors.New(fmt.Sprintf("invalid shape alias: %s; %s", tt[0], err))
		}
		// My shape
		mys, err := NewShapeFromAlias(tt[1])
		if err != nil {
			return Strategy{}, errors.New(fmt.Sprintf("invalid shape alias: %s; %s", tt[1], err))
		}

		s[idx] = Round{
			Ord: idx + 1,
			Op:  ops,
			Me:  mys,
		}
	}

	return s, nil
}

func ParseFlexibleStrategy(input string) (Strategy, error) {
	tokens, err := extractTokens(input)
	if err != nil {
		return Strategy{}, errors.New("parsing strategy failed: " + err.Error())
	}

	s := make(Strategy, len(tokens))

	for idx, tt := range tokens {
		// Opponent's shape
		ops, err := NewShapeFromAlias(tt[0])
		if err != nil {
			return Strategy{}, errors.New(fmt.Sprintf("invalid shape alias: %s; %s", tt[0], err))
		}
		// My shape - which should be created based on the token's suggested
		// result outcome. So we're not just blindly creating a shape from the
		// alias, we need to determine the target shape first, and then create
		// it.
		mys, err := NewShapeForOutcome(ops, tt[1])
		if err != nil {
			return Strategy{}, errors.New(fmt.Sprintf("invalid shape alias: %s; %s", tt[1], err))
		}

		s[idx] = Round{
			Ord: idx + 1,
			Op:  ops,
			Me:  mys,
		}
	}

	return s, nil
}
