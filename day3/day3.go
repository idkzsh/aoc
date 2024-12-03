package day3

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func ReadDay3File(filePath string) (string, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("error reading file: %v", err)
	}

	return string(content), nil
}

func ParseDay3FilePart1(content string) [][]string {
	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)

	matches := re.FindAllStringSubmatch(content, -1)

	return matches
}

func ParseDay3FilePart2(content string) [][]string {
	re := regexp.MustCompile(`do\(\)|don't\(\)|mul\((\d{1,3}),(\d{1,3})\)`)

	matches := re.FindAllStringSubmatch(content, -1)

	return matches
}

func Mul(mul [][]string) int {
	sum := 0
	for _, match := range mul {
		if len(match) < 3 {
			continue
		}
		num1, _ := strconv.Atoi(match[1])
		num2, _ := strconv.Atoi(match[2])
		sum += num1 * num2
	}
	return sum
}

func MulPart2(mul [][]string) int {
	disabled := false
	sum := 0
	for _, match := range mul {
		if match[0] == "don't()" {
			disabled = true
		} else if match[0] == "do()" {
			disabled = false
		}

		if strings.Contains(match[0], "mul") && !disabled {
			num1, _ := strconv.Atoi(match[1])
			num2, _ := strconv.Atoi(match[2])
			sum += num1 * num2
		}
	}
	return sum
}
