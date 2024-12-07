package day6

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Direction int

type MapDetails struct {
	column    int
	direction Direction
}

const (
	North Direction = iota
	East  Direction = iota
	South Direction = iota
	West  Direction = iota

	WALL    = 35 // '#'
	EMPTY   = 46 // '.'
	VISITED = 42 // '*'
)

func ReadDay6File(filePath string) ([][]rune, error) {
	var area [][]rune

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		var lineNumbers []rune
		parts := strings.Split(line, "")
		for _, numStr := range parts {
			lineNumbers = append(lineNumbers, rune(numStr[0]))
		}
		area = append(area, lineNumbers)

	}

	return area, nil
}

func TraversePart1(area [][]rune) bool {

	visitedMap := make(map[int][]MapDetails)
	count := 1
	pos := []int{55, 86} // 55, 86
	dir := North
	oob := false
	loop := false

	for !oob && !loop {
		// Check if we've visited this position before
		if _, exists := visitedMap[pos[0]]; exists {
			for _, detail := range visitedMap[pos[0]] {
				if detail.column == pos[1] && detail.direction == dir {
					count++
					loop = true
					break
				}
			}
		}

		// Only increment count if this is a new position (not already marked)
		// if area[pos[0]][pos[1]] == EMPTY {
		// 	count++
		// }
		// Mark current position as visited
		area[pos[0]][pos[1]] = VISITED
		visitedMap[pos[0]] = append(visitedMap[pos[0]], struct {
			column    int
			direction Direction
		}{
			column:    pos[1],
			direction: dir,
		})

		switch dir {
		case North:
			if pos[0]-1 < 0 {
				oob = true
				break
			}
			next := area[pos[0]-1][pos[1]]
			if next == WALL { // '#'
				dir = East
			} else if next == EMPTY || next == VISITED {
				pos[0]--
			} else {
				oob = true
			}
		case East:
			if pos[1]+1 >= len(area[0]) {
				oob = true
				break
			}
			next := area[pos[0]][pos[1]+1]
			if next == WALL {
				dir = South
			} else if next == EMPTY || next == VISITED {
				pos[1]++
			} else {
				oob = true
			}
		case South:
			if pos[0]+1 >= len(area) {
				oob = true
				break
			}
			next := area[pos[0]+1][pos[1]]
			if next == WALL {
				dir = West
			} else if next == EMPTY || next == VISITED {
				pos[0]++
			} else {
				oob = true
			}
		case West:
			if pos[1]-1 < 0 {
				oob = true
				break
			}
			next := area[pos[0]][pos[1]-1]
			if next == WALL {
				dir = North
			} else if next == EMPTY || next == VISITED {
				pos[1]--
			} else {
				oob = true
			}
		}

	}

	// Write the result to a file to visualize the path
	outputFile, err := os.Create("path_output.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return false
	}
	defer outputFile.Close()

	// Write each row of the area to the file
	for _, row := range area {
		outputFile.WriteString(string(row) + "\n")
	}

	return loop
}

func ObstructPart2(area [][]rune) int {
	count := 0
	for i := 0; i < len(area); i++ {
		for j := 0; j < len(area[i]); j++ {
			if area[i][j] != WALL {
				past := area[i][j]
				area[i][j] = WALL
				if TraversePart1(area) {
					count++
				}
				area[i][j] = past
			}
		}
	}
	return count
}
