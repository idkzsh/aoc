package main

import (
	"aoc/day6"
	"fmt"
	"time"
)

func main() {
	startTime := time.Now()

	// Example grid
	area, err := day6.ReadDay6File("day6/puzzle_6.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	count := day6.ObstructPart2(area)
	fmt.Printf("Total loops found: %d\n", count)

	elapsedTime := time.Since(startTime)
	fmt.Printf("Operation took %s\n", elapsedTime)
}
