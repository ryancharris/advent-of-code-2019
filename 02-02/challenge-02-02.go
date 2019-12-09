package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func readInput() []int {
	// Read input from .txt file]
	data, err := ioutil.ReadFile("inputs/challenge-02-02-input.txt")
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

func processIntcode(intCodes []int) int {
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
	return intCodes[0]

}

func main() {
	var valueA int
	var valueB int

	for i := 0; i <= 99; i++ {
		for j := 0; j <= 99; j++ {
			intCodes := readInput()
			intCodes[1] = i
			intCodes[2] = j

			fmt.Println("i", i)
			fmt.Println("j", j)

			output := processIntcode(intCodes)

			if output == 19690720 {
				valueA = i
				valueB = j
				fmt.Println("BREAKING")
				i = 99
				j = 99

			}
		}
	}

	fmt.Println("valueA")
	fmt.Println(valueA)
	fmt.Println("valueB")
	fmt.Println(valueB)

	fmt.Println("======> ")
	fmt.Println(100*valueA + valueB)

}
