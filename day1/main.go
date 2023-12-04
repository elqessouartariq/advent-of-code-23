package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

func main() {
	
	file, err := os.Open("input-day-one.txt")
	
	if err != nil {
			log.Fatal(err)
	}
	defer file.Close()


	scanner := bufio.NewScanner(file)
	
	if err := scanner.Err(); err != nil {
			log.Fatal(err)
	}
	var sum int

	mapOfDigits := map[string]string{
		"zero":"ze0ro",
		"one":"on1e",
		"two":"tw2o",
		"three":"th3ree",
		"four":"fo4ur",
		"five":"fi5ve",
		"six":"si6x",
		"seven":"se7ven",
		"eight":"ei8ght",
		"nine":"ni9ne",
	}

	for scanner.Scan() {
		sum	+= extractNumber(replaceDigits(scanner.Text(), mapOfDigits))
	}

	fmt.Println(sum)
}

func replaceDigits(line string, mp map[string]string) (string){
	for key, value := range mp {
		line = strings.Replace(line, key, value, -1)
	}
	return line
}


func extractNumber(line string) (int){
	var firstNumber int
	var secondNumber int

	found := 0
	
	for _, char := range line {
		if unicode.IsDigit(char) {
			if	found > 0 {
				secondNumber = int(char) - 48
			}else {
				firstNumber = int(char) - 48
				secondNumber = firstNumber
				found++
			}
			}
	}
	number := firstNumber * 10 + secondNumber
	return number
}

