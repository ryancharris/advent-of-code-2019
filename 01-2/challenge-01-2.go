package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func getMassesList() []int {
	massArr := []string{}
	massesIntArray := []int{}

	// Read input from .txt file]
	data, err := ioutil.ReadFile("inputs/challenge-1-input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return massesIntArray
	}

	// Split file string at by new lines
	massArr = strings.Split(string(data), "\n")

	// Interate through array of strings and build a new array of ints
	for index := 0; index < len(massArr)-1; index++ {
		intValue, err := strconv.Atoi(massArr[index])
		if err != nil {
			fmt.Println("Integer parsing error", err)
			return []int{}
		}

		massesIntArray = append(massesIntArray, intValue)
	}

	return massesIntArray
}

func main() {
	masses := getMassesList()
	fmt.Println(masses)
}
