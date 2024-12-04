package main

import (
	"aoc/day4"
	"fmt"
	"time"
)

func main() {
	startTime := time.Now()
	// day 3
	filePath := "day4/puzzle_4.txt"
	lines, err := day4.ReadDay4File(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// fmt.Println(day4.SearchPart1(lines))
	fmt.Println(day4.SearchPart2(lines))
	// day4.Array2d()

	elapsedTime := time.Since(startTime)
	fmt.Printf("Operation took %s\n", elapsedTime)
}
