package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	data, err := os.Open("input.txt")

	if err != nil {
		log.Println("Error opening file!")
		panic(err)
	}
	scanner := bufio.NewScanner(data)
	higherCount := 0
	var lastLine int64
	for i := 0; scanner.Scan(); i++ {
		currentLine, _ := strconv.ParseInt(scanner.Text(), 0, 64)

		if currentLine > lastLine && i != 0 {
			higherCount = higherCount + 1
		}
		lastLine = currentLine
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading file", err)
		}
	}
	fmt.Println(higherCount)

}
