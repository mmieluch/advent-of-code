package internal

import (
	"fmt"
	"strings"
)

func PrintPartHeading(message string) {
	hBorder := strings.Repeat("#", len(message)+4)

	fmt.Println(hBorder)
	fmt.Printf("# %s #\n", message)
	fmt.Println(hBorder)
}
