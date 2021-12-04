package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func inputToSlice(fileName string) []string {
	data, err := os.Open("input.txt")

	if err != nil {
		log.Println("Error opening file!")
		panic(err)
	}
	scanner := bufio.NewScanner(data)
	var readingList []string
	for scanner.Scan() {
		line := scanner.Text()
		readingList = append(readingList, line)

		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading file", err)
		}
	}
	return readingList
}
func countOnes(input []string) []int {
	//find most common bit for each position.  Least common is opposite.
	var one = make([]int, len(input[0]))
	for i := range input {
		for j := range input[i] {
			if string(input[i][j]) == "1" {
				one[j] = one[j] + 1
			}
		}
	}
	return one
}
func calGammaRate(counts []int, total int) int {
	var gamma = make([]string, len(counts))
	for i := range counts {
		if counts[i] > total/2 {
			gamma[i] = "1"
		} else {
			gamma[i] = "0"
		}
	}
	gammaDec, _ := strconv.ParseInt(strings.Join(gamma, ""), 2, 64)
	return int(gammaDec)
}

func calEpsilonRate(counts []int, total int) int {
	var epsilon = make([]string, len(counts))
	for i := range counts {
		if counts[i] < total/2 {
			epsilon[i] = "1"
		} else {
			epsilon[i] = "0"
		}
	}
	epsilonDec, _ := strconv.ParseInt(strings.Join(epsilon, ""), 2, 64)
	return int(epsilonDec)
}

func main() {
	input := inputToSlice("input.txt")
	numOnes := countOnes(input)
	gamaRate := calGammaRate(numOnes, len(input))
	epsilonRate := calEpsilonRate(numOnes, len(input))
	fmt.Println("Gama rate: ", gamaRate)
	fmt.Println("Epsilon rate: ", epsilonRate)
	fmt.Println("Power consumpiton: ", (gamaRate * epsilonRate))
}
