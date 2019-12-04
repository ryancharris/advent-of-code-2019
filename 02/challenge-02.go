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
		currentCode := intCodes[index]
		if currentCode == 99 {
			break
		}

		codeArray := intCodes[index : index+4]
		positionToUpdate := codeArray[3]

		switch currentCode {
		case 1:
			fmt.Println("codeArray", codeArray)
			fmt.Println("positionToUpdate", positionToUpdate)
			fmt.Println(fmt.Sprintf("%d + %d = %d", codeArray[1], codeArray[2], codeArray[1]+codeArray[2]))
			fmt.Println("**********")
			intCodes[positionToUpdate] = codeArray[1] + codeArray[2]
			index += 4
		case 2:
			fmt.Println("codeArray", codeArray)
			fmt.Println("positionToUpdate", positionToUpdate)
			fmt.Println(fmt.Sprintf("%d * %d = %d", codeArray[1], codeArray[2], codeArray[1]*codeArray[2]))
			fmt.Println("**********")
			intCodes[positionToUpdate] = codeArray[1] * codeArray[2]
			index += 4
		}
	}

	fmt.Println("intCodes", intCodes)
	fmt.Println("intCodes[0]", intCodes[0])
}
