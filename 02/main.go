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

	part1(lines)
	part2(lines)
}

func part1(lines []string) {
	pt1x, pt1y := 0, 0
	for _, line := range lines {
		split := strings.Split(line, " ")
		v, _ := strconv.Atoi(split[1])
		switch split[0] {
		case "forward":
			pt1x += v
		case "up":
			pt1y -= v
		case "down":
			pt1y += v
		}
	}

	fmt.Printf("Part 1: %d\n", pt1x*pt1y)
}

func part2(lines []string) {
	x, y, aim := 0, 0, 0
	for _, line := range lines {
		split := strings.Split(line, " ")
		v, _ := strconv.Atoi(split[1])
		switch split[0] {
		case "forward":
			x += v
			y += aim * v
		case "up":
			aim -= v
		case "down":
			aim += v
		}
	}

	fmt.Printf("Part 2: %d\n", x*y)
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
