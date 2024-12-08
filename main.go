package main

import (
	"fmt"
	"time"

	"aoc/day8"
)

func main() {
	startTime := time.Now()

	// Example grid
	eq, err := day8.ReadDay8File("day8/puzzle_8.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	count := day8.Day8(eq)
	fmt.Println(count)

	elapsedTime := time.Since(startTime)
	fmt.Printf("Operation took %s\n", elapsedTime)
}
