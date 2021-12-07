package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	lines := readLines("inp")
	coords := parseLines(lines)

	part1(coords)
	part2(coords)
}

func part1(coords [][]Coord) {
	seen := map[string]int{}
	for _, coords := range coords {
		if coords[0].x != coords[1].x && coords[0].y != coords[1].y {
			continue
		}

		if coords[0].x != coords[1].x {
			fromx, tox := coords[0].x, coords[1].x
			if fromx > tox {
				fromx, tox = tox, fromx
			}
			for x := fromx; x <= tox; x++ {
				seen[Coord{x, coords[0].y}.ID()]++
			}
		} else {
			fromy, toy := coords[0].y, coords[1].y
			if fromy > toy {
				fromy, toy = toy, fromy
			}
			for y := fromy; y <= toy; y++ {
				seen[Coord{coords[0].x, y}.ID()]++
			}
		}
	}

	c := 0
	for _, v := range seen {
		if v >= 2 {
			c++
		}
	}

	log.Printf("Part 1: %d", c)
}

func part2(coords [][]Coord) {
	seen := map[string]int{}
	for _, coords := range coords {
		if coords[0].x == coords[1].x || coords[0].y == coords[1].y {
			if coords[0].x != coords[1].x {
				fromx, tox := coords[0].x, coords[1].x
				if fromx > tox {
					fromx, tox = tox, fromx
				}
				for x := fromx; x <= tox; x++ {
					seen[Coord{x, coords[0].y}.ID()]++
				}
			} else {
				fromy, toy := coords[0].y, coords[1].y
				if fromy > toy {
					fromy, toy = toy, fromy
				}
				for y := fromy; y <= toy; y++ {
					seen[Coord{coords[0].x, y}.ID()]++
				}
			}
			continue
		}

		// handle diagonals
		fromx, tox := coords[0].x, coords[1].x
		fromy, toy := coords[0].y, coords[1].y

		if fromx < tox {
			if fromy < toy {
				for x := fromx; x <= tox; x++ {
					seen[Coord{x, fromy}.ID()]++
					fromy++
				}
			} else {
				for x := fromx; x <= tox; x++ {
					seen[Coord{x, fromy}.ID()]++
					fromy--
				}
			}
		}else{
			if fromy < toy {
				for x := fromx; x >= tox; x-- {
					seen[Coord{x, fromy}.ID()]++
					fromy++
				}
			} else {
				for x := fromx; x >= tox; x-- {
					seen[Coord{x, fromy}.ID()]++
					fromy--
				}
			}
		}
	}

	c := 0
	for _, v := range seen {
		if v >= 2 {
			c++
		}
	}

	log.Printf("Part 2: %d", c)
}

func parseLines(lines []string) [][]Coord {
	coords := make([][]Coord, 0, len(lines))
	for _, line := range lines {
		c := []Coord{}
		split := strings.Split(line, " -> ")
		s0 := strings.Split(split[0], ",")
		x1, _ := strconv.Atoi(s0[0])
		y1, _ := strconv.Atoi(s0[1])

		s1 := strings.Split(split[1], ",")
		x2, _ := strconv.Atoi(s1[0])
		y2, _ := strconv.Atoi(s1[1])

		c = append(c, Coord{x1, y1})
		c = append(c, Coord{x2, y2})
		coords = append(coords, c)
	}
	return coords
}

type Coord struct {
	x, y int
}

func (c Coord) ID() string {
	return "x" + strconv.Itoa(c.x) + "y" + strconv.Itoa(c.y)
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
