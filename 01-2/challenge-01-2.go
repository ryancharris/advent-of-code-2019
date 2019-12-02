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

	massArr = strings.Split(string(data), "\n")

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
