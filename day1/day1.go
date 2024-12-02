package day1

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func ReadDay1File(filePath string) ([]int, []int, error) {
	s1, s2 := []int{}, []int{}

	if _, err := os.Stat(filePath); err != nil {
		fmt.Printf("File status error: %v\n", err)
		return nil, nil, err
	}

	// Try to read the file
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return nil, nil, err
	}
	defer file.Close()

	// Read and print line by line with additional debugging
	scanner := bufio.NewScanner(file)
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error during scanning: %v\n", err)
	}

	for scanner.Scan() {
		line := scanner.Text()
		x := strings.Split(line, "   ")
		num1, _ := strconv.Atoi(x[0])
		num2, _ := strconv.Atoi(x[1])
		s1, s2 = append(s1, num1), append(s2, num2)
	}

	return s1, s2, nil
}

func DiffScore(s1 []int, s2 []int) int {
	count := 0

	for index, v := range s1 {
		if v > s2[index] {
			count += v - s2[index]
		} else {
			count += s2[index] - v
		}
	}

	return count
}

func SimilarityScore(s1 []int, s2 []int) int {
	simScore := []int{}

	for _, v := range s1 {
		// Use binary search to find the first occurrence of v in s2
		i := sort.SearchInts(s2, v)

		// Count how many times v appears in s2
		count := 0
		for j := i; j < len(s2) && s2[j] == v; j++ {
			count++
		}

		simScore = append(simScore, count*v)
	}

	result := 0
	for _, v := range simScore {
		result += v
	}

	return result
}
