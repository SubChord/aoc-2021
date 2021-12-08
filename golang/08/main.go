package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	lines := readLines("inp")
	part1(lines)
	part2(lines)

}

func part2(lines []string) {
	sum := 0
	for _, line := range lines {
		strToNum := map[string]int{}
		numToStr := map[int]string{}

		split := strings.Split(line, " | ")
		g1 := strings.Split(split[0], " ")
		g2 := strings.Split(split[1], " ")

		for _, s := range g1 {
			r := []rune(s)
			sort.Slice(r, func(i, j int) bool {
				return r[i] < r[j]
			})
			switch len(s) {
			case 2:
				strToNum[string(r)] = 1
				numToStr[1] = string(r)
			case 3:
				strToNum[string(r)] = 7
				numToStr[7] = string(r)
			case 4:
				strToNum[string(r)] = 4
				numToStr[4] = string(r)
			case 7:
				strToNum[string(r)] = 8
				numToStr[8] = string(r)
			}
		}

		dashCountMap := map[int][]string{}
		for _, s := range g1 {
			r := []rune(s)
			sort.Slice(r, func(i, j int) bool {
				return r[i] < r[j]
			})
			dashCountMap[len(s)] = append(dashCountMap[len(s)], string(r))
		}

		// find 0, 9 and 6
		for _, s := range dashCountMap[6] {
			if containsAll(s, numToStr[4]) {
				strToNum[s] = 9
				numToStr[9] = s
			} else if containsAll(s, numToStr[7]) && !containsAll(s, numToStr[4]) {
				strToNum[s] = 0
				numToStr[0] = s
			} else {
				strToNum[s] = 6
				numToStr[6] = s
			}
		}

		// so far we have 1, 7, 4, 8, 6, 0, 9
		// now we need to find the rest

		// find 2, 5 and 3
		for _, s := range dashCountMap[5] {
			if containsAll(s, numToStr[1]) {
				strToNum[s] = 3
				numToStr[3] = s
			} else if containsAll(numToStr[9], s) {
				strToNum[s] = 5
				numToStr[5] = s
			} else {
				strToNum[s] = 2
				numToStr[2] = s
			}
		}

		build := ""
		for _, s := range g2 {
			r := []rune(s)
			sort.Slice(r, func(i, j int) bool {
				return r[i] < r[j]
			})

			build += strconv.Itoa(strToNum[string(r)])
		}

		atoi, _ := strconv.Atoi(build)
		sum += atoi
	}

	fmt.Printf("Part 2: %d\n", sum)
}

func containsAll(s1, s2 string) bool {
	for _, c := range s2 {
		if !strings.Contains(s1, string(c)) {
			return false
		}
	}
	return true
}

func part1(lines []string) {
	occurences := map[int]int{}
	for _, line := range lines {
		split := strings.Split(line, " | ")
		g2 := strings.Split(split[1], " ")
		for _, s := range g2 {
			occurences[len(s)]++
		}
	}

	fmt.Printf("part 1: %d\n", occurences[2]+occurences[4]+occurences[3]+occurences[7])
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
