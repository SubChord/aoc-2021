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
	g := grid{}

	for _, line := range lines {
		ints := []int{}
		split := strings.Split(line, "")
		for _, s := range split {
			i, err := strconv.Atoi(s)
			if err != nil {
				log.Fatal(err)
			}
			ints = append(ints, i)
		}
		g = append(g, ints)
	}

	part1(g)
}

func part1(g grid) {
	s := 0
	for i := 0; i < 100; i++ {
		s += g.step()
		g.print()
	}
	fmt.Println(s)
}

type grid [][]int

func (g grid) step() int {
	flashCount := 0
	for i, ints := range g {
		for j := range ints {
			g[i][j]++
		}
	}

	flashed := map[string]bool{}
	i := 0
	firstLoop := true
	somethingFlashed := true
	for somethingFlashed || firstLoop {
		somethingFlashed = false
		for j := range g[i] {
			if g[i][j] > 9 && !flashed[fmt.Sprintf("%d,%d", i, j)] {
				// increment adjacent cells including diagonals
				flashCount++
				if i > 0 {
					g[i-1][j]++
				}
				if i < len(g)-1 {
					g[i+1][j]++
				}
				if j > 0 {
					g[i][j-1]++
				}
				if j < len(g[i])-1 {
					g[i][j+1]++
				}
				if i > 0 && j > 0 {
					g[i-1][j-1]++
				}
				if i > 0 && j < len(g[i])-1 {
					g[i-1][j+1]++
				}
				if i < len(g)-1 && j > 0 {
					g[i+1][j-1]++
				}
				if i < len(g)-1 && j < len(g[i])-1 {
					g[i+1][j+1]++
				}
				flashed[fmt.Sprintf("i%vj%v", i, j)] = true
				somethingFlashed = true
			}
		}
		i = (i + 1) % len(g)
		if i == 0 {
			firstLoop = false
		}
	}

	for k, _ := range flashed {
		// set flashed cells to 0
		i, j := splitKeyToIJ(k)
		g[i][j] = 0
	}

	return flashCount
}

func splitKeyToIJ(k string) (int, int) {
	split := strings.Split(k, "j")
	i, err := strconv.Atoi(split[0][1:])
	if err != nil {
		log.Fatal(err)
	}
	j, err := strconv.Atoi(split[1])
	if err != nil {
		log.Fatal(err)
	}
	return i, j
}

func (g grid) print() {
	for _, ints := range g {
		for _, i := range ints {
			fmt.Printf("%d", i)
		}
		fmt.Println()
	}
	// print blank
	fmt.Println()
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
