package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

func main() {
	file, err := os.Open("input-day-three.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines [][]rune
	for scanner.Scan() {
		lines = append(lines, []rune(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Sum:", partOne(lines))

}

func partOne(lines [][]rune) int {
	sum := 0
	for i, line := range lines {
		buffer := ""
		found := false
		for j, char := range line {
			if unicode.IsDigit(char) {
				buffer += string(char)
				if !found && isAdjacent(lines, i, j) {
					found = true
				}
			} else {
				if found && buffer != "" {
					num, err := strconv.Atoi(buffer)
					if err != nil {
						log.Fatal(err)
					}
					sum += num
				}
				buffer = ""
				found = false
			}
		}
		if found && buffer != "" {
			num, err := strconv.Atoi(buffer)
			if err != nil {
				log.Fatal(err)
			}
			sum += num
		}
	}
	return sum
}

func isDot(val rune) bool {
	if val == '.' {
		return true
	}
	return false
}

func isAdjacent(lines [][]rune, i int, j int) bool {
	for k := i - 1; k <= i+1; k++ {
		for l := j - 1; l <= j+1; l++ {
			if k >= 0 && k < len(lines) && l >= 0 && l < len(lines[i]) {
				if !isDot(lines[k][l]) && !unicode.IsDigit(lines[k][l]) {
					return true
				}
			}
		}
	}
	return false
}
