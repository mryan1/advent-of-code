package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func inputToSlice(fileName string) []string {
	data, err := os.Open(fileName)

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

func co2Rating(ones []int, input []string) int {
	fmt.Println("Processing co2 rating")
	var co2Canidate []string
	for i := range ones {
		for j := range input {
			half := int(math.Ceil((float64(len(input)) / 2)))
			if ones[i] >= half && input[j][i] == 48 {
				co2Canidate = append(co2Canidate, input[j])
			} else if ones[i] < half && input[j][i] == 49 {
				co2Canidate = append(co2Canidate, input[j])
			}
			if len(input) == 1 {
				break
			}
		}
		input = co2Canidate
		co2Canidate = []string{}
		ones = countOnes(input)
		if len(input) == 1 {
			break
		}
	}
	co2Dec, _ := strconv.ParseInt(strings.Join(input, ""), 2, 64)
	fmt.Println(input)
	return int(co2Dec)
}

// think I could/should have used recursion

func oxyRating(ones []int, input []string) int {
	fmt.Println("Processing oxy rating")
	var oxyCanidate []string
	for i := range ones {
		for j := range input {
			half := int(math.Ceil((float64(len(input)) / 2)))
			if ones[i] >= half && input[j][i] == 49 {
				oxyCanidate = append(oxyCanidate, input[j])
			} else if ones[i] < half && input[j][i] == 48 {
				oxyCanidate = append(oxyCanidate, input[j])
			}
			if len(input) == 1 {
				break
			}
		}
		input = oxyCanidate
		oxyCanidate = []string{}
		ones = countOnes(input)
		if len(input) == 1 {
			break
		}
	}
	oxyDec, _ := strconv.ParseInt(strings.Join(input, ""), 2, 64)
	fmt.Println(input)
	return int(oxyDec)
}

func main() {
	input := inputToSlice("input.txt")
	numOnes := countOnes(input)

	fmt.Println("Answer is : ", (oxyRating(numOnes, input) * co2Rating(numOnes, input)))

}
