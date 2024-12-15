package day10

import (
	"bufio"
	"fmt"
	"os"
)

func Day10(filePath string) ([][]int, error) {
	var lines [][]int

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		intSlice := []int{}
		for _, char := range line {
			intSlice = append(intSlice, int(char-'0'))
		}
		lines = append(lines, intSlice)
	}

	return lines, nil
}

func SearchPart1(arr [][]int) int {
	// Open debug file
	debugFile, err := os.Create("day10/debug.txt")
	if err != nil {
		fmt.Println("Error creating debug file:", err)
		return 0
	}
	defer debugFile.Close()

	count := 0
	trailHead := [][]int{}
	for i, row := range arr {
		for j, val := range row {
			if val == 0 {
				trailHead = append(trailHead, []int{i, j})
			}
		}
	}

	fmt.Fprintf(debugFile, "Found trailheads: %v\n\n", trailHead)
	for _, trail := range trailHead {
		fmt.Fprintf(debugFile, "Starting from trailhead [%d,%d]\n", trail[0], trail[1])
		count += dfsPart2(arr, trail[0], trail[1], 0, map[string]bool{}, debugFile)
	}

	return count
}

func dfs(grid [][]int, row, col int, currentHeight int, visited map[string]bool, debugFile *os.File) int {
	fmt.Fprintf(debugFile, "Visiting: row: %d col: %d\n", row, col)

	if row < 0 || row >= len(grid) || col < 0 || col >= len(grid[0]) ||
		visited[fmt.Sprintf("%d,%d", row, col)] {
		fmt.Fprintf(debugFile, "Rejecting: out of bounds or visited: row: %d col: %d\n", row, col)
		return 0
	}

	// For the first step (currentHeight == 0), we want to be at height 0
	// For all other steps, we want to be at currentHeight
	if (currentHeight == 0 && grid[row][col] != 0) ||
		(currentHeight > 0 && grid[row][col] != currentHeight) {
		fmt.Fprintf(debugFile, "Rejecting: wrong height (expected %d, got %d)\n", currentHeight, grid[row][col])
		return 0
	}

	visited[fmt.Sprintf("%d,%d", row, col)] = true

	if grid[row][col] == 9 {
		fmt.Fprintf(debugFile, "Found height 9!\n")
		return 1
	}

	nextHeight := grid[row][col] + 1
	sum := dfs(grid, row-1, col, nextHeight, visited, debugFile) +
		dfs(grid, row+1, col, nextHeight, visited, debugFile) +
		dfs(grid, row, col-1, nextHeight, visited, debugFile) +
		dfs(grid, row, col+1, nextHeight, visited, debugFile)

	// Only unmark if we didn't find a path to 9
	if sum == 0 {
		delete(visited, fmt.Sprintf("%d,%d", row, col))
	}

	return sum
}

func dfsPart2(grid [][]int, row, col int, currentHeight int, visited map[string]bool, debugFile *os.File) int {
	// Create a new visited map for this path
	newVisited := make(map[string]bool)
	for k, v := range visited {
		newVisited[k] = v
	}

	fmt.Fprintf(debugFile, "Visiting: row: %d col: %d\n", row, col)

	if row < 0 || row >= len(grid) || col < 0 || col >= len(grid[0]) ||
		newVisited[fmt.Sprintf("%d,%d", row, col)] {
		fmt.Fprintf(debugFile, "Rejecting: out of bounds or visited: row: %d col: %d\n", row, col)
		return 0
	}

	if (currentHeight == 0 && grid[row][col] != 0) ||
		(currentHeight > 0 && grid[row][col] != currentHeight) {
		fmt.Fprintf(debugFile, "Rejecting: wrong height (expected %d, got %d)\n", currentHeight, grid[row][col])
		return 0
	}

	newVisited[fmt.Sprintf("%d,%d", row, col)] = true

	if grid[row][col] == 9 {
		fmt.Fprintf(debugFile, "Found height 9!\n")
		return 1
	}

	nextHeight := grid[row][col] + 1
	sum := dfsPart2(grid, row-1, col, nextHeight, newVisited, debugFile) +
		dfsPart2(grid, row+1, col, nextHeight, newVisited, debugFile) +
		dfsPart2(grid, row, col-1, nextHeight, newVisited, debugFile) +
		dfsPart2(grid, row, col+1, nextHeight, newVisited, debugFile)

	return sum
}
