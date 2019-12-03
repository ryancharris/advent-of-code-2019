package main

import (
	"fmt"
	"io/ioutil"
	"math"
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

func calculateFuel(mass int) float64 {
	return math.Floor(float64(mass/3)) - 2
}

func getModuleFuelRequirement(moduleMass int) float64 {
	counter := float64(moduleMass)
	var fuelRequirements = []float64{}

	// Calculate the fuel requirement array for this module until
	// we reach a 0 or negative value.
	for {
		fuel := calculateFuel(int(counter))
		counter = fuel

		if counter <= 0 {
			break
		}

		fuelRequirements = append(fuelRequirements, fuel)
	}

	// Sum the array of fuel requirements for this given module
	var sum float64
	for _, req := range fuelRequirements {
		sum += req
	}

	return sum
}

func main() {
	moduleRequirementsArray := []float64{}
	moduleMasses := getMassesList()

	// Iterate through the list of module masses and create an array of fuel
	// requirements for each module with the weight of the fuel taken into account.
	for i := 0; i < len(moduleMasses); i++ {
		moduleFuelNeeded := getModuleFuelRequirement(moduleMasses[i])
		moduleRequirementsArray = append(moduleRequirementsArray, moduleFuelNeeded)
	}

	// Sum the array of fuel requirements by module
	var totalFuel float64
	for _, moduleFuel := range moduleRequirementsArray {
		totalFuel += moduleFuel
	}

	fmt.Println("Total fuel needed:")
	fmt.Println(fmt.Sprintf("%.2f", totalFuel))
}
