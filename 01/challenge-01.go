package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func calculateFuel(mass int) float64 {
	return math.Floor(float64(mass/3)) - 2
}

func main() {
	// Read input from .txt file]
	data, err := ioutil.ReadFile("01/challenge-1-input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	var requiredFuel float64
	masses := strings.Split(string(data), "\n")

	for index := 0; index < len(masses)-1; index++ {
		intValue, err := strconv.Atoi(masses[index])
		if err != nil {
			fmt.Println("Integer parsing error", err)
			return
		}

		requiredFuel += calculateFuel(intValue)
	}

	fmt.Println(fmt.Sprintf("%.2f", requiredFuel))
}
