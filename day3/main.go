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
	fmt.Println("Product:", partTwo(lines))
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
func partTwo(lines [][]rune) int {
	type Position struct {
		x int
		y int
	}

	prod := 0
	pairs := make(map[Position][]int)
	for i, line := range lines {
		buffer := ""
		found := false
		var pos Position
		for j, char := range line {
			if unicode.IsDigit(char) {
				buffer += string(char)
				isAdj, l, k := isAdjacent2(lines, i, j)
				if !found && isAdj {
					found = true
					pos = Position{l, k}
				}
			} else {
				if found && buffer != "" {
					num, err := strconv.Atoi(buffer)
					if err != nil {
						log.Fatal(err)
					}
					pairs[pos] = append(pairs[pos], num)
					found = false
				}
				buffer = ""
			}
		}
		if found && buffer != "" {
			num, err := strconv.Atoi(buffer)
			if err != nil {
				log.Fatal(err)
			}
			pairs[pos] = append(pairs[pos], num)
			found = false
		}
	}

	for key, nums := range pairs {
		fmt.Println(key, nums)
		if len(nums) == 2 {
			prod += nums[0] * nums[1]
			fmt.Println(prod)
		}
	}

	return prod
}
func isDot(val rune) bool {
	return val == '.'
}
func isAdjacent2(lines [][]rune, i int, j int) (bool, int, int) {
	directions := [][]int{
		{-1, 0},
		{-1, -1},
		{-1, 1},
		{0, -1},
		{0, 1},
		{1, 1},
		{1, 0},
		{1, -1},
	}

	for _, dir := range directions {
		ni, nj := i+dir[0], j+dir[1]
		if ni < 0 || nj < 0 || ni >= len(lines) || nj >= len(lines[ni]) {
			continue
		} else {
			if lines[ni][nj] == '*' {
				return true, ni, nj
			}
		}
	}
	return false, 0, 0
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
