package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func readInput() ([]string, []string) {
	// Read input from .txt file]
	data, err := ioutil.ReadFile("inputs/challenge-03-input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
	}

	// Split input by newline
	wireArray := strings.Split(string(data), "\n")

	// The line break separates the first and second wires, so we parse the here
	wireOne := strings.Split(wireArray[0], ",")
	wireTwo := strings.Split(wireArray[1], ",")

	return wireOne, wireTwo
}

func createCoordinateArray(wireCoords []string) [][]int {
	currentPosition := []int{0, 0}
	var coordinateArray = [][]int{}

	for index := 0; index < len(wireCoords)-1; index++ {
		fmt.Println("********* currentPosition", currentPosition)
		currentCoord := wireCoords[index]

		// Parse direction
		direction := string(currentCoord[0])
		fmt.Println("direction", direction)

		// Parse numerical value from string and convert it to an integer
		steps := string(currentCoord[1:len(currentCoord)])
		stepsInt, err := strconv.Atoi(steps)
		fmt.Println("stepsInt", stepsInt)
		if err != nil {
			fmt.Println("Error", err)
		}

		switch direction {
		case "U":
			var stepCoords = currentPosition

			for x := 0; x <= stepsInt; x++ {
				fmt.Println("x", x)
				stepCoords[1] = currentPosition[1] + 1
				fmt.Println("stepCoords", stepCoords)
				coordinateArray = append(coordinateArray, stepCoords)
			}
			currentPosition[1] += stepsInt
		case "D":
			currentPosition[1] -= stepsInt
		case "L":
			currentPosition[0] -= stepsInt
		case "R":
			currentPosition[0] += stepsInt
		}
	}

	return coordinateArray
}

func main() {
	// wireOneArr, wireTwoArr := readInput()

	// TEST CASE #1
	wireOneArr := []string{"R75", "D30", "R83", "U83", "L12", "D49", "R71", "U7", "L72"}
	wireTwoArr := []string{"U62", "R66", "U55", "R34", "D71", "R55", "D58", "R83"}

	// TEST CASE #2
	// wireOneArr := []string{"R98", "U47", "R26", "D63", "R33", "U87", "L62", "D20", "R33", "U53", "R51"}
	// wireTwoArr := []string{"U98", "R91", "D20", "R16", "D67", "R40", "U7", "R15", "U6", "R7"}

	wireOneCoords := createCoordinateArray(wireOneArr)
	fmt.Println("wireOneCoords", wireOneCoords)
	wireTwoCoords := createCoordinateArray(wireTwoArr)
	fmt.Println("wireTwoCoords", wireTwoCoords)

}
