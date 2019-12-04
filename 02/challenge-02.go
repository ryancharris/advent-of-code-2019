package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func readInput() []int {
	// Read input from .txt file]
	data, err := ioutil.ReadFile("inputs/challenge-02-input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
	}

	// Split file string at by new lines
	var opcodeInts = []int{}
	opcode := strings.Split(string(data), ",")

	// Create new array of ints
	for index := 0; index < len(opcode)-1; index++ {
		codeInt, err := strconv.Atoi(opcode[index])
		if err != nil {
			fmt.Println("Error", err)
		}

		opcodeInts = append(opcodeInts, codeInt)

	}

	return opcodeInts
}

func main() {
	intCodes := readInput()

	for index := 0; index < len(intCodes); index += 4 {
		if intCodes[index] == 99 {
			break
		}

		currentCode := intCodes[index]
		indexA := intCodes[index+1]
		indexB := intCodes[index+2]
		positionToUpdate := intCodes[index+3]

		valueA := intCodes[indexA]
		valueB := intCodes[indexB]

		switch currentCode {
		case 1:
			intCodes[positionToUpdate] = valueA + valueB
		case 2:
			intCodes[positionToUpdate] = valueA * valueB
		}
	}

	fmt.Println(intCodes)
	fmt.Println("=====>", intCodes[0])
}
