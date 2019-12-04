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
	// intCodes := readInput()
	// intCodes := [5]int{1, 0, 0, 0, 99}
	// intCodes := [5]int{2, 3, 0, 3, 99}
	intCodes := [6]int{2, 4, 4, 5, 99, 0}
	// intCodes := [9]int{1, 1, 1, 4, 99, 5, 6, 0, 99}

	fmt.Print("intCodes", intCodes)

	for index := 0; index < len(intCodes); index++ {
		if intCodes[index] == 99 {
			break
		}

		currentCode := intCodes[index]
		indexA := intCodes[index+1]
		indexB := intCodes[index+2]
		positionToUpdate := intCodes[index+3]

		valueA := intCodes[indexA]
		valueB := intCodes[indexB]

		if currentCode == 1 || currentCode == 2 {
			fmt.Println("*************")
			fmt.Println("indexA", indexA)
			fmt.Println("indexB", indexB)
			fmt.Println("valueA", valueA)
			fmt.Println("valueB", valueB)
			fmt.Println("positionToUpdate", positionToUpdate)
			fmt.Println("value at that position", intCodes[positionToUpdate])
		}

		switch currentCode {
		case 1:
			fmt.Println(fmt.Sprintf("%d + %d = %d", valueA, valueB, valueA+valueB))
			intCodes[positionToUpdate] = valueA + valueB
			index += 4
		case 2:
			fmt.Println(fmt.Sprintf("%d * %d = %d", valueA, valueB, valueA*valueB))
			intCodes[positionToUpdate] = valueA * valueB
			index += 4
		}
	}

	fmt.Println(intCodes)
	fmt.Println("=====>", intCodes[0])
}
