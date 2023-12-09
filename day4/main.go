package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {
	file, err := os.Open("input-day-four.txt")
	if err != nil {
		panic(err)
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

	fmt.Println("sum: ", partOne(lines))
}

func partOne(lines [][]rune) int {
	sum := 0
	for _, line := range lines {
		sets := strings.Split(string(line), "|")
		re := regexp.MustCompile(`\d+`)
		leftset := re.FindAllString(sets[0], -1)
		winningSet := leftset[1:]
		scartched := re.FindAllString(sets[1], -1)
		sum += getScore(winningSet, scartched)
	}
	return sum
}

func getScore(winningSet []string, scartched []string) int {
	score := 0
	for _, nums := range scartched {
		for _, num := range winningSet {
			if num == nums {
				if score == 0 {
					score = 1
				} else {
					score *= 2
				}
			}
		}
	}
	return score
}
