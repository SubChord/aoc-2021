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

	instructions := []instruction{}
	points := []point{}

	foldLines := false
	for _, line := range lines {
		if foldLines {
			split := strings.Split(line, " ")
			inst := strings.Split(split[len(split)-1], "=")
			atoi, _ := strconv.Atoi(inst[1])
			instructions = append(instructions, instruction{
				op: inst[0],
				v:  atoi,
			})
		}

		if line == "" {
			foldLines = true
		}

		if foldLines {
			continue
		}

		split := strings.Split(line, ",")
		x, _ := strconv.Atoi(split[0])
		y, _ := strconv.Atoi(split[1])
		points = append(points, point{x, y})
	}

	pts := fold(points, instructions[0])
	fmt.Printf("Part 1: %d\n", len(pts))

	for _, inst := range instructions[1:] {
		pts = fold(pts, inst)
	}

	printPoints(pts)
}

func fold(points []point, inst instruction) []point {
	m := map[point]bool{}
	if inst.op == "y" {
		for _, p := range points {
			if p.y < inst.v {
				m[p] = true
				continue
			}

			m[point{p.x, inst.v*2 - p.y}] = true
		}
	} else {
		for _, p := range points {
			if p.x < inst.v {
				m[p] = true
				continue
			}

			m[point{inst.v*2 - p.x, p.y}] = true
		}
	}

	ret := []point{}
	for p := range m {
		ret = append(ret, p)
	}

	return ret
}

func printPoints(points []point) {
	sort.Slice(points, func(i, j int) bool {
		return points[i].y < points[j].y
	})

	maxX := 0
	maxY := 0
	for _, p := range points {
		if p.x > maxX {
			maxX = p.x
		}
		if p.y > maxY {
			maxY = p.y
		}
	}

	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			found := false
			for _, p := range points {
				if p.x == x && p.y == y {
					found = true
					break
				}
			}

			if found {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}

		fmt.Println()
	}
}

type instruction struct {
	op string
	v  int
}

type point struct {
	x int
	y int
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
