package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
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
	fmt.Println("scratchcards : ", partTwo(lines))
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

func partTwo(lines [][]rune) int {
	// sum := 0
	mp := make(map[string]int)
	mps := make(map[string][][]string)
	for _, line := range lines {
		sets := strings.Split(string(line), "|")
		re := regexp.MustCompile(`\d+`)
		leftset := re.FindAllString(sets[0], -1)
		cardNum := leftset[0]
		mp[cardNum] = 1
		winningSet := leftset[1:]
		scartched := re.FindAllString(sets[1], -1)
		mps[cardNum] = [][]string{winningSet, scartched}
	}

	updateScratched(mp, mps)

	return countTotalScratchcards(mp)
}

func updateScratched(mp map[string]int, mps map[string][][]string) {
	// The map must be sorted before iterating cuz go map is unordered
	keys := make([]int, 0, len(mps))
	for k := range mps {
		kInt, _ := strconv.Atoi(k)
		keys = append(keys, kInt)
	}
	sort.Ints(keys)

	for _, key := range keys {
		keyStr := strconv.Itoa(key)
		nums := mps[keyStr]
		step := getSteps(nums[0], nums[1])
		if step > 0 {
			for i := 1; i <= step; i++ {
				ky := strconv.Itoa(key + i)
				mp[ky] += mp[keyStr]
			}
		}
	}
}
func getSteps(winningSet []string, scartched []string) int {
	step := 0
	for _, nums := range scartched {
		for _, num := range winningSet {
			if num == nums {
				step++
			}
		}
	}
	return step
}

func countTotalScratchcards(mp map[string]int) int {
	total := 0
	for _, v := range mp {
		total += v
	}
	return total
}
