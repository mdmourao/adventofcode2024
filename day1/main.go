package main

import (
	"bufio"
	"log"
	"os"
	"sort"
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

	var input1 = []int{}
	var input2 = []int{}

	for fileScanner.Scan() {
		line := fileScanner.Text()
		// remove all spaces
		result := strings.Join(strings.Fields(line), ",")
		resultSplit := strings.Split(result, ",")
		// get each value
		result1, err := strconv.Atoi(resultSplit[0])
		if err != nil {
			log.Fatal(err)
		}
		input1 = append(input1, result1)
		result2, err := strconv.Atoi(resultSplit[1])
		if err != nil {
			log.Fatal(err)
		}
		input2 = append(input2, result2)
	}

	// Part 1
	sort.Ints(input1)
	sort.Ints(input2)

	var distances = make([]int, len(input1))
	for i, v := range input1 {
		distances[i] = Abs(v - input2[i])
	}

	var totalDistance = 0
	for _, v := range distances {
		totalDistance += v
	}

	log.Println("part 1:", totalDistance)

	// Part 2
	var similarityScore = 0

	for _, v1 := range input1 {
		var count = 0
		for _, v2 := range input2 {
			if v1 == v2 {
				count++
			}
		}
		similarityScore += v1 * count
	}

	log.Println("part 2:", similarityScore)

}
