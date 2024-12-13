package main

import (
	"aoc/day9"
	"fmt"
	"time"
)

func main() {
	startTime := time.Now()

	data, err := day9.ReadDay9File("day9/puzzle_9.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fileblocks := day9.Day9(data)
	fmt.Println(day9.BlockMover2(fileblocks))

	elapsedTime := time.Since(startTime)
	fmt.Printf("Operation took %s\n", elapsedTime)
}
