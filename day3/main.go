package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

// find mul(1,1), do(), don't()
var regex = `(?:mul\(\d+,\d+\)|do\(\)|don't\(\))`

func getResult(str string) int {
	left, right := 0, 0
	fmt.Sscanf(str, "mul(%d,%d)", &left, &right)
	return left * right
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: go run main.go <input_file>")
	}
	fileName := os.Args[1]

	readFile, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	r := regexp.MustCompile(regex)

	isFirst := true
	allMatches := []string{}

	for fileScanner.Scan() {
		allMatches = append(allMatches, r.FindAllString(fileScanner.Text(), -1)...)
	}

	var sum = 0
	for i, match := range allMatches {
		if isFirst {
			if strings.Contains(match, "mul") {
				sum += getResult(match)
			} else {
				isFirst = false
			}
		} else {
			if !strings.Contains(match, "mul") {
				continue
			}
			var iterator = i
			for !strings.Contains(allMatches[iterator], "do") {
				iterator--
			}
			if allMatches[iterator] == "do()" {
				sum += getResult(match)
			}
		}
	}

	fmt.Println("Sum:", sum)
}
