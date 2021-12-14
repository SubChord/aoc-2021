package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"strings"
)

func main() {
	lines := readLines("inp")

	init := lines[0]
	pairs := make(map[string]string)
	for _, line := range lines[2:] {
		split := strings.Split(line, " -> ")
		pairs[split[0]] = split[1]
	}

	pt1 := pt1(init, pairs, 10)
	log.Printf("Part 1: %d", pt1)

	pt2 := pt2(init, pairs, 40)
	log.Printf("Part 2: %d", pt2)
}

func pt2(init string, pairs map[string]string, n int) int {
	pairCounter := map[string]int{}
	for i := 1; i < len(init); i++ {
		pairCounter[init[i-1:i+1]]++
	}

	letterCounter := map[string]int{}
	for i := 0; i < len(init); i++ {
		letterCounter[string(init[i])]++
	}

	for i := 0; i < n; i++ {
		pcCopy := map[string]int{}
		for k, v := range pairCounter {
			pcCopy[k] = v
		}

		for pair, n := range pairCounter {
			swap := pairs[pair]
			letterCounter[swap] += n

			pcCopy[string(pair[0])+string(swap)] += n
			pcCopy[string(swap)+string(pair[1])] += n

			pcCopy[pair] -= n
		}

		pairCounter = pcCopy
	}

	maxLetter := 0
	minLetter := math.MaxInt64
	for _, v := range letterCounter {
		if v > maxLetter {
			maxLetter = v
		}
		if v < minLetter {
			minLetter = v
		}
	}

	return maxLetter - minLetter
}

func pt1(init string, pairs map[string]string, n int) int {
	s := init
	for i := 0; i < n; i++ {
		s = swap(s, pairs)
	}

	countChars := make(map[rune]int)
	for _, c := range s {
		countChars[c]++
	}

	maxChars := 0
	minChars := math.MaxInt
	for _, c := range countChars {
		if c > maxChars {
			maxChars = c
		}
		if c < minChars {
			minChars = c
		}
	}

	return maxChars - minChars
}

func swap(init string, pairs map[string]string) string {
	strb := strings.Builder{}
	for i := 1; i < len(init); i++ {
		strb.WriteString(string(init[i-1]))
		s := pairs[init[i-1:i+1]]
		strb.WriteString(s)
	}
	strb.WriteString(string(init[len(init)-1]))

	return strb.String()
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
