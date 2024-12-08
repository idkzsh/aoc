package main

import (
	"aoc/day7"
	"fmt"
	"time"
)

func main() {
	startTime := time.Now()

	// Example grid
	eq, err := day7.ReadDay7File("day7/puzzle_7.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	count := day7.Part1(eq)
	fmt.Printf("Count: %d\n", count)

	elapsedTime := time.Since(startTime)
	fmt.Printf("Operation took %s\n", elapsedTime)
}
