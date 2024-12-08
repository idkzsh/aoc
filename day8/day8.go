package day8

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type Position struct {
	x, y int
}

func ReadDay8File(filePath string) ([][]rune, error) {
	data := [][]rune{}

	if _, err := os.Stat(filePath); err != nil {
		fmt.Printf("File status error: %v\n", err)
		return nil, err
	}

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error during scanning: %v\n", err)
	}

	for scanner.Scan() {
		line := scanner.Text()
		runeArray := []rune(line)
		data = append(data, runeArray)
	}
	return data, nil
}

func Day8(data [][]rune) int {
	antennas := make(map[rune][]Position)
	uniqueAntinodes := make(map[Position]bool)

	// First collect all antennas by their frequency
	for y, line := range data {
		for x, char := range line {
			if char != '.' {
				antennas[char] = append(antennas[char], Position{x, y})
			}
		}
	}

	// For each frequency that has multiple antennas
	for _, positions := range antennas {
		if len(positions) < 2 {
			continue // Skip frequencies with only one antenna
		}

		// Check each pair of antennas only once
		for i := 0; i < len(positions); i++ {
			for j := i + 1; j < len(positions); j++ {
				pos1 := positions[i]
				pos2 := positions[j]

				// Calculate vector from pos1 to pos2
				dx := pos2.x - pos1.x
				dy := pos2.y - pos1.y

				// Calculate GCD to get the smallest step size
				gcd := GCD(int(math.Abs(float64(dx))), int(math.Abs(float64(dy))))
				if gcd == 0 {
					continue
				}

				// Get the unit vector
				stepX := dx / gcd
				stepY := dy / gcd

				// Start from pos1 and go backwards until we hit a boundary
				currentPos := Position{pos1.x, pos1.y}
				for currentPos.x >= 0 && currentPos.y >= 0 &&
					currentPos.x < len(data[0]) && currentPos.y < len(data) {
					uniqueAntinodes[currentPos] = true
					currentPos.x -= stepX
					currentPos.y -= stepY
				}

				// Start from pos1 and go forwards until we hit a boundary
				currentPos = Position{pos1.x, pos1.y}
				for currentPos.x >= 0 && currentPos.y >= 0 &&
					currentPos.x < len(data[0]) && currentPos.y < len(data) {
					uniqueAntinodes[currentPos] = true
					currentPos.x += stepX
					currentPos.y += stepY
				}
			}
		}
	}

	return len(uniqueAntinodes)
}

func GCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
