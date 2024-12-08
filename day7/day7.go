package day7

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Equation struct {
	Value     int
	Variables []int
}

func ReadDay7File(filePath string) ([]Equation, error) {
	eq := []Equation{}

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
		val, vars := strings.Split(line, ": ")[0], strings.Split(line, ": ")[1]
		intVal, _ := strconv.Atoi(val)
		intVars := []int{}
		for _, v := range strings.Split(vars, " ") {
			intVar, _ := strconv.Atoi(v)
			intVars = append(intVars, intVar)
		}
		eq = append(eq, Equation{Value: intVal, Variables: intVars})
	}

	return eq, nil
}

func Part1(eq []Equation) int {
	count := 0
	for _, e := range eq {
		if findCombination(e.Value, e.Variables, 0, e.Variables[0]) {
			count += e.Value
		}
	}

	return count
}

func findCombination(target int, nums []int, index int, current int) bool {
	// Base case: if we've used all numbers, check if we've reached the target
	if index == len(nums)-1 {
		return current == target
	}

	// Try both operations with the next number and return true if either works
	concatenated, _ := strconv.Atoi(fmt.Sprintf("%d%d", current, nums[index+1]))
	return findCombination(target, nums, index+1, current+nums[index+1]) ||
		findCombination(target, nums, index+1, current*nums[index+1]) ||
		findCombination(target, nums, index+1, concatenated)
}
