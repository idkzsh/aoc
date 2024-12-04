package day4

import (
	"bufio"
	"os"
)

func ReadDay4File(filePath string) ([][]rune, error) {
	var lines [][]rune

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		runeSlice := []rune(line)
		lines = append(lines, runeSlice)
	}

	return lines, nil
}

func SearchPart1(arr [][]rune) int {
	count := 0
	word := "XMAS"

	// These directions cover all 8 possible directions
	directions := [][2]int{
		{-1, -1}, {-1, 0}, {-1, 1}, // up-left, up, up-right
		{0, -1}, {0, 1}, // left, right
		{1, -1}, {1, 0}, {1, 1}, // down-left, down, down-right
	}

	// Iterate through each position in the grid
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			// For each position, check all 8 directions
			for _, dir := range directions {
				if checkWord(arr, i, j, dir[0], dir[1], word) {
					count++
				}
			}
		}
	}

	return count
}

// Helper function to check if a word exists starting from a position in a given direction
func checkWord(arr [][]rune, row, col int, dRow, dCol int, word string) bool {
	// Check if the word would extend beyond the grid bounds
	if !canFitWord(arr, row, col, dRow, dCol, len(word)) {
		return false
	}

	// Check each character of the word
	for i, char := range word {
		currentRow := row + (dRow * i)
		currentCol := col + (dCol * i)
		if arr[currentRow][currentCol] != char {
			return false
		}
	}
	return true
}

// Helper function to check if a word of given length can fit from the starting position
func canFitWord(arr [][]rune, row, col int, dRow, dCol int, length int) bool {
	lastRow := row + (dRow * (length - 1))
	lastCol := col + (dCol * (length - 1))

	return lastRow >= 0 && lastRow < len(arr) &&
		lastCol >= 0 && lastCol < len(arr[0])
}

func SearchPart2(arr [][]rune) int {
	count := 0

	// Iterate through each position that could be the center 'A'.
	// No need to check the outer bounds as the center can only be on the inner layers
	for i := 1; i < len(arr)-1; i++ {
		for j := 1; j < len(arr[i])-1; j++ {
			if arr[i][j] == 'A' {
				// Check all four possible 'X' patterns with 'M' and 'S'
				if (arr[i-1][j-1] == 'M' && arr[i+1][j-1] == 'M' &&
					arr[i-1][j+1] == 'S' && arr[i+1][j+1] == 'S') ||
					(arr[i-1][j-1] == 'M' && arr[i-1][j+1] == 'M' &&
						arr[i+1][j-1] == 'S' && arr[i+1][j+1] == 'S') ||
					(arr[i-1][j+1] == 'M' && arr[i+1][j+1] == 'M' &&
						arr[i+1][j-1] == 'S' && arr[i-1][j-1] == 'S') ||
					(arr[i+1][j-1] == 'M' && arr[i+1][j+1] == 'M' &&
						arr[i-1][j-1] == 'S' && arr[i-1][j+1] == 'S') {
					count++
				}
			}
		}
	}

	return count
}
