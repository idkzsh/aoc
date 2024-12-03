package main

import (
	"aoc/day3"
	"fmt"
	"time"
)

func main() {
	startTime := time.Now()
	// day 1

	// filePath := "puzzle_2.txt"

	// // read file into int arrays
	// s1, s2, err := day1.ReadDay1File(filePath)
	// if err != nil {
	// 	fmt.Println("Error reading file:", err)
	// 	return
	// }

	// // sort the arrays with quicksort (o(n log n))
	// sort.Ints(s1)
	// sort.Ints(s2)

	// // part 1 o(n)
	// diffScore := day1.DiffScore(s1, s2)
	// fmt.Println(diffScore)

	// // part 2 o(n log n)
	// simScore := day1.SimilarityScore(s1, s2)
	// fmt.Println(simScore)

	// elapsedTime := time.Since(startTime)
	// fmt.Printf("Operation took %s\n", elapsedTime)

	// day 2
	// filePath := "day2/puzzle_2.txt"

	// lines, err := day2.ReadDay2File(filePath)
	// if err != nil {
	// 	fmt.Println("Error reading file:", err)
	// 	return
	// }

	// count := 0

	// for _, line := range lines {
	// 	nums := strings.Split(line, " ")
	// 	fmt.Printf("Processing line: %v\n", nums)
	// 	safe := day2.SafeChecker(nums)
	// 	if safe {
	// 		count++
	// 	}
	// }

	// fmt.Printf("Total safe sequences: %d\n", count)

	// day 3
	filePath := "day3/puzzle_3.txt"
	lines, err := day3.ReadDay3File(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// mul := day3.ParseDay3FilePart1(lines)
	mul := day3.ParseDay3FilePart2(lines)

	resultsPart2 := day3.MulPart2(mul)

	fmt.Println(resultsPart2)

	elapsedTime := time.Since(startTime)
	fmt.Printf("Operation took %s\n", elapsedTime)
}
