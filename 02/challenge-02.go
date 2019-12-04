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

	for index := 0; index < len(intCodes); index++ {
		if intCodes[index] == 99 {
			break
		}

		codeArray := intCodes[index : index+4]
		currentCode := codeArray[0]
		indexA := codeArray[1]
		indexB := codeArray[2]
		positionToUpdate := codeArray[3]

		valueA := intCodes[indexA]
		fmt.Println("valueA", valueA)
		valueB := intCodes[indexB]
		fmt.Println("valueB", valueB)

		switch currentCode {
		case 1:
			fmt.Println("codeArray", codeArray)
			fmt.Println(fmt.Sprintf("%d + %d = %d", valueA, valueB, valueA+valueB))
			intCodes[positionToUpdate] = valueA + valueB
			index += 4
		case 2:
			fmt.Println("codeArray", codeArray)
			fmt.Println(fmt.Sprintf("%d * %d = %d", valueA, valueB, valueA*valueB))
			intCodes[positionToUpdate] = valueA * valueB
			index += 4
		}
	}
	fmt.Println(intCodes)
	fmt.Println("=====>", intCodes[0])
}
