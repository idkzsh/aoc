package day9

import (
	"fmt"
	"os"
	"strconv"
)

func ReadDay9File(filePath string) ([][]int, error) {
	data := [][]int{}
	content, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}

	for i := 0; i < len(content); i += 2 {
		if i+1 < len(content) {
			data = append(data, []int{int(content[i] - '0'), int(content[i+1] - '0')})
		} else {
			data = append(data, []int{int(content[i] - '0'), 0})
		}
	}

	return data, nil

}

func Day9(data [][]int) [][]string {
	fileblocks := [][]string{}
	for index, row := range data {
		indexArray := []string{}
		for i := 0; i < row[0]; i++ {
			indexArray = append(indexArray, strconv.Itoa(index))
		}
		for i := 0; i < row[1]; i++ {
			indexArray = append(indexArray, ".")
		}
		fileblocks = append(fileblocks, indexArray)
	}

	outputFile, err := os.Create("day9/string_output.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return [][]string{}
	}
	defer outputFile.Close()

	// Write each row of the area to the file
	for _, line := range fileblocks {
		for _, char := range line {
			outputFile.WriteString(char)
		}
		outputFile.WriteString("\n")
	}

	return fileblocks
}

func BlockMover(fileblocks [][]string) int {
	blocks := []string{}

	// For each block from bottom to top
	for i := len(fileblocks) - 1; i >= 0; i-- {
		for j := len(fileblocks[i]) - 1; j >= 0; j-- {
			if fileblocks[i][j] != "." {
				// Find leftmost empty space
				moved := false
				for k := 0; k < len(fileblocks) && !moved; k++ {
					for l := 0; l < len(fileblocks[k]) && !moved; l++ {
						if fileblocks[k][l] == "." {
							fileblocks[k][l] = fileblocks[i][j]
							fileblocks[i][j] = "."
							moved = true // Stop after moving to first available space
						}
					}
				}
			}
		}
	}

	for _, v := range fileblocks {
		for i := 0; i < len(v); i++ {
			if v[i] != "." {
				blocks = append(blocks, v[i])
			}
		}
	}

	result := 0
	for index, v := range blocks {
		num, _ := strconv.Atoi(v)
		result += num * (index)
	}

	return result
}

func BlockMover2(fileblocks [][]string) int {
	// Convert to a single array for easier manipulation
	blocks := []string{}
	for _, row := range fileblocks {
		for _, val := range row {
			blocks = append(blocks, val)
		}
	}

	// Find max file ID
	maxID := -1
	for _, val := range blocks {
		if val != "." {
			id, _ := strconv.Atoi(val)
			if id > maxID {
				maxID = id
			}
		}
	}

	// Process each file ID from highest to lowest
	for fileID := maxID; fileID >= 0; fileID-- {
		idStr := strconv.Itoa(fileID)

		// Count file size
		fileSize := 0
		for _, val := range blocks {
			if val == idStr {
				fileSize++
			}
		}

		if fileSize == 0 {
			continue
		}

		// Find current position of file
		firstPos := -1
		for i, val := range blocks {
			if val == idStr {
				firstPos = i
				break
			}
		}

		// Find leftmost valid position
		bestPos := -1
		for i := 0; i < firstPos; i++ {
			if blocks[i] == "." {
				enough := true
				for j := 0; j < fileSize; j++ {
					if i+j >= len(blocks) || blocks[i+j] != "." {
						enough = false
						break
					}
				}
				if enough {
					bestPos = i
					break
				}
			}
		}

		// Move file if we found a valid position
		if bestPos != -1 {
			// Collect all blocks of this file
			fileBlocks := make([]string, fileSize)
			for i := range fileBlocks {
				fileBlocks[i] = idStr
			}

			// Clear old positions
			for i := range blocks {
				if blocks[i] == idStr {
					blocks[i] = "."
				}
			}

			// Place at new position
			copy(blocks[bestPos:], fileBlocks)
		}
	}

	// Calculate checksum
	result := 0
	for i, v := range blocks {
		if v != "." {
			num, _ := strconv.Atoi(v)
			result += num * i
		}
	}

	return result
}
