package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	lines := readLines("inp")
	numbers := parseNumbers(lines)

	pt1 := 0
	for i := 1; i < len(numbers); i++ {
		if numbers[i-1] < numbers[i] {
			pt1++
		}
	}

	fmt.Printf("Part 1: %d\n", pt1)

	pt2 := 0
	for i := 3; i < len(numbers); i++ {
		s3 := numbers[i-3] + numbers[i-2] + numbers[i-1]
		s2 := numbers[i-2] + numbers[i-1] + numbers[i]
		if s2 > s3 {
			pt2++
		}
    }

	fmt.Printf("Part 2: %d\n", pt2)
}

func parseNumbers(lines []string) []int {
	numbers := make([]int,0, len(lines))
    for _, line := range lines {
		atoi, _ := strconv.Atoi(line)
		numbers = append(numbers, atoi)
    }
    return numbers
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
