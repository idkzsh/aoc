package day5

import (
	"bufio"
	"os"
	"slices"
	"strconv"
	"strings"
)

func ReadDay5File(filePath string) (map[int][]int, [][]int, error) {
	order := map[int][]int{}
	var updates [][]int

	file, err := os.Open(filePath)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	parsingMap := true

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			parsingMap = false
			continue
		}

		if parsingMap {
			parts := strings.Split(line, "|")
			num1, _ := strconv.Atoi(strings.TrimSpace(parts[0]))
			num2, _ := strconv.Atoi(strings.TrimSpace(parts[1]))
			order[num1] = append(order[num1], num2)
		} else {
			var lineNumbers []int
			parts := strings.Split(line, ",")
			for _, numStr := range parts {
				num, _ := strconv.Atoi(strings.TrimSpace(numStr))
				lineNumbers = append(lineNumbers, num)
			}
			updates = append(updates, lineNumbers)
		}
	}

	return order, updates, nil
}

func SearchPart1(order map[int][]int, updates [][]int) int {
	count := 0
	for i := 0; i < len(updates); i++ {
		wrong := 0
		for j := 0; j < len(updates[i])-1; j++ {
			current := updates[i][j]
			next := updates[i][j+1]
			if slices.Contains(order[next], current) {
				wrong++
			}
		}

		if wrong == 0 {
			middleIndex := len(updates[i]) / 2
			count += updates[i][middleIndex]
		}
	}

	return count
}

func SearchPart2(order map[int][]int, updates [][]int) int {
	count := 0
	for _, sequence := range updates {
		needsSort := false
		for i := 0; i < len(sequence); i++ {
			for j := i + 1; j < len(sequence); j++ {
				if slices.Contains(order[sequence[j]], sequence[i]) {
					needsSort = true
					break
				}
			}
			if needsSort {
				break
			}
		}

		if needsSort {
			sorted := make([]int, len(sequence))
			copy(sorted, sequence)

			for i := 0; i < len(sorted)-1; i++ {
				for j := 0; j < len(sorted)-i-1; j++ {
					if slices.Contains(order[sorted[j+1]], sorted[j]) {
						sorted[j], sorted[j+1] = sorted[j+1], sorted[j]
					}
				}
			}

			middleIndex := len(sorted) / 2
			count += sorted[middleIndex]
		}
	}
	return count
}
