package main

import (
	"aoc/day10"
	"fmt"
	"time"
)

func main() {
	startTime := time.Now()

	data, err := day10.Day10("day10/puzzle_10.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	result := day10.SearchPart1(data)
	fmt.Println(result)

	elapsedTime := time.Since(startTime)
	fmt.Printf("Operation took %s\n", elapsedTime)
}
