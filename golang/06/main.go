package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	lines := readLines("inp")
	fish := []int{}

	for _, s := range strings.Split(lines[0], ",") {
		i, _ := strconv.Atoi(s)
		fish = append(fish, i)
	}

	for i := 0; i < 80; i++ {

		for i, v := range fish {
			if v == 0 {
				fish = append(fish, 8)
				fish[i] = 6
			} else {
				fish[i]--
			}
		}

	}

	fmt.Printf("part 1: %d\n", len(fish))

	// reset fish
	fish = []int{}
	for _, s := range strings.Split(lines[0], ",") {
		i, _ := strconv.Atoi(s)
		fish = append(fish, i)
	}

	ageMap := map[int]int{}
	for _, f := range fish {
		ageMap[f]++
	}

	zeros := ageMap[0]
	for i := 0; i < 256; i++ {
		newMap := map[int]int{}
		for i := 1; i < 9; i++ {
			newMap[i-1] = ageMap[i]
		}

		newMap[6] += zeros
		newMap[8] = zeros
		zeros = newMap[0]
		ageMap = newMap
	}

	c := 0
	for _, v := range ageMap {
		c += v
	}
	fmt.Printf("part 2: %d\n", c)
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
