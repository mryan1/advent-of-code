package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
)

func loadInput(fileName string, b *[][5][5]string, cn *[]string) {
	data, err := os.Open(fileName)
	//example of how to write to sice of arrays
	if err != nil {
		log.Println("Error opening file!")
		panic(err)
	}

	scanner := bufio.NewScanner(data)

	// parse first row for numbers
	scanner.Scan()
	(*cn) = strings.Split(scanner.Text(), ",")

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
}

func boardColumn(board [5][5]string, columnIndex int) (column []string) {
	column = make([]string, 0)
	for _, row := range board {
		column = append(column, row[columnIndex])
	}
	return
}

func updateBoards(b *[][5][5]string, cn string) int {
	winner := []string{"*", "*", "*", "*", "*"}
	for i, iv := range *b {
		for j, jv := range iv {

			for k, kv := range jv {
				if kv == cn {
					//check for winner in horizontal direction
					if reflect.DeepEqual((*b)[i][k][0:5], winner) {
						return i
					} else if reflect.DeepEqual(boardColumn((*b)[i], j), winner) {
						//check for winner in verticle direction
						return i
					}
					(*b)[i][j][k] = "*"
				}
			}
		}
	}
	return -1
}
func calculateWinner(wb [5][5]string, wn int) int {
	var sum int
	for _, iv := range wb {
		for _, kv := range iv {
			if kv != "*" {
				kvi, _ := strconv.Atoi(string(kv))
				sum = kvi + sum
			}
		}
	}
	fmt.Println("Sum: ", sum)
	return sum

}
func main() {
	boards := make([][5][5]string, 0)
	var calledNumbers []string
	//load numbers in slice and boards
	loadInput("input.txt", &boards, &calledNumbers)
	//loop though each number
	var p int
	//var cn int
	for i, iv := range calledNumbers {
		//iterate though each board and if it matches replace it with a *
		p = updateBoards(&boards, iv)
		if p != -1 {
			lastCalled, _ := strconv.Atoi(string(calledNumbers[i-1]))
			fmt.Println("Winning value: ", calculateWinner(boards[p], p)*lastCalled)
			break
		}

	}

	//fmt.Println(boards)

	//check each board as you go and see if there are 5 *'s in a row, if so we've found a winner

}
