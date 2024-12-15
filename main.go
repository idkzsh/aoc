package main

import (
	"aoc/day11"
	"fmt"
	"time"
)

func main() {
	startTime := time.Now()

	data := []int{2, 72, 8949, 0, 981038, 86311, 246, 7636740}
	// data := []int{125, 17}
	output := day11.Part1(data)
	fmt.Println(output)

	elapsedTime := time.Since(startTime)
	fmt.Printf("Operation took %s\n", elapsedTime)
}
