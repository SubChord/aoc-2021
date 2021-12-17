package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// target area: x=94..151, y=-156..-103
const targetX1 = 94
const targetX2 = 151
const targetY1 = -156
const targetY2 = -103

// target area: x=20..30, y=-10..-5
//const targetX1 = 20
//const targetX2 = 30
//const targetY1 = -10
//const targetY2 = -5

type pair struct {
	x, y int
}

func main() {
	pairsThatHitTarget := map[pair]int{}
	for x := 0; x <= targetX2; x++ {
		for y := 1000; y >= targetY1; y-- {
			posX := 0
			posY := 0
			velX := x
			velY := y
			maxYPos := posX

			for {
				posX, posY, velX, velY = step(posX, posY, velX, velY)

				if posY > maxYPos {
					maxYPos = posY
				}

				if inTarget(posX, posY) {
					pairsThatHitTarget[pair{x, y}] = maxYPos
				}

				if posX > targetX2 || posY < targetY1 {
					break
				}
			}
		}
	}

	maxY := 0
	for _, y := range pairsThatHitTarget {
		if y > maxY {
			maxY = y
		}
	}

	fmt.Printf("Part 1: %d\n", maxY)
	fmt.Printf("Part 2: %d\n", len(pairsThatHitTarget))
}

func inTarget(x int, y int) bool {
	if x >= targetX1 && x <= targetX2 && y >= targetY1 && y <= targetY2 {
		return true
	}
	return false
}

func step(posX, posY, velX, velY int) (int, int, int, int) {
	newX := posX + velX
	newY := posY + velY

	if velX < 0 {
		velX++
	}
	if velX > 0 {
		velX--
	}

	return newX, newY, velX, velY - 1
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
