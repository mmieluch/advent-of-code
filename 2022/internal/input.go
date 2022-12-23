package internal

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

// GetInput will try to load an input file for a specific mode (dev or prod),
// read it, and return as string.
//
// The expectation here is that there will always be two input files: input.dev
// (for hacking purposes) and input.prod (for computing actual values for the
// Advent of Code validator). Since the names only differ in the suffix, providing
// the filename as an argument is a bit pointless. Therefore, if a program is
// invoked without any flags, it will attempt to load the input.dev file by
// default. To load input.prod, invoke the program with the --production flag.
//
// Example:
// # Current working directory is {repoRoot}/2022/01
// $ go run main.go --production
func GetInput() (string, error) {
	productionModePtr := flag.Bool("production", false, "Whether to use production input file")
	flag.Parse()

	ext := "dev"
	if *productionModePtr == true {
		ext = "prod"
	}
	basename := "input"
	filename := fmt.Sprintf("%s.%s", basename, ext)

	cwd, _ := os.Getwd()
	abspath := strings.Join(
		[]string{cwd, filename},
		string(os.PathSeparator),
	)

	return LoadInput(abspath)
}

// LoadInput takes an absolute path to a file and tries to load it as string,
// with white space trimmed from both ends.
func LoadInput(abspath string) (string, error) {
	body, err := os.ReadFile(abspath)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
