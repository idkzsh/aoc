package main

import (
	"aoc/day5"
	"fmt"
	"time"
)

func main() {
	startTime := time.Now()
	// day 3
	filePath := "day5/puzzle_5.txt"
	order, updates, err := day5.ReadDay5File(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// fmt.Println(day5.SearchPart1(order, updates))
	fmt.Println(day5.SearchPart2(order, updates))

	elapsedTime := time.Since(startTime)
	fmt.Printf("Operation took %s\n", elapsedTime)
}
