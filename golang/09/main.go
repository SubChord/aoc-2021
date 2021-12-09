package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	lines := readLines("inp")
	grid := [][]int{}
	for _, line := range lines {
		row := []int{}
		for _, c := range line {
			atoi, _ := strconv.Atoi(string(c))
			row = append(row, atoi)
		}
		grid = append(grid, row)
	}

	pt1(grid)
	pt2(grid)
}

func pt2(grid [][]int) {
	seen := make(map[string]bool)
	basins := []int{}
	for y, row := range grid {
		for x, v := range row {
			if v < 9 && !seen[fmt.Sprintf("%d,%d", x, y)] {
				basins = append(basins, crawl(grid, x, y, seen))
			}
		}
	}

	sort.Ints(basins)
	fmt.Printf("part 2: %d\n", basins[len(basins)-1] * basins[len(basins)-2] * basins[len(basins)-3])
}

func crawl(grid [][]int, x, y int, seen map[string]bool) int {
	if seen[fmt.Sprintf("x%dy%d", x, y)] {
		return 0
	}

	seen[fmt.Sprintf("x%dy%d", x, y)] = true

	if x < 0 || y < 0 || x >= len(grid[0]) || y >= len(grid) {
		return 0
	}

	if grid[y][x] == 9 {
		return 0
	}

	return 1 + crawl(grid, x+1, y, seen) + crawl(grid, x, y+1, seen) + crawl(grid, x-1, y, seen) + crawl(grid, x, y-1, seen)
}

func pt1(grid [][]int) {
	s := 0
	for y := 0; y < len(grid); y++ {
		row := grid[y]
	outer:
		for x := 0; x < len(row); x++ {
			curr := row[x]
			adjacents := []int{}
			if x > 0 {
				adjacents = append(adjacents, row[x-1])
			}
			if x < len(row)-1 {
				adjacents = append(adjacents, row[x+1])
			}
			if y > 0 {
				adjacents = append(adjacents, grid[y-1][x])
			}
			if y < len(grid)-1 {
				adjacents = append(adjacents, grid[y+1][x])
			}

			for _, adjacent := range adjacents {
				if curr >= adjacent {
					continue outer
				}
			}

			s += curr + 1
		}
	}

	fmt.Printf("Part 1: %d\n", s)
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
