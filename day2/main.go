package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input-day-two.txt")
	if err != nil {
		log.Fatal(err)
	}
	
	defer file.Close()
	
	var sum int = 0

	colors_map := map[string]int {
		"red": 0,
		"green": 0,
		"blue": 0,
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(getGameData(scanner.Text(), colors_map))
		sum += getGameData(scanner.Text(), colors_map)
	}

	if err := scanner.Err(); err != nil {
			log.Fatal(err)
	}

	fmt.Println(sum)
}

func isGamePossible(colors_map map[string]int) bool {
	for key, value := range colors_map {
		if (key == "red" && value > 12) || (key == "green" && value > 13) || (key == "blue" && value > 14 ){
			return false
		}
	}
	return true
}

func getGameData(line string, colors_map map[string]int) int {
	id := 0

	re := regexp.MustCompile(`Game (\d+): (.*)`)
	matches := re.FindStringSubmatch(line)
	
	if len(matches) > 0 {
			id, _ = strconv.Atoi(matches[1])
			sets := strings.Split(matches[2], "; ")   

			reRed := regexp.MustCompile(`(\d+) red`)
			reGreen := regexp.MustCompile(`(\d+) green`)
			reBlue := regexp.MustCompile(`(\d+) blue`)

			for _, set := range sets {
					colors_map["red"] = 0
					colors_map["green"] = 0
					colors_map["blue"] = 0

					redMatches := reRed.FindStringSubmatch(set)
					greenMatches := reGreen.FindStringSubmatch(set)
					blueMatches := reBlue.FindStringSubmatch(set)

					if len(redMatches) > 0 {
							red, _ := strconv.Atoi(redMatches[1])   
							colors_map["red"] += red
					}

					if len(greenMatches) > 0 {
							green, _ := strconv.Atoi(greenMatches[1]) 
							colors_map["green"] += green
					}

					if len(blueMatches) > 0 {
							blue, _ := strconv.Atoi(blueMatches[1])  
							colors_map["blue"] += blue
					}

					if !isGamePossible(colors_map) {
							return 0
					}
			}
	}

	return id
}
