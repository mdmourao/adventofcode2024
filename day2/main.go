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
	fmt.Println(reports)
	for _, report := range reports {
		var isValid = true
		var countIsIncreasing = 0
		var countIsDecreasing = 0

		for i := 0; i < len(report)-1; i++ {
			if report[i+1] > report[i] {
				countIsIncreasing++
			}
		}

		for i := 0; i < len(report)-1; i++ {
			if report[i+1] < report[i] {
				countIsDecreasing++
			}
		}

		fmt.Println(report, countIsIncreasing, countIsDecreasing)

		if countIsDecreasing == len(report)-1 || countIsIncreasing == len(report)-1 {
			for i := 0; i < len(report)-1; i++ {
				var diff = Abs(report[i] - report[i+1])
				if diff <= 0 || diff > 3 {
					isValid = false
					break
				}
			}
		} else {
			isValid = false
		}

		fmt.Println(report, isValid)

		if isValid {
			validReportsCount++
		}

	}

	fmt.Println("P1", validReportsCount)

}
