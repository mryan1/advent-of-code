package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func loadInput(fileName string, b *[][5][5]string) {
	data, err := os.Open(fileName)
	//example of how to write to sice of arrays
	if err != nil {
		log.Println("Error opening file!")
		panic(err)
	}

	scanner := bufio.NewScanner(data)

	// parse first row for numbers
	scanner.Scan()
	calledNumbers := strings.Split(scanner.Text(), ",")
	fmt.Println(calledNumbers)

	//parse the rest to get the boards
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			//if current line is blank, get next 5 and enter into a 5x5 slice
			newBoard := [5][5]string{}
			for i := 0; i < 5; i++ {
				scanner.Scan()
				line = scanner.Text()
				for j, k := range strings.Fields(line) {
					newBoard[i][j] = k
				}
			}
			(*b) = append((*b), newBoard)
		}

		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading file", err)
		}
	}

	//board := [5][5]int{{1, 2, 3, 4, 5}, {6, 7, 8, 9, 10}, {11, 12, 13, 14, 15}, {16, 17, 18, 19, 20}, {21, 22, 23, 24, 25}}
	//boards = append(boards, board)
	//board = [5][5]int{{26, 27, 28, 29, 30}, {31, 32, 33, 34, 35}, {36, 37, 38, 39, 40}, {41, 42, 43, 44, 45}, {46, 47, 48, 49, 50}}
	//boards = append(boards, board)
	//fmt.Println(boards)

}

func main() {
	boards := make([][5][5]string, 0)

	//load numbers in slice and boards
	loadInput("input.txt", &boards)
	fmt.Println(boards)

	//loop though each number

	//iterate though each board and if it matches replace it with a *

	//check each board as you go and see if there are 5 *'s in a row, if so we've found a winner

}
