package operations

import (
	"regexp"
	"strings"
)

type Stack []string

func (s *Stack) GetTopItem() string {
	ll := len(*s)
	if ll == 0 {
		return ""
	}
	return (*s)[ll-1]
}

type Stacks map[int]Stack

func Reorder(ss Stacks, ii []Instruction) (Stacks, error) {
	for _, ins := range ii {
		for i := 0; i < ins.NumOps; i++ {
			item, popped := pop(ss[ins.Source])
			ss[ins.Source] = popped
			ss[ins.Target] = append(ss[ins.Target], item)
		}
	}

	return ss, nil
}

func pop(src Stack) (string, Stack) {
	ll := len(src)
	return src[ll-1], src[:ll-1]
}

func ParseStacks(input string) Stacks {
	drawing := extractDrawing(input)
	return tokeniseDrawing(drawing)
}

func tokeniseDrawing(input string) Stacks {
	ll := strings.Split(input, "\n")
	reverse(ll)

	// Predict number of tokens per line, whether they'd be empty or not.
	// Tokens are represented by 4 characters, expect the last one - which only
	// has 3 (no space after the closing bracket).
	numTokens := len(ll[0])/4 + 1
	ss := make(Stacks, numTokens)

	for _, l := range ll {
		for j := 1; j <= numTokens; j++ {
			off := (j*4 - 2) - 1
			if off <= len(l) {
				token := string(l[off])
				if token != " " {
					ss[j] = append(ss[j], token)
				}
			}
		}
	}

	return ss
}

func noop(a interface{}) {
}

type Token struct {
	raw string
}

func reverse(ss []string) {
	l := len(ss)
	for i := 0; i < l/2; i++ {
		cp := ss[i]
		ss[i] = ss[l-1-i]
		ss[l-1-i] = cp
	}
}

func extractDrawing(input string) string {
	// Remove all lines starting with "move" - these are instructions, not a
	// drawing of stacks.
	re := regexp.MustCompile("(?m)[\r\n]+^move .*$")
	res := re.ReplaceAllString(input, "")
	// Remove all empty lines.
	re = regexp.MustCompile(`(?m)[\r\n]+^\s*$`)
	res = re.ReplaceAllString(res, "")
	// Finally, return the extracted string without the last line containing
	// only stack ordinal numbers.
	lines := strings.Split(res, "\n")
	return strings.Join(lines[:len(lines)-1], "\n")
}
