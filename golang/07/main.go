package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	lines := readLines("inp")
	numbers := []int{}

	for _, s := range strings.Split(lines[0], ",") {
		i, _ := strconv.Atoi(s)
		numbers = append(numbers, i)
	}

	max := 0
	min := 0
	for _, number := range numbers {
		if number > max {
			max = number
		}
		if number < min {
			min = number
		}
	}

	minFuelCost := math.MaxInt
	for i := min; i <= max; i++ {
		fuelCost := 0
		for _, number := range numbers {
			fuelCost += int(math.Abs(float64(number - i)))
		}
		if fuelCost < minFuelCost {
			minFuelCost = fuelCost
		}
	}

	fmt.Printf("Part 1: %d\n", minFuelCost)

	stepCost := map[int]int{}
	maxStep := max
	if int(math.Abs(float64(min))) > maxStep {
		maxStep = int(math.Abs(float64(min)))
	}

	for i := 1; i <= maxStep; i++ {
		stepCost[i] = stepCost[i-1] + i
	}

	minFuelCost = math.MaxInt
	for i := min; i <= max; i++ {
		fuelCost := 0
		for _, number := range numbers {
			fuelCost += stepCost[int(math.Abs(float64(number - i)))]
		}
		if fuelCost < minFuelCost {
			minFuelCost = fuelCost
		}
	}

	fmt.Printf("Part 2: %d\n", minFuelCost)

}

func readLines(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}
