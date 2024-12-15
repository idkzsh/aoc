package day11

import (
	"strconv"
)

type Counter map[int]int // map[number]frequency

func Process(counts Counter) Counter {
	newCounts := make(Counter)
	for num, freq := range counts {
		if num == 0 {
			newCounts[1] += freq
		} else if len(strconv.Itoa(num))%2 == 0 {
			// Split number and add frequencies
			numStr := strconv.Itoa(num)
			firstHalf, _ := strconv.Atoi(numStr[:len(numStr)/2])
			secondHalf, _ := strconv.Atoi(numStr[len(numStr)/2:])
			newCounts[firstHalf] += freq
			newCounts[secondHalf] += freq
		} else {
			newCounts[num*2024] += freq
		}
	}
	return newCounts
}

func Part1(data []int) int {
	counts := make(Counter)
	for _, value := range data {
		counts[value]++
	}

	for i := 0; i < 75; i++ {
		counts = Process(counts)
	}

	total := 0
	for _, freq := range counts {
		total += freq
	}
	return total
}
