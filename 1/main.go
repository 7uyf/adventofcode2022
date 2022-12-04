package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const (
	fileName = "input.txt"
)

func main() {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var currentSum int = 0
	var biggestSum int = 0
	for scanner.Scan() {
		if scanner.Text() == "" {
			if biggestSum < currentSum {
				biggestSum = currentSum
			}
			currentSum = 0
		} else {
			calorie, err := strconv.Atoi(scanner.Text())
			if err != nil {
				panic(err)
			}

			currentSum += calorie
		}

	}
	fmt.Println(biggestSum)
}
