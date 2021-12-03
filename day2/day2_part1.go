package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := os.Open("input.txt")

	if err != nil {
		log.Println("Error opening file!")
		panic(err)
	}
	scanner := bufio.NewScanner(data)
	var horizontal int64 = 0
	var depth int64 = 0

	for scanner.Scan() {
		line := scanner.Text()
		commands := strings.Split(line, " ")
		d, _ := strconv.ParseInt(commands[1], 0, 64)

		if commands[0] == "forward" {
			horizontal = horizontal + d

		} else if commands[0] == "up" {
			depth = depth - d
		} else if commands[0] == "down" {
			depth = depth + d
		}

	}
	fmt.Println(horizontal * depth)
}
