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
	g2 := grid{}

	for _, line := range lines {
		ints := []int{}
		ints2 := []int{}
		split := strings.Split(line, "")
		for _, s := range split {
			i, err := strconv.Atoi(s)
			if err != nil {
				log.Fatal(err)
			}
			ints = append(ints, i)
			ints2 = append(ints2, i)
		}
		g = append(g, ints)
		g2 = append(g2, ints2)
	}

	part1(g)
	part2(g2)
}

func part1(g grid) {
	s := 0
	for i := 0; i < 100; i++ {
		s += g.step()
	}
	fmt.Println(s)
}

func part2(g2 grid) {
	size := len(g2) * len(g2[0])
	step := 0
	for {
		step++
		flashes := g2.step()
		if flashes == size {
			break
		}
	}

	fmt.Println(step)
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
	firstLoop := true
	flashInLastLoop := false
	for flashInLastLoop || firstLoop {
		firstLoop = false
		flashInLastLoop = false
		for i := range g {
			for j, _ := range g[i] {
				key := fmt.Sprintf("i%vj%v", i, j)
				if g[i][j] > 9 && !flashed[key] {
					flashInLastLoop = true
					flashCount++
					// increment adjacent cells including diagonals
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
					flashed[key] = true
				}
			}
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
