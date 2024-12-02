package day2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func ReadDay2File(filePath string) ([]string, error) {
	s1 := []string{}

	if _, err := os.Stat(filePath); err != nil {
		fmt.Printf("File status error: %v\n", err)
		return nil, err
	}

	// Try to read the file
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return nil, err
	}
	defer file.Close()

	// Read and print line by line with additional debugging
	scanner := bufio.NewScanner(file)
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error during scanning: %v\n", err)
	}

	for scanner.Scan() {
		line := scanner.Text()
		s1 = append(s1, line)
	}

	return s1, nil
}

func isAscending(nums []int) (bool, int) {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] >= nums[i+1] {
			return false, i
		}
	}
	return true, 0
}

func isDescending(nums []int) (bool, int) {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] <= nums[i+1] {
			return false, i
		}
	}
	return true, 0
}

func SafeChecker(s1 []string) bool {

	// Convert strings to ints
	nums := make([]int, len(s1))
	for i, s := range s1 {
		n, _ := strconv.Atoi(s)
		nums[i] = n
	}

	// Check both ascending and descending first
	isAsc, ascIdx := isAscending(nums)
	isDsc, dscIdx := isDescending(nums)

	// If either pattern is valid, check the differences
	if isAsc {
		for i := 0; i < len(nums)-1; i++ {
			if nums[i+1]-nums[i] > 3 {
				return Dampener(nums, i)
			}
		}
		return true
	} else if isDsc {
		for i := 0; i < len(nums)-1; i++ {
			if nums[i]-nums[i+1] > 3 {
				return Dampener(nums, i)
			}
		}
		return true
	}

	// If we get here, try removing one number to fix the sequence
	return Dampener(nums, ascIdx) || Dampener(nums, dscIdx)
}

func Dampener(s1 []int, index int) bool {
	if index < 0 || index >= len(s1) {
		return false
	}

	// Create both possible arrays by removing either number around the violation
	arr1 := append(append([]int{}, s1[:index+1]...), s1[index+2:]...)
	arr2 := append(append([]int{}, s1[:index]...), s1[index+1:]...)

	// Check both arrays for either ascending or descending patterns
	return (checkSequence(arr1) || checkSequence(arr2))
}

// helper function to reduce code duplication
func checkSequence(arr []int) bool {
	isAsc, _ := isAscending(arr)
	isDsc, _ := isDescending(arr)

	if isAsc {
		for i := 0; i < len(arr)-1; i++ {
			if arr[i+1]-arr[i] > 3 {
				return false
			}
		}
		return true
	} else if isDsc {
		for i := 0; i < len(arr)-1; i++ {
			if arr[i]-arr[i+1] > 3 {
				return false
			}
		}
		return true
	}
	return false
}
