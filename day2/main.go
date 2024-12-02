package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func remove(s []int, index int) []int {
	ret := make([]int, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}

func isValid(report []int) bool {
	var countIsIncreasing = 0
	var countIsDecreasing = 0

	for i := 0; i < len(report)-1; i++ {
		if report[i+1] > report[i] {
			countIsIncreasing++
		}

		if report[i+1] < report[i] {
			countIsDecreasing++
		}

		var diff = Abs(report[i] - report[i+1])
		if diff <= 0 || diff > 3 {
			return false
		}
	}

	if countIsDecreasing == len(report)-1 || countIsIncreasing == len(report)-1 {
		return true
	} else {
		return false
	}
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

	reportsRaw := []string{}

	for fileScanner.Scan() {
		line := fileScanner.Text()
		reportsRaw = append(reportsRaw, line)
	}

	reports := make([][]int, len(reportsRaw))

	for i, line := range reportsRaw {
		lineElements := strings.Split(line, " ")
		report := []int{}
		for _, v := range lineElements {

			intValue, err := strconv.Atoi(v)
			if err != nil {
				log.Fatal(err)
			}

			report = append(report, intValue)
		}
		reports[i] = report
	}

	var validReportsCount = 0
	for _, report := range reports {
		if isValid(report) {
			validReportsCount++
		}
	}

	fmt.Println("P1", validReportsCount)

	// Part 2
	validReportsCount = 0
	for _, report := range reports {
		for i := range report {
			var newArrayToEvaluate = remove(report, i)
			if isValid(newArrayToEvaluate) {
				validReportsCount++
				break
			}
		}
	}

	fmt.Println("P2", validReportsCount)
}
