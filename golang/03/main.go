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

	part1(lines)
	part2(lines)
}

type Pair struct {
	one  int
	zero int
}

func part1(lines []string) {
	m := map[int]Pair{}
	for _, line := range lines {
		for i, i2 := range line {
			if i2 == '1' {
				m[i] = Pair{m[i].one + 1, m[i].zero}
			} else {
				m[i] = Pair{m[i].one, m[i].zero + 1}
			}
		}
	}

	gamma := 0
	epsilon := 0
	for i := 0; i < len(m); i++ {
		v := m[i]

		if v.one > v.zero {
			gamma = gamma << 1
			gamma++
			epsilon = epsilon << 1
		} else {
			gamma = gamma << 1
			epsilon = epsilon << 1
			epsilon++
		}
	}

	fmt.Printf("Part 1: %d\n", gamma*epsilon)
}

func part2(lines []string) {
	oxyLines := make([]string, len(lines))
	co2lines := make([]string, len(lines))
	copy(oxyLines, lines)
	copy(co2lines, lines)

	calcM := func(lines []string) map[int]Pair {
		m := map[int]Pair{}
		for _, line := range lines {
			for i, i2 := range line {
				if i2 == '1' {
					m[i] = Pair{m[i].one + 1, m[i].zero}
				} else {
					m[i] = Pair{m[i].one, m[i].zero + 1}
				}
			}
		}
		return m
	}

	linesWithNumberInPosition := func(lines []string, v rune, position int) []string {
		var result []string
		for _, line := range lines {
			if rune(line[position]) == v {
				result = append(result, line)
			}
		}
		return result
	}

	m := calcM(oxyLines)

	idx := 0
	for len(oxyLines) > 1 {
		v := rune('0')
		if m[idx].one >= m[idx].zero {
			v = '1'
		}
		oxyLines = linesWithNumberInPosition(oxyLines, rune(v), idx)
		m = calcM(oxyLines)
		idx++
	}

	m = calcM(co2lines)

	idx = 0
	for len(co2lines) > 1 {
		v := rune('0')
		if m[idx].one < m[idx].zero {
			v = '1'
		}
		co2lines = linesWithNumberInPosition(co2lines, rune(v), idx)
		m = calcM(co2lines)
		idx++
	}

	oxy, _ := strconv.ParseInt(oxyLines[0], 2, 64)
	co2, _ := strconv.ParseInt(co2lines[0], 2, 64)

	fmt.Printf("Part 2: %d\n", oxy*co2)
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
