package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

	for scanner.Scan() {
		sum	+= extractNumber(scanner.Text())			
	}

	fmt.Println(sum)
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

