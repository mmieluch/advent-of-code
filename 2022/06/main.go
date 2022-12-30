package main

import (
	"fmt"
	"log"

	"github.com/mmieluch/advent-of-code/2022/internal"
)

func main() {
	input, err := internal.GetInput()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(input)
}
