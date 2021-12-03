package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}

func main() {

	data, err := os.Open("input.txt")

	if err != nil {
		log.Println("Error opening file!")
		panic(err)
	}
	scanner := bufio.NewScanner(data)
	var readingList []int
	for scanner.Scan() {
		line, _ := strconv.ParseInt(scanner.Text(), 0, 64)
		readingList = append(readingList, int(line))

		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading file", err)
		}
	}
	//now iterate through slice, calculate sliding window and cound increases
	var higherCounter int = 0
	var lastWindowSum int = 0
	var currentWindowSum int = 0
	for t := range readingList {
		p := t + 3
		currentWindowSum = sum(readingList[t:p])
		if currentWindowSum > lastWindowSum && t != 0 {
			higherCounter = higherCounter + 1
		}
		lastWindowSum = currentWindowSum
	}
	fmt.Println(higherCounter)

}
